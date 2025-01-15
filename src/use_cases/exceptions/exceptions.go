package exceptions

type TaskAlreadyExistsError struct {
	Message string
}

func (t TaskAlreadyExistsError) Error() string {
	return t.Message
}
