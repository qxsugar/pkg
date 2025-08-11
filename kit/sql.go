package kit

// WrapLike wraps the search value with SQL LIKE wildcards on both sides.
// Example: WrapLike("test") returns "%test%"
func WrapLike(searchValue string) string {
	return "%" + searchValue + "%"
}

// WrapLeftLike wraps the search value with SQL LIKE wildcard on the left side.
// Example: WrapLeftLike("test") returns "%test"
func WrapLeftLike(searchValue string) string {
	return "%" + searchValue
}
