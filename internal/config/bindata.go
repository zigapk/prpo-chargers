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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xdf\x8a\xea\x40\x0c\xc6\xef\xf3\x14\x79\x81\xa3\x15\xcf\x81\x22\xf4\x62\xb4\x39\x6d\x71\xda\x29\x93\xd6\xfd\x23\x32\x8c\x5a\xec\x42\xa1\x32\x1d\xf7\xf9\x97\x71\x71\x59\x90\x90\x8b\xef\x97\xe4\xe3\x23\xfb\x61\xbc\x4c\x07\x90\x2a\x33\x92\x76\x24\x31\xc1\x08\xfe\x17\x92\x8c\x54\x59\x56\x54\x19\x26\xe8\xdd\xad\x83\x8d\xaa\x58\x3d\xe3\x70\x58\x8b\x26\xc7\x04\x83\xd3\xdc\xde\x7c\x3f\x1b\xc6\x0b\x94\xe2\xd5\x70\xf1\x4e\x98\xe0\x22\xba\x2b\x91\xfd\x12\x6b\xb1\xd9\xb6\x35\x7f\x03\xd8\x9f\xad\xb7\x47\x3b\x75\x07\xc8\x15\x37\x77\xb3\x93\x1d\xfa\x71\xf2\x50\x2b\x1d\xc0\xbf\xbf\xcb\x25\xb4\x4c\x1a\x13\xbc\xba\xeb\xf8\xe7\xd4\x5b\x77\xe9\xdc\x04\xb5\x60\x7e\x51\x3a\xc5\x04\xdd\x38\xfa\xd0\x90\xae\x4d\x25\x4a\x7a\xda\x65\x96\xa6\x54\x69\x18\x9c\x3f\x26\x7b\x1c\x3a\x80\xfd\xd4\xb9\xcf\xce\x1d\x40\x16\xdc\x50\x65\x44\x9a\x6a\xe2\x10\x2d\x9a\xdd\xeb\x11\x21\x8e\xe2\x05\x88\xb6\xc9\x0d\x93\xde\x15\x1b\x32\xad\x0e\x0f\xeb\xbd\xbf\xae\xe6\xf3\x9f\xc8\xab\x38\x8a\x23\xc8\x4a\x51\xb3\x11\x75\x61\xb6\xf4\x86\x09\x16\x15\x93\x6e\x1e\xc0\xe4\xa4\xe9\x2b\x00\x00\xff\xff\x41\x32\xad\xb1\x7d\x01\x00\x00")

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

	info := bindataFileInfo{name: "config.ini", size: 381, mode: os.FileMode(420), modTime: time.Unix(1641077025, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd2\x41\x4e\xeb\x30\x10\x06\xe0\xbd\x4f\xf1\x2f\x5b\xa9\x37\xe8\xca\x4d\xa6\xef\x59\xa4\x6e\x71\x26\x82\xb0\x89\x4c\x63\xc0\x52\x9b\x20\xc7\xe5\xfc\xa8\xa1\x34\x14\x81\x44\xc0\x3b\x4b\x33\x9f\x46\xf3\x4f\x62\x48\x32\x81\xe5\x22\x23\x6c\x9f\x6c\x78\x74\xa1\x13\x13\x01\x00\xbe\xc6\xf0\x94\x66\xfc\x23\x4d\x46\x32\xa5\x58\x94\x48\x69\x29\x8b\x8c\x21\x73\xa8\x94\x34\x2b\x2e\xb1\x31\x6a\x25\x4d\x89\x2b\x2a\x67\xa2\x37\x1a\xbb\x77\x67\x83\xe9\x96\x31\xee\xe9\x35\x43\x17\x59\x86\x42\xab\xeb\x82\x66\x3d\xba\xb3\x71\xa8\xa8\xdb\xc3\xfd\xce\xe1\x39\xb8\xad\xef\x7c\xdb\x8c\x40\x4f\xda\xc7\x9e\x3f\x69\x3d\x57\xdb\xe8\xaa\x6d\x70\x36\xba\x1a\xac\x56\x94\xb3\x5c\x6d\x70\xa3\xf8\x7f\xff\xc5\xdd\x5a\xd3\x79\x7b\x49\x61\x0c\x69\xae\x86\xc2\x77\x4e\x4c\xe7\x42\x5c\xc4\x13\x5c\xe7\xc2\x8b\x8d\xbe\x6d\xbe\x88\x68\x7c\x40\xc7\xae\x43\xe7\x42\xf5\xa6\x8c\x4c\xe7\x72\x89\xa7\xd3\x39\x52\xc7\x41\x0c\x2d\xc9\x90\x4e\x28\x3f\x1f\x15\x26\xbe\x9e\x62\xad\x91\x52\x46\x4c\x48\x64\x9e\xc8\x94\x3e\x39\xd1\xef\x5d\xf5\x10\xda\x3d\xbe\xdf\xdd\x4f\xe6\xe9\x9d\x43\x13\xfd\xee\xb7\x8e\x98\xce\x5f\x03\x00\x00\xff\xff\x69\xc8\xec\x48\x1d\x03\x00\x00")

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

	info := bindataFileInfo{name: "migrations/01_initial.up.sql", size: 797, mode: os.FileMode(420), modTime: time.Unix(1641060210, 0)}
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
