package iteration

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	repeated := Repeat("a", 5)

	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 500)
	}
}

// it is good practice to add Example test for docs
// To try this out, run godoc -http=:6060 and navigate to http://localhost:6060/pkg/

// test will fail if output is not correct
func ExampleRepeat() {
	repeated := Repeat("a", 3)

	fmt.Println(repeated)

	// Output: aaa
}
