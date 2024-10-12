package maps

type Map map[string]interface{}

func (m Map) Get(key string) interface{} {
	return m[key]
}

func (m Map) Keys() (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

func (m Map) Values() (values []interface{}) {
	for _, value := range m {
		values = append(values, value)
	}
	return
}

func (m Map) Merge(other Map) {
	for key, value := range other {
		if _, ok := m[key]; ok {
			m[key] = value
		}
	}
}
