package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	simpleMaps "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
	"github.com/xuri/excelize/v2"
	"image/color"
	"log"
)

func main() {
	job_data := getData("JobData.xlsx")
	for loc, job := range job_data {
		if loc < 1 {
			continue
		}
		fmt.Println(job[0], " is hiring in ", job[4])
	}
}

func getData(fileName string) [][]string {
	excelFile, err := excelize.OpenFile(fileName)
	defer excelFile.Close()
	if err != nil {
		log.Fatalln("couldn't open excel file:", err)
	}
	all_rows, err := excelFile.GetRows("JobsInfo")
	if err != nil {
		log.Fatal(err)
	}
	return all_rows
}

func showMap(data [][]string, loc *geo.Location) {
	context := simpleMaps.NewContext()
	context.SetSize(1200, 1200)
	context.SetZoom(6)
	context.AddObject(
		simpleMaps.NewMarker(
			s2.LatLngFromDegrees(loc.Lat, loc.Lng),
			color.RGBA{
				R: 0,
				G: 0xff,
				B: 0,
				A: 0xff,
			},
			16,
		),
	)
	img, err := context.Render()
	if err != nil {
		log.Fatalln(err)
	}
	if err := gg.SavePNG("DEmoJobsMap.png", img); err != nil {
		log.Fatalln(err)
	}
}
