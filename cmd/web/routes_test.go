package main

import (
	"booking/internal/config"
	"fmt"
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//Do noting is the expected result
	default:
		t.Error(fmt.Sprintf("Type is not *chi.MUX, but is: %T", v))
	}
}
