/*
 * Mini Object Storage, (C) 2014 Minio, Inc.
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

package strbyteconv

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	UNIT_BYTE     = 1
	UNIT_KILOBYTE = 1024 * UNIT_BYTE
	UNIT_MEGABYTE = 1024 * UNIT_KILOBYTE
	UNIT_GIGABYTE = 1024 * UNIT_MEGABYTE
	UNIT_TERABYTE = 1024 * UNIT_GIGABYTE
)

func BytesToString(bytes uint64) string {
	var unit string = "B"
	var value uint64 = 0

	switch {
	case bytes >= UNIT_TERABYTE:
		unit = "TB"
		value = uint64(bytes / UNIT_TERABYTE)
	case bytes >= UNIT_GIGABYTE:
		unit = "GB"
		value = uint64(bytes / UNIT_GIGABYTE)
	case bytes >= UNIT_MEGABYTE:
		unit = "MB"
		value = uint64(bytes / UNIT_MEGABYTE)
	case bytes >= UNIT_KILOBYTE:
		unit = "KB"
		value = uint64(bytes / UNIT_KILOBYTE)
	case bytes < UNIT_KILOBYTE && bytes >= UNIT_BYTE:
		unit = "B"
		value = uint64(bytes / UNIT_BYTE)
	}

	return fmt.Sprintf("%d%s", value, unit)
}

func StringToBytes(s string) (uint64, error) {
	var bytes uint64

	StringPattern, err := regexp.Compile(`(?i)^(-?\d+)([KMGT])B?$`)
	if err != nil {
		return 0, err
	}

	parts := StringPattern.FindStringSubmatch(strings.TrimSpace(s))
	if len(parts) < 3 {
		return 0, errors.New("Incorrect string format must be K,KB,M,MB,G,GB")
	}

	value, err := strconv.ParseUint(parts[1], 10, 0)
	if err != nil || value < 1 {
		return 0, err
	}

	unit := strings.ToUpper(parts[2])
	switch unit {
	case "T":
		bytes = value * UNIT_TERABYTE
	case "G":
		bytes = value * UNIT_GIGABYTE
	case "M":
		bytes = value * UNIT_MEGABYTE
	case "K":
		bytes = value * UNIT_KILOBYTE
	case "B":
		bytes = value * UNIT_BYTE
	}

	return bytes, nil
}
