package Vehicle

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	testcases := []struct {
		desc           string
		input          Car
		expectedoutput bool
	}{
		{"Valid data", Car{Id: 123, Name: "Hyundai", Model: "Eon", EngineType: "Petrol"}, true},
		{"Invalid data ", Car{Id: -1, Name: "", Model: "", EngineType: ""}, false},
	}
	var s Store
	var err error

	s.Db, err = Connection(MySqlconfig{"root", "localhost", "Jason@470", "3306", "go"})
	if err != nil {
		fmt.Println(err)
	}
	for i, tc := range testcases {

		actualoutput := s.Set(tc.input)
		if actualoutput != tc.expectedoutput {
			t.Errorf("test case %s Expected %v Got %v testcase %v", tc.desc, tc.expectedoutput, actualoutput, i+1)
		}
	}

}

func TestGet(t *testing.T) {
	testcases := []struct {
		desc           string
		Id             int
		expectedoutput Car
	}{
		{"ValId Id", 123, Car{Id: 123, Name: "Hyundai", Model: "Eon", EngineType: "Petrol"}},
		{"InvalId Id", 0, Car{Id: 0, Name: "", Model: "", EngineType: ""}},
		{"Negative Id", -1, Car{Id: 0, Name: "", Model: "", EngineType: ""}},
	}
	var s Store
	var err error
	s.Db, err = Connection(MySqlconfig{"root", "localhost", "Jason@470", "3306", "go"})
	if err != nil {
		fmt.Println(err)
	}
	for i, tc := range testcases {
		actualoutput := s.Get(tc.Id)
		if actualoutput != tc.expectedoutput {
			t.Errorf("test case %s Expected %v Got %v testcase %v", tc.desc, tc.expectedoutput, actualoutput, i+1)
		}
	}
}

func TestDelete(t *testing.T) {
	testcases := []struct {
		desc           string
		Id             int
		expectedoutput bool
	}{
		{"If Id is in table", 123, true},
		{"If Id is not in table", 1, false},
	}

	var s Store
	var err error
	s.Db, err = Connection(MySqlconfig{"root", "localhost", "Jason@470", "3306", "go"})
	if err != nil {
		fmt.Println(err)
	}

	for i, tc := range testcases {
		actualoutput := s.Delete(tc.Id)
		if actualoutput != tc.expectedoutput {
			t.Errorf("test case %s Expected %v Got %v testcase %v", tc.desc, tc.expectedoutput, actualoutput, i+1)
		}
	}
}
