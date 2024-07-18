package slice

import (
	"reflect"
	"testing"
)

func TestFirst(t *testing.T) {
	testCases := []struct {
		slice []int
		want  int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{0, -1, -2}, 0},
		{[]int{5}, 5},
	}

	for _, tc := range testCases {
		got := First(tc.slice)
		if got != tc.want {
			t.Errorf("First(%v) = %v, want %v", tc.slice, got, tc.want)
		}
	}
}
func TestRest(t *testing.T) {
	testCases := []struct {
		slice []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{2, 3}},
		{[]int{0, -1, -2}, []int{-1, -2}},
		{[]int{5}, []int{}},
	}

	for _, tc := range testCases {
		got := Rest(tc.slice)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Rest(%v) = %v, want %v", tc.slice, got, tc.want)
		}
	}
}
func TestLast(t *testing.T) {
	testCases := []struct {
		slice []int
		want  int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{0, -1, -2}, -2},
		{[]int{5}, 5},
	}

	for _, tc := range testCases {
		got := Last(tc.slice)
		if got != tc.want {
			t.Errorf("Last(%v) = %v, want %v", tc.slice, got, tc.want)
		}
	}
}
func TestReverse(t *testing.T) {
	testCases := []struct {
		slice []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{0, -1, -2}, []int{-2, -1, 0}},
		{[]int{5}, []int{5}},
	}

	for _, tc := range testCases {
		got := Reverse(tc.slice)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Reverse(%v) = %v, want %v", tc.slice, got, tc.want)
		}
	}
}
func TestIndexOf(t *testing.T) {
	testCases := []struct {
		slice []int
		el    int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{0, -1, -2, -3}, -2, 2},
		{[]int{5, 10, 15, 20}, 25, -1},
	}

	for _, tc := range testCases {
		got := IndexOf(tc.slice, tc.el)
		if got != tc.want {
			t.Errorf("IndexOf(%v, %v) = %v, want %v", tc.slice, tc.el, got, tc.want)
		}
	}
}
func TestUnique(t *testing.T) {
	testCases := []struct {
		slice []int
		want  []int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
		{[]int{0, -1, -2, 0, -1}, []int{0, -1, -2}},
		{[]int{5, 5, 5, 5}, []int{5}},
	}

	for _, tc := range testCases {
		got := Unique(tc.slice)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Unique(%v) = %v, want %v", tc.slice, got, tc.want)
		}
	}
}
func TestMap(t *testing.T) {
	testCases := []struct {
		slice []int
		fnc   func(int) int
		want  []int
	}{
		{[]int{1, 2, 3}, func(x int) int { return x * 2 }, []int{2, 4, 6}},
		{[]int{0, -1, -2}, func(x int) int { return x + 1 }, []int{1, 0, -1}},
		{[]int{5}, func(x int) int { return x * x }, []int{25}},
	}

	for _, tc := range testCases {
		got := Map(tc.slice, tc.fnc)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Map(%v) = %v, want %v", tc.slice, got, tc.want)
		}
	}
}
func TestFilter(t *testing.T) {
	testCases := []struct {
		slice []int
		f     func(int) bool
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 0 }, []int{2, 4}},
		{[]int{0, -1, -2, -3}, func(x int) bool { return x > 0 }, []int{}},
		{[]int{5, 10, 15, 20}, func(x int) bool { return x < 10 }, []int{5}},
	}

	for _, tc := range testCases {
		got := Filter(tc.slice, tc.f)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Filter(%v) = %v, want %v", tc.slice, got, tc.want)
		}
	}
}
func TestSliceAfter(t *testing.T) {
	testCases := []struct {
		slice   []int
		element int
		want    []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{4, 5}},
		{[]int{0, -1, -2, -3}, -2, []int{-3}},
		{[]int{5, 10, 15, 20}, 25, []int{}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{}},
	}

	for _, tc := range testCases {
		got := SliceAfter(tc.slice, tc.element)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("SliceAfter(%v, %v) = %v, want %v", tc.slice, tc.element, got, tc.want)
		}
	}
}
func TestIntersect(t *testing.T) {
	testCases := []struct {
		slice1 []int
		slice2 []int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{4, 5}},
		{[]int{0, -1, -2, -3}, []int{-2, -3, -4, -5}, []int{-2, -3}},
		{[]int{5, 10, 15, 20}, []int{25, 30, 35, 40}, []int{}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{[]int{11, 12, 13}, []int{13, 5, 11}, []int{11, 13}},
	}

	for _, tc := range testCases {
		got := Intersect(tc.slice1, tc.slice2)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Intersect(%v, %v) = %v, want %v", tc.slice1, tc.slice2, got, tc.want)
		}
	}
}
func TestChunkBy(t *testing.T) {
	testCases := []struct {
		items     []int
		chunkSize int
		want      [][]int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}},
	}

	for _, tc := range testCases {
		got := ChunkBy(tc.items, tc.chunkSize)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ChunkBy(%v, %v) = %v, want %v", tc.items, tc.chunkSize, got, tc.want)
		}
	}
}
func TestRandomElement(t *testing.T) {
	testCases := []struct {
		slice []int
	}{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{0, -1, -2, -3}},
		{[]int{5, 10, 15, 20}},
	}

	for _, tc := range testCases {
		got := RandomElement(tc.slice)
		if !Contains(tc.slice, got) {
			t.Errorf("RandomElement(%v) = %v, not found in the original slice", tc.slice, got)
		}
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		slice   []int
		element int
		want    bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{0, -1, -2, -3}, -2, true},
		{[]int{5, 10, 15, 20}, 25, false},
		{[]int{1, 2, 3, 4, 5}, 6, false},
	}

	for _, tc := range testCases {
		got := Contains(tc.slice, tc.element)
		if got != tc.want {
			t.Errorf("Contains(%v, %v) = %v, want %v", tc.slice, tc.element, got, tc.want)
		}
	}
}
