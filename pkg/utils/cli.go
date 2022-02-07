package utils

import "strings"

func GetUrfaveFlagName(name, short string) string {
	return strings.Join([]string{name, short}, ", ")
}
