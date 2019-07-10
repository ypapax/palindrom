package polindrom_test

import (
	"testing"

	"github.com/ypapax/polindrom"

	"github.com/stretchr/testify/assert"
)

func TestPoly(t *testing.T) {
	type testCase struct {
		inp string
		exp bool
	}
	for _, c := range []testCase{
		{inp: "Казак", exp: true},
		{inp: "А роза упала на лапу Азора", exp: true},
		{inp: "Do geese see God?", exp: true},
		{inp: "Madam, I’m Adam", exp: true},
		{inp: "done", exp: false},
	} {
		t.Run(c.inp, func(t *testing.T) {
			as := assert.New(t)
			act := polindrom.IsPolindrom([]rune(c.inp))
			if !as.Equal(c.exp, act) {
				return
			}
		})
	}
}
