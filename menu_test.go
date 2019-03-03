package gopher

import "testing"

func TestLinkToString(t *testing.T) {
	link := Link{Type: TextFile, Display: "About", Selector: "/about", Hostname: "gopher.dev", Port: 70}
	str := link.String()
	expected := "0About\t/about\tgopher.dev\t70"
	if str != expected {
		t.Errorf("Link to string did not match expecations.\nExpect: %q\n   Got: %v", expected, str)
	}
}
