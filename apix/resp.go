package apix

type PageBody struct {
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Total  int64       `json:"total"`
	List   interface{} `json:"list"`
}

type RowAffectedBody struct {
	Rows int64 `json:"rows"`
}
