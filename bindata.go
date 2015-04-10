package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _migrations_20150410193710_create_fiscalyear_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x92\x4f\x4b\x03\x31\x10\xc5\xef\xf9\x14\x73\x6b\x8b\xd6\x93\x78\x29\x1e\x56\x9b\xc2\xe2\xf6\x0f\xbb\x29\x5a\x2f\x61\x9a\x4c\xd7\x60\x37\x5d\x93\xac\xa5\xdf\xde\xac\x28\xd4\xca\x22\xa2\xb9\x65\xc2\xef\x65\xde\xcc\x1b\x0e\xe1\xac\x32\xa5\xc3\x40\xb0\xac\x59\xc1\x05\xa8\xad\x21\x1b\x64\x65\xac\xac\xc8\x7b\x2c\xc9\xc3\x35\xf4\xf6\xe8\xac\xb1\x65\x6f\xc4\xd8\x6d\xce\x13\xc1\x41\x24\x37\x19\x87\x74\x02\xb3\xb9\x00\xfe\x90\x16\xa2\x80\xba\x59\x6f\x8d\xba\xd8\x18\xaf\x70\x2b\x0f\x84\xce\x43\x9f\x01\x18\x0d\x9d\xa7\xe0\x79\x9a\x64\xa7\xd5\x45\x9e\x4e\x93\x7c\x05\x77\x7c\x75\x1e\x05\x2c\x56\xd4\x25\xa0\x9e\xd0\xa1\x0a\xe4\xe0\x15\xdd\x21\x36\xd9\xbf\xba\x1c\xbc\x77\x35\x5b\x66\x59\x4b\x2b\x47\xd1\xa1\x96\x18\xbe\xd3\x22\x9d\xf2\x42\x24\xd3\x85\x78\x3c\xaa\x7e\xd2\x30\xe6\x93\x64\x99\xc5\xcb\xfc\xbe\x3f\x68\xb5\x9a\x5a\xff\x9b\x56\x3b\xa0\x2e\x57\xc6\x06\x2a\xe9\xf4\xf9\xd8\xd5\xba\xf1\xc6\xc6\x15\x49\x8d\x07\xff\x37\x5a\x6e\x8c\xf3\x41\xbe\x34\xe8\xda\x39\xfe\x4c\xc7\xf5\x92\xd5\xe8\xe4\x9e\xe8\xd9\xff\xe6\x6f\x36\x88\x11\x3a\xce\xdd\x78\xb7\xb7\x6c\x9c\xcf\x17\x1f\x89\xfa\x9a\xa1\x9a\x9c\xd9\x69\x3f\x62\x6f\x01\x00\x00\xff\xff\x56\xcd\x81\xed\xab\x02\x00\x00")

func migrations_20150410193710_create_fiscalyear_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_20150410193710_create_fiscalyear_sql,
		"migrations/20150410193710-create-fiscalYear.sql",
	)
}

func migrations_20150410193710_create_fiscalyear_sql() (*asset, error) {
	bytes, err := migrations_20150410193710_create_fiscalyear_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/20150410193710-create-fiscalYear.sql", size: 683, mode: os.FileMode(420), modTime: time.Unix(1428687678, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"migrations/20150410193710-create-fiscalYear.sql": migrations_20150410193710_create_fiscalyear_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"migrations": &_bintree_t{nil, map[string]*_bintree_t{
		"20150410193710-create-fiscalYear.sql": &_bintree_t{migrations_20150410193710_create_fiscalyear_sql, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

