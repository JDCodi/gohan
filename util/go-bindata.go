// Code generated by go-bindata.
// sources:
// etc/schema/core.json
// etc/schema/gohan.json
// etc/extensions/gohan_extension.yaml
// etc/templates/swagger.tmpl
// DO NOT EDIT!

package util

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

var _etcSchemaCoreJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\x3f\x8f\x9b\x30\x14\xdf\xef\x53\x58\x5c\xa7\xaa\x77\x74\xe8\x74\x5b\xa4\x2e\x95\x2a\x35\xea\x5a\x65\x70\xc0\x24\x3e\x11\xdb\xb5\x8d\xda\xe8\x94\xef\x5e\x1b\x83\x13\x8c\x9f\x81\x1c\x65\x02\xfb\xfd\xf9\xbd\xff\x8f\xb7\x07\x64\x9e\xec\x83\x2a\x8e\xe4\x84\xb3\x17\x94\x1d\xb5\x16\x2f\x79\xfe\xaa\x38\x7b\x72\xa7\xcf\x5c\x1e\xf2\x52\xe2\x4a\x3f\x7d\xfe\x92\xbb\xb3\xc7\xec\x93\xe3\xc4\x75\xfd\xa3\x32\x7c\xbf\xda\x4f\xfb\xbc\xf9\x37\x27\x5a\x92\x6a\x89\x5c\xcf\x7d\x69\xdf\x76\x9d\x9e\x92\x54\x94\x51\x4d\x39\x53\x46\xda\x55\x47\x26\x89\xe2\x8d\x2c\xc8\xe0\xb4\x63\x51\x85\xa4\xc2\xf2\x58\x00\x3f\x3b\x42\x74\x15\xd5\x19\xe1\x39\x84\xe4\x82\x48\x4d\x89\x1a\x49\x8b\x48\x1c\x13\xf4\x48\x71\x53\x6b\xab\x32\x90\x0f\x21\xfb\x7a\xf3\x09\x70\x68\xaa\x6b\x6b\xe2\x80\x15\xa2\x3d\x8b\x96\x54\x69\x49\xd9\x21\x1b\x11\x5d\xc6\x7c\x19\x2d\x53\x06\x0d\xd0\x1a\xd2\x29\x90\x09\x92\x3b\xb0\x09\x2c\x09\xd3\x6b\x3a\x7c\xdb\x4a\x44\xbc\x42\xfa\x48\x15\xea\x92\x7f\xca\xaa\x0e\xc7\x9a\x96\xd5\x8d\xc4\xf5\x6c\xcf\x6f\x5b\xf2\xa5\xb0\x1d\xd7\xaa\xb0\x4d\x49\xd3\xbf\xab\x06\xc4\x49\x9c\x8c\x40\x9a\xec\x0e\x53\x7c\xdf\x03\x4c\xe9\xbb\xd7\xe3\x3c\x71\x3d\xd2\x79\xf1\xfc\x8e\xf7\x64\x71\x38\xdd\xcb\x8a\x2e\xc0\xc5\xb8\xab\xa6\x50\xfb\x3e\xba\xe9\x18\x01\x2c\x02\x6b\x4d\x24\xdb\xa6\x7b\xaa\x27\x7f\xfe\x98\xbc\x77\x22\xe7\xc9\xf2\xf4\x94\x89\x06\x6e\x1c\x23\xf2\x44\xb0\xc3\x27\xe2\xc8\x91\x38\xde\xe8\xff\xa6\x3e\x49\x31\x01\xce\x67\x09\xdf\xbf\x92\x42\xc3\xda\xe2\x5a\x00\xe9\xd7\x0c\x9d\xc8\x8b\x29\xed\x43\xad\x81\x36\x33\xee\x7f\x37\x54\x92\x72\xb0\x72\xf8\xdb\xe8\xe0\x01\x4b\x26\x03\x4b\xae\xef\xcc\x70\xf3\x1b\x5c\xec\x02\x90\xde\x15\x07\x7e\xc4\x0c\xf9\x15\x25\x24\x03\x3c\x71\x63\xb3\xa9\x64\xac\xc9\x29\x36\xff\xba\xc2\x8d\x6f\x2a\xa9\x46\x10\xba\xb4\xe0\xac\xa4\xb0\x28\x6a\xf4\x27\xba\xc3\x3d\x2d\xa7\xe7\xc1\x52\xe2\x73\x1a\x5c\xd0\x7c\x04\xaf\x69\x71\x46\xde\x2b\xa8\xe2\x12\x39\x9a\xbd\x51\x8f\xba\x7b\x7b\xba\xd9\x7e\x0b\x1d\x4e\xaa\xca\x3a\x3a\x6a\x26\x61\xcd\x29\x9a\x54\xed\xad\x59\x73\xf9\x1f\x78\x98\xb1\xf3\xd8\xf0\x5d\xc2\xf0\x39\x61\x01\x56\xb2\x59\xeb\xd8\xc4\x2a\xb6\x08\x87\x30\x24\x05\x15\xc0\x9e\xb2\x48\x14\xb8\xab\x3b\x45\x44\x9e\xa8\x52\xce\x2e\x20\x0e\x85\x24\x26\xf2\x50\x20\x1a\x51\xda\xdb\x59\xa1\x98\x31\x4b\xec\x08\x3b\xa6\xa7\xd6\x54\xf6\xdb\x07\x6a\x98\x33\xa7\xd9\x44\x01\x2e\x82\x92\x80\x33\x90\x11\x29\x4c\xcf\xbe\xa8\xc2\x63\x4d\x3e\xcc\x8a\xeb\x7a\x17\x14\x37\xd4\x11\xad\x7d\x9b\x16\xe1\xa8\x27\xc2\xbe\x5a\x94\xa8\x80\x1f\x9c\xed\x17\xff\x4b\x3a\x28\xc4\x8d\xe9\xf3\x66\x38\x28\x03\x1c\xdb\x33\xbb\xd8\x61\x34\x9c\x00\xfd\x50\x84\x62\x9f\xfc\xc3\x04\x2d\xb8\xf5\x8e\xf3\xe1\x22\xc7\xf8\xbd\x23\xbf\xf9\xc9\xce\x23\x51\x58\xe2\xa9\x1b\x48\x57\xe3\xdf\x8f\xca\x77\x90\xf7\x82\x8a\x6f\xea\xb0\x8b\x07\x91\x0f\xb3\xfb\xf2\xf0\x2f\x00\x00\xff\xff\x95\x62\x9d\x2b\x48\x11\x00\x00")

