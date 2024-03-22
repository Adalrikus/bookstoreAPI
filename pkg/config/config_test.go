package config

import (
	"context"
	"reflect"
	"testing"

	"github.com/Adalrikus/bookstoreAPI/pkg/models"
	_ "github.com/mattn/go-sqlite3"
)

func TestConnect(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Queries
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
