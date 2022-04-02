package defs

type UnaryFunc[I any, O any] func(I) O

type UnaryPredicate[I any] UnaryFunc[I, bool]

type UnaryProc[I any] func(I)

type BinaryFunc[I1, I2, O any] func(I1, I2) O

type BinaryPredicate[I1, I2 any] BinaryFunc[I1, I2, bool]

func ZEqual[T comparable](x, y T) bool {
	return x == y
}
