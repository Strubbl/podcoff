package main

// Filter checks the given podcast item for a condition matching a keyword in a field of the feed item
type Filter struct {
	Condition string // IN or NOT
	Keyword   string // search keyword
	Field     string // currently only title
}
