package main

import (
	"ProjectDB/Vehicle"
	"fmt"
)

func main() {
	var s Vehicle.Store
	sql := Vehicle.MySqlconfig{"root", "localhost", "Jason@470", "3306","go"}
	s.Db, _= Vehicle.Connection(sql)


	input := Vehicle.Car{Id: 124, Name: "Hyundai", Model: "i10", EngineType: "Petrol"}

	res1:= s.Set(input)
	s.Set(Vehicle.Car{125,"maruti","swift","diesel"})
	fmt.Println(res1)

	res2:=s.Get(input.Id)
	fmt.Println(res2)

	res3 := s.Delete(input.Id)
	fmt.Println(res3)
}
