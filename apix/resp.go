package apix

type RespBody struct {
	Succeeded bool        `json:"succeeded"`      // 是否成功
	RespData  interface{} `json:"resp_data"`      // 返回数据
	Code      int         `json:"code,omitempty"` // 业务状态码
	Info      string      `json:"info,omitempty"` // 业务提示
	Desc      string      `json:"desc,omitempty"` // 异常提示，一般只出现在开发模式
}

type PageBody struct {
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Total  int64       `json:"total"`
	List   interface{} `json:"list"`
}

type RowAffectedBody struct {
	Rows int64 `json:"rows"`
}
