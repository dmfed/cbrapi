package cbrapi

import (
	"fmt"
	"testing"
)

func Test_Keeper_Connects_And_Unmarshalls_Dictionary(t *testing.T) {
	k, err := codeKeeperFromAPI()
	if err != nil {
		fmt.Println(k, err)
		t.Fail()
	}
	if k["USD"].NameENG != "US Dollar" {
		t.Fail()
	}
	fromfile, err := codeKeeperFromFile()
	if fromfile["USD"].NameENG != "US Dollar" ||
		err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
