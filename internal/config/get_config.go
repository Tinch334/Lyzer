package config

import (
	"os"
	"encoding/json"
)

type LanguageMap map[string]LanguageInfo

//Stores the data extracted from the JSON file.
type LanguageInfo struct {
	LanguageName string `json:"name"`
	LineComment []string `json:"line_comment"`
	MultiLineComment []string `json:"multi_line_comment"`
	Extensions []string `json:"extensions"`
}

type Root struct {
	Languages LanguageMap `json:"languages"`
}

/*
For more efficient access we create an auxiliary structure, where each file extension is a key in a dictionary, associated with the language
name. That way we can directly get the corresponding file info.
*/
type Info struct {
	languages *LanguageMap
}

func makeInfo(languages *LanguageMap) *Info {
	i := Info{languages: languages}

	return &i
}

//The following two functions can be used to get the file info structure associated with a language. The second argument indicates whether the
//key was found.

func GetConfig() (*LanguageMap, error ){
	var err error
	var data Root

	//Load and un-marshal data from JSON.
	jsonData, err := os.ReadFile("languages.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}

	extensionMap := make(LanguageMap)

	//Associate each file extension with it's corresponding language.
	for k, d := range data.Languages {
		li := data.Languages[k]

		for _, elem := range d.Extensions {	
			extensionMap[elem] = li
		}
	}

	return &extensionMap, nil
}