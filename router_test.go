package gopher

import (
	"testing"
)

func TestRouter(t *testing.T) {
	r := NewRouter(&testHandler{"Default"})
	r.Prefix("/~jmhobbs", &testHandler{"Home"})
	r.Exact("/blog/search", &testHandler{"Blog Search"})
	r.Prefix("/blog", &testHandler{"Blog"})
	r.Prefix("/blog/archives", &testHandler{"Blog Archives"})

	checks := []struct {
		selector string
		expected string
	}{
		{
			"/~jmhobbs/sub/file",
			"Home",
		},
		{
			"/nothing/matching",
			"Default",
		},
		{
			"",
			"Default",
		},
		{
			"/blog/search",
			"Blog Search",
		},
		{
			"/blog/search/nested",
			"Blog",
		},
		{
			"/blog/2019/post",
			"Blog",
		},
		{
			"/blog/archives/deeper",
			"Blog Archives",
		},
	}

	for _, check := range checks {
		h := r.selectHandler(check.selector)
		if h.(*testHandler).Name != check.expected {
			t.Errorf("Expected to get %q handler, got %q", check.expected, h.(*testHandler).Name)
		}
	}
}

type testHandler struct {
	Name string
}

func (h *testHandler) Handle(resp Response, req Request) {}
