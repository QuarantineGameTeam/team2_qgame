package main

import (
	"gopkg.in/fogleman/gg.v1"
)

func main() {
	dc := gg.NewContext(30, 30)
	dc.DrawCircle(15, 15, 15)
	dc.SetRGB(0, 200, 0)
	dc.Fill()

	dc.DrawPoint(14, 14, 14)
	dc.SetRGB(170, 190, 0)
	dc.Fill()

	dc.DrawPoint(13, 13, 13)
	dc.SetRGB(0, 200, 0)
	dc.Fill()

	dc.DrawPoint(12, 12, 12)
	dc.SetRGB(170, 190, 0)
	dc.Fill()

	dc.DrawPoint(11, 11, 11)
	dc.SetRGB(0, 200, 0)
	dc.Fill()

	dc.DrawPoint(10, 10, 10)
	dc.SetRGB(170, 190, 0)
	dc.Fill()

	dc.DrawPoint(9, 9, 9)
	dc.SetRGB(0, 200, 0)
	dc.Fill()

	dc.DrawPoint(8, 8, 8)
	dc.SetRGB(170, 190, 0)
	dc.Fill()

	dc.DrawPoint(7, 7, 7)
	dc.SetRGB(0, 200, 0)
	dc.Fill()

	dc.DrawPoint(6, 6, 6)
	dc.SetRGB(170, 190, 0)
	dc.Fill()

	dc.DrawPoint(5, 5, 5)
	dc.SetRGB(0, 200, 0)
	dc.Fill()

	dc.DrawPoint(4, 4, 4)
	dc.SetRGB(170, 190, 0)
	dc.Fill()

	dc.DrawPoint(3, 3, 3)
	dc.SetRGB(0, 200, 0)
	dc.Fill()

	dc.DrawPoint(2, 2, 2)
	dc.SetRGB(170, 190, 0)
	dc.Fill()

	dc.DrawPoint(1, 1, 1)
	dc.SetRGB(0, 200, 0)
	dc.Fill()

	dc.SavePNG("out.png")
}
