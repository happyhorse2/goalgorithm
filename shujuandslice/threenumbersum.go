package shujuandslice

import (
	"fmt"
	"math"
	"sort"
)

// ThreeSum 从num中找到三个数字加和为0
//思路：先对给定数组排序，然后使用双指针算法，从左到右，重点是，避免重复，避免重复的重点是，从后往前判断，注意0边界特殊点
func ThreeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	fmt.Println(nums)
	var threeSumnumber [][]int
	for i, number := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			firstindex := i + 1
			endindex := len(nums) - 1
			for {
				if firstindex >= endindex {
					break
				}
				fmt.Println(firstindex, endindex, i)
				if nums[firstindex]+nums[endindex] < (0 - number) {
					firstindex++
				} else if nums[firstindex]+nums[endindex] > (0 - number) {
					endindex--
				} else {
					if (firstindex == i+1) || nums[firstindex] != nums[firstindex-1] { //边界为0的特殊判断，firstindex=i+1,   nums[firstindex] != nums[firstindex-1]从index开始判断去重复

						tempslice := []int{number, nums[firstindex], nums[endindex]}
						threeSumnumber = append(threeSumnumber, tempslice)
					}
					firstindex++
					endindex--
				}
			}
		}
	}
	fmt.Println(threeSumnumber)
	return threeSumnumber
}

//Slicetest 切片联系，append函数能否创建空间？  可以创建空间
func Slicetest() {
	var testslice []int
	testslice = append(testslice, 5)
	fmt.Println(testslice, cap(testslice), len(testslice))
}

// ThreeSumClosest 从num中找到三个数字加和target最接近
func ThreeSumClosest(nums []int, target int) int {
	if nums == nil || len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	initialnumber := nums[0] + nums[1] + nums[2]
	for i, number := range nums {
		firstindex := i + 1
		endindex := len(nums) - 1
		for {
			if firstindex >= endindex {
				break
			}
			comparenumber := nums[firstindex] + nums[endindex] + number
			if int(math.Abs(float64(comparenumber-target))) < int(math.Abs(float64(initialnumber-target))) {
				initialnumber = comparenumber
			}
			if comparenumber > target {
				endindex--
			} else if comparenumber < target {
				firstindex++
			} else {
				break
			}
		}
	}
	return initialnumber
}
