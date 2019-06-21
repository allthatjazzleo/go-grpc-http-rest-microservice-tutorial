package main

import (
	"fmt"
	"reflect"
)

func doubleSlice(s interface{}) interface{} {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		fmt.Println("The interface is not a slice.")
		return nil
	}

	v := reflect.ValueOf(s)
	newLen := v.Len()
	newCap := (v.Cap() + 1) * 2
	typ := reflect.TypeOf(s).Elem()

	t := reflect.MakeSlice(reflect.SliceOf(typ), newLen, newCap)
	reflect.Copy(t, v)
	return t.Interface()
}

type b struct {
	c string
}

type a struct {
	something string
	*b
}

func (b *b) p() {
	fmt.Println(b.c)
}
func (b *b) set() {
	b.c = "bye"
}

func main() {
	xs := doubleSlice([]string{"foo", "bar"}).([]string)
	fmt.Println("data =", xs, "len =", len(xs), "cap =", cap(xs))

	ys := doubleSlice([]int{3, 1, 4}).([]int)
	fmt.Println("data =", ys, "len =", len(ys), "cap =", cap(ys))

	testa := a{
		"sdf",
		&b{c: "hi"},
	}
	testa.set()
	testa.p()
}
