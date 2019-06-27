package main

import "fmt"
import "math"

func main() {

	choice := 99
	switch choice {
	case 1:
		example1_1()
	case 2:
		example1_2()
	case 3:
		example1_3()
	case 4:
		example1_4()
	case 5:
		example1_5(5, 6)
	case 6:
		example1_6(3, 9)
	case 7:
		example1_7(3, 88, 300)
	case 8:
		example1_8()
	case 9:
		example1_9()
	case 10:
		example1_10(4)
	case 11:
		example1_11()
	case 12:
		example1_12()
	case 13:
		example1_13()
	case 14:
		example1_14()
	case 15:
		example1_15()
	case 16:
		example1_16_1()
	case 17:
		example1_17_1()
	case 18:
		example1_18()
	case 19:
		example1_19_1()
	case 20:
		example1_20()
	case 21:
		example1_21()
	case 22:
		example1_22()
	case 23:
		fmt.Println("bruh moment")
	case 99:
		test_sorting()
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
