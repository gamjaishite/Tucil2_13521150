package point

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

const float64EqualityTracehold = 1e-12

// Definition of point object
type Point struct {
	dimension  uint32
	coordinate []float64
}

// Creating a new point instance
func CreatePoint(coordinate []float64, dimension uint32) Point {
	return Point{
		coordinate: coordinate,
		dimension:  dimension,
	}
}

func (p *Point) GetCoors() string {
	var sCoor string = "("
	for i := 0; i < int(p.dimension); i++ {
		if i != 0 {
			sCoor += ","
		}
		sCoor += fmt.Sprintf("%0.3f", p.coordinate[i])
	}
	sCoor += ")"
	return sCoor
}

// Getter for coordinate
func (p *Point) GetCoor(index int) float64 {
	return p.coordinate[index-1]
}

// Getter for dimension
func (p *Point) GetDimension() uint32 {
	return p.dimension
}

// Generating randoms point
func GeneratePoints(N uint32, dimension uint32) []Point {
	var points []Point
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < int(N); i++ {
		coordinate := make([]float64, dimension)
		for j := 0; j < int(dimension); j++ {
			sign := rand.Intn(2)
			coordinate[j] = float64(rand.Float64() * 1000000 * math.Pow(-1, float64(sign)))
		}
		points = append(points, CreatePoint(coordinate, dimension))
	}

	pointsX := make([]Point, len(points))
	copy(pointsX, points)

	sort.SliceStable(pointsX, func(i, j int) bool {
		return pointsX[i].GetCoor(1) < pointsX[j].GetCoor(1)
	})

	return pointsX
}

// Calculate the Euclidian Distance between point p1 and p2
func CalculateDistance(p1, p2 Point) float64 {
	var sum float64 = 0
	for i := 1; i <= int(p1.GetDimension()); i++ {
		sum += math.Pow(p2.GetCoor(i)-p1.GetCoor(i), 2)
	}
	return math.Sqrt(sum)
}

func WritePointsToCSV(points []Point) {
	dataset := make([][]string, len(points)+1)
	var header = []string{
		"X", "Y", "Z",
	}
	dataset[0] = append(dataset[0], header...)
	for i := 1; i <= len(points); i++ {
		for j := 1; j <= int(points[i-1].dimension); j++ {
			dataset[i] = append(dataset[i], fmt.Sprintf("%v", points[i-1].GetCoor(j)))
		}
	}

	csvFile, err := os.Create("./plot/points.csv")
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

func WritePPointsToCSV(sdArray [][]Point) {
	dataset := make([][]string, len(sdArray)*2+1)
	var header = []string{
		"X", "Y", "Z",
	}
	dataset[0] = append(dataset[0], header...)
	for i := 1; i <= len(sdArray); i++ {
		for j := 1; j <= int(sdArray[i-1][0].dimension); j++ {
			dataset[i] = append(dataset[i], fmt.Sprintf("%v", sdArray[i-1][0].GetCoor(j)))
		}
		for j := 1; j <= int(sdArray[i-1][1].dimension); j++ {
			dataset[i+1] = append(dataset[i+1], fmt.Sprintf("%v", sdArray[i-1][1].GetCoor(j)))
		}
	}

	csvFile, err := os.Create("./plot/ppoints.csv")
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

func IsEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityTracehold
}

func PrintFormattedPoint(sdArray [][]Point) {
	for i := 0; i < len(sdArray); i++ {
		fmt.Printf("%*s %s %s\n", len(sdArray[i][0].GetCoors())+22, sdArray[i][0].GetCoors(), "<->", sdArray[i][1].GetCoors())
	}
}
