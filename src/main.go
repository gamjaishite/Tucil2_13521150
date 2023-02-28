package main

import (
	"fmt"
	"time"

	"github.com/klauspost/cpuid/v2"
	"github.com/matishsiao/goInfo"
	"github.com/pbnjay/memory"
	"github.com/sozyGithub/project/tucil_2/src/algorithm"
	"github.com/sozyGithub/project/tucil_2/src/io"
	"github.com/sozyGithub/project/tucil_2/src/point"
)

func main() {
	io.InitScreen()

	N, dimension := io.Input()
	pointsX := point.GeneratePoints(N, dimension)
	gi, _ := goInfo.GetInfo()

	fmt.Print("\n---\n\n")

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

	// Brute Force Data
	fmt.Println("==  BRUTE FORCE ==")
	fmt.Printf("%-20s: %f\n", "Shortest Distance", shortestDistanceBF)
	fmt.Printf("%-20s: \n", "Points")
	point.PrintFormattedPoint(sdArrayBF)
	fmt.Printf("%-20s: %d\n", "Total Operations", totalOptBF)
	fmt.Printf("%-20s: %v s\n", "Total Time", elapsedTimeBF)
	fmt.Print("\n---\n\n")

	// Divide and Conquer Data
	fmt.Println("==  DIVIDE AND CONQUER ==")
	fmt.Printf("%-20s: %f\n", "Shortest Distance", shortestDistanceDNC)
	fmt.Printf("%-20s: \n", "Points")
	point.PrintFormattedPoint(sdArrayDNC)
	fmt.Printf("%-20s: %d\n", "Total Operations", totalOptDNC)
	fmt.Printf("%-20s: %v s\n", "Total Time", elapsedTimeDNC)
	fmt.Print("\n---\n\n")

	// Computer Specifications
	fmt.Println("== COMPUTER SPECIFICATIONS ==")
	fmt.Printf("%-20s: %v\n", "Operating System", gi.OS)
	fmt.Printf("%-20s: %v GB\n", "Memory", float64(memory.TotalMemory())/float64(1000000000))
	fmt.Printf("%-20s: %s\n", "CPU Name", cpuid.CPU.BrandName)
	fmt.Printf("%-20s: %d\n", "Physical Cores", cpuid.CPU.PhysicalCores)
	fmt.Printf("%-20s: %d\n", "Logical Cores", cpuid.CPU.LogicalCores)
	if dimension == 3 {
		fmt.Println("\n---")
		fmt.Println("Visualisasi titik yang dihasilkan dapat dilihat dengan menjalankan perintah 'make run-visualizer' pada terminal")
		fmt.Println("---")
		io.WritePointsToCSV(pointsX)
		io.WritePPointsToCSV(sdArrayDNC)
	}
}
