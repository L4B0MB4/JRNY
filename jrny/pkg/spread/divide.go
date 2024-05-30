package spread

import "fmt"

var min = 0x00000000
var max = 0xFFFFFFFF

func EqualDivideSpace(workerAmount int) {
	extra := max % workerAmount
	spaceSize := (max - extra) / workerAmount
	fmt.Printf("%x / %x = %x and would have %v extra \n", max, workerAmount, spaceSize, extra)
}
