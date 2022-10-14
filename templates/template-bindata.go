// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/Dockerfile.gotmpl
// templates/README.md.gotmpl
// templates/Taskfile.yml.gotmpl
// templates/buf.gen.yaml.gotmpl
// templates/buf.yaml.gotmpl
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

var _templatesDockerfileGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcf\xc1\x6a\x02\x31\x14\x85\xe1\x7d\x9e\xe2\xa0\xeb\x24\x15\x5c\x59\x5c\x58\xb5\x22\xad\x13\x19\x2d\x45\x4a\x29\x19\xe7\x36\x0e\xc4\x44\x92\x4c\xa9\x88\xef\x5e\xec\x0c\x54\x28\x6e\x2f\xdc\xc3\xf7\x77\x51\xd4\x95\x2d\x91\x76\x84\x48\xe1\x8b\x02\x8a\xca\xe9\x70\x64\x8f\xb9\x5a\xc0\x78\xab\x9d\x19\xf4\x44\xaf\x2f\xfa\x5c\xdb\x43\xe5\x08\xa3\x15\x0a\x1d\x89\xff\x7e\x52\x60\xaf\x2a\x7f\x9a\xcc\x73\x48\xe3\x65\x0c\x5b\x79\x3a\x41\x2c\x7c\x59\x5b\x5a\xea\xb4\x3b\x9f\xd9\x58\x2d\x37\x10\x10\x2c\x7f\xc9\x30\x9e\xa9\x8f\x69\x36\x7a\x78\x9e\x4e\x86\x77\x98\x29\xb5\x1a\xda\xca\xd5\xdf\x30\xbe\xb5\x70\x7f\x31\xc8\x96\xc3\x93\x36\x11\x81\x2c\xe9\x48\x10\x72\xbb\x2f\x19\xeb\x62\xeb\x0f\xc7\xff\x6a\x7c\x06\xbf\x47\x0b\x43\x4c\xda\xd0\x3d\x42\xed\x6e\xf5\x35\x45\x03\xab\x13\xc5\xf4\x17\x52\x54\xae\x41\x73\x7e\x19\x1c\x5e\xe7\xde\xcc\x94\x57\x68\xc1\xa6\xd9\x3a\xdf\x2c\xd5\x3c\x5b\xe3\xad\xd3\x5c\x3b\xef\xec\x27\x00\x00\xff\xff\xed\x05\xf4\x39\x6f\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/Dockerfile.gotmpl", size: 367, mode: os.FileMode(436), modTime: time.Unix(1665748649, 0)}
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

var _templatesTaskfileYmlGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x4b\xc4\x30\x10\x85\xef\xf3\x2b\xde\x6d\x4f\x49\x04\x6f\x01\x0f\xdd\xb5\xe8\x82\x56\x59\x45\x10\x91\x6e\xda\x66\x4b\xb0\x9b\x94\xa4\x2d\x48\xc8\x7f\x97\xda\x55\xf6\xb0\xb7\x81\xe1\xfb\xde\xbc\x99\xb4\x0f\xc6\x59\x89\xd5\xf5\x8a\x68\x52\x3e\x48\x02\x5e\xf2\xdd\xdb\x76\x93\x97\x45\xf6\x98\x4b\xc4\x08\x5e\xa8\xa3\x4e\x89\x80\xf5\xb6\xc8\x76\xef\xe5\x73\xf6\x7a\x2f\xc1\x45\x65\xac\x88\x71\x1f\x23\x3f\x67\x52\xda\xa7\x44\x34\xa8\xf0\x15\x24\x11\xd0\x6a\xab\xbd\x1a\xf4\x2c\x07\xea\x63\x13\x96\x09\x60\xa8\xc6\xc3\xff\x1e\x5c\xa8\xde\x88\xde\xbb\xc1\x55\xe3\x81\x85\x5e\xd7\x82\x00\x65\x55\xf7\x1d\xcc\x89\x6a\x8c\x9f\xb3\x2f\xb8\x5a\xd7\x29\xdb\xd6\x86\x75\xc6\x0e\xf0\xa3\x05\x9b\xc0\x05\xe7\x7c\xbe\xa2\x1a\x4d\xd7\x9c\x14\xba\x0f\x12\x1f\x7f\xb9\x9f\x17\x5c\x9b\xbb\xa7\x32\x2f\xb2\xf5\x43\x7e\x7b\x73\x85\xd6\x2d\x38\x98\x02\x73\x58\x3a\x9f\x3d\xe3\xb7\x32\x38\xfd\x04\x00\x00\xff\xff\x0f\xae\x76\x70\x51\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/Taskfile.yml.gotmpl", size: 337, mode: os.FileMode(436), modTime: time.Unix(1665748739, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBufGenYamlGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\x31\x6a\x04\x31\x0c\x05\xd0\xde\xa7\xd0\x05\x36\x43\x5a\xc3\x1e\x21\x90\x1b\x18\x65\xfc\xc7\x6b\xe2\xb1\x1c\x59\x1e\x02\xcb\xde\x3d\x78\x52\x6d\x11\xd2\x09\xe9\xff\x57\xe8\x80\xf6\x2c\xd5\xd3\xf1\xea\x76\xae\x9c\x10\xbd\x23\x42\xe5\x8f\x82\xe8\xc9\x74\xc0\x11\x25\x09\x8d\xd7\x4f\x4e\x08\x4d\xb1\xe5\xef\x19\x22\x8a\xd8\x78\x14\xf3\x74\xbf\xd3\xcb\x9b\xc4\x51\xf0\xce\x76\x7b\x3c\x96\x84\xba\x34\x15\x93\x25\x89\x6b\x65\xa4\x5c\xfb\xec\x5c\xa8\xf2\x0e\x4f\x49\x4e\x40\x86\x79\x7a\xca\x9e\xdb\x66\x9e\x1a\xdb\xad\x5f\xbb\x0c\x5d\x11\x14\x85\x2d\x1f\x78\x12\x2e\x49\xdb\xfa\x0f\x73\x0e\xb3\xf4\x17\xf7\x7b\x55\x7c\x8d\xac\x08\xa3\xe6\xbd\x15\xec\xa8\x86\x18\x3a\x74\xfe\xe7\xba\x71\xe9\x70\x3f\x01\x00\x00\xff\xff\x6c\x33\xd5\xcb\x2b\x01\x00\x00")

func templatesBufGenYamlGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBufGenYamlGotmpl,
		"templates/buf.gen.yaml.gotmpl",
	)
}

func templatesBufGenYamlGotmpl() (*asset, error) {
	bytes, err := templatesBufGenYamlGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/buf.gen.yaml.gotmpl", size: 299, mode: os.FileMode(436), modTime: time.Unix(1665748572, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBufYamlGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4b\x2d\x2a\xce\xcc\xcf\xb3\x52\x28\x33\xe4\x4a\x2a\x4a\x4d\xcc\xce\xcc\x4b\xb7\xe2\x52\x50\x28\x2d\x4e\x05\x51\x0a\x0a\xba\x0a\x6e\x9e\x3e\xae\x5c\x39\x99\x79\x25\x68\xe2\x2e\xae\x6e\x8e\xa1\x3e\x21\x5c\x80\x00\x00\x00\xff\xff\x0a\x9b\x1f\xa4\x43\x00\x00\x00")

func templatesBufYamlGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBufYamlGotmpl,
		"templates/buf.yaml.gotmpl",
	)
}

func templatesBufYamlGotmpl() (*asset, error) {
	bytes, err := templatesBufYamlGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/buf.yaml.gotmpl", size: 67, mode: os.FileMode(436), modTime: time.Unix(1665745657, 0)}
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

var _templatesGoModGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xdb\x8e\xdb\x20\x14\x7c\xae\xbf\xc2\x8f\xdd\x87\x70\x73\x1c\x27\x1f\x51\xa9\xbf\x80\xc9\x31\xa6\xc1\x1c\x8a\xb1\x37\xd1\x6a\xff\xbd\x02\x4b\x55\xea\xc4\x9b\xbe\x21\xcd\x0c\xcc\x99\x33\x0c\x78\x9e\x2c\x94\x1f\x1f\x25\xf9\x91\x8f\x3f\x65\xec\x3f\x3f\x8b\x42\x63\xc9\x09\x3f\x16\x45\x80\xdf\x93\x09\x50\x7e\x2f\xbe\x69\x13\xfb\xa9\x25\x0a\x07\xaa\x71\x77\x31\x91\x5e\x4c\x2c\x67\x46\xb8\x20\xec\x19\x6c\x51\x27\xf8\x11\xb5\xd2\x69\xea\x03\x46\x6c\xa7\xae\x9c\x39\xa9\x89\x58\x51\x82\xb1\x56\xd2\x61\xba\x26\xf8\xb8\xba\x61\xf4\x1d\xaf\xa8\xef\xac\xd4\x09\x66\xa4\x7e\x02\xcf\xc6\x43\x48\x30\xaf\xb2\x3c\xbf\x4a\x30\x68\x7a\xa5\x0e\xb2\x71\x46\xd8\x4e\x30\x21\x58\x2d\x18\x63\xec\x54\x1d\x77\x02\x2a\x68\x9b\xf6\xb4\xaf\x55\xb2\x84\xa8\x2d\x90\x3b\xad\x06\x97\x8d\xff\x2b\xe7\x27\x5e\x57\x87\x5a\xec\x2a\xb9\x6f\xce\xd0\x40\x73\x6a\xcf\x4f\xe5\xc1\xab\xe4\x69\x7f\x20\xa2\x78\xdb\x8a\xb7\x1b\x1d\x46\xd3\xdd\xfe\x1e\x96\x8c\xf6\x25\xa5\xa5\x71\x67\x13\x40\xc5\x75\xe0\x16\x75\x37\xe4\xcc\xbb\x21\x0f\x57\x13\xbe\xc9\xef\xe5\xd8\x1b\x85\xc1\xd3\x5e\xd9\x25\x42\xb6\x49\x1e\xa4\x36\x0a\x9d\x34\x21\xed\xcc\x43\x88\x06\xc6\x65\x2d\x87\x6d\x91\x89\xaa\x07\x6b\x7b\x3a\x48\x3f\xc6\x30\xa9\x38\x05\x58\xe6\xd8\x7e\xca\x83\xb5\x10\x0d\x84\x34\x51\xc4\x21\x7b\x3b\x91\xfa\xff\x05\x74\x16\xe5\x2c\x52\x25\x36\x35\x4b\x3d\x64\x07\x01\x97\x29\xc4\x0b\xaa\x92\x63\x7c\xe5\x7c\x61\xfe\x7a\x97\x36\x42\x78\x07\x19\x7b\x08\x83\x74\xb9\x7f\x5f\xc9\xa6\x16\x47\x13\x91\x6a\x8c\xe0\xe6\x5c\x8d\x87\xbd\xdd\x37\x77\xbc\x8d\xeb\xe6\xf2\x9a\x57\x4c\xec\x5a\x25\xd4\xb1\x96\x67\xc9\x99\xfc\xe2\x82\x08\xd7\x5c\x8f\x8a\x34\x6b\xd6\xba\xac\xf7\x3f\x54\x1c\x1f\xa6\x40\x7f\xd1\xc4\x38\x6a\x9c\x21\x33\x4f\xa4\x43\xb3\x49\xba\xc9\xc1\x92\x65\x37\xfb\x17\xa4\xaa\x9c\x2b\xc2\x56\x29\xbc\x15\x7f\x02\x00\x00\xff\xff\xa6\x60\x01\xc6\xac\x04\x00\x00")

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

	info := bindataFileInfo{name: "templates/go.mod.gotmpl", size: 1196, mode: os.FileMode(436), modTime: time.Unix(1665748271, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHelloProtoGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xb1\xca\xc2\x40\x10\x84\xeb\x7f\x9f\x62\x48\x95\x34\x3f\x04\xcb\x90\xde\x5a\x9f\xe0\x38\x96\x33\x98\xdc\x9d\xbb\x9b\xa0\x04\xdf\x5d\x0e\x23\x18\xec\x76\x3f\x66\xe6\xd3\x47\x34\x77\x47\x8f\x2a\x4b\xb2\x74\xa8\x3a\xa2\xec\xfc\xd5\x05\xc6\x85\xc7\x31\xfd\x2f\x6d\x47\x34\xb1\x6a\x41\xc7\x82\x4e\x7c\x9b\x59\x0d\x2b\x01\x80\x9a\x0c\x31\x20\xba\x89\xd1\xa3\xed\xe8\xf9\x13\xd7\x9c\xa2\xf2\x3e\x1f\x84\xd9\xca\xf1\xe9\x28\xcb\x32\xf8\xad\x73\xde\x9e\x95\xfe\x24\xfb\x37\xab\xbf\xe5\x0d\x84\x6d\x96\xa8\xa8\x77\x92\xa6\x6c\xbd\x02\x00\x00\xff\xff\xfe\x3a\xe6\x79\xd5\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/hello.proto.gotmpl", size: 213, mode: os.FileMode(436), modTime: time.Unix(1665745643, 0)}
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
	"templates/buf.gen.yaml.gotmpl":       templatesBufGenYamlGotmpl,
	"templates/buf.yaml.gotmpl":           templatesBufYamlGotmpl,
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
		"buf.gen.yaml.gotmpl":       &bintree{templatesBufGenYamlGotmpl, map[string]*bintree{}},
		"buf.yaml.gotmpl":           &bintree{templatesBufYamlGotmpl, map[string]*bintree{}},
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
