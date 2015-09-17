// Code generated by go-bindata.
// sources:
// data/aws-vpc-public-private/build/template.json.tpl
// data/aws-vpc-public-private/deploy/main.tf.tpl
// data/aws-vpc-public-private/deploy/variables.tf
// data/common/dev/Vagrantfile.tpl
// data/common/dev-dep/Vagrantfile.fragment.tpl
// data/common/dev-dep/Vagrantfile.tpl
// data/common/dev-dep/build.sh.tpl
// data/common/dev-dep/upstart.conf.tpl
// DO NOT EDIT!

package goapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _dataAwsVpcPublicPrivateBuildTemplateJsonTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x51\x6e\xf3\x20\x10\x84\xdf\x7d\x0a\x84\x94\xb7\x24\xca\xff\x57\x6a\xaa\x5e\xa5\xaa\xc8\x42\x56\x29\x8a\xc1\x16\x0b\xa9\x52\xc4\xdd\x8b\x8d\xdc\x42\xd5\x26\x7e\xb1\xc5\x37\xde\x99\x1d\x62\xc7\xf2\xc3\x8d\xb6\x62\x04\x75\x46\x27\x2e\xe8\x48\x0f\x96\x3f\x33\xbe\xdb\x3e\x6d\x77\x7c\xdd\x15\xcd\x05\x9c\x06\xd9\x23\x65\x54\x7e\x9b\x8f\xe1\x9d\x04\x28\x85\x44\xe2\x8c\xd7\xcc\x6c\xe8\xfb\x75\xcb\x09\x95\x43\xff\x37\x77\x78\x2a\x96\x13\x9b\x51\x5a\x5c\x65\xd0\xfd\x31\x47\xca\xf0\xa5\x72\xb5\x60\x70\x8a\x38\x78\x3f\xf0\x6a\x9a\xbf\x8e\xf3\x39\x18\xf8\x18\xec\x06\x25\xd5\xb4\xc9\xc9\xe3\x8a\xe5\x65\x25\x78\x6d\xd8\x2a\xc5\xc8\x02\xa1\x63\x87\x76\xa1\x03\x4b\x29\x0b\xd1\x1e\x2b\x6d\x3d\xb3\xd9\xed\xd6\xcc\x6f\xe1\xdd\x99\x5f\x7d\xdc\x9a\x57\x44\xf7\xf3\x0d\xc1\x29\x14\x60\x74\xe9\x45\x6f\xf6\x8f\xf2\x3f\xec\xff\x61\xad\xd2\x96\x3c\xd8\xac\x5b\x0a\x54\x0f\xdb\x1e\xdc\xa9\x11\x11\xbd\x89\xc9\x7f\x29\x3f\xc8\x60\x7d\x68\x0a\x36\x5a\x2c\x34\xc6\xe9\x2b\x25\xf6\x73\x87\xfc\xc6\xec\x66\xc6\xdf\x92\x97\xdb\x7f\xed\x52\xf7\x19\x00\x00\xff\xff\x20\x45\x14\x15\x9b\x02\x00\x00"

func dataAwsVpcPublicPrivateBuildTemplateJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateBuildTemplateJsonTpl,
		"data/aws-vpc-public-private/build/template.json.tpl",
	)
}

func dataAwsVpcPublicPrivateBuildTemplateJsonTpl() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateBuildTemplateJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/build/template.json.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployMainTfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8f\x41\x4e\xc3\x30\x10\x45\xf7\x3e\xc5\x97\xcb\x12\x71\x83\xee\x90\xd8\xc1\x11\xa2\x69\x3c\x54\x16\x89\x1d\xd9\x93\x22\x2b\xf2\xdd\x99\x86\xa4\x26\x62\x96\x6f\x9e\xfd\xff\x9c\xf0\xc6\x81\x13\x09\x3b\x5c\x0a\x3e\x44\xe2\x33\x5c\x44\x88\x02\x76\x5e\x30\x52\x98\x69\x18\x8a\x31\x53\x8a\x37\xef\x38\xc1\xd2\x77\xb6\x58\x0c\x74\xa8\xef\x39\xe7\xee\x8b\x0b\xce\xb0\x4f\xcb\x8d\xd2\x8b\xae\xbb\xc6\xab\x5d\xc5\xcc\x7d\x62\xf9\x2f\x36\xbe\x89\x89\xaf\x3e\x86\xa3\xf4\xcb\x54\xa8\xc6\x9c\xf0\xca\xd3\x10\x0b\x48\xff\x14\xc4\x4f\xf8\x90\x85\x82\xe6\x99\xc4\x39\xce\xa9\xe7\xb5\x62\xb7\x73\x0b\xbb\x2c\x08\x34\x32\x6a\x7d\x14\x1f\xfd\x9f\x8c\xd1\x6f\xe9\xfb\x9b\x4e\xca\xc4\x4d\x38\xe0\xfd\xa2\xf9\x12\xb4\xb9\x77\x4d\x7b\x20\x55\x56\x47\xe8\x9a\xb7\xc0\xfb\xbc\xdf\x3b\x9c\x0f\x75\xd6\x5d\xd5\xc3\x7e\x02\x00\x00\xff\xff\xa1\x87\x8b\xff\x8a\x01\x00\x00"

