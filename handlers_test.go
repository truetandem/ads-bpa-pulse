package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateReturnsOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Update))
	defer server.Close()

	if resp, err := http.DefaultClient.Get(server.URL); err != nil || resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
}

func TestUpdateReturnsSolicitations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Update))
	defer server.Close()

	resp, err := http.DefaultClient.Get(server.URL)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
	defer resp.Body.Close()

	dump, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.FailNow()
	}

	var solicitations []Solicitation
	err = json.Unmarshal(dump, &solicitations)
	if err != nil {
		t.FailNow()
	}
}
