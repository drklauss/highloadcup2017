package main

import (
	"archive/zip"
	"github.com/buaazp/fasthttprouter"
	"github.com/drklauss/highloadcup2017/handlers"
	"github.com/drklauss/highloadcup2017/models"
	"github.com/valyala/fasthttp"
	"io"
	"os"
	"path/filepath"
)

func main() {
	Unzip()
	models.Init()
	router := fasthttprouter.New()

	router.GET("/users/:id", handlers.GetUser)
	router.GET("/users/:id/visits", handlers.GetUserVisits)
	router.GET("/locations/:id", handlers.GetUser)
	router.GET("/locations/:id/avg", handlers.GetUser)
	router.GET("/visits/:id", handlers.GetUser)
	router.POST("/visits/new", handlers.NewVisit)

	panic(fasthttp.ListenAndServe(":8080", router.Handler))
}

func Unzip() error {
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

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
