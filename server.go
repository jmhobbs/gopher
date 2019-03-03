package gopher

import (
	"fmt"
	"net"
	"strings"
)

type Handler interface {
	Handle(Response, Request)
}

type Server struct{}

type Request struct {
	Selector   string
	Query      string
	RemoteAddr string
}

func (r Request) String() string {
	return fmt.Sprintf("Request{%q, %q, %s}", r.Selector, r.Query, r.RemoteAddr)
}

type Response struct {
	conn net.Conn
}

func (r *Response) Write(p []byte) (n int, err error) {
	return r.conn.Write(p)
}

func (r *Response) WriteMenu(m Menu) error {
	if err := m.Encode(r.conn); err != nil {
		return err
	}
	if _, err := r.Write([]byte("\r\n\r\n")); err != nil {
		return err
	}
	return nil
}

func (r *Response) End() error {
	_, err := r.Write([]byte("\r\n.\r\n"))
	return err
}

func (s *Server) ListenAndServe(address string, handler Handler) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go s.handleConnection(conn, handler)
	}
}

func (s *Server) handleConnection(conn net.Conn, handler Handler) {
	defer conn.Close()

	// TODO: Switch to a scanner?
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	req := strings.TrimSuffix(string(buf[:n]), "\r\n")
	parts := strings.Split(req, "\t")

	request := Request{
		Selector:   parts[0],
		RemoteAddr: conn.RemoteAddr().String(),
	}

	if len(parts) == 2 {
		request.Query = parts[1]
	}

	handler.Handle(Response{conn}, request)
}
