//////////////////////////////////////////
// archivex_test.go
// Jhonathan Paulo Banczek - 2014
// jpbanczek@gmail.com - jhoonb.com
//////////////////////////////////////////

package archivex

import (
	"os"
	"testing"
)

func Test_archivex(t *testing.T) {

	// interface
	arcvx := []Archivex{&ZipFile{}, &TarFile{}}

	for _, arc := range arcvx {
		// create file
		if err := arc.Create("filetest"); err != nil {
			t.Fatalf("Error creating 'filetest': %v", err)
		}
		// create 50000 files
		dir, _ := os.Getwd()
		// absolute path
		if err := arc.AddAll(dir+"/testfolder/", true); err != nil {
			t.Fatalf("Error doing AddAll with '/testfolder/' and includeCurrentFolder = true: %v", err)
		}
		/*
			if err := arc.AddAll(dir+"/testfolder/", false); err != nil {
				t.Fatalf("Error doing AllAll with '/testfolder/' and includeCurrentFolder = false: %v", err)
			}
			// relative path
			if err := arc.AddAll("testfolder/", true); err != nil {
				t.Fatalf("Error doing AddAll with 'testfolder/' and includeCurrentFolder = true: %v", err)
			}
			if err := arc.AddAll("testfolder/", false); err != nil {
				t.Fatalf("Error doing AddAll with 'testfolder/' and includeCurrentFolder = false: %v", err)
			}
		*/
		arc.Close()
	}
}
