package gonfig

import (
	"fmt"
	"strconv"
)

func apply(data interface{}, cfg *Config) {
	switch x := data.(type) {
	case map[interface{}]interface{}:
		fmt.Printf("Got map: %s\n", cfg.Path)
		cfg.Exists = true
		cfg.Type = AnyMap // TODO: should be more specific!
		cfg.Children = []*Config{}
		i := 0
		for k, y := range x {
			i++
			ks := fmt.Sprint(k)
			child := &Config{
				Exists: true,
				Index:  i,
				Name:   ks,
				Parent: cfg,
				Path:   cfg.Path + "." + ks,
			}
			apply(y, child)
			cfg.Children = append(cfg.Children, child)
		}
	case []interface{}:
		fmt.Printf("Got list: %s\n", cfg.Path)
		cfg.Exists = true
		cfg.Type = AnyList // TODO: should be more specific!
		cfg.Children = []*Config{}
		for i, y := range x {
			child := &Config{
				Exists: true,
				Index:  i,
				Name:   strconv.Itoa(i),
				Parent: cfg,
				Path:   cfg.Path + "." + strconv.Itoa(i),
			}
			apply(y, child)
			cfg.Children = append(cfg.Children, child)
		}
	default:
		fmt.Printf("Got any: %s\n", cfg.Path)
		cfg.Exists = true
		cfg.Value = data
		cfg.Type = String // TODO: should be more specific!
	}
}

func merge(cfg ...*Config) {

}
