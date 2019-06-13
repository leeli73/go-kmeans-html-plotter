package clusters

import (
	"fmt"
	"math"
)

// 使用切片存储
type Coordinates []float64

// 定义一个n维数据点
type Observation interface {
	Coordinates() Coordinates
	Distance(point Coordinates) float64
}

// 使用切片存储
type Observations []Observation

// 实现一个坐标
func (c Coordinates) Coordinates() Coordinates {
	return Coordinates(c)
}

// 距离返回两个坐标之间的欧氏距离
func (c Coordinates) Distance(p2 Coordinates) float64 {
	var r float64
	for i, v := range c {
		r += math.Pow(v-p2[i], 2)
	}
	return r
}

// 返回指定数据的中心坐标
func (c Observations) Center() (Coordinates, error) {
	var l = len(c)
	if l == 0 {
		return nil, fmt.Errorf("there is no mean for an empty set of points")
	}

	cc := make([]float64, len(c[0].Coordinates()))
	for _, point := range c {
		for j, v := range point.Coordinates() {
			cc[j] += v
		}
	}

	var mean Coordinates
	for _, v := range cc {
		mean = append(mean, v/float64(l))
	}
	return mean, nil
}

// 返回某个点到指定数据的平均距离
func AverageDistance(o Observation, observations Observations) float64 {
	var d float64
	var l int

	for _, observation := range observations {
		dist := o.Distance(observation.Coordinates())
		if dist == 0 {
			continue
		}

		l++
		d += dist
	}

	if l == 0 {
		return 0
	}
	return d / float64(l)
}
