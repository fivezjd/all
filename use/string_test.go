package use

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// 字符串相关处理
//len(s string) int：返回字符串 s 的长度。

// 获取字符串长度(返回的是字节)
func TestLen(t *testing.T) {
	str := "abc"
	t.Log(len(str))
	// out : 3

	strChinese := "你好"
	t.Log(len(strChinese))
	// out: 6
}

// 返回unicode字符串长度
func TestUnicodeLen(t *testing.T) {
	str := "abc"
	t.Log(utf8.RuneCountInString(str))
	// out : 3

	strChinese := "你好"
	t.Log(utf8.RuneCountInString(strChinese))
	// out: 2
}

//
//s[i] byte：获取字符串 s 的第 i 个字符。
// 这样遍历unicode 字符串的方式就成了下面这样：
//

func TestForUnicode(t *testing.T) {
	str := "你好1234567"
	count := utf8.RuneCountInString(str)

	for i := 0; i < count; i++ {
		t.Log(str[i])
	}
	// 上面的遍历输出字符串会出现乱码
}

// 避免乱码遍历字符串
func TestStringToRune(t *testing.T) {
	str := "你好1234567"
	count := utf8.RuneCountInString(str)
	strRune := []rune(str)
	for i := 0; i < count; i++ {
		t.Log(string(strRune[i]))
	}
}

// 避免乱码遍历字符串
func TestStringToRune1(t *testing.T) {
	str := "你好1234567"
	for _, v := range str {
		t.Log(string(v))
	}
}

//s + t：连接字符串 s 和 t。 字符串的拼接 【简单】

func TestAddStr(t *testing.T) {
	str1 := "a"
	str2 := "b"
	t.Log(str1 + str2)
}

// 字符串拼接 【高效】

func TestAddStrB(t *testing.T) {
	str1 := "a"
	str2 := "b"
	var strB strings.Builder
	strB.WriteString(str1)
	strB.WriteString(str2)
	t.Log(strB.String())
}

//s[i:j]：获取 s 中从第 i 个字符到第 j-1 个字符组成的子串。
//
//strings.Contains(s, substr string) bool：判断字符串 s 是否包含子串 substr。
//
//strings.HasPrefix(s, prefix string) bool：判断字符串 s 是否以 prefix 开头。
//
//strings.HasSuffix(s, suffix string) bool：判断字符串 s 是否以 suffix 结尾。
//
//strings.Index(s, sep string) int：查找字符串 s 中第一次出现子串 sep 的位置。
//
//strings.LastIndex(s, sep string) int：查找字符串 s 中最后一次出现子串 sep 的位置。
//
//strings.Split(s, sep string) []string：将字符串 s 按照分隔符 sep 分割成多个子串，并返回一个字符串切片。
//
//strings.Join(a []string, sep string) string：将字符串切片 a 按照分隔符 sep 连接成一个字符串。
//
//strconv.Atoi(s string) (int, error)：将字符串 s 转换为 int 类型。
//
//strconv.ParseFloat(s string, bitSize int) (float64, error)：将字符串 s 转换为 float64 类型。
//
//strconv.FormatInt(i int64, base int) string：将 int64 类型的 i 转换为字符串。
//
//strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string：将 float64 类型的 f 转换为字符串。
//
//fmt.Sprintf(format string, a ...interface{}) string：将多个值格式化为字符串，类似于 C 语言中的 sprintf 函数。
//
//使用 strings.Builder 类型：strings.Builder 是 Golang 提供的一个专门用于字符串拼接的类型。它支持高效的字符串拼接操作，避免了每次创建新的字符串对象，从而提高了性能。可以使用 strings.Builder 的 WriteString 方法将多个字符串拼接起来。
//
//使用 bytes.Buffer 类型：bytes.Buffer 是 Golang 提供的一个类似于 strings.Builder 的类型。它也支持高效的字符串拼接操作，但是与 strings.Builder 不同的是，它可以处理二进制数据，而不仅仅是字符串。
//
//使用 fmt.Sprintf 函数：fmt.Sprintf 函数可以将多个值格式化为字符串，并返回一个字符串。可以使用这个函数来进行字符串拼接操作。
//
//使用 strings.Join 函数：strings.Join 函数可以将一个字符串切片拼接成一个字符串，使用这个函数可以高效地进行字符串拼接操作。
//
//以上是一些在 Golang 中高效进行字符串拼接的方法，其中使用 strings.Builder 和 bytes.Buffer 类型可以获得最好的性能。
//
//以下是 Golang 中常用的字符串处理官方包以及它们的主要功能点：
//
//strings 包：提供了许多字符串处理的基本函数，包括字符串拼接、替换、比较、截取、分割等。
//
//strconv 包：提供了字符串和基本数据类型之间的相互转换函数，包括字符串转整型、浮点型、布尔型等，以及整型、浮点型、布尔型等转字符串的函数。
//
//regexp 包：提供了正则表达式的支持，可以用来进行字符串匹配、查找、替换等操作。
//
//unicode 包：提供了 Unicode 字符集的支持，可以用来进行 Unicode 字符集的转换、查找、分类等操作。
//
//bytes 包：提供了对字节切片的操作，包括字节切片的拼接、替换、查找、截取等。
//
//bufio 包：提供了对缓冲区的支持，可以用来进行高效的文件读取和写入。
