package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функия, создающая случайный срез чисел длиной 10 элементов.
func sliceCreate() []int {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	numbers := make([]int, 10)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = random.Intn(100)
	}
	return numbers
}

// Функция, которая принимает на вход срез чисел, проверяет их на чётные\нечётные
// и возвращает два среза с полученными значениями этой проверки.
func evenOdd(slice []int) ([]int, []int) {
	var even []int
	var odd []int

	for _, v := range slice {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return even, odd
}

func main() {
	currentSlice := sliceCreate()
	even, odd := evenOdd(currentSlice)
	fmt.Printf("Изначальный срез чисел: %v\n", currentSlice)
	fmt.Printf("Срез чётных чисел %v\n", even)
	fmt.Printf("Срез нечётных чисел: %v\n", odd)
}
