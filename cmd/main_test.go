package main

import (
	"testing"
)

type NormalizeTestStruct struct {
	input string
	want  string
}

func TestNormalize(t *testing.T) {
	testcases := []NormalizeTestStruct{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			actual := Normalize(tc.input)
			want := tc.want

			if actual != want {
				t.Errorf("Want %s got %s", want, actual)
			}

		})
	}

}
