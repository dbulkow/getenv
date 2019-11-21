package getenv_test

import (
	"os"
	"testing"

	"github.com/dbulkow/getenv"
)

var varname = "GETENV_TEST_VARIABLE"

func TestGetenv(t *testing.T) {
	os.Clearenv()
	os.Setenv(varname, "somevalue")

	value := getenv.Getenv(varname, "foo")
	if value != "somevalue" {
		t.Fatalf("expected \"somevalue\" got \"%s\"", value)
	}
}

func TestGetenvUnset(t *testing.T) {
	os.Clearenv()

	value := getenv.Getenv("GETENV_UNSET_VARIABLE", "unset")
	if value != "unset" {
		t.Fatalf("expected \"somevalue\" got \"%s\"", value)
	}
}

func TestGetenvInt(t *testing.T) {
	os.Clearenv()
	os.Setenv(varname, "5150")

	value := getenv.GetenvInt(varname, 1234)
	if value != 5150 {
		t.Fatalf("expected \"5150\" got \"%d\"", value)
	}
}

func TestGetenvIntUnset(t *testing.T) {
	os.Clearenv()
	os.Setenv(varname, "5150")

	value := getenv.GetenvInt("GETENV_UNSET_VARIABLE", 1234)
	if value != 1234 {
		t.Fatalf("expected \"somevalue\" got \"%d\"", value)
	}
}

func TestGetenvIntNaN(t *testing.T) {
	os.Clearenv()
	os.Setenv(varname, "abcdefg")

	value := getenv.GetenvInt(varname, 1234)
	if value != 1234 {
		t.Fatalf("expected \"1234\" got \"%d\"", value)
	}
}

func TestGetenvBool(t *testing.T) {
	os.Clearenv()
	os.Setenv(varname, "TRUE")

	value := getenv.GetenvBool(varname, false)
	if value != true {
		t.Fatalf("expected \"true\" got \"%t\"", value)
	}
}

func TestGetenvBoolUnset(t *testing.T) {
	os.Clearenv()

	value := getenv.GetenvBool("GETENV_UNSET_VARIABLE", false)
	if value != false {
		t.Fatalf("expected \"somevalue\" got \"%t\"", value)
	}
}

func TestGetenvBoolBadValue(t *testing.T) {
	os.Clearenv()
	os.Setenv(varname, "whatever")

	value := getenv.GetenvBool(varname, true)
	if value != false {
		t.Fatalf("expected \"false\" got \"%t\"", value)
	}
}

func TestGetenvBoolBadValueFalseDefault(t *testing.T) {
	os.Clearenv()
	os.Setenv(varname, "whatever")

	value := getenv.GetenvBool(varname, false)
	if value != false {
		t.Fatalf("expected \"false\" got \"%t\"", value)
	}
}

func TestEnvGetenv(t *testing.T) {
	os.Clearenv()
	os.Setenv(varname, "somevalue")

	env := getenv.NewEnv("GETENV")

	value := env.Getenv("TEST_VARIABLE", "somedefault")
	if value != "somevalue" {
		t.Fatalf("expected \"somevalue\" got \"%s\"", value)
	}
}
