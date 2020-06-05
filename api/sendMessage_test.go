package api

<<<<<<< HEAD
import (
	"testing"
)
=======
import "testing"
>>>>>>> f42b3f3afd86bec62aa8bc6a094df7066142ab24

func TestClient_SendMessage(t *testing.T) {
	type fields struct {
		token string
	}
	type args struct {
		m Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Raw message",
<<<<<<< HEAD
			fields{testToken},
			args{Message{
				ChatID: testSendToID,
=======
			fields{"1285255270:SomeCorrectToken"},
			args{Message{
				ChatID: 123456789,
>>>>>>> f42b3f3afd86bec62aa8bc6a094df7066142ab24
				Text:   "Hello, it's a raw text.",
			}},
			false,
		},
		{
			"Message with markup",
<<<<<<< HEAD
			fields{testToken},
			args{Message{
				ChatID: testSendToID,
=======
			fields{"1285255270:SomeCorrectToken"},
			args{Message{
				ChatID: 123456789,
>>>>>>> f42b3f3afd86bec62aa8bc6a094df7066142ab24
				Text:   "Hello, it's text with inline markup.",
				InlineMarkup: InlineKeyboardMarkup{
					[][]InlineKeyboardButton{
						{
							InlineKeyboardButton{
								Text:     "Row 1",
								Callback: "Callback 1",
							},
						},
						{
							InlineKeyboardButton{
								Text: "Row 2",
								URL:  "google.com",
							},
						},
					},
				},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				token: tt.fields.token,
			}
			if err := c.SendMessage(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
