package main

import (
	"encoding/json"
	"fmt"
	"github.com/teamcity/main/api"
	"github.com/teamcity/main/model"
	"os"
)



func main() {

	var projects = make(chan []*model.Item)
	var buildTypes = make(chan []*model.Item)
	go func() {
		projects <- api.FetchProjects().ToItems()
	}()
	go func() {
		buildTypes <- api.FetchBuildTypes().ToItems()
	}()

	items := filterItemsByCommandLineArg(projects, buildTypes)
	printItemsAsJson(items)
}

func printItemsAsJson(items model.Items) {
	itemsJson, _ := json.Marshal(items)
	fmt.Println(string(itemsJson))
}

func filterItemsByCommandLineArg(projects chan []*model.Item, buildTypes chan []*model.Item) model.Items {
	searchToken := commandLineArg()
	unfilteredItems := append(<-projects, <-buildTypes...)
	items := model.Items{Items: model.FilterItems(unfilteredItems, searchToken)}
	return items
}

func commandLineArg() string {
	if len(os.Args) == 1 {
		return "ddv"
	} else {
		return  os.Args[1]
	}
}



