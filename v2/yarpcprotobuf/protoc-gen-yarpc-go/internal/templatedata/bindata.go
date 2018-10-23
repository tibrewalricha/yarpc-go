// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by go-bindata. DO NOT EDIT.
// sources:
// internal/template/.nocover (0)
// internal/template/base.tmpl (470B)
// internal/template/client.tmpl (1.055kB)
// internal/template/client_impl.tmpl (2.094kB)
// internal/template/client_stream.tmpl (2.656kB)
// internal/template/fx.tmpl (2.475kB)
// internal/template/server.tmpl (1.915kB)
// internal/template/server_impl.tmpl (1.615kB)
// internal/template/server_stream.tmpl (1.743kB)

package templatedata

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

var _internalTemplateNocover = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func internalTemplateNocoverBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateNocover,
		"internal/template/.nocover",
	)
}

func internalTemplateNocover() (*asset, error) {
	bytes, err := internalTemplateNocoverBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/.nocover", size: 0, mode: os.FileMode(420), modTime: time.Unix(1539792228, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var _internalTemplateBaseTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8c\x41\x4b\xc3\x40\x10\x85\xef\xfb\x2b\x9e\xa1\x07\x05\x93\xdc\x0b\x9e\xac\x4a\x2f\x6d\xc1\xde\x65\xdc\x4e\xe3\x62\x92\x5d\x36\x6b\xb1\x0c\xf3\xdf\xc5\x6e\xaa\x82\xc6\xdb\xcc\xfb\xde\xf7\xea\x1a\xb7\x7e\xc7\x68\xb8\xe7\x48\x89\x77\x78\x3e\x22\x44\x9f\xbc\x2d\x1b\xee\xcb\x23\xc5\x60\xcb\xc6\x9b\xba\xc6\xe0\xdf\xa2\xe5\x39\x44\xaa\x7b\xd7\x72\xb5\xa2\x8e\x55\x3f\xc9\x62\x8d\xd5\x7a\x8b\xbb\xc5\x72\x7b\x61\x4c\x20\xfb\x4a\x0d\x7f\xf5\x36\xf9\xaf\x1e\xfc\x78\xa9\x1a\x23\xe2\xf6\xc8\xfc\x91\xe3\xc1\x59\x1e\x50\xaa\x1a\xc0\x75\xc1\xc7\x84\x4b\x03\x00\x22\x91\xfa\x86\x31\xcb\xe9\x86\xd2\xcb\x35\x66\xd4\x3a\x1a\x30\xbf\x41\xb5\x3c\xc5\x67\x35\x1b\x19\xab\xa2\x10\xf9\xe1\xa9\x16\xe3\x24\xf7\xbb\x51\xb8\x32\xdf\x9f\x11\x49\xdc\x85\x96\x12\xa3\xb0\xad\xe3\x3e\x15\xa8\x4e\xe8\x37\x79\x72\x5d\x68\xff\xc1\x43\x8a\x4c\xdd\x5f\x85\x81\xe3\x81\xe3\x34\x99\x5c\x1e\xf1\xf4\xf2\xfe\xfd\x9c\x7e\x04\x00\x00\xff\xff\xdc\x7d\xb8\x98\xd6\x01\x00\x00")

func internalTemplateBaseTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateBaseTmpl,
		"internal/template/base.tmpl",
	)
}

func internalTemplateBaseTmpl() (*asset, error) {
	bytes, err := internalTemplateBaseTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/base.tmpl", size: 470, mode: os.FileMode(420), modTime: time.Unix(1539227189, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x4, 0xb8, 0x46, 0xd, 0x49, 0xa6, 0x68, 0x60, 0xd, 0xcf, 0x2, 0xfb, 0x70, 0x8c, 0x9, 0x3d, 0x42, 0xe2, 0x5b, 0xa6, 0x56, 0xea, 0xa3, 0x44, 0x3a, 0x93, 0x36, 0xea, 0x2e, 0xbb, 0xe5}}
	return a, nil
}

