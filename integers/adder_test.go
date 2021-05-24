package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4
	if sum != expect {
		t.Errorf("expect %d, got %d", expect, sum)
	}
}

func ExampleAdd() {
	// Output: 6
	fmt.Println(Add(1, 5))
}
