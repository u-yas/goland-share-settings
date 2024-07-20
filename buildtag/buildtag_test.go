//go:build buildtag

package buildtag

import (
	"os"
	"testing"
)

func TestBuildTag(t *testing.T) {
	t.Log("Testing BuildTag")

}

func TestHogeBuildTag(t *testing.T) {
	t.Log("Testing HogeBuildTag")

}

func TestEnvTest(t *testing.T) {
	e := os.Getenv("APP_ENV")
	if e != "test" {
		t.Errorf("APP_ENV isn't set")
	}
}
