package main

import "errors"

// Filter checks the given podcast item for a condition matching a keyword in a field of the feed item
type Filter struct {
	Condition string // IN or NOT
	Keyword   string // search keyword
	Field     string // currently only title
}

func getFilter(condition, field, keyword string) (Filter, error) {
	var f Filter
	if condition == "" || field == "" || keyword == "" {
		return f, errors.New("condition, field and keyword of filter shall not be empty")
	}
	f.Condition = condition
	f.Field = field
	f.Keyword = keyword
	return f, nil
}

func addFilterToPostcast(f Filter, name string, p []Podcast) ([]Podcast, error) {
	for i := 0; i < len(p); i++ {
		if p[i].Name == name {
			p[i].Filter = f
			return p, nil
		}
	}
	return p, errors.New("addFilterToPostcast: no podcast with that name found")
}
