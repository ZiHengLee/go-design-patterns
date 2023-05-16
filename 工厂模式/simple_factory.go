package 工厂模式

import "fmt"

type FactoryInterface interface {
	CreateDumpling(t string) DumplingInterface
}

type HaidianShop struct {
}

func (*HaidianShop) CreateDumpling(t string) DumplingInterface {
	switch t {
	case "sucai":
		return new(SucaiDumpling)
	}
	return nil
}

type DumplingInterface interface {
	Create()
}

type SucaiDumpling struct {
}

func (*SucaiDumpling) Create() {
	fmt.Println("我可以生产蔬菜饺子")
}
