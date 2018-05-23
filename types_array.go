package gotype

import "fmt"

func newTypeArray(v Type, l int) Type {
	return &typeArray{
		val: v,
		le:  l,
	}
}

type typeArray struct {
	typeBase
	le  int
	val Type
}

func (t *typeArray) String() string {
	return fmt.Sprintf("[%v]%v", t.le, t.val)
}

func (t *typeArray) Kind() Kind {
	return Array
}

func (t *typeArray) Elem() Type {
	return t.val
}

func (t *typeArray) Len() int {
	return t.le
}
