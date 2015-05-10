//////////////////////////////////////////
// archivex_test.go
// Jhonathan Paulo Banczek - 2014
// jpbanczek@gmail.com - jhoonb.com
//////////////////////////////////////////

package archivex

import (
	"fmt"
	// "log"
	"os"
	"path"
	"reflect"
	"testing"
)

type archTest struct {
	addPath string
	include bool
	name    string
}

type archTypeTest struct {
	tests []archTest
	arch  Archivex
}

func Test_archivex(t *testing.T) {

	dir, _ := os.Getwd()

	// All the different tests we want to run with different combinations of input paths and the includeCurrentFolder flag
	tests := []archTest{
		// absolute path
		archTest{dir + "/testfolder/", true, "absTrailInclude"},
		archTest{dir + "/testfolder/", false, "absTrailExclude"},
		// relative path
		archTest{"testfolder/", true, "relTrailInclude"},
		archTest{"testfolder/", false, "relTrailExclude"},
		// without trailing slashes
		archTest{dir + "/testfolder", true, "absInclude"},
		archTest{dir + "/testfolder", false, "absExclude"},
		archTest{"testfolder", true, "relInclude"},
		archTest{"testfolder", false, "relExclude"},
	}

	// We want to execute the batch of tests on both Zip and Tar
	typeTests := []archTypeTest{
		archTypeTest{tests, &ZipFile{}},
		archTypeTest{tests, &TarFile{}},
	}

	// Run all tests
	for _, typeTest := range typeTests {

		currentType := reflect.TypeOf(typeTest.arch)
		t.Logf("Running tests for archive type: %s", currentType.Elem())

		for i, test := range typeTest.tests {

			t.Logf("Running %s...", test.name)

			// Create the archive
			filename := fmt.Sprintf("%d_%s_test", i+1, test.name)
			arch := reflect.ValueOf(typeTest.arch).Interface().(Archivex)
			if err := arch.Create(path.Join("testresults", filename)); err != nil {
				t.Fatalf("Error creating '%s': %v", filename, err)
			}

			// Add the files to the archive
			if err := arch.AddAll(test.addPath, test.include); err != nil {
				t.Fatalf("Error doing AddAll with '%s' and includeCurrentFolder = %v: %v", test.addPath, test.include, err)
			}

			// Close the archive
			arch.Close()

		}
	}
}
