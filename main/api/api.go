package api

import (
	"encoding/xml"
	"github.com/teamcity/main/model"
	"io/ioutil"
	"net/http"
	"os"
)



func FetchProjects() *model.Projects {
	var projects model.Projects
	unmarshal("build.innovid.com/httpAuth/app/rest/projects",&projects)
	return &projects
}

func FetchBuildTypes() *model.BuildTypes {
	var buildTypes model.BuildTypes
	unmarshal("build.innovid.com/httpAuth/app/rest/buildTypes",&buildTypes)
	return &buildTypes
}

func unmarshal(route string, model interface{}) {
	usr := os.Getenv("user")
	pass := os.Getenv("pass")

	_ = xml.Unmarshal(
		fetch("https://"+usr+":"+pass+"@"+route),
		model)

}


func fetch(route string) []byte {
	ret := make(chan []byte) // its not really async....
	go func() {
		cacheFileName, file, err := readCacheFile(route)
		if err != nil || notOldEnough(file) {
			res, _ := http.Get(route)
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			_ = ioutil.WriteFile(cacheFileName, body, 777)
			ret <- body
		} else {
			bytes, _ := ioutil.ReadFile(cacheFileName)
			ret <- bytes
		}
	}()
	return <-ret
}


