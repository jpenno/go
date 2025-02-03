package main

import (
	"fmt"
	"strings"
)

func StringToNumber(str string) int {
	i := 0
	min := 1
	if str[0] == '-' {
		min = -1
		i++
	}

	r := 0

	for ; i < len(str); i++ {
		r *= 10
		r += int(str[i] - '0')
	}

	return r * min
}
func RepeatStr(repetitions int, value string) string {
	var str string
	for range repetitions {
		str += value
	}
	return str
}

func CountBy(x, n int) []int {
	list := make([]int, n)
	// for i := 1; i <= n; i++ {
	// 	fmt.Printf("%v * %v = %v\n", x, i, x*i)
	// 	list[i-1] = x * i
	// }
	for i := range list {
		list[i] = x * (i + 1)
	}
	return list
}

func Digitize(n int) []int {
	var ret []int
	for {
		fmt.Printf("%v mod 10 = %v\n", n, n%10)
		ret = append(ret, n%10)
		n /= 10
		if n == 0 {
			return ret
		}
	}
}

func AbbrevName(name string) string {
	names := strings.Split(name, " ")
	return strings.ToUpper(
		fmt.Sprintf(
			"%v.%v",
			names[0][:1],
			names[1][:1],
		),
	)
}

func SmallestIntegerFinder(numbers []int) int {
	lowest := numbers[0]

	for _, numbre := range numbers {
		if lowest > numbre {
			lowest = numbre
		}
	}

	return lowest
}

func DNAtoRNA(dna string) string {
	var response string

	for _, letter := range dna {
		if letter == 'T' {
			letter = 'U'
		}
		response += string(letter)
	}

	return response
}

func Goals(goals ...int) (res int) {
	for _, i := range goals {
		res += i
	}
	return res
}

func ZeroFuel(distanceToPump int, mpg int, fuelLeft int) bool {
	return mpg*fuelLeft >= distanceToPump
}

func SpinWords(str string) string {
	s := strings.Split(str, " ")

	for i, v := range s {
		if len(v) >= 5 {
			res := ""
			for _, r := range v {
				res = string(r) + res
			}
			s[i] = res
		}
	}

	return strings.Join(s, " ")
}

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func Divisors(n int) int {
	//Put your code here
	var divisors_count int
	for i := 1; i <= n; i++ {
		fmt.Printf("i: %v, n: %v\n", i, n)
		if n%i == 0 {
			divisors_count++
		}
	}

	return divisors_count
}

type MyString string

func (s MyString) IsUpperCase() bool {
	// Your code here!
	for _, char := range s {
		if char >= 'a' {
			fmt.Printf("char: %v\n", string(char))
			return false
		}
	}
	return true
}

func main() {
	test := MyString.IsUpperCase("a")
	test2 := MyString.IsUpperCase("ABCDEFGHIJKLMNOPQRSTUVWXYz")

	fmt.Printf("test a: %v\n", test)
	fmt.Printf("test A: %v\n", test2)
	fmt.Printf("a: %v, g: %v\n", 'a', 'g')
}
