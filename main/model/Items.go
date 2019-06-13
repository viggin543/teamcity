package model

import "strings"

type Item struct {
	Title string  `json:"title"`
	Subtitle string  `json:"subtitle"`
	Arg string  `json:"arg"`
	Id string  `json:"id"`
}
type Items struct {
	Items []*Item `json:"items"`
}

func (items *Items) FilterByString(token string) []*Item {
	return items.Filter(func(item *Item) bool {
		title := strings.ToLower(item.Title)
		searchToken := strings.ToLower(token)
		return strings.Contains(title, searchToken)
	})
}

func (items *Items ) Filter( pred func(item *Item) bool) []*Item {
	needed := items.Count(pred)
	ret := make([]*Item, needed)
	idx := 0
	for _, item := range items.Items {
		if pred(item) {
			ret[idx] = item
			idx++
		}
	}
	return ret
}

func (items *Items ) Count( pred func(item *Item) bool) int {
	var count = 0
	for _,item := range items.Items {
		if pred(item){ count ++ }
	}
	return count
}



