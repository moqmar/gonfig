package main

import "github.com/moqmar/gonfig"

/*

Supported formats: YAML, JSON, Env() (SOME_THING_GREAT=test), Arg() (--some-thing-great test), Default(), @json:<json>, @yaml:<yaml>
Paths may only contain letters, numbers and dots and are case-insensitive
For more complex keys, use the case-sensitive Child method - `c.Get("some.thing.great")` is the same as `c.Child("some").Child("thing").Child("great")` (except for case sensitivity)
c.Children always contains all children, including line breaks and comments!
Accessing a comment: if e.children[2].Type == config.Comment { return e.Children[2].Value.(string) }

*/

func main() bool {

	c := gonfig.Open(gonfig.Arg(), gonfig.Env(), "config.yaml", "/etc/myprogram.yaml", gonfig.Default(map[string]string{"hello": "world"}))
	c.Get("some.thing.great").Default("hello world").String()

	// When not checking with `Is(...)`, will return default values if non-castable

	// anything:
	// - example: test
	//   whatever: 2
	// - example: hello world
	//   whatever: 9001
	for _, d := range c.Get("anything").AnyList() {
		for k, e := range d.AnyMap() {
			if k == "whatever" && e.Int() > 9000 {
				println(e.Parent.Get("example").String())
			}
		}
	}

	// and: [true, true, true]
	x := c.Get("and").Default([]bool{false})
	if x.Is(gonfig.BoolList) {
		for b := range x.BoolList() {
			if x == false {
				return false
			}
		}
		return true
	} else {
		return false
	}

}
