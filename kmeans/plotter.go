package kmeans

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"github.com/leeli73/go-kmeans-html-plotter/clusters"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

// 绘制图片
type Plotter interface {
	Plot(cc clusters.Clusters, iteration int) (string,error)
}

// SimplePlotter绘制二维数据集
type SimplePlotter struct {
}

// 定义调色板
var colors = []drawing.Color{
	drawing.ColorFromHex("f92672"),
	drawing.ColorFromHex("89bdff"),
	drawing.ColorFromHex("66d9ef"),
	drawing.ColorFromHex("67210c"),
	drawing.ColorFromHex("7acd10"),
	drawing.ColorFromHex("af619f"),
	drawing.ColorFromHex("fd971f"),
	drawing.ColorFromHex("dcc060"),
	drawing.ColorFromHex("545250"),
	drawing.ColorFromHex("4b7509"),
}

// 绘制点并输出到png
func (p SimplePlotter) Plot(cc clusters.Clusters, iteration int) (string,error) {
	var series []chart.Series

	// 绘制数据点
	for i, c := range cc {
		series = append(series, chart.ContinuousSeries{
			Style: chart.Style{
				Show:        true,
				StrokeWidth: chart.Disabled,
				DotColor:    colors[i%len(colors)],
				DotWidth:    8,
			},
			XValues: c.PointsInDimension(0),
			YValues: c.PointsInDimension(1),
		})
	}

	// 绘制聚类中心点
	series = append(series, chart.ContinuousSeries{
		Style: chart.Style{
			Show:        true,
			StrokeWidth: chart.Disabled,
			DotColor:    drawing.ColorBlack,
			DotWidth:    16,
		},
		XValues: cc.CentersInDimension(0),
		YValues: cc.CentersInDimension(1),
	})

	// 定义输出png的参数
	graph := chart.Chart{
		Height: 1024,
		Width:  1024,
		Series: series,
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return "",err
	}
	filename := fmt.Sprintf("data/images/%d_%d.png", len(cc), iteration)
	return filename,ioutil.WriteFile(filename, buffer.Bytes(), 0644)
}