func etcSchemaCoreJsonBytes() ([]byte, error) {
	return bindataRead(
		_etcSchemaCoreJson,
		"etc/schema/core.json",
	)
}

func etcSchemaCoreJson() (*asset, error) {
	bytes, err := etcSchemaCoreJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/schema/core.json", size: 4424, mode: os.FileMode(420), modTime: time.Unix(1445535064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcSchemaGohanJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\x4f\x6f\xdc\xba\x11\xbf\xfb\x53\x08\x8b\x1e\xda\x20\x48\x9a\x6b\x6f\x45\x63\x14\x2e\xdc\xd8\xb5\x9d\x5c\x02\xd7\xa0\x77\xb9\xb6\x52\x2d\xa5\xe8\x8f\x93\x45\xe0\xef\x5e\x91\x94\xb8\x92\x2c\x0e\x87\xa4\xa4\x95\xfc\x22\xbc\x83\x5f\x34\x1c\x91\x3f\x0e\xc9\xdf\x0c\xc9\xd9\x5f\x27\x41\xf9\xac\x92\x38\x0a\xd7\x21\xcd\x56\x7f\x0b\xbe\x8a\x7f\xe1\xcf\x2f\xf5\x97\x90\x21\xeb\x3c\x8c\x59\x29\xb1\x7a\xb3\x7a\xdb\x7e\x45\xb7\x5b\xba\xce\xf9\x2b\x12\x45\xf1\x8f\xee\xeb\x70\x23\x5e\x6d\x76\x21\xbb\xcb\x72\x92\xd3\x1d\x65\x79\x57\x28\x49\x43\xb6\x0e\x13\x12\x29\xd9\xae\x44\x4a\xb3\xb8\x48\xd7\xb4\x14\x68\x57\x4d\x96\x27\xf9\x23\x2f\xfa\xee\xcd\xaa\xf5\xf2\xf9\xa4\xfd\xd7\xad\xd4\xba\xca\xd6\x8f\x74\x47\xc0\x16\x6f\x68\xb6\x4e\xc3\xa4\x6e\xf6\xcd\x23\x0d\x76\x24\x64\xc1\x8e\xe6\x44\x16\xef\x6f\x69\xff\x3b\x5e\x6a\x43\x72\xd2\x5f\xfd\x7c\x9f\xf0\x86\xad\x1a\xba\xdb\xcd\xe8\xc2\x15\x15\xa9\xc4\xaa\x6e\xc8\x0b\x3c\xe9\x36\xfc\xc9\x05\xde\x3f\xc4\x8f\x84\xbd\x7f\xfa\xeb\xbb\x0f\x5d\xa1\xea\x4b\xbd\x35\x22\x9b\x4d\xc8\x9b\x4e\xa2\xcb\x34\x4e\x68\x9a\x4b\xfb\xd8\x92\x28\xa3\x6f\x7b\xf0\x6f\x0a\xbd\x54\xd7\x03\x68\xbf\x50\x25\xb8\x25\x45\x24\x2c\x6a\xf5\xf2\x5b\x1a\x7d\xab\x8f\x8d\xff\x05\x4a\x95\xd5\xdc\x85\x59\x26\x0b\x7d\xd5\xca\x09\xd9\x75\x4a\x4b\x7b\x05\xb4\x09\xa9\x22\xd9\x70\x29\xad\xd0\x2d\x50\x9b\x3c\xcc\x23\x6a\x51\xfb\xda\x50\xb2\xbc\x1c\x30\x0f\xfd\xdf\x7c\xee\x2f\x2f\xcd\x13\x84\xbd\x85\xe7\xb5\xb0\x8e\xe0\xec\xe3\xc0\x68\xfa\xe1\x04\x57\xa7\x03\x0f\x20\x59\xb0\xf0\x7b\xc1\x65\xf3\xb4\xa0\x56\x30\x56\xdf\xb0\x06\xf2\x86\x97\x9b\x13\x94\xa6\x0a\xf9\xd8\x1a\xfd\x99\x53\xb6\xd1\x4f\x06\x42\xa8\x83\xd3\x69\x55\x66\x4e\x18\x21\xea\x54\xc3\x44\xd2\x94\xec\x21\xc1\xb0\x5c\xfa\x60\x44\x5a\xfa\x20\xd8\xf9\xf3\x6c\xd5\x21\xe0\xfa\xa3\xa4\x0e\x33\xef\x2f\x8d\xa2\x4a\xac\xd5\x71\xb5\xee\x60\x1b\xa7\x01\x49\x92\x92\x4d\x10\xfe\x2e\xd8\xd0\x27\x1a\xf1\x85\x01\x82\xa5\x2c\xb4\x23\x62\xb2\xdf\x93\x5d\xb4\xc0\xa9\xfb\xdf\x35\xb4\x08\x23\x89\xef\xbf\x71\xae\x64\xd5\x75\x8c\xec\x68\x96\x10\x0d\xf5\x51\x62\x6e\xab\xe6\xa7\x5a\x79\x90\x97\x04\x47\x32\x82\xe0\xbe\xec\x35\xf6\x90\x05\x79\x3c\xab\xc1\xa8\xea\x3a\xd6\xac\x95\x90\x94\x73\xd3\x11\x60\xbe\x14\x9a\x83\x78\x5b\xc2\x1c\x66\x41\x2f\x47\x6c\x57\x65\x9e\xb6\x2e\xdb\x31\x1a\xfe\x35\xad\xc5\x2f\x1a\x97\xa2\xc8\x2b\x00\x56\x36\x7d\x2c\x60\x6b\x77\x60\x70\xc3\xae\x34\x2f\x10\x70\x63\xc5\x7d\x00\x07\x5c\x2b\x25\xd3\x58\x6c\xe1\xf6\x22\x9c\xab\x96\xbc\x81\xe8\xb7\x75\xe3\x7b\xa7\x55\xce\x34\xa1\x37\x1f\xa0\x3f\x5a\x3a\xb1\xcc\xa7\x7e\xfa\x19\x90\x7a\x6b\xb0\x22\xcc\x92\x6c\xd0\xd3\x17\x23\x08\x19\xa3\xa9\x3e\x48\xd0\x29\x9e\x93\x30\xba\x7b\x0a\xe9\x8f\xca\x13\x59\xcc\x38\xb2\xb0\x49\x5d\x2c\x01\x61\xc9\x75\x1f\xdd\xc7\x71\x44\x09\x83\x4d\xc2\xd4\xe1\xb6\xc3\x28\x21\x79\x4e\x53\x66\x59\x67\x51\xf4\xbf\xef\xde\xfc\x09\x2d\xed\x52\xb9\x56\x59\xec\x44\xd2\x5b\x98\xb0\xfd\xc5\xd6\x6a\xec\x37\x1f\xfb\x0f\xaa\x0f\xb7\xfd\x26\x27\x3d\x86\x1e\xd7\x3d\x03\x54\x1a\x65\x90\xba\xe7\x78\xd5\x0e\x59\x4e\x1f\x4a\x77\x6c\x61\xd5\x66\xc5\xee\x7e\x79\xb5\x36\x2d\x2c\xd0\x73\xbc\x5a\x63\x17\xdf\xbe\x07\x5e\x90\xfb\x1e\x24\x35\x68\x3e\xcd\x98\xa9\x9c\xf7\xac\x54\x58\x22\x8b\x8e\x5a\x6b\x15\x58\x86\x78\xf5\x7a\x3c\x3a\xc8\xbe\xd1\x4d\x62\xe2\xd7\x68\xae\x29\xf8\x52\x6a\x0a\x2e\x58\x04\x05\xc8\xf4\xca\x7c\x26\x5d\xdb\xa6\xd3\x72\xb2\x71\x6b\xf3\x2e\x64\x67\x55\x88\xef\x83\x8f\x59\x9f\xf2\x1a\x78\xc0\x64\x0a\x44\x6a\x15\xc8\x98\x78\xdd\x06\x6d\x60\x5c\xf7\xd8\x22\xad\x02\x7f\x3e\xf6\x55\x29\x99\xf9\x78\xb2\xf0\xc8\x5a\xe5\x70\xdb\x1e\xfa\xe2\x53\x36\x11\x15\xdd\xee\x2d\x8a\x8e\x00\x6b\x35\x1c\x70\x12\xb5\xf0\x80\xca\x65\xd1\xb6\x85\xaa\x72\x2c\x7c\xc1\x4a\x4b\x2a\x07\x05\x30\xb4\x2a\x1a\xd1\x3c\x59\x91\x99\x9b\x56\xcb\xf3\x75\x80\xcc\xdd\x34\x45\xf1\x6a\x45\x70\x73\x90\x84\x06\x94\x3b\x0e\x6a\x30\xb9\xea\xd0\xe3\xc0\xb1\xc4\x37\x7d\x19\xa1\x03\x87\x6d\x98\xe6\xa1\xcf\xbd\x97\xc2\x71\x8d\xd3\xdd\x71\x17\xe5\x07\x9c\xfc\x1a\x91\x8a\x65\xcc\x80\x0e\xa1\x95\x96\x9a\x21\xb1\x93\x35\x0a\x96\x83\x61\x4a\x23\x32\x80\x57\x72\x55\xab\x99\xf9\x22\x50\x37\xf7\xae\x1a\x70\xfb\x81\xda\x5d\xf7\xb8\x97\x67\x32\x0d\x00\xdf\x8b\x30\xa5\x8e\x4c\xd2\x73\x0d\x3c\xea\x5a\x70\x55\xb5\x3c\xd8\x86\x34\x02\x0f\x63\xe8\x75\x2d\xcc\x37\xca\xbe\xc3\xdb\xa0\xda\x82\x0a\xb4\xeb\xff\x9c\x07\x39\x7c\xc2\x47\xaf\x64\x42\xb3\xae\x2b\xec\xd3\xd6\x1b\xf1\xc7\xdc\x1b\x6a\x3a\x37\xa6\x2d\x19\x33\x7a\x9c\x10\xbd\x37\xf3\x15\x5a\x5c\xc7\x5c\x4b\x49\x1d\xfb\xf1\x54\x53\x07\xc0\x3d\xd5\xb0\x22\x72\xa1\x1b\x1d\x1d\x22\xa8\xed\xa9\xa5\x62\x1a\x9e\x5a\x7c\x82\xbe\xfc\x71\x74\x31\xc4\xb7\xb1\xe7\x11\xcd\x8a\x86\x88\x60\x4f\x1e\x77\xf7\x5b\x98\x95\x9a\x41\x86\xaa\xd0\x34\xc4\x70\x15\x8a\x86\x19\xb2\x42\xd5\x30\xc3\x56\xa8\x1a\x60\xe8\x56\x7a\x06\x18\xbe\x42\xd3\x20\x43\x58\x68\xf2\x1d\xc6\xfc\xb9\x75\x2e\xed\x38\x76\xf8\x33\xfc\x24\xe0\x6b\xc5\xbe\x84\xaf\x7e\x1c\xb6\xc4\x46\x65\x21\xea\xdc\xbb\x0f\xdf\xfa\x2c\x95\xcc\x7d\x2b\xa7\xdd\x85\xfe\x0d\x0e\x42\xdf\xe8\xf2\x24\xcd\x1e\x60\xd7\x8e\xab\x08\xd6\x31\xdb\x86\x0f\x45\x4a\xb8\x97\x5c\xfe\xf7\xf9\x6c\xea\x50\x1c\x5a\xda\x02\xa2\x46\xc0\xee\x22\xdd\x94\x33\xb8\xed\xa2\xc9\x37\x73\x9c\x7c\x1c\x17\x7f\xc1\x6d\x4b\xcd\x7a\xb7\x57\xee\xa3\xbb\x87\xe5\xdd\x26\x21\xf7\x48\xab\x77\xb8\xd1\x69\x73\xcb\x63\xfa\x71\xdc\x21\x72\xd9\xbb\x76\x0d\x23\x7a\x1c\x11\x70\x08\x6e\xb8\x06\xbf\x26\x18\x44\xf5\xdd\x25\x74\x29\x0b\xef\xa7\x19\xc5\xb3\x9c\x76\x36\xce\xc7\x47\x9c\x50\x90\x73\xd6\x28\x18\xb8\x44\xc9\xcd\x6b\x01\x62\x0d\x70\xdc\x2a\xb1\xaa\x30\xfe\x14\x6a\xbd\xfc\x60\x4e\x74\x5b\x71\x18\xfb\x83\xd4\x6e\xd0\x05\xb2\x05\x16\x00\x62\x59\xb9\x2d\xfb\x36\x81\x6e\x15\x3b\x9f\x15\xda\x2a\xf6\x6d\xb5\xfa\x1d\x1d\x70\x74\xb0\xd3\xda\xf3\xb3\x82\x5b\x3f\x6b\x40\xa7\xfa\x55\x95\xfe\x75\x7d\xf1\x09\x71\xa5\xc8\xe7\x72\x5f\x56\xb6\xa1\x88\x08\x3c\x0f\xbc\xb8\x51\x5c\x15\x0a\xf8\xd5\xc0\xe5\xdf\x7d\xaa\x9b\x33\xd6\x65\x1c\xf3\x16\x43\x17\xe0\x73\x72\x4f\x5f\xc1\xa5\x32\xd3\xbe\x88\x17\xaa\x32\x23\x88\xf1\x8a\xb7\xcb\x85\xe2\xab\x2a\xd7\x47\x50\x7f\x63\x31\xd0\xdb\x1f\x6c\xe0\x29\x4b\x70\xd7\x51\x1c\xee\xa1\x84\x2c\x29\xec\x4e\x58\xba\x91\x33\x1c\xf9\x5b\xc5\x45\x3e\x4d\x7d\x86\x59\x83\x6d\x58\x9f\xdf\x42\xf3\x77\xb3\x9d\x57\xb5\x31\x58\x38\x2a\x96\x2c\x63\xe0\x7a\xfb\xb6\x9a\x06\xc2\x2c\xae\x0e\x66\x44\xfc\xb6\xff\x18\xd3\xc1\x59\xfd\x89\xe0\x5c\x7c\xe2\xf7\x6c\x30\x29\x0b\x52\xf0\x1b\xd1\x3f\xba\x8d\xbe\xf8\xd7\x9e\x06\xa2\x03\x80\xfa\x40\xdf\x81\xb5\x69\xde\x27\xd0\x6d\x6e\x30\x1c\x08\x5f\xa8\xee\x4f\x2e\xa5\xde\x9a\x43\x04\x75\x8a\x03\xcd\x5b\x66\xc8\xb0\x70\x48\x23\x62\x20\x05\xc8\xc9\xe2\x85\x54\xcf\x68\x32\x87\x4c\x46\xeb\x28\x18\x6d\x4d\x84\xa4\xaf\x09\x07\xaa\x2b\x13\x6c\x7c\x2c\xfb\x98\x85\x9a\x7e\x02\x17\x9e\x6e\x26\xb0\x86\xff\xa0\xc9\x3c\xa6\xbe\xfd\x4f\x9e\x08\x2c\xb8\xee\xa4\x17\x6b\x28\x34\x27\x40\x13\x99\xe2\xf6\xc6\x14\x68\x52\xac\xfb\x6e\xac\x14\x68\x2a\x7d\x5d\x57\xc2\x37\x07\x1a\x26\xa9\x99\xca\x8b\x87\x77\x2e\xaa\x22\x8b\x59\xc5\x3a\x2c\x65\x2c\x7f\x62\x1d\x33\x79\x4d\x1c\x4b\x21\xbe\x42\x95\xee\x60\x7e\x50\x0e\x94\xa9\xe3\x3e\x30\xe2\x43\x2f\xce\x70\xcf\x59\x7c\xcd\x8f\xa7\x82\x3c\x66\x9e\xf6\xf8\x0f\x4c\xa7\x22\xf6\x1f\xb5\x49\xcc\xea\xc4\x96\xf8\xb1\x5d\x15\x01\xea\x83\x3a\xa8\xd3\x9b\x49\xb3\xe7\xdb\x0c\xd8\x51\x5d\x60\x87\x9e\x1a\xd1\x9b\x30\xff\x21\xb8\xcd\x3b\xf3\xc4\x87\x56\xb8\x34\x73\xaf\xe2\xe1\x39\x94\x5a\x9e\x9d\x5d\x62\xea\xee\x03\x29\x98\xac\x56\x49\x75\x10\x55\x85\x16\x03\xa8\x4d\x12\x98\x2a\x43\xaf\xcd\x8e\xc0\x25\x2f\x32\xf4\x8e\x00\x7a\x7b\x6e\x79\x7b\x45\x6e\x9b\x9a\x88\x83\x39\x9e\x01\x84\x2b\x84\x61\xa3\xf6\x50\xa6\x72\xef\x4d\x73\x9b\x69\xa8\xc2\x14\x1f\x26\x09\x0d\xaa\x8a\x73\x2d\x1d\x9d\xc5\x7e\x1f\xad\xe3\x2c\x5e\x4a\x21\x37\x67\x51\x9c\xa3\xe0\x53\x95\xd1\x5f\x54\x92\x53\xb9\x8c\xea\x83\xc7\x71\x1a\xd7\xf1\x66\x94\x64\x9e\x42\x2f\x20\x7e\xb8\x2f\xf9\x8d\x3c\x11\x59\x74\x39\xab\x4d\x83\xf7\xc3\xcd\xf4\xf3\x42\x37\xf4\x0e\x93\xe2\x59\xf5\x0f\x0e\xca\x9e\x9e\x32\x5d\xcf\x9a\x33\xfc\xc6\x6c\xd6\xbf\x89\xba\x46\xcc\x2b\x3d\xad\x81\x42\xbd\xe0\xe8\x30\x81\x9a\xab\x81\x19\x78\x9f\x17\x86\x45\x8a\xde\xa2\xb2\x99\x79\x3f\x5f\x9d\xe3\x26\x5e\xfe\xfd\xe5\xf5\x88\xa1\x79\xa8\x0e\x99\x8a\xbc\x1d\x66\x70\xc0\x04\xa0\x11\x06\xe8\x1d\x97\x92\x69\x69\x50\x87\x95\x9d\x2a\x39\x47\x62\xf6\xc4\x73\x51\x1b\x49\xd9\x53\xcf\xef\xb4\xc0\x84\x8c\xc5\xd9\x9e\xad\x57\xba\x54\xa6\x3e\x84\x8d\x57\xe6\x38\x64\xed\x3e\xde\xc0\xb9\x06\xba\x10\x8b\x02\xa8\xc9\x60\xa1\x09\xef\x4f\x85\xfd\x98\x9a\xe9\x73\x2a\xee\x89\xa6\xc6\x74\x3f\x7d\x96\x5d\x95\x93\x27\xb7\x68\x75\x83\x45\x1e\xe2\x92\x46\x2f\x51\x1a\x9a\x30\x4c\x14\xec\xe6\x6d\xa9\x91\x41\x00\x0f\x26\xfb\x7c\x55\xdc\x4b\x66\x4b\xa8\x1b\x1c\x94\xee\xfb\x8e\xa4\xfb\xe0\x7f\x74\x1f\x90\x22\x8f\xef\x4a\x6f\x3e\x15\x3f\x3b\x15\x80\xb8\xd9\x31\x38\x27\x78\x31\x14\xce\x85\x7f\xc8\x21\xf9\x9b\xef\xf5\x7c\x64\x47\xb3\x9c\xec\x92\xf1\x50\x57\x9f\x08\xfe\x5c\xb0\xf0\x27\xff\xdf\xbf\x2c\xb0\x17\x6e\x14\x52\x63\x19\xbf\x95\x5b\x6d\xdf\x0d\xcb\xf4\xa6\x07\x71\xa4\xa7\xe2\xd6\x10\xad\x86\xb8\x73\x6e\xb0\xad\x15\xbc\xac\x49\x4a\x35\x2e\xf7\xee\x63\xbb\x5d\xde\x2d\xec\xec\x3c\x7e\x70\xe4\xdd\xea\x74\x55\xff\x89\xf2\x8a\x77\xeb\xce\x60\x35\xd8\xb0\x12\x71\x63\xc4\xd3\x92\x78\x5f\xfe\x3d\xf2\xcf\x06\xe2\xae\xd5\xcd\x75\xf6\x98\xdf\xcf\x06\xce\x83\x11\x4e\x10\x8d\x9b\xe2\x17\xd5\x5e\xb5\x1b\x39\xc5\xef\xa6\x8d\x31\x63\x08\xbd\xcb\x83\xfb\x93\xa1\xda\x8b\xf8\xe1\x34\xd3\x01\x65\x59\x9b\xa9\x27\x19\xf0\x50\xb5\x14\xfd\xe3\xfc\x78\x97\x27\x94\xa3\xfe\x2c\x17\x32\xae\xe4\x82\x25\x22\x30\x33\xd3\x89\xe1\x0b\x3e\xa4\x34\x0b\x27\x04\x98\x7f\x51\x37\x20\xc0\x0b\x16\xf0\xfd\x08\x83\x9b\xa2\x96\xed\x51\x5d\x15\xad\x83\xd0\x71\x57\x0e\x3f\x97\x79\x70\x57\xc4\x5f\xb7\x27\xcf\x27\xff\x0f\x00\x00\xff\xff\xfc\x14\x87\xca\xfb\x7d\x00\x00")

