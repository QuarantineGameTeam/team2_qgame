package drawers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"gopkg.in/fogleman/gg.v1"
	"image"
	"image/color"
	"log"
	"os"
	"github.com/nfnt/resize"
)

const (
	windowConfig int = 1024 //didnt figure out how to get 2d const.

	//s represents float64 value of windowConfig.
	s float64 = 1024
	gridThickness float64 = 3

	//defaultDimension is a default size for a full map.
	defaultDimension int = 9 //probably will be unused later. idk atm.
)

//NOTE: these 3 funcs are UNEXPORTED(they will be used only here).

//Scale translates value v1 from one range(min1; max1) to another(min2; max2).
func scale(v1, min1, max1, min2, max2 float64) float64 {
	v2 := min2+(max2-min2)*((v1-min1)/(max1-min1))
	return v2
}

//DrawBackground draws a background with selected color.
func drawBackground(context *gg.Context, c color.Color) {
	context.DrawRectangle(0, 0, s, s)
	context.SetColor(c)
	context.Fill()
}

//DrawGrid draws a (dimension*dimension) grid with selected dimensions.
func drawGrid(context *gg.Context, dimension int) {
	for v := 1; v < dimension; v++ {
		context.DrawRectangle((s/float64(dimension))*float64(v), 0, gridThickness, s)
		context.SetColor(color.Black)
		context.Fill()
		context.DrawRectangle(0, (s/float64(dimension))*float64(v), s, gridThickness)
		context.SetColor(color.Black)
		context.Fill()
	}
}

//NOTE 2: these 3 funcs are EXPORTED(will be used in main.go).

//CreatePartViewPhoto draws a part-view map where objects are displayed on players horizon.
func CreatePartViewPhoto(locations []models.Location, players []models.Player, drawingCenterX, drawingCenterY, drawingHorizon int, saveTo string) error {
	context := gg.NewContext(windowConfig, windowConfig)
	drawBackground(context, color.White)
	horizon := 2*drawingHorizon+1
	drawGrid(context, horizon)

	objects := locations
	for i := 0; i < len(players); i++ {
		objects = append(objects, &players[i])
	}

	for _, l := range objects {
		locX, locY := l.GetLocation()
		f, err := os.Open(l.GetSmallPic())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}	//finally found a bit confusing but ok way to convert string to image.Image.
		crop := resize.Resize(uint(s)/uint(horizon), uint(s)/uint(horizon), img, resize.Lanczos3)
		if locX >= (drawingCenterX - drawingHorizon) && locX <= (drawingCenterX + drawingHorizon) {
			if locY >= (drawingCenterY - drawingHorizon) && locY <= (drawingCenterY + drawingHorizon) {
				if locX > drawingCenterX && locY > drawingCenterY {
					context.DrawImage(crop, int(s)*2/3, int(s)*2/3)
				} else if locX == drawingCenterX && locY > drawingCenterY {
					context.DrawImage(crop, int(s)/3, int(s)*2/3)
				} else if locX < drawingCenterX && locY > drawingCenterY {
					context.DrawImage(crop, 0, int(s)*2/3)
				} else if locX > drawingCenterX && locY == drawingCenterY {
					context.DrawImage(crop, int(s)*2/3, int(s)/3)
				} else if locX > drawingCenterX && locY < drawingCenterY {
					context.DrawImage(crop, int(s)*2/3, 0)
				} else if locX == drawingCenterX && locY == drawingCenterY {
					context.DrawImage(crop, int(s)/3, int(s)/3)
				} else if locX == drawingCenterX && locY < drawingCenterY {
					context.DrawImage(crop, int(s)/3, 0)
				} else if locX < drawingCenterX && locY == drawingCenterY {
					context.DrawImage(crop, 0, int(s)/3)
				} else if locX < drawingCenterX && locY < drawingCenterY {
					context.DrawImage(crop, 0, 0)
				}
			}
		}
		if drawingCenterX == defaultDimension {
			context.DrawRectangle(s*2/3, 0, s*2/3, s)
			context.SetRGB(255, 0, 0)
			context.Fill()
		} else if drawingCenterY == defaultDimension {
			context.DrawRectangle(0, s*2/3, s, s/3)
			context.SetRGB(255, 0, 0)
			context.Fill()
		} else if drawingCenterX == 0 {
			context.DrawRectangle(0, 0, s/3, s)
			context.SetRGB(255, 0, 0)
			context.Fill()
		} else if drawingCenterY == 0 {
			context.DrawRectangle(0, 0, s, s/3)
			context.SetRGB(255, 0, 0)
			context.Fill()
		} 
		if drawingCenterX == defaultDimension && drawingCenterY == defaultDimension {
			context.DrawRectangle(s*2/3, 0, s*2/3, s)
			context.SetRGB(255, 0, 0)
			context.Fill()
			context.DrawRectangle(0, s*2/3, s, s*2/3)
			context.SetRGB(255, 0, 0)
			context.Fill()
		} else if drawingCenterX == 0 && drawingCenterY == 0 {
			context.DrawRectangle(0, 0, s/3, s)
			context.SetRGB(255, 0, 0)
			context.Fill()
			context.DrawRectangle(0, 0, s, s/3)
			context.SetRGB(255, 0, 0)
			context.Fill()
		} else if drawingCenterX == defaultDimension && drawingCenterY == 0 {
			context.DrawRectangle(s*2/3, 0, s*2/3, s)
			context.SetRGB(255, 0, 0)
			context.Fill()
			context.DrawRectangle(0, 0, s, s/3)
			context.SetRGB(255, 0, 0)
			context.Fill()
		} else if drawingCenterX == 0 && drawingCenterY == defaultDimension {
			context.DrawRectangle(0, 0, s/3, s)
			context.SetRGB(255, 0, 0)
			context.Fill()
			context.DrawRectangle(0, s*2/3, s, s*2/3)
			context.SetRGB(255, 0, 0)
			context.Fill()
		}
	}
	return context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}

