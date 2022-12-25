// Copyright (C) 2022 Tatsuto YAMAMOTO
/*
Memoは、ファイルシステム上のディレクトリと、配下のファイル/ディレクトリをメモリ上に読み込むパッケージです。
*/
package memo

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// File is file on fs.
type File struct {
	// file name
	Path string

	// file contents
	Contents []byte

	// file size
	Size int64
}

func ReadDir(dirPath string, ignorePath []string) ([]File, error) {
	files := []File{}
	if err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, _ := d.Info()

		for _, v := range ignorePath {
			if d.IsDir() && filepath.Base(path) == v {
				return filepath.SkipDir
			}
		}

		if info.IsDir() {
			return nil
		}

		fileReader, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileReader.Close()

		fileContents, err := io.ReadAll(fileReader)
		if err != nil {
			return err
		}

		fileData := File{
			Path:     filepath.ToSlash(path),
			Size:     info.Size(),
			Contents: fileContents,
		}

		// fmt.Println(fileData)
		files = append(files, fileData)

		return nil
	}); err != nil {
		return []File{}, err
	} else {
		return files, nil
	}
}
