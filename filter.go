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
			err := isFilterAlreadyKnown(f, p[i])
			if err != nil {
				return p, err
			}
			p[i].Filters = append(p[i].Filters, f)
			return p, nil
		}
	}
	return p, errors.New("addPostcast: a podcast with that name or feed url is already in the database")
}

func isFilterAlreadyKnown(f Filter, p Podcast) error {
	for i := 0; i < len(p.Filters); i++ {
		if p.Filters[i].Condition == f.Condition && p.Filters[i].Field == f.Field && p.Filters[i].Keyword == f.Keyword {
			return errors.New("This filter is already known to that podcast")
		}
	}
	return nil
}
