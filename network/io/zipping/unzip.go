package zipping

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(zipFile string) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	var errs []error

	for _, f := range r.File {
		path := f.Name
		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(path, f.Mode()); err != nil {
				errs = append(errs, err)
			}
			continue
		}

		dir, _ := os.Getwd()
		if err = os.MkdirAll(filepath.Dir(filepath.Join(dir, path)), os.ModePerm); err != nil {
			errs = append(errs, err)
			continue
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			errs = append(errs, err)
			continue
		}

		inFile, err := f.Open()
		if err != nil {
			outFile.Close()
			errs = append(errs, err)
			continue
		}

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			inFile.Close()
			outFile.Close()
			errs = append(errs, err)
			continue
		}

		inFile.Close()
		outFile.Close()
	}

	if len(errs) > 0 {
		var errMsgs []string
		for _, e := range errs {
			errMsgs = append(errMsgs, e.Error())
		}
		return fmt.Errorf("unzip failed with error_const: %s", strings.Join(errMsgs, "; "))
	}

	return nil
}
