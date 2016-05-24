// +build appengine

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestUpdateReturnsOK(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/update", nil)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	Update(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected Status Code [%v] Got [%v]", w.Code, http.StatusOK)
	}
}

func TestUpdateReturnsSolicitations(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/update", nil)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	Update(w, r)

	dump, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.FailNow()
	}

	var solicitations []Solicitation
	err = json.Unmarshal(dump, &solicitations)
	if err != nil {
		t.FailNow()
	}
}
