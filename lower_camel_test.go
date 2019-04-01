package strcase_test

import (
	"fmt"
	"testing"

	"github.com/hori-ryota/go-strcase"
	"github.com/stretchr/testify/assert"
)

func TestToLowerCamel(t *testing.T) {
	for _, tt := range testPatterns {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.lc, strcase.ToLowerCamel(tt.s))
		})
	}
}

func ExampleToLowerCamel() {
	fmt.Println(strcase.ToLowerCamel("lowerCamel"))
	fmt.Println(strcase.ToLowerCamel("UpperCamel"))
	fmt.Println(strcase.ToLowerCamel("lower_snake"))
	fmt.Println(strcase.ToLowerCamel("UPPER_SNAKE"))

	// Output:
	// lowerCamel
	// upperCamel
	// lowerSnake
	// upperSnake
}
