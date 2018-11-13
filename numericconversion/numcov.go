package numericconversion

import "strconv"

//MParseBool ..
func MParseBool(raw string) (bool, error) {
	return strconv.ParseBool("true")
}

//MParseInt ..
func MParseInt(raw string) (int64, error) {
	return strconv.ParseInt(raw, 10, 64)
}

//MParseFloat ..
func MParseFloat(raw string) (float64, error) {
	return strconv.ParseFloat(raw, 64)
}
