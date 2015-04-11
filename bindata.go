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

var _migrations_20150411142140_add_fiscalyear_foreignkey_to_fiscalperiods_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xcd\xb1\xae\xc2\x20\x14\x80\xe1\x9d\xa7\x38\x5b\x87\x9b\xde\x17\x68\x1c\xaa\x74\x43\x6b\x2a\x9d\x09\xd2\x23\x39\x09\x9c\x12\xc0\x34\xbe\xbd\x0e\x0e\x0e\x0e\xce\x7f\xf2\xfd\x6d\x0b\x7f\x91\x7c\xb6\x15\x61\x4e\xe2\x32\x68\x70\x81\x90\xab\x89\xc4\x26\x62\x29\xd6\x63\x81\x1d\x34\x9b\xcd\x4c\xec\x9b\x4e\x88\x5e\xe9\x61\x02\xdd\xef\xd5\x00\xe9\x7e\x0d\xe4\xfe\x6f\x54\x9c\x0d\x26\x61\xa6\x75\x29\xd0\x4b\x09\x87\x51\xcd\xc7\x13\xbc\xcb\x03\x6d\x36\xb4\x00\x71\x45\x8f\xf9\xa5\x7c\xae\xe5\xba\xf1\x0f\xac\x9c\xc6\xf3\x77\xb7\x13\xcf\x00\x00\x00\xff\xff\x61\xe0\xac\x37\xca\x00\x00\x00")

func migrations_20150411142140_add_fiscalyear_foreignkey_to_fiscalperiods_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_20150411142140_add_fiscalyear_foreignkey_to_fiscalperiods_sql,
		"migrations/20150411142140-add-fiscalYear-foreignKey-to-fiscalPeriods.sql",
	)
}

func migrations_20150411142140_add_fiscalyear_foreignkey_to_fiscalperiods_sql() (*asset, error) {
	bytes, err := migrations_20150411142140_add_fiscalyear_foreignkey_to_fiscalperiods_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/20150411142140-add-fiscalYear-foreignKey-to-fiscalPeriods.sql", size: 202, mode: os.FileMode(420), modTime: time.Unix(1428755005, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations_20150411171420_create_planyears_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x91\xcd\x6e\xea\x30\x10\x85\xd7\xd7\x4f\x31\x3b\x40\x17\xfa\x02\xa8\x8b\xb4\x18\x29\x6a\xf8\x51\x62\xd4\xd2\x8d\x35\x75\x86\x74\x24\xc7\x44\xb6\x03\xe2\xed\x1b\x28\x95\x22\x81\xba\xa9\x57\x63\x69\xce\x99\x99\xf3\x4d\x26\xf0\xbf\xe6\xca\x63\x24\xd8\x34\xa2\x90\x0a\x8c\x65\x72\x51\xd7\xec\x74\x4d\x21\x60\x45\x01\x1e\x61\x70\x44\xef\xd8\x55\x83\xa9\x10\xcf\xb9\x4c\x94\x04\x95\x3c\x65\x12\xd2\x39\x2c\x57\x0a\xe4\x5b\x5a\xa8\x02\x9a\xf6\xc3\xb2\x79\x68\x2c\x3a\x7d\x22\xf4\x01\x86\x02\x80\x4b\xf8\xe5\x15\x32\x4f\x93\xec\xbb\x5e\xe7\xe9\x22\xc9\xb7\xf0\x22\xb7\xe3\x4e\x68\x3c\x75\x9b\x95\x1a\xe3\x3d\xa1\x4a\x17\xb2\x50\xc9\x62\xad\xde\x2f\x3b\x2c\x37\x59\x06\x33\x39\x4f\x36\x59\xf7\x59\xbd\x0e\x47\x67\x8f\xb6\x29\xff\xec\xb1\xe3\x60\xd0\x5e\x2e\xd2\xb7\xc7\xb0\x8b\x54\x91\xbf\xd4\x3f\x1e\x63\xf1\x0f\x0f\xe4\xbb\xf4\x74\x89\xa7\xa0\xf7\x3b\xcd\xd6\xba\x2e\xd0\xab\xc8\xb5\x35\x79\x36\x77\xfa\xcc\x27\xdb\xd2\x93\xd3\x06\x3d\xf5\xfa\x4a\xda\x61\x6b\xa3\x3e\xa0\xc1\xc8\x7b\xa7\xcf\x73\x3d\x85\xd8\xf7\x13\xa3\x0e\x50\x9f\xea\x6c\x7f\x74\x62\x96\xaf\xd6\x57\x5e\x37\x84\xa6\xe2\x2b\x00\x00\xff\xff\x77\x16\xbc\xba\x05\x02\x00\x00")

func migrations_20150411171420_create_planyears_sql_bytes() ([]byte, error) {
	return bindata_read(
		_migrations_20150411171420_create_planyears_sql,
		"migrations/20150411171420-create-planYears.sql",
	)
}

func migrations_20150411171420_create_planyears_sql() (*asset, error) {
	bytes, err := migrations_20150411171420_create_planyears_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "migrations/20150411171420-create-planYears.sql", size: 517, mode: os.FileMode(420), modTime: time.Unix(1428765563, 0)}
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
	"migrations/20150411142140-add-fiscalYear-foreignKey-to-fiscalPeriods.sql": migrations_20150411142140_add_fiscalyear_foreignkey_to_fiscalperiods_sql,
	"migrations/20150411171420-create-planYears.sql": migrations_20150411171420_create_planyears_sql,
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
		"20150411142140-add-fiscalYear-foreignKey-to-fiscalPeriods.sql": &_bintree_t{migrations_20150411142140_add_fiscalyear_foreignkey_to_fiscalperiods_sql, map[string]*_bintree_t{
		}},
		"20150411171420-create-planYears.sql": &_bintree_t{migrations_20150411171420_create_planyears_sql, map[string]*_bintree_t{
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

