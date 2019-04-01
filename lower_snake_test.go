package strcase_test

import (
	"fmt"
	"testing"

	"github.com/hori-ryota/go-strcase"
	"github.com/stretchr/testify/assert"
)

func TestToLowerSnake(t *testing.T) {
	for _, tt := range testPatterns {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.ls, strcase.ToLowerSnake(tt.s))
		})
	}
}

func ExampleToLowerSnake() {
	fmt.Println(strcase.ToLowerSnake("lowerCamel"))
	fmt.Println(strcase.ToLowerSnake("UpperCamel"))
	fmt.Println(strcase.ToLowerSnake("lower_snake"))
	fmt.Println(strcase.ToLowerSnake("UPPER_SNAKE"))

	// Output:
	// lower_camel
	// upper_camel
	// lower_snake
	// upper_snake
}
