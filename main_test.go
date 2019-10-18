package main

import (
	c "countWebService/count"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCount(t *testing.T) {
	// test is the input is nil
	t.Run("Nil input", func(t *testing.T) {
		input := ""
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 0,"
		expectedVowels := "Huruf hidup: 0"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})

	// test is the input is no vowels
	t.Run("No vowels input", func(t *testing.T) {
		input := "thsnptsnvwls"
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 12,"
		expectedVowels := "Huruf hidup: 0"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})

	// test is the input is no vowels
	t.Run("No vowels input", func(t *testing.T) {
		input := "thsnptsnvwls"
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 12,"
		expectedVowels := "Huruf hidup: 0"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})

	// test is the input is no consonants
	t.Run("No vowels input", func(t *testing.T) {
		input := "aiueo"
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 0,"
		expectedVowels := "Huruf hidup: 5"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})

	main()
}

func TestServer(t *testing.T) {
	// test for homepage
	t.Run("Homepage", func(t *testing.T) {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll // pass 'nil' as the third parameter.
		req, err := http.NewRequest("GET", "/", nil)

		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(homepage)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := rr.Body.String()

		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	// test for count page
	t.Run("Count page", func(t *testing.T) {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll // pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/count", nil)

		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(countController)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Check the response body is what we expect.
		input := ""
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 0,"
		expectedVowels := "Huruf hidup: 0"

		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "aiueo"
		c.Count(input)
	}
}
