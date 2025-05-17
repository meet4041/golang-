package main

func main() {

	// 'SLICES' -> dynamic
	// Most used construct in GO
	// useful methods

	// Declare
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// Capacity - max no.of element can be fit
	// fmt.Println(cap(nums))

	// Length - no.of element in an array
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// nums := []int{}
	// nums = append(nums, 1, 2)
	// fmt.Println(nums)
	// fmt.Println("Length: ", len(nums))
	// fmt.Println("Capacity: ", cap(nums))

	// var a = make([]int, 0, 5)
	// a = append(a, 2, 4, 5)
	// var b = make([]int, len(a), 5)

	// Copy Function
	// copy(b, a)
	// fmt.Println(a, b)

	// Slice Operator
	// var c = []int{1, 2, 3, 4, 5}
	// fmt.Println(c[2:3])
	// fmt.Println(c[:3])

	// Slice
	// var abc = []int{1, 2, 3}
	// var bbc = []int{1, 2, 3}
	// var bbc = []int{1, 2, 5}

	// fmt.Println("And the array is:", slices.Equal(abc, bbc))

	// 2D slice
	// var _2d = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(_2d)

}

// Commmand to run this file -
// go run 12.go
