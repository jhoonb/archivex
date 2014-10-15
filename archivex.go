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
	"os"
	"strings"
)

// interface
type Archivex interface {
	Create(name string) error
	Add(name string, file []byte) error
	AddFile(name string) error
	AddAll(dir string) error
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
func (z *ZipFile) AddAll(dir string) error {

	// capture all name files in dir
	listFile, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var names []string
	var bdatas [][]byte

	for _, arq := range listFile {
		bytearq, err := ioutil.ReadFile(dir + arq.Name())
		if err != nil {
			return err
		}
		names = append(names, arq.Name())
		bdatas = append(bdatas, bytearq)
	}

	for i, file := range bdatas {
		filep, err := z.Writer.Create(names[i])
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
func (t *TarFile) AddAll(dir string) error {

	// capture all name files in dir
	listFile, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var names []string
	var bdatas [][]byte

	for _, arq := range listFile {
		bytearq, err := ioutil.ReadFile(dir + arq.Name())
		if err != nil {
			return err
		}
		names = append(names, arq.Name())
		bdatas = append(bdatas, bytearq)
	}

	for i, file := range bdatas {
		hdr := &tar.Header{Name: names[i], Size: int64(len(file))}
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
