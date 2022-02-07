package routes

import "regexp"

var IsValidUsername = &regexp.Regexp{}
var IsValidPassword = &regexp.Regexp{}
var IsValidObjectID = &regexp.Regexp{}

func CompileRegexp() {
	IsValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	IsValidPassword = regexp.MustCompile(`^.{6,20}$`)
	IsValidObjectID = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)
}
