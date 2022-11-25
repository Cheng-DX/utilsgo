package utilsgo

import (
	"fmt"
	"testing"
)

var collection = []int{1, 2, 3, 4}

func TestEvery(t *testing.T) {
	r := Every(collection, func(item int) bool {
		return item < 0
	})
	if !r {
		t.Errorf("Every should return true")
	}
}

func TestSome(t *testing.T) {
	r := Some(collection, func(item int) bool {
		return item < 5
	})
	if !r {
		t.Errorf("Some should return true")
	}
}

func TestJoin(t *testing.T) {
	r := Join(collection, ",")
	if r != "1,2,3,4" {
		t.Errorf("Join should return '1,2,3,4'")
	}
}

func TestSlice(t *testing.T) {
	fmt.Println(Slice(collection, 0, 4))
}

func TestPop(t *testing.T) {
	c, r := Pop(collection)
	fmt.Println(r, c)
}

func TestShift(t *testing.T) {
	c, r := Shift(collection)
	fmt.Println(c, r)
}

func TestReverse(t *testing.T) {
	c := Reverse(collection)
	fmt.Println(c)
}
