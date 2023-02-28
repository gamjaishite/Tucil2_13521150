package algorithm

import (
	"math"

	"github.com/sozyGithub/project/tucil_2/src/point"
)

// Divide and Conquer Algorithm implementation
func DoDNC(pointsX []point.Point, totalOpt *int) (float64, [][]point.Point) {
	n := len(pointsX)
	dimension := pointsX[0].GetDimension()

	var sPoint1 point.Point
	var sPoint2 point.Point
	var shortestDistance float64
	var sdArray = make([][]point.Point, 0)

	// Base cases
	if n == 2 {
		*totalOpt++
		var pPoint = []point.Point{pointsX[0], pointsX[1]}
		sdArray = append(sdArray, pPoint)
		return point.CalculateDistance(pointsX[0], pointsX[1]), sdArray
	}
	if n == 3 {
		var pPoint = []point.Point{sPoint1, sPoint2}
		sdArray = append(sdArray, pPoint)
		shortestDistance, sdArray = DoBruteForce(pointsX, totalOpt)
		return shortestDistance, sdArray
	}

	// Divide
	midPoint := pointsX[n/2]
	deltaLeft, sdArrayL := DoDNC(pointsX[0:n/2], totalOpt)
	deltaRight, sdArrayR := DoDNC(pointsX[n/2:n], totalOpt)
	if deltaLeft < deltaRight {
		shortestDistance = deltaLeft
		temp := make([][]point.Point, len(sdArrayL))
		copy(temp, sdArrayL)
		sdArray = append(sdArray, temp...)
	} else if deltaLeft > deltaRight {
		shortestDistance = deltaRight
		temp := make([][]point.Point, len(sdArrayR))
		copy(temp, sdArrayR)
		sdArray = append(sdArray, temp...)
	} else {
		sdArrayL = append(sdArrayL, sdArrayR...)
		sdArray = append(sdArray, sdArrayL...)
	}

	leftPoints := pointsX[0 : n/2]
	rightPoints := pointsX[n/2 : n]

	indxL := 0
	indxR := 0
	inLeftPoints := make([]point.Point, n/2+1)
	inRightPoints := make([]point.Point, n/2+1)
	for _, pointL := range leftPoints {
		if pointL.GetCoor(1) >= midPoint.GetCoor(1)-shortestDistance {
			inLeftPoints[indxL] = pointL
			indxL++
		}
	}
	for _, pointR := range rightPoints {
		if pointR.GetCoor(1) <= midPoint.GetCoor(1)+shortestDistance {
			inRightPoints[indxR] = pointR
			indxR++
		}
	}

	// Conquer
	for i := 0; i < indxL; i++ {
		for j := 0; j < indxR; j++ {
			skip := false
			for k := 1; k <= int(dimension); k++ {
				if math.Abs(inLeftPoints[i].GetCoor(k)-inRightPoints[j].GetCoor(k)) > shortestDistance {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			*totalOpt++
			pointDist := point.CalculateDistance(inLeftPoints[i], inRightPoints[j])
			if pointDist < shortestDistance {
				shortestDistance = pointDist
				sdArray = nil
				sdArray = append(sdArray, []point.Point{inLeftPoints[i], inRightPoints[j]})
			} else if point.IsEqual(pointDist, shortestDistance) {
				sdArray = append(sdArray, []point.Point{inLeftPoints[i], inRightPoints[j]})
			}
		}
	}

	return shortestDistance, sdArray
}
