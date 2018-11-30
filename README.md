# gonfig
The easiest way to use YAML configuration files for your Go programs. Inspired by the Bukkit configuration API.

***THIS IS NOT YET DONE AND MISSES A LOT OF THE FUNCTIONALITY LISTED HERE! DON'T USE THIS YET, UNLESS YOU WANT TO CONTRIBUTE TO GONFIG ITSELF IF YOU EXPERIENCE A PROBLEM!***

The main principle: it *always* returns exactly what you tell it - no errors possible. If something's required, you have to check if it exists manually, but the preferred way are default values to handle unexpected behaviour.

**How the hell do I debug this thing?**  
That's a valid question - set GONFIG_DEBUG=1, and gonfig will print messages like the following to `log`:
```
[example.hello] Want STRING, got LIST, using first element (IMPORTANT)
[example.hello]  -> got BOOL, casting to STRING (IMPORTANT)
[example.hello]  -> false
[sum] Want INT LIST, got LIST, using all elements
[sum.0] Want INT, got INT -> 5
[sum.1] Want INT, got INT -> 12
[sum.2] Not a number -> skipping (IMPORTANT)
[sum.3] Want INT, got INT -> 39
[sum] -> [5 12 39]
```

Usage example:
```golang
package main

import "github.com/moqmar/gonfig"

var cfg = config.Open(config.Arg(), config.Env(), "config.yaml", "/etc/myprogram.yaml", `---
# Default configuration file
# - Provides defaults for any options
# - If none of the files listed in config.Open() exist, will be copied to the
#   last writable path listed there (probably "config.yaml" in this example,
#   "/etc/myprogram.yaml" if run as root)
example:
  hello: world
sum:
- 5
- 12
- hello
- 39
`)

func main() {
    fmt.Printf("Hello %s\n", cfg.Get("example.hello").String()) // "Hello world"
    n := 0
    for _, x := range cfg.Get("sum").IntList() {
        n += x
    }
    fmt.Printf("The sum of all numbers in sum is %d\n", n) // 5+12+39 = 56 (things not convertible to int are ignored)
}
```

# What happens if the file contains...

- a **list**, and you want...
  - a **list**: you get a list containing all elements
  - a **map**: you get a map, the list index casted to a string is used as the key
  - a **value**: you get the first element

- a **map**, and you want...
  - a **list**: you get a list containing all elements, in no particular order
  - a **map**: you get a map containing all elements with the keys used in the YAML file
  - a **value**: the element is handled as if it didn't exist (using the in-place-default or null value)

- a **value**, and you want...
  - a **list**: you get a list containing exactly one element
  - a **map**: you get an empty map
  - a **value**: you get the value casted to your target type using the rules below

- a **string**, and you want...
  - a **string**: you get the `string` value
  - a **bool**: 
  - an **int**: 
  - a **float**:
  - a **binary**: you get the string as a `[]byte` slice
  - a **time**: the string is casted to a `time.Time` value, if that fails, the in-place-default or null value is used

- a **bool**, and you want...
  - a **string**: 
  - a **bool**: you get the `bool` value
  - an **int**: 
  - a **float**:
  - a **binary**: you get the string as a `[]byte` slice
  - a **time**: the in-place-default or null value is used

- an **int**, and you want...
  - a **string**: 
  - a **bool**: 
  - an **int**: you get the `int` value
  - a **float**:
  - a **binary**: 
  - a **time**: you get a `time.Time` value, the original value is used as a unix timestamp

- a **float**, and you want...
  - a **string**: 
  - a **bool**: 
  - an **int**: 
  - a **float**: you get the `float` value
  - a **binary**: you get the string as a `[]byte` slice
  - a **time**: you get a `time.Time` value, the original value is used as a unix timestamp including nanoseconds

- a **binary**, and you want...
  - a **string**: 
  - a **bool**: 
  - an **int**: 
  - a **float**: 
  - a **binary**: you get the `binary` value
  - a **time**: the value is handled like a string (see above)

- a **time**, and you want...
  - a **string**: 
  - a **bool**: 
  - an **int**: 
  - a **float**: 
  - a **binary**: 
  - a **time**: you get the `time.Time` value
