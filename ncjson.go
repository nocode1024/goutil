package main

import "github.com/tidwall/gjson"

// GetStrArrB  returns the string array for the given path.
func GetStrArrB(jsonStr []byte, path string) []string {
	return GetStrArr(string(jsonStr), path)
}

// GetStrArr  returns the string array for the given path.
func GetStrArr(jsonStr string, path string) []string {
	var result []string
	gjson.Get(jsonStr, path).ForEach(func(key, value gjson.Result) bool {
		result = append(result, value.String())
		return true
	})
	return result
}

// GetStr  returns the string value for the given path.
func GetStr(jsonStr string, path string) string {
	return GetStrB([]byte(jsonStr), path)
}

// GetStrB  returns the string value for the given path.
func GetStrB(jsonStr []byte, path string) string {
	result := gjson.GetBytes(jsonStr, path)
	if result.Exists() {
		return result.String()
	}
	return ""
}

// GetInt  returns the int value for the given path.
func GetInt(jsonStr string, path string) int64 {
	result := gjson.Get(jsonStr, path)
	if result.Exists() {
		return result.Int()
	}
	return 0
}

// GetIntB  returns the int value for the given path.
func GetIntB(jsonStr []byte, path string) int64 {
	return GetInt(string(jsonStr), path)
}

// GetObject  returns the interface{} for the given path.
func GetObject(jsonStr string, path string) interface{} {
	result := gjson.Get(jsonStr, path)
	if result.Exists() {
		return result
	}
	return nil
}

// GetObjectB  returns the interface{} for the given path.
func GetObjectB(jsonStr []byte, path string) interface{} {
	return GetObject(string(jsonStr), path)
}
