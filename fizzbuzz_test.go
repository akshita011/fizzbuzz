package fizzbuzz_test

import (
	"testing"

	"fmt"

	"github.com/akshita011/fizzbuzz"
)

func TestFizz(t *testing.T) {
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("The value %d should not have returned true", 1)
		t.Fail()
	}
	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("The value %d should not have returned true", 3)
		t.Fail()
	}
	if result != "fizz" {
		t.Logf("Result should have been fizz")
		t.Fail()
	}

}

func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	// Output: 3
}