var _internalTemplateClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x52\xcd\x6e\xdb\x3c\x10\xbc\xeb\x29\x06\xc6\xf7\xb5\xb6\xe1\xd2\xf7\x00\x3d\x14\x06\x52\xf4\x50\x37\x4d\x7a\xe9\x91\xa1\x57\x32\x11\x99\x64\x48\xaa\x6e\xb0\xe0\xbb\x17\x12\x15\xfd\xd8\x05\x5a\xa0\xba\x88\xdc\xc5\xce\xcc\x0e\x87\xf9\x40\xa5\x36\x84\x85\xaa\x35\x99\xb8\xc0\xbb\x94\x0a\xe6\xb3\x8e\x47\x88\x5b\x5d\x53\x77\xfd\xaf\xb2\xee\xa9\xc2\xcd\x7b\x88\x3b\xa9\x9e\x64\x45\xe2\xa3\xed\x4f\x29\x15\x05\xb3\x97\xa6\x22\x88\x07\xf2\x3f\xb4\xa2\xd0\xc1\x14\x00\xf3\x76\x8d\x5d\x07\x0d\x6d\x22\xf9\x52\x2a\xc2\x7a\x9b\xbb\xdb\x2d\x98\x45\x6e\xa7\x04\x1d\x10\x8f\xd4\x96\xf6\xf2\x44\x29\x21\x64\xb4\xb7\x01\xea\x02\x42\x14\x40\x7c\x71\x34\x9f\x1f\x08\xb8\x00\x5a\xf2\x5e\xd5\x67\x8a\x47\x7b\xc8\xa2\x80\xdc\xd2\x25\xac\x47\x3f\xfc\x10\x3d\xc9\x93\x36\x55\xde\x80\xfc\x58\x18\x67\x30\x2a\x5b\x0e\x25\x40\x59\x13\xe9\x67\x14\xbb\xfc\xdf\x4c\x5a\x1d\x8b\xb1\xf1\x9a\x66\x8a\xda\x7e\x6b\xe6\xca\x7e\x6b\x17\x12\xf7\xf4\xdc\x50\x88\xc8\x9e\xa7\x34\x07\x24\x73\xb8\x18\x16\x42\xbc\x48\xef\x94\xd8\xc9\xba\xfe\xe2\xa2\xb6\x66\x1c\x59\x61\xc9\x2c\x32\xef\xab\x4f\x1b\x90\xf7\xd6\xaf\x06\x2b\xa8\x0e\xf4\x4f\x7b\xfe\xa5\xfa\x3f\x29\x9d\xc1\x04\x67\x4d\xa0\x11\xe7\x4a\xf4\xc4\x88\xe9\xed\x2a\x76\xca\x9a\x10\x7d\xa3\x5a\xba\x69\xf2\xf6\x74\x9e\x86\xe7\xb1\xd1\xf5\x21\x40\xc2\xd0\x19\xdf\x3f\xdc\xdf\xed\x5e\x43\x57\x5a\xff\xfb\x5c\xb6\x21\x2c\x1b\xa3\x2e\xa0\x96\x0a\xfd\x9a\x5d\x61\x03\xeb\x62\x18\x96\x77\xde\x46\xfb\xd8\x94\x7d\x37\xdb\xb0\x9a\xe5\x38\xa7\xd7\x53\x6c\xbc\xc1\x9b\xa1\xf3\xe9\xe4\xea\x94\x38\x74\x8f\x79\x83\x39\xda\x9e\xce\xd3\x57\x5e\xaa\x0d\x98\x9d\xd7\x26\x96\x58\xfc\xff\xbc\x80\xb8\xfd\xba\x6f\x6d\x6c\xc5\x08\x21\x56\xbd\x57\xa3\x75\xe3\x71\xac\xfd\x0a\x00\x00\xff\xff\xd6\x5e\x65\x4c\x1f\x04\x00\x00")

func internalTemplateClientTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateClientTmpl,
		"internal/template/client.tmpl",
	)
}

func internalTemplateClientTmpl() (*asset, error) {
	bytes, err := internalTemplateClientTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/client.tmpl", size: 1055, mode: os.FileMode(420), modTime: time.Unix(1539227189, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5c, 0x33, 0xaf, 0x80, 0x4b, 0xdc, 0x36, 0x4b, 0x61, 0x90, 0x50, 0xc0, 0x17, 0xd9, 0xb0, 0x9b, 0x57, 0x3e, 0xe3, 0x64, 0x1f, 0x78, 0x3b, 0x22, 0x31, 0x91, 0xf1, 0xfa, 0xd7, 0x1b, 0x64, 0x5a}}
	return a, nil
}

var _internalTemplateClient_implTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\xc9\x6e\xdb\x30\x10\xbd\xeb\x2b\x5e\x8c\xb4\x90\x0c\x85\xb9\xbb\xf0\xc9\x68\x8b\x1e\xba\xa0\xee\x3d\x60\xe5\x91\x42\x58\x22\x65\x92\xce\x02\x82\xff\x5e\x90\x94\xf7\xc4\xbd\xa4\x2d\x0a\xe4\x64\x72\xcc\xb7\xcc\x1b\x51\x72\x6e\x41\xb5\x90\x84\x51\xd5\x0a\x92\xf6\x46\x74\x7d\x3b\xc2\x95\xf7\x99\x73\xf7\xc2\xde\x82\x7d\x10\x2d\xc5\xed\x65\xa3\xfa\x65\x83\xc9\x14\xec\x1b\xaf\x96\xbc\x21\xf6\x51\x0d\x2b\xef\xb3\xcc\x39\xcd\x65\x43\x60\x73\xd2\x77\xa2\x22\x13\x69\x00\xe7\x2e\x13\xf9\xa7\xae\x6f\x23\x7c\xb6\xdd\x06\x5c\x38\x71\x3d\x46\x2a\x22\x18\xa0\x8e\xa4\xe5\x56\x28\x89\xf1\x75\x3a\x62\x1f\x7b\x3a\x64\xf2\x1e\xc6\xea\x75\x65\xe1\x32\x00\x61\x43\xbc\xc3\x23\xd7\x7d\xd5\x6b\x65\xd5\xcf\x75\xcd\xe6\xb1\x98\xa8\x33\x20\x52\xdd\x71\x8d\x1b\x38\x37\xd8\xf0\x1e\x53\xe4\xe3\x23\xee\x22\x97\xa2\x2d\x92\xb9\xa1\xad\xcf\x64\x6f\xd5\xc2\xc4\x9e\x42\x59\xd4\xe0\x72\xb1\xe9\x26\x29\x09\xd9\xa4\xfe\x49\xef\x0a\x57\x03\x04\xa8\xd7\xb2\x42\x5e\xe1\x44\x2d\xd8\xf9\xc2\x3b\xf2\x3e\xaf\xec\x03\x2a\x25\x2d\x3d\x58\x36\x4b\xbf\x25\x54\x6f\x0d\x18\x63\xb1\x3b\x36\xe3\x6d\xfb\xb5\x0f\xf9\x14\xc8\x9d\x3b\xe8\xd2\xfb\x12\xa4\xb5\xd2\xc5\x90\x4b\xcc\x26\xd6\x42\xf6\x15\x4b\x39\x45\x8e\x84\x0b\x8a\x25\x9c\xeb\xb5\x90\xb6\xc6\xe8\xcd\x6a\x84\xc1\x4c\x12\x66\x8c\x15\x5b\x2a\x51\x47\xaa\x8b\x29\xa4\x68\xf7\x24\x00\x4d\x76\xad\x65\x28\x47\xb5\xed\x3f\x3e\x3b\x3a\xf1\xf6\xc8\x72\xca\xc0\x25\x63\x13\x18\x5f\x06\x92\x6c\x1f\xed\x1c\xb5\x86\x82\xf8\x49\xdc\xaf\xe9\xbe\x64\xba\x7f\xe0\xd9\xd5\xb4\x0a\xa0\x46\xfd\x08\x77\x98\x7d\xa7\xd5\x9a\x8c\x45\x7a\x9d\x6c\x52\xf8\xff\xe3\x1f\xb0\x93\x29\x0c\x9b\x93\x5c\xe4\x9a\x56\xc5\xbb\xbf\x3b\xce\x93\x91\xfe\xe3\xf1\x1d\xe0\x4c\xaf\xa4\xa1\x3d\xe0\xc9\x24\x3b\xd3\x3c\x3d\xcb\xb3\x53\xd4\xb4\x2a\x21\xe9\x3e\x3f\x23\x56\xbc\xf8\x5d\x33\x25\xd4\x32\x18\xed\x4c\xc3\xce\x36\x7a\xa0\x78\xa1\x96\xcf\x4a\x1d\x7e\xba\x66\xdc\xd8\xf7\x21\xa1\xfc\xf7\xbd\x69\x32\xc5\xf3\x8f\x50\x34\xfb\xc4\x9d\x97\x8b\xed\x17\x7a\xb3\xde\xad\x76\xcb\x5d\xed\x57\x00\x00\x00\xff\xff\x86\x35\xf1\x9f\x2e\x08\x00\x00")

