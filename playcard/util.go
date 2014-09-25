package main

func panicIfNil(o interface{}) {
	if o == nil {
		panic("is nil")
	}
}
