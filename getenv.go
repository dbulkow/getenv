/* Copyright (c) 2019 David Bulkow */

package getenv

import (
	"os"
	"strconv"
	"strings"
)

func Getenv(varname, defvalue string) string {
	env := os.Getenv(varname)

	if env == "" {
		return defvalue
	}

	return env
}

func GetenvInt(varname string, defvalue int) int {
	env := os.Getenv(varname)

	if env == "" {
		return defvalue
	}

	v, err := strconv.Atoi(env)
	if err != nil {
		return defvalue
	}

	return int(v)
}

func GetenvBool(varname string, defvalue bool) bool {
	env := os.Getenv(varname)

	if env == "" {
		return defvalue
	}

	if strings.ToLower(env) == "true" {
		return true
	}

	return false
}
