package polindrom_test

import (
	"testing"

	"github.com/ypapax/polindrom"

	"github.com/stretchr/testify/assert"
)

func TestPoly1(t *testing.T) {
	as := assert.New(t)
	act := polindrom.IsPolindrom([]rune("Казак"))
	if !as.Equal(true, act) {
		return
	}
}

func TestPoly2(t *testing.T) {
	as := assert.New(t)
	act := polindrom.IsPolindrom([]rune("А роза упала на лапу Азора"))
	if !as.Equal(true, act) {
		return
	}
}

func TestPoly3(t *testing.T) {
	as := assert.New(t)
	act := polindrom.IsPolindrom([]rune("Do geese see God?"))
	if !as.Equal(true, act) {
		return
	}
}

func TestPoly4(t *testing.T) {
	as := assert.New(t)
	act := polindrom.IsPolindrom([]rune("Madam, I’m Adam"))
	if !as.Equal(true, act) {
		return
	}
}
