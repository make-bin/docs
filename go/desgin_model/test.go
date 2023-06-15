package main

import (
	"fmt"
	"reflect"
)

type Object interface {
	Do()
}

type Dog struct{}

func (d *Dog) Do() {
	fmt.Println("Dog does something")
}

type Cat struct{}

func (c *Cat) Do() {
	fmt.Println("Cat does something")
}

// 定义接口类型
type ObjectIface interface {
	Do()
}

// 实现接口类型
type DogIface struct{}

func (d *DogIface) Do() {
	fmt.Println("Dog does something")
}

type CatIface struct{}

func (c *CatIface) Do() {
	fmt.Println("Cat does something")
}

// 定义工厂函数
func makeObject(iface ObjectIface) Object {
	obj := reflect.New(iface).Interface().(Object)
	obj.Do()
	return obj
}

func main() {

	// 创建对象
	dogIface := &DogIface{}
	catIface := &CatIface{}
	obj := makeObject(dogIface)
	obj.Do()
	obj = makeObject(catIface)
	obj.Do()

	// 输出结果
	fmt.Println("Dog does something")
	fmt.Println("Cat does something")
}