func etcSchemaGohanJsonBytes() ([]byte, error) {
	return bindataRead(
		_etcSchemaGohanJson,
		"etc/schema/gohan.json",
	)
}

func etcSchemaGohanJson() (*asset, error) {
	bytes, err := etcSchemaGohanJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/schema/gohan.json", size: 32251, mode: os.FileMode(420), modTime: time.Unix(1456945742, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcExtensionsGohan_extensionYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xad\x28\x49\xcd\x2b\xce\xcc\xcf\x2b\xb6\xe2\x52\x00\x02\x5d\x85\xcc\x14\x2b\x85\xe2\xe4\x8c\xd4\xdc\x44\xb0\x80\x82\x42\x72\x7e\x4a\x6a\x7c\x49\x65\x41\xaa\x95\x42\x7a\x3e\x92\x98\x95\x42\x46\x62\x5e\x4a\x4e\x6a\x3c\x8a\xea\x82\xc4\x92\x0c\x2b\x05\xfd\xf4\x7c\xa0\xa4\x7e\x99\x81\x9e\xa1\x3e\x58\x5a\x4f\x0b\x10\x00\x00\xff\xff\xc8\x52\x27\xbc\x6a\x00\x00\x00")

func etcExtensionsGohan_extensionYamlBytes() ([]byte, error) {
	return bindataRead(
		_etcExtensionsGohan_extensionYaml,
		"etc/extensions/gohan_extension.yaml",
	)
}

