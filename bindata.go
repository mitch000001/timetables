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

var _migrations_20150410193710_create_fiscalyear_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x91\xd1\x4a\xc3\x30\x18\x85\xef\xf3\x14\xff\xdd\x36\x74\xbe\xc0\xf0\x22\xda\x0c\x8a\xed\x5a\xda\x14\x9d\x37\x21\x6b\xfe\x95\x60\x1b\x6b\x92\x52\xf6\xf6\xa6\x4e\x41\xd0\x22\xb8\x5c\x25\x24\xdf\xc9\xf9\xcf\x59\xaf\xe1\xaa\xd3\x8d\x95\x1e\xa1\xea\x49\xc9\x38\xd4\xad\x46\xe3\x45\xa7\x8d\xe8\xd0\x39\xd9\xa0\x83\x5b\x58\x8c\xd2\x1a\x6d\x9a\xc5\x86\x90\xfb\x82\x51\xce\x80\xd3\xbb\x84\x41\xbc\x85\x5d\xc6\x81\x3d\xc5\x25\x2f\xa1\x1f\x0e\xad\xae\x6f\x8e\xda\xd5\xb2\x15\x27\x94\xd6\xc1\x92\x00\x68\x05\xb3\xab\x64\x45\x4c\x93\xf3\x3e\x2f\xe2\x94\x16\x7b\x78\x60\xfb\xeb\x80\xd5\x16\x83\x33\x25\xa4\xff\x89\xf1\x38\x65\x25\xa7\x69\xce\x9f\x3f\x1c\xec\xaa\x24\x81\x88\x6d\x69\x95\x84\x43\xf6\xb8\x5c\x4d\x0a\x43\xaf\x2e\x54\x98\x86\x98\xb3\xae\x8d\xc7\x06\xcf\xd7\x5f\x0a\x13\x73\x18\x9c\x36\x21\x3c\xa1\xe4\xc9\xfd\x87\x11\x47\x6d\x9d\x17\x6f\x83\xb4\x3e\xbc\x9d\x63\x42\xc8\x68\x94\xb4\x62\x44\x7c\x71\x7f\xff\x43\x56\xa1\xbe\xef\x9d\x47\xaf\xa3\x21\x51\x91\xe5\x9f\x6d\xfe\xd2\xdf\x86\xbc\x07\x00\x00\xff\xff\x66\x62\xbc\x32\x25\x02\x00\x00")

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

	info := bindata_file_info{name: "migrations/20150410193710-create-fiscalYear.sql", size: 549, mode: os.FileMode(420), modTime: time.Unix(1428691900, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_20150410204602_create_fiscalperiods_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\xbf\x6e\xeb\x20\x14\x87\x77\x9e\xe2\x6c\x49\x74\x6f\xfa\x02\x55\x07\x5a\x13\xc9\xaa\x1d\x5b\x80\xd5\xa6\x0b\x22\x86\x5a\x47\xb2\x89\x05\x58\x51\xdf\xbe\x38\x4d\x2a\x0f\x59\xca\x04\xe2\xfb\x9d\x3f\xdf\x76\x0b\xff\x06\xec\xbc\x8e\x16\x9a\x91\x08\x26\xa1\xed\xd1\xba\xa8\x06\x74\x6a\xb0\x21\xe8\xce\x06\x78\x82\xd5\x59\x7b\x87\xae\x5b\x3d\x12\xf2\xc2\x19\x95\x0c\x24\x7d\x2e\x18\xe4\x3b\xd8\x57\x12\xd8\x7b\x2e\xa4\x80\x71\x3a\xf6\xd8\x3e\x7c\x62\x68\x75\xaf\x46\xeb\xf1\x64\x02\xac\x09\x00\x1a\x58\x1c\xc1\x78\x4e\x8b\x9f\x7b\xcd\xf3\x92\xf2\x03\xbc\xb2\xc3\xff\x04\xb6\xde\xa6\x69\x8c\xd2\x71\xfe\x94\x79\xc9\x84\xa4\x65\x2d\x3f\x2e\x7d\xf6\x4d\x51\x40\xc6\x76\xb4\x29\xd2\xa3\x7a\x5b\x6f\xe6\xcc\x34\x9a\x3f\x67\x8e\x53\x40\x97\x16\x54\x46\x7f\x05\x40\x17\x6d\x67\xfd\x65\xa0\x5b\x66\xa6\x42\xd4\x3e\x86\x6b\x61\x98\xbb\xdc\x56\x58\x52\xd6\x99\x5f\xe6\x3e\x45\x36\x49\xdc\xd2\x76\x76\x3a\x3b\x92\xf1\xaa\xbe\x7a\xbc\x6b\x2e\x65\xbe\x03\x00\x00\xff\xff\xa0\x6e\x9f\xc8\xa2\x01\x00\x00")

func migrations_20150410204602_create_fiscalperiods_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_20150410204602_create_fiscalperiods_sql,
		"migrations/20150410204602-create-fiscalPeriods.sql",
	)
}

func migrations_20150410204602_create_fiscalperiods_sql() (*asset, error) {
	bytes, err := migrations_20150410204602_create_fiscalperiods_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/20150410204602-create-fiscalPeriods.sql", size: 418, mode: os.FileMode(420), modTime: time.Unix(1428691878, 0)}
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
	"migrations/20150410204602-create-fiscalPeriods.sql": migrations_20150410204602_create_fiscalperiods_sql,
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
		"20150410204602-create-fiscalPeriods.sql": &_bintree_t{migrations_20150410204602_create_fiscalperiods_sql, map[string]*_bintree_t{
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

