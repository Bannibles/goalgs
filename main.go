package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/set"
	"github.com/golang-collections/collections/stack"
)

func main() {
	var choice float64
	choice = 555
	switch choice {
	case 11:
		example1_1()
	case 12:
		example1_2()
	case 13:
		example1_3()
	case 14:
		example1_4()
	case 15:
		example1_5(5, 6)
	case 16:
		example1_6(3, 9)
	case 17:
		example1_7(3, 88, 300)
	case 18:
		example1_8()
	case 19:
		example1_9()
	case 110:
		example1_10(4)
	case 111:
		example1_11()
	case 112:
		example1_12()
	case 113:
		example1_13()
	case 114:
		example1_14()
	case 115:
		example1_15()
	case 116:
		example1_16_1()
	case 117:
		example1_17_1()
	case 118:
		example1_18()
	case 119:
		example1_19_1()
	case 120:
		example1_20()
	case 121:
		example1_21()
	case 122:
		example1_22()
	case 124:
		example1_24()
	case 125:
		example1_25()
	case 127:
		example1_27()
	case 132:
		example1_32_1()
	case 134:
		fmt.Println(example1_34(4))
	case 137:
		fmt.Println(example1_37(9))
	case 199:
		test_sorting()
	case 21:
		example2_1(100)
	case 22:
		example2_2(100)
	case 23:
		example2_3(100)
	case 24:
		example2_4(100)
	case 25:
		example2_5(100)
	case 26:
		example2_6(100)
	case 27:
		example2_7(100)
	case 28:
		example2_8(100)
	case 29:
		example2_9(100)
	case 210:
		example2_10(100)
	case 211:
		example2_11(100)
	case 212:
		example2_12(100)
	case 213:
		example2_13(100)
	case 214:
		example2_14()
	case 41:
		example4_1()
	case 42:
		example4_2()
	case 43:
		example4_3()
	case 44:
		example4_4()
	case 51:
		arr := []int{2, 4, 5, 8, 5, 3, 5, 7, 8, 5, 34, 33, 5, 6, 8, 9, 7, 5, 12}
		linearSearchInArray(arr, 12)
	case 52:
		arr := []int{2, 4, 5, 8, 5, 3, 5, 7, 8, 5, 34, 33, 5, 6, 8, 9, 7, 5, 12}
		example5_2(arr, 12)
	case 55:
		arr := []int{1, 7, 3, 8, 5, 3, 8, 0, 6, 4, 3, 8, 0, 7, 4, 2, 5, 8}
		example5_5(arr)
	case 56:
		arr := []int{1, 7, 3, 8, 5, 3, 8, 0, 6, 4, 3, 8, 0, 7, 4, 2, 5, 8}
		example5_6(arr)
	case 57:
		arr := []int{1, 7, 3, 8, 5, 3, 8, 0, 6, 4, 3, 8, 0, 7, 4, 2, 5, 8}
		example5_7(arr)
	case 58:
		arr := []int{1, 7, 3, 8, 5, 3, 8, 0, 6, 4, 3, 8, 0, 7, 4, 2, 5, 8}
		example5_8(arr, 19)
	case 59:
		arr := []int{1, 7, 3, 8, 5, 3, 8, 0, 6, 4, 3, 8, 0, 7, 4, 2, 5, 8}
		example5_9(arr)
	case 510:
		arr := []int{1, 7, 3, 8, 5, 3, 8, 0, 6, 4, 3, 8, 0, 7, 4, 2, 5, 8}
		example5_10(arr)
	case 511:
		arr := []int{1, 7, 3, 8, 5, 3, 8, 0, 6, 4, 3, 8, 0, 7, 4}
		example5_11(arr, 19)
	case 513:
		arr := []int{1, 1, 1, 1, 1, 1, 3, 5, 1, 1, 3, 4, 1, 43, 5, 1, 1, 1, 4, 2, 3}
		example5_13(arr)
	case 514:
		//arr := []int{1, 1, 1, 1, 2, 1, 1, 2, 2, 2, 1, 1, 1, 1, 2, 1, 1, 1, 1}
		//arrayMajority2(arr)
	case 517:
		data := []int{1432, 4312, 1324, 1324, 3241, 4132, 1243, 1423, 2341, 3421, 2314, 2314, 2314, 3421, 4132, 4312, 3412, 3412, 3412, 4132}
		findPair(data, 1324)
	case 518:
		data := []int{132, 123, 132, 123, 123, 231, 321, 123, 312, 132, 321, 123, 123, 321, 123, 123, 123, 123, 123, 231, 213, 213, 231, 213, 213, 123, 123, 123, 213, 132, 321}
		findPair2(data, 123)
	case 520:
		data := []int{11, 30, 56, 76, 54, 23, 1254, 67, 98, 65, 23, 21, 32, 54, 1121, 12, 534}
		minAbsSumPair(data)
	case 519:
		data := []int{11, 30, 56, 76, 54, 23, 1254, 67, 98, 65, 23, 21, 32, 54, 1121, 12, 534}
		minAbsSumPair2(data)
	case 522:
		data := []int{1, 3, 5, 7, 9, 12, 8, 6, 4, 2, -1}
		SearchBitonicArrayMax(data)
	case 523:
		data := []int{1, 3, 5, 7, 9, 12, 8, 6, 4, 2, -1}
		SearchBitonicArray(data, 6)
	default:
		fmt.Println("Invalid function")
	}
}

