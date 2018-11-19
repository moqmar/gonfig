package config

import (
	"strconv"
	"time"
)

func (c *Config) AnyMap() map[string]*Config {
	r := map[string]*Config{}
	for _, v := range c.Children {
		r[v.Name] = v
	}
	return r
}
func (c *Config) AnyList() []*Config {
	return c.Children
}

type stringable interface {
	String() string
}

func (c *Config) String() string {
	// FIXME: Type-Switch?!

	s1, ok := c.Value.(string)
	if ok {
		return s1
	}

	s2, ok := c.Value.(stringable)
	if ok {
		return s2.String()
	}

	// FIXME: Nur mit int64, nicht mit int?
	s3, ok := c.Value.(int64)
	if ok {
		return strconv.FormatInt(s3, 10)
	}

	s4, ok := c.Value.(float64)
	if ok {
		return strconv.FormatFloat(s4, 'f', 3, 64)
	}

	s5, ok := c.Value.(bool)
	if ok {
		return strconv.FormatBool(s5)
	}

	return ""
}
func (c *Config) StringMap() map[string]string {
	str := map[string]string{}
	for _, v := range c.Children {
		str[v.Name] = v.String()
	}
	return str
}
func (c *Config) StringList() []string {
	if c.Type == String {
		return []string{c.String()}
	}
	str := []string{}
	for _, v := range c.Children {
		str = append(str, v.String())
	}
	return str
}

func (c *Config) Bool() bool {
	return false
}
func (c *Config) BoolMap() map[string]bool {
	return map[string]bool{}
}
func (c *Config) BoolList() []bool {
	return []bool{}
}

func (c *Config) Int() int {
	v, _ := strconv.Atoi(c.String())
	return v
}
func (c *Config) IntMap() map[string]int {
	return map[string]int{}
}
func (c *Config) IntList() []int {
	return []int{}
}

func (c *Config) Float() float64 {
	return 0
}
func (c *Config) FloatMap() map[string]float64 {
	return map[string]float64{}
}
func (c *Config) FloatList() []float64 {
	return []float64{}
}

func (c *Config) Binary() []byte {
	return []byte{}
}
func (c *Config) BinaryMap() map[string][]byte {
	return map[string][]byte{}
}
func (c *Config) BinaryList() [][]byte {
	return [][]byte{}
}

func (c *Config) Time() time.Time {
	return time.Unix(0, 0)
}
func (c *Config) TimeMap() map[string]time.Time {
	return map[string]time.Time{}
}
func (c *Config) TimeList() []time.Time {
	return []time.Time{}
}
