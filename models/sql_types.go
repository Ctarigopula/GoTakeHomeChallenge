package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Metadata map[string]interface{}

// Value marshals Metadata
func (m Metadata) Value() (value driver.Value, err error) {
	marshalled, err := json.Marshal(m)
	if err != nil {
		return
	}

	value = driver.Value(string(marshalled))
	if value == "null" {
		value = "{}"
	}

	return
}

// Scan unmarshals the stored string into Metadata
func (m *Metadata) Scan(src interface{}) (err error) {
	switch b := src.(type) {
	case string:
		var values Metadata
		err = json.Unmarshal([]byte(b), &values)
		if err == nil {
			*m = values
		}
	case []byte:
		var values Metadata
		err = json.Unmarshal(b, &values)
		if err == nil {
			*m = values
		}
	default:
		return errors.New("data from sql drive was not a []byte")
	}

	return
}

// IntList is wrapper around a integer list. An array of integers is saved in
// the database like so: {1,2,3}
type IntList []int

// Contains checks whether or not the value is present
func (m IntList) Contains(i int) bool {
	for _, val := range m {
		if i == val {
			return true
		}
	}

	return false
}

// Value marshals Metadata
func (m IntList) Value() (value driver.Value, err error) {
	if len(m) == 0 {
		value = "{}"
		return
	}

	value = "{" + strings.Trim(strings.Replace(fmt.Sprint(m), " ", ",", -1), "[]") + "}"

	return
}

// Scan unmarshals the stored string into Metadata
func (m *IntList) Scan(src interface{}) (err error) {
	switch src := src.(type) {
	case []byte:
		stringValue := string(src)
		trimmed := strings.Trim(stringValue, "{")
		trimmed = strings.Trim(trimmed, "}")
		if trimmed == "" {
			return
		}

		values := strings.Split(trimmed, ",")
		ints := make([]int, len(values))
		for idx, s := range values {
			i, _ := strconv.Atoi(strings.Trim(s, " "))
			ints[idx] = i
		}
		*m = ints
	case string:
		trimmed := strings.Trim(src, "{")
		trimmed = strings.Trim(trimmed, "}")
		if trimmed == "" {
			return
		}

		values := strings.Split(trimmed, ",")
		ints := make([]int, len(values))
		for idx, s := range values {
			i, _ := strconv.Atoi(strings.Trim(s, " "))
			ints[idx] = i
		}
		*m = ints
	default:
		return errors.New("data from sql drive was not a []byte")
	}

	return
}
