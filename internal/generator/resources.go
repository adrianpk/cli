// Code generated for package generator by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/templates/migration.tmpl
package generator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsTemplatesMigrationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x41\x6b\xf2\x40\x10\x86\xcf\x3b\xbf\x62\xbe\xe0\x21\xf9\xd0\x08\x3d\x0a\x1e\x8a\x2d\xa5\xd0\x83\x25\x1e\x0b\x75\x4d\xc7\xb0\x38\xd9\xd8\xc9\x08\x29\x61\xff\x7b\xd9\xd4\x83\x6d\x11\x14\x7a\xcc\x3b\x99\x67\x5f\xe6\xd9\xdb\x72\x67\x2b\xc2\xda\x55\x62\xd5\x35\x1e\xc0\xd5\xfb\x46\x14\x13\x6e\xaa\x04\x60\x3a\xc5\x85\x90\x55\xea\xfb\x09\xe6\x8f\x7e\xcb\xf9\x92\x0f\x62\x79\x69\xdb\xd2\xf2\xc2\xb6\x84\x93\x10\x56\x76\xc3\xa7\x94\xed\xc1\x97\x98\xd6\xf8\xbf\x76\x55\x76\x39\x21\xcd\x90\x44\x1a\xc1\x1e\x8c\x76\x38\x9b\x63\x9d\x3f\x90\xae\xba\x34\x03\x30\xad\xc6\x64\x3d\x70\x8a\xe7\xa7\xfc\x0b\x5b\x68\x5c\x5f\x03\x98\x38\x48\x5e\x7c\x12\xbf\x01\xcc\xeb\x38\xc2\x70\x8e\xda\xe5\xf7\x1d\x95\x69\xab\x19\x18\xb7\x1d\xd2\x7f\x73\xf4\x8e\xe3\x3b\x46\x48\x0f\xe2\x63\x0a\x26\xfc\xc6\xf4\x3d\x8a\xf5\x15\xe1\x68\x47\x1f\x37\x63\x1c\xb5\xef\x5c\xa8\x38\x5f\xc5\x36\x43\x91\x5b\x56\x92\x42\x31\x04\x40\x6c\x15\x8f\x25\x4f\xfe\xfc\xcb\x86\xb1\x11\xf9\x37\x1c\x10\xc7\x91\x77\x0c\x61\xb0\x75\x27\xcd\xfe\x12\x57\xd2\x30\x6f\x6c\xb9\xfb\xa1\xea\xc2\xf5\xeb\x44\x45\xe8\x19\x4d\x67\x4e\x32\xbb\xda\xda\xb7\x43\x7c\x06\x00\x00\xff\xff\xa2\x64\xfa\x07\xd7\x02\x00\x00")

func assetsTemplatesMigrationTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesMigrationTmpl,
		"assets/templates/migration.tmpl",
	)
}

func assetsTemplatesMigrationTmpl() (*asset, error) {
	bytes, err := assetsTemplatesMigrationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/migration.tmpl", size: 727, mode: os.FileMode(436), modTime: time.Unix(1572532722, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/templates/migration.tmpl": assetsTemplatesMigrationTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"migration.tmpl": &bintree{assetsTemplatesMigrationTmpl, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
