package server

import (
	"fmt"
	"net/url"
	"strconv"
)

func intFromForm(key string, form url.Values) (int, error) {
	strVal := form.Get(key)
	val, err := strconv.Atoi(strVal)
	if err != nil {
		return val, fmt.Errorf("Got value from form with key %s, but it is not an integer as expected. Got: '%s'", key, strVal)
	}
	return val, nil
}

func boolFromForm(key string, form url.Values) (bool, error) {
	strVal := form.Get(key)
	val, err := strconv.ParseBool(strVal)
	if err != nil {
		return val, fmt.Errorf("Got value from form with key %s, but it is not a boolean as expected. Got: '%s'", key, strVal)
	}
	return val, nil
}
