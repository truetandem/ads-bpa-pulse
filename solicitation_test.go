package main

import (
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestSolicitationToString(t *testing.T) {
	expected := "Title: Test\n\n  - ID: TBD"
	s := Solicitation{
		Title: "Test",
		Properties: map[string]string{
			"ID": "TBD",
		},
	}

	if s.ToString() != expected {
		t.FailNow()
	}
}

func TestSolicitationChecksum(t *testing.T) {
	expected := "f866a997e7b28e85a438e7415908e4f51edaf43e"
	s := Solicitation{
		Title: "Test",
		Properties: map[string]string{
			"ID": "TBD",
		},
	}

	if s.Checksum() != expected {
		t.FailNow()
	}
}

func TestSolicitationGet(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	s := Solicitation{
		Title: "Test",
		Properties: map[string]string{
			"ID": "TBD",
		},
	}

	_, err = s.Get(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSolicitationSave(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	s := Solicitation{
		Title: "Test",
		Properties: map[string]string{
			"ID": "TBD",
		},
	}

	err = s.Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
