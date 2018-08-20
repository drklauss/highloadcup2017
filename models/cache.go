package models

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	UCache   *Users
	VCache   *Visits
	LocCache *Locations
	UvlCache *UserVisitLinks
	LvlCache *LocationVisitLinks
)

func Init() {
	//err := unzipData()
	//if err != nil {
	//	panic("cannot read data")
	//}
	// Таблицы связей
	t := time.Now()
	UvlCache = new(UserVisitLinks).Init()
	LvlCache = new(LocationVisitLinks).Init()
	// Основные таблицы
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		LocCache = new(Locations).Init()
		wg.Done()
	}()
	go func() {
		UCache = new(Users).Init()
		wg.Done()
	}()
	go func() {
		VCache = new(Visits).Init()
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("All data: %+v", time.Since(t))
}

func unzipData() error {
	r, err := zip.OpenReader("data.zip")
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll("data", 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join("data", f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}
	for k, f := range r.File {
		extractAndWriteFile(f)
	}

	for _, f := range r.File {
		extractAndWriteFile(f)
	}

	return nil
}
