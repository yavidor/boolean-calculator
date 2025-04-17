package main

func and(A, B bool) bool {
	return A && B
}

func or(A, B bool) bool {
	return A || B
}

func xor(A, B bool) bool {
	return A != B
}

func nand(A, B bool) bool {
	return !(A && B)
}

func nor(A, B bool) bool {
	return !(A || B)
}

func xnor(A, B bool) bool {
	return A == B
}
