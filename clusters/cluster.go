package clusters

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义聚类簇的结构体
type Cluster struct {
	Center       Coordinates
	Observations Observations
}

// 使用切片存储
type Clusters []Cluster

// 新设置一组新的簇并随机种子其初始位置
func New(k int, dataset Observations) (Clusters, error) {
	var c Clusters
	if len(dataset) == 0 || len(dataset[0].Coordinates()) == 0 {
		return c, fmt.Errorf("there must be at least one dimension in the data set")
	}
	if k == 0 {
		return c, fmt.Errorf("k must be greater than 0")
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++ {
		var p Coordinates
		for j := 0; j < len(dataset[0].Coordinates()); j++ {
			p = append(p, rand.Float64())
		}

		c = append(c, Cluster{
			Center: p,
		})
	}
	return c, nil
}

// 添加新聚类簇
func (c *Cluster) Append(point Observation) {
	c.Observations = append(c.Observations, point)
}

// 返回最接近点的聚类簇索引
func (c Clusters) Nearest(point Observation) int {
	var ci int
	dist := -1.0

	// 找到近邻聚类数据点
	for i, cluster := range c {
		d := point.Distance(cluster.Center)
		if dist < 0 || d < dist {
			dist = d
			ci = i
		}
	}

	return ci
}

// 返回一个点的相邻簇以及到其点的平均距离
func (c Clusters) Neighbour(point Observation, fromCluster int) (int, float64) {
	var d float64
	nc := -1

	for i, cluster := range c {
		if i == fromCluster {
			continue
		}

		cd := AverageDistance(point, cluster.Observations)
		if nc < 0 || cd < d {
			nc = i
			d = cd
		}
	}

	return nc, d
}

// 重置聚类簇
func (c *Cluster) Recenter() {
	center, err := c.Observations.Center()
	if err != nil {
		return
	}

	c.Center = center
}

// 重置所有聚类簇
func (c Clusters) Recenter() {
	for i := 0; i < len(c); i++ {
		c[i].Recenter()
	}
}

// 重置
func (c Clusters) Reset() {
	for i := 0; i < len(c); i++ {
		c[i].Observations = Observations{}
	}
}

// 返回给定区间中的所有坐标
func (c Cluster) PointsInDimension(n int) Coordinates {
	var v []float64
	for _, p := range c.Observations {
		v = append(v, p.Coordinates()[n])
	}
	return v
}

// 返回指定区间的聚类簇
func (c Clusters) CentersInDimension(n int) Coordinates {
	var v []float64
	for _, cl := range c {
		v = append(v, cl.Center[n])
	}
	return v
}
