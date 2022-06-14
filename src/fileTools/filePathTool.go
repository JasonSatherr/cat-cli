package fileTools

import (
	"os"
	"path/filepath"
)

//geExecutionPath is a function that will return the path the the program is executing from
func GetExecutionPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
