package main

// Cách 1
// Độ phức tạp về thời gian thực thi: O(N) -> (Bởi vì cho các số trong khoảng low, high là n thì nó phải duyệt n lần)
// Độ phức tạp về bộ nhớ: O(1) -> (Bởi vì chỉ khai báo hằng số mặc định)
func countOdds1(low int, high int) int {
	result := 0
	i := low
	for i <= high {
		if i%2 == 1 {
			result++
		}
		i++
	}
	return result
}

// Cách 2:
// 3 4 5 6 7 | Số phần tử: 5  = 7 - 3 + 1
// 8 9 10 | Số phần tử: 3 = 10 - 8 + 1

// Số lượng phần tử trong dãy là số chẵn
// 3 4 5 6 | 4 / 2 = Số phần tử lẻ
// 8 9 10 11 12 13 | 6 / 2 = Số phần tử lẻ

// Số lượng phần tử trong dãy số là số lẻ
// 3 4 5 6 7 | 5 / 2 = 2 + 1
// 2 3 4 5 6 | 5 / 2 = 2
// O(1) -> (Vì khi dãy số có dài bao nhiêu số nhưng thuật toán chỉ thực hiện các phép toán cộng trừ nhân chia)
// O(1)
func countOdds(low int, high int) int {
	result := 0
	numOfElement := high - low + 1
	result = numOfElement / 2
	if (numOfElement % 2 == 1 && low % 2 == 1){
		result++
	}

	return result
}

func main() {
	// Cách 1
	print(countOdds1(3,7))
	// Cách 2
	print(countOdds(3,7))

}