package ginx

type RespBody struct {
	Succeeded bool        `json:"succeeded"`      // Whether the operation was successful
	RespData  interface{} `json:"resp_data"`      // Returned data
	Code      int         `json:"code,omitempty"` // Business status code
	Info      string      `json:"info,omitempty"` // Business hints
	Desc      string      `json:"desc,omitempty"` // Exception hints, typically only appear in development mode
}

type PageBody struct {
	Offset int         `json:"offset"` // Offset
	Limit  int         `json:"limit"`  // Limit on the number of items
	Total  int64       `json:"total"`  // Total number of items
	List   interface{} `json:"list"`   // Data list
}

type RowAffectedBody struct {
	Rows int64 `json:"rows"` // Number of affected rows
}

type HttpError interface {
	HttpCode() int
}

type BusinessError interface {
	Code() int
	Info() string
	Desc() string
}