func internalTemplateClient_implTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateClient_implTmpl,
		"internal/template/client_impl.tmpl",
	)
}

func internalTemplateClient_implTmpl() (*asset, error) {
	bytes, err := internalTemplateClient_implTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/client_impl.tmpl", size: 2094, mode: os.FileMode(420), modTime: time.Unix(1539227189, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbe, 0x90, 0xab, 0xb3, 0x5c, 0x11, 0x3f, 0xdf, 0xbd, 0x70, 0x1d, 0x8b, 0x6a, 0x7c, 0x5f, 0x67, 0x92, 0x53, 0xc, 0x4d, 0x37, 0x26, 0x3, 0xbb, 0x96, 0xe6, 0xb8, 0xd4, 0x7, 0xb1, 0x84, 0x8f}}
	return a, nil
}

var _internalTemplateClient_streamTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\xbd\x6e\xdb\x30\x10\x9e\x9b\xa7\xb8\x04\x19\x24\x23\xa5\xf7\x16\x19\x8a\xa0\x2d\x3a\x14\x2d\xe2\xee\x81\x2a\x9d\x64\xc2\xd2\x51\x21\x29\xa7\x01\xa1\x77\x2f\x48\x51\xb4\x64\x4b\x8e\x63\xa4\x5b\xa7\x30\xa7\xbb\xef\xe7\xf2\x91\x31\x26\xc3\x9c\x13\xc2\x55\x5a\x72\x24\xfd\xa0\xb4\xc4\xa4\xba\x82\xf7\x6d\x7b\x61\xcc\x13\xd7\x6b\x60\x5f\x78\x89\xee\xd7\xeb\x42\xd4\x9b\x02\x3e\xdc\x02\xfb\x99\xa4\x9b\xa4\x40\xf6\x55\xf8\x53\xdb\x5e\x5c\x18\x23\x13\x2a\x10\xd8\x0a\xe5\x96\xa7\xa8\x1c\x0c\x80\x31\xd7\x1d\xbc\x1b\xbd\x73\x47\xdb\x6f\xbf\x2c\x17\xd0\x15\xa0\xa3\x06\x4e\x1a\x65\x9e\xd8\xe9\xc5\xb2\xef\xf2\xb8\xdf\x51\xaf\x45\xd6\xc3\xda\x0f\x3c\x07\x21\x7b\xcc\x95\x43\xe0\x54\x74\x0a\x50\x86\x82\xef\x5f\x2e\xc1\x18\xd6\x55\x7b\x19\xc0\x15\x24\x9e\xdc\x8e\x06\x7e\x68\x14\x66\xc0\x09\xf4\x1a\x77\x16\xec\x40\xdf\xc1\x1c\xa8\x7e\xae\x71\x0a\x36\xe0\x18\xd7\x06\x70\x27\x48\xe3\x1f\x1d\xc5\x90\x76\x27\xe6\x2b\x3b\x2f\x07\x46\x7a\xa7\x00\x2b\xa4\x2c\x5a\x18\x53\x88\x5f\x96\x90\xdd\xe3\x63\x83\x4a\x43\xf7\x47\x69\xdb\x1b\x60\x8c\x3d\x27\xb2\x4e\xbd\x92\x1f\xb5\xe6\x82\x62\x40\x29\x85\xf4\x14\x48\xd9\x78\x79\xfb\x8b\x1a\x10\xde\x63\xba\x8d\x66\x30\xc7\x42\x54\x2d\x48\xe1\x40\x89\xa3\x8c\x7b\xdb\xa5\x50\xe8\xd4\x9f\x23\x30\xa1\xec\x70\x2b\x11\x09\x7d\x20\x3d\x1e\x68\x77\x9c\x9f\x28\x7b\x1b\x0f\x63\x59\xed\x41\x6d\x77\x9e\x89\x74\x55\x97\x58\x21\xe9\xc4\x32\xff\x83\x5c\x4f\x45\xf0\x5b\x55\x97\x6d\x6b\x25\x34\xa9\x0e\x19\xf4\x8a\x16\x6e\x21\xb5\x14\x5a\xfc\x6e\xf2\x11\x8d\xf7\xe8\x7e\x6c\x13\x09\x0f\x13\xd9\xbe\x75\xcb\x9b\xa0\x8b\x23\xe2\x65\xdc\x0d\xe7\x0d\xa5\x10\xa5\x30\xd3\x39\x7f\x1d\x82\x58\x89\xba\x91\x04\x29\xeb\x54\xb3\x30\x31\xd4\x38\x79\x6f\xfc\x5e\x5e\x92\xe0\x52\x29\xf1\x11\x8e\xde\x2b\x51\x6b\x75\xf4\x72\xcd\x0a\xee\xf1\x3b\x0c\xc6\x58\x3c\x99\xa0\xf9\xeb\x78\xa2\x0f\x17\xf3\x63\x32\x4f\xca\x7a\xb0\x51\xa9\xc2\xd5\xec\x73\x1d\xbc\xdc\x63\x8a\x7c\x8b\x11\xe1\x53\x74\x04\x2c\xde\xf3\x0a\xc0\x73\x87\x75\x79\x0b\xc4\xcb\xc0\x11\x96\x45\xbc\x74\x64\xbe\xde\x86\x55\xaa\x1b\x10\x1b\x2b\xa1\x52\x05\x3b\x6a\x60\xc0\x74\x29\x36\x33\x14\x7b\x89\x4f\x94\xfe\x6c\x5d\x9f\xe0\xa7\x52\x45\x7c\xa0\xce\x01\x3b\x91\xc4\xcb\x3e\x8f\xef\x5e\xcc\x7c\x78\x0b\xcf\x0f\x95\xc3\x88\x4e\x0c\xd4\x2b\x9e\xcf\x13\xb3\x36\x7a\x5a\xdf\x2e\x73\x3e\x25\xc3\xc4\x8d\x8d\x7e\x3c\x27\x45\xff\x93\xfc\x9a\x24\xf7\xda\x67\xe2\x34\xf9\x0f\x6f\xea\xb8\xab\xfd\x0d\x00\x00\xff\xff\x9e\x66\x58\x33\x60\x0a\x00\x00")

