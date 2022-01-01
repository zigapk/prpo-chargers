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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\x61\x6b\xe2\x40\x10\x86\xbf\xcf\xaf\xd8\x3f\x70\xba\x9e\xe7\x5d\x10\xf2\x61\x63\xd6\xb8\x98\x64\x73\x99\xc4\xda\x8a\x2c\xab\xa6\x49\x21\x25\xb2\x59\x15\xfd\xf5\x65\x2d\x2d\x05\x19\xe6\xc3\xfb\xce\xcc\xcb\xc3\x6c\xda\xae\xee\xb7\x10\xcb\x48\xc5\x7c\xc5\x63\xe2\x13\x0a\x73\x11\x73\x15\xcb\x28\x12\x69\x44\x7c\x62\xcd\xa9\x82\x99\x4c\x51\x3e\xda\xee\x30\x63\xc5\x82\xf8\xc4\x25\x0d\xf5\xc9\x36\x83\xb6\xab\x21\x61\x6b\x85\xe2\x85\x13\x9f\x8c\xe8\x5d\xb1\xe8\x87\x08\xd8\x6c\x59\x66\xf8\x69\xc0\xe6\xa0\xad\xde\xe9\xbe\xda\xc2\x42\x62\x71\x0f\xdb\xeb\xb6\xe9\x7a\x0b\x99\xcc\x9d\x31\xf9\x33\x1e\x43\x89\x3c\x27\x3e\x39\x9a\x63\xf7\x6b\xdf\x68\x53\x57\xa6\x87\x8c\x21\x3e\xc9\x3c\x24\x3e\x31\x5d\x67\x5d\x43\x18\xa8\x94\x25\xfc\x61\x17\x31\x56\x89\x0c\xdd\xe0\xf0\xd6\xeb\x5d\x5b\x01\x6c\xfa\xca\x9c\x2b\xb3\x85\x58\x60\xc1\x53\xc5\xc2\x30\xe7\xe8\xd0\xe8\xe0\x5e\x5f\x08\x1e\xa5\x23\x60\x65\xb1\x50\xc8\xf3\x95\x98\x71\x55\xe6\xee\x61\x8d\xb5\xc7\xe9\x70\xf8\x8d\x3c\xf5\x28\xa5\x10\x25\x2c\x43\xc5\x32\xa1\x96\xfc\x99\xf8\x84\x89\x9b\xc6\x6b\xe0\x5d\xff\x07\xa6\xfe\xdd\xda\x75\x2b\xff\xca\xb9\x48\xff\x5d\xb0\x7d\x17\xeb\xf3\xab\x9a\x78\xb7\x4b\xf9\x11\x00\x00\xff\xff\x51\xab\x0c\xd2\x91\x01\x00\x00")

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

	info := bindataFileInfo{name: "config.ini", size: 401, mode: os.FileMode(420), modTime: time.Unix(1640870996, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd2\xd1\x4a\xfb\x30\x14\x06\xf0\xeb\x7f\x9e\xe2\xbb\xdc\x60\x6f\xb0\xab\xac\x3d\xfb\x1b\xec\xb2\x99\x9e\xa2\xf5\xa6\xc4\x35\x6a\x60\x6b\x25\xcd\x7c\x7e\x59\xdd\x56\x26\x08\x56\x3c\x77\x85\xc3\x8f\xe6\xfb\x4e\x62\x48\x32\x81\xe5\x22\x23\x6c\x5f\x6d\x78\x71\xa1\x13\x13\x01\x00\xbe\xc6\x30\x4a\x33\xfe\x93\x26\x23\x99\x52\x2c\x4a\xa4\xb4\x94\x45\xc6\x90\x39\x54\x4a\x9a\x15\x97\xd8\x18\xb5\x92\xa6\xc4\x2d\x95\x33\xd1\x1b\x8d\xdd\xbb\x8b\xc1\xf4\xc0\x18\x37\x7a\xcd\xd0\x45\x96\xa1\xd0\xea\xae\xa0\x59\x8f\xee\x6c\x1c\x36\xea\xf6\xf0\xb4\x73\x78\x0b\x6e\xeb\x3b\xdf\x36\xff\xc6\xd2\x27\xb3\x6d\xfe\xd4\xec\xd1\xda\x46\x57\x6d\x83\xb3\xd1\xd5\x60\xb5\xa2\x9c\xe5\x6a\x83\x7b\xc5\x37\xfd\x27\x1e\xd7\x9a\x2e\x49\x26\x85\x31\xa4\xb9\x1a\x16\xcf\x9c\x98\xce\x85\xb8\xaa\x2a\xb8\xce\x85\x77\x1b\x7d\xdb\x9c\xeb\x3a\x74\x2e\x54\x9f\x9d\x8d\x0c\xfa\x3a\x89\xd3\x15\x1c\xa9\x63\xe9\x86\x96\x64\x48\x27\x94\x5f\xee\x03\x13\x5f\x4f\xb1\xd6\x48\x29\x23\x26\x24\x32\x4f\x64\x4a\x5f\x9c\xe8\xf7\xae\x7a\x0e\xed\x1e\xdf\x3f\xfd\x27\xff\xd3\x3b\x87\x26\xfa\xdd\x6f\x1d\x31\x9d\x7f\x04\x00\x00\xff\xff\xdf\xb0\xa6\x09\xe8\x02\x00\x00")

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

	info := bindataFileInfo{name: "migrations/01_initial.up.sql", size: 744, mode: os.FileMode(420), modTime: time.Unix(1640873985, 0)}
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
