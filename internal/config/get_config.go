package config

import (
	"os"
	"encoding/json"
)

//Stores the data extracted from the JSON file.
type LanguageInfo struct {
	LineComment []string `json:"line_comment"`
	MultiLineComment []string `json:"multi_line_comment"`
	Extensions []string `json:"extensions"`
}

type Root struct {
	Languages map[string]LanguageInfo `json:"languages"`
}

/*
For more efficient access we create an auxiliary structure, where each file extension is a key in a dictionary, associated with the language
name. That way we can directly get the corresponding file info.
*/
type Info struct {
	extensionMap *map[string]string
	languages *map[string]LanguageInfo
}

func makeInfo(extensionMap *map[string]string, languages *map[string]LanguageInfo) *Info {
	i := Info{extensionMap: extensionMap, languages: languages}

	return &i
}

//The following two functions can be used to get the file info structure associated with a language. The second argument indicates whether the
//key was found.
func (i* Info) GetInfoFromExtension(extension string) (LanguageInfo, bool) {
	value, prs := (*i.languages)[(*i.extensionMap)[extension]]

	return value, prs
}

func (i* Info) GetInfoFromName(name string) (LanguageInfo, bool) {
	value, prs := (*i.languages)[name]

	return value, prs
}

func GetConfig() (*Info, error ){
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

	extensionMap := make(map[string]string)

	//Associate each file extension with it's corresponding language.
	for k, d := range data.Languages {
		for _, elem := range d.Extensions {
			extensionMap[elem] = k
		}
	}

	return makeInfo(&extensionMap, &data.Languages), nil
}