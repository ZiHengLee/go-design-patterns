package 工厂模式

import "fmt"

type ShopInterface2 interface {
	GenerateSucaiDumpling() *DumplingInterface2
	GenerateSanxianDumpling() *DumplingInterface2
}

type DumplingInterface2 interface {
	Create()
}

type HaidianShop2 struct {
}

func (*HaidianShop2) GenerateSucaiDumpling() DumplingInterface1 {
	return new(HaidianSucaiDumpling)
}

func (*HaidianShop2) GenerateSanxianDumpling() DumplingInterface1 {
	return new(HaidianSanxianDumpling)
}

type CaoyangShop2 struct {
}

func (*CaoyangShop2) GenerateSucaiDumpling() DumplingInterface1 {
	return new(CaoyangSucaiDumpling)
}

func (*CaoyangShop2) GenerateSanxianDumpling() DumplingInterface1 {
	return new(CaoyangSanxianDumpling)
}

type HaidianSucaiDumpling struct {
}

func (*HaidianSucaiDumpling) Create() {
	fmt.Println("我可以生产海淀蔬菜饺子")
}

type HaidianSanxianDumpling struct {
}

func (*HaidianSanxianDumpling) Create() {
	fmt.Println("我可以生产海淀三鲜饺子")
}

type CaoyangSucaiDumpling struct {
}

func (*CaoyangSucaiDumpling) Create() {
	fmt.Println("我可以生产朝阳蔬菜饺子")
}

type CaoyangSanxianDumpling struct {
}

func (*CaoyangSanxianDumpling) Create() {
	fmt.Println("我可以生产朝阳三鲜饺子")
}
