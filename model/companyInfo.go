package model

type CompanyListStruct struct {
	CompanyList	[]CompanyStruct	`json:"company"`
}

type CompanyStruct struct {
	Name	string	`json:"name"`
	Code	string	`json:"code"`
 }

 type CompanyXmlResult struct {
	 Result []CompanyXmlList `xml:"result"`
 }

 type CompanyXmlList struct {
	List CompanyXml `xml:"list"`
}
type CompanyXml struct {
	corp_code	string	`xml:"corp_code"`
	corp_name	string	`xml:"corp_name"`
	stock_code	string	`xml:"stock_code"`
	modify_date	string	`xml:"modify_date"`
}
