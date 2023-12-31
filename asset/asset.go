// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/2013.json (2.718kB)
// data/2014.json (1.883kB)
// data/2015.json (2.133kB)
// data/2016.json (2.129kB)
// data/2017.json (1.944kB)
// data/2018.json (2.26kB)
// data/2019.json (2.053kB)
// data/2020.json (2.124kB)
// data/2021.json (2.304kB)
// data/2022.json (2.465kB)
// data/2023.json (2.379kB)

package asset

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _data2013Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\xdf\x6a\xe2\x40\x14\xc6\xdf\x65\xae\x33\x30\x33\x89\x7f\x92\xb7\x09\x6c\x74\x17\xdc\x5d\x88\x59\x96\x45\xbc\xd9\x62\x89\xb6\xda\x5e\xb4\xd4\x22\x5e\x14\xb4\x2d\x14\xe9\x45\x6f\x44\xe8\xdb\x38\xd1\xbe\x45\x51\x4f\x12\xe3\x24\x27\xa9\xe0\xfd\xfc\xce\x97\x73\xbe\xf9\xe6\xa4\x45\x18\xa7\x8c\x13\xab\x45\xbe\xff\x6e\xfc\xf8\x66\xff\x23\x96\xe7\xfe\x71\x34\xf2\xcb\xfe\xe9\x10\x8b\xc8\xce\x59\x70\xf7\x48\x34\xf2\xd7\xae\x3b\xc4\xd2\xdb\xda\x96\x10\x85\x09\x01\x84\xfe\x65\xa2\x94\x20\x6a\x76\xa3\xe9\x68\xc4\xae\x79\x8e\x9b\xca\xcb\xeb\xc1\xfa\x61\xba\xba\x9a\x45\x85\xb8\x46\x3c\xdb\xad\x3b\x5e\xac\x01\xa5\xcb\x27\x2a\x2d\x28\x33\xb3\xfb\xfc\xb8\x9f\xc8\xc9\x6d\xdc\xa7\x46\x5c\xa7\xe9\xc1\x50\x05\xe5\x2c\x1b\x0d\x86\xd3\x75\xef\x7f\xd2\x06\x41\x39\x66\x9c\x3f\x5e\x2e\x2e\x15\x02\x33\xce\x1f\x2f\xe7\x5d\x85\xc0\x8c\xf3\xc7\x72\x34\x4a\x1a\x27\x28\x37\x72\xbe\xea\x46\x21\x4a\xb8\x46\x67\xa6\x10\xa9\x0e\x26\x86\x95\x6a\x5a\xc2\xe3\xc8\x41\x98\x2e\x94\xae\x14\xbd\x1c\x88\x4e\x5a\x69\x83\x32\x64\x32\xc1\xbc\x13\x0c\x07\x8a\xc9\xc6\x61\x0e\x72\x20\x01\x50\xf9\x18\xa8\x78\xe7\x61\x89\x9c\xe6\x23\xa5\x9d\x80\xc0\x04\x92\x1e\xca\xde\x9b\xec\x3d\x6f\x14\xba\x7d\x2c\x7b\xe1\xb1\x50\xa1\x7a\x3a\x85\x30\xae\x1c\xa4\x90\xa4\xef\x43\x60\x66\x48\xc3\xac\x75\x24\xec\x2a\x2d\x12\xda\x25\xfc\xc5\xce\xa1\xcb\x94\x15\x1f\xd2\xea\xe5\x55\xf6\xfd\xdc\x21\x45\xc7\xa2\x21\xe9\x55\xd0\x32\x4f\xa7\xb5\x53\xc0\x9e\xcd\xfd\xef\x3a\x98\x85\xc1\x00\x47\x46\x99\x82\x03\x84\x3c\xa3\x2a\xb4\x0d\xb2\x49\x39\x72\x61\x96\xf3\xd9\xea\xe9\x42\x51\x32\xa9\x40\xba\xcb\x86\x90\x9e\x54\x08\x3e\x4f\x88\xa2\xe9\x8f\x4a\xe0\xe9\x8f\x95\x40\xa0\xf8\x55\x90\xa3\x77\xb9\x38\xcf\xcf\x66\x78\x6c\xa3\xc0\x19\x1e\x8b\xe8\xec\x7e\xdf\x1b\x08\x5b\x89\x99\x10\xb6\x15\x15\x48\x00\x84\x2d\xc6\x4c\x08\xdb\x8d\x99\x10\xf2\xfc\x67\x43\x95\x23\x20\x5e\xf8\xd6\xc4\x9e\xe2\x7f\x53\xb1\xa7\xed\xcf\x00\x00\x00\xff\xff\xbf\x2c\xf5\x35\x9e\x0a\x00\x00")

func data2013JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2013Json,
		"data/2013.json",
	)
}

func data2013Json() (*asset, error) {
	bytes, err := data2013JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2013.json", size: 2718, mode: os.FileMode(0644), modTime: time.Unix(1691332962, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xea, 0x37, 0x1f, 0x65, 0xad, 0xb8, 0x41, 0xe6, 0x4f, 0x8, 0x5a, 0x2, 0xe6, 0x91, 0xd9, 0xda, 0x83, 0x1, 0x14, 0xa3, 0x46, 0x55, 0x43, 0x54, 0x48, 0xd7, 0x9, 0x2e, 0xd6, 0xc5, 0x95, 0xd1}}
	return a, nil
}

