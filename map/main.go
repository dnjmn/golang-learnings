package main

import "fmt"

func main() {
	v := []int{2, 2, -1, 2, 2, 2}
	fmt.Println(numOfMinutes(2, v, []int{0, 0, 1, 0, 0, 0}))
}

func numOfMinutes(headID int, manager []int, informTime []int) int {
	managerToEmps := getAdjList(manager)
	fmt.Println(managerToEmps)
	return dfs(headID, managerToEmps, informTime)
}

func dfs(currentId int, managerToEmps map[int][]int, informTime []int) int {
    if l, ok := managerToEmps[currentId]; !ok || len(l) == 0 {
		return 0
	}
	max := 0
	for _, manager := range managerToEmps[currentId] {
		current := dfs(manager, managerToEmps, informTime)
		if current > max {
			max = current
		}
	}

	return max + informTime[currentId]
}

func getAdjList(manager []int) map[int][]int {
	adjList := make(map[int][]int)

	for emp, manager := range manager {
		if manager == -1 {
			continue
		}
		if list, ok := adjList[manager]; !ok {
			adjList[manager] = []int{emp}
		} else {
			adjList[manager] = append(list, emp)
		}
	}
	return adjList
}
