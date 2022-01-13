package gotcha

import "fmt"

/*
The First Use of iota Doesn't Always Start with Zero

It may seem like the iota identifier is like an increment operator.
You start a new constant declaration and the first time you use iota you get zero,
the second time you use it you get one and so on. It's not always the case though.

The iota is really an index operator for the current line in the constant declaration block,
so if the first use of iota is not the first line in the constant declaration block
the initial value will not be zero.
*/

const (
	azero = iota
	aone  = iota
)

const (
	foo   = "foo"
	bzero = iota
	bone  = iota
)

const (
	czero = 0
	cone  = 1
	ctwo  = iota
	cthree
	cfour
	cfive = 5
	csix
	cseven
)

//0 1
//1 2
//0 1 2 3 4 5 5 5

func UseIota() {
	fmt.Println(azero, aone)
	fmt.Println(bzero, bone)
	fmt.Println(czero, cone, ctwo, cthree, cfour, cfive, csix, cseven)
}
