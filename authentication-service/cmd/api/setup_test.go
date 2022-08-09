package main

import (
	"authentication/test"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	repo := test.NewPostgresTestRepository(nil)
	testApp.Repo = repo

	os.Exit(m.Run())
}