var _data2014Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\xd1\x4a\x02\x41\x14\x86\xdf\xe5\x5c\x6f\xe0\x8e\xab\xe9\xbe\xcd\x40\xa3\x05\x56\x60\x1b\x11\xe2\x4d\x62\xa8\xa5\x75\x11\x24\x48\x17\x5d\x58\x41\x48\x17\xdd\x88\xd0\xdb\x38\xab\x8f\x11\xe6\x34\x26\x67\xe6\xdf\xbc\x5c\x38\xdf\x39\xff\xce\xff\x9f\xd3\xa0\x5c\xb8\x97\x0b\x29\x6e\xd0\xe1\x69\xed\xe8\x40\x5e\x52\x9c\xd4\xcf\x55\x40\x27\xf2\x58\x51\x4c\xba\xdd\x4a\x1f\x5f\x28\xa0\x0b\x59\x55\x14\xe7\x9b\xc1\x8a\x10\xc5\x2d\xa2\x22\x6b\x67\x2a\x20\x59\x49\x54\xdd\x7e\x99\x06\xe9\x70\xbc\xec\x5d\xe9\x6e\x7f\xf9\x3c\x5e\xdc\x4d\x6c\xa7\x30\xa0\x44\xd6\xab\x2a\xb1\x35\xb4\xee\x9d\x07\x6a\x4c\xe1\x96\x1a\x81\xf5\x77\x9e\xe6\xb3\x5b\x46\x08\x4c\x4c\xbb\x8c\xc8\x43\x42\x8f\x46\x96\x10\x86\x88\x32\x54\x3d\x30\xa2\x80\x67\xb4\x27\x8c\x28\x66\xfc\x47\x8b\x11\x25\x97\x73\xdb\x5e\xdd\x0f\xb8\x57\xc6\xdb\xf5\x04\x97\x71\x11\x94\x9f\x4e\xdb\xe9\x70\xc0\xbc\x8b\xe0\x1f\x70\x48\x18\x68\x7f\x67\xa8\x80\x53\xd2\xfb\xd4\xbd\x37\x26\xaf\x80\x83\xc2\xa0\xdf\x49\x28\x2b\x5e\x28\x02\x2b\xe5\x6e\xe1\x34\xca\x7a\xb3\x99\xb4\x1e\x80\xf6\x6a\xf1\xfe\xa1\xfb\x1d\xa6\xaa\x08\x1f\xcd\x0f\x81\x47\xe3\xd0\xcf\x4b\x97\x61\x10\xe6\xd3\xc9\xe2\xf5\x86\x4d\x2a\xc3\x20\xf8\xa1\xd2\x2e\x90\x91\x27\x9c\x7b\xe3\xbc\x78\x7a\xf4\xa5\x67\xd7\x99\x47\xcf\x96\xad\xfc\x09\x73\x38\x9f\xb6\xf6\xaf\xac\x15\x84\xf2\xe9\x85\x50\x3e\x19\x24\x0c\x84\xce\x99\x17\x42\x17\xcd\x0b\xa1\xa3\xe6\x85\x40\x12\xbc\x50\x18\xfe\x7b\xe7\xac\xa7\x78\xe7\x36\x9e\x36\xbf\x03\x00\x00\xff\xff\x79\xa2\xa0\x08\x5b\x07\x00\x00")

func data2014JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2014Json,
		"data/2014.json",
	)
}

func data2014Json() (*asset, error) {
	bytes, err := data2014JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2014.json", size: 1883, mode: os.FileMode(0644), modTime: time.Unix(1691332960, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8f, 0x25, 0xf9, 0x1, 0x2f, 0xf0, 0xdb, 0x92, 0x74, 0x3d, 0x93, 0x3b, 0xdd, 0xeb, 0xd1, 0x5e, 0xd2, 0xb6, 0x7d, 0x83, 0x3b, 0xb5, 0x6f, 0xa3, 0x39, 0xc, 0x1d, 0x29, 0x19, 0x92, 0xfd, 0xdd}}
	return a, nil
}

var _data2015Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xcf\x4e\xea\x40\x18\xc5\xdf\x65\xd6\x6d\xd2\x19\x4a\x2f\xf4\x6d\x9a\xdc\xc2\xbd\x09\xf7\x9a\x94\x1a\x63\x08\x1b\x09\xa6\x54\x41\x17\x1a\x30\xc8\xc2\x04\xfc\x13\x43\x5c\xb8\x21\x24\xbe\x0d\x53\xf0\x2d\x0c\x32\xb4\xb6\x9d\x39\x82\xfb\xfe\xce\xe9\xf7\x9d\xf3\x4d\x83\x18\x54\x37\x28\xb1\x1b\xe4\xcf\x41\xed\xef\x6f\xe7\x98\xd8\xbe\x77\xe8\x6a\xe4\xbf\xf3\xcf\x25\x36\xe1\xed\x56\xd4\xbf\x27\x1a\x39\x72\xaa\x2e\xb1\x0b\x4d\xed\x93\x60\x3b\x13\x4c\x10\x85\xbd\x09\x33\x45\x54\x9c\x5a\x3d\x8b\xf0\xcb\xde\xea\x6e\xb2\xbc\x98\xc6\x2c\xd5\x88\x53\xf1\x5d\x2f\xfe\xde\x77\xbc\xaa\xeb\x27\x26\x6b\x6d\xa6\xd3\xa2\x4c\x3b\x4d\x0a\xa7\x68\x30\x59\x85\x27\xbc\xd3\xcd\x3b\xc5\xda\x9b\x6f\xb6\xda\x25\xf5\xa4\xef\x37\x63\x3e\xbe\x4e\x4f\xca\x74\x5a\x56\x13\x42\x3a\xb5\x7f\xa6\x33\x03\x6c\x33\x18\x2d\xe6\xe7\x39\x02\x65\x1c\x8c\x16\xb3\x4e\x8e\x40\x19\x07\x23\x3e\x1c\x66\xe7\x60\x28\xe3\xf5\x5f\x5d\xe5\x08\x13\x7b\xb4\xa7\x39\xa2\x04\x5a\x21\xb2\x02\xad\xd8\x38\xc8\x82\x33\xb3\x85\x4b\xc7\x30\x6b\x47\x83\xde\xd7\x24\x98\x80\x8a\x3f\x81\xac\xbd\xa1\x22\xbe\xd2\xf0\x95\x87\x8f\x32\x08\x85\xa8\x84\x50\x8e\x72\xc8\x82\x95\x5c\x3e\xbf\xf0\x6e\x90\x82\x34\xe2\xb9\x75\x9f\xd8\x66\x41\xe0\x60\x3a\x09\x2e\x20\x30\x9d\x02\x2a\xc3\xe9\xa2\xb0\x1f\x05\x83\x55\xeb\x96\x07\x4f\x51\x7f\x92\x9c\xc4\xf6\x77\x99\x25\x44\x50\x57\x14\x22\x5b\x7f\xd4\x98\x6f\x50\x0b\x3c\x5b\x48\x48\x7a\x12\xc9\x15\x64\x5c\x37\x66\x0c\x94\x74\x31\x9b\x2e\x1f\xce\x24\xcb\x65\xbf\xf6\x81\xd6\xd9\x53\x03\x37\x7b\xf8\xc6\xe7\xa7\x32\x08\x35\x5b\x09\xa1\x66\xe7\x20\x26\x20\xf4\x48\x29\x21\x90\xb2\x1a\x02\x2b\x57\x43\x60\xe5\x4a\x88\x1a\xbb\x36\x29\x96\xc0\x1d\x4a\x9c\x9a\xcd\x8f\x00\x00\x00\xff\xff\x59\x63\x1b\x14\x55\x08\x00\x00")

