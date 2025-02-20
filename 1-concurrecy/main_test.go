package main

import (
	"testing"
)

func TestGenerateNumbers(t *testing.T) {
	ch := make(chan int)
	go generateNumbers(ch)

	// Читаем числа из канала и проверяем их количество
	count := 0
	for range ch {
		count++
	}

	// Ожидаем, что будет сгенерировано 10 чисел
	if count != 10 {
		t.Errorf("Expected 10 numbers, got %d", count)
	}
}

func TestSquareNumbers(t *testing.T) {
	inCh := make(chan int)
	outCh := make(chan int)

	go squareNumbers(inCh, outCh)

	// Отправляем тестовые числа в канал
	testNumbers := []int{2, 3, 4}
	expectedSquares := []int{4, 9, 16}

	go func() {
		for _, num := range testNumbers {
			inCh <- num
		}
		close(inCh)
	}()

	// Читаем квадраты из выходного канала и проверяем их
	for i := 0; i < len(testNumbers); i++ {
		squared := <-outCh
		if squared != expectedSquares[i] {
			t.Errorf("Expected %d, got %d", expectedSquares[i], squared)
		}
	}
}
