package main

import (
	"fmt"

	"reflect"
)

// interface
type Product interface {
	Name() string
	Description() string
}

// factory
type factory struct{}

func NewFactory() *factory {
	return &factory{}
}

func (f *factory) CreateProduct(name string) (Product, error) {
	var product Product
	fmt.Println(reflect.ValueOf(product).Elem())

	return product, nil
}

// impl
type dog struct{}

func (a *dog) Name() string {
	return "dog"
}

func (a *dog) Description() string {
	return "A friendly dog"
}

type cat struct{}

func (c *cat) Name() string {
	return "cat"
}

func (c *cat) Description() string {
	return "A fluffy cat"
}

func main() {
	f := NewFactory()
	animalProduct, err := f.CreateProduct("dog")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(animalProduct.Name())

	catProduct, err := f.CreateProduct("cat")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(catProduct.Name())
}
