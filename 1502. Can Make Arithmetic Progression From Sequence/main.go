package main

// min, max
// min = 1
// max = 9
// N = ?
// diff = 1 -> 1 2 3 4 5 6 7 8 9 | N = 9 | (9-1)/(9-1)=1
// diff = 2 -> 1 3 5 7 9 | N = 5 | (9-1)/(5-1)=2
// Chắc chắn khi chia kết quả là số chẵn
// diff = (max - min) / (N - 1)
// (max - min) / (N-1) không chia chẵn -> False
// Cần kiểm tra xem các phần tử trong dãy có phải là cấp số cộng không ví dụ
// 1 2 5 7 9 | N = 5 | 8/4 = 2
// 1 + 2 = 3 <> 2 -> False

func canMakeArithmeticProgression(arr []int) bool {
	N := len(arr)
	min := arr[0]
	max := arr[0]
	// Sử dụng cái map để check phần tử tồn tại hay chưa
	elementExist := make(map[int]bool)
	for i := 0; i < N; i++ {
		if (arr[i] < min){
			min = arr[i]
		}
		if (arr[i] > max){
			max = arr[i]
		}
		elementExist[arr[i]] = true
	}
	totalDiff := max - min
	if totalDiff % (N - 1) != 0 {
		return false
	}
	diff := totalDiff / (N - 1)
	for min < max {
		min += diff
		_, isExist := elementExist[min]
		if !isExist{
			return false
		}
	}
	return true
}

func main() {

}