package cli

import "github.com/spf13/pflag"

func AddStringFlags(flagSet *pflag.FlagSet, flagList []string) {
	for _, flag := range flagList {
		description := "Optional flag '" + flag + "'"
		flagSet.StringP(flag, "", "", description)
	}
}

func ParseStringFlagList(flagSet *pflag.FlagSet, flagList []string) map[string]string {
	list := make(map[string]string)

	for _, flag := range flagList {
		flagValue, err := flagSet.GetString(flag)

		if err == nil && flagValue != "" {
			list[flag] = flagValue
		}
	}

	return list
}
