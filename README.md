archivex
========

archivex is a golang package that implements facilities to create files zip and tar
using standard library.

Installation
-------------

::

   $ go get github.com/jhoonb/archivex


Example 
-------------

::


package main

import (
	"github.com/jhoonb/archivex"
)

// Example using only func zip
func zip() {
	zip := new(archivex.ZipFile)
	zip.Create("filezip")
	zip.Add("testadd.txt", []byte("test 1"))
	zip.AddFile("<input_path_file_here>")
	zip.AddAll("<input_dir_here")
	zip.Close()
}

// Example using only func tar
func tar() {
	tar := new(archivex.TarFile)
	tar.Create("filetar")
	tar.Add("testadd.txt", []byte("test 1"))
	tar.AddFile("<input_path_file_here>")
	tar.AddAll("<input_dir_here")
	tar.Close()
}

// Example using interface
func usingInterface() {

	archx := []archivex.Archivex{&archivex.TarFile{}, &archivex.ZipFile{}}

	for _, arch := range archx {
		arch.Create("fileinterface")
		arch.Add("testadd.txt", []byte("file 1 :) "))
		arch.AddFile("<input_path_file_here>")
		arch.AddAll("<input_dir_here")
		arch.Close()
	}
}

func main() {

	zip()
	tar()
	usingInterface()
}


:)
