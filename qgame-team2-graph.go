package main

import (
	"gopkg.in/fogleman/gg.v1"
	"image/color"
)
//S - random const that represents our image
var S float64 = 1024 
func main() {
    dc := gg.NewContext(1024, 1024)
	dc.DrawRectangle(0, 0, S, S)
	dc.SetColor(color.White)
	dc.Fill()
	var v float64
	v = 1
	for i := 1; i < 5; i++ {
		dc.DrawRectangle(S/5 * v, 0, 2, S)
		dc.SetColor(color.Black)
		dc.Fill()
		dc.DrawRectangle(0, S/5 * v, S, 2)
		dc.SetColor(color.Black)
		dc.Fill()
		v++
	}
	dc.DrawRectangle(S/5*2, S/5*2, S/5, S/5)
	dc.SetColor(color.RGBA{255, 0, 0, 255})
	dc.Fill()

	if err := dc.LoadFontFace("/Library/Fonts/Minecraft/minecraft.ttf", 26); err != nil {
		panic(err)
	}
	dc.SetColor(color.White)
	dc.DrawStringAnchored("you are here", S/2, S/2, 0.5, 0.5)

	dc.SavePNG("out.png")
}