package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 1. Сумма цифр числа
// Напишите программу, которая принимает целое число и вычисляет сумму его цифр.
func sum_of_digits(num int) int {
	sum := 0
	for _, digit := range strconv.Itoa(num) {
		sum += int(digit - '0')
	}
	return sum
}

// 2. Преобразование температуры
// Напишите программу, которая преобразует температуру из градусов Цельсия в Фаренгейты и обратно.
func celsius_to_fahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// 3. Удвоение каждого элемента массива
// Напишите программу, которая принимает массив чисел и возвращает новый массив, где каждое число удвоено.
func double_array(arr []int) []int {
	doubled := make([]int, len(arr))
	for i := range arr {
		doubled[i] = arr[i] * 2
	}
	return doubled
}

// 4. Объединение строк
// Напишите программу, которая принимает несколько строк и объединяет их в одну строку через пробел.
func join_strings(strings_array []string) string {
	return strings.Join(strings_array, " ")
}

// 5. Расчет расстояния между двумя точками
// Напишите программу, которая вычисляет расстояние между двумя точками в 2D пространстве.
func distance_between_points(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// 6. Проверка на четность/нечетность
// Напишите программу, которая проверяет, является ли введенное число четным или нечетным.
func check_even_odd(num int) string {
	if num%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}

// 7. Проверка высокосного года
// Напишите программу, которая проверяет, является ли введенный год високосным.
func check_leap_year(year int) string {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return "Високосный"
	}
	return "Не високосный"
}

// 8. Определение наибольшего из трех чисел
// Напишите программу, которая принимает три числа и выводит наибольшее из них.
func max_of_three(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

// 9. Категория возраста
// Напишите программу, которая принимает возраст человека и выводит, к какой возрастной группе он относится.
func age_category(age int) string {
	switch {
	case age >= 0 && age <= 12:
		return "Ребенок"
	case age >= 13 && age <= 17:
		return "Подросток"
	case age >= 18 && age <= 64:
		return "Взрослый"
	default:
		return "Пожилой"
	}
}

// 10. Проверка делимости на 3 и 5
// Напишите программу, которая проверяет, делится ли число одновременно на 3 и 5.
func check_divisibility(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "Делится"
	}
	return "Не делится"
}

// 11. Факториал числа
// Напишите программу, которая вычисляет факториал числа.
func factorial(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}

// 12. Числа Фибоначчи
// Напишите программу, которая выводит первые n чисел Фибоначчи.
func fibonacci(n int) []int {
	fib := make([]int, n)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fib[i] = a
		a, b = b, a+b
	}
	return fib
}

// 13. Реверс массива
// Напишите программу, которая переворачивает массив чисел.
func reverse_array(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// 14. Поиск простых чисел
// Напишите программу, которая выводит все простые числа до заданного числа.
func prime_numbers(n int) []int {
	var primes []int
	for num := 2; num <= n; num++ {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num)
		}
	}
	return primes
}

// 15. Сумма чисел в массиве
// Напишите программу, которая вычисляет сумму всех чисел в массиве.
func sum_array(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println("Сумма цифр числа 1234:", sum_of_digits(1234))
	fmt.Println("25 градусов Цельсия в Фаренгейтах:", celsius_to_fahrenheit(25))
	fmt.Println("Удвоение массива [1, 2, 3, 4]:", double_array([]int{1, 2, 3, 4}))
	fmt.Println("Объединение строк ['Hello', 'world']:", join_strings([]string{"Hello", "world"}))
	fmt.Println("Расстояние между точками (1, 1) и (4, 5):", distance_between_points(1, 1, 4, 5))
	fmt.Println("Число 4 является:", check_even_odd(4))
	fmt.Println("Год 2020:", check_leap_year(2020))
	fmt.Println("Наибольшее из трех чисел (4, 9, 7):", max_of_three(4, 9, 7))
	fmt.Println("Возраст 25:", age_category(25))
	fmt.Println("Число 15:", check_divisibility(15))
	fmt.Println("Факториал числа 5:", factorial(5))
	fmt.Println("Первые 7 чисел Фибоначчи:", fibonacci(7))
	fmt.Println("Реверс массива [1, 2, 3, 4, 5]:", reverse_array([]int{1, 2, 3, 4, 5}))
	fmt.Println("Простые числа до 20:", prime_numbers(20))
	fmt.Println("Сумма массива [1, 2, 3, 4, 5]:", sum_array([]int{1, 2, 3, 4, 5}))
}
