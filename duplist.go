package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func search(path string) {
	if fi, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "does not exist: %s\n", path)
		return
	} else if f, err := os.Open(path); err != nil {
		fmt.Fprintf(os.Stderr, "can not open path: %s\n", path)
		return
	} else {
		if fi.IsDir() {
			if childs, err := f.Readdir(0); err != nil {
				fmt.Fprintf(os.Stderr, "can not read dir: %s\n", fi.Name())
				return
			} else {
				for _, child := range childs {
					search(filepath.Join(path, child.Name()))
				}
			}
		} else {
			h := md5.New()
			if _, err := io.Copy(h, f); err != nil {
				fmt.Fprintf(os.Stderr, "create hash fail: %s\n", fi.Name())
				return
			} else {
				fmt.Fprintf(os.Stdout, "%s\t%s\n", filepath.Join(path, fi.Name()), hex.EncodeToString(h.Sum(nil)))
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Stderr, "need args (target dir)\n")
		return
	}
	for _, path := range os.Args[1:] {
		search(path)
	}
}
