package leetcode

import (
	"sort"
)

//最优解
func sortedSquares(nums []int) []int {
	left,right,k := 0, len(nums)-1,len(nums)-1
	ans := make([]int, len(nums))
	for ;left<=right;k-- {
		if abs(nums[left])>=abs(nums[right]) {
			ans[k] = nums[left]*nums[left]
			left++
		} else {
			ans[k] = nums[right]*nums[right]
			right--
		}
	}
	return ans
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sortedSquares2(nums []int) []int {
	ans := make([]int, len(nums))
	for i,k,j := 0,len(nums)-1,len(nums)-1;i<=j;k--{
		if nums[i]*nums[i] > nums[j]*nums[j] {
			ans[k] = nums[i]*nums[i]
			i++
		} else {
			ans[k] = nums[j]*nums[j]
			j--
		}
	}
	return ans
}

func sortedSquares3(nums []int) []int {
	stack := []int{}
	ans := []int{}
	index := 0

	for k, v := range nums {
		index = k
		if v <= 0 {
			stack = append(stack, v)
		} else {
			break
		}
	}

	if len(stack) == len(nums) {
		index += 1
	}
		for i := index; i<len(nums); {
			if len(stack) == 0 || nums[i]*nums[i] <= stack[len(stack)-1]*stack[len(stack)-1] {
				ans = append(ans, nums[i]*nums[i])
				i++
			} else {
				ans = append(ans, stack[len(stack)-1]*stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}

	if len(stack) > 0{
		for j:=len(stack)-1;j>=0;j--{
			ans = append(ans, stack[len(stack)-1]*stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return ans
}

func sortedSquares4(nums []int) []int {
	ans := []int{}
	for _, v := range nums {
		ans = append(ans, v*v)
	}
	sort.Ints(ans)
	return  ans
}