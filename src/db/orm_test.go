package db

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

// test dbLocation not provided
// test dbLocation not absolute path
// test dbLocation not a file
// test dbLocation not a DB file

func TestInputSanitization(t *testing.T) {
	var tests = []struct{
		name string
		path string
		noErrorType string
		expectedError string
	}{
		{"empty path test","","for missing dbLocation","dbLocation not provided to OnwardDb constructor"},
		{"non absolute path test", "not/absolute/path/test.db","for non absolute path dbLocation","provided dbLocation is invalid (relative path provided)"},
		{"non file path test", "/absolute/path/to/folder/","for non file dbLocation", "provided dbLocation is invalid (not a file)"},
		{"incorrect file name test", "/absolute/path/but/not_correct_file.txt","for wrong file name in dbLocation", "provided dbLocation is invalid (not an sqlite DB file)"},

	}

	for _, tt := range tests {
		_, err := NewOnwardDb(tt.path)
		if err == nil {
			t.Errorf("%v: Expected but did not receive an error from NewOnwardDb %v",tt.name, tt.noErrorType)
		}
		if !strings.Contains(err.Error(),tt.expectedError) {
			t.Errorf("%v: Provided error %v does not include correct error text", tt.name, err)
		}
	}
}

func TestNewOnwardDbCreation(t *testing.T) {
	db, err := NewOnwardDb(getDbLocation())
	if err != nil {
		t.Errorf("Received unexpected error when trying to create DB with valid location: %v", err)
	}
	if fmt.Sprintf("%v",reflect.TypeOf(db.db)) != "*gorm.DB" {
		t.Errorf("Expected type of received DB to be gorm.DB but received %v",reflect.TypeOf(db.db))
	}
}

func getDbLocation() string {
	dirName, err := os.UserHomeDir()
	if err != nil {
		panic("could not find user homedir")
	}
	return fmt.Sprintf("%v/git/onward/test/db/onwardTest.db",dirName)
}
