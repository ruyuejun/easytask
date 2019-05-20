package common

type Result struct {
	Code int
	Msg string
	Data []interface{}
}

type CommonI interface {
	Find() *Result
}
