package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mmcquillan/lawsg/util"
)

func ReadGroups(cacheDir string) (groups []string, exists bool) {
	gc := util.MakePath(cacheDir, "g")
	if !util.FileExists(gc) {
		return groups, false
	}
	fileStat, err := os.Stat(gc)
	if err != nil {
		return groups, false
	}
	if fileStat.ModTime().Unix()*1000 < util.ParseDate("30m ago") {
		os.Remove(gc)
		return groups, false
	}
	file, err := ioutil.ReadFile(gc)
	if err != nil {
		return groups, false
	}
	err = json.Unmarshal(file, &groups)
	if err != nil {
		fmt.Println("ERROR: Group File Unmarshal - ", err)
		os.Exit(1)
	}
	return groups, true
}

func WriteGroups(groups []string, cacheDir string) {
	if !util.DirExists(cacheDir) {
		os.MkdirAll(cacheDir, 0700)
	}
	gc := util.MakePath(cacheDir, "g")
	if !util.FileExists(gc) {
		jsonGroups, _ := json.Marshal(groups)
		err := ioutil.WriteFile(gc, jsonGroups, 0644)
		if err != nil {
			fmt.Println("ERROR: Group Cache Write - ", err)
		}
	}
}