func example1_1() {
	s := "hello, world!"
	fmt.Printf("%s\n", s)
}

func example1_2() {
	v1 := 100
	fmt.Println("Value stored in v1 :: ", v1)
}

func example1_3() {
	maxInt8 := math.MaxInt8
	minInt8 := math.MinInt8
	fmt.Println("Rang of Int8 :: ", minInt8, " to ", maxInt8)
}

func example1_4() {
	s := "hello, world!"
	r := []rune(s)
	r[0] = 'H'
	s2 := string(r)
	fmt.Println(s2)
}

func example1_5(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

func example1_6(x, y int) int {
	var max int
	if x > y {
		max = x
	} else {
		max = y
	}
	return max
}

func example1_7(length, width, limit int) bool {
	if area := length * width; area < limit {
		return true
	} else {
		return false
	}
}

func example1_8() {
	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("something else")
	}
}

func example1_9() {
	i := 2
	switch i {
	case 1, 2, 3:
		fmt.Println("One, two, or three")
	default:
		fmt.Println("something else")
	}
}

func example1_10(value int) {
	switch {
	case value%2 == 0:
		fmt.Println("value is even")
	default:
		fmt.Println("value is odd")
	}
}

func example1_11() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("sum is :: ", sum)
}

func example1_12() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	i := 0
	n := len(numbers)
	for i < n {
		sum += numbers[i]
		i++
	}
	fmt.Println("Sum is :: ", sum)
}

func example1_13() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	i := 0
	n := len(numbers)
	for {
		sum += numbers[i]
		i++
		if i >= n {
			break
		}
		fmt.Println("Sum is :: ", sum)
	}
}

func example1_14() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println("Sum is :: ", sum)
}

func example1_15() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for index, val := range numbers {
		sum += val
		fmt.Println("[", index, ",", val, "]")
	}
	fmt.Println("\nSum is :: ", sum)
	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Println(k, " -> ", v)
	}
	str := "Hello World!"
	for index, c := range str {
		fmt.Println("[", index, ",", string(c), "]")
	}
}

func example1_16(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func example1_16_1() {
	fmt.Println(example1_16(10, 20))
}

func example1_17(x int) {
	x++
}

func example1_17_1() {
	i := 10
	fmt.Println("Value of i before increment is : ", i)
	example1_17(i)
	fmt.Println("Value of i after increment is : ", i)
}

func example1_18() {
	data := 10
	ptr := &data
	fmt.Println("Value stored at variable car is ", data) //value of data
	fmt.Println("Value stored at variable car is ", *ptr)
	fmt.Println("The address of variable car is ", &data) //memory address of data
	fmt.Println("The address of variable car is ", ptr)
}

func example1_19(ptr *int) {
	(*ptr)++
}

func example1_19_1() {
	i := 10
	fmt.Println("Value of i before increment is: ", i) //value of i
	example1_19(&i)
	fmt.Println("Value of i after increment is : ", i) //value of i w/ increment
}

type student struct {
	rollNo int
	name   string
	count  int
}

func (s *student) Count() {
	s.count++
}

func countStudent(s *student) {
	s.count++
}

func example1_20() {
	stud := student{1, "Johnny", 1}
	fmt.Println(stud)
	fmt.Println("Student name ::", stud.name)
	fmt.Println("Student count ::", stud.count)
	pstud := &stud
	countStudent(pstud)
	fmt.Println("Student name ::", pstud.name)
	fmt.Println("Student count ::", stud.count)
	fmt.Println(student{rollNo: 2, name: "Ann"})
	fmt.Println(student{name: "Ann", rollNo: 2})
	fmt.Println(student{name: "Alice"})
}

type Rect struct {
	width  float64
	length float64
}

func (r Rect) Area() float64 {
	return r.width * r.length
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.length)
}

