package ww

func Value[T any](pointer *T, defaultValue T) T {
	if pointer == nil {
		return defaultValue
	}
	return *pointer
}

func ValueOrNil[T any](pointer *T) any {
	if pointer == nil {
		return pointer
	}
	return *pointer
}

func Pointer[T any](value T) *T {
	return &value
}