func data2015JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2015Json,
		"data/2015.json",
	)
}

func data2015Json() (*asset, error) {
	bytes, err := data2015JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2015.json", size: 2133, mode: os.FileMode(0644), modTime: time.Unix(1691332811, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7a, 0x2c, 0xb1, 0x12, 0x2b, 0x14, 0x61, 0x1a, 0x28, 0xbe, 0xd0, 0xc6, 0x41, 0xd9, 0x91, 0xf3, 0xc5, 0x75, 0x34, 0x75, 0xdd, 0xc0, 0x67, 0x8e, 0xdd, 0xb, 0xce, 0xdf, 0x58, 0xa0, 0x97, 0xb2}}
	return a, nil
}

var _data2016Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xd6\x51\x4b\xc2\x50\x14\x07\xf0\xef\x72\x9e\x37\xd8\xbd\x53\x73\xfb\x36\x83\xa6\x05\x56\x60\x8b\x08\xd9\x4b\x61\x4c\x4b\xeb\xa1\xc8\x10\x1f\x02\xad\x20\xa4\x87\x5e\x44\xe8\xdb\x78\xa7\x7d\x8b\xb0\x2e\x77\x6d\x77\xe7\x34\x83\xde\xfd\x9d\x73\x8f\xff\x73\x2f\x6b\x81\xc5\x4c\x8b\x81\xdb\x82\x9d\x83\xc6\xee\xb6\x77\x02\x6e\xd0\x3c\xf2\x0d\xd8\xf7\xf6\x7c\x70\x41\xb4\xcf\xe2\xbb\x47\x30\xe0\xd8\xab\xfb\xe0\xda\xa1\xf1\x25\x78\x61\xc1\xa5\xb0\x37\x13\xdc\xb4\x2a\x29\x51\xf3\x1a\x87\x09\x89\x07\x93\x55\xf7\x54\x74\x7a\xab\x87\xc9\xf2\x6a\xaa\x2c\x33\xc0\xab\x05\x7e\x53\xfd\x3e\xf0\x9a\x75\x3f\x50\x02\x64\xed\x2d\xfc\x34\x1f\xf7\x63\x31\xbe\xd5\x4e\x53\xc5\x85\x2c\x9d\xfa\x8f\xb8\x69\x39\xc4\xc4\xd1\x68\x31\xbf\xcc\x0a\x66\xd1\x62\xd6\xd1\x04\x95\x5c\x34\x12\xc3\x61\x76\x0e\x46\x25\xb7\x3e\xd5\x8d\x26\xa8\xe4\xa2\x91\x68\x4f\x35\x51\xfa\x3d\xb9\xeb\x3e\x9a\xdc\x77\x87\xbc\xe0\x4a\xe4\xe2\xc5\xb3\x76\x3c\xe8\xff\x4c\x82\x4b\x44\x4c\x80\xa3\xd2\x1f\x90\x4d\x25\xd8\x7d\x13\xdd\x67\x0d\x95\xe9\xeb\x87\x22\x2a\xc7\x7c\x54\x21\x57\x72\xf9\xf2\x2a\x7a\x51\x0e\xa2\xb6\x12\x47\xc4\x4c\x38\xe2\x79\x7b\x93\xda\x8a\x6c\x89\xdc\x3d\x52\xab\x93\x74\x5a\x37\x70\x4c\x56\xc6\x4f\xb5\x98\x4d\x97\x4f\x17\xda\x3d\x76\x4c\x56\xd9\x04\x71\x89\x88\x07\x06\x47\xd5\xa2\xf3\xab\x12\xf4\xfc\x49\xa7\xd0\x00\x66\xd1\x9b\x36\x7c\x17\xf3\xf3\xec\xfc\x6b\x44\x6d\x1a\x8a\xa8\x47\x43\x43\x5c\x22\xe2\xca\xe1\x88\xc8\x14\x47\x44\xa6\x38\x22\x32\xc5\x51\xe1\x4c\x55\x09\x3a\xd3\xa4\x93\x6c\xe0\xfc\x6b\x03\x6e\xda\xc5\xbf\x0e\x78\x18\x7e\x06\x00\x00\xff\xff\x65\xb9\x78\xd1\x51\x08\x00\x00")

func data2016JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2016Json,
		"data/2016.json",
	)
}

func data2016Json() (*asset, error) {
	bytes, err := data2016JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2016.json", size: 2129, mode: os.FileMode(0644), modTime: time.Unix(1691332809, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x49, 0xc2, 0x65, 0xda, 0x1c, 0x9d, 0x5d, 0xdd, 0xee, 0xf3, 0x3b, 0xd7, 0x0, 0xf8, 0xf7, 0xf3, 0x10, 0xcd, 0x46, 0xae, 0xde, 0xb, 0xbd, 0xa5, 0x59, 0x86, 0x46, 0x85, 0x28, 0x6a, 0x91, 0x62}}
	return a, nil
}

