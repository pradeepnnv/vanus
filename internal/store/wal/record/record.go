// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package record

import (
	// standard libraries.
	"bytes"
	"encoding/binary"
	"hash/crc32"
)

var crc32q = crc32.MakeTable(crc32.Castagnoli)

type Type uint8

const (
	Zero Type = iota
	Full
	First
	Last
	Middle
)

type Record struct {
	// CRC is crc32c of Type and Data
	CRC uint32
	// Length is len(Data). optimize?
	Length uint16
	Type   Type
	Data   []byte
}

func (r *Record) Size() int {
	return 4 + 2 + 1 + len(r.Data)
}

func (r *Record) Marshal() ([]byte, error) {
	data := make([]byte, r.Size())
	if _, err := r.MarshalTo(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Record) MarshalTo(data []byte) (int, error) {
	sz := r.Size()
	if len(data) < sz {
		// TODO(james.yin): correct error.
		return 0, bytes.ErrTooLarge
	}
	binary.BigEndian.PutUint16(data[4:6], r.Length)
	data[6] = byte(r.Type)
	ds := len(r.Data)
	if ds > 0 {
		copy(data[7:7+ds], r.Data)
	}
	// calculate CRC
	if r.CRC == 0 {
		r.CRC = crc32.Checksum(data[6:7+ds], crc32q)
	}
	binary.BigEndian.PutUint32(data[0:4], r.CRC)
	return sz, nil
}