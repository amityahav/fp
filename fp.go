package fp

// Map function takes as arguments a callback function and a slice of type T and returns a new slice of type Z
func Map[T any, Z any](callback func(T, int) Z, arr []T) []Z {
	var newSlice []Z
	for i, elem := range arr {
		newSlice = append(newSlice, callback(elem, i))
	}

	return newSlice
}
