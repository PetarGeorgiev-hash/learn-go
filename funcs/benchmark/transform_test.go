package bench

import (
	"testing"
)

func transformNumbersCopy(numbers []int, transform func(int) int) []int {
	result := make([]int, len(numbers))
	for i, val := range numbers {
		result[i] = transform(val)
	}
	return result
}

func transformNumbersInPlace(numbers []int, transform func(int) int) {
	for i, val := range numbers {
		numbers[i] = transform(val)
	}
}

func double(n int) int {
	return n * 2
}

// Benchmarks
func BenchmarkTransformCopy(b *testing.B) {
	numbers := make([]int, 1000000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transformNumbersCopy(numbers, double)
	}
}

func BenchmarkTransformInPlace(b *testing.B) {
	numbers := make([]int, 1000000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Copy input so original isn't permanently mutated
		nums := make([]int, len(numbers))
		copy(nums, numbers)
		transformNumbersInPlace(nums, double)
	}
}
