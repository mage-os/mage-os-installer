package detector

import (
	"strings"
	"testing"
)

func TestParseThemeID_ReturnsTheTrimmedID(t *testing.T) {
	id, err := parseThemeID("4\n", hyvaThemePath)

	if err != nil {
		t.Fatalf("expected an ID, got error: %v", err)
	}
	if id != "4" {
		t.Errorf("ID is %q, expected \"4\"", id)
	}
}

func TestParseThemeID_FailsWhenTheThemeIsNotRegistered(t *testing.T) {
	_, err := parseThemeID("\n", hyvaThemePath)

	if err == nil {
		t.Fatal("expected an error for an empty query result")
	}
	if !strings.Contains(err.Error(), hyvaThemePath) {
		t.Errorf("error %q should name the missing theme", err)
	}
}

func TestParseThemeID_FailsOnAMultiRowResult(t *testing.T) {
	_, err := parseThemeID("4\n7\n", hyvaThemePath)

	if err == nil {
		t.Fatal("expected an error for a multi-row result")
	}
}

func TestParseThemeID_FailsOnNoiseInTheOutput(t *testing.T) {
	_, err := parseThemeID("Warning: something\n4\n", hyvaThemePath)

	if err == nil {
		t.Fatal("expected an error when the output is not a bare ID")
	}
}
