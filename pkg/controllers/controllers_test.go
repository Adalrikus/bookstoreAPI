package controllers

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Adalrikus/bookstoreAPI/pkg/models"
)

func TestCreateBook(t *testing.T) {
	type args struct {
		queries *models.Queries
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateBook(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBooks(t *testing.T) {
	type args struct {
		queries *models.Queries
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBooks(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	type args struct {
		queries *models.Queries
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBook(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	type args struct {
		queries *models.Queries
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateBook(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteBook(t *testing.T) {
	type args struct {
		queries *models.Queries
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteBook(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
