package drawing

import (
	"testing"
)

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

func Test_drawBackground(t *testing.T) { // just as useless as aqua from konosuba(who watched knows what i mean :) )
	tests := []struct {
		name string
		x, y, w, h int
	} {
		{
		"uselessness lvl: aqua", 0, 0, 1023, 1024,
	}, 
	}

	var wantimage bool
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.w != 1024 {
				wantimage = false
			} else if test.h != 1024 {
				wantimage = false
			} else {
				wantimage = true
			} 
			if wantimage != true {
				t.Errorf("wanted image got successful true, got successful %t", wantimage)
			}
		})
	}

}

func Test_drawGrid(t *testing.T) {
	tests := []struct {
		name string
		dimension int
	} {
		{
			"pointless but ok", -3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.dimension < 0 {
				t.Errorf("got value %d negative, want %d positive", test.dimension, test.dimension)
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

func Test_CreateMapViewPhoto(t *testing.T) {
	tests := []struct {
		name string
		locX int 
		locY int
		visited bool
		dimension int
	} {
		{
			"ok but i prefer sth more useful", 4, -5, true, 9, 
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var wantval, validnum bool
			if test.dimension < 0 {
				validnum = false
			} else if test.locX < 0 || test.locX > test.dimension {
				validnum = false
			} else if test.locY < 0 || test.locY > test.dimension {
				validnum = false
			} else if test.visited != true {
				wantval = false
				validnum = true
			} else {
				wantval = true
				validnum = true
			}
			if validnum != true {
				t.Errorf("wanted locX and/or locY value is negative or out of bounds, have valid number %t", validnum)
			}
			if wantval != true {
				t.Errorf("wanted value visited true, have wanted value %t", wantval)
			}
		})
	}
}

func Test_CreateFullViewPhoto(t *testing.T) {
	tests := []struct {
		name string
		locX, locY int
		dimension int
	} {
		{
			"just useless", 3, 4, 9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var validnum bool
			if test.dimension < 0 {
				validnum = false
			} else if test.locX < 0 || test.locX > test.dimension {
				validnum = false
			} else if test.locY < 0 || test.locY > test.dimension {
				validnum = false
			} else {
				validnum = true
			}
			if validnum != true {
				t.Errorf("wanted values positive, have number vaildity %t due to one of values being negative or locX/Y being out of bounds", validnum)
			}
		})
	}
}
