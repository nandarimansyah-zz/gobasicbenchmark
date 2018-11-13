package strregex

import "regexp"

var testRegexp = `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+$`

var myRegex *regexp.Regexp

func init() {
	r, err := regexp.Compile(testRegexp)
	if err != nil {
		panic(err)
	}
	myRegex = r
}

//IsMatchUsingMatchString ..
func IsMatchUsingMatchString(raw string) (bool, error) {
	return regexp.MatchString(testRegexp, raw)
}

//IsMatchUsingMatchStringCompiled ..
func IsMatchUsingMatchStringCompiled(raw string) (bool, error) {
	return myRegex.MatchString(raw), nil
}
