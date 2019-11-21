/* Copyright (c) 2019 David Bulkow */

package getenv

import (
	"os"
	"strconv"
	"strings"
)

type Env struct {
	prefix string
}

func NewEnv(prefix string) *Env {
	return &Env{prefix: prefix}
}

func (e *Env) varname(suffix string) string {
	if e.prefix == "" {
		return suffix
	}
	return strings.Join([]string{e.prefix, suffix}, "_")
}

func (e *Env) Getenv(suffix, defvalue string) string {
	env := os.Getenv(e.varname(suffix))

	if env == "" {
		return defvalue
	}

	return env
}

func (e *Env) GetenvInt(suffix string, defvalue int) int {
	env := os.Getenv(e.varname(suffix))

	if env == "" {
		return defvalue
	}

	v, err := strconv.Atoi(env)
	if err != nil {
		return defvalue
	}

	return int(v)
}

func (e *Env) GetenvBool(suffix string, defvalue bool) bool {
	env := os.Getenv(e.varname(suffix))

	if env == "" {
		return defvalue
	}

	if strings.ToLower(env) == "true" {
		return true
	}

	return false
}

var defEnv = &Env{}

func Getenv(varname, defvalue string) string {
	return defEnv.Getenv(varname, defvalue)
}

func GetenvInt(varname string, defvalue int) int {
	return defEnv.GetenvInt(varname, defvalue)
}

func GetenvBool(varname string, defvalue bool) bool {
	return defEnv.GetenvBool(varname, defvalue)
}
