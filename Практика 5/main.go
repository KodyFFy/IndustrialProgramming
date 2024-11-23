package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 1. Перевод чисел из одной системы счисления в другую
func convert_base(numStr string, fromBase int, toBase int) (string, error) {
	num, err := strconv.ParseInt(numStr, fromBase, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(num, toBase), nil
}

// 2. Решение квадратного уравнения
func solve_quadratic(a, b, c float64) (complex128, complex128) {
	d := b*b - 4*a*c
	if d > 0 {
		root1 := complex((-b+math.Sqrt(d))/(2*a), 0)
		root2 := complex((-b-math.Sqrt(d))/(2*a), 0)
		return root1, root2
	} else if d == 0 {
		root := complex(-b/(2*a), 0)
		return root, root
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-d) / (2 * a)
		return complex(realPart, imaginaryPart), complex(realPart, -imaginaryPart)
	}
}

// 3. Сортировка чисел по модулю
func sort_by_absolute(arr []int) []int {
	sortSlice := make([]int, len(arr))
	copy(sortSlice, arr)
	for i := 0; i < len(sortSlice)-1; i++ {
		for j := i + 1; j < len(sortSlice); j++ {
			if abs(sortSlice[i]) > abs(sortSlice[j]) {
				sortSlice[i], sortSlice[j] = sortSlice[j], sortSlice[i]
			}
		}
	}
	return sortSlice
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 4. Слияние двух отсортированных массивов
func merge_sorted_arrays(arr1, arr2 []int) []int {
	merged := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}
	merged = append(merged, arr1[i:]...)
	merged = append(merged, arr2[j:]...)
	return merged
}

// 5. Нахождение подстроки в строке без использования встроенных функций
func find_substring(haystack, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)

	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}

// Условия

// 1. Калькулятор с расширенными операциями
func calculator(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	case "^":
		return math.Pow(a, b), nil
	case "%":
		return math.Mod(a, b), nil
	default:
		return 0, fmt.Errorf("недопустимая операция")
	}
}

// 2. Проверка палиндрома
func is_palindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 3. Нахождение пересечения трех отрезков
func segments_intersection(seg1, seg2, seg3 [2]int) bool {
	start := max(max(seg1[0], seg2[0]), seg3[0])
	end := min(min(seg1[1], seg2[1]), seg3[1])
	return start <= end
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 4. Выбор самого длинного слова в предложении
func longest_word(sentence string) string {
	words := strings.Fields(sentence)
	longest := ""
	for _, word := range words {
		word = strings.Trim(word, ",.!?;:")
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

// 5. Проверка высокосного года
func is_leap_year(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// Циклы

// 1. Числа Фибоначчи до определенного числа
func fibonacci_up_to(n int) []int {
	fib := []int{0, 1}
	for {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		if next > n {
			break
		}
		fib = append(fib, next)
	}
	return fib
}

// 2. Определение простых чисел в диапазоне
func prime_numbers_in_range(start, end int) []int {
	var primes []int
	for num := start; num <= end; num++ {
		if is_prime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

func is_prime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 3. Числа Армстронга в заданном диапазоне
func armstrong_numbers_in_range(start, end int) []int {
	var armstrongs []int
	for num := start; num <= end; num++ {
		if is_armstrong(num) {
			armstrongs = append(armstrongs, num)
		}
	}
	return armstrongs
}

func is_armstrong(num int) bool {
	sum := 0
	digits := strconv.Itoa(num)
	length := len(digits)
	for _, digit := range digits {
		d, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(d), float64(length)))
	}
	return sum == num
}

// 4. Реверс строки
func reverse_string(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 5. Нахождение наибольшего общего делителя (НОД)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	result, err := convert_base("101", 2, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Перевод числа '101' из двоичной системы в десятичную:", result)
	}
	root1, root2 := solve_quadratic(1, -3, 2)
	fmt.Println("Корни квадратного уравнения 1x^2 - 3x + 2:", root1, root2)
	fmt.Println("Сортировка по модулю:", sort_by_absolute([]int{-3, 1, -2, 4}))
	fmt.Println("Слияние массивов [1, 3, 5] и [2, 4, 6]:", merge_sorted_arrays([]int{1, 3, 5}, []int{2, 4, 6}))
	fmt.Println("Нахождение подстроки 'world' в 'Hello world':", find_substring("Hello world", "world"))
	resultCalc1, err := calculator(10, 2, "/")
	if err != nil {
		fmt.Println("Результат калькулятора: ошибка", err)
	} else {
		fmt.Println("Результат калькулятора:", resultCalc1)
	}
	resultCalc2, err := calculator(10, 0, "/")
	if err != nil {
		fmt.Println("Результат калькулятора: ошибка", err)
	} else {
		fmt.Println("Результат калькулятора:", resultCalc2)
	}
	fmt.Println("Проверка палиндрома 'A man, a plan, a canal, Panama':", is_palindrome("A man, a plan, a canal, Panama"))
	fmt.Println("Пересечение отрезков [1, 5], [2, 6], [3, 4 ]:", segments_intersection([2]int{1, 5}, [2]int{2, 6}, [2]int{3, 4}))
	fmt.Println("Самое длинное слово в 'Hello, world!':", longest_word("Hello, world!"))
	fmt.Println("Год 2024 високосный:", is_leap_year(2024))
	fmt.Println("Числа Фибоначчи до 10:", fibonacci_up_to(10))
	fmt.Println("Простые числа в диапазоне от 10 до 30:", prime_numbers_in_range(10, 30))
	fmt.Println("Числа Армстронга в диапазоне от 1 до 1000:", armstrong_numbers_in_range(1, 1000))
	fmt.Println("Реверс строки 'Hello':", reverse_string("Hello"))
	fmt.Println("НОД чисел 48 и 18:", gcd(48, 18))
}
