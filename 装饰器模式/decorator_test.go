package 装饰器模式

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecorator(t *testing.T) {
	sq := Square{}
	//给矩形加个颜色
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)
}
