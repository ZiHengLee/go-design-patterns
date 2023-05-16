package 工厂模式

import "fmt"

type FactoryInterface1 interface {
	CreateShop(t string) ShopInterface1
}

type ShopInterface1 interface {
	CreateDumpling1(t string) DumplingInterface1
}

type Beijing struct {
}

func (*Beijing) CreateShop(t string) ShopInterface1 {
	switch t {
	case "haidian":
		return new(HaidianShop1)
	case "caoyang":
		return new(CaoyangShop1)
	}
	return nil
}

type HaidianShop1 struct {
}

func (*HaidianShop1) CreateDumpling1(t string) DumplingInterface1 {
	switch t {
	case "sucai":
		return new(SucaiDumpling1)
	}
	return nil
}

type CaoyangShop1 struct {
}

func (*CaoyangShop1) CreateDumpling1(t string) DumplingInterface1 {
	switch t {
	case "sucai":
		return new(SucaiDumpling1Pro)
	case "sanxian":
		return new(SanxianDumpling1)
	}
	return nil
}

type DumplingInterface1 interface {
	Create()
}

type SucaiDumpling1 struct {
}

func (*SucaiDumpling1) Create() {
	fmt.Println("我可以生产蔬菜饺子1")
}

type SucaiDumpling1Pro struct {
}

func (*SucaiDumpling1Pro) Create() {
	fmt.Println("我可以生产更高级的蔬菜饺子1")
}

type SanxianDumpling1 struct {
}

func (*SanxianDumpling1) Create() {
	fmt.Println("我可以生产三鲜饺子1")
}
