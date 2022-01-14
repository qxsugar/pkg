package gormx

func LikeString(searchValue string) string {
	return "%" + searchValue + "%"
}