func example1_21() {
	r := Rect{width: 10, length: 10}
	fmt.Println("Area: ", r.Area())
	fmt.Println("Perimeter: ", r.Perimeter())

	ptr := &Rect{width: 10, length: 5}
	fmt.Println("Area: ", ptr.Area())
	fmt.Println("Perimeter: ", ptr.Perimeter())
}

type MyInt int

func (data MyInt) increment1() {
	data = data + 1
}

func (data *MyInt) increment2() {
	*data = *data + 1
}

func example1_22() {
	var data MyInt = 1
	fmt.Println("value before increment1() call :", data)
	data.increment1()
	fmt.Println("value after increment1() call :", data)
	data.increment2()
	fmt.Println("value after increment2() call :", data)
}

func example1_24() {
	var arr [10]int
	fmt.Println(arr)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	count := len(arr)
	fmt.Println("the length of arr is ", count)
}

func example1_25() {
	var bruh []int
	for i := 0; i < 100; i++ {
		bruh = append(bruh, i)
		PrintSlice(bruh)
	}
}

// PrintSlice prints slice
func PrintSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

func example1_27() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50

	for key, val := range m {
		fmt.Println("[ ", key, " -> ", val, " ]")
	}

	fmt.Println("Apple price: ", m["Apple"])
	delete(m, "Apple")

	value, ok := m["Apple"]
	fmt.Println("Apple price:", value, "Present", ok)

	value2, ok2 := m["Banana"]
	fmt.Println("Banana price:", value2, "Present", ok2)
}

func example1_32(data []int) int {
	size := len(data)
	maxRN := 0
	maxEnding := 0
	for i := 0; i < size; i++ {
		maxEnding = maxEnding + data[i]
		if maxEnding < 0 {
			maxEnding = 0
		}
		if maxRN < maxEnding {
			maxRN = maxEnding
		}
	}
	return maxRN
}

func example1_32_1() {
	data := []int{1, 3, -5, -1, 0, -4, 3, 6, -293}
	fmt.Println("The max sum of a subarray is ", example1_32(data))
}

func example1_34(num int) int {
	count := 1
	for i := 1; i <= num; i++ {
		count = count * i
	}
	return count
}

func example1_37(n int) int {
	if n <= 1 {
		return n
	}
	return (n - 1) + (n - 2)
}

func example2_1(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m++
	}
	return m
}

func example2_2(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			m++
		}
	}
	return m
}

func example2_3(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			m++
		}
	}
	return m
}

func example2_4(n int) int {
	m := 0
	i := 1
	for i < n {
		m++
		i = i * 2
	}
	return m
}

func example2_5(n int) int {
	m := 0
	i := n
	for i > 0 {
		m++
		i = i / 2
	}
	return m
}

func example2_6(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < i; k++ {
				m++
			}
		}
	}
	return m
}

func example2_7(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			m++
		}
	}
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			m++
		}
	}
	return m
}

func example2_8(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		sq := math.Sqrt(float64(n))
		for j := 0; j < int(sq); j++ {
			m++
		}
	}
	return m
}

func example2_9(n int) int {
	m := 0
	for i := 0; i > 0; i /= 2 {
		for j := 0; j < i; j++ {
			m++
		}
	}
	return m
}

func example2_10(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			m++
		}
	}
	return m
}

func example2_11(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j + 1; k < n; k++ {
				m++
			}
		}
	}
	return m
}

func example2_12(n int) int {
	m := 0
	j := 0
	for i := 0; i < n; i++ {
		for ; j < n; j++ {
			m++
		}
	}
	return m
}

func example2_13(n int) int {
	m := 0
	for i := 1; i <= n; i *= 2 {
		for j := 0; j <= i; j++ {
			m++
		}
	}
	return m
}

