package 单例模式

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	g1 := GetSingleton()
	g2 := GetSingleton()
	fmt.Println(g1 == g2)
}
