package utilsgo

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func EntriesMap[K comparable, V any](m map[K]V) []Entry[K, V] {
	result := make([]Entry[K, V], 0, len(m))
	for key, value := range m {
		result = append(result, Entry[K, V]{
			Key:   key,
			Value: value,
		})
	}
	return result
}