func example2_14() {
	fmt.Println("N = 100, Number of instructions O(n):: ", example2_1(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", example2_2(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", example2_3(100))
	fmt.Println("N = 100, Number of instructions O(log(n):: ", example2_4(100))
	fmt.Println("N = 100, Number of instructions O(log(n):: ", example2_5(100))
	fmt.Println("N = 100, Number of instructions O(n^3):: ", example2_6(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", example2_7(100))
	fmt.Println("N = 100, Number of instructions O(n^3/2):: ", example2_8(100))
	fmt.Println("N = 100, Number of instructions O(log(n)):: ", example2_9(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", example2_10(100))
	fmt.Println("N = 100, Number of instructions O(n^3):: ", example2_11(100))
	fmt.Println("N = 100, Number of instructions O(n):: ", example2_12(100))
	fmt.Println("N = 100, Number of instructions O(n):: ", example2_13(100))
}

func example4_1() {
	s := stack.New()
	s.Push(2)
	s.Push(4)
	s.Push(5)
	for s.Len() != 0 {
		val := s.Pop()
		fmt.Println(val, " ")
	}
}

func example4_2() {
	q := queue.New()
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	for q.Len() != 0 {
		val := q.Dequeue()
		fmt.Print(val, " ")
	}
}

func example4_3() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50

	for key, val := range m {
		fmt.Println("[", key, "->", val, "]")
	}

	fmt.Println("Apple price: ", m["Apple"])
	delete(m, "Apple")
	fmt.Println("Apple price: ", m["Apple"])

	v, ok := m["Apple"]
	fmt.Println("Apple price:", v, "Present:", ok)

	v2, ok2 := m["Banana"]
	fmt.Println("Banana price:", v2, "Present:", ok2)
}

func example4_4() {
	st := set.New()
	st.Insert(1)
	st.Insert(2)
	fmt.Println(st.Has(1))
	fmt.Println(st.Has(3))
}

func linearSearchInArray(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func example5_2(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		} else if value < data[i] {
			return false
		}
	}
	return false
}

func example5_5(data []int) {
	size := len(data)
	fmt.Println("Repeating elements are : ")
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				fmt.Println(" ", data[i])
			}
		}
	}
	fmt.Println()
}

func example5_6(data []int) {
	size := len(data)
	sort.Ints(data)
	fmt.Println("Repeating elements are : ")

	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Println(" ", data[i])
		}
	}
	fmt.Println()
}

func example5_7(data []int) {
	s := set.New()
	size := len(data)
	fmt.Println("Repeating elements are : ")
	for i := 0; i < size; i++ {
		if s.Has(data[i]) {
			fmt.Println(" ", data[i])
		} else {
			s.Insert(data[i])
		}
	}
	fmt.Println()
}

// This function prints repeating numbers in a range
func example5_8(data []int, intrange int) {
	count := make([]int, intrange)
	//for i := 0; i < size; i++ {
	//	count[i] = 0
	//}
	for i, size := 0, len(data); i < size; i++ {
		val := data[i]
		if count[val] == 1 {
			fmt.Println(" ", val)
		} else {
			count[val]++
		}
	}
	fmt.Println()
}

func example5_9(data []int) int {
	size := len(data)
	max := data[0]
	count := 1
	maxCount := 1
	for i := 0; i < size; i++ {
		count = 1
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	return max
}

func example5_10(data []int) int {
	size := len(data)
	max := data[0]
	maxCount := 1
	curr := data[0]
	currCount := 1
	sort.Ints(data)
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			currCount++
		} else {
			currCount = 1
			curr = data[i]
		}
		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	return max
}

func example5_11(data []int, dataRange int) int {
	max := data[0]
	maxCount := 1
	size := len(data)
	count := make([]int, dataRange)
	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	return max
}

func example5_13(data []int) int {
	size := len(data)
	majIndex := size / 2
	sort.Ints(data)
	candidate := data[majIndex]
	count := 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return majIndex
	}
	// fmt.Println("MajorityDoesNotExist")
	return -1
}

func arrayMajority(data []int) (int, bool) {
	size := len(data)
	count := 0
	majority := size / 2
	sort.Ints(data)
	for i := 0; i < size-1; i++ {
		if data[i] == data[i+1] {
			count++
		} else if data[i] != data[i+1] {
			if count > majority {
				return data[i], true
			}
		}
	}
	return 0, false
}

func getMax(data []int) int {
	size := len(data)
	maxsoFar := data[0]
	for i := 1; i < size; i++ {
		if data[i] > maxsoFar {
			maxsoFar = data[i]
		}
	}
	return maxsoFar
}

func getMin(data []int) int {
	size := len(data)
	minsoFar := data[0]
	for i := 1; i < size; i++ {
		if data[i] < minsoFar {
			minsoFar = data[i]
		}
	}
	return minsoFar
}

func printRepeating(data []int) {
	size := len(data)
	sort.Ints(data)
	fmt.Println("Repeating elements are: ")

	for i := 0; i < size-1; i++ {
		if data[i] == data[i+1] {
			fmt.Println(data[i])
		}
	}
}

func equalsNumber(data []int, num int) {
	size := len(data)
	for i := 0; i < size; i++ {
		fmt.Println(num - data[i])
	}
}

func addsToNum(data []int, num int) {
	size := len(data)
	for i := 0; i < size; i++ {
		for j := 1; j < size; j++ {
			if (data[i] + data[j]) == num {
				fmt.Println("Array [", i, "] + Array [", j, "] = ", num)
			}
		}
	}
}

