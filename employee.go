package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName   string  `json:"Name"`
	EmpSalary float64 `json:"salary"`
	Email     string  `json:"email"`
}

// EmpName   string  `json:"Name"`
// EmpSalary float64 `json:"salary"`
// Email     string  `json:"email"`

// EmpName   string  `gorm:"type:varchar(20)"`
// EmpSalary float64 `gorm:"type:float"`
// Email     string  `gorm:"type:longtext"`
