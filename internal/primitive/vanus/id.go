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

package vanus

import (
	"strconv"
	"time"
)

type ID uint64

var (
	emptyID = ID(0)
)

func EmptyID() ID {
	return emptyID
}
func NewID() ID {
	return ID(time.Now().UnixNano())
}

func NewIDFromUint64(id uint64) ID {
	return ID(id)
}

func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func (id ID) Uint64() uint64 {
	return uint64(id)
}

func GenerateID() ID {
	return ID(time.Now().UnixNano())
}

func StringToID(str string) (ID, error) {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return ID(id), nil
}
