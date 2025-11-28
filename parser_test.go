package growel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRouterSplitPath(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"/", []string{""}},
		{"/hello", []string{"hello"}},
		{"/hello/world", []string{"hello", "world"}},
		{"//hello//world//", []string{"hello", "", "world"}},
		{"/hello?world=true", []string{"hello?world=true"}},
	}

	for _, tt := range tests {
		got := splitPath(tt.input)
		if len(got) != len(tt.want) {
			t.Fatalf("splitPath(%q) length mismatch: got %v want %v", tt.input, got, tt.want)
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Fatalf("splitPath(%q) mismatch: got %v want %v", tt.input, got, tt.want)
			}
		}
	}
}

func TestExtractQueryStr(t *testing.T) {
	tests := []struct {
		input string
		want1  string
		want2 string
	}{
		{"", "", ""},
		{"hello", "hello", ""},
		{"user/23/", "user/23/", ""},
		{"api?str=123&yrs=23", "api", "str=123&yrs=23"},
		{"w?st?r", "w", "st"},
		{"w?str", "w", "str"},
	}

	for _, tt := range tests {
		got1, got2 := extractQueryStr(tt.input)
		if got1 != tt.want1 && got2 != tt.want2 {
			fmt.Printf("want1: %s, got1: %s\n", tt.want1, got1)
			fmt.Printf("want2: %s, got2: %s\n", tt.want2, got2)
			t.Fail()
		}
	}
}

func TestParseQueryStr(t *testing.T) {
	tests := []struct {
		input string
		want  map[string]string
	}{
		{"str=123&yrs=23", map[string]string{
			"str": "123",
			"yrs": "23",
		}},
		{"str=abc", map[string]string{
			"str": "abc",
		}},
	}

	for _, tt := range tests {
		got := parseQueryStr(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("want: %s, got: %s\n", tt.want, got)
			t.Fail()
		}
	}
}
