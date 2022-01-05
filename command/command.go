package command

type Validatable interface {
	Validate() error
}

type Command interface {
	Validatable

	Prepare()
}

func Process(cmd Command) error {
	cmd.Prepare()

	return Validate(cmd, "request")
}
