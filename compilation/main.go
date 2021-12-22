package main

// Void does nothing and returns nothing.
func Void(arg interface{}) {
	// Do nothing
}

// VoidT does nothing and returns nothing.
// The entire purpose of this function is to see how calling it with an
// increasing number of new types for T increases the overall compile time.
func VoidT[T any](arg T) {
	// Do nothing
}

const (
	zeroInt   int   = 0
	zeroInt8  int8  = 0
	zeroInt16 int16 = 0
	zeroInt32 int32 = 0
	zeroInt64 int64 = 0
)

func main() {
	// Invoke the following no-op functions fives times each.
	for i := 0; i < 5; i++ {
		Void(zeroInt)
		Void(zeroInt8)
		Void(zeroInt16)
		Void(zeroInt32)
		Void(zeroInt64)

		VoidT(zeroInt)
		VoidT(zeroInt8)
		VoidT(zeroInt16)
		VoidT(zeroInt32)
		VoidT(zeroInt64)
	}
}
