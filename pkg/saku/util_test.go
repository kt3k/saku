package saku

import "testing"

func TestPrependEmoji(t *testing.T) {
	if prependEmoji("🔎", "mag", true) != "🔎  mag" {
		t.Error("prepend emoji when the flag is true")
	}

	if prependEmoji("🔎", "mag", false) != "mag" {
		t.Error("remove emoji when the flag is false")
	}
}
