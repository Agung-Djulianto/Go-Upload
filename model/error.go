package model

type MyError struct {
	Err string `json:"error"`
}

func (me MyError) Error() string {
	return me.Err
}

var (
	ErrorNotFound = MyError{
		Err: "Not Found!",
	}

	ErrorForbiddenAccess = MyError{
		Err: "Forbidden Access!",
	}
)
