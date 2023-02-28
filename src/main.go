package main

import (
	"fmt"
	"time"

	"github.com/sozyGithub/project/tucil_2/src/algorithm"
	"github.com/sozyGithub/project/tucil_2/src/io"
	"github.com/sozyGithub/project/tucil_2/src/point"
)

func main() {
	N, dimension := io.Input()
	pointsX := point.GeneratePoints(N, dimension)

	fmt.Println("---")

	// Brute Force
	startTimeBF := time.Now()
	totalOptBF := 0
	shortestDistanceBF, sdArrayBF := algorithm.DoBruteForce(pointsX, &totalOptBF)
	elapsedTimeBF := time.Since(startTimeBF).Seconds()

	// Divide and Conquer
	startTimeDNC := time.Now()
	totalOptDNC := 0
	shortestDistanceDNC, sdArrayDNC := algorithm.DoDNC(pointsX, &totalOptDNC)
	elapsedTimeDNC := time.Since(startTimeDNC).Seconds()

	fmt.Println("==  BRUTE FORCE ==")
	fmt.Printf("%-20s: %f\n", "Shortest Distance", shortestDistanceBF)
	fmt.Printf("%-20s: \n", "Points")
	point.PrintFormattedPoint(sdArrayBF)
	fmt.Printf("%-20s: %d\n", "Total Operations", totalOptBF)
	fmt.Printf("%-20s: %v s\n", "Total Time", elapsedTimeBF)
	fmt.Println("---")
	fmt.Println("==  DIVIDE AND CONQUER ==")
	fmt.Printf("%-20s: %f\n", "Shortest Distance", shortestDistanceDNC)
	fmt.Printf("%-20s: \n", "Points")
	point.PrintFormattedPoint(sdArrayDNC)
	fmt.Printf("%-20s: %d\n", "Total Operations", totalOptDNC)
	fmt.Printf("%-20s: %v s\n", "Total Time", elapsedTimeDNC)
	if dimension == 3 {
		point.WritePointsToCSV(pointsX)
		point.WritePPointsToCSV(sdArrayDNC)
	}
}
