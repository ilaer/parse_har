package util

import "golang.org/x/text/encoding/simplifiedchinese"

/*
description:convert the command`s stdout byte data to string
*/

func ConvertByte2String(byte []byte, charset string) string {

	var str string
	switch charset {
	case "GB18030":
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case "UTF-8":
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
