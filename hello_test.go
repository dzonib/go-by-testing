package main

import "testing"

//Write a test
//Make the compiler pass
//Run the test, see that it fails and check the error message is meaningful
//Write enough code to make the test pass
//Refactor

func TestHello(t *testing.T) {
	//	got := Hello("King")
	//
	//	want := "Yo, what's up King"
	//
	//	if got != want {
	//		t.Errorf("got %q, want %q", got, want)
	//	}

	// adding subtests
	t.Run("say greeting plus provided name", func(t *testing.T) {
		got := Hello("King", "EN")

		want := "Hello King"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say default greeting if name is empty", func(t *testing.T) {
		got := Hello("", "EN")

		want := "Hello Sir/Madam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greeting in Spanish", func(t *testing.T) {
		got := Hello("Kingo", "Spanish")

		want := "Hola Kingo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greeting in French", func(t *testing.T) {
		got := Hello("Kongoo", "French")

		want := "YO Kongoo"

		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to
//accept a testing.TB which is an interface that *testing.T and
//*testing.B both satisfy, so you can call helper functions from a test, or a benchmark
func assertCorrectMessage(t testing.TB, got, want string) {
	//	t.Helper() is needed to tell the test suite that this method is a helper.
	//By doing this when it fails the line number reported will be in our function call rather than inside our test helper. 
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
