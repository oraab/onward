package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/hashicorp/go-multierror"
	"strings"
)

type OnwardDb struct {
	db *gorm.DB
}

func NewOnwardDb(dbLocation string) (*OnwardDb, error) {
	err := sanitizeInput(dbLocation)
	if err != nil {
		return &OnwardDb{}, err
	}
	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	if err != nil {
		return &OnwardDb{}, errors.New(fmt.Sprintf("Failed to open DB %v: %v",dbLocation, err))
	}
	return &OnwardDb{db: db}, nil
}



func sanitizeInput(dbLocation string) error {
	var sanitizeResults error
	if len(dbLocation) == 0 {
		sanitizeResults = multierror.Append(sanitizeResults,
			errors.New("dbLocation not provided to OnwardDb constructor"))
	} else {
		dbLocationVerification := strings.Split(dbLocation,"/")
		if dbLocationVerification[0] != "" {
			sanitizeResults = multierror.Append(sanitizeResults,
				errors.New("provided dbLocation is invalid (relative path provided)"))
		}
		if strings.LastIndex(dbLocation,"/") == len(dbLocation)-1 {
			sanitizeResults = multierror.Append(sanitizeResults,
				errors.New("provided dbLocation is invalid (not a file)"))
		}
		if strings.LastIndex(dbLocationVerification[len(dbLocationVerification)-1],".db") !=
			len(dbLocationVerification[len(dbLocationVerification)-1])-3 {
			sanitizeResults = multierror.Append(sanitizeResults,
				errors.New("provided dbLocation is invalid (not an sqlite DB file)"))
		}
	}
	return sanitizeResults
}

