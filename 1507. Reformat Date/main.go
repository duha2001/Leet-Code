package main

func reformatDate(date string) string {
	result := ""

	// Mapping month
	monthMap := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	// Xử lý trường hợp khi số ngày từ 1 -> 9 thì độ dài date là 12 và khi từ 10 trở lên thì độ dài là 13 nên xử lý độ dài date khi bằng 12 sẽ cộng thêm số 0 ở trước
	if len(date) == 12 {
		date = "0" + date
	}
	year := date[9:]
	month := monthMap[date[5:8]]
	day := date[:2]
	result = year + "-" + month + "-" + day
	return result
}
// O(1) -> (Vì thuật toán không sử dụng vòng lặp và chỉ sử dụng cắt chuỗi mà trong golang đã hỗ trợ nên độ phức tạp chỉ là hằng số)
// O(1) -> (Thuật toán sử dụng map để lưu trữ và chỉ lưu trữ cố định 12 tháng nên độ phức tạp chỉ là hằng số)
func main() {
	var date string = "20th Oct 2052"
	print(reformatDate(date))
}