var _data2017Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xd5\xcb\x4a\xf3\x40\x14\x00\xe0\x77\x39\xeb\x14\x92\x49\xfa\xff\x36\x6f\x33\x60\x5a\x85\xaa\x10\x23\x22\x25\x1b\xa5\xa2\xd1\x56\x11\xc4\x8a\x2b\xa1\x55\x41\x8a\x0b\x37\x45\xf0\x6d\x9c\xa4\xbe\x85\xe4\xe2\xb4\x99\xf4\x1c\x3b\xd9\xcf\x37\x39\xd7\x4c\x0f\x4c\xab\x61\x5a\xe0\xf6\x60\x6b\xaf\xbb\xbd\xc9\x8f\xc0\x0d\xfc\x03\xcf\x80\x5d\xbe\xe3\x81\x0b\xa2\x7f\x12\xdf\x3d\x81\x01\x87\xbc\xe3\x81\x6b\x87\x46\x26\xd8\xda\x82\xe5\x82\x95\x45\x9b\x77\xf7\x17\x24\x1e\x4d\xe6\xd1\xb1\x38\x1f\xcc\x1f\x27\xc9\xd5\x54\x5a\xcb\x00\xde\x0e\x3c\x5f\x9e\x0f\xb8\xdf\xf1\x02\x29\xa0\xb8\xfb\x3f\x1e\xcd\xf7\xfd\x58\x8c\x6f\x2b\xd1\x6c\xe0\xa2\xb8\x5a\xc9\x98\xb5\x74\x85\x6d\x6a\x0b\xa2\x0f\x8a\xc8\xf2\x60\x64\xe7\x10\x41\x74\x0e\x11\xce\xdf\x9d\xbb\x1e\xa2\x9d\xcb\xbf\xb0\xaa\x71\x8e\x1a\xbe\x72\xf5\xac\x1f\x8f\x86\x5a\x43\x91\x89\xa5\xeb\xa9\x5c\xe5\xd9\x52\xba\x4e\xc3\xb4\xeb\x20\xa7\x06\xa2\x46\x4a\x44\xef\x22\x7a\x59\x81\xa8\xa9\x42\x50\x93\x5e\x70\x0c\x29\x5b\x55\x6e\x4e\xf2\xfa\x26\x06\x67\x9a\x4b\x2b\x11\x14\x5f\x20\xb6\x70\x71\x56\x0d\x8b\xa8\x1a\x8a\xa8\xaa\x21\xa8\xa5\xa2\x72\x01\xbe\x66\xd3\xe4\xf9\x42\xdc\x5c\x8a\x87\x4f\xf1\x71\xaa\x53\x86\x9c\xa6\xb5\xfb\xd5\x45\x49\x2c\x93\xee\x94\x3c\xbb\x1c\x68\x8a\xa8\xbf\x31\x8a\x88\x39\xc7\x11\x31\xe7\x32\xad\x0a\x6a\xd6\xf9\xd2\xbf\x3a\x88\x78\x08\x70\x44\x4c\x21\x82\x18\xbd\x86\xd5\x07\x30\x15\xeb\x3f\xb2\x2c\x0c\x7f\x02\x00\x00\xff\xff\x2b\x5f\xf0\xf5\x98\x07\x00\x00")

func data2017JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2017Json,
		"data/2017.json",
	)
}

func data2017Json() (*asset, error) {
	bytes, err := data2017JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2017.json", size: 1944, mode: os.FileMode(0644), modTime: time.Unix(1691332806, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb7, 0x2e, 0xe8, 0x56, 0x69, 0x30, 0xb0, 0xd7, 0x33, 0xfa, 0x72, 0xc1, 0x25, 0xd9, 0x72, 0xc6, 0xfa, 0x5f, 0xea, 0xba, 0x4d, 0x4d, 0x63, 0xdb, 0xc6, 0x2c, 0xd, 0xd9, 0xd, 0xce, 0x8d, 0xf2}}
	return a, nil
}

var _data2018Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd6\xc1\x4a\xc3\x40\x10\x06\xe0\x77\x99\xf3\x06\xb2\x9b\xb4\x26\x79\x9b\x80\x69\x15\xaa\x42\x1a\x11\x29\xb9\x28\x15\xad\xb4\x7a\xb3\xe0\xc9\x43\x55\x90\xe2\xc1\x4b\x11\x7c\x1b\xd3\xf6\x31\xc4\x64\xb2\x35\x9b\xee\x34\x0b\xde\xf3\xed\x6c\xf6\x9f\x49\x76\x00\x36\xb7\x6c\x0e\xc1\x00\x0e\x4e\x7a\x87\xfb\xe1\x39\x04\x49\x7c\x1a\x31\x38\x0e\x8f\x22\x08\x20\x1b\x5e\x2e\x1f\x9e\x81\xc1\x59\xd8\x8d\x20\x70\x52\x06\xb6\xb0\x78\x55\x74\xc2\x5e\x7f\x43\x96\xd3\xd9\x7a\x74\x91\xdd\x8c\xd7\x4f\xb3\xd5\xdd\x1c\x18\x84\x9d\x24\x8a\xe5\x63\xc5\x4a\x9c\x41\x12\xc6\xdd\x28\x91\x02\x70\xed\x96\x7e\x37\xf8\x60\xb9\x86\x40\xd1\x6e\x2c\xca\xfd\xef\x19\x0b\xcf\x58\xf8\xa6\xef\x21\x6c\x63\x41\x24\xa7\x11\xee\xee\xe4\xee\x27\x32\x39\x99\x15\x46\x58\x54\x50\x83\x63\x10\x47\xfd\x04\x02\xfe\x5b\xc3\xb5\x6c\x2a\xc1\xc5\x70\x39\x9d\x54\x8e\xab\xd4\x6e\xc9\xa9\x38\x6b\x5c\x20\xa2\x12\xd5\x22\x8f\x3a\x8c\x52\x35\x3e\x0c\x59\xa6\x58\x5d\x50\xab\x67\xa3\x8f\x6c\xf4\xaa\xcc\x89\x5a\x00\x89\xac\x20\x51\x59\x81\x68\xb1\xcd\xb3\xca\x4b\x3b\x44\x97\x69\x50\x8b\xfe\x44\xd4\x50\x3e\x01\x6d\x72\x2e\x57\x6f\xef\xd9\xf8\xba\x52\x49\x76\x51\x0b\x39\x11\xe9\x16\x8e\x88\x98\xd3\x3a\xca\x37\xea\x5b\x42\xe8\xd1\xf7\x62\xbe\x7a\xb9\xdd\xba\x51\x5f\x20\x77\x8c\x38\x22\xd7\x04\x39\x95\x11\xf3\xd5\xe4\x95\xde\x7a\xfc\xca\x3e\xaf\x8c\x1a\x2b\x17\xb5\x51\xf6\xd5\x66\xf9\xbf\x3a\x29\x03\x6e\xd3\x6d\xf5\x67\x4f\xca\x21\x38\xa8\x89\xd8\x08\xcd\x51\x13\xa9\x6d\xd1\x88\x88\xd4\xea\x48\x20\x22\xbe\x87\x7a\x44\x0c\x8f\x1e\x11\x23\xa3\x41\x62\x47\x2f\xe5\xff\x7f\x93\x9f\xb9\xbc\x31\x14\xa7\xed\x61\x15\xf2\xb3\x53\xbd\x64\xc8\x01\xf3\x1c\xb4\xcd\x2f\x28\x22\x4d\x7f\x02\x00\x00\xff\xff\x01\xb5\x37\xd6\xd4\x08\x00\x00")

