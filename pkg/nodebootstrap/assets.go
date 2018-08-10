// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/10-eksclt.al2.conf
// assets/bail.sh
// assets/bootstrap.al2.sh
// assets/kubeconfig.yaml

package nodebootstrap

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __10EkscltAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4f\x4f\xfb\x46\x10\xbd\xfb\x53\xac\x04\x87\x56\x62\x6d\xc8\x2f\xe5\x80\xe4\x43\x44\x02\x42\x4d\x01\x11\x10\x95\xda\x2a\x9a\xec\x8e\xc3\x28\xf6\xac\x35\x3b\x4e\xf8\x23\xbe\x7b\xe5\x38\x46\x01\xaa\xea\x27\x5f\xbc\xfb\x66\xdf\xbe\x79\xf3\xec\x03\x83\xab\xe8\xb4\xb4\xb1\x46\x47\x05\x39\x13\x5f\xa2\x62\xe5\x8d\x97\x50\x5b\x62\xd3\x30\xa9\x29\x82\x98\x55\xb3\xc0\x12\xf5\x68\xbb\x18\x55\xf0\x1a\xd8\x4c\x89\x9b\x67\x33\x30\xbf\x8c\xa6\x83\x5f\x93\xe4\xaf\x19\xca\x9a\x1c\xfe\x93\x1c\x98\x69\x70\x50\x9a\x0a\x15\x3c\x28\x98\x1a\x04\x2a\x54\x94\x78\x66\xee\x26\x97\x57\x37\xd7\x47\x66\xf4\x38\x9b\x8f\x27\x17\xa3\x87\xe9\xfd\xbc\xdb\x4b\x26\xbc\x26\x09\x5c\x21\xeb\x05\x95\x98\x67\xa8\x2e\xeb\x24\x66\x3d\x57\x8a\xbc\x4e\x0e\xcc\x65\x19\x16\x50\x1a\x60\x6f\xa2\x82\x92\xfb\x74\xc7\x1f\xa3\x3f\xe7\xb7\x37\xe3\xd9\x91\x39\x9f\x3e\xcc\xee\x27\x77\xf3\xf1\xf5\xec\x7f\xe9\x77\xfd\xed\xd8\x3b\xf9\x1c\xd8\xfe\x07\xf9\xf5\xcd\x78\x32\xbf\xba\xfd\x29\xba\xb2\x25\xda\x92\x26\x93\x67\x74\x33\x05\xd1\x7c\xef\x35\x6b\xa2\x64\x0b\xe2\xfe\x80\xf9\x3b\x31\xc6\x5a\xf0\x5e\x30\xc6\xfc\x38\xdd\x3e\xbb\x5d\x0e\x1e\x2d\xd5\xf9\xe1\xdb\x4e\xc2\xfb\x0e\x70\x65\x13\x15\xc5\x7a\x8e\xf9\xe1\xdb\x5e\xcb\x7d\x41\x05\xcf\xb6\x0e\xbe\x45\x7b\x6b\x7a\x08\x1a\x7d\x42\x56\x72\xa0\x14\xd8\x6a\x58\x21\xdb\x0d\x2e\x9e\x42\x58\xed\x95\x04\xa1\xd7\xae\xa2\x0a\x1e\xf3\xc7\xcf\x05\x65\x19\x36\xb6\x16\x5a\x53\x89\x4b\xf4\xb9\x4a\x83\x3b\xac\x0e\xde\x12\x17\x02\xd6\x05\x56\x20\x46\xb1\x54\xc1\x12\xf3\xd3\xe3\xc1\xf0\xf8\xe4\x64\xf8\x63\xf8\xdb\x20\xf5\x2b\x49\xd1\x49\x7a\xf8\xf6\x3d\x17\xef\x29\x6c\x03\x07\x9b\x98\xba\x50\xb5\x1e\x67\x35\x34\x11\x2d\x54\xfe\x74\x78\xf6\x23\x3d\xf9\x30\x22\x34\xde\xd6\x12\xd6\xe4\x51\x72\xd8\xc4\xaf\x0e\x85\x0a\x88\xf3\xdd\xb2\x1b\x4f\x5f\xc2\x64\x17\xc4\xd6\x93\xe4\x59\xa8\x35\x73\x4c\xed\x68\xf6\x60\x17\xb8\xe8\xf0\x76\xd4\x2d\xce\xa8\xa9\xef\x2b\x3e\xfa\x93\x86\x95\x2a\xcc\x7d\x70\x2b\x94\x7e\x7a\xa8\x9b\x20\x2b\x5b\x97\xcd\xb2\x95\xc0\xd4\x9f\x5b\x4a\x68\x6a\xeb\x85\xd6\x28\x79\xb7\x2a\x7a\xe1\x82\x4b\xda\x2a\x6f\x87\xbf\xef\x6b\x1b\x98\x56\x0f\x2d\xbf\x05\xaf\xdb\x4e\x5f\xa0\xea\x7b\x2b\x10\xb4\x11\xb4\x4b\x50\x8c\xf9\x5d\x50\x50\xfc\xbd\x4b\x5c\xfb\xd1\xa2\x9c\xa3\x68\xfb\x03\x00\xfd\x74\x09\x70\xe0\x97\x2a\x34\x71\x9b\x81\xbc\x80\x32\xe2\x87\xa3\x84\xac\xd6\x81\x2d\xbe\x86\xdf\x41\xea\x44\x93\x7f\x03\x00\x00\xff\xff\xa8\xe6\x02\x07\x5e\x04\x00\x00")

