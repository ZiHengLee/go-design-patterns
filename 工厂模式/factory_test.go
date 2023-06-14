package 工厂模式

import (
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	haidianShop := HaidianShop{}
	haidianShop.CreateDumpling("sucai").Create()
}

func TestFactoryMethod(t *testing.T) {
	beijing := Beijing{}
	beijing.CreateShop("haidian").CreateDumpling1("sucai").Create()
	beijing.CreateShop("caoyang").CreateDumpling1("sucai").Create()
	beijing.CreateShop("caoyang").CreateDumpling1("sanxian").Create()
}

func TestAbstractFactoryMethod(t *testing.T) {
	haidian := HaidianShop2{}
	caoyang := CaoyangShop2{}
	haidian.GenerateSucaiDumpling().Create2()
	haidian.GenerateSanxianDumpling().Create2()
	caoyang.GenerateSucaiDumpling().Create2()
	caoyang.GenerateSanxianDumpling().Create2()
}
