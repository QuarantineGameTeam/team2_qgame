package graph

import "testing"

func Test_scale(t *testing.T) {
	tests := []struct {
		name string
		v1 float64
		min1 float64
		min2 float64
		max1 float64
		max2 float64
		wantv2 float64
	} {
		{
		"positive", 2, 0, 0, 10, 100, 20,
	},
	{
		"negative", 3, 0, 0, 10, 100, 40,
	},
	}
	
	for _, test := range tests {
		t.Run (test.name, func(t *testing.T) {
			v2 := test.min2+(test.max2-test.min2)*((test.v1-test.min1)/(test.max1-test.min1))
			if v2 != test.wantv2 {
				t.Errorf("got value 2 %v, want value 2 %v", v2, test.wantv2)
			}
		})
	}
}

func Test_CreatePartViewPhoto(t *testing.T) {
	tests := []struct {
		name string
		locX, locY int
		drawingcenterX int
		drawingcenterY int
		drawingHorizon int
		wantonmap bool
	}{
		{
		"photo test", 5, 2, 4, 3, 1, true,
	}, 
	}

	var onmap bool

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.locX >= (test.drawingcenterX - test.drawingHorizon) && test.locX <= (test.drawingcenterX + test.drawingHorizon) {
				if test.locY >= (test.drawingcenterY - test.drawingHorizon) && test.locY <= (test.drawingcenterY + test.drawingHorizon) {
					onmap = true
				} else {
					onmap = false
				}
			} else {
				onmap = false
			}
			if onmap != test.wantonmap {
				t.Errorf("want shown %t, have %t", test.wantonmap, onmap)
			}
		})
	}
}
