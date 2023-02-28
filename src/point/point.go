package point

import (
	"fmt"
	"math"
	"math/rand"
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

// Getting coordinates of a point in string format
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
			coordinate[j] = float64(rand.Float64() * 1000 * math.Pow(-1, float64(sign)))
		}
		points = append(points, CreatePoint(coordinate, dimension))
	}

	pointsX := make([]Point, len(points))
	copy(pointsX, points)

	sortedPoints := QuickSortPoints(pointsX)

	return sortedPoints
}

// Sorting []Point using Quick Sort Algorithm
func QuickSortPoints(points []Point) []Point {
	if len(points) < 2 {
		return points
	}

	left, right := 0, len(points)-1
	pivot := rand.Int() % len(points)

	points[pivot], points[right] = points[right], points[pivot]
	for i := range points {
		if points[i].GetCoor(1) < points[right].GetCoor(1) {
			points[left], points[i] = points[i], points[left]
			left++
		}
	}

	points[left], points[right] = points[right], points[left]
	QuickSortPoints(points[:left])
	QuickSortPoints(points[left+1:])

	return points
}

// Calculate the Euclidian Distance between point p1 and p2
func CalculateDistance(p1, p2 Point) float64 {
	var sum float64 = 0
	for i := 1; i <= int(p1.GetDimension()); i++ {
		sum += math.Pow(p2.GetCoor(i)-p1.GetCoor(i), 2)
	}
	return math.Sqrt(sum)
}

func IsEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityTracehold
}

func PrintFormattedPoint(sdArray [][]Point) {
	for i := 0; i < len(sdArray); i++ {
		fmt.Printf("%*s %s %s\n", len(sdArray[i][0].GetCoors())+22, sdArray[i][0].GetCoors(), "<->", sdArray[i][1].GetCoors())
	}
}
