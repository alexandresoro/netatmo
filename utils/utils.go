package utils

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func CreateFile(filePath string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(filePath), 0770); err != nil {
		return nil, err
	}
	return os.Create(filePath)
}

func WriteToDestinationPath(filePath string, data []byte) error {
	file, err := CreateFile(ResolvePathWithTilde(filePath))
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ResolvePathWithTilde(path string) string {

	usr, _ := user.Current()
	dir := usr.HomeDir

	var resolvedPath string
	if path == "~" {
		// In case of "~", which won't be caught by the "else if"
		resolvedPath = dir
	} else if strings.HasPrefix(path, "~/") {
		resolvedPath = filepath.Join(dir, path[2:])
	} else {
		resolvedPath = path
	}

	return resolvedPath
}
