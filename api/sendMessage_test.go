package api

import "testing"

const token = "1285255270:SOME-CORRECT-TOKEN"
const sendToID = 123456789 // some correct user/chat id

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
			fields{token},
			args{Message{
				ChatID: sendToID,
				Text:   "Hello, it's a raw text.",
			}},
			false,
		},
		{
			"Message with markup",
			fields{token},
			args{Message{
				ChatID: sendToID,
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
		// TODO: Message with photo test
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
