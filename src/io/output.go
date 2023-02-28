package io

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/sozyGithub/project/tucil_2/src/point"
)

func WritePointsToCSV(points []point.Point) {
	dataset := make([][]string, len(points)+1)
	var header = []string{
		"X", "Y", "Z",
	}
	dataset[0] = append(dataset[0], header...)
	for i := 1; i <= len(points); i++ {
		for j := 1; j <= int(points[i-1].GetDimension()); j++ {
			dataset[i] = append(dataset[i], fmt.Sprintf("%v", points[i-1].GetCoor(j)))
		}
	}

	csvFile, err := os.Create("../src/plot/points.csv")
	if err != nil {
		fmt.Println("Failed to export points.", err)
	}

	csvWriter := csv.NewWriter(csvFile)
	for _, row := range dataset {
		_ = csvWriter.Write(row)
	}
	csvWriter.Flush()
	csvFile.Close()
}

func WritePPointsToCSV(sdArray [][]point.Point) {
	dataset := make([][]string, len(sdArray)*2+1)
	var header = []string{
		"X", "Y", "Z",
	}
	dataset[0] = append(dataset[0], header...)
	for i := 1; i <= len(sdArray); i++ {
		for j := 1; j <= int(sdArray[i-1][0].GetDimension()); j++ {
			dataset[i] = append(dataset[i], fmt.Sprintf("%v", sdArray[i-1][0].GetCoor(j)))
		}
		for j := 1; j <= int(sdArray[i-1][1].GetDimension()); j++ {
			dataset[i+1] = append(dataset[i+1], fmt.Sprintf("%v", sdArray[i-1][1].GetCoor(j)))
		}
	}

	csvFile, err := os.Create("../src/plot/ppoints.csv")
	if err != nil {
		fmt.Println("\n!!Failed to export ppoints.!!")
	}

	csvWriter := csv.NewWriter(csvFile)
	for _, row := range dataset {
		_ = csvWriter.Write(row)
	}
	csvWriter.Flush()
	csvFile.Close()
}
