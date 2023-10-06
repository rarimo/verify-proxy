package data

type MasterQ interface {
	New() MasterQ

	VerifyRequestsQ() VerifyRequestsQ

	Transaction(func() error) error
}
