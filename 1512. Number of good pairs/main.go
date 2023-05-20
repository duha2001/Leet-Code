package main

import (
	"fmt"
)

// Cách giải 1 (Sử dụng 2 vòng lặp)
// Độ phức tạp thuật toán: O(n^2) -> (Bởi vì thuật toán phải duyệt qua 2 vòng for nếu mảng có n phần tử nên độ phức tạp sẽ là n^2)
// Độ phức tạp lưu trữ bộ nhớ: O(1) -> (Bởi vì thuật toán chỉ cần gán 1 biến const là result := 0 nên chỉ là O(1))
func numIdenticalPairs(nums []int) int {
	result := 0
	for i := 0; i < len(nums) - 1; i++{
		for j:= i+1; j<len(nums); j++{
			if(nums[i] == nums[j]){
				result++;
			}
		}
	}
	return result
}

// Cách giải 2 (Sử dụng map)
// Tính theo số lần xuất hiện để trở thành cặp tốt
// 1 -> 0 Good
// 2 -> 1 Good
// 3 -> 3 Good
// 4 -> 6 Good
// 5 -> 10 Good
// Key: Phần tử trong mảng
// Value: Số lần xuất hiện của phần tử
// Key: 1 2 3
// Value: 3 1 2
// 1 2 3 1 1 3
// i
// Result: 0 + 1 = 1 + 2 = 3 + 1 = 4
// Độ phức tạp thuật toán: O(n) -> (Bởi vì bài toán chỉ duyệt qua 1 vòng for với n phần tử trong mảng nên độ phức tạp chỉ là On(n))
// Độ phức tạp của lưu trữ bộ nhớ: O(n) -> (Bởi vì bài toán sử dụng cấu trúc dữ liệu lưu trữ là map, vì vậy nếu key trong map là khác nhau thì nó sẽ là n nên độ phức tạp sẽ là O(n))
func numIdenticalPairsMap(nums []int) int {
	result := 0
	numberCountMap := make(map[int]int)
	for i := 0; i < len(nums); i++{
		count, isExist := numberCountMap[nums[i]]
		if isExist{
			result += count
			numberCountMap[nums[i]] = count + 1
		}else{
			numberCountMap[nums[i]] = 1
		}
	}
	return result
}

func main() {
	var exampleArr = [...]int{1,2,3,1,1,3}
	fmt.Println(numIdenticalPairs(exampleArr[:]))
	fmt.Println(numIdenticalPairsMap(exampleArr[:]))
}
