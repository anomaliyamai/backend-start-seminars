package unittests

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"both positive", 2, 3, 5},
		{"positive and zero", 4, 0, 4},
		{"both negative", -2, -3, -5},
		{"negative and positive", -1, 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func FuzzDiv(f *testing.F) {
	f.Add(10, 2)
	f.Add(0, 1)
	f.Add(5, -1)

	f.Fuzz(func(t *testing.T, a int, b int) {
		result, err := Div(a, b)
		if b == 0 {
			if err == nil {
				t.Errorf("Expected error for division by zero, got result %d", result)
			}
			return
		}
		if err != nil {
			t.Errorf("Unexpected error for Div(%d,%d): %v", a, b, err)
		}
		if a != b*result+(a%b) {
			t.Errorf("Invariant failed: %d != %d*%d + %d", a, b, result, a%b)
		}
	})
}

func SumLoop(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumRec(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + SumRec(nums[1:])
}

func BenchmarkSum(b *testing.B) {
	cases := []struct {
		name string
		nums []int
	}{
		{"small", []int{1, 2, 3, 4, 5}},
		{"medium", makeRange(1, 100)},
		{"large", makeRange(1, 1000)},
	}

	for _, tt := range cases {
		b.Run(tt.name+"/loop", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = SumLoop(tt.nums)
			}
		})
		b.Run(tt.name+"/rec", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = SumRec(tt.nums)
			}
		})
	}
}

func makeRange(start, end int) []int {
	nums := make([]int, end-start+1)
	for i := range nums {
		nums[i] = start + i
	}
	return nums
}
