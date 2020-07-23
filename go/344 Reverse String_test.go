package main

import (
	"fmt"
	"testing"
)

func generateReverseString() []struct{ input, want []byte } {

	return nil
}

func BenchmarkReverseString(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var tests = []struct {
			input, want []byte
		}{
			{[]byte{'1', '2', '3', '4'}, []byte{'4', '3', '2', '1'}},
		}
		for _, v := range tests {
			echo := fmt.Sprintf("input:%s", v.input)
			ReverseString(v.input)
			for index, vv := range v.want {
				if vv != v.input[index] {
					fmt.Printf("%s, want:%s, ret:%s\n", echo, v.want, v.input)
				}
			}
		}
	}
}

func TestReverseString(t *testing.T) {

	var tests = []struct {
		input, want []byte
	}{
		{[]byte{'1', '2', '3', '4'}, []byte{'4', '3', '2', '1'}},
	}

	for _, v := range tests {
		echo := fmt.Sprintf("input:%s", v.input)
		ReverseString(v.input)
		for index, vv := range v.want {
			if vv != v.input[index] {
				fmt.Printf("%s, want:%s, ret:%s\n", echo, v.want, v.input)
			}
		}
	}
}

func TestReverseStringV2(t *testing.T) {
	var tests = []struct {
		input, want []byte
	}{
		{[]byte{'1', '2', '3', '4'}, []byte{'4', '3', '2', '1'}},
	}

	for _, v := range tests {
		echo := fmt.Sprintf("input:%s", v.input)
		ReverseStringV2(v.input)
		for index, vv := range v.want {
			if vv != v.input[index] {
				fmt.Printf("%s, want:%s, ret:%s\n", echo, v.want, v.input)
			}
		}
	}
}

func TestReverseStringV3(t *testing.T) {
	var tests = []struct {
		input, want []byte
	}{
		{[]byte{'1', '2', '3', '4'}, []byte{'4', '3', '2', '1'}},
	}

	for _, v := range tests {
		echo := fmt.Sprintf("input:%s", v.input)
		ReverseStringV3(v.input)
		for index, vv := range v.want {
			if vv != v.input[index] {
				fmt.Printf("%s, want:%s, ret:%s\n", echo, v.want, v.input)
			}
		}
	}
}
