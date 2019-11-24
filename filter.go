package main

import (
	"errors"
	"log"
	"podcoff/cmd"
	"strings"
)

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

// doesFilterMatch returns true if a given filter matches the given item in the podcast
func doesFilterMatch(item PodcastItem, f Filter) bool {
	if f.Field == "" || f.Condition == "" || f.Keyword == "" {
		// no filter or no valid filter means we match all titles
		return true
	}
	var field string
	if f.Field == "title" {
		field = item.Title
	} else {
		log.Fatalln("doesFilterMatch: Given filter field is not equal to title. It's:", f.Field)
	}

	if f.Keyword == "" {
		log.Fatalln("doesFilterMatch: Given filter keyword is empty")
	}
	contains := strings.Contains(strings.ToLower(field), strings.ToLower(f.Keyword))
	if cmd.Debug {
		log.Printf("doesFilterMatch: condition=%s, field=%s, keyword=%s, contains=%v\n", f.Condition, field, f.Keyword, contains)
	}
	if f.Condition == "IN" {
		return contains
	} else if f.Condition == "NOT" {
		return !contains
	} else {
		log.Fatalln("doesFilterMatch: Given filter condition is not IN or NOT. It's:", f.Condition)
	}
	// if we really reach this return accept the item and report filter does match
	return true
}
