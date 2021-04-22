package leetcode

import (
	"fmt"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	//nums := []int{-4, -1, 0, 3, 10}
	//nums := []int{-5,-3,-2,-1}
	//nums := []int{-4,-1,0,3,10}
	nums := []int{1}
	fmt.Printf("para:%v\n",nums)
	res := sortedSquares(nums)
	fmt.Printf("result:%v\n",res)
}