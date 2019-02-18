package cli

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

const cliDefaultValueTagName = "default"

func AssignFlagsFromInterfaceFields(flagSet *pflag.FlagSet, v interface{}) {
	values := reflect.ValueOf(v).Elem()

	for i := 0; i < values.NumField(); i++ {
		typeField := values.Type().Field(i)
		assignFieldAsFlag(typeField, flagSet)
	}
}

func assignFieldAsFlag(field reflect.StructField, flagSet *pflag.FlagSet) {
	flagName := flagNameForField(field)

	// TODO: Support all types of pflag here in the future
	switch t := field.Type.Name(); t {
	case "bool":
		flagSet.BoolP(flagName, "", defaultBoolForField(field), descriptionForFlag(flagName))
	case "int":
		flagSet.IntP(flagName, "", defaultIntForField(field), descriptionForFlag(flagName))
	case "string":
		flagSet.StringP(flagName, "", defaultStrForField(field), descriptionForFlag(flagName))
	}
}

func SetInterfaceFieldsFromFlags(flagSet *pflag.FlagSet, v interface{}, requireExplicitChange bool) {
	values := reflect.ValueOf(v).Elem()

	for i := 0; i < values.NumField(); i++ {
		valueField := values.Field(i)
		typeField := values.Type().Field(i)
		flagName := flagNameForField(typeField)
		flagValueChanged := flagSet.Changed(flagName)

		if requireExplicitChange == true && flagValueChanged == false {
			continue
		}

		switch t := typeField.Type.Name(); t {
		case "bool":
			boolVal, _ := flagSet.GetBool(flagName)
			valueField.SetBool(boolVal)
		case "int":
			intVal, _ := flagSet.GetInt(flagName)
			valueField.SetInt(int64(intVal))
		case "string":
			strVal, _ := flagSet.GetString(flagName)
			valueField.SetString(strVal)
		}
	}
}

func defaultBoolForField(field reflect.StructField) bool {
	defaultValueStr := field.Tag.Get(cliDefaultValueTagName)
	defaultValueBool, err := strconv.ParseBool(defaultValueStr)
	if err != nil {
		log.Fatalf("Error while parsing CLI default value: %s\n", err)
	}

	return defaultValueBool
}

func defaultIntForField(field reflect.StructField) int {
	defaultValueStr := field.Tag.Get(cliDefaultValueTagName)
	defaultValueInt, err := strconv.Atoi(defaultValueStr)
	if err != nil {
		log.Fatalf("Error while parsing CLI default value: %s\n", err)
	}

	return defaultValueInt
}

func defaultStrForField(field reflect.StructField) string {
	return field.Tag.Get(cliDefaultValueTagName)
}

func descriptionForFlag(flagName string) string {
	return flagName + " argument (is maybe required to execute this command)"
}

func flagNameForField(field reflect.StructField) string {
	return strings.ToLower(field.Name)
}
