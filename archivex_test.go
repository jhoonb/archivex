//////////////////////////////////////////
// archivex_test.go
// Jhonathan Paulo Banczek - 2014
// jpbanczek@gmail.com - jhoonb.com
//////////////////////////////////////////

package archivex

import (
	"testing"
	"os"
)

func Test_archivex(t *testing.T) {

	// interface
	arcvx := []Archivex{&ZipFile{}, &TarFile{}}

	for _, arc := range arcvx {
		// create file
		err := arc.Create("filetest")
		checkError(t, err)
		// create 50000 files
		dir, _ := os.Getwd()
		// absolute path
		err = arc.AddAll(dir+"/testfolder/", true)
		err = arc.AddAll(dir+"/testfolder/", false)
		// relative path
        	err = arc.AddAll("testfolder/", true)
        	err = arc.AddAll("testfolder/", false)
		checkError(t, err)
		arc.Close()
		checkError(t, err)
	}
}

// func for check errors
func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
