package strcase_test

import (
	"fmt"
	"testing"

	"github.com/hori-ryota/go-strcase"
	"github.com/stretchr/testify/assert"
)

func TestToUpperSnake(t *testing.T) {
	for _, tt := range testPatterns {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.us, strcase.ToUpperSnake(tt.s))
		})
	}
}

func ExampleToUpperSnake() {
	fmt.Println(strcase.ToUpperSnake("lowerCamel"))
	fmt.Println(strcase.ToUpperSnake("UpperCamel"))
	fmt.Println(strcase.ToUpperSnake("lower_snake"))
	fmt.Println(strcase.ToUpperSnake("UPPER_SNAKE"))

	// Output:
	// LOWER_CAMEL
	// UPPER_CAMEL
	// LOWER_SNAKE
	// UPPER_SNAKE
}
