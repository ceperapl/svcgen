// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/Dockerfile.gotmpl
// templates/README.md.gotmpl
// templates/Taskfile.yml.gotmpl
// templates/docker-compose.yml.gotmpl
// templates/gitignore.gotmpl
// templates/go.mod.gotmpl
// templates/hello.proto.gotmpl
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

var _templatesTaskfileYmlGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\x51\x6b\xdb\x3c\x14\x7d\xf7\xaf\x38\x18\x3e\x0a\x1f\xc8\x1e\xec\xcd\x90\x07\x37\xd5\x92\x40\xe7\x14\x27\x74\x1d\x63\x38\x8a\x2d\xab\xa2\xb6\x64\x24\xc5\xa5\x0b\xf9\xef\xc3\x96\xbd\x76\xa5\x7b\xb3\xb9\xf7\xdc\x73\xee\x39\x57\x3d\x37\x56\x6a\x95\xe0\xea\xf3\x55\x10\xf4\xcc\xd8\x24\x00\xd6\xe9\x6e\x9d\xe0\x0c\xfb\x98\x40\x48\x07\xc3\x7b\xd2\x31\x63\x39\x08\xb1\x8f\xda\x38\xac\x69\x7a\x83\x4b\x00\x5c\xe7\x69\xb6\xfc\x67\x33\x3b\x1e\x87\x5f\xc3\xeb\x57\xc4\x8e\xe6\xf7\x9b\x25\x2d\xb2\xf4\x2b\x4d\x60\xb9\xe9\x65\xc9\x87\x49\x9b\x2c\xcd\xbf\x17\x77\xe9\x7e\x9d\x20\x8a\x8f\x52\xc5\xe7\xf3\xe1\x7c\x8e\xde\x02\x2e\x97\xc3\xc5\x0f\x49\xf3\xe5\xba\xb8\xcb\xb7\xfb\xed\x97\xcd\x2d\xdd\xcd\x0a\x6a\xa9\x2a\x44\x31\xeb\x64\xdc\x19\xed\xf4\xf1\x54\x13\xdb\xf1\x12\xc4\xbd\x74\x1c\x35\x88\x62\x2d\x47\xf8\x7f\x34\x96\x43\x90\xce\x48\xe5\x6a\x84\xff\xd5\x08\x47\x81\xf9\x76\xbb\x9f\x64\xf8\x99\xdd\x73\x85\x4b\x10\x38\x66\x9f\x6c\x12\x04\x80\xe0\x8a\x1b\xe6\xf8\xe0\x15\xc0\x55\xef\x3f\x80\xdd\xb7\x74\xb5\xa2\x79\xb1\xa2\x19\xcd\xd3\x3d\x2d\xe8\xc3\x9e\x66\xbb\xcd\x36\x4b\x10\xd6\xac\xb1\x3c\x1c\x3b\xcb\xb6\xb2\x33\x86\xa0\xd2\xe5\x13\x37\x30\x27\x05\x42\x4c\x0b\xd2\xc3\x6f\xfe\x47\xc9\xb8\x76\xe2\x37\xc2\x2f\xd5\xbc\xf8\xcf\x12\x84\x08\x5d\xe8\x93\x5b\x74\xcd\x49\x48\x65\x17\xc2\x74\xe5\xd4\x18\x77\x4f\x22\x76\x86\x29\xdb\x69\xe3\xe2\xa1\x32\xe4\xf7\xcc\x84\xe0\x66\x04\x35\x5a\x38\x6d\x5d\xc5\x8d\x59\x38\x73\xe2\x33\x70\x70\x6f\xea\x9b\xcc\xdb\x2c\xde\x94\xfe\x36\x76\x0e\xe9\x5d\x20\x3e\xa9\x00\x60\x8a\x35\x2f\x56\x4e\xfb\x56\xd2\x0c\xe9\x7e\xe0\x82\xd0\x0d\x53\xa2\x94\xa4\x91\xca\x79\x33\x7a\x10\xa2\xf8\x33\xa9\x8d\x6e\x89\xe1\x3d\x0e\xc3\x81\xb5\xdc\x08\x4e\x8e\xcc\x72\xb4\xcc\x3a\x6e\x26\x09\xfe\x14\x47\xde\x03\xa2\x38\x8a\xa2\x81\xfe\x78\x92\x4d\x35\x71\xf3\xce\x26\xf8\x31\xa7\xf7\xf3\x03\x11\xcb\xd5\xb6\xa0\x59\x7a\x7d\x4b\x6f\x16\x9f\x20\xb4\x87\x83\x34\x55\xdd\x30\x61\x11\x92\x07\xb4\x4c\xaa\xe8\xde\x3f\x9c\x85\xa7\x1e\x1e\xcc\x48\x8c\xb9\xbe\x66\xf6\xf1\x7d\x31\x04\x61\x20\x7a\x96\xfb\x7a\xef\x1e\x1a\x05\xbf\x03\x00\x00\xff\xff\xa5\x1b\xf9\xd9\x8e\x03\x00\x00")

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

	info := bindataFileInfo{name: "templates/Taskfile.yml.gotmpl", size: 910, mode: os.FileMode(436), modTime: time.Unix(1665489015, 0)}
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
