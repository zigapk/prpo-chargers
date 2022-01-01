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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xef\x6e\x9b\x30\x14\xc5\xbf\xfb\x29\xfc\x02\x21\xce\xbf\x0e\x55\xe2\x83\x09\x2e\x58\x05\xcc\x30\x64\x6c\x51\x64\x39\xe0\x41\x26\xb7\xce\x8c\x69\xd5\x3c\xfd\x44\xa7\x4a\x95\xaa\xfb\xe9\x9c\xab\xdf\xd1\xbd\xe7\xa8\x4d\x3f\x9e\x40\xca\x62\x91\x92\x03\x49\x61\x00\x11\x78\xa0\x29\x11\x29\x8b\x63\x9a\xc7\x30\x80\xce\x4e\x0a\xec\x59\xce\xd9\x57\x7b\x06\x0b\x5c\x25\x30\x80\x73\xd2\x52\x4e\x6e\xf0\xb4\xe9\x41\x86\x1b\xc1\xe9\x2f\x02\x03\xb8\x42\xef\x0a\xc7\x9f\x44\x88\xf7\x8f\x75\xc1\xff\x1b\xe0\xd8\x49\x27\xcf\x72\x54\x27\x90\x30\x5e\xc1\x00\x5e\xed\xd5\x2c\x3a\xb3\x98\x46\x65\x17\xab\xed\x66\x83\xd0\x6e\x81\xbc\xb3\xd7\x9d\x3d\xf3\xdc\x5d\xfa\x8b\x93\xda\xb4\x4a\x3e\x7b\xad\x79\x02\x05\x2b\x67\x6a\xbd\x43\x77\x08\xd4\x9c\x94\x1f\x11\xed\x20\x6d\xaf\xec\x08\x0a\xcc\xf9\x0f\x56\x46\x30\x80\x9b\xe4\x30\xfe\x69\xd6\x61\xbf\x45\xda\x6d\x6f\x04\x44\xa1\xc8\x71\x46\xbe\x30\x9c\xa7\x22\x63\xd1\xbc\xb0\xea\xef\x74\xb1\x0a\x80\xe3\xa8\xec\x8b\xb2\x27\x90\x52\x5e\x91\x5c\xe0\x28\x2a\x09\x9f\x1f\x41\xde\xfb\x7c\xdc\xe2\x23\xb4\x02\xb8\xae\x12\xc1\x49\x79\xa0\x7b\x22\xea\x72\xae\x77\x70\xee\x7a\xbf\x5c\x6a\xd3\x4a\x3d\x98\xd1\xdd\xfb\x08\x21\x10\x67\xb8\xe0\x02\x17\x54\x3c\x92\x9f\x30\x80\x98\xde\x24\x7f\x0b\xfd\xb7\xef\xa1\xed\xd7\xda\x35\x9a\xdd\xb1\x07\x9a\x7f\x7b\xe5\xfa\x89\x36\x2f\xbf\xc5\xce\xbf\xbd\xd6\xff\x02\x00\x00\xff\xff\xd0\x59\x2d\x40\xbf\x01\x00\x00")

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

	info := bindataFileInfo{name: "config.ini", size: 447, mode: os.FileMode(420), modTime: time.Unix(1641066913, 0)}
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
