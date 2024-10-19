package utils

import "strings"

const ShortImageIDLength = 12

// ShortImageID - конвертирует длинное представление id образа в короткое
// sha256:2503e324e27050f4d3fd21d25147ca108840864941d96097c2633cd9232f5088
// 2503e324e270
func ShortImageID(id string) string {

	split_result := strings.Split(id, ":")

	if len(split_result) != 2 {
		return id
	}

	if len(split_result[1]) < 12 {
		return id
	}

	return split_result[1][:ShortImageIDLength]
}
