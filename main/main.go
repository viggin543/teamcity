package main

import (
	"encoding/json"
	"fmt"
	"github.com/teamcity/main/api"
	"github.com/teamcity/main/model"
	"os"
)



func main() {
	searchToken := commandLineArg()
	projects := api.FetchProjects().ToItems()
	buildTypes := api.FetchBuildTypes().ToItems()

	items := &model.Items{Items: append(projects, buildTypes...)}
	itemsJson, _ := json.Marshal(model.Items{Items: items.FilterByString(searchToken)})
	fmt.Println(string(itemsJson))
}

func commandLineArg() string {
	if len(os.Args) == 1 {
		return "ddv"
	} else {
		return  os.Args[1]
	}
}



