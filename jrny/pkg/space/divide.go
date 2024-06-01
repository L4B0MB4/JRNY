package space

import (
	"fmt"
	b "math/big"
)

var MaxSpaceSize *b.Int

func Init() {
	if MaxSpaceSize != nil {
		return
	}
	onebyte := b.NewInt(0x10)
	exponent := b.NewInt(32)
	MaxSpaceSize = onebyte.Exp(onebyte, exponent, nil)
	fmt.Printf("Max space is %x in hex\n", MaxSpaceSize)
}

func equalDivideSpace(workerAmount int64) (first *b.Int, space *b.Int) {
	Init()
	workerAmountBig := b.NewInt(workerAmount)

	leftOver := b.NewInt(0)
	leftOver.Mod(MaxSpaceSize, workerAmountBig)
	spaceSizeDividable := b.NewInt(0)
	spaceSizeDividable.Sub(MaxSpaceSize, leftOver)
	spaceSize := b.NewInt(0)
	spaceSize.Div(spaceSizeDividable, workerAmountBig)

	fmt.Printf("Equal space per worker is %x in hex\n", spaceSize)
	firstSpace := b.NewInt(0)
	firstSpace.Add(spaceSize, leftOver)

	return firstSpace, spaceSize

}
