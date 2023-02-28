package algorithm

import (
	"math"

	"github.com/sozyGithub/project/tucil_2/src/point"
)

func DoBruteForce(points []point.Point, totalOpt *int) (float64, [][]point.Point) {
	var shortestDistance float64 = math.MaxFloat64
	var sdArray [][]point.Point
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			*totalOpt++
			distance := point.CalculateDistance(points[i], points[j])
			if distance < shortestDistance {
				shortestDistance = distance
				sdArray = nil
				sdArray = append(sdArray, []point.Point{points[i], points[j]})
			} else if point.IsEqual(distance, shortestDistance) {
				sdArray = append(sdArray, []point.Point{points[i], points[j]})
			}
		}
	}

	return shortestDistance, sdArray
}
