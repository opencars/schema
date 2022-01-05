package command_test

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/opercars/schema"
	"github.com/opercars/schema/command"
)

type Example struct {
	UserID     string
	ExampleID  string
	Nested     NestedExample
	ExampleIDs []string
}

func (ex *Example) Validate() error {
	return validation.ValidateStruct(ex,
		validation.Field(&ex.UserID,
			validation.Required.Error("required"),
		),
		validation.Field(&ex.ExampleID,
			validation.Required.Error("required"),
		),
		validation.Field(&ex.Nested,
			validation.Required.Error("required"),
		),
		validation.Field(&ex.ExampleIDs,
			validation.Required.Error("required"),
		),
	)
}

type NestedExample struct {
	HelloWorld string
}

func (ex *NestedExample) Validate() error {
	return validation.ValidateStruct(ex,
		validation.Field(&ex.HelloWorld,
			validation.Required.Error("required"),
		),
	)
}

func TestToSnakeCase(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		{
			input:  "simple",
			output: "simple",
		},
		{
			input:  "JohnDoe",
			output: "john_doe",
		},
		{
			input:  "TestHelloWorld",
			output: "test_hello_world",
		},
	}

	for _, test := range tt {
		output := command.ToSnakeCase(test.input)
		assert.Equal(t, test.output, output)
	}
}

func TestValidate(t *testing.T) {
	example := Example{
		UserID:    "",
		ExampleID: "",
		Nested: NestedExample{
			HelloWorld: "",
		},
		ExampleIDs: []string{},
	}

	err := command.Validate(&example, "request")
	assert.Error(t, err)

	res, ok := err.(schema.ValidationError)
	require.True(t, ok)

	expected := map[string][]string{
		"request.example_id":  {"required"},
		"request.user_id":     {"required"},
		"request.example_ids": {"required"},
	}

	assert.Equal(t, expected, res.Messages)
}
