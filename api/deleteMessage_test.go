package api

import (
	"testing"
)

func TestClient_DeleteMessage(t *testing.T) {
	c := &Client{
		token: testToken,
	}
	msg, err := c.SendMessage(Message {
		ChatID: testSendToID,
		Text: "test DeleteMessage",
	})
	if err != nil {
		t.Errorf("SendMessage error %v", err)
	}
	
	tests := []struct {
		name    string
		args   UpdateMessage
		wantErr bool
	}{
		{
			"Ok test",
			msg,
			false,
		},
		{
			"Not ok test",
			UpdateMessage{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				token: testToken,
			}
			if err := c.DeleteMessage(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}