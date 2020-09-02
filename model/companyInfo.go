package model

type CompanyListStruct struct {
	CompanyList	[]CompanyStruct	`json:"company"`
}

type CompanyStruct struct {
	Name	string	`json:"name"`
	Code	string	`json:"code"`
 }