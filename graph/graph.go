package graph

import (
	"image/color"
	"gopkg.in/fogleman/gg.v1"
	"team2_qgame/models"
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

//ScaleForPlayer converts players game coordinates into pixels for map.
func ScaleForPlayer(int, int) (float64, float64) {
	p := models.Player{}
	mappX := S/9*float64(p.X) //<- fancy double P cuz forgot that i had to do same for locations.
	mappY := S/9*float64(p.Y)
	return mappX, mappY
}

//ScaleForLocation - pretty much like ScaleForPlayer, but for locations. tbh much harder((
func ScaleForLocation(int, int) (float64, float64) {
	var (loc models.Location	
	 x, y int)

	m := models.Matrix{}
	loc = m.GetLocation(x, y)
	locX, locY := models.Location.GetLocation(loc)
	maplX := S/9*float64(locX)
	maplY := S/9*float64(locY)
	return maplX, maplY
	//WIP: somehow connect those variables x and y to actual coordinates. By far no clue about that.
}