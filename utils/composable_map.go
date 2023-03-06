package utils

type ComposableMap[T any] struct {
	defaultVal    T
	directMapping map[string]T
	nestedMapping map[string]ComposableMap[T]
}

func (c *ComposableMap[T]) Get(key string) (out T, ok bool) {
	ok = false

	if key == "" {
		return c.defaultVal, true
	}

	if val, ok := c.directMapping[key]; ok {
		return val, ok
	}

	for _, val := range c.nestedMapping {
		if v, ok := val.Get(key); ok {
			return v, ok
		}
	}

	return
}

func (c *ComposableMap[T]) SetDefault(val T) {
	c.defaultVal = val
}

func (c *ComposableMap[T]) SetDirect(key string, val T) {
	c.directMapping[key] = val
}

func (c *ComposableMap[T]) SetNested(key string, val ComposableMap[T]) {
	c.nestedMapping[key] = val
}

func (c *ComposableMap[T]) ReplaceDirect(in map[string]T) {
	c.directMapping = in
}

func (c *ComposableMap[T]) ReplaceNested(in map[string]ComposableMap[T]) {
	c.nestedMapping = in
}