// return index of the majority element or -1 if none
func mooresMajorityAlgorithm(data []int) int {
	majIndex := 0
	count := 1
	size := len(data)

	for i := 1; i < size; i++ {
		if data[majIndex] == data[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			majIndex = i
			count = 1
		}
	}

	//if count > 1 {
	//	return majIndex
	//}

	candidate := data[majIndex]
	count = 0

	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		//fmt.Println(data[majIndex])
		return majIndex
	}
	//fmt.Println("There is no majority.")

	return -1
}

// return index of the majority element or -1 if none
func arrayMajorityWithMap(data []int) int {
	// map: element => count
	var m = map[int]int{}
	maxIdx := -1
	maxCount := 0
	for idx, element := range data {
		m[element]++

		if m[element] > maxCount {
			maxIdx = idx
			maxCount = m[element]
		}
	}

	if maxCount > len(data)/2 {
		return maxIdx
	}

	return -1
}

func findPair(data []int, value int) bool {
	size := len(data)
	ret := false

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == value && data[j] == value {
				fmt.Println("The pair is : ", data[i], ",", data[j])
				ret = true
			}
		}
	}
	return ret
}

func findPair2(data []int, value int) bool {
	size := len(data)
	first := 0
	second := size - 1
	ret := false
	sort.Ints(data)

	for first < second {
		curr := data[first] + data[second]
		if curr == value {
			fmt.Println("The pair is ", data[first], ",", data[second])
			ret = true
		}
		if curr < value {
			first++
		} else {
			second--
		}
	}
	return ret
}

func minAbsSumPair(data []int) {
	var sum int
	size := len(data)

	if size < 2 {
		fmt.Println("Invalid Input")
	}

	minFirst := 0
	minSecond := 1
	minSum := data[0] + data[1]
	for l := 0; l < size-1; l++ {
		for r := l + 1; r < size; r++ {
			sum = data[l] + data[r]
			if sum < minSum {
				minSum = sum
				minFirst = l
				minSecond = r
			}
		}
	}
	fmt.Println("The two elements with minimum sum are : ", data[minFirst], " , ", data[minSecond])
}

func minAbsSumPair2(data []int) {
	sort.Ints(data)
	size := len(data)
	if size < 2 {
		fmt.Println("Invalid Input")
	} else {
		fmt.Println("Valid Input")
	}
	fmt.Println("The two elements with minimum sum are : ", data[0], " , ", data[1])
}

// SearchBitonicArrayMax retturns maximum element
func SearchBitonicArrayMax(data []int) (int, bool) {
	size := len(data)
	start := 0
	end := size - 1
	mid := (start + end) / 2
	maximaFound := 0

	for start < end {
		mid := (start + end) / 2
		if data[mid-1] < data[mid] && data[mid+1] < data[mid] {
			maximaFound = 1
			break
		} else if data[mid-1] < data[mid] && data[mid] < data[mid+1] {
			start = mid + 1
		} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] {
			end = mid - 1
		} else {
			break
		}
	}
	if maximaFound == 0 {
		fmt.Println("NoMaximaFound")
		return -1, false
	}
	return mid, true
}

// SearchBitonicArray returns index of the `key` in `data` if found, otherwise -1
// we assume that `data` is a bitonic array
func SearchBitonicArray(data []int, key int) int {
	size := len(data)
	if size < 3 {
		fmt.Println("Invalid input: array must have at least 3 elements")
		return -1
	}

	bitonicCheck := isBitonicArray(data)
	if bitonicCheck == false {
		return -1
	}

	maxIndex, _ := SearchBitonicArrayMax(data)
	k := BinarySearch(data, 0, maxIndex, key, true)
	if k != -1 {
		return k
	} else {
		return BinarySearch(data, maxIndex+1, size-1, key, false)
	}
}

// BinarySearch returns ...
func BinarySearch(data []int, start int, end int, key int, isInc bool) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if key == data[mid] {
		return mid
	}
	if isInc != false && key < data[mid] || isInc == false && key > data[mid] {
		return BinarySearch(data, start, mid-1, key, isInc)
	} else {
		return BinarySearch(data, mid+1, end, key, isInc)
	}
}

//isBitonicArray returns bool if array is bitonic or not
func isBitonicArray(data []int) bool {
	size := len(data)

	// TODO check monotonic increase
	// check monotonic decrease

	i := 0
	for i < size-1 && data[i] < data[i+1] {
		i++
	}

	increased := i > 0 && i < size-1
	if !increased {
		return false
	}

	for i < size-1 && data[i] > data[i+1] {
		i++
	}

	decreased := i == size-1

	return increased && decreased
}
