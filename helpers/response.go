package helpers

type Response struct {
	Message string `json:"msg"`
}

// return example { Message: "Error 1062: Duplicate entry 'bro@gmail.com' for key 'users.email'"}
func Res(msg string) *Response {
	res := &Response{Message: msg}
	return res
}
