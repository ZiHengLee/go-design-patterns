package 原型模式

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtoType(t *testing.T) {
	words := Keywords{
		"testA": &Keyword{
			word:  "testA",
			visit: 1,
		},
		"testB": &Keyword{
			word:  "testB",
			visit: 2,
		},
	}

	updatedWords := []*Keyword{
		{
			word:  "testB",
			visit: 10,
		},
	}

	got := words.Clone(updatedWords)

	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updatedWords[0], got["testB"])
}
