package strcase_test

import (
	"fmt"
	"testing"

	"github.com/hori-ryota/go-strcase"
	"github.com/stretchr/testify/assert"
)

func TestToUpperCamel(t *testing.T) {
	for _, tt := range testPatterns {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.uc, strcase.ToUpperCamel(tt.s))
		})
	}
}

func ExampleToUpperCamel() {
	fmt.Println(strcase.ToUpperCamel("lowerCamel"))
	fmt.Println(strcase.ToUpperCamel("UpperCamel"))
	fmt.Println(strcase.ToUpperCamel("lower_snake"))
	fmt.Println(strcase.ToUpperCamel("UPPER_SNAKE"))

	// Output:
	// LowerCamel
	// UpperCamel
	// LowerSnake
	// UpperSnake
}
