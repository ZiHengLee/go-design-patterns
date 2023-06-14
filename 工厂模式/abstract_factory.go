package 工厂模式

import "fmt"

type ShopInterface2 interface {
	GenerateSucaiDumpling() *DumplingInterface2
	GenerateSanxianDumpling() *DumplingInterface2
}

type DumplingInterface2 interface {
	Create2()
}

type HaidianShop2 struct {
}

func (*HaidianShop2) GenerateSucaiDumpling() DumplingInterface2 {
	return new(HaidianSucaiDumpling)
}

func (*HaidianShop2) GenerateSanxianDumpling() DumplingInterface2 {
	return new(HaidianSanxianDumpling)
}

type CaoyangShop2 struct {
}

func (*CaoyangShop2) GenerateSucaiDumpling() DumplingInterface2 {
	return new(CaoyangSucaiDumpling)
}

func (*CaoyangShop2) GenerateSanxianDumpling() DumplingInterface2 {
	return new(CaoyangSanxianDumpling)
}

type HaidianSucaiDumpling struct {
}

func (*HaidianSucaiDumpling) Create2() {
	fmt.Println("我可以生产海淀蔬菜饺子")
}

type HaidianSanxianDumpling struct {
}

func (*HaidianSanxianDumpling) Create2() {
	fmt.Println("我可以生产海淀三鲜饺子")
}

type CaoyangSucaiDumpling struct {
}

func (*CaoyangSucaiDumpling) Create2() {
	fmt.Println("我可以生产朝阳蔬菜饺子")
}

type CaoyangSanxianDumpling struct {
}

func (*CaoyangSanxianDumpling) Create2() {
	fmt.Println("我可以生产朝阳三鲜饺子")
}
