package main

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(CheckTwiceNumbers(nums))
// 	fmt.Println(5 ^ 3) // 101 ^ 011 => 110
// }

func CheckTwiceNumbers(nums []int) bool {
	numsCount := make(map[int]int)
	for _, val := range nums {
		numsCount[val]++
		if numsCount[val] > 1 {
			return true
		}
	}
	return false
}
