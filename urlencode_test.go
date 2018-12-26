package main

import (
	"fmt"
	"testing"
)

func TestUrlencode(t *testing.T) {
	var tests = []struct {
		in       string
		expected string
	}{
		{"http://www.hp.com", "http://www.hp.com"},
		{"www", "www"},
		{"", ""},
		{"http://www.hp.com/path/to/target", "http://www.hp.com/path/to/target"},
		{"http://www.hp.com/path to target", "http://www.hp.com/path%20to%20target"},
		{"http://borax.gre.hp.com/canna/registry/ca/CANNA Root Certificate Authority/cert/1f541a81/cert.pem",
		 "http://borax.gre.hp.com/canna/registry/ca/CANNA%20Root%20Certificate%20Authority/cert/1f541a81/cert.pem"},
	}
	for _, test := range tests {
		got := urlencode(test.in)
		if got != test.expected {
			t.Errorf("From %q, expected %q, but got %q", test.in, test.expected, got)
		}
	}
}

func ExampleUrlencode() {
	fmt.Println(urlencode("http://borax.fra.hp.com/canna/registry/CANNA Root Certificate Authority/"))
	// Output:
	// http://borax.fra.hp.com/canna/registry/CANNA%20Root%20Certificate%20Authority/
}
