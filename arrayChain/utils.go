package arrayChain

func panicIfNil(this interface{}) {
	if this == nil {
		panic("the object is nil")
	}
}
