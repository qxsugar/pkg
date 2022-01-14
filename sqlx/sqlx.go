package sqlx

func LikeString(searchValue string) string {
	return "%" + searchValue + "%"
}
