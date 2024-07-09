package api_test

import (
	"testing"

	"examble.com/crypto/api"
)

func TestGetRate(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not found")
	}
}
