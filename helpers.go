package gopher

type HandlerFunc struct {
	handler func(Response, Request)
}

func (h HandlerFunc) Handle(resp Response, req Request) {
	h.handler(resp, req)
}

func HandleFunc(handler func(Response, Request)) Handler {
	return HandlerFunc{handler}
}
