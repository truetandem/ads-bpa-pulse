package main

import (
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestSolicitationString(t *testing.T) {
	expected := "Title: Test\n\n  - ID: TBD"
	s := Solicitation{
		Title: "Test",
		Properties: map[string]string{
			"ID": "TBD",
		},
	}

	if s.String() != expected {
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

	if s.Sum() != expected {
		t.FailNow()
	}
}

func TestSolicitationChecksumAfterUpdate(t *testing.T) {
	expected := "f866a997e7b28e85a438e7415908e4f51edaf43e"
	s := Solicitation{
		Title: "Test",
		Properties: map[string]string{
			"ID": "TBD",
		},
	}

	if s.Sum() != expected {
		t.FailNow()
	}

	s.Properties["ID"] = "10000"
	if s.Sum() == expected {
		t.FailNow()
	}

	s.Properties["ID"] = "TBD"
	s.Properties["New"] = "Test"
	if s.Sum() == expected {
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
	if err != nil && err != datastore.ErrNoSuchEntity {
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

func TestSolicitationUpdate(t *testing.T) {
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

	s.Properties["ID"] = "TBD"
	err = s.Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
