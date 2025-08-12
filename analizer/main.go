package main

import (
	"fmt"
	"strings"
	"bufio"
	"io"
	"os"
	"path/filepath"

	"github.com/Tinch334/file_analizer/internal/utils"
	"github.com/Tinch334/file_analizer/internal/config"
)

type FileMap map[string]FileInfo

type FileInfo struct {
	files int32
	lines int64
	code int64
	comments int64
	blanks int64
}

type FileResponse struct {
	languageName string
	lines int64
	code int64
	comments int64
	blanks int64
}

func newFileResponse(languageName string) FileResponse {
	fi := FileResponse{languageName: languageName, lines: 0, code: 0, comments: 0, blanks: 0}

	return fi
}


//Generic error checker.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	var err error

	fileLst, err := utils.GetFiles(".")
	check(err)
	languageInfo, err := config.GetConfig()
	check(err)

	for _, elem := range fileLst {
		//Check that file has an extension.
		if strings.Contains(elem.Name, ".") {
			//Remove dot from extension and get info associated with it.
			extension := filepath.Ext(elem.Name)[1:]
			info, prs := (*languageInfo)[extension]

			if prs {
				getFileData(elem.Path, info)
			} else {
				fmt.Println("Extension not found")
			}
		}
	}
}


func getFileData(path string, lInfo config.LanguageInfo) {
	fInfo := newFileResponse(lInfo.LanguageName)

	f, err := os.Open(path)
	check(err)
	defer closeFile(f)

	reader := bufio.NewReader(f)
	inMultilineComment := false

	for {
		line, err := reader.ReadString('\n')
		//Check for EOF, otherwise check for errors.
		if err == io.EOF {
			break
		}
		check(err)

		//Check for start and end of multiline comments
		multlineStart:= strings.LastIndex(line, lInfo.MultilineComment[0])
		if multlineStart != (-1) {
			inMultilineComment = true
		}

		multilineEnd := strings.LastIndex(line, lInfo.MultilineComment[1])
		//Check to see if there is a multiline comment starting and ending on the same line, if so don't stop counting.
		if multilineEnd != (-1) && multilineEnd > multlineStart {
			inMultilineComment = false
		}

		foundComment := false

		if inMultilineComment {
			fInfo.comments += 1
		} else {
			//Check all single line comment possibilities, if we are not in a multiline comment to avoid recounting.
			for _, lc := range lInfo.LineComment {
				if strings.Contains(line, lc) {
					fInfo.comments += 1
					foundComment = true
				}
			}
		}

		//ADD: Check for tabs
		if line != "" {
			if !inMultilineComment || foundComment {
				//If line isn't empty and we aren't in a comment then we are in a code line.
				fInfo.code += 1
			}
		} else {
			fInfo.blanks += 1
		}

		lInfo.lines += 1
	}

	fmt.Println("\n")
}

func closeFile(f *os.File) {
    err := f.Close()
    check(err)
}