// Code generated by thriftgo (0.3.2). DO NOT EDIT.

package option_test

import (
	"reflect"

	"github.com/cloudwego/thriftgo/thrift_reflection"
)

// IDL Name: test
// IDL Path: option_idl/test.thrift

var file_test_thrift_go_types = []interface{}{
	(*IDCard)(nil), // Struct 0: option_test.IDCard
	(*Person)(nil), // Struct 1: option_test.Person
	(*MyEnum)(nil), // Enum 0: option_test.MyEnum
}
var file_test_thrift *thrift_reflection.FileDescriptor
var file_idl_test_rawDesc = []byte{
	0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xcc, 0x56, 0xdf, 0x6b, 0xeb, 0x36,
	0x14, 0x3e, 0x71, 0xda, 0xdc, 0x36, 0x6e, 0xe9, 0xdb, 0xde, 0x6, 0x9e, 0x61, 0xf4, 0x61, 0xa1,
	0x17, 0xff, 0x68, 0xe8, 0x34, 0xee, 0xd8, 0xee, 0xb6, 0x87, 0xb, 0xcb, 0x7e, 0x5d, 0x18, 0x17,
	0x42, 0x56, 0x94, 0xf8, 0xa4, 0x11, 0xb5, 0xe5, 0xcc, 0x96, 0xc3, 0x42, 0xc8, 0xff, 0x3e, 0x24,
	0xcb, 0x89, 0x1d, 0x3b, 0xc1, 0x81, 0x3d, 0xcc, 0x2f, 0x4a, 0xa4, 0xef, 0xe8, 0x7c, 0xdf, 0xa7,
	0x63, 0xeb, 0x98, 0xd0, 0x1, 0x80, 0xcf, 0xe2, 0xa5, 0x60, 0x31, 0x7f, 0x66, 0x41, 0xf8, 0x56,
	0x60, 0x2a, 0x1e, 0xc4, 0x22, 0x61, 0x73, 0x71, 0xb, 0x86, 0x69, 0x2, 0x80, 0x1, 0x0, 0x3d,
	0xe4, 0x82, 0x89, 0x35, 0x0, 0x7c, 0x55, 0x2, 0x53, 0xce, 0x63, 0x41, 0xe5, 0xdf, 0xf4, 0x6d,
	0xe, 0xd0, 0x83, 0xde, 0x1, 0x0, 0xfa, 0x2b, 0x1a, 0xb2, 0x40, 0x61, 0x0, 0xc0, 0x3b, 0x12,
	0xbc, 0x7, 0x95, 0x7e, 0xee, 0x68, 0x74, 0x15, 0xd, 0xc9, 0xd4, 0x78, 0x89, 0x1, 0xc0, 0xd4,
	0xbb, 0x48, 0xae, 0x77, 0x70, 0x71, 0x23, 0x17, 0x4f, 0x2a, 0x31, 0x95, 0x86, 0xeb, 0xd1, 0xfa,
	0x23, 0x26, 0x2b, 0x36, 0xc3, 0x3b, 0xe8, 0xca, 0x20, 0xa3, 0x45, 0x90, 0x31, 0x72, 0x6e, 0xa0,
	0xdb, 0x2, 0xd8, 0x4b, 0x45, 0xc2, 0xf8, 0xb, 0x68, 0x42, 0x70, 0xb, 0x97, 0xe6, 0x9d, 0xa6,
	0xdd, 0xcb, 0x3, 0xb, 0x15, 0x93, 0x92, 0xc6, 0x8, 0xc5, 0x22, 0xe, 0x9e, 0x19, 0x9f, 0xc7,
	0xd6, 0x3b, 0x6b, 0xd3, 0xb7, 0x4a, 0xf, 0xa7, 0x11, 0x12, 0x6b, 0xa4, 0x10, 0x1f, 0xf8, 0x3c,
	0xfe, 0x85, 0x46, 0x58, 0x5, 0x64, 0xd1, 0x14, 0x13, 0x62, 0x3d, 0x3e, 0x3e, 0xee, 0xe6, 0xb7,
	0x26, 0xf4, 0x64, 0xfe, 0x3b, 0x78, 0xa3, 0x78, 0x18, 0x70, 0x5, 0xd0, 0x4a, 0xa8, 0xfb, 0xff,
	0x17, 0xea, 0xfb, 0xfe, 0x69, 0xa1, 0xb7, 0x70, 0x71, 0x8c, 0xcc, 0xaf, 0x25, 0x32, 0xe9, 0x6a,
	0x56, 0x67, 0x92, 0xb3, 0xd0, 0x25, 0x52, 0xa3, 0x51, 0x50, 0x18, 0xe, 0x87, 0x7d, 0x9d, 0xfe,
	0x12, 0xf2, 0xfc, 0x97, 0x6d, 0x8b, 0xa9, 0xf7, 0xe1, 0xc7, 0x1f, 0x68, 0x12, 0x9c, 0x51, 0x7e,
	0xbd, 0x3c, 0xed, 0x79, 0x27, 0x63, 0xc2, 0x5, 0x0, 0x5c, 0xfd, 0x81, 0x7f, 0x67, 0x2c, 0xc1,
	0xe0, 0x4a, 0x11, 0xed, 0xdc, 0xc2, 0x1b, 0x65, 0xd, 0x98, 0xd2, 0x28, 0x68, 0x55, 0x14, 0x5d,
	0xfa, 0x82, 0xed, 0x72, 0x1b, 0xec, 0xa9, 0x39, 0xaf, 0x51, 0xcb, 0x5b, 0x9c, 0x11, 0x14, 0xe,
	0xb6, 0x91, 0xf6, 0x1b, 0x26, 0x69, 0xcc, 0xcf, 0x70, 0xee, 0x42, 0x1e, 0xe7, 0x7f, 0xe9, 0x5b,
	0x43, 0x49, 0xb9, 0xfa, 0x6b, 0xb7, 0x54, 0xe4, 0x9e, 0xe7, 0xc, 0xc3, 0xbc, 0xc2, 0xdf, 0xdd,
	0x8b, 0x5, 0xaa, 0x82, 0xb2, 0xe2, 0xb9, 0x25, 0x16, 0x2c, 0xb5, 0x72, 0xcc, 0xfd, 0x19, 0xe6,
	0x1b, 0x2c, 0x68, 0xc9, 0x3f, 0x2f, 0xab, 0x73, 0xfd, 0x3f, 0x10, 0x74, 0xd, 0xd0, 0xf9, 0xbc,
	0x2a, 0x68, 0x4a, 0x53, 0xd6, 0xf0, 0xa2, 0xc8, 0x67, 0x45, 0xc3, 0xc, 0xd9, 0x13, 0x79, 0x6a,
	0x98, 0x75, 0x86, 0xc4, 0x19, 0x36, 0xcc, 0x7b, 0x2e, 0xf1, 0xdc, 0x86, 0xf9, 0xa1, 0x4f, 0x86,
	0x7e, 0x7d, 0x3e, 0x3f, 0x15, 0x72, 0x8f, 0xff, 0xd0, 0x68, 0x19, 0xe2, 0x77, 0x18, 0x51, 0x16,
	0x3e, 0xcc, 0xe2, 0xe8, 0xbe, 0x8e, 0x9d, 0xae, 0x5, 0x12, 0xcb, 0x69, 0x58, 0x60, 0x9c, 0x26,
	0x6b, 0x62, 0x39, 0xd, 0x89, 0x83, 0x38, 0x9b, 0x86, 0x48, 0xbc, 0x7, 0xc7, 0x77, 0x1e, 0xbf,
	0x6e, 0x88, 0x8d, 0xe3, 0x90, 0x58, 0x22, 0xc9, 0xf2, 0xaf, 0xc0, 0x16, 0xa0, 0xf3, 0x57, 0xd5,
	0xa0, 0x54, 0x24, 0xd9, 0x4c, 0x9c, 0x70, 0x28, 0x7, 0x90, 0x8d, 0xe2, 0x4e, 0x6c, 0x8c, 0x96,
	0x62, 0x6d, 0xa9, 0x3f, 0xf6, 0xb6, 0x8e, 0x96, 0x27, 0x5b, 0x44, 0x54, 0x56, 0xf5, 0xb7, 0xc9,
	0xe, 0x11, 0xed, 0xea, 0x2, 0xe3, 0x1c, 0x93, 0x8f, 0x8d, 0x31, 0xf2, 0xd1, 0x79, 0x79, 0xac,
	0x93, 0x1e, 0x20, 0xaa, 0x1c, 0x1a, 0x18, 0x21, 0xcf, 0x22, 0x62, 0xbd, 0x3f, 0x26, 0x4c, 0xac,
	0x97, 0x18, 0xe0, 0xbc, 0xad, 0x3e, 0x55, 0x4f, 0x45, 0x8c, 0x65, 0x2f, 0x30, 0xc, 0x63, 0x4b,
	0x2c, 0x30, 0xd1, 0xaa, 0xa4, 0xc3, 0xbf, 0x57, 0x1d, 0x9e, 0xc5, 0x5c, 0x50, 0xc6, 0x31, 0x39,
	0x61, 0x72, 0x44, 0x97, 0x64, 0x63, 0x2f, 0x70, 0xed, 0xd8, 0xc4, 0x56, 0x33, 0x4e, 0x53, 0xf6,
	0x90, 0xa5, 0x82, 0x8c, 0x6d, 0x39, 0x38, 0xf6, 0x40, 0x8d, 0xae, 0x3d, 0x69, 0x90, 0x86, 0x5,
	0xcc, 0xd3, 0x30, 0xbf, 0x9, 0x26, 0x17, 0x14, 0x74, 0x4c, 0x7, 0xd3, 0xc1, 0x6c, 0x32, 0x18,
	0x7, 0x3, 0x1c, 0xcc, 0x27, 0xc7, 0xa1, 0xfa, 0x68, 0xc7, 0x63, 0x6d, 0x17, 0x3a, 0xdb, 0x41,
	0xf1, 0xd3, 0xdd, 0x4e, 0x6, 0xbb, 0x79, 0x6f, 0x3f, 0xef, 0x6f, 0x9b, 0x36, 0x8c, 0xe8, 0xb2,
	0xa8, 0x93, 0x57, 0x87, 0xec, 0xb7, 0xb3, 0x5e, 0x5d, 0xb2, 0xdf, 0x71, 0x5b, 0x98, 0xa, 0x5f,
	0x96, 0xee, 0xbe, 0x7d, 0xe9, 0x32, 0xfe, 0x52, 0xb8, 0xaa, 0x8e, 0x2, 0x0, 0x1e, 0xea, 0xc0,
	0x88, 0x2e, 0x77, 0xde, 0x1f, 0x9a, 0xc, 0x0, 0x5f, 0xd4, 0x23, 0x64, 0xd1, 0x14, 0x21, 0x9f,
	0x3e, 0xfd, 0xc, 0x0, 0x7e, 0x1d, 0x94, 0x7f, 0x5b, 0x74, 0x31, 0x14, 0xe8, 0x4a, 0x49, 0x0,
	0xc0, 0x37, 0x8d, 0xbc, 0xe5, 0x2b, 0x77, 0x10, 0xb8, 0x51, 0x2f, 0x88, 0xae, 0x3f, 0xf9, 0x5b,
	0x71, 0xfb, 0xf6, 0x68, 0x78, 0x80, 0x73, 0x9a, 0x85, 0xe2, 0x59, 0x9, 0xd9, 0x6d, 0xb2, 0x72,
	0x88, 0xbd, 0x72, 0xac, 0xdc, 0x1a, 0xbb, 0x74, 0xc3, 0xf7, 0x6e, 0x4a, 0x9d, 0x46, 0x9b, 0x5e,
	0xb3, 0x37, 0x5a, 0xff, 0xc4, 0xb3, 0xe8, 0x8c, 0xfb, 0xaa, 0xf3, 0x7d, 0x1f, 0xba, 0xa0, 0x9f,
	0x13, 0x6d, 0x4c, 0x50, 0xd2, 0xa4, 0x8c, 0xae, 0x4a, 0x68, 0xe8, 0xab, 0x24, 0x91, 0x3f, 0x25,
	0xe8, 0x64, 0x6b, 0xe5, 0xba, 0x6e, 0xb9, 0xb5, 0x6a, 0x7d, 0x33, 0x77, 0xde, 0xef, 0x79, 0x77,
	0x6a, 0x57, 0xfb, 0x9, 0x21, 0xa3, 0x43, 0x21, 0x47, 0x1a, 0x32, 0x49, 0xff, 0x68, 0x37, 0xe6,
	0x79, 0xde, 0x61, 0x37, 0x76, 0xa5, 0xcf, 0xea, 0x5a, 0x8f, 0x7d, 0x35, 0xc2, 0xbf, 0x1, 0x0,
	0x0, 0xff, 0xff, 0x7, 0x9e, 0x5e, 0x2b, 0xeb, 0xc, 0x0, 0x0,
}

func init() {
	if file_test_thrift != nil {
		return
	}
	type x struct{}
	builder := &thrift_reflection.FileDescriptorBuilder{
		Bytes:         file_idl_test_rawDesc,
		GoTypes:       file_test_thrift_go_types,
		GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
	}
	file_test_thrift = thrift_reflection.BuildFileDescriptor(builder)
}

func GetFileDescriptorForTest() *thrift_reflection.FileDescriptor {
	return file_test_thrift
}
func (p *IDCard) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_test_thrift.GetStructDescriptor("IDCard")
}
func (p *Person) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_test_thrift.GetStructDescriptor("Person")
}
func (p MyEnum) GetDescriptor() *thrift_reflection.EnumDescriptor {
	return file_test_thrift.GetEnumDescriptor("MyEnum")
}