func data2018JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2018Json,
		"data/2018.json",
	)
}

func data2018Json() (*asset, error) {
	bytes, err := data2018JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2018.json", size: 2260, mode: os.FileMode(0644), modTime: time.Unix(1691332805, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x70, 0x9d, 0x71, 0x44, 0xc0, 0x5a, 0x76, 0x77, 0x6d, 0xf1, 0xd3, 0x84, 0xfc, 0x1f, 0x6d, 0xed, 0x8e, 0x49, 0x6e, 0x98, 0x34, 0xb0, 0x72, 0x34, 0xa7, 0x3, 0x31, 0xe2, 0x74, 0x75, 0x6, 0x46}}
	return a, nil
}

var _data2019Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xd5\x5f\x4b\x7a\x31\x18\x07\xf0\xf7\xf2\x5c\xef\xc0\x36\xff\xfc\x3c\xe7\xdd\x0c\x7e\xd3\x02\x2b\x38\x9e\x88\x10\x6f\x0a\xa3\x0c\xad\x9b\xc8\xe8\x2a\xd0\x0a\x42\xba\xe8\x46\x82\xde\x4d\x53\x7b\x17\x61\xce\xe9\x71\x3c\x0f\x2e\xe8\x7e\x9f\xcd\x7d\xf7\x7d\x8e\x4d\xe0\x22\xe2\x02\x92\x26\xec\x1c\xd4\x77\xff\xab\x63\x48\xb2\xf4\x50\x33\xd8\x57\x7b\x1a\x12\x30\xed\xd3\xc9\xed\x23\x30\x38\x52\x35\x0d\x49\xa1\xc5\x80\xcb\x88\xcb\x9c\xa8\xaa\x7a\x63\x45\x26\xfd\xe1\xac\x73\x62\x2e\xba\xb3\x87\xe1\xf4\x6a\x04\x0c\x54\x35\xd3\xa9\x5b\xb6\xd8\x49\x30\xc8\x54\x5a\xd3\x99\x13\x60\xf7\x2e\xfc\xe1\xde\x45\xfc\xa6\x5f\x77\x03\x33\xb8\x71\x37\x95\x56\x94\x70\x61\xb7\xde\xcc\xa6\x1c\x2c\xfe\x05\x8b\xca\xd6\x62\x79\x8f\x38\x54\x08\x1e\x26\x8a\x74\x56\xe3\xf6\xa4\xdf\xf3\xae\x52\xa4\xe3\xf2\xd0\xf2\x24\x2a\x31\x0c\xc9\x0a\x51\x2c\xd3\x79\x33\x9d\xe7\x8d\x6e\xb9\x36\xe5\x4b\xe6\xba\xe5\xd0\x4f\xbd\x4a\xf4\x20\xb9\xb5\x2e\x00\x06\xa9\x6e\x64\x90\x08\xab\x65\x88\x96\x16\x15\x7e\x83\x88\x31\xc0\x51\x69\xab\xf8\xae\x7b\x68\x7c\x8b\x73\x90\xf4\xca\xe4\xa3\x4e\x5f\x5e\x4d\xf7\xdc\xab\x4f\x99\x9c\x04\x1f\x49\x8b\x88\x61\x40\x50\x1c\x09\x22\xe9\xcf\xf1\x68\xfa\x74\xe9\xfd\xbc\x38\x12\x44\xd2\x3e\x5a\x9e\x44\xcc\x11\x8a\x64\x4c\x3d\xcf\xfd\x87\x79\x3f\x0b\xfc\x72\x3a\x34\x7f\x1f\xc1\xe9\x76\xbb\xb5\xeb\x01\xcc\x11\x55\x6a\x1f\xad\x8f\xc4\x5c\x53\xed\x46\x8f\xa4\xda\xed\x21\x69\x11\x91\x39\x8e\x88\x6f\x17\x8e\x88\x9a\xa3\x48\x50\x7f\xb8\xab\xd7\x5d\x1b\xbe\xdc\xcc\x91\x8f\xdb\xfa\x0e\x00\x00\xff\xff\xae\x26\x67\xab\x05\x08\x00\x00")

func data2019JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2019Json,
		"data/2019.json",
	)
}

func data2019Json() (*asset, error) {
	bytes, err := data2019JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2019.json", size: 2053, mode: os.FileMode(0644), modTime: time.Unix(1691332803, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x69, 0xc1, 0x9, 0xd1, 0x2b, 0x45, 0xa0, 0x3b, 0x5, 0x20, 0x3e, 0xab, 0x4e, 0x18, 0xa2, 0x72, 0xb0, 0xe8, 0xd5, 0xe1, 0x43, 0x4c, 0x5e, 0xa3, 0x6d, 0xb5, 0x38, 0x4b, 0x67, 0xe4, 0xe3, 0x93}}
	return a, nil
}

