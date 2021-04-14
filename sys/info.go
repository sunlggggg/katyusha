spackage sys

import (
	"os"
	"path/filepath"
)

func BinSize(bin string) int64 {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fi , _ :=os.Stat(dir + "/" + bin)
	return fi.Size()
}
