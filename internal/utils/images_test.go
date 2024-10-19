package utils

import "testing"

func TestShortImageID(t *testing.T) {
	src := "sha256:2503e324e27050f4d3fd21d25147ca108840864941d96097c2633cd9232f5088"
	expected := "2503e324e270"

	result := ShortImageID(src)

	if result != expected {
		t.Errorf("got: %s, want: %s", result, expected)
	}
}
