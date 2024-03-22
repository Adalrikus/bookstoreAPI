// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package models

import (
	"context"
	_ "encoding/json"
	"reflect"
	"testing"
)

func TestQueries_CreateBook(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg CreateBookParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.CreateBook(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.CreateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_DeleteBook(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			if err := q.DeleteBook(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Queries.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueries_GetBook(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.GetBook(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_GetBooks(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.GetBooks(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_UpdateBook(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg UpdateBookParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			if err := q.UpdateBook(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("Queries.UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
