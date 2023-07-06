package 模版模式

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplate(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", 8888)
	assert.NoError(t, err)
}
