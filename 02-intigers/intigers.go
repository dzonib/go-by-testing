package integers

import "fmt"

func main() {

}

func Add(i, i2 int) int {
	return i + i2
}

// Please note that the example function will not be executed if you remove the comment
// Output: 6. Although the function will be compiled, it won't be executed.

// By adding this code the example will appear in the documentation inside godoc, making your code even more accessible.

// To try this out, run godoc -http=:6060 and navigate to http://localhost:6060/pkg/

//Inside here you'll see a list of all the packages and you'll be able to find your example documentation.

//If you publish your code with examples to a public URL, you can share the documentation of your code at pkg.go.net

//This web interface allows you to search for documentation of standard library packages and third-party packages.
//EXAMPLE: https://pkg.go.dev/github.com/quii/learn-go-with-tests/integers/v2

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
