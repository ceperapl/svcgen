// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/Dockerfile.gotmpl
// templates/README.md.gotmpl
// templates/Taskfile.yml.gotmpl
// templates/docker-compose.yml.gotmpl
// templates/gitignore.gotmpl
// templates/go.mod.gotmpl
// templates/hello.proto.gotmpl
// templates/placeholder.gotmpl
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

var _templatesDockerfileGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcf\xc1\x4b\xc3\x30\x14\xc7\xf1\x7b\xfe\x8a\xc7\x76\x4e\xe2\x60\xa7\x49\x0f\x73\xab\x63\xa8\xcd\xe8\x26\x32\x44\xe4\xb5\x7b\x66\xc1\x34\x29\x49\x2a\xee\xbf\x97\xd9\x82\x03\x11\xaf\x0f\xde\x8f\xcf\x77\x0c\x55\x67\xec\x01\xd2\x91\x20\x52\xf8\xa0\x00\x95\x71\x18\x4e\xec\xb6\x54\x0f\xa0\xbd\x45\xa7\x67\x13\x31\x99\x8a\x29\x47\xdb\x1a\x47\x30\xdf\x42\x85\x91\xf8\xf7\x27\x05\xf6\xa4\xca\xbb\xe5\xba\x04\xa9\xbd\x8c\xa1\x96\xda\xa4\x63\x57\x89\xda\x37\xb2\xa6\x96\x02\xb6\x56\x26\x8c\xef\x91\x2d\xd4\x66\x0f\x02\x04\x2b\x1f\x0b\x58\xac\xd4\x6b\x5e\xcc\x6f\xee\xf3\x65\x76\x05\x2b\xa5\xb6\x99\x35\xae\xfb\x04\xed\x07\x14\xf7\x67\x8c\x1c\x5c\x3c\xa1\x8e\x10\xc8\x12\x46\x02\x21\xeb\xe6\xc0\xd8\x18\x6a\xdf\x9e\x7e\xf3\xe1\x2d\xf8\x06\x06\x21\xc4\x84\x9a\xae\x21\x74\xee\xaf\xd0\x3e\x6d\x66\x31\x51\x4c\x3f\x45\x95\x71\x3d\x9a\xf3\xf3\x60\x76\xd9\xfd\x7f\xaf\xbc\xd0\x0b\x96\x17\xbb\x72\xbf\x51\xeb\x62\x07\xcf\xa3\xfe\x3a\x7a\x61\x5f\x01\x00\x00\xff\xff\x97\x75\x3c\xcd\x81\x01\x00\x00")

func templatesDockerfileGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDockerfileGotmpl,
		"templates/Dockerfile.gotmpl",
	)
}

func templatesDockerfileGotmpl() (*asset, error) {
	bytes, err := templatesDockerfileGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/Dockerfile.gotmpl", size: 385, mode: os.FileMode(436), modTime: time.Unix(1665498257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesReadmeMdGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templatesReadmeMdGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesReadmeMdGotmpl,
		"templates/README.md.gotmpl",
	)
}

