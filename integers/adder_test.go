package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

// Testable examples - The comment in the last line is required to make the function run when the tests are run
// Docs can be viewed by running pkgsite -open .

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

	// Output: 6
}
