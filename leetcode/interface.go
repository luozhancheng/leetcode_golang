package leetcode

import "fmt"

type A interface {
	a() bool
}


type B struct {
	b int
}

func (b *B) a() bool {
	fmt.Println("b impl = ", b.b)
	return true
}

//region Description
//<editor-fold desc="Description">
/*func (b B) a() bool {
	fmt.Println("bb impl = ", b.b)
	return false
}*/
//</editor-fold>
//endregion

func test(a A)  {
	a.(*B).a()
}

func Test_interface() {
	b := B{10}
	b.a()
	test(&b)
}
