/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2017 by authors and contributors.
 */

package datadog

import "encoding/json"

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// GetBool is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetBool(v *bool) (bool, bool) {
	if v != nil {
		return *v, true
	}

	return false, false
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// GetInt is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetIntOk(v *int) (int, bool) {
	if v != nil {
		return *v, true
	}

	return 0, false
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// GetString is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetStringOk(v *string) (string, bool) {
	if v != nil {
		return *v, true
	}

	return "", false
}

// JsonNumber is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func JsonNumber(v json.Number) *json.Number { return &v }

// GetJsonNumber is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetJsonNumberOk(v *json.Number) (json.Number, bool) {
	if v != nil {
		return *v, true
	}

	return "", false
}

// Precision is a helper routine that allocates a new precision value
// to store v and returns a pointer to it.
func Precision(v PrecisionT) *PrecisionT { return &v }

// GetPrecision is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetPrecision(v *PrecisionT) (PrecisionT, bool) {
	if v != nil {
		return *v, true
	}

	return PrecisionT(""), false
}

// Width is a helper routine that allocates a new width value
// to store v and returns a pointer to it.
func Width(v WidthS) *WidthS { return &v }

// GetWidthOk is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetWidthOk(v *WidthS) (WidthS, bool) {
	if v != nil {
		return *v, true
	}

	return WidthS(""), false
}

// Height is a helper routine that allocates a new width value
// to store v and returns a pointer to it.
func Height(v HeightS) *HeightS { return &v }

// GetHeightOk is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetHeightOk(v *HeightS) (HeightS, bool) {
	if v != nil {
		return *v, true
	}

	return HeightS(""), false
}
