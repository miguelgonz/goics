// Code generated by go-bindata.
// sources:
// templates/episode.html
// DO NOT EDIT!

package templates

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

var _episodeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x4b\x8e\x9c\x30\x10\x86\xf7\x39\x85\xe5\xbd\xe1\x02\x30\x8b\x28\x59\x8c\x94\xc9\x66\x72\x01\xd3\x54\xd3\xa5\xd8\xc6\x72\x99\x96\x5a\x88\xbb\xc7\x36\xd0\xcd\x53\x19\x56\xc6\xf5\xff\xf5\xf8\xca\x85\x64\x37\x07\xd7\x92\xe7\xd6\xb5\x8d\x93\x5a\x03\xe5\x7d\x9f\xbd\xff\x18\x06\xce\x3c\x7a\x05\x25\x0f\xff\x7f\xe2\x29\x5e\x5d\x94\x24\x2a\xb9\x42\xf2\x02\x3d\x68\xa1\xd0\xfc\x65\xe4\xa5\xe7\x6f\xdf\x58\xf8\xfa\x1e\xaf\x2c\xfb\xdd\xe9\x0a\xdc\x30\xa4\xab\xf8\x15\x35\xde\x67\xb3\x49\x41\xce\xa4\x43\x29\x6e\x58\xd7\x60\x4a\xee\x5d\x07\xfc\x2d\x94\x9a\xad\x45\x1e\x2c\x73\x4e\x30\xf5\x94\x6c\x99\x28\xf5\xc7\x7c\x6b\x45\x3a\x25\xfb\xd4\xe9\xe4\xde\x59\xac\x43\x2d\xdd\x63\x6a\x76\x1b\x75\x22\x44\x1b\x58\x44\x93\x02\x75\x33\x2b\xc2\x91\x33\x72\x97\x44\xe5\x3d\x8a\x23\x15\xa9\xfc\x1a\x53\xbe\xc8\xff\x9a\x63\x5b\xae\xbd\x83\x53\x32\x34\xb3\xd1\x8c\x0c\x7f\xc9\x0a\x14\x2d\x18\x6e\xed\x84\x8d\xb1\x2d\x79\x06\x35\xfa\x36\xc0\x54\x89\xc0\xe8\xcb\x7e\xce\x97\x2b\x94\x5b\x9c\x3b\x3e\xa1\x1f\x81\x97\xd6\x84\xa6\xf0\x60\x43\x4f\x0e\x41\xc2\xfc\x1d\xa3\x1e\xdc\xfe\x4f\x4c\xe7\x38\x1b\x2e\xe7\x3b\x5b\x0c\x41\x70\xd5\xe7\xab\x39\xdb\xf0\x9a\xd9\x67\x57\xf9\x31\x7c\x4e\x6d\x92\xa4\x5c\x2f\xfd\xff\x18\xd9\xa7\xff\x61\x5a\x4b\x48\xa3\x3f\xfd\x00\x65\x1f\x61\x07\x9d\x8e\x69\xec\xa2\x7f\xbb\x7e\x48\x63\x8f\x1f\x92\x3c\xb8\xef\x4e\xae\x0a\x3c\x3d\x64\xa5\x99\x6b\xe9\x24\x15\x55\xd4\x6e\x1e\xe5\xa1\x9e\xb4\x54\xe3\x23\x58\x14\x19\x81\x51\xf6\x19\x83\xb1\xc3\xe8\xf8\x42\x32\x9d\x46\x3a\xcb\xf6\x1a\xf8\x38\xdd\xd1\xfd\x0e\xea\x0c\x6b\x82\x5f\xe4\x32\xbc\x8c\x7f\x01\x00\x00\xff\xff\x1d\xb3\x37\x4c\x94\x04\x00\x00")

func episodeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_episodeHtml,
		"episode.html",
	)
}

func episodeHtml() (*asset, error) {
	bytes, err := episodeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "episode.html", size: 1172, mode: os.FileMode(420), modTime: time.Unix(1454080408, 0)}
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
	"episode.html": episodeHtml,
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
	"episode.html": &bintree{episodeHtml, map[string]*bintree{}},
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