func _10EkscltAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		__10EkscltAl2Conf,
		"10-eksclt.al2.conf",
	)
}

func _10EkscltAl2Conf() (*asset, error) {
	bytes, err := _10EkscltAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10-eksclt.al2.conf", size: 1118, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x41, 0x3, 0xc, 0x11, 0x90, 0x3a, 0x28, 0x60, 0x6f, 0x34, 0x51, 0x44, 0x54, 0x45, 0x93, 0x1f, 0x12, 0x5b, 0xff, 0x54, 0xcc, 0xac, 0x2e, 0xd6, 0x74, 0x3b, 0xca, 0x88, 0x99, 0x9a, 0x1a, 0xae}}
	return a, nil
}

var _bailSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x6f\xf2\x30\x0c\x86\xef\xf9\x15\xfe\x0a\xd7\x12\x7d\x3f\x01\x8d\x0a\xa1\xb1\x22\xc1\xd0\x2e\x93\x8a\x9b\x18\xb0\x08\x49\x95\x38\xb0\x69\xe2\xbf\x4f\x45\xed\x6d\xc7\xd7\xef\x63\xeb\xf1\xe4\x9f\x6e\xd9\xeb\x16\xd3\x59\xa9\x09\x6c\x5a\x41\xf6\x60\x5c\x4e\x42\x11\x4c\x24\x4b\x5e\x18\x5d\x02\xf4\x16\x72\x67\x51\x08\x0e\x97\xdc\x92\x09\xfe\xc8\xa7\xd9\x37\x5e\xdd\x01\xee\x2c\x67\xc0\x8e\x13\xc5\x1b\x45\xd8\x6f\xd7\x4a\x25\x12\x28\x03\x50\x8c\xf4\xc5\x32\xc6\x8e\x3b\x3a\x22\xbb\x31\xfb\x90\x7d\x22\x51\x2a\x85\x1c\x0d\x81\x26\x31\x9a\x2e\xc9\x88\xd3\x57\x12\xb4\x28\x38\x23\x7f\xfb\xab\xef\x35\x1c\xc9\xcc\x05\x83\xee\x09\x29\x1d\x3a\xd1\x78\x4f\xcf\xaf\xcc\xd1\x97\x89\x4f\x1e\x1d\x7c\x2a\x80\xb2\xec\x45\x4a\x13\x2c\x41\x31\xfd\x5f\x0c\xc3\x24\x68\x2e\x50\x4c\x7f\xe6\x1f\xbb\xa6\x7a\xdd\x35\x2f\xeb\xfd\xee\xbd\xda\x36\xf5\xfc\xad\x7a\x8c\x54\xa4\x41\xa0\x0e\x96\x96\x31\xe4\x6e\x28\xd8\xf6\xbb\xf5\x62\x53\x35\xab\xc5\xa3\x18\xe0\x13\x07\x3f\xde\xdc\x56\xcb\xd5\xa6\x7e\x14\xea\x37\x00\x00\xff\xff\xf5\x38\xeb\x60\x6d\x01\x00\x00")

func bailShBytes() ([]byte, error) {
	return bindataRead(
		_bailSh,
		"bail.sh",
	)
}

func bailSh() (*asset, error) {
	bytes, err := bailShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bail.sh", size: 365, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x56, 0x4c, 0x8b, 0x5c, 0xe2, 0x2, 0x3d, 0x32, 0x4b, 0xfa, 0x3f, 0x60, 0x15, 0x41, 0x57, 0xab, 0x5c, 0x9c, 0xbd, 0x40, 0xd4, 0x38, 0x1a, 0x59, 0x58, 0xdb, 0xdb, 0x48, 0xcf, 0x3f, 0x10, 0xb3}}
	return a, nil
}

var _bootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcb\x31\x0e\xc2\x30\x0c\x46\xe1\x3d\xa7\x30\x85\x01\x86\x34\x27\x80\x09\x06\x16\xe0\x06\xc8\x49\x7f\x29\x55\x9d\x44\xaa\x5d\x24\x6e\xcf\x82\xaa\xae\xdf\xd3\xdb\xef\x42\x1c\x6b\xd0\x4c\x1e\x8b\x73\x48\xb9\x51\xf7\x78\x5e\x6f\xef\xfb\xeb\x7c\x38\xe6\xa6\x56\xb9\x80\xfc\x78\xea\xe8\x42\x01\x96\x02\x26\x4d\x26\x61\x5a\x22\x04\xd6\x4b\x4b\x2c\x3d\xea\xc7\x39\xfd\xaa\xa1\x24\x13\x1a\x18\xa5\x55\x3f\x43\x1a\x0f\x1b\x47\xe5\x28\xa0\xff\xbb\x09\x6a\x3c\xdb\xea\xbf\x00\x00\x00\xff\xff\xef\x6e\xff\x33\x97\x00\x00\x00")

func bootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAl2Sh,
		"bootstrap.al2.sh",
	)
}

func bootstrapAl2Sh() (*asset, error) {
	bytes, err := bootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.al2.sh", size: 151, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb3, 0x29, 0xcc, 0x4a, 0xb5, 0x89, 0x68, 0x6a, 0xda, 0xc1, 0xe9, 0x96, 0xb6, 0xb8, 0x89, 0x27, 0x40, 0xe9, 0x4f, 0xd6, 0xbc, 0x86, 0xbb, 0x9, 0xab, 0x70, 0xea, 0x2a, 0xf2, 0x9b, 0xef, 0x43}}
	return a, nil
}

var _kubeconfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x41\x8f\xd3\x30\x10\x85\xef\xfe\x15\x4f\xdb\x2b\x4d\xd5\x1b\xca\x0d\xad\x38\x20\x21\x90\x10\x70\xa5\x83\xf3\xda\x8c\xea\xd8\x91\x67\xd2\xdd\xfd\xf7\xc8\x21\x0b\x3d\xec\x6d\x34\x33\xef\xbd\xf1\xe7\x1d\x3e\x65\x75\x95\x84\xd3\x75\xf9\xcd\x58\xf2\x59\x2f\xdd\x8b\x4c\xe9\x04\xe7\x34\x27\x71\x86\x1d\x4e\x17\xfa\xaf\x58\x39\x30\xb7\x65\xeb\x6c\x3c\xe1\xa9\xaa\xd3\x50\x16\xc7\xe3\x07\x44\x56\x7f\x07\xc9\x03\x8c\x6e\x90\x59\x8d\xf5\xc6\x8a\x1f\xdf\x3e\x87\x1d\xbe\x7c\xfd\xfe\xb1\x87\x8f\x34\x22\x96\x69\x62\x76\xc3\x93\xa6\x84\x5c\x1c\x32\xcf\x94\x0a\x39\x3b\x2b\x7c\x54\xc3\x59\x13\x0d\x97\x66\xb5\xcc\x83\x38\x87\x20\xb3\xfe\x64\x35\x2d\xb9\xc7\xed\x18\xae\x9a\x87\x1e\x8f\xeb\xc5\x21\xa6\xc5\x9c\xd5\xfa\xb0\xc7\x56\xf7\x01\xc0\x7a\x96\x9e\x35\x8a\x73\x2f\x8b\x8f\xa5\xaa\xbf\xf4\x78\x38\xd0\xe3\x81\x57\x8b\x9e\x0e\x51\xba\x58\xfd\x21\x00\x59\x26\xf6\x68\x24\x6a\xa6\xd3\x42\x2c\xd9\xf9\xec\xab\xef\xff\x61\xa2\x07\x60\x9b\x6d\x39\x5b\xe8\xbd\xb8\xf5\x17\x7b\x6d\x26\x3a\x10\xe2\x52\x2b\xb3\xef\x5f\xc5\xff\xfc\xda\xe2\x9b\x31\xab\xc3\xea\xc5\x67\xc6\xbf\x15\x70\x0f\x23\x26\x65\xf6\xae\x3d\xaf\xfd\x4f\x14\xd7\x92\xbb\xeb\x7b\xeb\xb4\x1c\x6e\x47\x49\xf3\x28\xc7\x4d\xd7\xd8\x4b\x03\x77\x0f\xe0\x4e\x59\x6a\x67\xe3\x9f\x00\x00\x00\xff\xff\xf3\x8f\xab\xb4\x16\x02\x00\x00")

func kubeconfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubeconfigYaml,
		"kubeconfig.yaml",
	)
}

func kubeconfigYaml() (*asset, error) {
	bytes, err := kubeconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubeconfig.yaml", size: 534, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa0, 0x93, 0xd0, 0x68, 0x60, 0x69, 0xf4, 0x6d, 0x37, 0xdf, 0x47, 0x73, 0x4f, 0x94, 0xa7, 0xdf, 0xed, 0xf1, 0x52, 0x91, 0x84, 0x23, 0x18, 0x24, 0xdb, 0x99, 0x4c, 0x6a, 0xc4, 0x5d, 0xee, 0x8}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"10-eksclt.al2.conf": _10EkscltAl2Conf,

	"bail.sh": bailSh,

	"bootstrap.al2.sh": bootstrapAl2Sh,

	"kubeconfig.yaml": kubeconfigYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"10-eksclt.al2.conf": &bintree{_10EkscltAl2Conf, map[string]*bintree{}},
	"bail.sh":            &bintree{bailSh, map[string]*bintree{}},
	"bootstrap.al2.sh":   &bintree{bootstrapAl2Sh, map[string]*bintree{}},
	"kubeconfig.yaml":    &bintree{kubeconfigYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}