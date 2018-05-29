package gotype

func newTypeVar(name string, typ Type) *typeVar {
	v, ok := typ.(*typeVar)
	if ok {
		return &typeVar{
			name: name,
			elem: v.elem,
		}
	}
	return &typeVar{
		name: name,
		elem: typ,
	}
}

type typeVar struct {
	typeBase
	elem Type
	name string
}

func (t *typeVar) String() string {
	return t.name
}

func (t *typeVar) Name() string {
	return t.name
}

func (t *typeVar) Kind() Kind {
	return Var
}

func (t *typeVar) Elem() Type {
	return t.elem
}