var _data2020Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\xdd\x4a\x32\x41\x18\xc7\xef\xe5\x39\x5e\x61\x66\xd6\xcf\xbd\x9b\x85\x77\xf5\x7d\xc1\xb7\x60\xdd\x88\x90\x85\x28\x0c\xb5\xb4\x0e\x8a\x0c\xe9\x20\xd0\x0a\x42\x3a\xe8\x44\x84\xee\xc6\x59\xed\x2e\xc2\x1c\xc7\x9d\x99\xe6\xd9\xec\x7c\x7e\xf3\x3c\xfb\xff\xd8\x69\x02\xa1\x39\x42\xc1\x6b\xc2\xdf\xfd\xfa\xbf\x3f\xfe\x11\x78\x51\x78\x10\x38\xb0\xe7\xff\x0f\xc0\x03\xde\x3a\x4d\x6e\x1f\xc1\x81\x43\xbf\x16\x80\xe7\xc6\xce\x8a\xa0\x15\x85\xa8\xfa\xf5\xc6\x16\x49\x06\xe3\x65\xf7\x84\x77\x7a\xcb\x87\xf1\xe2\x72\x02\x0e\xf8\xd5\x28\x08\xe5\xb1\xf5\x4d\xd4\x81\xc8\x0f\x6b\x41\x24\x09\x58\xdf\xcd\xf2\xf6\x6d\x3e\xee\x46\x7c\x74\x23\xb7\x61\x82\x28\x20\xfb\xb7\xef\xe7\xd3\x63\x7d\x7f\x56\xc4\x89\xd9\x85\x41\x94\x32\x66\x74\x0c\xa2\x8c\x12\x7c\x38\x34\xbe\xa3\x92\xb1\xd5\xb5\x4e\xb8\x04\x9f\xd1\x9a\xa8\x04\xd3\xbd\x16\x96\x08\x83\x14\x5e\xd8\x78\xd5\x97\x36\xea\x7e\x49\x2b\x57\x57\xe7\x73\x04\x31\x2e\x99\xb6\x92\x41\x3f\x0d\xb9\x02\x42\xbc\x33\x21\x26\x20\xc4\x3e\x2b\xa4\x79\xae\x66\x96\x77\xdf\x78\xf7\x79\xc7\xd8\x4a\xe8\x2b\xb9\x05\xbc\x47\xf2\xac\x22\x40\x21\x47\xd8\x2e\x10\x13\x90\xfb\x1b\x08\xf1\xc7\x0e\x61\xdd\xb2\x42\xd8\xef\x61\x2b\x75\x2a\x5a\x4a\x00\xb3\x94\x2e\xa2\x8d\x5f\xbc\xbc\xf2\x5e\xdb\x50\xba\x88\x96\xde\x84\x98\x80\x90\xde\xdb\xa1\xf2\x4f\x4b\x26\xaf\x48\x8b\x61\x7e\xff\x76\xd2\x6a\x40\x45\xdf\x4a\x1d\xa0\xc9\x3d\x7c\xe7\xb3\x33\x2d\xd9\xdf\x28\xbc\x39\x06\x0e\x84\x41\x23\x02\xaf\x14\x3b\x40\x09\x1a\xea\xf9\x74\xb2\x78\x3a\x57\xa4\xde\xd0\x94\x0a\x1c\x8b\x77\x6a\xa6\x8e\x0b\x1a\xcb\xb9\x49\x0b\x08\xcb\x79\xf6\x48\x2c\xf0\x06\xcd\x04\x84\xbd\x27\x56\x08\x7b\x52\xac\x10\xf6\xaa\xd8\x20\x4a\xb0\x3e\xca\x80\xec\xd0\x47\x23\x2d\x2c\x8e\x3f\x03\x00\x00\xff\xff\x12\x60\xb5\x94\x4c\x08\x00\x00")

func data2020JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2020Json,
		"data/2020.json",
	)
}

func data2020Json() (*asset, error) {
	bytes, err := data2020JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2020.json", size: 2124, mode: os.FileMode(0644), modTime: time.Unix(1691332801, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8d, 0xd5, 0xe4, 0x2e, 0xd5, 0x4d, 0x59, 0x4d, 0xbb, 0x8a, 0x7a, 0x90, 0x9f, 0x78, 0xa5, 0xe9, 0xdc, 0x3f, 0x31, 0x3, 0x6f, 0x5c, 0x8c, 0xf4, 0x70, 0x7c, 0x5d, 0xa6, 0xe1, 0x30, 0x8a, 0x84}}
	return a, nil
}

var _data2021Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x95\xc1\x4a\xc3\x40\x10\x86\xdf\x65\xce\x09\x64\xb7\x4d\x6a\xf2\x36\x01\xd3\x2a\x54\x85\x34\x22\x52\x0a\xa2\x54\xda\x6a\xab\x07\xc5\x4a\xe9\x41\x68\x55\x90\xe2\xc1\x4b\x29\xf8\x36\xdd\xb4\xbe\x85\xa4\xd9\x6e\x4c\xd6\x9d\xa4\xf7\xf9\x76\x66\xfe\xf9\x7f\xb6\x09\x06\xd1\x0d\x02\x4e\x13\x0e\x4e\xea\x87\xfb\xee\x39\x38\x81\x7f\xea\x69\x70\xec\x1e\x79\xe0\x00\x6b\x5f\x85\x4f\xaf\xa0\xc1\x99\x5b\xf3\xc0\x29\x69\xe0\x7b\x8d\x00\x9c\x0a\x69\x69\x1b\x96\x16\x66\xe9\x96\xdd\xa2\xa5\xe2\x68\x44\x50\xdd\xa8\xa4\x88\xaa\x5b\x6f\x24\x48\x38\x9c\xae\x7b\x97\xac\xdb\x5f\xbf\x4c\x57\x77\x33\xd0\xc0\xad\x06\x9e\x2f\xca\xe2\x97\x88\x06\x81\xeb\xd7\xbc\x40\x10\x10\xbf\x4d\x10\x11\x7e\x9e\x27\x6c\xf2\x98\x9d\x86\x60\xab\x77\xc6\xcb\xf9\x45\x22\x1b\x27\xb0\x8d\x3b\xe3\xe5\xe2\x56\x22\xca\x39\x3d\xba\x12\x61\xa2\x04\x1b\x8d\xa4\x3d\xac\x9c\xa9\x1e\x24\xa2\x82\xf7\x68\xcf\xb2\x04\x35\xf2\x2f\x77\x3f\xc8\x5e\x2e\x7e\x18\x3d\x5c\x19\xb5\x51\x38\x6f\x87\xc3\x41\x54\x9b\x9a\xa7\xac\x1b\x88\xae\x32\x54\xe2\x10\x22\xad\xb2\x13\x35\x91\xcd\x59\xef\x8b\xf5\xde\x77\xb4\xad\x80\x36\x02\x98\x78\x7c\x45\x6d\x6a\x17\x13\xcf\xad\x04\x51\x0e\x61\xfe\x55\x42\x98\x85\x95\x10\xe6\x62\x25\xb4\x57\x48\xea\xe2\x3e\x4b\x2b\x6d\xa1\x89\x5f\x7d\x7c\xb2\x7e\x47\x9a\xca\x42\x43\xaf\x86\x10\xd1\x64\x68\x73\x53\x5b\x27\xff\xee\x9f\xb6\x13\x7f\x63\x39\x9f\xad\xde\x6e\x32\xc6\x93\x05\x10\x65\xc0\x3b\xd8\xea\xb1\x92\xda\xd4\x2e\x76\x36\xfb\x05\x21\xc4\xd4\x32\xc4\x05\xa0\x56\x61\x01\xd8\xe8\x9b\x2d\xae\x73\x05\x10\x65\x91\x00\xc4\xc0\xb3\x26\x6a\xff\x8e\x15\x41\x58\xd6\x94\x10\x96\x35\x25\x84\x65\x4d\x82\x28\x87\xb0\xac\x29\x21\xec\xd3\x50\x42\xd8\xbf\xa1\x84\x6c\x2c\xd5\xe2\x8c\x3b\xa4\x3a\xb9\x69\xeb\x37\x00\x00\xff\xff\xf6\x2c\xe6\x17\x00\x09\x00\x00")

