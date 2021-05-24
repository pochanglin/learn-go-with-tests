package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	// t.Run("collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}
	// 	got := Sum(numbers)
	// 	expect := 15
	// 	if got != expect {
	// 		t.Errorf("expect %d, got %d, given %v", expect, got, numbers)
	// 	}
	// })

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expect := 6
		if got != expect {
			t.Errorf("expect %d, got %d, given %v", expect, got, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3, 5}, []int{2, 4, 6})
	want := []int{9, 12}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 3, 5}, []int{2, 4, 6})
		want := []int{8, 10}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but want %v", got, want)
		}
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 4, 6})
		want := []int{0, 10}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but want %v", got, want)
		}
	})

}
