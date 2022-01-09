package counting

type Counter struct {
	Count int
}

func (c *Counter) Next() int {
	result := c.Count
	c.Count++
	return result
}
