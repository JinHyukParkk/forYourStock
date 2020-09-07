package model

// json 형식
type CompanyDisclosure struct {
	Status    string       `json:"status"`
	Message   string       `json:"message"`
	PageNo    string       `json:"page_no"`
	PageCnt   string       `json:"page_count"`
	TotalCnt  string       `json:"total_count"`
	TotalPage string       `json:"total_page"`
	List      []Disclosure `json:"list"`
}

type Disclosure struct {
	CorpCode  string `json:"corp_code"`
	CorpName  string `json:"corp_name"`
	StockCode string `json:"stock_code"`
	CorpCls   string `json:"corp_cls"`
	ReportNm  string `json:"report_nm"`
	RceptNo   string `json:"rcept_no"`
	FlrNm     string `json:"flr_nm"`
	RceptDt   string `json:"rcept_dt"`
	Rm        string `json:"rm"`
}
