package utils

import "path"

type ComposableMap[T any] struct {
	defaultVal T
	mapping    map[string]T
}

func (c *ComposableMap[T]) GetKeys() (out []string) {
	for k := range c.mapping {
		out = append(out, k)
	}
	return
}

func (c *ComposableMap[T]) GetDefault() T {
	return c.defaultVal
}

func (c *ComposableMap[T]) GetMapping() map[string]T {
	return c.mapping
}

func (c *ComposableMap[T]) Set(key string, val T) {
	c.mapping[key] = val
}

func (c *ComposableMap[T]) Get(key string) (out T) {
	if val, ok := c.mapping[key]; ok {
		return val
	}

	return c.GetDefault()
}

func (c *ComposableMap[T]) Add(prefix string, in ComposableMap[T]) {
	c.Set(prefix, in.GetDefault())

	for k, v := range in.GetMapping() {
		c.Set(path.Join(prefix, k), v)
	}
}