func dataAwsVpcPublicPrivateDeployMainTfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployMainTfTpl,
		"data/aws-vpc-public-private/deploy/main.tf.tpl",
	)
}

func dataAwsVpcPublicPrivateDeployMainTfTpl() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployMainTfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/main.tf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployVariablesTf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x91\xb1\x4e\xc4\x30\x0c\x86\xf7\x3c\x85\x95\x9b\xb9\x81\x9d\xe1\x24\x96\x0e\x2c\xdc\xc0\x58\xf9\x72\x2e\x58\xe4\x92\x28\x76\xa9\x2a\xc4\xbb\x93\x46\x95\x90\x2a\x05\x96\x62\x75\xaa\xbf\x7c\xb6\x7e\x1f\xee\x76\x28\x73\x80\x93\x73\x24\x02\x5d\x18\xa2\xd9\xc7\x69\x3e\x30\x33\x5e\x3c\x81\xc5\x49\x7a\xac\x03\xfa\x77\x9a\x2d\x7c\x1a\x28\x75\x25\x71\x99\x93\x72\x0c\xf0\x00\x76\xdd\xa0\x00\x30\xc4\x0c\xa7\x97\xb3\x35\x5f\x5b\x8b\x90\xcb\xa4\xbf\x58\xce\x15\xf8\xc3\x92\xe9\xb5\xe0\x0d\xc3\x73\x6d\xc2\xf4\x46\x99\x60\x2a\x1f\x7b\x0f\x31\x51\x46\xa5\x63\x95\xed\x95\xf9\x23\x25\x1f\xe7\xff\xca\xfc\xc6\xad\xa0\x9f\x3a\xd0\x58\xfe\x2e\xd3\x37\xe9\x70\x10\xc5\xe0\xa8\xd7\x39\x51\xe3\x7d\xb7\x32\x50\x99\x95\x18\x70\xf4\xba\x74\xf5\xfe\x28\x37\xf4\x7e\x23\x96\xf1\x12\xca\xe1\xf8\xda\xba\x5b\xed\xff\xec\x05\x1c\x34\x2e\x8e\xef\x00\x00\x00\xff\xff\xc5\x62\x66\xc0\xe1\x02\x00\x00"

func dataAwsVpcPublicPrivateDeployVariablesTfBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployVariablesTf,
		"data/aws-vpc-public-private/deploy/variables.tf",
	)
}

