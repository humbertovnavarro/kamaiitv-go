package api

import "regexp"

var IsValidUsername = &regexp.Regexp{}
var IsValidPassword = &regexp.Regexp{}

func compileRegex() {
	IsValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	IsValidPassword = regexp.MustCompile(`^.{6,20}$`)
}
