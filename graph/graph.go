package graph

import (
	"image/color"
	"gopkg.in/fogleman/gg.v1"
)
//S equals width and heigth of map image. Made as float64 to use it in drawing funcs.
const S float64 = 1024

//GetMapGrid gets game map grid.
func GetMapGrid() {
	dc := gg.NewContext(1024, 1024) //<- cannot use S as mismatched type int and float64.
	dc.DrawRectangle(0, 0, S, S)
	dc.SetColor(color.White)
	dc.Fill()
	//creates plain white field.
	for i := 1; i < 9; i++ {
		//TO DO: connect grid with matrix and block dimensions. By far idk how to do it, so here is default 9x9 :p
		dc.DrawRectangle(0, S/9*float64(i), S, 2)
		dc.SetColor(color.Black)
		dc.Fill()
		dc.DrawRectangle(S/9*float64(i), 0, 2, S)
		dc.SetColor(color.Black)
		dc.Fill()
	}
	dc.SavePNG("gamemap_grid.png")
	//TO DO: get filepath.Abs("./gamemap_grid.png") and make tg bot send it as PhotoConfig or sth like that.
}
