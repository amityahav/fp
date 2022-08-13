package fp

// Map creates a new array populated with the results of calling a provided callback on every element in the calling array.
//
// Example:
//		s := []int{1, 2, 3, 4}
//		cb := func(elem int) int {
//			return elem * 2
//		}
//
//		ns := Map(cb, s) // output: ns = [2, 4, 6, 8]
//
func Map[T any, Z any](callback func(T) Z, s []T) []Z {
	var newSlice []Z
	for _, elem := range s {
		newSlice = append(newSlice, callback(elem))
	}

	return newSlice
}

// MapI is same as Map except from the callback function which also receives the index of the current element in the slice.
// Example:
//		s := []int{1, 2, 3, 4}
// 		cb := func(elem int, index int) int {
//			return index * elem
//		}
//
//		ns = MapI(cb, s) // output: ns = [0, 2, 6, 12]
func MapI[T any, Z any](callback func(T, int) Z, s []T) []Z {
	var newSlice []Z
	for i, elem := range s {
		newSlice = append(newSlice, callback(elem, i))
	}

	return newSlice
}

// Filter creates a shallow copy of a portion of a given slice,
// filtered down to just the elements from the given slice that pass the test implemented by the provided predicate.
//
// Example:
// 		s := []int{1, 3, 5, 4}
//		cb := func(elem int) bool {
//			return elem%2 == 0
//		}
//
//		ns := Filter(cb, s) // output: ns = [4]
func Filter[T any](pred func(T) bool, s []T) []T {
	var newSlice []T
	for _, elem := range s {
		if pred(elem) {
			newSlice = append(newSlice, elem)
		}
	}

	return newSlice
}

// FilterI is same as Filter except from the predicate function which also receives the index of the current element in the slice.
//
// Example:
// 		s := []int{1, 3, 5, 4}
//		cb := func(elem int, index int) bool {
//			return elem%2 == 0 || index%2 != 0
//		}
//
//		ns := FilterI(cb, s) // output: ns = [3, 4]
func FilterI[T any](pred func(T, int) bool, s []T) []T {
	var newSlice []T
	for i, elem := range s {
		if pred(elem, i) {
			newSlice = append(newSlice, elem)
		}
	}

	return newSlice
}

// Reduce executes a user-supplied "reducer" callback function on each element of the slice, in order,
// passing in the return value from the calculation on the preceding element.
// The final result of running the reducer across all elements of the slice is a single value.
//
// Example:
// 		s := []int{1, 2, 3, 4}
//		cb := func(acc int, curr int, index int) int {
//			return acc + curr
//		}
//
//		ns := Reduce(cb, 0, s) // output: ns = 10
func Reduce[T any](reducer func(T, T) T, acc T, s []T) T {
	for _, elem := range s {
		acc = reducer(acc, elem)
	}

	return acc
}

// ReduceI is same as Reduce except from the reducer function which also receives the index of the current element in the slice.
//
// Example:
// 		s := []int{1, 2, 3, 4}
//		cb := func(acc int, curr int, index int) int {
//			return acc + curr + index
//		}
//
//		ns := ReduceI(cb, 0, s) // output: ns = 16
func ReduceI[T any](reducer func(T, T, int) T, acc T, s []T) T {
	for i, elem := range s {
		acc = reducer(acc, elem, i)
	}

	return acc
}
