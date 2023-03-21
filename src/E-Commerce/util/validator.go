package util

import "regexp"

func IsEmail(value string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(value)
}

func IsMobile(value string) bool {
	mobileRegex := regexp.MustCompile(`^[1-9]\d{9}$`)
	return mobileRegex.MatchString(value)
}
