package gonfig

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
	switch val := c.Value.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', 3, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', 3, 64)
	case bool:
		return strconv.FormatBool(val)
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
