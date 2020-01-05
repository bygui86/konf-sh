package flags

import "flag"

func GetString(flagKey, flagDefault, errorMsg string) (string,error) {
	return *flag.String(flagKey, flagDefault, errorMsg), nil
}
