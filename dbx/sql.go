package dbx

func WrapLike(searchValue string) string {
	return "%" + searchValue + "%"
}

func WrapLeftLike(searchValue string) string {
	return "%" + searchValue
}
