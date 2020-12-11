package main

import (
	"fmt"
	"math"
)

func main() {
	graph := make(map[string]map[string]int)
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = map[string]int{}
	graph["a"]["finish"] = 1

	graph["b"] = map[string]int{}
	graph["b"]["a"] = 3
	graph["b"]["finish"] = 5

	graph["finish"] = map[string]int{}

	fmt.Println(graph)
	costs, parents := findShortestPath(graph, "start", "finish")
	fmt.Println(costs, parents)
}

func findShortestPath(graph map[string]map[string]int, start, finished string) (map[string]int, map[string]string) {
	costs := make(map[string]int)
	costs[finished] = math.MaxInt32

	parents := make(map[string]string)
	parents[finished] = ""

	processed := make(map[string]bool)
	for node, cost := range graph[start] {
		costs[node] = cost
		parents[node] = start
	}

	fmt.Println(costs, parents)

	lowestNode := findLowestCostNode(costs, processed)

	for lowestNode != "" {
		for node, cost := range graph[lowestNode] {
			newCost := costs[lowestNode] + cost
			if newCost < costs[node] {
				costs[node] = newCost
				parents[node] = lowestNode
			}
		}
		processed[lowestNode] = true
		lowestNode = findLowestCostNode(costs, processed)
	}

	return costs, parents
}

func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := math.MaxInt32
	lowestNode := ""

	for node, cost := range costs {
		if _, ok := processed[node]; !ok && cost < lowestCost {
			lowestCost = cost
			lowestNode = node
		}
	}
	fmt.Println("lowestNode=", lowestNode)
	return lowestNode
}
