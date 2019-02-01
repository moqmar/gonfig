package gonfig

import (
	"strings"
)

type Type string

const (
	AnyMap     Type = "any-map"
	AnyList    Type = "any-list"
	String     Type = "string"
	StringMap  Type = "string-map"
	StringList Type = "string-list"
	Bool       Type = "bool"
	BoolMap    Type = "bool-map"
	BoolList   Type = "bool-list"
	Int        Type = "int"
	IntMap     Type = "int-map"
	IntList    Type = "int-list"
	Float      Type = "float"
	FloatMap   Type = "float-map"
	FloatList  Type = "float-list"
	Binary     Type = "binary"
	BinaryMap  Type = "binary-map"
	BinaryList Type = "binary-list"
	Time       Type = "time"
	TimeMap    Type = "time-map"
	TimeList   Type = "time-list"
	Null       Type = "null"
	Comment    Type = "comment"
	LineBreak  Type = "line-break"
	Dummy      Type = "dummy"
)

type Config struct {
	Parent            *Config
	Index             int
	Path              string
	Name              string
	Value             interface{}
	Exists            bool
	Type              Type
	Children          []*Config
	hasDefault        bool
	changeHandler     func(*Config, string)
	deepChangeHandler func(*Config, string)
}

func leaf(c *Config) *Config {
	return &Config{
		Parent: c,
		Type:   Dummy,
		Exists: false,
	}
}

//////////////
// Querying //
//////////////

func (c *Config) Get(path ...string) *Config {
	// TODO: why are there multiple paths?!
	s := strings.SplitN(strings.Join(path, "."), ".", 2)
	if len(s[0]) > 0 {
		for _, x := range c.Children {
			if x.Name == s[0] {
				if len(s) > 1 {
					return x.Get(s[1])
				}
				return x
			}
		}
		return leaf(c)
	}
	return c
}

//Return Child of c
func (c *Config) Child(name string) *Config {
	for _, c := range c.Children {
		if c.Name == name && c.Exists {
			return c
		}
	}
	return leaf(c)
}

/////////////
// Editing //
/////////////

//TODO: später
// Set the value to `new` and update all children. Will remove comments and line breaks, and might reorder maps.
func (c *Config) Update(new interface{}) error {
	return nil
}

//TODO: später
// Change so the new Value is `new`, but try to keep comments, map order and line breaks.
func (c *Config) Apply(new interface{}) error {
	return nil
}

//TODO: später
// Replace the element by another one. Must not exist in the configuration tree yet!
func (c *Config) Replace(new *Config) error {
	return nil
}

//TODO: später
func (c *Config) Delete() error {
	return nil
}

//TODO: später
func (c *Config) OnChange(handler func(*Config, string), deep bool) {
	// handler(c, reason)
}

/////////////
// Casting //
/////////////

//TODO: Und was tut es?!?
func (c *Config) Default(value interface{}) *Config {
	c2 := &*c
	if !c2.Exists {
		apply(value, c2)
	}
	return c2
}

//TODO: Return True if native type ist t
func (c *Config) Is(t Type) bool {
	return true
}
