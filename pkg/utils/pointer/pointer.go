package pointer

import "reflect"

func GetPointerOrDefaultValue[I any](pointer *I) (value I) {
	if pointer != nil {
		return *pointer
	}
	return value
}

func GetNilIfDefaultValue[I any](value I) *I {
	var zeroValue I
	if reflect.DeepEqual(value, zeroValue) {
		return nil
	}
	return &value
}

func Ptr[T any](in T) *T {
	return &in
}
