package Vehicle

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
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

	db, mock, err := sqlmock.New()
	defer db.Close()

	s.Db = db
	if err != nil {
		fmt.Println(err)
	}
	for i, tc := range testcases {
		mock.ExpectExec("INSERT INTO car values;").WithArgs(tc.input.Id, tc.input.Name, tc.input.Model, tc.input.EngineType).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(err)
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
	db, mock, err := sqlmock.New()
	s.Db = db
	if err != nil {
		fmt.Println(err)
	}
	for i, tc := range testcases {
		row := mock.NewRows([]string{"id", "name", "model", "engineType"}).AddRow(tc.expectedoutput.Id, tc.expectedoutput.Name, tc.expectedoutput.Model, tc.expectedoutput.EngineType)
		mock.ExpectQuery("select (.+) from car where Id=?").WithArgs(tc.Id).WillReturnRows(row).WillReturnError(err)
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
		{"If Id is not in table", 1, true},
	}
	db, mock, err := sqlmock.New()
	var s Store

	s.Db = db
	if err != nil {
		fmt.Println(err)
	}

	for i, tc := range testcases {
		mock.ExpectExec("delete from car where id=?").WithArgs(tc.Id).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(err)
		actualoutput := s.Delete(tc.Id)
		if actualoutput != tc.expectedoutput {
			t.Errorf("test case %s Expected %v Got %v testcase %v", tc.desc, tc.expectedoutput, actualoutput, i+1)
		}
	}
}
