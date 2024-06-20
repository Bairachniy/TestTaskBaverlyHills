// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// 000001_create_customers_table.down.sql
// 000001_create_customers_table.up.sql
// 000002_create_customers_settings_table.down.sql
// 000002_create_customers_settings_table.up.sql
// 000003_create_lead_table.down.sql
// 000003_create_lead_table.up.sql
// 000004_create_lead_settings_table.down.sql
// 000004_create_lead_settings_table.up.sql
// generate.go
package migrations

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

var __000001_create_customers_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2e\xc9\xcf\x4d\x2d\x2a\xb6\x06\x04\x00\x00\xff\xff\xe0\xad\x88\x66\x1f\x00\x00\x00")

func _000001_create_customers_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_create_customers_tableDownSql,
		"000001_create_customers_table.down.sql",
	)
}

func _000001_create_customers_tableDownSql() (*asset, error) {
	bytes, err := _000001_create_customers_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_create_customers_table.down.sql", size: 31, mode: os.FileMode(438), modTime: time.Unix(1718892525, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_create_customers_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2e\xc9\xcf\x4d\x2d\x2a\xd6\xe0\xe5\x52\x50\x50\x50\xc8\x4c\x51\x28\x4e\x2d\xca\x4c\xcc\x51\x08\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\xd4\x81\xc8\x95\x16\xa7\x16\xe5\x25\xe6\xa6\x2a\x84\x39\x06\x39\x7b\x38\x06\x29\x68\x98\x1a\x68\x2a\x84\xfa\x79\x06\x86\xba\x82\xcd\xf4\x0b\xf5\xf1\x81\xaa\x2d\x48\x2c\x2e\x2e\xcf\x2f\x4a\x41\x55\x8b\xa6\x28\x35\x37\x31\x33\x07\xa1\xc2\xd8\x00\xd3\x38\x5e\x2e\x4d\x6b\x40\x00\x00\x00\xff\xff\xbc\x33\x73\xf5\xbf\x00\x00\x00")

func _000001_create_customers_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_create_customers_tableUpSql,
		"000001_create_customers_table.up.sql",
	)
}

func _000001_create_customers_tableUpSql() (*asset, error) {
	bytes, err := _000001_create_customers_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_create_customers_table.up.sql", size: 191, mode: os.FileMode(438), modTime: time.Unix(1718892525, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_create_customers_settings_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2e\xc9\xcf\x4d\x2d\x2a\x8e\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\xb6\x06\x04\x00\x00\xff\xff\x58\xea\xc4\x3e\x28\x00\x00\x00")

func _000002_create_customers_settings_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_create_customers_settings_tableDownSql,
		"000002_create_customers_settings_table.down.sql",
	)
}

func _000002_create_customers_settings_tableDownSql() (*asset, error) {
	bytes, err := _000002_create_customers_settings_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_create_customers_settings_table.down.sql", size: 40, mode: os.FileMode(438), modTime: time.Unix(1718895113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_create_customers_settings_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcd\xcd\x4a\x03\x31\x14\xc5\xf1\x7d\xa1\xef\x70\x96\x2d\xf8\x06\xae\xe2\xf4\x16\x06\x63\x2a\x99\x08\x16\x91\x30\x3a\xa9\x5e\x48\x26\x25\xf7\x8a\xfa\xf6\x82\x1f\x20\x4a\xd7\x3f\xce\xff\x74\x9e\x4c\x20\x04\x73\x61\x09\xfd\x16\x6e\x17\x40\xb7\xfd\x10\x06\x3c\xbe\x88\xd6\x92\x9a\x44\x49\xaa\x3c\x3f\xc9\x6a\xb9\x00\x00\x9e\x20\xa9\xf1\x98\x71\xed\xfb\x2b\xe3\xf7\xb8\xa4\xfd\xd9\x97\xfd\x8c\x22\x4f\xe0\x59\xe1\x69\x4b\x9e\x5c\x47\xbf\x7a\x2b\x9e\xd6\xd8\x39\x6c\xc8\x52\x20\x74\x66\xe8\xcc\x86\x3e\xaf\xdd\x8d\xb5\xdf\xa5\x03\x67\x4d\x2d\xbe\x3e\xc7\x43\xab\x05\xca\x25\x89\x8e\xe5\xf8\x8f\xb5\x9e\xc2\x63\xe3\xda\x58\xdf\x21\x65\xcc\x99\x67\xfd\xc3\x55\x84\x1f\x38\xb3\x72\x12\x68\x7a\xd3\xbb\xfb\xe5\x62\x7d\xfe\x11\x00\x00\xff\xff\xb5\xce\x91\xb5\x14\x01\x00\x00")

