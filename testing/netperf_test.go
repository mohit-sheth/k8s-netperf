package netperf

import (
	"testing"

	"github.com/jtaleric/k8s-netperf/pkg/config"
)

// TestParseConf Test for success. Ensure we successfully parse a good config file
func TestParseConf(t *testing.T) {
	file := "test-config.yml"
	_, err := config.ParseConf(file)
	if err != nil {
		t.Fatal("Parsing config file failed")
	}
}

// TestParseConf Test for success. Ensure we successfully parse the default config
func TestShippingConf(t *testing.T) {
	file := "../netperf.yml"
	_, err := config.ParseConf(file)
	if err != nil {
		t.Fatal("Parsing config file failed")
	}
}

// TestMissingParseConf Testing for failure. Test profile regex
func TestMissingParseConf(t *testing.T) {
	file := "test-bad-missing-config.yml"
	_, err := config.ParseConf(file)
	if err == nil {
		t.Fatal("Parsing config file should have failed but succeeded")
	}
}

// TestBadParseConf Test for failure. User leaves out a config field
func TestBadParseConf(t *testing.T) {
	file := "test-bad-profile-config.yml"
	_, err := config.ParseConf(file)
	if err == nil {
		t.Fatal("Parsing config file should have failed but succeeded")
	}
}
