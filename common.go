package gocommon

import (
	"log"
	"regexp"
	"runtime"
)

// GetFileLine is a debug utility that get the file&line of caller.
func GetFileLine() (string, int) {
	_, fn, line, _ := runtime.Caller(1)
	return fn, line
}

// CheckedDeferFunc0 is useful for report error in deferred function without argument.
func CheckedDeferFunc0(f func() error) {
	err := f()
	if err != nil {
		log.Println(err)
		_, fn, line, _ := runtime.Caller(1)
		log.Println(fn, line)
	}
}

// MatchPattern function checks if the str matches specified regular expression pattern.
func MatchPattern(str, pattern string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(str)
}
