package gonfig

import (
	"strconv"
	"strings"
	"time"
)

//TODO:
func (c *Config) AnyMap() map[string]*Config {
	r := map[string]*Config{}
	for _, v := range c.Children {
		r[v.Name] = v
	}
	return r
}

//TODO:
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

//TODO: muss hier überhaupt noch was gemacht werden?
func (c *Config) StringMap() map[string]string {
	str := map[string]string{}
	for _, v := range c.Children {
		str[v.Name] = v.String()
	}
	return str
}

//TODO: muss hier noch was gemacht werden?
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
	switch val := c.Value.(type) {
	case string:
		if strings.ToLower(val) == "true" {
			return true
		}
	case int:
		if val == 1 {
			return true
		}
	case int16:
		if val == 1 {
			return true
		}
	case int32:
		if val == 1 {
			return true
		}
	case int8:
		if val == 1 {
			return true
		}
	case int64:
		if val == 1 {
			return true
		}
	case float32:
		if val == 1 {
			return true
		}
	case float64:
		if val == 1 {
			return true
		}
	case bool:
		return val
	}
	return false
}

//TODO:
func (c *Config) BoolMap() map[string]bool {
	return map[string]bool{}
}

//TODO:
func (c *Config) BoolList() []bool {
	return []bool{}
}

func (c *Config) Int() int {
	switch val := c.Value.(type) {
	case string:
		res, err := strconv.Atoi(val)
		if err != nil {
			return -999999
		}
		return res
	case int:
		return val
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int8:
		return int(val)
	case int64:
		return int(val)
	case float32:
		return int(val)
	case float64:
		return int(val)
	case bool:
		if val {
			return 1
		}
		return 0
	}
	return -999999
}

//TODO:
func (c *Config) IntMap() map[string]int {
	return map[string]int{}
}

//TODO:
func (c *Config) IntList() []int {
	return []int{}
}

func (c *Config) Float() float64 {
	switch val := c.Value.(type) {
	case string:
		res, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return -9.99999
		}
		return res
	case int:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int8:
		return float64(val)
	case int64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	case bool:
		if val {
			return 1.0
		}
		return 0.0
	}
	return -9.9999
}

//TODO:
func (c *Config) FloatMap() map[string]float64 {
	return map[string]float64{}
}

//TODO:
func (c *Config) FloatList() []float64 {
	return []float64{}
}

func (c *Config) Binary() []byte {
	switch val := c.Value.(type) {
	case string:
		return []byte(val)
	return []byte(c.String())
}

//TODO:
func (c *Config) BinaryMap() map[string][]byte {
	return map[string][]byte{}
}

//TODO:
func (c *Config) BinaryList() [][]byte {
	return [][]byte{}
}

//TODO:
func (c *Config) Time() time.Time {
	return time.Unix(0, 0)
}

//TODO:
func (c *Config) TimeMap() map[string]time.Time {
	return map[string]time.Time{}
}

//TODO:
func (c *Config) TimeList() []time.Time {
	return []time.Time{}
}
