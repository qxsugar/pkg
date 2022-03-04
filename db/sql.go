package db

func LikeString(searchValue string) string {
	return "%" + searchValue + "%"
}
