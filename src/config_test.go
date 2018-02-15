package src

import (
	"testing"
	"reflect"
)

func compare(t *testing.T, fromFile *Config, fromCode *Config, equal bool) {
	if !reflect.DeepEqual(fromFile, fromCode) {
		if equal {
			t.Log("Mismatch between read and expected:",
				"\n-- (file): ", *fromFile,
			    "\n-- (code): ", *fromCode)
			t.Fail()
		}
	} else {
		if !equal {
			t.Log("Data structures were equal:")
			t.Log("-- (file): %+v\n", *fromFile)
			t.Log("-- (code): %+v\n", *fromCode)
			t.Fail()
		}
	}
}

func TestValidFile(t *testing.T) {
	config, _ := ReadYamlFile("./resources/certmaster_valid.yml")

	cmp := Config{
		Meta: Meta_S{
			Email: "someguy@example.com",
			Poll_Interval: "5s",
			Dry_Run: true,
		},
		Digital_Ocean: Digital_Ocean_S{
			Token: "token1",
		},
		Domains: []Domain_S{
			{
				Name: "example.com",
				Subdomains: []string{
					"app.example.com",
				},
			},
			{
				Name: "example2.com",
			},
		},
	}

	compare(t, config, &cmp, true)
}

func TestIncompleteFile(t *testing.T) {
	config, _ := ReadYamlFile("./resources/certmaster_incomplete.yml")

	cmp := Config{
		Digital_Ocean: Digital_Ocean_S{
			Token: "",
		},
		Domains: []Domain_S{
			{
				Name: "example.com",
				Subdomains: []string{},
			},
			{
				Name: "example2.com",
				Subdomains: []string{},
			},
		},
	}

	compare(t, config, &cmp, false)
}
