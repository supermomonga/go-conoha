package identity

import (
	"testing"
)

func TestGetVersions(t *testing.T) {
	cli := NewClient("")
	versions := cli.GetVersions()
	existence := false
	for _, version := range versions {
		if version.ID == "v2.0" {
			existence = true
		}
	}
	if existence == false {
		t.Error("Version 2.0 doesn't exist.")
	}
}

func TestGetVersion(t *testing.T) {
	cli := NewClient("")
	version := cli.GetVersion("v2.0")
	if version.ID != "v2.0" {
		t.Error("Wrong version returned.")
	}
}
