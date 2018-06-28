/* Copyright (c) 2018 David Bulkow */

package getenv

import (
	"os"
	"strconv"
)

func Getenv(varname, defvalue string) string {
	env := os.Getenv(varname)

	if env == "" {
		return defvalue
	}

	return env
}

func GetenvInt(varname string, defvalue int) (int, error) {
	env := os.Getenv(varname)

	if env == "" {
		return defvalue, nil
	}

	v, err := strconv.Atoi(env)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}