func etcExtensionsGohan_extensionYaml() (*asset, error) {
	bytes, err := etcExtensionsGohan_extensionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/extensions/gohan_extension.yaml", size: 106, mode: os.FileMode(420), modTime: time.Unix(1456945742, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcTemplatesSwaggerTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\x51\x6f\xdb\x36\x10\x7e\xef\xaf\xb8\x69\x0b\xb0\x01\x85\x9d\x16\x7b\x2a\xd0\x87\x62\x05\x86\x0c\xcd\x12\x2c\xe8\xd3\x50\x0c\x8c\x74\xb2\x59\x28\xa2\x4a\x52\xcb\x0a\x47\xff\x7d\x47\x52\x92\x29\x8b\xb4\x15\xbb\xf1\xd6\xc2\x7a\xb2\xc9\xe3\x1d\xf5\xdd\xdd\xc7\xd3\x71\x75\x06\xac\xd6\x02\x55\xca\x2a\x04\x91\xe7\x70\xd6\xac\x9e\x01\x3d\x89\xba\x67\x8b\x05\xca\xe4\x15\x24\x2f\x67\xe7\xc9\x73\x37\xca\xcb\x5c\xd0\x90\x93\xb1\x23\x7f\xa3\x54\x5c\x94\x46\xee\x7c\xf6\xa2\x95\xb3\x33\x9a\xeb\x02\xcd\xf8\x42\x2c\x59\x09\x6f\xae\x2f\x12\x3b\xd9\xb4\xba\x6e\x99\xc2\x6b\xa6\x97\x46\x64\xde\x19\x50\xe9\x12\xef\x50\xd1\xd8\x9f\xad\xa6\x64\xa9\x75\xe5\x56\x7e\x68\x85\x52\x51\xaa\x7a\x28\x45\xa3\xac\xaa\x0a\x9e\x32\x4d\xbb\x99\x7f\x54\xb4\xa5\xc1\x9a\x4a\x8a\xac\x4e\x1f\xb9\x86\x76\x67\x16\xac\x80\x80\xca\x85\x04\xbb\x3b\x06\xbc\x6c\x7f\x29\x83\xd7\x19\xf0\xbc\xfd\x3f\xbb\x44\xcd\x32\xa6\xd9\x4c\x7f\x26\x40\xbf\x7b\x0d\xc9\x1d\x8d\xb8\xc9\x84\x84\xd7\x96\x57\xab\x6e\xcd\xaf\xa8\xaf\x8b\x5a\xb2\xe2\xfd\x1f\xef\x7e\xfc\x09\x9a\x66\x00\xb0\x15\x5e\xa0\x1e\x0d\xda\x89\x8c\x5c\x27\x79\x65\xf6\x9f\x00\xe1\x48\xba\xa0\xe0\x4a\x93\x2f\x61\x6d\xe1\xe2\x6d\xd3\x80\x44\x25\x6a\x69\x10\x78\x3e\xd6\x13\x44\x67\x20\x11\x46\xca\x7f\x3e\x04\xf4\x92\xd1\x8a\x9c\x65\x15\x8f\xb7\x6f\x45\x5e\x9e\x9f\x47\x27\x47\xef\xf8\xca\x07\xce\x9b\x30\xa8\x8d\xcd\xf7\x2a\x5a\x07\x6c\x33\x63\xe5\x8c\xd7\x8c\x0d\x26\x25\xfb\xbc\x45\xa1\x15\xe6\x1a\xef\xe2\xef\x35\x10\xfd\x41\x62\x6e\xf4\x7e\x3f\xcf\x30\xe7\x25\x37\x5b\x56\x73\xdf\x41\xe6\x05\xb6\xea\x69\xa2\xb3\xe1\x99\x26\xbc\x7b\x42\x33\x67\x75\x11\x0e\x27\x4f\x68\x00\x79\x5d\xe2\x3f\x15\xa6\x1a\x33\x40\x29\x85\x3c\x10\xea\x08\x1c\x56\xf5\xa5\xc8\xb0\x88\x23\x11\x79\xd7\xd1\xe8\x70\x64\x03\x8b\xa4\x12\x6a\x62\x3e\xfd\x22\x91\x69\x84\x12\xef\x61\xc3\x5b\x7d\x3e\x1d\x33\x9d\x2a\x26\x19\x11\x0a\x51\x6e\x50\x73\x24\xc3\x4a\x5a\x33\xcc\x1c\x17\x6f\x91\x08\xe1\xd6\xe7\xb7\x22\x8b\x65\x40\x3c\x23\x3b\x64\x84\x32\xc8\x10\x4f\x56\xb5\x8e\x29\x91\xf8\xa9\xe6\x12\x33\xd2\xa0\x65\x8d\x11\xa9\x1d\xe1\x34\x2d\xb3\x2e\xec\x3e\xa6\x86\xce\xde\x3c\xf6\x62\x4f\x1e\x73\xa8\xa5\x36\xd4\xb2\x2f\x41\x63\x07\xd2\xcd\x89\x50\x76\x12\x4a\xff\xcf\x43\xc5\xf3\x28\x9d\xe4\x84\xf0\x7c\xc5\xb3\x43\x0e\xf3\x9b\xa5\xb8\x07\xf6\xbf\x38\xc7\x0f\x23\x1e\x1e\x8b\xea\x96\x6b\x4c\x9d\x35\x91\x6b\x28\x74\x37\x4a\x1b\x93\x3b\x5a\x40\x8e\x3a\x8d\x2a\x99\xc6\x35\xdd\xe9\xaf\xb4\xe4\xe5\x62\x0c\xcc\x38\x2a\xbe\x81\x92\xe7\xc4\x15\xc7\x2e\x3e\xea\x89\xe9\xff\xbe\xca\x4c\xed\xf1\xf5\xd7\x1d\xc7\x48\xff\xda\x82\xf5\xc4\xf9\x3f\x5e\x77\xaa\xb8\xc2\x84\xe1\x62\xf7\xe9\x4b\xae\x7d\x79\xf4\x54\x72\x7d\xed\x34\x4a\x26\x88\x90\xa6\x31\xe9\x5b\x2b\xbb\x51\x4a\xc1\xa9\x96\xfa\x36\x6a\xa9\x9f\x0f\xe2\x00\xaa\xc6\xc1\xc5\xd2\x89\x07\xf6\x78\xe9\xff\xf2\xd3\xab\x6d\xca\xb2\xd4\x96\xc3\x7d\x53\x76\xf6\xc6\x0e\x98\xde\x6c\xf0\xe3\x6c\xd4\x6f\xa5\x29\xa7\x63\x66\xda\xd1\xf0\x00\x6d\xef\xfb\x2f\x93\x4b\xc1\x7e\xec\x7a\xc5\x25\xea\xa5\xc8\x1e\x0a\x71\x8f\xd2\x88\xc2\x4e\x42\x22\xac\xdc\x06\x61\xad\xc5\x7c\xd8\x9d\x28\xe8\xd8\x14\xe4\x3a\xf7\x03\x47\x92\xf3\x9d\x2b\x5f\xbf\x6e\xdb\x84\xd0\x47\x58\x50\x84\xaa\x2d\x18\x04\xda\x44\x0c\x7d\xdf\xbb\x93\xe8\x53\x8d\xa6\x6f\x7f\xfb\x91\x92\xf3\x4b\x16\x8a\x7b\x98\x79\x5c\xa1\x18\x9c\xf5\x0c\x9b\x92\xf4\xc6\x5d\x9c\xf4\xa9\x45\x5b\x0b\xea\x6c\x02\xe9\x4f\x6e\xc2\x32\xe3\xb9\x7f\x7d\xd2\x3d\x47\x2a\x1c\x59\x9f\xb0\x9b\xe5\x83\xb5\x34\x85\x43\xa3\x12\x6b\xa4\xae\x6a\x3d\x1d\x2a\xab\x3c\x00\x97\x79\x4e\xe7\x45\xe8\x5f\x1b\x48\xe6\xc8\xe8\x6e\xed\x4a\xa1\xcd\x11\x52\x08\x51\xcd\xde\x31\xca\x0b\xca\x64\x3f\xdc\xbc\xdf\xfe\x62\xab\xb2\xbb\xc6\xf4\x5e\xe3\xc9\xaf\x0a\x5d\xf5\x00\xeb\x60\x5a\x4f\xfd\x76\x73\xf5\xfb\x96\xd0\x49\x68\x61\xf0\x34\xf4\x1a\xe4\x3b\xf4\x5e\x95\xed\x6d\xc8\x03\x3c\xda\x40\xfb\x3d\xb8\xd3\x42\xdb\xf3\xe8\xdf\x00\x7c\x0b\xcd\xc0\x39\x23\x87\x58\x19\x2f\x90\x46\xa7\x76\x77\x20\x04\xc9\xcf\x9c\xb5\x15\x4a\xcd\x23\xbc\xe1\x34\xc7\x29\x65\x23\x6d\xac\x34\xdc\xa1\x52\x6c\x11\x6d\x4b\xf4\x37\xe4\xdb\x32\xeb\x91\xa5\xb4\x17\xef\x2e\x4c\x9f\x75\x61\xec\xdd\xf4\x9f\x35\xff\x06\x00\x00\xff\xff\xa9\xf9\xf2\xf0\xfa\x1f\x00\x00")

