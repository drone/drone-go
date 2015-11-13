package drone

import (
	"encoding/json"
	"strconv"
)

// StringSlice representes a string or an array of strings.
type StringSlice struct {
	parts []string
}

func (e *StringSlice) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	p := make([]string, 0, 1)
	if err := json.Unmarshal(b, &p); err != nil {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		p = append(p, s)
	}

	e.parts = p
	return nil
}

func (e *StringSlice) Len() int {
	if e == nil {
		return 0
	}
	return len(e.parts)
}

func (e *StringSlice) Slice() []string {
	if e == nil {
		return nil
	}
	return e.parts
}

// StringInt representes a string or an integer value.
type StringInt struct {
	value string
}

func (e *StringInt) UnmarshalJSON(b []byte) error {
	var num int
	err := json.Unmarshal(b, &num)
	if err == nil {
		e.value = strconv.Itoa(num)
		return nil
	}
	return json.Unmarshal(b, &e.value)
}

func (e StringInt) String() string {
	return e.value
}

// StringMap representes a string or a map of strings.
// StringMap representes a string or a map of strings.
type StringMap struct {
	parts map[string]string
}

func (e *StringMap) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	p := map[string]string{}
	if err := json.Unmarshal(b, &p); err != nil {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		p[""] = s
	}

	e.parts = p
	return nil
}

func (e *StringMap) Len() int {
	if e == nil {
		return 0
	}
	return len(e.parts)
}

func (e *StringMap) String() (str string) {
	if e == nil {
		return
	}
	for _, val := range e.parts {
		return val // returns the first string value
	}
	return
}

func (e *StringMap) Map() map[string]string {
	if e == nil {
		return nil
	}
	return e.parts
}

func NewStringMap(parts map[string]string) StringMap {
	return StringMap{parts}
}
