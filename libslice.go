/*
	libslice
	author: fbac <me@fbac.dev>

	example of slice library to add
	some missing funcionality in stdlib
*/
package libslice

// sliceType is generic type
// so the functions can be called
// using a slice of any type
type sliceType interface {
	int | float32 | float64 | string
}

// Delete the item in the first position (i = 0)
func PopFirst[sT sliceType](slice []sT) []sT {
	ret := slice[1:]
	return ret
}

// Delete the item in the last position (i = len(slice))
func PopLast[sT sliceType](slice []sT) []sT {
	ret := slice[:len(slice)-1]
	return ret
}

// PopN deletes the item in the position n
// WARNING: as this function uses append
// the original slice will be modified too
// since an slice is a struct in the form of
// a pointer to the underlying array, len and cap
// and we modify it in memory
// example:
// [a] original slice:     	  [0 1 2 3 4]
// [b] result of PopN(a, 1):      [0 2 3 4]
// [a] original slice after PopN: [0 2 3 4 4]
// a is modified, it preserves len(5)
// the last item is duplicated
func PopN[sT sliceType](slice []sT, n int) []sT {
	ret := append(slice[:n], slice[n+1:]...)
	return ret
}

// Return a slice containing the first n items
func GetFirstN[sT sliceType](slice []sT, n int) []sT {
	ret := slice[:n]
	return ret
}

// Return a slice containing the last n items
func GetLastN[sT sliceType](slice []sT, n int) []sT {
	ret := slice[len(slice)-n:]
	return ret
}