func data2021JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2021Json,
		"data/2021.json",
	)
}

func data2021Json() (*asset, error) {
	bytes, err := data2021JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2021.json", size: 2304, mode: os.FileMode(0644), modTime: time.Unix(1691332797, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1f, 0xbb, 0x94, 0xc6, 0x8a, 0x30, 0x45, 0x4b, 0xc3, 0x99, 0x86, 0x53, 0xff, 0x9d, 0x8d, 0xee, 0x13, 0xe2, 0x76, 0xe3, 0xe, 0xf8, 0xf4, 0xf4, 0xb3, 0xdd, 0x93, 0x3f, 0x1e, 0xdc, 0x1f, 0xde}}
	return a, nil
}

var _data2022Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\x51\x4b\xc2\x50\x14\xc7\xbf\xcb\x79\x9e\xb0\x7b\xa7\x33\xf7\x6d\x06\x4d\x0b\xac\x60\x2e\x22\x44\x88\xc2\x50\x4b\xeb\xa1\xc8\x10\x1f\x02\xad\x20\xa4\x87\x5e\x44\xe8\xdb\x78\xd5\xbe\x45\xd8\xee\xee\xdc\xbd\xbb\x47\x8d\x7a\xdf\xef\xfc\xcf\xf9\x9f\xff\xd9\x56\x05\x93\x64\x4c\x02\x4e\x15\xf6\x8e\xca\xfb\xbb\xee\x29\x38\x81\x7f\xec\x19\x70\xe8\x1e\x78\xe0\x00\xab\x5f\xcc\x1e\x9e\xc1\x80\x13\xb7\xe4\x81\x63\x19\xe0\x7b\x95\x00\x1c\x52\x33\x7e\x50\xba\x31\x4a\x65\xd4\xfa\x35\x4a\x0b\x09\xb4\xe8\x96\x2b\x31\x3b\xeb\x0e\x17\xad\x73\xd6\x6c\x2f\x9e\x86\xf3\x9b\x11\x18\xe0\x16\x03\xcf\x17\x8f\x85\x25\x89\x01\x81\xeb\x97\xbc\x40\x10\x10\x89\xd8\xa1\x88\x65\xfe\x87\x08\xaf\x8d\x38\xfe\xf5\x38\x60\x83\x7b\xcd\xec\x14\x5f\x56\xa3\x3f\x1d\x9f\xc5\xcb\xe2\x04\xb6\xa3\x46\x7f\x3a\xb9\xd6\xac\x97\xe2\x3b\x5a\x8a\x35\x15\xb1\x2c\x4a\xb0\x5e\x2f\x9e\x8c\x13\xb9\x35\xed\xdd\x29\x84\x8d\x6b\xd4\x47\x49\x22\x2b\x5b\x20\xed\x72\x5c\x9f\x75\x3b\xdb\xae\x33\x82\x80\x2b\x20\x4e\xc5\xcf\xca\x6d\x21\x66\xe9\x21\xc4\x2f\x15\xb2\x42\x88\x66\x11\x03\x58\xeb\x83\xb5\x5e\xb7\x34\x40\x40\xe2\x6e\x48\x21\xd4\x92\x0e\x27\xb9\x9e\x15\x4a\xca\x37\xcd\x2d\xf1\x1c\x1e\x70\x05\xb7\x38\x84\x65\x5c\xd5\xe4\x10\x96\x6e\x2d\x84\x05\x5c\x0b\xe5\x37\xb2\xff\xb6\x23\xdb\x1f\x96\xc7\xdc\x5f\x0a\xd8\xe8\x28\xf3\xb7\x77\xd6\x6e\x28\xa6\xd9\xe8\x28\x2a\x44\x39\x84\xc4\x4f\x03\x15\x32\x04\x89\xc4\x74\x3c\x9a\xbf\x5c\x25\xda\x13\x89\x8a\x70\x24\x12\x2a\x1e\x41\x48\x24\xd2\x21\x62\xe2\xe1\xeb\x7d\xb2\xc9\x65\x6a\xa3\x11\x8d\xa5\x50\x4f\x13\x4e\x63\x71\x5c\x4f\x63\xb9\x54\xe8\xa8\x61\xec\xdd\xab\x85\xb0\xd7\xaf\x16\xca\x6f\x05\xc9\xc3\xed\xa4\xdd\x4f\xfa\x8d\xc8\x45\x57\xaf\x2a\xbe\x1f\x21\xc9\x05\x52\xff\x28\xfe\x4c\x80\xa2\x5f\x7c\xdd\xdf\x8e\x6d\xd6\x6a\xdf\x01\x00\x00\xff\xff\x85\x55\xf5\x32\xa1\x09\x00\x00")

func data2022JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2022Json,
		"data/2022.json",
	)
}

