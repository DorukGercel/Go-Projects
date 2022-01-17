package cli_parser

import (
	"TODO_CLI/definitions"
	t "TODO_CLI/services/todo_handler"
	"errors"
	"flag"
	"os"
	"strconv"
)

// BoolFlagAttributes flags wout args
type BoolFlagAttributes struct {
	service     t.Service
	description string
	value       *bool
}

// StrFlagAttributes flags with args
type StrFlagAttributes struct {
	service     t.Service
	description string
	value       *string
	argValid    strFlagArgValid
}

// Arg validator
type strFlagArgValid func(arg t.StrServiceArg) bool

// Parse performs the validation and parsing of flags
func Parse() (t.Service, t.StrServiceArg, error) {
	if err := checkArgs(); err != nil {
		return t.None, nil, err
	}
	boolFlags, strFlags := initFlags()
	flag.Parse()
	return getService(boolFlags, strFlags)
}

// Check flag correctness
func checkArgs() error {
	argLen := len(os.Args)
	if argLen < 2 {
		return errors.New(definitions.ErrNfp)
	} else if os.Args[1][0] != '-' {
		return errors.New(definitions.ErrNfp)
	} else {
		arg := os.Args[1]
		if (isHelpFlag(arg) || isBoolFlag(arg)) && argLen == 2 {
			return nil
		} else if isStrFlag(arg) && argLen == 3 {
			return nil
		}
		return errors.New(definitions.ErrNarg)
	}
}

// Check if help flag
func isHelpFlag(arg string) bool {
	return arg[1:] == "h"
}

// Check if bool flag
func isBoolFlag(arg string) bool {
	boolFlagNames := getBoolFlags()
	_, ok := boolFlagNames[arg[1:]]
	return ok
}

// Check if str flag
func isStrFlag(arg string) bool {
	strFlagNames := getStrFlags()
	_, ok := strFlagNames[arg[1:]]
	return ok
}

// Returns map of bool flags
func getBoolFlags() map[string]BoolFlagAttributes {
	return map[string]BoolFlagAttributes{
		"v": {description: "version", value: new(bool), service: t.Version},
		"l": {description: "list all items", value: new(bool), service: t.ListAll},
		"c": {description: "list completed items", value: new(bool), service: t.ListCompleted},
	}
}

// Returns map of str flags
func getStrFlags() map[string]StrFlagAttributes {
	return map[string]StrFlagAttributes{
		"a": {description: "add new item", value: new(string), service: t.AddNew, argValid: validateStrArg},
		"m": {description: "mark as complete", value: new(string), service: t.MarkComplete, argValid: validateIntArg},
		"d": {description: "delete item", value: new(string), service: t.Delete, argValid: validateIntArg},
	}
}

// Returns the user inputted service
func getService(boolFlags map[string]BoolFlagAttributes, strFlags map[string]StrFlagAttributes) (t.Service, t.StrServiceArg, error) {
	boolService := getBoolService(boolFlags)
	if boolService != t.None {
		return boolService, nil, nil
	}
	strService, strArg := getStrService(strFlags)
	if strService != t.None {
		if err := validateStrFlagArg(strFlags, strService, strArg); err != nil {
			return t.None, nil, err
		}
		return strService, strArg, nil
	}
	return t.None, nil, errors.New(definitions.ErrNsf)
}

// Check if user triggered bool flag
func getBoolService(boolFlags map[string]BoolFlagAttributes) t.Service {
	for _, val := range boolFlags {
		if *val.value == true {
			return val.service
		}
	}
	return t.None
}

// Check if user triggered str flag
func getStrService(strFlags map[string]StrFlagAttributes) (t.Service, t.StrServiceArg) {
	for _, val := range strFlags {
		if len(*val.value) > 0 {
			return val.service, *val.value
		}
	}
	return t.None, nil
}

// Validate the arg for string flag
func validateStrFlagArg(strFlags map[string]StrFlagAttributes, s t.Service, arg t.StrServiceArg) error {
	for _, val := range strFlags {
		if val.service == s {
			if val.argValid(arg) {
				return nil
			} else {
				return errors.New(definitions.ErrNarg)
			}
		}
	}
	return errors.New(definitions.ErrInter)
}

// Validate the arg for string flag with string arg
func validateStrArg(arg t.StrServiceArg) bool {
	argStr, ok := arg.(string)
	if ok && len(argStr) > 0 {
		return true
	}
	return false
}

// Validate the arg for string flag with int arg
func validateIntArg(arg t.StrServiceArg) bool {
	argStr, ok := arg.(string)
	if !ok {
		return false
	}
	n, err := strconv.Atoi(argStr)
	if n >= 0 && err == nil {
		return true
	}
	return false
}

// Initialize accepted flags
func initFlags() (map[string]BoolFlagAttributes, map[string]StrFlagAttributes) {
	boolFlags := getBoolFlags()
	for key, val := range boolFlags {
		flag.BoolVar(boolFlags[key].value, key, false, val.description)
	}

	strFlags := getStrFlags()
	for key, val := range strFlags {
		flag.StringVar(strFlags[key].value, key, "", val.description)
	}

	return boolFlags, strFlags
}
