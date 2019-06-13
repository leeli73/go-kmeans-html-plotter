package kmeans

import (
	"fmt"
	"math/rand"
	"github.com/leeli73/go-kmeans-html-plotter/clusters"
)

// Kmeans结构体
type Kmeans struct {
	// 定义绘图类，用于每次计算后绘图
	plotter Plotter
	// 终止精度
	deltaThreshold float64
	// 计算最大的终止次数
	iterationThreshold int
}

// 初始化
func New(deltaThreshold float64, plotter Plotter) (Kmeans, error) {
	if deltaThreshold <= 0.0 || deltaThreshold >= 1.0 {
		return Kmeans{}, fmt.Errorf("threshold is out of bounds (must be >0.0 and <1.0, in percent)")
	}

	return Kmeans{
		plotter:            plotter,
		deltaThreshold:     deltaThreshold,
		iterationThreshold: 96,
	}, nil
}

// 数据集进行k-means计算，并将其划分为k个簇。
func (m Kmeans) Partition(dataset clusters.Observations, k int) (clusters.Clusters,[]string,error) {
	AllFile := []string{}
	if k > len(dataset) {
		return clusters.Clusters{},AllFile,fmt.Errorf("the size of the data set must at least equal k")
	}

	cc, err := clusters.New(k, dataset)
	if err != nil {
		return cc,AllFile,err
	}

	points := make([]int, len(dataset))
	changes := 1

	for i := 0; changes > 0; i++ {
		changes = 0
		cc.Reset()

		for p, point := range dataset {
			ci := cc.Nearest(point)
			cc[ci].Append(point)
			if points[p] != ci {
				points[p] = ci
				changes++
			}
		}

		for ci := 0; ci < len(cc); ci++ {
			if len(cc[ci].Observations) == 0 {
				// 在迭代过程中，如果任何一个集群中心没有与之相关联的数据点，则为其分配一个随机数据点。
				var ri int
				for {
					// 找到至少有两个数据点的数据集
					ri = rand.Intn(len(dataset))
					if len(cc[points[ri]].Observations) > 1 {
						break
					}
				}
				cc[ci].Append(dataset[ri])
				points[ri] = ci

				// 确保在将数据点随机分配
				changes = len(dataset)
			}
		}

		if changes > 0 {
			cc.Recenter()
		}
		if m.plotter != nil {
			filename,_ := m.plotter.Plot(cc, i)
			AllFile = append(AllFile,filename)
		}
		if i == m.iterationThreshold ||
			changes < int(float64(len(dataset))*m.deltaThreshold) {
			break
		}
	}

	return cc,AllFile,nil
}
