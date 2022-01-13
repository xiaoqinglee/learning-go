package gotcha

import "fmt"

/*
Breaking Out of "for switch" and "for select" Code Blocks

A "break" statement without a label only gets you out of the inner switch/select block.
If using a "return" statement is not an option then defining a label for the outer loop
is the next best thing.

A "goto" statement will do the trick too...
*/

func RightBreakOut() {
loopFoo:
	for {
		switch {
		default:
			fmt.Println("breaking out...")
			break loopFoo
		}
	}
	fmt.Println("out of loopFoo!")

loopBar:
	for {
		select {
		default:
			fmt.Println("breaking out...")
			break loopBar
		}
	}
	fmt.Println("out of loopBar!")
}