func templatesReadmeMdGotmpl() (*asset, error) {
	bytes, err := templatesReadmeMdGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/README.md.gotmpl", size: 0, mode: os.FileMode(436), modTime: time.Unix(1665498257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTaskfileYmlGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x5b\x6b\xdb\x30\x18\x7d\xf7\xaf\x38\x18\x46\x61\x20\x7b\xb0\x37\x43\x1e\xdc\x4c\x4b\x03\x9d\x5d\x9c\xb0\x0b\x63\xb8\x8a\x2d\x6b\x22\x8e\x24\x24\xd9\x25\x0b\xfe\xef\xc3\x97\x8c\xac\xf4\x4d\xf0\x9d\x73\x38\x17\xf5\xdc\x3a\xa9\x55\x82\xbb\x8f\x77\x41\xd0\x33\xeb\x92\x00\xd8\xd1\xe2\xeb\x76\x4d\xcb\x2c\xfd\x42\x13\x38\x6e\x7b\x59\xf1\x00\xb8\xdf\x66\x69\xf1\xa3\x7c\x4a\xf7\x0f\x09\xa2\xf8\x20\x55\x7c\xb9\x3c\x5f\x2e\xd1\x2d\x61\x18\x9e\x87\x61\x12\x49\x8b\xf5\x43\xf9\x54\xe4\xfb\xfc\xf3\xf6\x91\xee\x12\x5c\xe0\x7e\x27\x68\xa4\xaa\x11\xc5\xcc\xc8\xd8\x58\xed\xf5\xa1\x6b\x88\x33\xbc\x02\xf1\x67\xc3\xd1\x80\x28\x76\xe2\x08\xdf\x47\xd3\x39\x04\x31\x56\x2a\xdf\x20\x7c\xd7\x20\xc4\xa8\x5d\xe4\xf9\x7e\xb1\x31\x6b\x9a\x97\x1a\x43\x10\x78\xe6\x8e\x2e\x09\x02\x40\x70\xc5\x2d\xf3\x7c\x8c\x03\x70\xd5\xcf\x0f\x60\xf7\x2d\xdd\x6c\x68\x51\x6e\x68\x46\x8b\x74\x4f\x4b\xfa\x7d\x4f\xb3\xdd\x36\xcf\x12\x84\x0d\x6b\x1d\x0f\x27\x64\x75\xaa\xdd\x95\x43\x50\xeb\xea\xc8\x2d\x6c\xa7\x40\x88\x3d\x81\xf4\x98\x93\xff\x73\x32\xc5\x4e\xe6\x44\xf8\xa3\xda\xf3\xfc\xac\x40\x88\xd0\xa5\xee\xfc\xca\xb4\x9d\x90\xca\xad\x84\x35\xd5\x02\x8c\xcd\x51\xc4\xde\x32\xe5\x8c\xb6\x3e\x1e\x2f\x20\xc4\xbd\x30\x21\xb8\x9d\x48\xad\x16\x5e\x3b\x5f\x73\x6b\x57\xde\x76\xfc\x4a\x1c\xdb\x5b\x70\x4b\x79\xdb\xd5\xcd\xe9\xff\x62\xaf\x23\xbd\x1a\x64\x5e\x2a\x00\x98\x62\xed\xd9\xc9\x25\x6f\x2d\xed\xb8\xee\x1b\x2d\x08\xdd\x32\x25\x2a\x49\x5a\xa9\xfc\x5c\x46\x8f\x28\x8e\xa2\x68\x54\x39\x74\xb2\xad\x17\x09\x6e\x5c\x82\x9f\xd7\x11\x7e\xbd\xa1\xb5\xde\xe4\x25\xcd\xd2\xfb\x47\xfa\x69\xf5\x01\x42\xcf\x74\x10\x06\xa2\x17\xc3\x37\xdf\x6d\xb2\x8a\x28\xf8\x1b\x00\x00\xff\xff\x4e\x43\x3a\x3b\xb0\x02\x00\x00")

func templatesTaskfileYmlGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesTaskfileYmlGotmpl,
		"templates/Taskfile.yml.gotmpl",
	)
}

func templatesTaskfileYmlGotmpl() (*asset, error) {
	bytes, err := templatesTaskfileYmlGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/Taskfile.yml.gotmpl", size: 688, mode: os.FileMode(436), modTime: time.Unix(1665654554, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDockerComposeYmlGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templatesDockerComposeYmlGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDockerComposeYmlGotmpl,
		"templates/docker-compose.yml.gotmpl",
	)
}

func templatesDockerComposeYmlGotmpl() (*asset, error) {
	bytes, err := templatesDockerComposeYmlGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/docker-compose.yml.gotmpl", size: 0, mode: os.FileMode(436), modTime: time.Unix(1665498257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGitignoreGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd2\xd2\x4f\xca\xcc\xe3\x02\x04\x00\x00\xff\xff\x39\x47\x85\xde\x07\x00\x00\x00")

func templatesGitignoreGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesGitignoreGotmpl,
		"templates/gitignore.gotmpl",
	)
}