func _000002_create_customers_settings_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_create_customers_settings_tableUpSql,
		"000002_create_customers_settings_table.up.sql",
	)
}

func _000002_create_customers_settings_tableUpSql() (*asset, error) {
	bytes, err := _000002_create_customers_settings_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_create_customers_settings_table.up.sql", size: 276, mode: os.FileMode(438), modTime: time.Unix(1718895068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_create_lead_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\x49\x4d\x4c\x29\xb6\x06\x04\x00\x00\xff\xff\x2c\x44\xfe\x09\x1b\x00\x00\x00")

func _000003_create_lead_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_create_lead_tableDownSql,
		"000003_create_lead_table.down.sql",
	)
}

func _000003_create_lead_tableDownSql() (*asset, error) {
	bytes, err := _000003_create_lead_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_create_lead_table.down.sql", size: 27, mode: os.FileMode(438), modTime: time.Unix(1718895824, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_create_lead_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcc\x41\x8b\x82\x40\x18\xc6\xf1\xbb\xe0\x77\x78\x8e\x0a\x7b\x59\x41\x58\xd8\xd3\xac\xcc\xd2\x90\x59\x8d\x63\xe4\x29\x5e\xf2\xad\x06\x46\x83\x71\x88\xe8\xd3\x07\xe5\xc9\x53\xd7\xe7\xf9\xf1\x2f\xb4\x14\x46\xc2\x88\xbf\x52\x42\xfd\xa3\x5a\x1b\xc8\xbd\xaa\x4d\x0d\xc7\xd4\x8d\x49\x1c\x01\x80\xed\x30\xb2\xb7\xe4\xb0\xd1\x6a\x25\x74\x8b\xa5\x6c\xbf\xde\xdf\xc9\xfa\x31\x1c\x06\xea\x19\x3b\xa1\x8b\x85\xd0\x48\xbe\xb3\x9f\xf4\xd5\xaa\x9a\xb2\x9c\x9c\xa3\x8f\x18\xf7\x64\xdd\x8c\x34\x95\xda\x36\x72\x2e\xcf\x3c\x74\xec\x11\xf8\x1e\xa6\x25\xd8\x9e\x1f\xd7\x81\x71\x23\x7f\xbc\x90\x4f\xb2\x3c\x4f\xe3\x28\xfd\x7d\x06\x00\x00\xff\xff\x55\x91\x0b\xc0\xe7\x00\x00\x00")

func _000003_create_lead_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_create_lead_tableUpSql,
		"000003_create_lead_table.up.sql",
	)
}

func _000003_create_lead_tableUpSql() (*asset, error) {
	bytes, err := _000003_create_lead_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_create_lead_table.up.sql", size: 231, mode: os.FileMode(438), modTime: time.Unix(1718895572, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_create_lead_settings_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\x49\x4d\x4c\x29\x8e\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\xb6\x06\x04\x00\x00\xff\xff\xd0\x7c\xb2\x8a\x24\x00\x00\x00")

func _000004_create_lead_settings_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_create_lead_settings_tableDownSql,
		"000004_create_lead_settings_table.down.sql",
	)
}

