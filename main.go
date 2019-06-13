package main

import (
	"io"
	"os"
	"log"
	"fmt"
	"bufio"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/leeli73/go-kmeans-html-plotter/clusters"
	"github.com/leeli73/go-kmeans-html-plotter/kmeans"
)
// 读取数据
func InitData() clusters.Observations{
	fi, err := os.Open("data/iris.dat")
    if err != nil {
        log.Fatalln(err)
        return nil
    }
    defer fi.Close()

	var d clusters.Observations
    br := bufio.NewReader(fi)
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
		temp := string(a)
		if temp[0] != '@'{
			data := strings.Split(temp,", ")
			num1,_ := strconv.ParseFloat(data[0], 64)
			num2,_ := strconv.ParseFloat(data[1], 64)
			num3,_ := strconv.ParseFloat(data[2], 64)
			num4,_ := strconv.ParseFloat(data[3], 64)
			d = append(d,clusters.Coordinates{
				num1,
				num2,
				num3,
				num4,
			})
		}
	}
	return d
}
func main() {
	d := InitData()
	//定义一个k-means
	km, _ := kmeans.New(0.01, kmeans.SimplePlotter{}) 
	//进行聚类运算
	clusters,files, _ := km.Partition(d, 4)

	//生成演示html
	strfiles := "\"" + files[0] + "\","
	for i:=1;i<len(files) - 1;i++{
		strfiles = strfiles + "\"" + files[i] + "\","
	}
	strfiles = strfiles + "\"" + files[len(files)-1] + "\""

	clusterout := []string{}
	for i, c := range clusters {
		log.Printf("Cluster: %d %+v", i, c.Center)
		str := fmt.Sprintf("Cluster: %d %+v", i, c.Center)
		clusterout = append(clusterout,str)
	}

	clustersstr := "\"" + clusterout[0] + "\","
	for i:=1;i<len(clusterout)-1;i++{
		clustersstr = clustersstr + "\"" + clusterout[i] + "\","
	}
	clustersstr = clustersstr + "\"" +clusterout[len(clusterout)-1] + "\""
	var Data,err = ioutil.ReadFile("data/web.html")
	if err != nil{
		log.Fatal(err)
	}
	html := string(Data)
	html = strings.Replace(html,"{{images}}",strfiles,-1)
	html = strings.Replace(html,"{{clusters}}",clustersstr,-1)
	ioutil.WriteFile("k-means.html",[]byte(html), 0644)
}