func internalTemplateClient_streamTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateClient_streamTmpl,
		"internal/template/client_stream.tmpl",
	)
}

func internalTemplateClient_streamTmpl() (*asset, error) {
	bytes, err := internalTemplateClient_streamTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/client_stream.tmpl", size: 2656, mode: os.FileMode(420), modTime: time.Unix(1539227189, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2e, 0xfe, 0x2c, 0xa6, 0x89, 0xd1, 0x46, 0x7e, 0x2, 0x55, 0xb6, 0xeb, 0x20, 0xca, 0x5a, 0x15, 0x16, 0x21, 0xe2, 0x37, 0xc7, 0x4d, 0xe7, 0xd7, 0x3a, 0x96, 0xab, 0x85, 0x4b, 0xf9, 0x7b, 0x94}}
	return a, nil
}

var _internalTemplateFxTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x4d\x6f\x1b\x37\x10\xbd\xef\xaf\x98\x2e\x1a\x40\x0a\x14\xfa\x2e\x20\x97\x36\x51\x91\x43\x1d\x37\x01\x72\x31\x0c\x84\xe6\xce\x6e\x08\xcb\x24\x3d\xe4\x2a\x32\x88\xfd\xef\x05\x3f\xf6\x33\x9b\x5a\x8d\x4e\xd4\x68\xf8\xe6\x0d\xdf\xf0\x89\xde\x57\x58\x4b\x85\x50\xd6\xe7\x12\xde\x74\x5d\xe1\xfd\x77\xe9\xbe\x01\x3b\xc8\x23\xc6\xaf\xbf\x37\xda\x3c\x34\xb0\x7f\x0b\xec\x86\x8b\x07\xde\x20\xfb\x4b\xe7\x55\x4a\x20\xac\x8f\x28\x9c\xd4\xea\x0b\xa7\x98\xf8\x69\x88\xb0\x2f\x9c\xba\xae\x28\xbc\x27\xae\x1a\x04\xf6\x19\xe9\x24\x05\xda\x58\xac\x00\xf0\xfe\xea\x35\x1c\xce\x20\x8e\x12\x95\xb3\xf0\xfa\x2a\xc5\xaf\xae\xc0\x7b\x76\x38\xff\x19\xe3\x5d\x77\xc3\x89\x3f\x5a\x48\x74\x2d\xb8\x6f\x08\x26\x84\xd0\x21\xd9\x94\x4f\xf8\xd4\x4a\xc2\x0a\x9c\x06\x43\xfa\x24\x2b\x04\x1e\x50\x7a\x0c\x90\xca\x69\xe0\x2a\xa5\x1f\xce\xc0\x8d\x39\x4a\xc1\x23\xd1\x02\xc0\x3d\x1b\x5c\xad\x6a\x1d\xb5\xc2\x81\x2f\x00\x00\xea\x33\xfb\xa0\x8a\xb8\x4c\x69\x37\xa9\x16\xc1\x33\x27\x23\xd8\x3c\x58\x00\xac\xf6\xf3\x09\x6d\x7b\x74\x3d\x4f\x3b\x27\x9a\xf2\x33\xdb\x8b\x88\x66\xb8\x25\xd1\x8f\xad\x9b\x32\x5d\xd4\xe8\x89\x5d\xe3\xf7\x19\xd8\x2f\xb0\xda\x41\x6b\xa5\x6a\xa2\x2e\x8d\x3c\x61\x3e\x63\xc5\x1f\x11\x6a\x4d\x40\xba\x75\x52\x35\x2c\x86\xd3\x6f\x81\x5f\x3e\xa5\x4d\x8e\x84\x71\x48\xf3\xd6\x75\x6c\xc9\x6a\x53\xda\x34\x3b\x6f\x02\x6a\xb9\xdd\x0d\x9b\x18\x63\x79\xbd\x2d\x00\xea\x56\x89\x1f\x5a\xda\x44\x26\xd6\x91\x54\xcd\x0e\xb4\x71\x36\xec\x8a\x82\x19\xd2\x4e\xdf\xb7\x75\xee\xf3\xa3\x09\xfd\x6c\x43\x9f\x48\x35\x17\xe8\xbb\x7c\x9e\x84\xae\x25\x15\xf1\x37\x66\x6d\x4e\xb6\xb0\x59\x11\x65\x07\x48\xa4\x69\x9b\x51\x20\x8f\xfa\x0e\xf4\x43\xb8\x2c\x66\x31\x30\xf9\x6b\x24\xbc\xcd\x3b\x64\x0d\xbf\xe9\x87\x01\x60\xa0\xb2\x52\xcd\x77\xbb\x21\x0b\xa0\x7e\x74\xec\x7d\xa8\x5e\x6f\xca\x06\x15\x12\x77\x58\x81\xd0\x15\x82\xd0\xed\xb1\x02\xa5\x5d\x00\x23\x89\x27\xcc\xc4\xa2\x5e\xaf\x9e\xca\x1d\x4c\x29\x74\xc5\x8b\x95\x87\xba\x29\xbe\xcf\x22\x0c\x12\x0c\x7d\x1b\x67\x19\x63\xdb\x9e\x68\xb7\x03\x25\x8f\x45\x5f\x65\x66\x0b\x41\x72\xa4\x1f\x6d\xe1\x73\x8c\xff\xdc\x16\xf8\x7f\xdb\x42\xc8\xf3\x9e\xf5\x28\x21\x2e\xb0\x6a\x09\xed\xe5\x57\x6f\x41\xe1\x67\x1e\x91\xd2\xa6\xd5\x96\x96\xd0\xc7\x97\x96\xb0\x20\x99\x76\x8c\x4c\xff\x17\xc9\x97\xfc\xe1\x66\x84\x0d\x9f\xdb\xbb\xe4\x65\x43\x38\x46\xbf\x36\xa4\x5b\xb3\x2f\xe3\x6f\xf5\xb9\xfc\x1a\xb7\x8e\x66\xff\x37\x3a\x0e\xe3\xbf\x41\x66\x1e\xa3\x6b\x5b\x97\xf6\x33\x55\xe3\xb2\x13\x58\xed\x1f\x3e\x38\xc0\xb3\x41\xe1\xb2\x98\x7c\x26\xb5\xd3\x70\x8f\x60\x08\x6d\x98\x75\xa9\x62\x11\xa1\x95\xe3\x52\x21\xfd\x8a\x41\xf5\xd0\x9b\xcb\x2c\x69\x4c\xbf\xcc\x63\xe6\x73\xb6\x5d\xd5\xd5\xaf\xdc\xce\x79\xca\x78\x3b\x47\xa9\xf7\xf0\x47\x2b\x8f\x95\xf7\xa3\xce\xb6\xeb\x36\x26\x1f\xd6\x76\x74\x92\xb9\xc6\xfb\x75\x91\xfd\xc4\x78\xf2\xdf\xfc\x35\x7f\xc4\x3d\x78\x6f\x48\x2a\x57\x43\xf9\xea\xa9\x04\x76\xf8\xe7\xba\x9b\xb9\x54\x78\x6b\xbc\x43\x2b\x48\x1a\xa7\xc9\x86\x0d\xf3\x47\xc5\x34\x7d\x58\x76\x53\xc7\xf0\x1e\x55\x95\xde\x14\xd1\x3a\x46\xc2\xd9\x42\x92\x83\x9c\x38\xad\x80\xc3\x5b\xb8\xbd\xbb\xbd\xbb\x7f\x76\xe8\x87\x7b\x19\xa8\xc7\x91\xf3\x7e\xfa\x9e\x79\xaf\x84\xae\xa4\x6a\x12\xa5\xe1\x59\xf3\x0e\x0d\xaa\x0a\x95\x90\xf9\x69\x13\xb8\x2d\x80\x5e\x82\xea\x3b\x98\xb6\x33\x46\xff\x0d\x00\x00\xff\xff\xb4\x9c\x5f\xb4\xab\x09\x00\x00")

func internalTemplateFxTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateFxTmpl,
		"internal/template/fx.tmpl",
	)
}

func internalTemplateFxTmpl() (*asset, error) {
	bytes, err := internalTemplateFxTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/fx.tmpl", size: 2475, mode: os.FileMode(420), modTime: time.Unix(1539792370, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc4, 0x75, 0xf4, 0xca, 0x9f, 0x1f, 0xe, 0x7e, 0xd9, 0x6e, 0x88, 0x98, 0xd9, 0x4e, 0xe8, 0x31, 0x9f, 0x44, 0x62, 0xc7, 0x67, 0x68, 0xe8, 0xa, 0x3e, 0x92, 0xd0, 0xdd, 0xc2, 0xb4, 0xc1, 0xed}}
	return a, nil
}

var _internalTemplateServerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\x41\x8f\xd3\x3c\x10\xbd\xe7\x57\x8c\xaa\xfd\x3e\xda\xaa\xb8\xf7\x48\x1c\xa0\xd2\x02\x07\xaa\xb2\x0b\x07\x84\x38\x84\x64\x92\x5a\x9b\xda\x59\xdb\x69\xa9\x2c\xff\x77\xe4\xd8\x8d\x9d\xb4\xd5\x56\x02\xd1\x4b\x13\xcf\x9b\xe7\x37\x33\x6f\xa2\x75\x81\x25\x65\x08\x13\x89\x62\x8f\x62\x02\xaf\x8d\x49\xb4\x3e\x50\xb5\x05\x72\x4f\x6b\xec\x5e\xef\x2a\xde\x3c\x55\x90\xbe\x01\xb2\xc9\xf2\xa7\xac\x42\xf2\x9e\xfb\x27\x63\x92\x44\x6b\x91\xb1\x0a\x81\x3c\xa2\xd8\xd3\x1c\x65\x47\x03\xa0\xf5\x9d\xdc\xe7\x5d\xde\x3a\xdb\x75\x50\x7b\xb8\x9c\x83\x07\x02\x65\x0a\x45\x99\xe5\x08\xf3\xa5\x0b\x2f\x97\xa0\x75\x47\x84\xc2\x18\xa0\x12\xd4\x16\x3d\x93\x31\x20\x5d\xe2\x2b\x09\x4e\x71\x60\x20\x09\x80\x3a\x36\x38\x4c\xef\xf9\x75\x02\x60\x2f\xf7\x4a\x3f\xa1\xda\xf2\xe2\x24\xd4\x85\x68\x09\x5c\x00\x59\xd5\x14\x99\x7a\x54\x02\xb3\x1d\x65\x15\x78\xb6\x70\x10\x72\x6c\x96\x2f\x6d\xda\x1f\x79\x2a\xc6\xd5\x39\x57\x9c\x6a\x7f\x73\xad\x2b\xfe\xc5\xaa\x26\x0f\xf8\xdc\xa2\x54\xe0\x9a\x6d\xcc\x62\x40\x88\xac\x18\x25\xdb\x32\x3b\xde\x53\xb1\x21\x61\x16\x29\x18\x89\x37\x06\xa6\x83\x4b\x65\xc3\x99\xc4\x70\x2b\xa0\x10\x5c\x58\x06\xac\x25\x1a\xe3\xde\x9d\x82\xa8\x57\x36\xf8\x72\x23\x72\xce\x14\xfe\x52\x64\xe5\xfe\xe3\x92\x6e\xaa\x7c\x76\x9b\xd8\xe4\x52\x97\xe2\xb7\x60\xbc\x8d\xe0\x39\x16\xad\x40\xab\x4d\x2a\xd1\xe6\x8a\x72\x16\xbb\xef\x5d\x4b\xeb\x42\x6b\xd2\x23\xa5\x31\x01\xec\xec\xf8\xed\xed\xc3\x66\x05\x4d\x8f\x80\x92\x8b\x8b\x3e\xb5\xa6\x2c\x5b\x96\x5f\x64\x9d\xca\xd8\xab\x33\xf8\xfe\xe3\x98\x89\x26\x0f\x18\x6f\xda\xad\xdd\xa0\xff\x7b\xe8\xc7\x5d\x53\x1b\xa3\x9d\xff\x53\x90\xae\x5c\x81\xaa\x15\x0c\x3a\x86\x46\x70\xc5\x7f\xb6\x65\x74\xdb\x69\x2a\xd7\xe2\x9b\x4c\x64\x3b\xa9\xfb\xce\xfb\x05\x4d\x41\xeb\x46\x50\xa6\x4a\x98\xfc\xf7\x3c\x01\x72\xff\x79\x1d\x0f\xe8\x2b\xcb\xc4\x31\x3d\x29\xef\x79\xbb\xe3\xeb\xe4\x61\x0f\x5b\x0b\x3c\xed\x22\x19\x3b\x7c\xb0\x2a\x0e\x75\x26\xc8\x99\x6e\x31\x80\x7e\xc8\x58\x51\xdb\xde\x0c\x55\xad\xf1\xd0\x09\xf3\xe1\xe9\x20\x67\xdc\x9a\x18\x79\xae\x3f\xbe\x27\x85\x2d\xe9\xdd\xbf\x38\x03\x79\x77\x5b\x0b\xa7\xc0\xf0\x30\xbd\x6e\xfc\xd9\x38\x7b\x44\x37\x88\xbf\xf0\x81\x88\xc2\x6e\xf9\xcf\xa7\xe4\xce\x6f\x19\x93\xec\x90\xff\x70\x4e\x4e\xda\x4d\x83\x1a\x40\xff\x60\x52\x7f\xa3\xd7\xfe\x61\xd6\x7d\x72\x02\x2e\x3c\x86\xb3\xdf\x01\x00\x00\xff\xff\x98\x10\x28\x14\x7b\x07\x00\x00")

func internalTemplateServerTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateServerTmpl,
		"internal/template/server.tmpl",
	)
}

func internalTemplateServerTmpl() (*asset, error) {
	bytes, err := internalTemplateServerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/server.tmpl", size: 1915, mode: os.FileMode(420), modTime: time.Unix(1539227189, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa8, 0x5c, 0xf0, 0x21, 0x62, 0x13, 0xd, 0xf7, 0xeb, 0xec, 0x5, 0x76, 0x10, 0x21, 0xec, 0x65, 0xb, 0x8, 0x65, 0xc6, 0x7f, 0xe1, 0xa8, 0x9c, 0xaf, 0xd8, 0xa4, 0x88, 0xea, 0xba, 0xaa, 0xb3}}
	return a, nil
}

var _internalTemplateServer_implTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\xc1\x6e\xdb\x3a\x10\xbc\xfb\x2b\xe6\x05\xc1\x83\x64\x28\xcc\x3d\x80\x4f\x46\x5b\xf4\x90\xa2\x88\x7b\x0f\x58\x6a\x25\x13\x91\x28\x99\xa4\x1c\x07\x04\xff\xbd\x20\x25\x4b\x56\x52\xb7\x48\x9a\x56\x17\x2d\x97\xdc\xd9\x99\x59\x4a\xce\xe5\x54\x48\x45\xb8\x30\xa4\xf7\xa4\xef\x65\xdd\x56\x17\xb8\xf2\x7e\xe1\xdc\xa3\xb4\x5b\xb0\x8f\xb2\xa2\xb8\xbc\x2c\x9b\xf6\xa1\xc4\xcd\x0a\xec\x2b\x17\x0f\xbc\x24\xf6\xa9\x19\x22\xef\x17\x0b\xe7\x34\x57\x25\x81\x6d\x48\xef\xa5\x20\x13\x61\x00\xe7\x2e\x7b\xf0\xcf\x75\x5b\xc5\xf2\xcd\xb8\x0c\x75\xe1\xc4\xf5\x12\x7d\x12\x81\x00\xd5\xa4\x2c\xb7\xb2\x51\x58\x5e\xf7\x47\xec\x53\x4b\x73\x24\xef\x61\xac\xee\x84\x85\x5b\x00\x40\xbf\x03\xe7\x06\xf8\xd8\x7b\x80\x1f\x88\xdd\x92\xdd\x36\xb9\x89\x3b\x21\x2d\x0b\x70\x95\x83\xad\x2b\x49\xca\x6e\xac\x26\x5e\x4b\x55\x1e\x09\x4e\x89\xab\xa1\x04\x28\x3a\x25\x90\x6c\xb1\x7c\xc6\x25\x0d\x8d\xbf\xf0\x9a\xbc\x4f\x0c\x96\x4f\x5c\xb7\xa2\xd5\x8d\x6d\xbe\x77\xc5\x0c\x2e\x05\x69\xdd\xe8\x81\x74\x78\x34\xd9\x4e\x2b\x6c\x59\x8f\xc7\x26\xa0\xff\x83\x98\x58\x75\xea\x98\x33\x31\x75\x03\xe3\xd3\x01\xe4\x28\x88\x2a\x43\x90\xc5\x4b\x45\x7f\x59\x80\xc9\x42\x32\xcc\xf6\x4f\x54\x20\x50\x0f\x38\xff\xad\xa0\x64\x75\xd2\x61\x34\x89\xb4\x1e\x93\xfe\xb9\x83\x86\x6d\x48\xe5\x89\x26\x73\xd6\x97\x7f\x3c\x58\xb1\x1f\x8d\x31\xec\x8e\x04\xc9\x3d\x25\x8a\x1e\x13\xe7\xca\xe6\x5b\xb8\xd2\xec\x8e\x76\x1d\x19\x8b\xfe\xeb\xf2\x3e\x7d\x07\x37\x76\x19\xee\x43\xcf\x40\x80\x25\xcb\x5f\x34\x3b\xed\xa5\x69\x87\xd5\xd9\x5e\x73\xe1\x6b\x6e\xec\x87\x20\xf7\x77\x6a\xb2\x48\x22\x3d\x3f\xb4\x9f\x5c\x98\x28\xe0\x6d\x77\xff\x75\xf3\x14\xf6\x00\xd1\x28\x4b\x07\xcb\xd6\xfd\x3b\x43\x8d\xa8\x92\xdd\x92\x31\xbc\xa4\x14\xc9\x6c\x9d\xf5\x63\x4e\x67\x73\x3e\x1a\x5e\xbf\x97\xdb\x4a\x56\xd9\xdb\x2d\xaf\x5f\xe7\xb7\xb0\x87\x30\xa6\xdd\x0b\x47\x55\x3e\xfe\xc2\x8f\xf1\x14\x4d\xe1\x94\xfb\x11\x00\x00\xff\xff\xe2\xd4\x05\xc7\x4f\x06\x00\x00")

