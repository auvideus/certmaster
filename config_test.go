package certmaster

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
		Meta: Meta{
			Email: "someguy@example.com",
			Poll_Interval: "5m",
			Dry_Run: true,
		},
		Digital_Ocean: DigitalOcean{
			Token: "token1",
		},
		Domains: []Domain{
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
	_, err := ReadYamlFile("./resources/certmaster_incomplete.yml")

	if err == nil {
		t.Error("missing required fields, should have failed")
	}
}
