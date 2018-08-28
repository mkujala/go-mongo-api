package models

import (
	"fmt"
	"strings"
)

// Validator is used to validate fields
type Validator struct {
	err error
}

type value interface{}

// Required validates that required field is not empty
func (v *Validator) Required(val value) bool {
	// fmt.Printf("type of value: %T\n", val) 		// DEBUG

	if v.err != nil {
		return false
	}

	switch val.(type) {
	default:
		fmt.Printf("unexpected type %T", val)
		return false
	case string:
		if strings.TrimSpace(val.(string)) == "" {
			return false
		}
		return true
	}
}