func internalTemplateServer_implTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateServer_implTmpl,
		"internal/template/server_impl.tmpl",
	)
}

func internalTemplateServer_implTmpl() (*asset, error) {
	bytes, err := internalTemplateServer_implTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/server_impl.tmpl", size: 1615, mode: os.FileMode(420), modTime: time.Unix(1539227189, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8d, 0xb9, 0xde, 0x81, 0x67, 0x75, 0x8f, 0x37, 0xf2, 0x1d, 0xdd, 0x3a, 0x82, 0x6, 0xa0, 0xa2, 0x50, 0xcd, 0xe7, 0xde, 0xaf, 0xf8, 0x3f, 0x44, 0xe6, 0xc9, 0x12, 0x9c, 0x83, 0x76, 0x93, 0xcd}}
	return a, nil
}

var _internalTemplateServer_streamTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xc1\x6e\xdb\x30\x0c\xbd\xf7\x2b\xd8\xa2\x07\x3b\xc8\x94\xfb\x80\x9c\x82\x6d\xd8\x61\xd8\x90\xec\x5e\x78\x0e\xed\x08\xb1\x25\x57\x92\xdd\x15\x82\xfe\x7d\x10\x2d\x2b\x49\x63\xa5\xbd\xec\x14\x85\x26\xdf\x7b\x24\x9f\x64\xed\x1e\x2b\x2e\x10\x1e\x34\xaa\x01\xd5\x93\x36\x0a\x8b\xf6\x01\x3e\x39\x77\x67\xed\x0b\x37\x07\x60\x5f\x79\x83\xf4\xf7\xb1\x96\xdd\xb1\x86\xcf\x6b\x60\xbf\x8a\xf2\x58\xd4\xc8\xbe\xc9\x70\x72\xee\xee\xce\x5a\x55\x88\x1a\x81\xed\x50\x0d\xbc\x44\x4d\x30\x00\xd6\x3e\x8e\xf0\x54\xba\xa3\xa3\xcf\xf7\x5f\x56\x0b\x18\x03\x30\x52\x03\x17\x06\x55\x55\xf8\xea\xc5\x6a\xca\x0a\xb8\x3f\xd0\x1c\xe4\x7e\x82\xf5\x1f\x78\x05\x52\x01\xdb\x34\x1c\x85\xd9\x11\x02\x17\xf5\x44\x12\x03\x21\x7f\xb5\x02\x6b\xd9\x18\x9d\x64\x00\xd7\x50\x04\x72\x5f\x1a\xf9\xa1\xd7\xb8\x07\x2e\xc0\x1c\xf0\xd4\x82\x2f\x98\x32\x18\x81\x9a\xd7\x0e\xe7\x60\x23\x8e\xa5\x34\x80\x8d\x14\x06\xff\x9a\x2c\x87\x72\x3c\xb1\x10\x39\xf5\x72\xd5\xc8\xd4\x29\xc0\x16\xcb\x21\x63\x8c\xbd\x16\xaa\x2b\x03\xd9\xcf\xce\x70\x29\x72\xc8\x16\xd6\xd6\xf2\xb7\x17\xc2\xb6\xf8\xdc\xa3\x36\x30\x2e\xcb\xb9\x25\xa0\x52\x52\xe5\x81\x04\xc5\xfe\x72\x7c\x6f\x47\x75\x46\xb9\x43\xb1\xbf\x84\xd6\x9d\x14\x1a\xcf\xb0\x13\x82\x88\x72\x86\xd1\x5d\xc5\x4e\xe7\x84\x1d\xda\xae\xc1\x16\x85\x29\x3c\xf2\x7f\xf0\xc4\xdc\xfa\xbe\xb7\x5d\xe3\x9c\x97\xd0\x97\x26\xee\x2f\x28\x5a\x50\xc3\x9d\x92\x46\xfe\xe9\xab\x0b\xd4\xd0\x23\xfd\x0c\x85\x82\xa7\x19\x5f\xac\x69\x5b\x33\x74\x79\x26\x78\x93\x8f\xc5\x55\x2f\x4a\xc8\x34\x24\x32\xd3\x56\x8a\x62\x15\x9a\x5e\x09\xd0\x6c\x54\xcd\x62\xc5\xb9\xc6\x59\xcf\x85\xb9\xbc\x27\x81\xfc\x28\x3b\xa3\x53\x1e\xf8\x88\x29\xa3\xdc\x56\xd7\x14\xf3\x2f\x44\xd4\xbc\xc5\x12\xf9\x80\x99\xc0\x97\x2c\x8d\x95\x2f\xc1\xcb\x60\x8c\xe5\x01\x8c\x57\x04\x75\xbf\x06\xc1\x9b\x48\x11\x67\x22\x78\x43\x5c\x21\xee\xe2\xc4\x9e\x97\x20\x8f\x5e\x41\xab\x6b\x76\x4b\xfe\x19\xd1\xbd\x3c\x26\x18\x2e\x6d\xb2\x29\xb4\xf9\xe2\x7b\x7e\xbf\x9b\x56\xd7\xf9\x95\x36\xc2\x25\x89\x82\x37\xb3\x77\x29\x7d\xa5\x3f\xb8\x51\xba\xee\x0a\x29\xe1\xc6\x95\xbf\xb5\x73\xda\x6a\xd2\x83\x13\xc1\x9b\x85\x25\x1a\x99\x7d\x23\xe6\x8e\xa7\xd8\xbf\x00\x00\x00\xff\xff\x41\x80\x9a\x3a\xcf\x06\x00\x00")

