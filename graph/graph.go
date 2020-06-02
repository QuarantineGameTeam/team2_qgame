package graph

import (
	"fmt"
	"image"
	"image/color"
	"gopkg.in/fogleman/gg.v1"
)

const (
	windowConfig int = 1024 //didnt figure out how to get 2d const.

	//S represents float64 value of windowConfig.
	S float64 = 1024
	gridThickness float64 = 3

	//DefaultDimension is a default size for a full map.
	DefaultDimension int = 9 //probably will be unused later. idk atm.
)

//Location is coordinates of a location. probably i guess :p
type Location struct {
	X int
	Y int
	SmallPic image.Image
}

//NOTE: these 3 funcs are UNEXPORTED(they will be used only here).

//Scale translates value v1 from one range(min1; max1) to another(min2; max2).
func Scale(v1, min1, max1, min2, max2 float64) float64 {
	v2 := min2+(max2-min2)*((v1-min1)/(max1-min1))
	return v2
}

//DrawBackground draws a background with selected color.
func DrawBackground(context *gg.Context, c color.Color) {
	context.DrawRectangle(0, 0, S, S)
	context.SetColor(c)
	context.Fill()
}

//DrawGrid draws a (dimension*dimension) grid with selected dimensions.
func DrawGrid(context *gg.Context, dimension int) {
	for v := 1; v < dimension; v++ {
		context.DrawRectangle((S/float64(dimension))*float64(v), 0, gridThickness, S)
		context.SetColor(color.Black)
		context.Fill()
		context.DrawRectangle(0, (S/float64(dimension))*float64(v), S, gridThickness)
		context.SetColor(color.Black)
		context.Fill()
	}
}

//NOTE 2: these 3 funcs are EXPORTED(will be used in main.go).

//CreatePartViewPhoto draws a part-view map where objects are displayed on players horizon.
func CreatePartViewPhoto(locations []Location, drawingCenterX, drawingCenterY, drawingHorizon int, saveTo string) {
	context := gg.NewContext(windowConfig, windowConfig)
	DrawBackground(context, color.White)
	horizon := 2*drawingHorizon+1
	DrawGrid(context, horizon)
	for _, l := range(locations) {
		if l.X >= (drawingCenterX - drawingHorizon) && l.X <= (drawingCenterX + drawingHorizon) {
			if l.Y >= (drawingCenterY - drawingHorizon) && l.Y <= (drawingCenterY + drawingHorizon) {
				context.DrawImage(l.SmallPic, int(Scale(float64(l.X), 0, float64(horizon*3), 0, S)), int(Scale(float64(l.Y), 0, float64(horizon*3), 0, S)))
			}
		}
	}
	context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}

//CreateMapViewPhoto draws a full map but only areas that have been visited will be displayed.
func CreateMapViewPhoto(locations []Location, visited[][]bool, saveTo string) {
	context := gg.NewContext(windowConfig, windowConfig)
	DrawBackground(context, color.White)
	DrawGrid(context, DefaultDimension)
	for _, l := range(locations) {
		if visited[l.X][l.Y] == true {
			context.DrawImage(l.SmallPic, int(Scale(float64(l.X), 0, float64(DefaultDimension), 0, S)), int(Scale(float64(l.Y), 0, float64(DefaultDimension), 0, S)))
		}
	}
	context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}

//CreateFullViewPhoto draws a full map with all the locations no matter if they are not visible. Only for admins.
func CreateFullViewPhoto(locations []Location, saveTo string) {
	context := gg.NewContext(windowConfig, windowConfig)
	DrawBackground(context, color.White)
	DrawGrid(context, DefaultDimension)
	for _, l := range(locations) {
		context.DrawImage(l.SmallPic, int(Scale(float64(l.X), 0, float64(DefaultDimension), 0, S)), int(Scale(float64(l.Y), 0, float64(DefaultDimension), 0, S)))
	}
	context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}