package main

func restoreString(s string, indices []int) string {
	result := make([]byte, len(indices))
	for index, value := range indices {
		result[value] = s[index]
	}

	return string(result)

}

// O(N)
// O(N)
// Solution
// Thuật toán ở đây chỉ cần dùng make để tạo ra 1 mảng byte muốn xử lý trên từng ký tự trên chuỗi với độ dài mảng là độ dài của mảng input ban đầu, sau đó dùng range để lặp dễ dàng với 2 chỉ số là index và value
func main() {

}