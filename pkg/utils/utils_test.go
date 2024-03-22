package utlis

import (
	"net/http"
	"testing"
)

func TestParseBody(t *testing.T) {
	type args struct {
		r *http.Request
		x interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseBody(tt.args.r, tt.args.x); (err != nil) != tt.wantErr {
				t.Errorf("ParseBody() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
