package api

import (
	"testing"
)

func TestClient_SetToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"NormalTest",
			args{testToken},
			false,
		},
		{
			"ErrorTest",
			args{"1285234123:SomeIncorrectToken"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Client
			if err := c.SetToken(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("SetToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
