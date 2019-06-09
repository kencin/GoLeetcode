package main

import (
	"fmt"
	"sort"
)

/*
给定一个整数数组，返回所有数对之间的第 k 个最小距离。一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。

示例 1:

输入：
nums = [1,3,1]
k = 1
输出：0
解释：
所有数对如下：
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
因此第 1 个最小距离的数对是 (1,1)，它们之间的距离为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
思路分析：
首先不管你是要求第几小的距离对，我先二分。
求出mid 是第几小的距离对，然后mid 与k比较，调整上下界，
重复以上步骤。

*/

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	length := len(nums)
	left, right := 0, nums[length-1]-nums[0] //距离对的最小/大值
	for left < right {
		mid := (left + right) / 2
		j := 0
		count := 0
		for i := 1; i < length; i++ {
			for nums[i]-nums[j] > mid {
				j++
			}
			count += i - j
		}
		if count < k { // 说明mid太小，第k小的距离对在二分右边
			left = mid + 1
		} else { // mid太大，第k小的距离对在二分左边
			right = mid
		}
	}

	return left //left和right都行

	//length := len(nums)
	//result := []int{}
	//for i:=0;i<length;i++{
	//	for j:=i+1;j<length;j++{
	//		AbsoluteValue := 0
	//		if AbsoluteValue = nums[i] - nums[j];AbsoluteValue < 0{
	//			AbsoluteValue = -AbsoluteValue
	//		}
	//		result = append(result,AbsoluteValue)
	//	}
	//}
	////排序
	//sort.Ints(result)
	////去重
	//endResult := make([]int, 0, len(result))
	//temp := map[int]struct{}{}
	//for _,item:=range result{
	//	if _,ok:=temp[item]; !ok{
	//		temp[item] = struct{}{}
	//		endResult = append(endResult, item)
	//	}
	//}
	//fmt.Println(result)
	//if k > len(endResult){
	//	return endResult[len(endResult)-1]
	//}
	//return endResult[k-1]
}
func main() {
	input := []int{1, 2, 3, 4, 5}
	fmt.Println(smallestDistancePair(input, 2))
}
