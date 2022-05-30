package db

import (
	"testing"
)

func TestSelect(t *testing.T) {
	crud := NewCrud("../../db/crud_test.db")
	// test no columns
	// test one column
	// test two columns
	// test no conditions
	// test one condition
	// test two conditions
	// test no table name provided

	result := crud.Select(make([]string,0),"LISTS",[]string{"TYPE=LIST"});
	if len(result) != 1 {
		t.Errorf("expected to get 1 row of result, got %v rows",len(result))
	}

}
