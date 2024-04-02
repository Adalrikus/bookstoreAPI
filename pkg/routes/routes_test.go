package routes

import (
	"net/http"
	"testing"

	"github.com/Adalrikus/bookstoreAPI/pkg/models"
)

func TestRegisterBookStoreRoutes(t *testing.T) {
	type args struct {
		router  *http.ServeMux
		queries *models.Queries
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterBookStoreRoutes(tt.args.router, tt.args.queries)
		})
	}
}
