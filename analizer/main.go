package main

import (
	"fmt"
	"path/filepath"

	"github.com/Tinch334/file_analizer/internal/utils"
	"github.com/Tinch334/file_analizer/internal/config"
)

//Generic error checker.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	var err error

	lst, err := utils.GetFiles(".")
	check(err)
	languageInfo, err := config.GetConfig()
	check(err)

	for _, elem := range lst {
		extension := filepath.Ext(elem.Name)[1:]
		info, prs := languageInfo.GetInfoFromExtension(extension)

		if prs {
			fmt.Println(extension)
			fmt.Println(info)
		} else {
			fmt.Println("Extension not found")
		}
	}
}