package schema

type Subject string

const (
	CustomerRegistrationActions Subject = "customer.registration.actions"
	OperationCustomerActions    Subject = "customer.operation.actions"
	VinDecodingCustomerActions  Subject = "customer.vin-decoding.actions"
	KoatuuCustomerActions       Subject = "customer.koatuu.actions"
	CoreCustomerActions         Subject = "customer.core.actions"

	StatisfyAuctionActions Subject = "statisfy.auction.actions"
)
