package p12

import (
	"ctci/container/tree"
)

func getSumPath(t *tree.Tree, target int) int {
	cnt := 0
	getPath(t, target, &cnt)
	return cnt
}

func getPath(t *tree.Tree, target int, cnt *int) []int {
	if t == nil {
		return []int{}
	}

	left := getPath(t.Left, target, cnt)
	right := getPath(t.Right, target, cnt)

	for idx := range left {
		left[idx] += t.Value
		if left[idx] == target {
			*cnt++
		}
	}
	for idx := range right {
		right[idx] += t.Value
		if right[idx] == target {
			*cnt++
		}
	}
	if t.Value == target {
		*cnt++
	}
	result := make([]int, len(left)+len(right)+1)
	copy(result, left)
	copy(result[len(left):len(result)-1], right)
	result[len(result)-1] = t.Value
	// fmt.Println(result)
	return result
}

//CountPathsWithSum solves ch04-p12
func CountPathsWithSum(root *tree.Tree, targetSum int) int {
	return countPathsWithSumRecursive(root, targetSum, 0, make(map[int]int))
}

func countPathsWithSumRecursive(node *tree.Tree, targetSum, runningSum int, pathCount map[int]int) int {
	if node == nil {
		return 0
	}

	runningSum += node.Value
	sum := runningSum - targetSum
	var totalPaths int
	if v, ok := pathCount[sum]; ok {
		totalPaths = v
	} else {
		totalPaths = 0
	}

	if runningSum == targetSum {
		totalPaths++
	}

	incrementHashTable(pathCount, runningSum, 1)
	totalPaths += countPathsWithSumRecursive(node.Left, targetSum, runningSum, pathCount)
	totalPaths += countPathsWithSumRecursive(node.Right, targetSum, runningSum, pathCount)
	incrementHashTable(pathCount, runningSum, -1)

	return totalPaths
}

func incrementHashTable(hashTable map[int]int, key, delta int) {
	newCount := delta
	if v, ok := hashTable[key]; ok {
		newCount += v
	}
	if newCount == 0 {
		delete(hashTable, key)
	} else {
		hashTable[key] = newCount
	}
}
