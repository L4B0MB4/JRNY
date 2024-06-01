package space

var min = 0x00000000
var max = 0xFFFFFFFF

func equalDivideSpace(workerAmount int) (firstSpace int, remainingSpaces int) {
	maxSpaceSize := max + 1 //+1 due to including min into space size
	leftOver := maxSpaceSize % workerAmount
	spaceSize := (maxSpaceSize - leftOver) / workerAmount

	return spaceSize + leftOver, spaceSize
}
