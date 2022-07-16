package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
	//Do nothing is the correct return type value
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is of type: %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
	//Do nothing is the correct return type value
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is of type: %T", v))
	}
}
