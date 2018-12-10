// Code generated by go-bindata.
// sources:
// assets/sql/001_create_books_table.sql
// DO NOT EDIT!

package assets

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsSql001_create_books_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x4e\x83\x40\x14\x45\xf7\xf3\x15\x37\x6c\xaa\x51\xbe\xa0\x2b\xb4\x63\x42\xa4\x2d\xc2\x90\xb4\x2b\x33\x30\x2f\xf2\x52\x64\x26\xcc\x90\xea\xdf\x1b\x50\xab\x0b\x4c\xdc\xde\x73\xee\xe6\xc4\x31\x6e\x5e\xf9\x65\xd0\x81\x50\x39\x11\xc7\x28\x9f\x32\x70\x0f\x4f\x4d\x60\xdb\x63\x55\xb9\x15\xd8\x83\xde\xa8\x19\x03\x19\x9c\x5b\xea\x11\x5a\xf6\xf8\xfc\x4d\x12\x7b\x68\xe7\x3a\x26\x23\xee\x0b\x99\x28\x09\x95\xdc\x65\x12\xe9\x03\x76\x7b\x05\x79\x48\x4b\x55\xa2\xb6\xf6\xe4\x71\x25\x80\x88\x4d\x84\x52\x16\x69\x92\xcd\xc2\xae\xca\x32\xe4\x45\xba\x4d\x8a\x23\x1e\xe5\xf1\x76\x72\x02\x87\x8e\x22\x28\x79\x50\x17\x69\x06\x7a\x0c\xad\x1d\x96\x88\x1b\xeb\x8e\x7d\x4b\x8b\xf0\x9d\xf4\xe2\xde\x0c\xa4\x03\x99\x67\x1d\x22\xa8\x74\x2b\x4b\x95\x6c\xf3\x1f\x3c\x3a\xf3\x07\x16\xd7\x6b\x21\x7e\x07\xdc\xd8\x73\xff\x9d\xf0\xd2\x6f\x1a\xff\x55\x70\xb0\x5d\x47\x06\xb5\x6e\x4e\x62\x53\xec\xf3\xaf\x86\x73\xb5\xb5\xf8\x08\x00\x00\xff\xff\x3d\xf0\x67\x68\xa9\x01\x00\x00")

func assetsSql001_create_books_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_assetsSql001_create_books_tableSql,
		"assets/sql/001_create_books_table.sql",
	)
}

func assetsSql001_create_books_tableSql() (*asset, error) {
	bytes, err := assetsSql001_create_books_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/sql/001_create_books_table.sql", size: 425, mode: os.FileMode(436), modTime: time.Unix(1544247950, 0)}
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
	"assets/sql/001_create_books_table.sql": assetsSql001_create_books_tableSql,
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
		"sql": &bintree{nil, map[string]*bintree{
			"001_create_books_table.sql": &bintree{assetsSql001_create_books_tableSql, map[string]*bintree{}},
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

