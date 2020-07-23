package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderscoreToUpperCamelCase(s string) string { //下划线单词转大写驼峰单词
	s = strings.Replace(s, "_", " ", -1)   //将下划线替换成空格
	s = strings.Title(s)                   //将所有字符修改为其对应的首字母大写的形式
	return strings.Replace(s, " ", "", -1) //将空格字符替换为空
}

func UnderscoreToLowerCamelCase(s string) string { //下划线单词转小写驼峰单词
	s = UnderscoreToUpperCamelCase(s)                  //直接复用大写驼峰单词的转换方式
	return string(unicode.ToLower(rune(s[0]))) + s[1:] //将字符串的第一位取出来进行小写转换
}

func CamelCaseToUnderscore(s string) string { //驼峰单词转下划线单词
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
