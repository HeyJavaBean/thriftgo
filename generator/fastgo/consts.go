/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fastgo

import (
	"strings"

	"github.com/cloudwego/thriftgo/parser"
	"github.com/cloudwego/thriftgo/version"
)

const cloudwegoRepoPrefix = "github.com/cloudwego/"

var fixedFileHeader string

func init() {
	fixedFileHeader = strings.Replace(
		`// Code generated by thriftgo ({{Version}}) (fastgo). DO NOT EDIT.`,
		"{{Version}}",
		version.ThriftgoVersion, 1)
}

const ( // wiretypes
	tSTOP   = 0
	tVOID   = 1
	tBOOL   = 2
	tBYTE   = 3
	tI08    = 3
	tDOUBLE = 4
	tI16    = 6
	tI32    = 8
	tI64    = 10
	tSTRING = 11
	tUTF7   = 11
	tSTRUCT = 12
	tMAP    = 13
	tSET    = 14
	tLIST   = 15
	tUTF8   = 16
	tUTF16  = 17
)

var category2ThriftWireType = [16]int{
	// 0-15, panic if Category_Typedef or Category_Service
	parser.Category_Bool:      tBOOL,
	parser.Category_Byte:      tI08,
	parser.Category_I16:       tI16,
	parser.Category_I32:       tI32,
	parser.Category_I64:       tI64,
	parser.Category_Double:    tDOUBLE,
	parser.Category_String:    tSTRING,
	parser.Category_Binary:    tSTRING,
	parser.Category_Map:       tMAP,
	parser.Category_List:      tLIST,
	parser.Category_Set:       tSET,
	parser.Category_Enum:      tI32,
	parser.Category_Struct:    tSTRUCT,
	parser.Category_Union:     tSTRUCT,
	parser.Category_Exception: tSTRUCT,
}

var category2GopkgConsts = [16]string{
	// 0-15, panic if Category_Typedef or Category_Service
	parser.Category_Bool:      "thrift.BOOL",
	parser.Category_Byte:      "thrift.I08",
	parser.Category_I16:       "thrift.I16",
	parser.Category_I32:       "thrift.I32",
	parser.Category_I64:       "thrift.I64",
	parser.Category_Double:    "thrift.DOUBLE",
	parser.Category_String:    "thrift.STRING",
	parser.Category_Binary:    "thrift.STRING",
	parser.Category_Map:       "thrift.MAP",
	parser.Category_List:      "thrift.LIST",
	parser.Category_Set:       "thrift.SET",
	parser.Category_Enum:      "thrift.I32",
	parser.Category_Struct:    "thrift.STRUCT",
	parser.Category_Union:     "thrift.STRUCT",
	parser.Category_Exception: "thrift.STRUCT",
}

var category2WireSize = [16]int{
	parser.Category_Bool:   1,
	parser.Category_Byte:   1,
	parser.Category_I16:    2,
	parser.Category_I32:    4,
	parser.Category_Enum:   4,
	parser.Category_I64:    8,
	parser.Category_Double: 8,
}