func internalTemplateServer_streamTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateServer_streamTmpl,
		"internal/template/server_stream.tmpl",
	)
}

func internalTemplateServer_streamTmpl() (*asset, error) {
	bytes, err := internalTemplateServer_streamTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/server_stream.tmpl", size: 1743, mode: os.FileMode(420), modTime: time.Unix(1539227189, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0xf8, 0x11, 0xaa, 0x40, 0x44, 0x82, 0x90, 0xdf, 0x8b, 0x86, 0xcd, 0x10, 0xd, 0xd5, 0xaf, 0xd3, 0x83, 0x6, 0xc6, 0x20, 0xc0, 0xda, 0x92, 0xe0, 0x32, 0xf7, 0x5c, 0x83, 0x1c, 0xf, 0x4c}}
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
	"internal/template/.nocover": internalTemplateNocover,

	"internal/template/base.tmpl": internalTemplateBaseTmpl,

	"internal/template/client.tmpl": internalTemplateClientTmpl,

	"internal/template/client_impl.tmpl": internalTemplateClient_implTmpl,

	"internal/template/client_stream.tmpl": internalTemplateClient_streamTmpl,

	"internal/template/fx.tmpl": internalTemplateFxTmpl,

	"internal/template/server.tmpl": internalTemplateServerTmpl,

	"internal/template/server_impl.tmpl": internalTemplateServer_implTmpl,

	"internal/template/server_stream.tmpl": internalTemplateServer_streamTmpl,
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
	"internal": &bintree{nil, map[string]*bintree{
		"template": &bintree{nil, map[string]*bintree{
			".nocover":           &bintree{internalTemplateNocover, map[string]*bintree{}},
			"base.tmpl":          &bintree{internalTemplateBaseTmpl, map[string]*bintree{}},
			"client.tmpl":        &bintree{internalTemplateClientTmpl, map[string]*bintree{}},
			"client_impl.tmpl":   &bintree{internalTemplateClient_implTmpl, map[string]*bintree{}},
			"client_stream.tmpl": &bintree{internalTemplateClient_streamTmpl, map[string]*bintree{}},
			"fx.tmpl":            &bintree{internalTemplateFxTmpl, map[string]*bintree{}},
			"server.tmpl":        &bintree{internalTemplateServerTmpl, map[string]*bintree{}},
			"server_impl.tmpl":   &bintree{internalTemplateServer_implTmpl, map[string]*bintree{}},
			"server_stream.tmpl": &bintree{internalTemplateServer_streamTmpl, map[string]*bintree{}},
		}},
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