func data2022Json() (*asset, error) {
	bytes, err := data2022JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2022.json", size: 2465, mode: os.FileMode(0644), modTime: time.Unix(1691332794, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3f, 0x37, 0xff, 0xc0, 0xe2, 0xe1, 0xbb, 0x25, 0xa6, 0x86, 0xe9, 0x2a, 0xf1, 0x49, 0x16, 0x35, 0x42, 0x9b, 0x5a, 0xd1, 0x23, 0x36, 0xb9, 0x42, 0xdb, 0x95, 0xe5, 0x28, 0x44, 0x85, 0x86, 0xc3}}
	return a, nil
}

var _data2023Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\xd1\x4a\x3a\x41\x14\xc6\xdf\xe5\x5c\xaf\xb0\x3b\xeb\xae\xba\x6f\xb3\xf0\x5f\xfd\x07\x56\xb0\x6e\x44\x88\x10\x85\xa1\x96\xd6\x45\x91\x21\x5d\x04\x5a\x41\x48\x17\xdd\x88\xd0\xdb\x38\xab\xbd\x45\xa8\xeb\x38\x33\xbb\x73\x46\xa2\xee\xfd\xcd\xf7\xed\x39\xdf\x39\xc7\x3a\x98\x56\xce\xb4\xc0\xab\xc3\xff\xc3\xea\xde\x3f\xff\x04\xbc\x28\x3c\x0a\x0c\x38\xf0\xf7\x03\xf0\x80\x36\xcf\xe3\xfb\x67\x30\xe0\xd8\xaf\x04\xe0\xd9\x06\x84\x41\x2d\x02\xcf\x6a\x18\x2b\x94\xec\x8c\x12\x09\x25\x88\xea\xd7\xc3\x90\x0e\xef\xd2\xaa\x24\x41\x31\xd5\xd6\xe3\x6c\x72\xaa\x34\x4c\x6c\x1c\x9d\x5e\xa9\xd1\xbc\x46\xb5\xad\x46\x1d\x14\xa5\x83\x81\xba\x4c\xae\xc6\xf0\xad\x1a\x2d\xe0\xaa\xcd\xb1\x1a\x2d\x0a\x68\xd9\xaf\xd6\xb6\x6c\xdc\x1f\x2d\x3a\x67\xf4\xa6\xb7\x78\x1a\xcd\xaf\xb7\x8f\x58\x06\xf8\xe5\x28\x08\x37\x52\x91\x1f\x56\x82\x88\x01\x20\x6b\x94\xfe\x5a\x23\x9f\x33\x91\xc2\xc7\x93\x66\xdc\xef\xad\x21\xa9\x6d\x6e\x61\x8d\x4b\x69\x11\x2d\xd2\xce\x07\xed\xbc\x2e\x5d\xb6\xbb\x69\x97\xcc\x18\xfb\x19\x30\xeb\xc9\x3b\x82\x53\xa9\x1a\x62\xb3\xb8\x27\xa4\x7e\x25\x46\x6d\xf3\x47\xf4\x4a\xdb\xc1\x17\x40\x8a\xb6\x65\x1a\x9b\x46\x3d\x8d\x0d\xa4\x9e\x76\x77\x6a\x10\x17\x23\x21\x3d\x78\xb7\xb8\x9d\xe3\xa2\x3b\x67\xfe\xf6\x4e\xbb\xad\x4c\x9f\x1b\x1a\xf9\x4a\x84\xb6\x12\x1a\xd9\x3c\x69\x9a\xc8\xb4\x83\xd4\x88\xe1\x99\xa3\xc6\xca\xc2\xab\x08\x15\xe4\x94\x4a\x68\x86\x67\x93\xf1\xfc\xe5\x32\xf3\x2b\x8b\xf6\x1a\xc7\x42\x8c\xe0\x4b\x71\xcb\xc4\x43\x3c\xf8\xa4\xd3\x0b\x94\xc6\x42\xac\xa7\xb1\x10\xa7\x68\x22\xd3\xd8\x61\xd1\xd3\xd8\x6d\xd1\xd3\xd8\x79\xd1\xd3\x85\xac\x68\x65\x8f\x98\xfc\x28\x1f\xb8\xed\xf8\x71\x92\xa2\x52\xe6\x41\xfa\x7d\x25\x92\xb3\x77\xff\x37\xc4\xca\x51\x74\x1a\x8d\xef\x00\x00\x00\xff\xff\x92\x99\x4f\x65\x4b\x09\x00\x00")

func data2023JsonBytes() ([]byte, error) {
	return bindataRead(
		_data2023Json,
		"data/2023.json",
	)
}

func data2023Json() (*asset, error) {
	bytes, err := data2023JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/2023.json", size: 2379, mode: os.FileMode(0644), modTime: time.Unix(1691332793, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x77, 0x19, 0x4, 0x7, 0x7, 0x54, 0x69, 0xb1, 0x76, 0x8c, 0xf6, 0xad, 0xfe, 0xa5, 0x48, 0x65, 0x19, 0x89, 0x37, 0xf9, 0x73, 0x5c, 0x17, 0xec, 0x3c, 0xde, 0x67, 0x31, 0x6d, 0xaf, 0xe5, 0x96}}
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
	"data/2013.json": data2013Json,
	"data/2014.json": data2014Json,
	"data/2015.json": data2015Json,
	"data/2016.json": data2016Json,
	"data/2017.json": data2017Json,
	"data/2018.json": data2018Json,
	"data/2019.json": data2019Json,
	"data/2020.json": data2020Json,
	"data/2021.json": data2021Json,
	"data/2022.json": data2022Json,
	"data/2023.json": data2023Json,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"data": {nil, map[string]*bintree{
		"2013.json": {data2013Json, map[string]*bintree{}},
		"2014.json": {data2014Json, map[string]*bintree{}},
		"2015.json": {data2015Json, map[string]*bintree{}},
		"2016.json": {data2016Json, map[string]*bintree{}},
		"2017.json": {data2017Json, map[string]*bintree{}},
		"2018.json": {data2018Json, map[string]*bintree{}},
		"2019.json": {data2019Json, map[string]*bintree{}},
		"2020.json": {data2020Json, map[string]*bintree{}},
		"2021.json": {data2021Json, map[string]*bintree{}},
		"2022.json": {data2022Json, map[string]*bintree{}},
		"2023.json": {data2023Json, map[string]*bintree{}},
	}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
