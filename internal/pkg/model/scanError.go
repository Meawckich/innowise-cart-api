package model

type ScanError struct {
	Msg string `json:"message"`
}

func (e ScanError) Error() string {
	return e.Msg
}