func templatesGitignoreGotmpl() (*asset, error) {
	bytes, err := templatesGitignoreGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/gitignore.gotmpl", size: 7, mode: os.FileMode(436), modTime: time.Unix(1665498257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGoModGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xcb\x92\x9b\x30\x10\x3c\x87\xaf\xe0\x98\x3d\x58\x2f\x8c\xb1\x3f\x47\xc8\x83\x50\x2c\x69\x14\x21\xd8\xf5\xdf\xa7\x24\xaa\x52\x0e\x36\xeb\xdc\xa8\xea\x6e\x4d\x4f\x4f\xe3\xf0\x3a\x5b\xa8\xb5\x49\xe3\xdc\x13\x85\x8e\x2a\x74\x41\xfa\x3b\xed\xad\xf4\xb7\x69\x51\x55\xa5\xb1\xe6\x84\x9f\xab\x2a\xc2\xef\xd9\x44\xa8\x7f\x56\x3f\x1e\x04\x1a\x0f\x37\x93\xe8\xcd\xa4\x7a\x61\x84\x0b\xc2\x5e\xc1\x16\x75\x86\x9f\x51\x2b\xbd\xa6\x21\x62\xc2\x7e\x1e\xea\x85\x93\x96\x88\x0d\x25\x1a\x6b\x25\x75\xf3\x57\x86\xcf\x9b\x17\xa6\x30\xf0\x86\x86\xc1\x4a\x9d\x61\x46\xda\x17\xf0\x62\x02\xc4\x0c\xf3\xa6\xc8\xcb\x54\x82\x51\xd3\x2f\xea\xa1\x18\x67\x84\x1d\x04\x13\x82\xb5\x82\x31\xc6\x2e\xcd\xf9\x20\xa0\x81\xbe\xeb\x2f\xc7\x56\x65\x4b\x88\xda\x02\x79\xd0\x6a\xf0\xc5\xf8\xbf\x72\x7e\xe1\x6d\x73\x6a\xc5\xa1\x91\xc7\xee\x0a\x1d\x74\x97\xfe\xfa\x52\x1e\x83\xca\x9e\x8e\x27\x22\xaa\x8f\xbd\x78\x87\xc9\x63\x32\xc3\xfd\xef\xc7\x9a\xd1\xb1\xa6\xb4\x36\xfe\x6a\x22\xa8\xb4\x0d\xdc\xa2\x1e\x5c\xc9\x7c\x70\x65\xb9\x96\xf0\x5d\xfe\x28\xa7\xd1\x28\x8c\x81\x8e\xca\xae\x11\xb2\x5d\xb2\x93\xda\x28\xf4\xd2\xc4\x7c\xb3\x00\x31\x19\x98\xd6\xb3\x9c\xf6\x45\x26\xa9\x11\xac\x1d\xa9\x93\x61\x4a\x71\x56\x69\x8e\xb0\xee\xb1\x3f\x2a\x80\xb5\x90\x0c\xc4\xbc\x51\x42\x57\xbc\x5d\x48\xfb\xff\x02\xba\x88\x7a\x11\xb9\x12\xbb\x9a\xb5\x1e\x72\x80\x88\xeb\x16\xe2\x0d\x55\xc9\x29\xbd\x73\xbe\x32\x7f\x7d\x4a\x9b\x20\x7e\x82\x4c\x23\x44\x27\x7d\xe9\xdf\x77\xb2\xb9\xc7\xc9\x24\xa4\x1a\x13\xf8\xa5\x54\xe3\xe9\x6e\x8f\xcd\x9d\xee\xd3\xb6\xb9\xbc\xe5\x0d\x13\x87\x5e\x09\x75\x6e\xe5\x55\x72\x26\xbf\x79\x20\xc1\x57\xa9\x47\x43\xba\x2d\x6b\x5b\xd6\xc7\x3f\x54\x9c\x9f\xb6\xc0\x70\xd3\xc4\x78\x6a\xbc\x21\x0b\xcf\xa4\x53\xb7\x4b\xba\x4b\x67\xc9\x7a\x9b\xe3\x1b\x52\x53\x2f\x0d\x61\x9b\x14\x3e\xaa\x3f\x01\x00\x00\xff\xff\x38\x71\x16\xd6\xb7\x04\x00\x00")

func templatesGoModGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesGoModGotmpl,
		"templates/go.mod.gotmpl",
	)
}

