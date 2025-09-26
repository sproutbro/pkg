package converter

import (
	"encoding/json"
	"strconv"
)

func ByteToStruct(b []byte, t any) error {
	return json.Unmarshal(b, &t)
}

func StructToByte(t any) ([]byte, error) {
	return json.Marshal(t)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
