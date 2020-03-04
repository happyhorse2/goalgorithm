package shujuandslice

import "fmt"

// GetLeastNumbers 获取无需数组中最小k个数
func GetLeastNumbers(arr []int, k int) []int {
	if k > len(arr) {
		return nil
	}
	var result []int
	quickstort(0, len(arr)-1, arr)
	for i := 0; i < k; i++ {
		result = append(result, arr[i])
		fmt.Printf("%p\n", result)
		fmt.Println(cap(result))
	}
	fmt.Println(arr)
	return result
}

func quickstort(startindex, endindex int, arr []int) {
	if startindex < endindex {
		partindex := partition(startindex, endindex, arr)
		quickstort(startindex, partindex-1, arr)
		quickstort(partindex+1, endindex, arr)
	}
}

func partition(startindex, endindex int, arr []int) int {
	comparenumber := arr[startindex]
	for startindex < endindex {
		for arr[endindex] >= comparenumber && startindex < endindex {
			endindex--
		}
		arr[startindex] = arr[endindex]
		for arr[startindex] <= comparenumber && startindex < endindex {
			startindex++
		}
		arr[endindex] = arr[startindex]
	}
	arr[startindex] = comparenumber
	return startindex
}
