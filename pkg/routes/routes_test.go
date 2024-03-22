package routes

import (
	"net/http"
	"testing"
)

func TestRegisterBookStoreRoutes(t *testing.T) {
	type args struct {
		router *http.ServeMux
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterBookStoreRoutes(tt.args.router)
		})
	}
}
