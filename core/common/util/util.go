// Package util contains some common functions of GO_SPIDER project.
package util

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"

	"golang.org/x/net/html/charset"
)

// JsonpToJson modify jsonp string to json string
// Example: forbar({a:"1",b:2}) to {"a":"1","b":2}
func JsonpToJson(json string) string {
	start := strings.Index(json, "{")
	end := strings.LastIndex(json, "}")
	start1 := strings.Index(json, "[")
	if start1 > 0 && start > start1 {
		start = start1
		end = strings.LastIndex(json, "]")
	}
	if end > start && end != -1 && start != -1 {
		json = json[start : end+1]
	}
	json = strings.Replace(json, "\\'", "", -1)
	regDetail, _ := regexp.Compile("([^\\s\\:\\{\\,\\d\"]+|[a-z][a-z\\d]*)\\s*\\:")
	return regDetail.ReplaceAllString(json, "\"$1\":")
}

// Load json_str -> simplejson.Json
func Load(jsonStr string) *simplejson.Json {
	res, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		log.Error(err)
		return nil
	}
	return res
}

// Dump simplejson.Json -> json_str
func Dump(jsonObject *simplejson.Json) string {
	jsonStr, err := json.Marshal(jsonObject)
	// res, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		log.Error(err)
		return ""
	}
	return string(jsonStr)
}

//GetWDPath The GetWDPath gets the work directory path.
func GetWDPath() string {
	wd := os.Getenv("GOPATH")
	if wd == "" {
		panic("GOPATH is not setted in env.")
	}
	return wd
}

//IsDirExists The IsDirExists judges path is directory or not.
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("util isDirExists not reached")
}

//IsFileExists The IsFileExists judges path is file or not.
func IsFileExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return !fi.IsDir()
	}

	panic("util isFileExists not reached")
}

//IsNum The IsNum judges string is number or not.
func IsNum(a string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(a)
}

// XML2mapstr simple xml to string  support utf8
func XML2mapstr(xmldoc string) map[string]string {
	var t xml.Token
	var err error
	inputReader := strings.NewReader(xmldoc)
	decoder := xml.NewDecoder(inputReader)
	decoder.CharsetReader = func(s string, r io.Reader) (io.Reader, error) {
		return charset.NewReader(r, s)
	}
	m := make(map[string]string, 32)
	key := ""
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			key = token.Name.Local
		case xml.CharData:
			content := string([]byte(token))
			m[key] = content
		default:
			// ...
		}
	}

	return m
}

//MakeHash string to hash
func MakeHash(s string) string {
	const IEEE = 0xedb88320
	IEEETable := crc32.MakeTable(IEEE)
	hash := fmt.Sprintf("%x", crc32.Checksum([]byte(s), IEEETable))
	return hash
}
