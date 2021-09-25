package main

import (
	"testing"
	"time"
)

//simple test
func TestGreet(t *testing.T) {
	actual := greet("abhishek")
	expected := "hello, abhishek!"
	if actual != expected {
		t.Errorf("expected %s, but was %s", expected, actual)
	}
}

func TestGreet_(t *testing.T) {
	actual := greet("")
	expected := "hello, there!"
	if actual != expected {
		t.Errorf("expected %s, but was %s", expected, actual)
	}
}

//sub-tests
func TestGreet2(t *testing.T) {
	t.Run("test blank value", func(te *testing.T) {
		actual := greet("")
		expected := "hello, there!"
		if actual != expected {
			te.Errorf("expected %s, but was %s", expected, actual)
		}
	})

	t.Run("test valid value", func(te *testing.T) {
		actual := greet("abhishek")
		expected := "hello, abhishek!"
		if actual != expected {
			te.Errorf("expected %s, but was %s", expected, actual)
		}
	})
}

//table driven tests
func TestGreet3(t *testing.T) {
	type testCase struct {
		name             string
		input            string
		expectedGreeting string
	}

	testCases := []testCase{
		{name: "test blank value", input: "", expectedGreeting: "hello, there!"},
		{name: "test valid value", input: "abhishek", expectedGreeting: "hello, abhishek!"},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(te *testing.T) {
			actual := greet(test.input)
			expected := test.expectedGreeting
			if actual != expected {
				te.Errorf("expected %s, but was %s", expected, actual)
			}
		})
	}
}

//parallel tests
func TestGreet4(t *testing.T) {
	type testCase struct {
		name             string
		input            string
		expectedGreeting string
	}

	testCases := []testCase{
		{name: "test blank value", input: "", expectedGreeting: "hello, there!"},
		{name: "test valid value", input: "abhishek", expectedGreeting: "hello, abhishek!"},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(te *testing.T) {
			te.Parallel()
			//simulating time taking test: total time for test execution little > 3s which means tests ran parallely

			time.Sleep(3 * time.Second)
			actual := greet(test.input)
			expected := test.expectedGreeting
			if actual != expected {
				te.Errorf("expected %s, but was %s", expected, actual)
			}
		})
	}
}
