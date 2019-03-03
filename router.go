package gopher

import (
	"strings"
)

type Router struct {
	defaultHandler Handler
	exactHandlers  map[string]Handler
	prefixHandlers map[string]Handler
}

func NewRouter(defaultHandler Handler) *Router {
	return &Router{defaultHandler, make(map[string]Handler), make(map[string]Handler)}
}

func (r *Router) Prefix(prefix string, handler Handler) {
	r.prefixHandlers[prefix] = handler
}

func (r *Router) Exact(selector string, handler Handler) {
	r.exactHandlers[selector] = handler
}

func (r *Router) Handle(resp Response, request Request) {
	handler := r.selectHandler(request.Selector)
	handler.Handle(resp, request)
}

func (r *Router) selectHandler(selector string) Handler {
	handler, ok := r.exactHandlers[selector]
	if ok {
		return handler
	}

	handler, ok = r.prefixHandlers[selector]
	if ok {
		return handler
	}

	var (
		bestLen     int = 0
		bestHandler Handler
	)

	for prefix, handler := range r.prefixHandlers {
		if strings.HasPrefix(selector, prefix) {
			if len(prefix) > bestLen {
				bestLen = len(prefix)
				bestHandler = handler
			}
		}
	}

	if 0 != bestLen {
		return bestHandler
	}

	return r.defaultHandler
}
