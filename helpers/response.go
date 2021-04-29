package helpers

type Response struct {
	Message string `json:"msg"`
}

// return example { Message: "Error 1062: Duplicate entry 'bro@gmail.com' for key 'users.email'"}
func Res(msg string) *Response {
	return &Response{Message: msg}
}
