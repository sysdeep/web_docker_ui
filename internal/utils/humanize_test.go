package utils

import "testing"

func TestByteCountToDisplaySize(t *testing.T) {

	dataset := map[int64]string{
		10:                 "10 bytes",
		1024:               "1 KB",
		1024 * 2:           "2 KB",
		1024 * 1024:        "1 MB",
		1024 * 1024 * 1024: "1 GB",
	}

	for key, expected := range dataset {

		result := ByteCountToDisplaySize(key)
		if result != expected {
			t.Errorf("have: %s, want: %s", result, expected)
		}
	}

}
