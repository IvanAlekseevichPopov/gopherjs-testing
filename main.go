package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	//fmt.Println(js.Global.Get("document"))
	js.Global.Call("alert", "Test")

	js.Global.Get("document").Call("write", "Hello world!")
	//js.Global.Get("document").Call("", "Hello world!")
}
//
//type Pet struct {
//	name string
//}
//
//func New(name string) *js.Object {
//	return js.MakeWrapper(&Pet{name})
//}
//
//func (p *Pet) Name() string {
//	return p.name
//}
//
//func (p *Pet) SetName(name string) {
//	p.name = name
//}