func templatesGoModGotmpl() (*asset, error) {
	bytes, err := templatesGoModGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/go.mod.gotmpl", size: 1207, mode: os.FileMode(436), modTime: time.Unix(1665501384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHelloProtoGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xb1\x4a\xc0\x40\x0c\x40\x67\xef\x2b\xc2\x4d\xed\xd2\x52\x1c\x4b\x77\x67\x7f\x40\x8e\x23\x1c\x87\x6d\x12\x93\xb4\x28\xc5\x7f\x97\x53\x0b\x16\xb7\x84\xbc\xf7\x62\x1f\xe4\xe9\x1d\x16\x88\xa2\xec\xfc\x18\xe7\x10\x24\xe5\xd7\x54\x10\x90\xbc\x7a\x45\x1b\x8e\x69\x0e\x81\xc5\x2b\x13\x14\x7e\xb9\xee\x0b\xc4\xa2\x92\x9b\x52\x37\x61\x75\x88\x85\xb9\xac\x38\x26\xa9\x63\x22\x62\x4f\xcd\xb1\xe1\xbb\xdd\xb8\x0d\xcd\x9a\xfa\x84\xeb\xca\xcf\xf8\xb6\xa3\x39\x9c\x01\x00\xc0\x5c\x2b\x15\xa0\xb4\xb5\xf2\x34\x87\xcf\x7f\xb8\x09\x93\xe1\x9d\x2f\x8a\xe8\x6d\xb8\x1c\x43\x3d\x6a\xfe\x75\xe0\x0c\x0f\x2a\xf9\x67\xe9\xfe\x7e\xed\x41\xd1\x77\x25\x83\xee\x56\xef\x5b\xe4\x2b\x00\x00\xff\xff\x14\x08\xd5\xa1\x16\x01\x00\x00")

func templatesHelloProtoGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHelloProtoGotmpl,
		"templates/hello.proto.gotmpl",
	)
}

func templatesHelloProtoGotmpl() (*asset, error) {
	bytes, err := templatesHelloProtoGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/hello.proto.gotmpl", size: 278, mode: os.FileMode(436), modTime: time.Unix(1665497527, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPlaceholderGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templatesPlaceholderGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPlaceholderGotmpl,
		"templates/placeholder.gotmpl",
	)
}

func templatesPlaceholderGotmpl() (*asset, error) {
	bytes, err := templatesPlaceholderGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/placeholder.gotmpl", size: 0, mode: os.FileMode(436), modTime: time.Unix(1665617933, 0)}
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
	"templates/Dockerfile.gotmpl":         templatesDockerfileGotmpl,
	"templates/README.md.gotmpl":          templatesReadmeMdGotmpl,
	"templates/Taskfile.yml.gotmpl":       templatesTaskfileYmlGotmpl,
	"templates/docker-compose.yml.gotmpl": templatesDockerComposeYmlGotmpl,
	"templates/gitignore.gotmpl":          templatesGitignoreGotmpl,
	"templates/go.mod.gotmpl":             templatesGoModGotmpl,
	"templates/hello.proto.gotmpl":        templatesHelloProtoGotmpl,
	"templates/placeholder.gotmpl":        templatesPlaceholderGotmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"Dockerfile.gotmpl":         &bintree{templatesDockerfileGotmpl, map[string]*bintree{}},
		"README.md.gotmpl":          &bintree{templatesReadmeMdGotmpl, map[string]*bintree{}},
		"Taskfile.yml.gotmpl":       &bintree{templatesTaskfileYmlGotmpl, map[string]*bintree{}},
		"docker-compose.yml.gotmpl": &bintree{templatesDockerComposeYmlGotmpl, map[string]*bintree{}},
		"gitignore.gotmpl":          &bintree{templatesGitignoreGotmpl, map[string]*bintree{}},
		"go.mod.gotmpl":             &bintree{templatesGoModGotmpl, map[string]*bintree{}},
		"hello.proto.gotmpl":        &bintree{templatesHelloProtoGotmpl, map[string]*bintree{}},
		"placeholder.gotmpl":        &bintree{templatesPlaceholderGotmpl, map[string]*bintree{}},
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
