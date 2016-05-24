// +build appengine

package main

import (
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestSendEmail(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("POST", "/mail", nil)
	if err != nil {
		t.FailNow()
	}

	err = sendEmail(r, "sender@test.com", []string{"rcpt@test.com"}, "Test subject", "", "")
	if err != nil {
		t.FailNow()
	}
}
