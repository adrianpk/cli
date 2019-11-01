// Code generated for package generator by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/templates/migration.tmpl
// assets/templates/model.tmpl
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

var _assetsTemplatesModelTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x10\x02\xc4\xfe\x90\x8f\xba\x07\xe8\x21\x70\x7b\xf0\x21\x45\x50\xb7\x3d\x67\x25\xae\xe4\xad\x29\xae\xc2\x1f\x24\xae\xe1\x77\x2f\xf4\x93\xc4\xb1\x0b\x37\x87\xf6\x66\x2f\x67\x67\x86\xc3\xb1\x3b\xac\x36\xd8\x10\xb4\x62\xc8\x2a\xc5\x6d\x27\x3e\xc2\x4c\x65\xb9\xc1\x88\x25\x06\x2a\xc2\x83\xcd\x15\x40\xde\x70\x5c\xa7\x52\x57\xd2\x16\x01\xa3\x78\x2e\x1a\xd1\x29\xb1\xc9\x95\xca\x0e\x4f\x2d\x97\x45\xf7\x90\xab\xac\x1d\x96\x2c\x8e\xe3\x96\x37\x5e\x1e\xe9\x27\xd9\xa2\xc4\x6a\x43\xce\x14\x83\x6a\xae\xe6\x4a\xc5\x6d\x47\xbd\x6c\x51\xc0\x6e\xa7\x97\xae\xb6\x7a\xc5\xae\x49\x16\xfd\x1d\x86\x0a\xed\x02\x03\xed\xf7\x93\xcf\xec\x2c\x26\x44\x9f\xaa\x08\x3b\x95\x65\xad\x5e\x1a\x72\x91\x6b\xae\x30\xb2\x38\x95\x65\xbb\xdd\xff\xe0\xd1\x35\x04\x17\x1b\xda\x5e\xc1\x45\xe7\xa5\x83\xeb\x0f\xa0\xef\xbc\x74\x1f\xa9\x0e\xfb\xfd\x04\xe3\x1a\x9c\xc4\x11\xa1\x97\xe1\x53\x5b\x92\x31\x64\xa6\xf3\x71\xfc\x19\xdb\x5e\xf3\xe5\x6b\xb2\xf6\xeb\xb6\xeb\x47\xf7\xa6\xbc\xce\x9f\xe7\x0b\xb1\x13\x34\x87\x1f\x41\xdc\xeb\xc9\x9b\x7b\x2c\xb0\xa5\xe9\x1a\x57\xd2\x72\xa4\xb6\x8b\xdb\xfc\x7e\xf2\x43\xce\xbc\x78\x1b\x3f\x03\x00\xb4\xfa\x26\x19\x8e\x2a\xdb\xf7\x41\x16\x05\xac\x28\x2e\x3c\x61\xa4\xef\x68\x13\x05\xa8\xc5\x8f\xb9\x69\x55\x27\x57\xc1\xec\x38\xbd\x03\x55\xf8\xef\x5c\xb4\xf3\x63\xee\xd9\x1c\xc8\x7b\xf1\xb0\x53\x00\x45\x01\x5f\xa8\xb3\x58\x11\x5c\x9e\x91\x18\x72\xd0\xab\xe8\xd9\x35\x97\x83\x39\x74\x12\xd7\xe4\x21\x39\x7e\x48\x04\x3c\x3d\x19\xf9\x91\x74\x59\x43\xbf\x02\x46\x28\x0c\x0f\x42\x4f\x1c\x22\x88\x87\xad\x24\xe8\x3c\xd5\xf4\xca\x21\x8e\xf4\xb4\x16\xe1\x91\xad\x85\x92\x20\x05\x32\x83\xd2\xca\xa6\x06\x1a\x72\xe4\x87\x3a\x4c\xc8\x1b\xe8\x2c\xb2\xeb\x8b\xc3\xae\x19\xb7\xd0\x06\x81\x47\xf1\x1b\x28\x53\x04\x9e\xb8\x5a\xdc\x10\x58\x91\x0d\x04\x6e\xd9\xa2\x07\xb4\x16\x82\x4d\xcd\x18\x33\x61\xb5\x06\x76\x21\xa2\xab\x68\x24\x97\x1a\xe2\x9a\x03\x78\x0a\x92\x7c\x45\x7a\x28\x39\x6b\xd2\xd7\x27\x55\x3f\x48\xe9\xae\x7e\xea\x4b\x99\xf7\xdc\x5d\xfd\x94\x9f\x76\xfe\x14\xfc\xce\xd0\xcf\x52\x1d\xfd\x60\xf4\xf1\x83\x9f\x77\x31\x3f\xcf\x3d\x14\xf5\x84\x72\xae\x32\x4f\x31\x79\x07\x8e\xad\xda\x3f\x77\xf8\x5b\x67\xfe\x59\x87\x0f\xb9\x0f\x3a\xfc\x3e\xf3\x6f\x97\x7f\x63\xfe\x16\x63\xb5\x86\x4a\x9c\xe1\x3e\xc3\xbf\x6a\x7e\xe0\x9e\xc5\xea\x4f\xb8\x52\xc4\xf6\x37\xf2\x53\x31\x6e\x07\xfd\x61\x7b\x21\xe3\xff\xc8\xe4\xdb\xab\xbd\xfa\x15\x00\x00\xff\xff\xad\xd9\xd3\xef\x07\x06\x00\x00")

func assetsTemplatesModelTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesModelTmpl,
		"assets/templates/model.tmpl",
	)
}

func assetsTemplatesModelTmpl() (*asset, error) {
	bytes, err := assetsTemplatesModelTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/model.tmpl", size: 1543, mode: os.FileMode(436), modTime: time.Unix(1572621878, 0)}
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
	"assets/templates/model.tmpl":     assetsTemplatesModelTmpl,
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
			"model.tmpl":     &bintree{assetsTemplatesModelTmpl, map[string]*bintree{}},
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
