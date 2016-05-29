package main

import "os"

type Uploader interface {
	Upload(*os.File) (string, error)
}
