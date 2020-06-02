package api

import (
	"fmt"
	"testing"
)

// Uses a different testing method as the input&output are usually different
func TestClient_GetUpdates(t *testing.T) {
	firstUpdate := 0
	lastUpdate := 0

	client := &Client{
		"1285255270:",
	}

	// terminates after receiving 10 updates
	for lastUpdate-firstUpdate < 10{
		updates := client.GetUpdates(lastUpdate+1)
		if len(updates) != 0{
			fmt.Println(updates[0])

			if firstUpdate == 0{
				firstUpdate = updates[0].UpdateID
			}
			lastUpdate = updates[0].UpdateID
		}
	}
}