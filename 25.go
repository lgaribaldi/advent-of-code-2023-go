package main

import (
	"io"
	"os"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func day25(c chan string) {

	graph := make(map[string][]string)
	var graphNodes = []opts.GraphNode{}
	links := make([]opts.GraphLink, 0)

	for line := range c {
		if len(line) == 0 {
			continue
		}
		node, destList := splitString(line, ": ")
		dest := strings.Split(destList, " ")
		for _, d := range dest {
			links = append(links, opts.GraphLink{Source: node, Target: d})
			links = append(links, opts.GraphLink{Source: d, Target: node})
			if existingNode, ok := graph[d]; ok {
				graph[d] = append(existingNode, node)
			} else {
				graph[d] = append(make([]string, 0), node)
			}
		}

		if existingNode, ok := graph[node]; ok {
			graph[node] = append(existingNode, dest...)
		} else {
			graph[node] = dest
		}
	}
	for key := range graph {
		graphNodes = append(graphNodes, opts.GraphNode{Name: key})
	}

	visited := make(map[string]struct{})
	countNodes("krj", &graph, &visited)
	count := 0
	for key := range visited {
		if _, ok := visited[key]; ok {
			count++
		}
	}
	println("krj: ", count)

	visited = make(map[string]struct{})
	countNodes("xcm", &graph, &visited)
	count = 0
	for key := range visited {
		if _, ok := visited[key]; ok {
			count++
		}
	}
	println("xcm: ", count)

	// I used this to render the graph and see where to disconect
	// not my proudest moment
	page := components.NewPage()
	page.AddCharts(
		graphBase(&graphNodes, &links),
	)
	f, err := os.Create("graph.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}

func countNodes(node string, graph *map[string][]string, visited *map[string]struct{}) {
	if _, ok := (*visited)[node]; ok {
		return
	}
	(*visited)[node] = struct{}{}
	for _, v := range (*graph)[node] {
		countNodes(v, graph, visited)
	}
}

func graphBase(nodes *[]opts.GraphNode, links *[]opts.GraphLink) *charts.Graph {
	graph := charts.NewGraph()
	graph.AddSeries("graph", *nodes, *links,
		charts.WithGraphChartOpts(
			opts.GraphChart{Force: &opts.GraphForce{Repulsion: 6}},
		),
	)
	return graph
}
