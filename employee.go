package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName   string  `json:"Name"`
	EmpSalary float64 `json:"salary"`
	Email     string  `json:"email"`
}


