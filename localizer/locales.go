// Code generated for package localizer by go-bindata DO NOT EDIT. (@generated)
// sources:
// localizer/locales/en.json
package localizer

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

var _localizerLocalesEnJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\x8a\xe4\x46\x0c\xbd\xcf\x57\x88\xbe\xf4\x65\x70\xb3\xd7\xb9\x65\x09\x0b\x03\xd9\xb0\x0c\xb3\x0b\x81\x06\x23\x57\xc9\xed\x4a\xec\x92\x53\x92\xdb\xdd\x0c\xfd\xef\x41\x55\xb6\x67\xd3\xd9\x1c\xf6\xd6\xad\x92\x9e\xde\x7b\x92\xfc\xf6\x00\xb0\x43\xef\x43\x3c\xd5\xe7\x0f\x75\xc3\x97\xdd\x13\xec\x8e\xf1\x97\x1c\x82\x6f\x1f\xe0\x1b\x9e\x12\x46\x85\x86\x2f\x8f\xb0\x7f\x7b\xab\x3e\xf2\xe5\x77\x1c\xe8\x76\xdb\x57\xdb\xe3\x1c\xfa\x1e\x70\x52\x1e\x50\x83\xc3\xbe\xbf\xc2\x34\x9e\x12\x7a\x02\xed\x82\x58\x71\x05\x5f\x7a\x42\x21\x70\x1c\x25\x78\x4a\x4b\x86\xb5\xd1\x8e\x40\x78\x4a\x8e\x4a\x66\xa7\x3a\xca\xd3\xe1\x30\xcf\x73\x75\x2e\x2d\xa6\xb1\x72\x3c\x1c\x3c\x3b\x39\x34\x7c\x21\x39\xb4\x9c\x06\xd4\x63\xdc\x3d\x9a\x86\x86\x2f\x35\x7a\x6f\xec\x17\xee\xa6\x65\x7b\xf2\x4c\x52\x47\xd6\xba\xc3\x33\xd5\x03\x29\x7a\x54\xac\xff\x14\x8e\x75\x1b\x7a\xb2\xba\xd7\x8e\x60\xbf\x3e\x55\xf6\xb4\x07\x7b\x83\x96\x53\xa6\xd8\xf0\xe5\xce\x01\x98\x51\x20\xb2\x42\xcb\x53\xf4\x15\x7c\x34\x66\x90\xe8\xef\x29\xa4\x45\x7a\x46\x08\x11\x38\x99\x66\x83\x5a\x4d\x53\x06\x4f\x4a\x69\x08\x91\x32\xfe\x98\xf8\x9c\x9d\x09\x9a\x81\x07\xf3\xaf\xe5\x54\xc1\x73\x0b\x57\x9e\x4a\x60\x61\xf2\x08\x63\xf1\x13\xbd\x07\xfc\x31\x71\x65\x08\x9a\xab\x85\x07\xe2\x48\x40\xbd\xd0\x8f\x61\x22\x6b\x68\xaf\x9b\x4e\x97\x08\x35\xeb\x46\xdd\x82\x41\xc0\x71\x4a\xd3\xa8\xd5\x31\x1e\xe3\xaf\xec\xa6\x81\xa2\xa2\x06\x8e\x59\x9a\x25\xad\x96\x0d\xa8\xe0\x30\x42\x43\xc5\x1d\x58\x80\xbe\xbe\xfc\xf6\xf4\x33\x13\xae\x3a\x1d\xfa\xf7\x49\x6e\xb3\x1b\x82\x88\xad\xed\xe2\xb6\xaf\xdb\x40\xbd\x97\x75\x92\x6b\x1e\xa0\x08\xbb\x80\x4a\x1e\xe6\xa0\xdd\xff\x4d\x12\xc7\x91\x30\x89\x79\xd6\x10\x2c\xe0\x39\x79\x6d\x00\xb9\x41\xae\x7b\x59\x42\x9f\x2c\x92\x2f\x61\x59\x6e\x8a\x32\xa5\xff\xae\x51\x87\x02\xd8\xf7\x77\x50\x92\x5d\x7c\xf9\x77\xec\x09\xee\xf1\xe5\x76\x2b\xf2\x3d\x8d\x89\x9c\x29\xa9\xdb\x1e\x4f\xab\xd2\x65\x6d\x3c\x58\x30\xd3\xfb\xd4\xe3\xc9\x34\x05\x81\xf7\x9a\x0a\x9e\x23\xb4\x93\x1a\xbf\x44\x99\xae\x00\xb7\xdb\x3a\x86\xe5\x8c\x6d\x9d\x1b\x02\x3c\x63\xe8\xb1\xe9\xa9\x2a\xcd\x29\x25\x4e\x75\xe4\x18\xa2\x52\x42\xa7\xe1\x4c\xf5\x14\x8c\xc4\x86\x20\x80\xaa\x34\x8c\x9a\xad\x63\xc8\xa9\x2d\x3a\x7a\x77\xfe\xeb\xb3\x1d\x03\xc2\x8c\xd7\xb2\x5b\x8b\x25\x02\x08\xaf\xaf\x7f\x54\xf0\x99\x45\xc1\xe0\x39\x8a\xa5\x6e\xd7\xf2\x5d\x72\xc9\x05\xbb\x65\xfb\x9c\xb4\xe1\x34\xa5\xb2\x84\x32\x07\x75\x1d\xe5\x31\xfa\x20\xc6\xbf\x9c\xe1\x52\x69\xeb\xba\x0d\xcb\x73\x41\xe5\x04\x69\x8a\xdf\x7d\xcc\xb4\xcb\x5c\x8a\xf0\xf5\x2a\xf3\xe7\x63\xca\x90\x77\xce\xa7\x6c\xfa\x97\xe5\x8f\x19\x9f\x61\xed\x84\xad\x2d\x89\xed\x9e\xed\x15\xba\xbf\xb2\x09\x03\xba\xce\x8e\xde\xca\x3e\x97\xdf\xcb\xb8\x12\x8d\x9c\x8a\x7d\x06\x11\xcc\xd4\xb8\x57\x28\x7d\x81\x63\x51\x23\x57\x51\x1a\x76\x0f\xb7\x87\x7f\x02\x00\x00\xff\xff\x8e\x83\x51\xc2\xc1\x05\x00\x00")

func localizerLocalesEnJsonBytes() ([]byte, error) {
	return bindataRead(
		_localizerLocalesEnJson,
		"localizer/locales/en.json",
	)
}

func localizerLocalesEnJson() (*asset, error) {
	bytes, err := localizerLocalesEnJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localizer/locales/en.json", size: 1473, mode: os.FileMode(420), modTime: time.Unix(1664208268, 0)}
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
	"localizer/locales/en.json": localizerLocalesEnJson,
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
	"localizer": &bintree{nil, map[string]*bintree{
		"locales": &bintree{nil, map[string]*bintree{
			"en.json": &bintree{localizerLocalesEnJson, map[string]*bintree{}},
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