//CreateMapViewPhoto draws a full map but only areas that have been visited will be displayed.
func CreateMapViewPhoto(locations []models.Location, players []models.Player, visited[][]bool, saveTo string) error {
	context := gg.NewContext(windowConfig, windowConfig)
	drawBackground(context, color.White)
	drawGrid(context, defaultDimension)

	objects := locations
	for i := 0; i < len(players); i++ {
		objects = append(objects, &players[i])
	}

	for _, l := range objects {
		locX, locY := l.GetLocation()
		f, err := os.Open(l.GetSmallPic())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		crop := resize.Resize(uint(s)/9, uint(s)/9, img, resize.Lanczos3)

		if visited[locX][locY] == true {
			context.DrawImage(crop, int(scale(float64(locX), 0, float64(defaultDimension), 0, s)), int(scale(float64(locY), 0, float64(defaultDimension), 0, s)))
		}
	}
	return context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}

//CreateFullViewPhoto draws a full map with all the locations no matter if they are not visible. Only for admins.
func CreateFullViewPhoto(locations []models.Location, players []models.Player, saveTo string) error {
	context := gg.NewContext(windowConfig, windowConfig)
	drawBackground(context, color.White)
	drawGrid(context, defaultDimension)

	objects := locations
	for i := 0; i < len(players); i++ {
		objects = append(objects, &players[i])
	}

	for _, l := range objects {
		locX, locY := l.GetLocation()
		f, err := os.Open(l.GetSmallPic())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		crop := resize.Resize(uint(s)/9, uint(s)/9, img, resize.Lanczos3)

		context.DrawImage(crop, int(scale(float64(locX), 0, float64(defaultDimension), 0, s)), int(scale(float64(locY), 0, float64(defaultDimension), 0, s)))
	}

	return context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}