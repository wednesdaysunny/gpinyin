package gpinyin

import (
	"bytes"
	"regexp"
)

const (
	data_Chinese_tas  = "chinese"
	data_pinyin       = "pinyin"
	data_multi_pinyin = "multi_pinyin"
)

var traditionalChinese map[string]string
var simplifiedChinese map[string]string
var multiPinyinChinese map[string]string
var pinyinChinese map[string]string
var chineseRegex *regexp.Regexp

func init() {
	chineseRegex = regexp.MustCompile("[\u4e00-\u9fa5]")
}

func ConvertToSimplifiedChinese(source string) string {
	var buf bytes.Buffer
	for _, runeValue := range source {
		buf.WriteString(toSimplifiedChinese(string(runeValue)))
	}
	return buf.String()
}

func ConvertToTraditionalChinese(source string) string {
	var buf bytes.Buffer
	for _, runeValue := range source {
		buf.WriteString(toTraditionalChinese(string(runeValue)))

	}
	return buf.String()
}

func loadMapFromResource(resourceName string, reverse bool) map[string]string {
	v := make(map[string]string)
	err := loadResource(resourceName, v, reverse)
	if err != nil {
		panic(err)
	}
	return v
}

func toSimplifiedChinese(source string) string {
	v := traditionToSimpleMapping[source]
	if len(v) == 0 {
		return source
	}
	return v
}

func toTraditionalChinese(source string) string {
	v := simpleToTraditionMapping[source]
	if len(v) == 0 {
		return source
	}
	return v
}

func toMultiPinyin(source string) string {
	if multiPinyinChinese == nil {
		multiPinyinChinese = loadMapFromResource(data_multi_pinyin, false)
	}
	return multiPinyinChinese[source]
}

func toPinyin(source string) string {
	if pinyinChinese == nil {
		pinyinChinese = loadMapFromResource(data_pinyin, false)
	}
	return pinyinChinese[source]
}

func isChinese(char string) bool {
	return chineseRegex.MatchString(char)
}
