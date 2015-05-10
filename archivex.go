//////////////////////////////////////////
// archivex.go
// Jhonathan Paulo Banczek - 2014
// jpbanczek@gmail.com - jhoonb.com
//////////////////////////////////////////

package archivex

import (
	"archive/tar"
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// interface
type Archivex interface {
	Create(name string) error
	Add(name string, file []byte) error
	AddFile(name string) error
	AddAll(dir string, includeCurrentFolder bool) error
	Close() error
}

// ZipFile implement *zip.Writer
type ZipFile struct {
	Writer *zip.Writer
	Name   string
}

// TarFile implement *tar.Writer
type TarFile struct {
	Writer *tar.Writer
	Name   string
}

func isDir(path string) bool {
	src, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	return src.IsDir()
}

// Create new file zip
func (z *ZipFile) Create(name string) error {
	// check extension .zip
	if strings.HasSuffix(name, ".zip") != true {
		if strings.HasSuffix(name, ".tar.gz") == true {
			name = strings.Replace(name, ".tar.gz", ".zip", -1)
		} else {
			name = name + ".zip"
		}
	}
	z.Name = name
	file, err := os.Create(z.Name)
	if err != nil {
		return err
	}
	z.Writer = zip.NewWriter(file)
	return nil
}

// Add add byte in archive zip
func (z *ZipFile) Add(name string, file []byte) error {

	iow, err := z.Writer.Create(name)
	if err != nil {
		return err
	}
	_, err = iow.Write(file)
	return err
}

// AddFile add file from dir in archive
func (z *ZipFile) AddFile(name string) error {
	bytearq, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	filep, err := z.Writer.Create(name)
	if err != nil {
		return err
	}
	_, err = filep.Write(bytearq)
	if err != nil {
		return err
	}
	return nil
}

// AddAll add all files from dir in archive
func (z *ZipFile) AddAll(dir string, includeCurrentFolder bool) error {
	dir = path.Clean(dir)
	return z.addAll(dir, dir, includeCurrentFolder)
}

func (z *ZipFile) addAll(dir string, rootDir string, includeCurrentFolder bool) error {
	// capture all name files in dir
	listFile, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var names []string
	var bdatas [][]byte

	for _, arq := range listFile {
		full := path.Join(dir, arq.Name())
		if isDir(full) {
			z.addAll(full, rootDir, includeCurrentFolder)
		} else {
			bytearq, err := ioutil.ReadFile(full)
			if err != nil {
				return err
			}
			names = append(names, arq.Name())
			bdatas = append(bdatas, bytearq)
		}
	}

	subDir := getSubDir(dir, rootDir, includeCurrentFolder)
	for i, file := range bdatas {
		filep, err := z.Writer.Create(path.Join(subDir, names[i]))
		if err != nil {
			return err
		}
		_, err = filep.Write(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *ZipFile) Close() error {
	err := z.Writer.Close()
	return err
}

// Create new Tar file
func (t *TarFile) Create(name string) error {
	// check extension .zip
	if strings.HasSuffix(name, ".tar.gz") != true {
		if strings.HasSuffix(name, ".zip") == true {
			name = strings.Replace(name, ".zip", ".tar.gz", -1)
		} else {
			name = name + ".tar.gz"
		}
	}
	t.Name = name
	file, err := os.Create(t.Name)
	if err != nil {
		return err
	}
	t.Writer = tar.NewWriter(file)
	return nil
}

// Add add byte in archive tar
func (t *TarFile) Add(name string, file []byte) error {

	hdr := &tar.Header{Name: name, Size: int64(len(file))}
	if err := t.Writer.WriteHeader(hdr); err != nil {
		return err
	}
	_, err := t.Writer.Write(file)
	return err
}

// AddFile add file from dir in archive tar
func (t *TarFile) AddFile(name string) error {
	bytearq, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	hdr := &tar.Header{Name: name, Size: int64(len(bytearq))}
	err = t.Writer.WriteHeader(hdr)
	if err != nil {
		return err
	}
	_, err = t.Writer.Write(bytearq)
	if err != nil {
		return err
	}
	return nil

}

// AddAll add all files from dir in archive
func (t *TarFile) AddAll(dir string, includeCurrentFolder bool) error {
	dir = path.Clean(dir)
	return t.addAll(dir, dir, includeCurrentFolder)
}

func (t *TarFile) addAll(dir string, rootDir string, includeCurrentFolder bool) error {

	// capture all name files in dir
	listFile, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var names []string
	var bdatas [][]byte

	for _, arq := range listFile {
		if isDir(path.Join(dir, arq.Name())) {
			t.addAll(path.Join(dir, arq.Name()), rootDir, includeCurrentFolder)
		} else {
			bytearq, err := ioutil.ReadFile(path.Join(dir, arq.Name()))
			if err != nil {
				return err
			}
			names = append(names, arq.Name())
			bdatas = append(bdatas, bytearq)
		}
	}

	subDir := getSubDir(dir, rootDir, includeCurrentFolder)
	for i, file := range bdatas {
		hdr := &tar.Header{Name: path.Join(subDir, names[i]), Size: int64(len(file))}
		log.Print("Writing header for: ", hdr.Name)
		if err := t.Writer.WriteHeader(hdr); err != nil {
			return err
		}
		_, err = t.Writer.Write(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close the file Tar
func (t *TarFile) Close() error {
	err := t.Writer.Close()
	return err
}

func getSubDir(dir string, rootDir string, includeCurrentFolder bool) (subDir string) {

	// log.Printf("dir: %v", dir)
	// log.Printf("rootDir: %v", rootDir)

	subDir = strings.Replace(dir, rootDir, "", 1)

	if includeCurrentFolder {
		parts := strings.Split(rootDir, string(os.PathSeparator))
		subDir = path.Join(parts[len(parts)-1], subDir)
	}

	// log.Printf("getSubDir: %v", subDir)

	return
}
