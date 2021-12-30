// Code generated for package config by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.ini
// configs/migrations/01_initial.up.sql
package config

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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\x71\x6b\xc2\x30\x10\xc5\xff\xbf\x4f\x71\x5f\x60\x1a\x71\x03\x11\xfa\x47\xb4\x59\x2d\x8b\x4d\xe9\xb5\x6e\x4c\x24\x44\x0d\xed\xa0\x90\x92\xc4\x7d\xfe\x11\xc7\xc6\x60\x1c\xf7\xc7\xfb\xc1\xbd\xf7\xb8\xe3\xe8\xfa\x70\x02\xa9\x0a\x2d\xc5\x41\x48\xcc\x90\xc1\x73\x29\x85\x96\xaa\x28\xca\xaa\xc0\x0c\xa3\xbf\x59\xd8\xaa\x8a\xd4\x7f\x9c\x0e\x6b\xde\xee\x30\xc3\xe4\x34\x37\xb7\x38\xcc\x46\xd7\xc3\x9e\xbf\x69\x2a\xdf\x05\x66\xb8\x60\x77\xc5\x8b\x3f\x62\xc3\xb7\x2f\x5d\x4d\xdf\x00\x8e\x57\x13\xcd\xd9\x04\x7b\x82\x9d\xa2\xf6\x6e\x76\x31\xe3\xe0\x42\x84\x5a\x35\x09\x3c\x3d\x2e\x97\xd0\x91\x68\x30\xc3\xc9\x4f\xee\xe1\x32\x18\xdf\x5b\x1f\xa0\xe6\x44\xaf\xaa\xc9\x31\x43\xef\x5c\x4c\x0b\xf9\x46\x57\x7c\x9f\xf2\x26\x17\x62\xef\x6d\x00\x22\xa9\xf7\x2a\x4f\xec\xfa\x11\xcc\x79\xb4\x00\xc7\x60\xfd\xa7\xf5\x27\x90\x25\xb5\xa2\xd2\x3c\xcf\x1b\x41\xa9\x15\x9b\xdd\xe7\x27\x7d\xc5\xd8\x02\x78\xd7\xee\x34\x89\xe6\x50\x6e\x85\xee\x9a\xf4\xab\x21\xc6\x69\x3d\x9f\xff\xb6\x5d\xaf\x18\x63\x5f\x01\x00\x00\xff\xff\xc3\xad\x2d\xcd\x54\x01\x00\x00")

func configIniBytes() ([]byte, error) {
	return bindataRead(
		_configIni,
		"config.ini",
	)
}

func configIni() (*asset, error) {
	bytes, err := configIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.ini", size: 340, mode: os.FileMode(420), modTime: time.Unix(1640865712, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd1\xc1\x4b\xc3\x30\x14\x06\xf0\x7b\xfe\x8a\x77\x6c\x61\x17\xbd\xee\x14\xd3\x37\x0d\xa6\xe9\x4c\x5f\xd1\x79\x29\x61\x8d\x5a\x58\x5b\x48\x33\xff\x7e\x31\x5b\x57\x14\x04\xc7\xf6\x6e\x81\x8f\x1f\xe1\xfb\x84\x41\x4e\x08\xc4\xef\x14\xc2\xf6\xc3\xfa\x77\xe7\x47\x96\x30\x00\x80\xb6\x81\xf9\xa4\x26\xbc\x47\x03\x6b\x23\x73\x6e\x36\xf0\x88\x9b\x05\x8b\xb1\xde\x76\xee\x14\x23\x7c\x21\x38\xef\x74\x41\xa0\x2b\xa5\xa0\xd2\xf2\xa9\xc2\x45\x44\x77\x36\xcc\x89\x0c\x85\xcc\xb9\x4a\x6e\x6e\xd3\x73\xd1\xa3\x36\xf4\x57\xd2\x22\xd7\xd8\xe0\xea\xad\x77\x36\xb8\x06\x48\xe6\x58\x12\xcf\xd7\xf0\x2c\xe9\x21\x3e\xe1\xb5\xd0\x08\x19\xae\x78\xa5\x08\x44\x65\x0c\x6a\xaa\xe7\xe0\xc4\xb1\x74\xc9\xd8\x8f\x05\xbc\x1b\x9d\xff\xb4\xa1\x1d\xfa\x69\x85\xfd\xe8\x7c\x7d\x98\xe2\x82\x72\x0f\x3d\x1c\x07\xfe\xe6\xa6\x3d\x0d\xae\xd0\xa0\x16\x58\x9e\xe6\x87\xa4\x6d\x52\x28\x34\x64\xa8\x90\x10\x04\x2f\x05\xcf\xf0\x97\x15\xda\xce\xd5\x6f\x7e\xe8\xe0\xef\x0a\xfe\xfb\xaf\x68\xed\xfb\xd0\xee\x2e\xb1\x58\xba\xfc\x0a\x00\x00\xff\xff\xdf\xbe\xfb\xd3\xcf\x02\x00\x00")

func migrations01_initialUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initialUpSql,
		"migrations/01_initial.up.sql",
	)
}

func migrations01_initialUpSql() (*asset, error) {
	bytes, err := migrations01_initialUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_initial.up.sql", size: 719, mode: os.FileMode(420), modTime: time.Unix(1640863184, 0)}
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
	"config.ini":                   configIni,
	"migrations/01_initial.up.sql": migrations01_initialUpSql,
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
	"config.ini": &bintree{configIni, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"01_initial.up.sql": &bintree{migrations01_initialUpSql, map[string]*bintree{}},
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
