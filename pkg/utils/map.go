package utils

func (m Map) Get(key string) any {
	val, ok := m[key]
	if ok {
		return val
	}
	return nil
}