func _000004_create_lead_settings_tableDownSql() (*asset, error) {
	bytes, err := _000004_create_lead_settings_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_create_lead_settings_table.down.sql", size: 36, mode: os.FileMode(438), modTime: time.Unix(1718895824, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_create_lead_settings_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\xc1\x4a\x03\x31\x14\x45\xf7\x85\xfe\xc3\x5d\xb6\xe0\x1f\xb8\x8a\xd3\x57\x18\x1c\x53\xc9\x44\xb0\x88\x84\x91\xc4\xf1\x61\x9a\x48\xde\x13\xf5\xef\xc5\x5a\xc4\x45\xb7\xf7\x70\xce\xed\x1c\x19\x4f\xf0\xe6\x6a\x20\xf4\x5b\xd8\x9d\x07\xdd\xf7\xa3\x1f\x91\xd3\x14\x25\x48\x52\xe5\x32\xcb\x6a\xb9\x00\x00\x8e\x90\xd4\x78\xca\xb8\x75\xfd\x8d\x71\x7b\x5c\xd3\xfe\xe2\x97\xfd\x08\x81\x23\xb8\x68\x9a\x53\x83\xa3\x2d\x39\xb2\x1d\x9d\x5a\x2b\x8e\x6b\xec\x2c\x36\x34\x90\x27\x74\x66\xec\xcc\x86\x8e\x97\xf6\x6e\x18\x4e\x95\x8f\xda\x5e\xb9\xcc\xe1\xa5\xbe\x37\x09\xcf\xad\x1e\xa0\xe9\x53\xcf\x52\xad\xff\xd9\x5b\xe3\xda\x58\xbf\x20\x87\x29\x67\x2e\x7f\x7b\x15\xe1\x27\xce\xac\x9c\xe4\x28\x3c\x3c\x2e\x17\xeb\xcb\xef\x00\x00\x00\xff\xff\xbd\xd8\xd4\x6b\xfc\x00\x00\x00")

func _000004_create_lead_settings_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_create_lead_settings_tableUpSql,
		"000004_create_lead_settings_table.up.sql",
	)
}

func _000004_create_lead_settings_tableUpSql() (*asset, error) {
	bytes, err := _000004_create_lead_settings_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_create_lead_settings_table.up.sql", size: 252, mode: os.FileMode(438), modTime: time.Unix(1718896022, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generateGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\xb1\x0d\xc5\x20\x0c\x04\xd0\x9e\x29\x6e\x01\xa0\xff\xdb\xdc\x4f\xac\x13\x42\xb1\x11\xf1\xfe\x4a\x93\x22\xf5\x7b\x8b\xc7\xa4\x0c\xd7\xd0\x66\x8e\xf0\xbb\x94\xde\x15\x3f\x99\xdb\x66\x1a\x14\xf5\x3f\xfc\x64\x12\x75\x4d\x7d\x26\x6a\xe0\xa5\xa6\x00\x5a\x79\x02\x00\x00\xff\xff\xe6\x24\xd8\x86\x4e\x00\x00\x00")

func generateGoBytes() ([]byte, error) {
	return bindataRead(
		_generateGo,
		"generate.go",
	)
}

func generateGo() (*asset, error) {
	bytes, err := generateGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generate.go", size: 78, mode: os.FileMode(438), modTime: time.Unix(1718897028, 0)}
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
	"000001_create_customers_table.down.sql":          _000001_create_customers_tableDownSql,
	"000001_create_customers_table.up.sql":            _000001_create_customers_tableUpSql,
	"000002_create_customers_settings_table.down.sql": _000002_create_customers_settings_tableDownSql,
	"000002_create_customers_settings_table.up.sql":   _000002_create_customers_settings_tableUpSql,
	"000003_create_lead_table.down.sql":               _000003_create_lead_tableDownSql,
	"000003_create_lead_table.up.sql":                 _000003_create_lead_tableUpSql,
	"000004_create_lead_settings_table.down.sql":      _000004_create_lead_settings_tableDownSql,
	"000004_create_lead_settings_table.up.sql":        _000004_create_lead_settings_tableUpSql,
	"generate.go": generateGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"000001_create_customers_table.down.sql":          &bintree{_000001_create_customers_tableDownSql, map[string]*bintree{}},
	"000001_create_customers_table.up.sql":            &bintree{_000001_create_customers_tableUpSql, map[string]*bintree{}},
	"000002_create_customers_settings_table.down.sql": &bintree{_000002_create_customers_settings_tableDownSql, map[string]*bintree{}},
	"000002_create_customers_settings_table.up.sql":   &bintree{_000002_create_customers_settings_tableUpSql, map[string]*bintree{}},
	"000003_create_lead_table.down.sql":               &bintree{_000003_create_lead_tableDownSql, map[string]*bintree{}},
	"000003_create_lead_table.up.sql":                 &bintree{_000003_create_lead_tableUpSql, map[string]*bintree{}},
	"000004_create_lead_settings_table.down.sql":      &bintree{_000004_create_lead_settings_tableDownSql, map[string]*bintree{}},
	"000004_create_lead_settings_table.up.sql":        &bintree{_000004_create_lead_settings_tableUpSql, map[string]*bintree{}},
	"generate.go": &bintree{generateGo, map[string]*bintree{}},
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