func etcTemplatesSwaggerTmplBytes() ([]byte, error) {
	return bindataRead(
		_etcTemplatesSwaggerTmpl,
		"etc/templates/swagger.tmpl",
	)
}

func etcTemplatesSwaggerTmpl() (*asset, error) {
	bytes, err := etcTemplatesSwaggerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/templates/swagger.tmpl", size: 8186, mode: os.FileMode(420), modTime: time.Unix(1458247777, 0)}
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
	"etc/schema/core.json":                etcSchemaCoreJson,
	"etc/schema/gohan.json":               etcSchemaGohanJson,
	"etc/extensions/gohan_extension.yaml": etcExtensionsGohan_extensionYaml,
	"etc/templates/swagger.tmpl":          etcTemplatesSwaggerTmpl,
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
	"etc": &bintree{nil, map[string]*bintree{
		"extensions": &bintree{nil, map[string]*bintree{
			"gohan_extension.yaml": &bintree{etcExtensionsGohan_extensionYaml, map[string]*bintree{}},
		}},
		"schema": &bintree{nil, map[string]*bintree{
			"core.json":  &bintree{etcSchemaCoreJson, map[string]*bintree{}},
			"gohan.json": &bintree{etcSchemaGohanJson, map[string]*bintree{}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			"swagger.tmpl": &bintree{etcTemplatesSwaggerTmpl, map[string]*bintree{}},
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