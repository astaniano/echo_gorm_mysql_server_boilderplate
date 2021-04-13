package helpers

type ErrRes struct {
	Message string `json:"message"`
}

// return example { Message: "Error 1062: Duplicate entry 'bro@gmail.com' for key 'users.email'"}
func Error_res(errMsg string) *ErrRes {
	res := &ErrRes{Message: errMsg}
	return res
}
