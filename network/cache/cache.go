package cache

import "fmt"

type Cache struct {
	Objects []ObjectCache
}

type ObjectCache struct {
	Identifier string // is proxy-framework or server-framework
	Content    []interface{}
}

func (c *Cache) Clear() {
	c.Objects = make([]ObjectCache, 0)
}

func NewCache() *Cache {
	return &Cache{
		Objects: []ObjectCache{},
	}
}

func (c *Cache) GetAll() []ObjectCache {
	return c.Objects
}

func (c *Cache) Get(identifier string) []interface{} {
	for _, object := range c.Objects {
		if object.Identifier == identifier {
			return object.Content
		}
	}
	return nil
}

func (c *Cache) Set(identifier string, content []interface{}) {
	for i, object := range c.Objects {
		if object.Identifier == identifier {
			c.Objects[i].Content = content
			return
		}
	}
	c.Objects = append(c.Objects, ObjectCache{
		Identifier: identifier,
		Content:    content,
	})
}

func (c *Cache) Remove(identifier string, content interface{}) {
	for i, object := range c.Objects {
		if object.Identifier == identifier {
			for j, element := range object.Content {
				if element == content {
					c.Objects[i].Content = append(c.Objects[i].Content[:j], c.Objects[i].Content[j+1:]...)
					return
				}
			}
		}
	}
}

func (c *Cache) Delete(identifier string) {
	for i, object := range c.Objects {
		if object.Identifier == identifier {
			c.Objects = append(c.Objects[:i], c.Objects[i+1:]...)
			return
		}
	}
}

func (c *Cache) Exists(identifier string) bool {
	for _, object := range c.Objects {
		if object.Identifier == identifier {
			return true
		}
	}
	return false
}

func (c *Cache) GetOrSet(identifier string, content []interface{}) []interface{} {
	if c.Exists(identifier) {
		return c.Get(identifier)
	}
	c.Set(identifier, content)
	return content
}

func (c *Cache) AddMultipleElements(identifier string, content []interface{}) {
	if !c.Exists(identifier) {
		c.Set(identifier, content)
		return
	}
	for i, object := range c.Objects {
		if object.Identifier == identifier {
			c.Objects[i].Content = append(c.Objects[i].Content, content...)
			return
		}
	}
}

func (c *Cache) Add(identifier string, content interface{}) error {
	slice, err := convertToSlice(content)
	if err != nil {
		return err
	}
	c.AddMultipleElements(identifier, slice)
	return nil
}

func convertToSlice(content interface{}) ([]interface{}, error) {
	// Check if the content is already a slice
	if slice, ok := content.([]interface{}); ok {
		return slice, nil
	}
	// Check if the content is a single element
	if element, ok := content.(interface{}); ok {
		return []interface{}{element}, nil
	}
	// If the content is not a slice or single element, return an error
	return nil, fmt.Errorf("content is not a slice or single element")
}
