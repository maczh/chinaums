package chinaums

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func ToJSON(o any) string {
	j, err := json.Marshal(o)
	if err != nil {
		return "{}"
	} else {
		js := string(j)
		js = strings.Replace(js, "\\u003c", "<", -1)
		js = strings.Replace(js, "\\u003e", ">", -1)
		js = strings.Replace(js, "\\u0026", "&", -1)
		return js
	}
}

func ToJSONPretty(o any) string {
	j, err := json.Marshal(o)
	if err != nil {
		return "{}"
	} else {
		js := string(j)
		js = strings.Replace(js, "\\u003c", "<", -1)
		js = strings.Replace(js, "\\u003e", ">", -1)
		js = strings.Replace(js, "\\u0026", "&", -1)
		js = JSONPretty(js, "", "  ")
		return js
	}
}

func FromJSON(j string, o any) *any {
	err := json.Unmarshal([]byte(j), &o)
	if err != nil {
		fmt.Printf("数据转换错误:%s", err.Error())
		return nil
	} else {
		return &o
	}
}

// JSONPrettyPrint pretty print raw json string to indent string
func JSONPretty(in, prefix, indent string) string {
	var out bytes.Buffer
	if err := json.Indent(&out, []byte(in), prefix, indent); err != nil {
		return in
	}
	return out.String()
}

// CompactJSON compact json input with insignificant space characters elided
func CompactJSON(in string) string {
	var out bytes.Buffer
	if err := json.Compact(&out, []byte(in)); err != nil {
		return in
	}
	return out.String()
}

func getRandomIntString(l int) string {
	str := "0123456789"
	return generateRandString(str, l)
}

func generateRandString(source string, l int) string {
	bytes := []byte(source)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
