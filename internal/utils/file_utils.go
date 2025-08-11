package utils

import (
	"io/fs"
	"path/filepath"
)


type fileInfo struct {
	Name string
	Path string
}

func newFileInfo(name string, path string) *fileInfo {
	fi := fileInfo{Name: name, Path: path}

	return &fi
}


//Returns a slice with all the files in the given directory.
func GetFiles(dir string) ([]*fileInfo, error) {
	var fileList []*fileInfo

	//Walk given directory, store all files in file list.
	err := filepath.WalkDir(dir, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !file.IsDir() {
			fileList = append(fileList, newFileInfo(filepath.Base(file.Name()), path))
		}

		return nil;
	})

	//If an error occurred return it and no list.
	if err != nil {
		return nil, err
	}

	return fileList, nil
}