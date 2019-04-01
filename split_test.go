package strcase_test

import (
	"testing"

	"github.com/hori-ryota/go-strcase"
	"github.com/stretchr/testify/assert"
)

func TestSplitWord(t *testing.T) {
	for _, tt := range []struct {
		s string
		t []string
	}{
		{
			s: "",
			t: []string{},
		},
		{
			s: "a",
			t: []string{
				"a",
			},
		},
		{
			s: "aA",
			t: []string{
				"a",
				"A",
			},
		},
		{
			s: "aaAA",
			t: []string{
				"aa",
				"AA",
			},
		},
		{
			s: "Aa",
			t: []string{
				"Aa",
			},
		},
		{
			s: "AAaa",
			t: []string{
				"A",
				"Aaa",
			},
		},
		{
			s: "a_a",
			t: []string{
				"a",
				"a",
			},
		},
		{
			s: "aa_aa",
			t: []string{
				"aa",
				"aa",
			},
		},
		{
			s: "A_A",
			t: []string{
				"A",
				"A",
			},
		},
		{
			s: "AA_AA",
			t: []string{
				"AA",
				"AA",
			},
		},
		{
			s: "a_A",
			t: []string{
				"a",
				"A",
			},
		},
		{
			s: "aa_AA",
			t: []string{
				"aa",
				"AA",
			},
		},
		{
			s: "A_a",
			t: []string{
				"A",
				"a",
			},
		},
		{
			s: "AA_aa",
			t: []string{
				"AA",
				"aa",
			},
		},
	} {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.t, strcase.SplitWord(tt.s))
		})
	}
}
