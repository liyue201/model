package util

import "math"

func CalculateUserLevel(point float64) (uint64, float64, float64) {

	if point < 100 {
		return 0, 0, 100
	}

	var pointsInLevel = []float64{100}
	var level, maxLevel uint64 = 1, 999999
	for level <= maxLevel {
		nextPointInLevel := pointsInLevel[level-1] + 100*math.Pow(1.3, float64(level))
		pointsInLevel = append(pointsInLevel, nextPointInLevel)
		if point < nextPointInLevel {
			return level, pointsInLevel[level-1], nextPointInLevel
		}
		level++
	}
	return maxLevel, pointsInLevel[level-1], 0
}
