package main

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	simpleMaps "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
	"github.com/xuri/excelize/v2"
	"image/color"
	"log"
)

var LocationCache map[string]*geo.Location
var JobCounts map[string]int

func main() {
	job_data := getData("JobData.xlsx")
	LocationCache = make(map[string]*geo.Location)
	JobCounts = make(map[string]int)
	univ_loc := findLocation("St. Joseph, MO")
	showMap(job_data, univ_loc)
}

func findLocation(place string) *geo.Location {
	geoLookup := openstreetmap.Geocoder()
	locationData, err := geoLookup.Geocode(place)
	if err != nil {
		log.Fatalln("Couldn't look up goecoding info:", err)
	}
	return locationData
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
	processJobs(data, loc)
	context := simpleMaps.NewContext()
	context.SetSize(1200, 1200)
	context.SetZoom(6)
	for city, numJobs := range JobCounts {
		cityLoc := LocationCache[city]
		context.AddObject(
			simpleMaps.NewMarker(
				s2.LatLngFromDegrees(cityLoc.Lat, cityLoc.Lng),
				getColor(numJobs),
				16,
			),
		)
	}
	context.SetCenter(s2.LatLngFromDegrees(loc.Lat, loc.Lng))
	img, err := context.Render()
	if err != nil {
		log.Fatalln(err)
	}
	if err := gg.SavePNG("DEmoJobsMap.png", img); err != nil {
		log.Fatalln(err)
	}
}

func getColor(numjobs int) color.Color {
	if numjobs > 50 {
		return color.RGBA{
			R: 0,
			G: 0xff,
			B: 0,
			A: 0xff,
		}
	}
	if numjobs > 20 {
		return color.RGBA{
			R: 0x11,
			G: 0xee,
			B: 0x11,
			A: 0xff,
		}
	}
	if numjobs > 10 {
		return color.RGBA{
			R: 0x11,
			G: 0xbb,
			B: 0xbb,
			A: 0xff,
		}
	}
	if numjobs > 3 {
		return color.RGBA{
			R: 0x11,
			G: 0x11,
			B: 0xff,
			A: 0xff,
		}
	}
	if numjobs >= 1 {
		return color.RGBA{
			R: 0xff,
			G: 0x11,
			B: 0x11,
			A: 0xff,
		}
	}
	return color.RGBA{}
}

func processJobs(job_data [][]string, defaultLoc *geo.Location) {
	for row_number, job := range job_data {
		if row_number < 1 {
			continue
		}
		cityName := job[4]
		_, ok := LocationCache[cityName]
		if ok {
			JobCounts[cityName] += 1
		} else {
			loc := findLocation(cityName)
			if loc == nil {
				loc = defaultLoc
			}
			LocationCache[cityName] = loc
			JobCounts[cityName] = 1
		}
	}
}
