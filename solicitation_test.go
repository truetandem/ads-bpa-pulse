package main

import "testing"

func TestSolicitationToString(t *testing.T) {
	s := Solicitation{
		Title: "Test",
		Properties: map[string]string{
			"ID": "TBD",
		},
	}
	expected := "Title: Test\n\n  - ID: TBD"

	if s.ToString() != expected {
		t.FailNow()
	}
}

func TestSolicitationChecksum(t *testing.T) {
	s := Solicitation{
		Title: "Test",
		Properties: map[string]string{
			"ID": "TBD",
		},
	}
	expected := "f866a997e7b28e85a438e7415908e4f51edaf43e"

	if s.Checksum() != expected {
		t.FailNow()
	}
}