func dataAwsVpcPublicPrivateDeployVariablesTf() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployVariablesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/variables.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x56\x6d\x6f\xd4\xb8\x13\x7f\x9f\x4f\x31\xa4\x2d\xb4\x52\xe3\x14\x84\x78\x51\x68\x05\x2a\x7f\xa0\xd2\xff\x54\xee\xda\xe3\x0d\x42\x8b\x37\x71\x12\xab\x8e\xed\xb3\x9d\x7d\xa0\xdd\xef\x7e\x33\x76\xb6\xdb\x05\xca\xdd\x49\xdd\xd6\x3b\x9e\xf9\xcd\xd3\x6f\xc6\xdd\x81\xf7\x42\x0b\xc7\x83\xa8\x61\xba\x84\x8b\x10\xcc\x21\xd4\x06\xb4\x09\x20\x6a\x19\x1e\x65\x3b\xd9\x0e\x5c\x75\xd2\x03\xfe\x84\x4e\xc0\x27\xde\x3a\xae\x43\x23\x95\x80\xf6\x7b\x5b\x68\x8c\x8b\x5a\xb5\x98\x09\x65\x6c\x2f\x74\x00\xd3\x20\x44\x20\x08\x6e\xad\x92\x15\x0f\xd2\xe8\xd2\x0b\x37\x93\x95\x60\x70\x1e\xc0\x77\x66\x50\x75\x74\x3a\x15\xd0\x71\x5d\x17\xe4\x5c\xd4\x0c\xae\x0c\xf4\xa6\x96\xcd\x92\x60\x11\xe7\x9e\xfb\x43\x18\xbc\x88\xde\xde\x58\x4b\x02\x96\x65\xe3\x35\xab\x8c\x6e\x64\x3b\x38\xb1\x9f\x3f\xcb\x0f\x28\xa3\xdb\x24\xba\xcd\x00\xd2\x89\xcd\x7a\x36\x35\x0b\x38\x81\xbc\xe3\xbe\x93\x95\x71\xb6\xb4\x4e\x54\xd2\x8b\x17\xcf\xf3\x0c\x15\x77\xe0\x52\x84\xc1\x02\x07\xbf\xd4\x15\xa6\xd9\x18\x55\x0b\x07\x8d\x33\x3d\x98\xc1\xc1\xdc\xb8\x6b\xa9\x5b\xa8\x25\xda\x05\xe3\x30\x4a\x03\xe5\x2c\x05\xb1\xe5\x29\x01\x4c\x46\x80\xfc\xe6\x06\x2c\x0f\x1d\x5b\x03\xac\x56\xf9\x61\x94\xfa\x8e\xbb\x3b\xbd\x09\xe9\xc4\x3b\xc4\x02\x30\x73\x2c\xf7\x31\xe4\x23\x3e\x5a\xb4\xce\x0c\xf6\x9e\x84\x82\xbe\xd9\x03\xd9\x80\xec\xad\x71\x21\x01\x3c\xc2\x14\x73\xd8\x5b\xc5\x8c\xde\x4a\xcf\xa7\x4a\x8c\x5d\x6a\xf8\xa0\xc2\x76\x76\xbf\x0a\x9b\x51\x94\xe5\xc6\x7f\x9d\xc0\xea\x63\x08\x6e\x10\xc9\xb9\xd0\xd8\x2e\xf2\x16\xdd\xfd\x4f\x47\x6f\x97\x97\x1f\x80\xb7\xc4\x06\x64\xc8\x9c\xbb\x9a\x92\xf6\x06\x19\x14\x02\x1d\xad\x93\x33\x24\x12\x46\x64\xd1\x5e\xe8\x4a\x0a\x1f\xab\xeb\x37\xe1\x78\xdf\xb1\xd1\x7a\x92\xb0\x4e\x92\xdb\xe8\xe8\x9d\x19\x74\x1d\xa9\x05\xeb\xe6\xa7\x6f\xfb\x18\x0d\xd7\xcb\x83\x14\x1d\x11\x14\xbb\x05\x52\xe3\x71\x6d\x31\x41\x89\x67\xc8\xd9\x54\x24\xba\xc7\x92\x95\x06\x19\x5d\x6e\xb4\x0a\x6c\x0f\x9a\x2b\x63\x2c\x3b\x43\x69\xc0\x8a\x60\x6f\xfe\xa1\xcd\x04\x16\xbb\x8b\x87\x2d\x55\xeb\xcc\x4c\x7a\x8a\x30\xf7\x9d\x50\x0a\x55\xa4\x56\x52\x0b\xec\x67\x55\xc3\xce\x0d\x1a\xac\xe0\xf1\x63\x98\x22\x3b\xc7\xaf\x65\xcf\xa5\x66\xbe\xcb\xef\x4a\x4d\xf9\xac\x6b\xfd\x7f\xc3\x6b\xe0\x4a\x45\x6a\x36\x8e\xb7\x34\x7e\x1e\x3a\xe1\x44\xcc\x1b\xab\xb0\x55\x60\xb6\x29\xc9\x5a\x9b\xea\x82\x75\x98\x6c\xac\x63\x45\x28\xf3\x51\x72\xeb\x04\x7a\x59\xad\x7e\x1a\xc1\xb9\xf6\x81\x02\x78\x6f\x60\x3a\x48\x1c\x69\xa1\x67\xd2\x19\x4d\x86\xff\x36\xf9\x5d\x5f\x39\x69\xc3\xa4\x35\x8a\xeb\x36\xe1\xfe\xc6\xaf\x05\xc8\x40\x8c\x09\x1d\x0f\xf0\x75\xa4\x20\x20\x27\xbe\x42\x6b\x90\x2c\x69\x04\x55\x9c\x40\xe2\x36\x8e\x33\x09\xfe\x43\xd9\xe3\x8c\xed\xfd\xfe\x59\x54\x9d\x89\x2d\x78\x70\x1c\xe1\xf4\x14\xca\xce\xf4\x62\x3d\x0a\x25\xa3\x26\xb9\xea\x4b\x0a\xf7\x6e\x5f\x9a\xc8\x61\xe0\x8e\x48\x84\xd1\xf7\x02\xeb\xd2\x7a\x70\xb2\xed\x02\xee\xbb\x39\xaa\x7f\xce\x67\x3d\x72\x5a\x4c\x9a\x81\xc2\xa2\x09\x1b\x05\x91\xff\x21\x72\x2f\xff\xc2\x04\xaf\xba\xb8\xc8\x34\xef\xc5\x6d\x0c\xf6\xbb\xac\x30\xc4\x7d\xba\x4c\xfb\xce\x26\x1d\x00\xcb\x44\x1c\xc2\xc9\xac\x77\x83\x9e\x48\x3b\x41\x0a\x5f\xe3\x5a\x3b\x81\x86\x2b\x2f\xa2\x1a\x36\x32\x4b\xbf\xe9\x93\x6d\x37\x01\x35\x5f\xbd\xba\x3c\xfb\xe3\xfc\xe3\x55\xe6\x45\x80\x02\x67\xce\x88\xfd\x03\xb8\x81\xdd\xd7\xf0\xec\xf4\xf1\x53\xb8\x05\x65\xda\x16\x39\x5f\xe0\xc2\xa7\x97\x00\x2b\x84\x44\x2a\xf5\xa0\xd4\x4b\x58\x65\x46\x45\xf5\x54\xdb\xcf\xa4\xf1\x05\x6d\x73\xba\xc2\xad\x7e\xde\xc0\x9c\x16\xff\x4c\x20\x75\x0e\xa9\x7f\x7a\x7c\x85\x3a\xdc\x0e\x19\x4e\x70\x65\xfa\x1e\xdf\x05\x28\x66\xd8\x6e\x38\xbd\xc3\x8e\xde\x5f\x46\x8b\xb4\x23\x15\xe4\xc8\x3e\xae\x88\xa6\x4b\x6c\x6d\x24\xa4\xa8\x1f\xa5\xf7\x69\x6e\xf4\x93\xb0\x96\xa2\x2f\x96\xa7\xec\x17\x48\xae\xa3\xac\x91\x98\x17\x02\xbc\xc5\x4d\xab\x70\x98\x68\x33\x21\x18\x4d\x31\xce\x44\x6b\x26\x33\xe1\x22\x75\x56\x2b\xc6\xd0\xd4\x08\x98\xb7\x54\x8f\xbf\xa0\xb8\xf8\x8e\x11\xad\x61\x81\x3b\xd6\x7e\x83\x2e\x04\xeb\x8f\xcb\xd2\xe3\x03\x81\x7b\x8b\xb5\xc6\xb4\x4a\x70\x2b\x3d\x3e\x53\x7d\x99\x4a\x8c\x7f\x7e\xea\x06\x99\x39\x2c\x0a\xde\xd7\x2f\x9e\x8f\x78\x29\xc4\x3f\x35\x7e\x73\x29\xc0\x75\x2c\x7e\xc0\x9a\xa1\x18\x8a\x33\x28\x07\xef\x4a\x65\x2a\xae\xa0\x58\x7c\x6b\x1e\x0a\x2e\x81\xe1\x78\x45\xa4\x8b\x8f\x6f\xae\x3e\x6c\xa1\xf5\xd7\xb4\xbf\x0a\x0b\xa5\xb1\x64\x46\x13\x90\x35\x3e\x2c\xad\x38\xd9\xdd\x6f\x24\x76\xe4\xde\x0d\x14\x3d\x4a\x84\xc5\xd3\x11\x9e\xf9\xe2\xee\x4c\x06\x80\xdd\xc3\x3d\x8f\x8f\x37\xe4\x7b\xef\xf2\x83\xec\x47\xf3\x84\x0c\xbb\x37\xe9\xb0\x1a\x0d\x8e\x90\x5e\x0b\xee\x70\x70\x0a\xc4\xd2\xf0\xf4\xe8\x08\x90\x48\x73\x0d\x63\x42\xc7\xeb\x27\x37\xa6\x73\x39\x3e\x2a\x48\xf1\xbb\x84\x22\xf1\x9e\x88\x05\xbd\x8a\x51\x7a\x72\xcf\x71\x39\x95\xfa\x78\x53\x31\x94\x46\xc9\x2e\xe9\x3d\x79\x70\xd8\xb7\x31\x53\xf1\xee\xa3\xfe\xc2\x32\x86\x39\x2e\x4b\x8a\xf4\xd3\xd9\xa5\x8f\x7b\xb8\x8d\x4f\xe2\x56\x0b\xb8\x0d\x05\x91\x6c\xb0\x35\xbd\x90\xc5\xf2\x87\x9b\x35\x9d\x8b\x25\xb4\xc8\xe3\xe9\x37\x07\xbd\x70\xd5\xe0\x24\x57\xc9\xd5\xd9\xf8\x22\x8e\x84\xc6\x41\xa0\x7f\x9d\xe8\x51\x26\x5b\x5a\xe8\xa6\x81\x0f\x57\x57\x1f\xa3\x67\x02\x49\xab\x05\x8a\xa2\x55\x66\x8a\x24\x1a\x9c\x62\x39\x5e\xbc\xc6\x4f\x37\x4c\x89\xb9\xc7\x39\x1b\xad\x2f\xb0\xa3\x6b\x96\x6f\xee\xcb\x3c\x1b\x77\xc6\xdf\x01\x00\x00\xff\xff\xd1\x74\x56\xa0\x66\x0a\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepVagrantfileFragmentTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\xbb\x4e\x2b\x31\x10\x86\x7b\x3f\xc5\x9c\xcd\x29\xc9\xba\x8f\x42\x95\x82\x12\x24\x68\xa8\x90\x63\x4f\xd8\x91\x7c\x93\x3d\x5e\xb4\x8a\xf2\xee\xd8\x9b\x48\x59\xc4\x45\x50\x5a\x1e\x7f\xdf\xfc\xbf\xbc\x82\x3b\xf4\x98\x14\xa3\x81\xfd\x04\x06\x23\x7a\x83\x5e\x4f\x70\x08\xa9\x1e\x47\xb4\x21\x3a\xf4\xbc\x81\xe3\x11\xbc\x72\x08\xa7\x93\x10\xff\xaf\x87\x97\x8c\x5c\x22\xdc\xc2\x76\xbb\xbb\x7f\x78\x16\x2b\xd8\x85\x38\x41\x28\x09\xf6\xe4\x55\x9a\x44\x2e\x26\x80\x1b\x41\xb2\x8b\xb2\x1a\xd6\xd7\xc7\x20\x4b\x4e\xd2\x06\xad\xac\xac\xd3\x72\xe9\xb8\x80\x78\x40\x28\x31\xb3\x4a\x0c\x07\xb2\xf8\x13\xae\xbf\x0c\xf6\x3a\xf8\x03\x48\x64\x2d\xc9\x13\x2f\xb0\xf3\x4d\x63\x3f\xce\x40\xe2\x7f\x67\xde\x99\xbf\xd0\x2f\xb3\x34\x6d\x06\xf2\x1c\x40\x01\xa3\x8b\x60\x28\xa1\xe6\x90\xa6\x1e\x9e\xea\x7e\x59\x27\x8a\x0c\x6f\x64\x2d\xb8\x30\x62\x5b\xda\xf5\xa2\xb9\xe8\xb5\x1f\x5d\x1f\x53\x18\x29\x53\xf0\xd0\x35\x58\x77\x23\x00\x72\xad\x48\xe3\x06\xba\x6a\x8d\x8a\x87\x5e\x2b\x3d\x34\x77\x0d\x35\xae\x5b\xb0\x50\x38\x16\x9e\x87\x0d\x66\xae\x6d\x72\x45\xd4\x17\x5f\x44\xef\xc4\x9f\x75\xc1\xc5\x7a\x69\x16\x46\xb9\x2f\x64\x8d\x5c\xb6\xf8\x3b\xfb\x87\xe2\xbf\x5b\x25\x0f\x68\xed\xcc\x23\x6f\xc9\xd7\x5d\x3e\x7d\x23\xf1\x1e\x00\x00\xff\xff\x97\x51\x37\x37\x90\x02\x00\x00"

func dataCommonDevDepVagrantfileFragmentTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepVagrantfileFragmentTpl,
		"data/common/dev-dep/Vagrantfile.fragment.tpl",
	)
}

func dataCommonDevDepVagrantfileFragmentTpl() (*asset, error) {
	bytes, err := dataCommonDevDepVagrantfileFragmentTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/Vagrantfile.fragment.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\xdd\x4e\xf3\x46\x10\xbd\xf7\x53\x4c\x9d\x4f\xa5\x95\xf0\x5a\x20\xc4\x45\x04\x91\x10\x48\x29\x17\x15\xa8\xa4\xf4\x12\x36\xf6\xc4\x5e\x61\xef\x6c\x77\xd7\xf9\x01\xf2\xee\x9d\x5d\x3b\x40\x90\xe8\x87\x84\x88\x3d\x9e\x39\xe7\xcc\xcc\xd9\x1d\xc1\x14\x35\x5a\xe9\xb1\x84\xf9\x06\x6e\xbc\xa7\x43\x28\x09\x34\x79\xc0\x52\xf9\x5f\x92\x51\x32\x82\x59\xad\x1c\xf0\x9f\xaf\x11\xee\x65\x65\xa5\xf6\x0b\xd5\x20\x54\x9f\x6b\x61\x41\x16\xe6\x9d\x6a\x4a\xa5\xab\x90\xce\xc5\x73\xa5\xa5\xdd\xf0\x8b\xf4\x01\xa3\x73\x9c\x2e\x1d\x48\x28\xd1\xa0\x2e\x51\x17\x9b\x58\x56\xe2\x12\x1b\x32\x2d\x6a\x2f\x22\xeb\x55\x2f\xa3\x96\xba\xcc\x82\x16\x86\xe0\xfa\x40\x2c\x60\x46\xd0\x52\xa9\x16\x9b\x18\x3c\x0c\xa8\x51\xdd\x85\x31\x31\x21\x49\x06\x9d\xa2\x20\xbd\x50\x55\x67\xf1\xb7\xf4\x38\xfd\x3d\xf4\xf6\xda\x87\x5e\x13\x80\xfe\x49\x2c\x5b\x31\xa7\x35\x9c\x43\x5a\x4b\x57\xab\x82\xac\xc9\x8d\xc5\x42\x39\x3c\x3d\x49\x13\x4e\x1c\xc1\x1d\xfa\xce\xb0\x6a\xb7\xd1\x05\x77\xb0\xa0\xa6\x44\x0b\x0b\x4b\x2d\x50\x67\x61\x45\xf6\x29\xf4\x5c\x2a\xae\xf3\x14\x1a\x26\xc8\x97\xbd\x88\x3d\xa6\x1e\xe0\x61\x00\x48\x5f\x5e\xc0\x48\x5f\x8b\x1d\xc0\x76\x9b\x1e\x42\xba\xab\xfc\x39\xf9\xaa\x46\x8b\x51\x42\x41\xad\xe1\xde\x4b\x28\xa5\x97\x71\x5d\x14\x8b\x73\xe2\xcd\x08\xf8\x07\x43\xf3\x71\x86\x2c\x4d\x16\x05\xba\x7e\xa3\x71\x5f\xe0\x0a\xab\x0c\x4f\xfe\x1b\x52\xdf\x88\xb6\xdb\x9c\xb7\x96\xf1\x22\xf3\x08\x12\x95\x07\xb6\xef\xca\x0e\xf4\x85\x2c\xf8\x3f\xcf\x8d\x25\x7f\x8b\x3d\xe6\x0f\x63\x0a\x64\x59\x8c\x0c\x94\xd7\xda\x79\xd9\x34\x30\xa5\xa1\x2f\xd4\x4b\x65\x49\x07\x5b\xed\xa1\x1b\x4b\x4b\xe5\x14\x69\x48\x5d\x8d\x4d\xc3\x70\x4a\x37\x4a\xe3\x18\x7e\xf4\xb3\x78\xa8\xa8\x91\xba\x4a\xd8\xa4\x49\xb2\x1f\x63\xab\x9c\x9d\xdd\x5d\xfe\x75\x7d\x3b\x4b\x1c\x7a\xc8\x30\x49\xb0\xa8\x09\xd2\x2b\x5a\xe9\x86\x64\xf4\xff\x94\x84\x10\x69\xb2\xaa\x42\xc6\xbf\x90\xdd\x40\x5e\x53\x8b\xbb\xdd\xe6\x15\x09\x2f\xad\xa8\x9e\xa1\xf6\xde\xb8\x71\x9e\x3b\x76\x8e\xac\x50\x54\x44\x55\x83\xd2\x28\x17\x86\x9d\xf7\xa4\xfc\x73\x24\x4e\xc4\xb1\x60\x95\xdd\x3a\x93\x6d\x79\x7a\x32\x00\xec\xd8\xff\xd6\xfc\x6e\x3f\x70\xbb\x2e\xec\x5c\x5a\xc8\x2e\x21\xef\x9c\xcd\x1b\x2a\x64\x03\xd9\xfa\x79\xf1\x95\x98\xc4\xb6\x5f\x7e\x1a\x68\xfe\x94\xd1\xaa\xd3\x9b\xdb\x8b\xd9\x1f\xef\x3c\xed\x53\xd8\x62\x66\xd8\x71\x26\x54\x85\x75\xf5\x5f\xb8\x6a\xa5\xe1\x71\x55\x93\x6c\xd5\xe3\x78\xf7\xb0\x97\x38\x60\xb3\x67\x7c\x00\x67\xe3\xbc\xa1\xc7\x2f\x07\xb8\x36\x64\x7d\x8c\x9e\x7f\x28\xcc\xf9\x82\x19\xbf\xf7\xc6\xd1\x18\xf9\x11\xf2\x0e\x60\x32\xf9\xd4\x8c\x98\xf3\x11\xb7\xc5\x3e\x66\xdf\xc9\x47\xd4\xff\xa9\x1c\x84\x0e\x4e\x0b\x5a\xef\x2f\xef\x5c\xbc\xc3\x2a\xe2\x3b\xd1\xbf\x4f\x44\x1a\x9f\x85\xed\x77\x86\xcf\x24\x42\xb6\x81\x49\x38\x31\xb9\xee\xd8\xa2\xc7\x93\x5f\x8f\xf6\xd3\xd4\x60\x5e\xce\xab\xf8\xbe\x9b\x3f\x5b\x68\xd1\x16\x9d\x55\xb2\x49\x06\xbb\xfd\x17\x00\x00\xff\xff\x84\x08\xd5\xa4\xb2\x05\x00\x00"

func dataCommonDevDepVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepVagrantfileTpl,
		"data/common/dev-dep/Vagrantfile.tpl",
	)
}

func dataCommonDevDepVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevDepVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildShTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x91\xc1\x6e\xdb\x30\x10\x44\xef\xfa\x8a\xa9\x7d\xc8\x49\x12\xfa\x01\x2d\xd0\x5e\x72\x6c\x81\xb4\x1f\xb0\xa2\x56\x12\x1b\x99\x4b\xac\x96\x76\xfc\xf7\x5d\x4a\x0e\x12\xc0\x30\x04\x2e\x67\xe6\xed\xf0\xfc\xa5\x1f\x62\xea\x07\xda\x96\xe6\xdc\x9c\xf1\xa3\x98\xb4\x33\x27\x56\x32\x1e\x31\xdc\xf1\xcb\x4c\xba\x7d\xf6\x67\x89\x1b\xfc\x67\x0b\x63\x28\x71\x1d\xb1\x05\x8d\xd9\x30\x89\x82\xf0\x2c\xad\xdb\xb8\x28\xab\xfc\xe3\x60\x5d\xb3\xb1\xa1\xe5\xc6\xa5\x2f\xfe\x65\x55\x6e\x82\x89\x5e\xf9\xf0\xf0\x50\x0d\x90\x84\x4d\x2e\x8c\xbc\x92\xb9\xd3\xa5\x06\x90\xe1\xc6\x4f\xca\x88\xc9\x1c\x25\x58\xbc\xb2\x43\xe0\xef\x50\x92\x15\x3f\x45\x26\xb5\x18\xca\x4a\x5a\x21\x47\x9e\xa8\xac\x2e\x92\xf4\x64\x58\x85\xc6\xcf\x09\x71\x3a\xc2\xe3\xe6\x53\x77\x71\xae\xae\xe1\xb7\x2c\x6a\xf8\xfd\xf2\xf5\xdb\xe9\x3b\x4e\x3b\xa5\x14\x0d\x0c\xff\xaf\x3b\x4c\x71\x65\x27\x73\x10\xcc\x8e\xbf\x9f\x92\x2d\x75\x94\x59\xd7\x7b\xb5\x29\xb9\xe9\xd0\x2f\x8e\xdf\x5f\x69\x56\x4a\xd6\x77\x47\x68\xf5\x7b\x96\xca\x2f\xbb\xf4\x26\xfa\x1a\xd3\x8c\x31\xaa\x77\x23\x7a\x6f\xc2\x88\x77\x51\xbd\xfc\x73\x6f\xb4\x52\x3f\xfa\x03\xa5\x11\x37\x8d\x76\xb4\x25\xc5\x72\xb1\x0f\xc3\x6d\x21\xf5\xb2\x3f\xfc\xce\xb8\x45\xc7\xab\x77\x83\x5c\xb2\xd3\x7f\x9a\xd6\x45\x1e\xb5\x22\x50\x02\xd3\x16\x7d\x05\x7e\xb3\xda\x2e\xa2\x17\x32\xcb\xe3\x55\x5b\xc1\xa9\x17\x7f\xf6\x36\x50\x58\xb8\x1f\xf9\xda\x8e\x9c\xdb\x83\xe0\xd4\xfc\x0f\x00\x00\xff\xff\x13\xa6\xc6\xb3\x34\x02\x00\x00"

func dataCommonDevDepBuildShTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildShTpl,
		"data/common/dev-dep/build.sh.tpl",
	)
}

func dataCommonDevDepBuildShTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildShTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build.sh.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepUpstartConfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcf\x31\x4f\x03\x31\x0c\x05\xe0\xdd\xbf\xe2\xa9\x48\x30\x95\xd0\xc2\x31\xde\xca\xc8\xc8\x80\x3a\xa4\x17\x17\x45\x4a\xed\x28\x76\x5b\x50\xd5\xff\xce\x11\x81\x54\xa6\x28\xef\x49\x9f\xed\xc4\x36\xb5\x5c\x3d\xab\x60\x71\x3e\x43\xe2\x9e\x71\xb9\x60\x89\x17\x16\x6e\xd1\x39\x61\xfb\x85\x57\x77\x5d\x10\x35\xb6\x1a\x4f\xf2\xf7\xa2\xe4\x7d\x76\xac\x06\x0c\x44\xe6\xb1\x39\x66\xa6\x1d\xa4\xf0\x91\x0b\xde\xd7\x8f\x4f\xc3\x66\x2e\xb4\xfe\xcf\x1f\x9e\x37\x44\x37\x78\x63\x24\x95\x3b\xc7\x29\x8a\xc3\x15\x33\xdb\x11\x57\xc5\x2e\x9a\x63\xa7\x0d\x89\xab\x51\x55\xf3\x65\x87\xf8\x93\x27\x58\x61\xae\x7d\x68\x5f\x9e\x80\x70\xb0\x16\x8a\x4e\xb1\x84\x6d\x96\x70\x75\xc9\x38\x86\x63\xfc\xe9\x3e\xae\xd2\xfb\xf9\x8b\xf5\x78\xbb\x22\x96\x84\x5f\xe5\x3b\x00\x00\xff\xff\x9a\xf6\x9d\x7c\x0c\x01\x00\x00"

func dataCommonDevDepUpstartConfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepUpstartConfTpl,
		"data/common/dev-dep/upstart.conf.tpl",
	)
}

func dataCommonDevDepUpstartConfTpl() (*asset, error) {
	bytes, err := dataCommonDevDepUpstartConfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/upstart.conf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/aws-vpc-public-private/build/template.json.tpl": dataAwsVpcPublicPrivateBuildTemplateJsonTpl,
	"data/aws-vpc-public-private/deploy/main.tf.tpl": dataAwsVpcPublicPrivateDeployMainTfTpl,
	"data/aws-vpc-public-private/deploy/variables.tf": dataAwsVpcPublicPrivateDeployVariablesTf,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
	"data/common/dev-dep/Vagrantfile.fragment.tpl": dataCommonDevDepVagrantfileFragmentTpl,
	"data/common/dev-dep/Vagrantfile.tpl": dataCommonDevDepVagrantfileTpl,
	"data/common/dev-dep/build.sh.tpl": dataCommonDevDepBuildShTpl,
	"data/common/dev-dep/upstart.conf.tpl": dataCommonDevDepUpstartConfTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"aws-vpc-public-private": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"template.json.tpl": &bintree{dataAwsVpcPublicPrivateBuildTemplateJsonTpl, map[string]*bintree{
				}},
			}},
			"deploy": &bintree{nil, map[string]*bintree{
				"main.tf.tpl": &bintree{dataAwsVpcPublicPrivateDeployMainTfTpl, map[string]*bintree{
				}},
				"variables.tf": &bintree{dataAwsVpcPublicPrivateDeployVariablesTf, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
			"dev-dep": &bintree{nil, map[string]*bintree{
				"Vagrantfile.fragment.tpl": &bintree{dataCommonDevDepVagrantfileFragmentTpl, map[string]*bintree{
				}},
				"Vagrantfile.tpl": &bintree{dataCommonDevDepVagrantfileTpl, map[string]*bintree{
				}},
				"build.sh.tpl": &bintree{dataCommonDevDepBuildShTpl, map[string]*bintree{
				}},
				"upstart.conf.tpl": &bintree{dataCommonDevDepUpstartConfTpl, map[string]*bintree{
				}},
			}},
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

