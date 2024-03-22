package controllers

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Adalrikus/bookstoreAPI/pkg/models"
)

func TestGetQueries(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want *models.Queries
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQueries(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseWriter(t *testing.T) {
	type args struct {
		w        http.ResponseWriter
		r        *http.Request
		status   int
		response []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResponseWriter(tt.args.w, tt.args.r, tt.args.status, tt.args.response)
		})
	}
}

func TestLogError(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogError(tt.args.w, tt.args.r, tt.args.err)
		})
	}
}

func TestCreateBook(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBook(tt.args.w, tt.args.r)
		})
	}
}

func TestGetBooks(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetBooks(tt.args.w, tt.args.r)
		})
	}
}

func TestGetBook(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetBook(tt.args.w, tt.args.r)
		})
	}
}

func TestUpdateBook(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateBook(tt.args.w, tt.args.r)
		})
	}
}

func TestDeleteBook(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteBook(tt.args.w, tt.args.r)
		})
	}
}
