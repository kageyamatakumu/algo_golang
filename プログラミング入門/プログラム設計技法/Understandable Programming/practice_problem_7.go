package main

import (
	"fmt"
	"math"
)

func main() {
	scores := [] int{70, 50, 85, 69, 63}

	ave := average(scores)

	sd := standardDeviation(scores, ave)

	dv := deviationValue(scores, ave, sd)

	fmt.Println(dv)
}

// 平均値
func average(scores []int) float32 {
	var sum int
	var ave float32

	for _, v := range scores {
		sum += v
	}

	ave = float32(sum / len(scores))

	return ave
}

// 標準偏差
func standardDeviation(scores []int, ave float32) float32 {
	var sd float32
	for _, v := range scores {
		sd += (float32(v) - ave) * (float32(v) - ave)
	}

	sd = sd / float32(len(scores))

	return float32(math.Sqrt(float64(sd)))
}

// 偏差値
func deviationValue(scores []int, ave float32, sd float32) float32 {
	var dv float32

	dv = 50 + 10 * (float32(scores[0]) - ave) / sd

	return dv
}
