package ginx

type SuccessfulRespBody struct {
	Succeeded bool        `json:"succeeded"`
	RespData  interface{} `json:"resp_data"`
}

type FailedRespBody struct {
	Succeeded bool        `json:"succeeded"`
	RespData  interface{} `json:"resp_data"`
	Code      int         `json:"code,omitempty"`
	Info      string      `json:"info"`
	Desc      string      `json:"desc,omitempty"`
}

type PageBody struct {
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Total  int         `json:"total"`
	List   interface{} `json:"list"`
}

type RowAffectedBody struct {
	Rows int64 `json:"rows"`
}
