package schema

type Stream string

const (
	BackofficeStream = "backoffice"
)

type Subject string

const (
	RegistrationCustomerActions = "registration.customer.actions"

	OperationCustomerActions = "operation.customer.actions"

	BackofficeCustomerActions = "backoffice.customer.actions"
)
