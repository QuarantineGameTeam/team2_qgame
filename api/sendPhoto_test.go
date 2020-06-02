package api

import (
	"testing"
)

func TestClient_SendPhoto(t *testing.T) {
	type fields struct {
		token string
	}
	type args struct {
		chatID   int
		fileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Sending non-existing photo",
			fields{
				testToken,
			},
			args{
				testSendToID,
				"/some/wrong/path",
			},
			true,
		},
		{
			"Sending existing photo",
			fields{
				testToken,
			},
			args{
				testSendToID,
				"/some/correct/path",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				token: tt.fields.token,
			}
			if err := c.SendPhoto(tt.args.chatID, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("SendPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
