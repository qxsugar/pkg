package kit

type RespBody struct {
	Succeeded bool   `json:"succeeded"`      // Whether the operation was successful
	RespData  any    `json:"resp_data"`      // Returned data
	Code      int    `json:"code,omitempty"` // Business status code
	Info      string `json:"info,omitempty"` // Business hints
	Desc      string `json:"desc,omitempty"` // Exception hints, typically only appear in development mode
}

type PageBody struct {
	Offset int   `json:"offset"` // Offset
	Limit  int   `json:"limit"`  // Limit on the number of items
	Total  int64 `json:"total"`  // Total number of items
	List   any   `json:"list"`   // Data list
}

type RowAffectedBody struct {
	Rows int64 `json:"rows"` // Number of affected rows
}

type BusinessError interface {
	Code() int    // Business error code
	Info() string // Business error message
	Desc() string // Business error description
}
