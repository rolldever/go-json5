package json5

import (
	"encoding/json"

	"github.com/robertkrimen/otto"
)

// Unmarshal parses the JSON5-encoded data `src`
// and stores the result in the value pointed to by v.
func Unmarshal(src []byte, v interface{}) error {
	vm := otto.New()
	obj, err := vm.Object("x=" + string(src))
	if err != nil {
		return err
	}

	// m here is something a `map[string]interface{}` object
	m, err := obj.Value().Export()
	if err != nil {
		return err
	}

	switch t := v.(type) {
	case *map[string]interface{}:
		*t = m.(map[string]interface{})
		return nil
	}

	// jb - JSON Bytes
	var jb []byte
	if jb, err = json.Marshal(m); err != nil {
		return err
	}

	if err = json.Unmarshal(jb, v); err != nil {
		return err
	}

	return nil
}
