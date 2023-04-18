package use

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"
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
/**
字符串的截取
*/

// 获取字符串全部数据
func TestAllStr(t *testing.T) {

	// 英文
	str := "abcdesf"
	t.Log(str[:])

	// 中文
	strC := "中文"
	t.Log(strC[:])

	// 中文+英文
	strMix := "中文abc"

	t.Log(strMix[:])

}

//strings.Contains(s, substr string) bool：判断字符串 s 是否包含子串 substr。
//

func TestStrContains(t *testing.T) {
	str := "abcdefs"
	subStr := "de"
	t.Log(strings.Contains(str, subStr))
	// bool
}

//strings.HasPrefix(s, prefix string) bool：判断字符串 s 是否以 prefix 开头。
//

func TestHasPrefix(t *testing.T) {
	str := "abcdeddkff"
	t.Log(strings.HasPrefix(str, "ab"))
}

//strings.HasSuffix(s, suffix string) bool：判断字符串 s 是否以 suffix 结尾。
//

func TestHasSuffix(t *testing.T) {
	str := "abcdeddkff"
	t.Log(strings.HasSuffix(str, "kff"))
}

//strings.Index(s, sep string) int：查找字符串 s 中第一次出现子串 sep 的位置。
// tips: 按字节来查找的
//

func TestIndex(t *testing.T) {
	str := "abcdefghijklmn"
	t.Log(strings.Index(str, "cd"))

	// 测试中文
	strC := "中文abc你好"
	t.Log(strings.Index(strC, "你"))
	// out: 9
	// 在使用中文的时候应该注意
}

// 上面按字节来查找可能无法兼容rune,所以可以先将字符串转为[]rune 然后再获取查找到的子字符串索引

func TestRuneStr(t *testing.T) {
	// 测试中文
	strC := "中文abc你好"

	// search
	strS := "你"
	runeC := []rune(strC)

	for i, v := range runeC {
		if string(v) == strS {
			t.Log(i)
			break
		}
	}
}

//strings.LastIndex(s, sep string) int：查找字符串 s 中最后一次出现子串 sep 的位置。
//

func TestLastIndex(t *testing.T) {
	str := "abcdefghijk"
	t.Log(strings.LastIndexFunc(str, func(r rune) bool {
		return string(r) == "i"
	}))
}

func TestLastIndex1(t *testing.T) {
	str := "abcde中国ghijk"
	t.Log(strings.LastIndexFunc(str, func(r rune) bool {
		return string(r) == "国"
	}))
}

func TestLastIndex2(t *testing.T) {
	str := "中国ABC哈哈"
	t.Log(strings.IndexByte(str, 'A'))

	t.Log(strings.IndexAny(str, "哈"))

	t.Log(strings.IndexFunc(str, func(r rune) bool {
		return string(r) == "国"
	}))

	t.Log(len(str))
}

//strings.Split(s, sep string) []string：将字符串 s 按照分隔符 sep 分割成多个子串，并返回一个字符串切片。
//

func TestSplit(t *testing.T) {
	str := "a,b,c,d,e,f"
	t.Log(strings.Split(str, ","))
}

//strings.Join(a []string, sep string) string：将字符串切片 a 按照分隔符 sep 连接成一个字符串。
//

func TestJoin(t *testing.T) {
	str := []string{"a", "b", "c"}
	t.Log(strings.Join(str, ","))
}

//strconv.Atoi(s string) (int, error)：将字符串 s 转换为 int 类型。

// 字符串转为int  tips : 只能是数字字符串

func TestAtoi(t *testing.T) {
	str := "123"
	v, err := strconv.Atoi(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

//
//strconv.ParseFloat(s string, bitSize int) (float64, error)：将字符串 s 转换为 float64 类型。
//

func TestParseFloat(t *testing.T) {
	//s必须是数字字符串
	str := "1234"
	t.Log(strconv.ParseFloat(str, 64))
}

//strconv.FormatInt(i int64, base int) string：将 int64 类型的 i 转换为字符串。

func TestFormatInt(t *testing.T) {
	// strconv.FormatInt 函数的第二个参数指定了整数应该被表示成字符串的进制。
	i := 123456677
	t.Log(strconv.FormatInt(int64(i), 10))
}

// strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string：将 float64 类型的 f 转换为字符串。
//
// 其中，参数 f 表示要转换的浮点数，参数 fmt 表示输出格式，取值范围为 'f', 'e', 'E', 'g', 'G'，
// 分别表示小数格式、科学计数法格式、科学计数法格式（大写E）、根据实际情况自动选择格式、根据实际情况自动选择格式（大写G）；参数 prec 表示保留的小数位数；参数 bitSize 表示浮点数类型的位数，取值范围为 32 或 64。
func TestFormatFloat(t *testing.T) {
	t.Log(strconv.FormatFloat(3.1415926, 'f', 2, 64))
}

//
//fmt.Sprintf(format string, a ...interface{}) string：将多个值格式化为字符串，类似于 C 语言中的 sprintf 函数。

func TestSprintf(t *testing.T) {
	t.Log(fmt.Sprintf("%s %s", "hello", "world"))
}

//
//使用 strings.Builder 类型：strings.Builder 是 Golang 提供的一个专门用于字符串拼接的类型。它支持高效的字符串拼接操作，避免了每次创建新的字符串对象，从而提高了性能。可以使用 strings.Builder 的 WriteString 方法将多个字符串拼接起来。

func TestBuilder(t *testing.T) {
	var str strings.Builder
	str.WriteString("字符串A")
	str.WriteString("字符串B")
	t.Log(str.String())
}

//
//使用 bytes.Buffer 类型：bytes.Buffer 是 Golang 提供的一个类似于 strings.Builder 的类型。它也支持高效的字符串拼接操作，但是与 strings.Builder 不同的是，它可以处理二进制数据，而不仅仅是字符串。

func TestBytesBuffer(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("写入字符串")
	// 写入字节数组
	b.Write([]byte{65, 23, 4})
	t.Log(b.String())
	t.Log(b.Bytes())
}

//使用 strings.Join 函数：strings.Join 函数可以将一个字符串切片拼接成一个字符串，使用这个函数可以高效地进行字符串拼接操作。

func TestStringsJoin(t *testing.T) {
	p := []string{"ab", "cd", "ef", "gh"}
	res := strings.Join(p, "+")
	t.Log(res)
	// output: ab+cd+ef+gh
}

//
//以上是一些在 Golang 中高效进行字符串拼接的方法，其中使用 strings.Builder 和 bytes.Buffer 类型可以获得最好的性能。
//

//以下是 Golang 中常用的字符串处理官方包以及它们的主要功能点：
//
//strings 包：提供了许多字符串处理的基本函数，包括字符串拼接、替换、比较、截取、分割等。

//字符串替换

func TestStringsReplace(t *testing.T) {
	str := "中国 你好 你好"
	t.Log(strings.Replace(str, "你好", "哈哈", 1))
	// 替换全部可以指定n为-1
	// 也可以使用下面的方法
	t.Log(strings.ReplaceAll(str, "你好", "哈哈"))
	// tips:替换不会在原字符串中操作，会返回一个新的字符串
}

//字符串的比较，比较字符串的大小

func TestStringsCompare(t *testing.T) {
	str1 := "acd"
	str2 := "bcd"
	com := strings.Compare(str1, str2)
	if com == 0 {
		t.Log("相等")
	} else if com == 1 {
		t.Log("大于")
	} else {
		t.Log("小于")
	}
}

// 字符串的截取、分割
// 常用方法，按指定分隔符分割成切片

func TestSplitString(t *testing.T) {
	str := "a,b,c,d,e,f,g,h,i,j"
	t.Log(strings.Split(str, ","))
	// 相反的方法 join

	// 截取
	t.Log(strings.Cut(str, ","))
}

//strconv 包：提供了字符串和基本数据类型之间的相互转换函数，包括字符串转整型、浮点型、布尔型等，以及整型、浮点型、布尔型等转字符串的函数。
//
//regexp 包：提供了正则表达式的支持，可以用来进行字符串匹配、查找、替换等操作。
//
//unicode 包：提供了 Unicode 字符集的支持，可以用来进行 Unicode 字符集的转换、查找、分类等操作。

func TestUnicode(t *testing.T) {
	// 判断字符是否被试字母
	t.Log(unicode.IsLetter('A'))

	// 转换字符大小写
	t.Log(string(unicode.ToLower('A')))
	t.Log(string(unicode.ToUpper('a')))

	// 转换字符编码
	// 转换为utf8编码
	r := 'A'
	buf := make([]byte, 4)
	utf8.EncodeRune(buf, r)
	t.Log(buf)

	// 计算字节数
	t.Log(len("hello,世界"))

	// 计算字符数
	t.Log(utf8.RuneCountInString("hello,世界"))

	// 检查字符是否可打印
	t.Log(unicode.IsPrint('A'))

}

//
//bytes 包：提供了对字节切片的操作，包括字节切片的拼接、替换、查找、截取等。

func TestBytes(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("写入字符串")
	t.Log(b.String())

	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64}
	var r bytes.Reader
	r.Reset(data)
	var buf [6]byte
	n, err := r.Read(buf[:])
	if err != nil {
		return
	}
	t.Log(n)
	t.Log(buf)

}

//
//bufio 包：提供了对缓冲区的支持，可以用来进行高效的文件读取和写入。

//bufio 包是 Go 语言标准库中的一个包，提供了缓冲读写操作的支持。下面是 bufio 包中常用的一些功能：
//
//bufio.NewReader()：创建一个新的带有默认大小缓冲区的 Reader 对象，用于从输入源中读取数据。

func TestBufioNewReader(t *testing.T) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)

	file, _ := os.OpenFile(currentDir+"/tmp/tmp.log", os.O_RDONLY, 0644)
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(line)
}

//
//bufio.NewWriter()：创建一个新的带有默认大小缓冲区的 Writer 对象，用于向输出源中写入数据。

func TestBufioNewWriter(t *testing.T) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)

	file, _ := os.OpenFile(currentDir+"/tmp/tmp.log", os.O_RDWR, 0644)
	r := bufio.NewWriter(file)
	_, err := r.Write([]byte("addd"))

	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 10)

	// 返回已经缓存的字节数
	t.Log(r.Buffered())

	// 刷新缓存
	r.Flush()

	t.Log(r.Buffered())

	//bufio.NewWriter 创建的缓冲区会在以下情况下自动刷新：
	//
	//当缓冲区已满时，调用写入方法会自动将缓冲区中的内容刷新到底层的 io.Writer 接口中。
	//
	//当调用 Flush 方法时，缓冲区中的内容会被强制刷新到底层的 io.Writer 接口中。
	//
	//当 bufio.NewWriter 关联的 io.Writer 对象被关闭时，缓冲区中的内容会被刷新到底层的 io.Writer 接口中。
	//
	//需要注意的是，在某些情况下，缓冲区中的内容可能并不会被立即刷新到底层的 io.Writer 接口中。例如，当使用文件作为底层的 io.Writer 接口时，由于文件缓冲机制的存在，缓冲区中的内容可能会在一定条件下才被刷新到磁盘中。因此，如果程序需要确保缓冲区中的内容能够及时刷新到底层的 io.Writer 接口中，可以在必要的时候手动调用 Flush 方法。
}

//
//bufio.NewReaderSize() 和 bufio.NewWriterSize()：分别创建具有指定缓冲区大小的 Reader 和 Writer 对象。
//

func TestNewReaderSize(t *testing.T) {
	// 指定缓冲区大小
	_ = bufio.NewReaderSize(os.Stdin, 100)

}

//bufio.Scanner()：创建一个 Scanner 对象，用于逐行读取输入源中的数据。一行一行读取数据
//

// 逐行读取

func TestScanner(t *testing.T) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t.Log(s.Text())
	}
}

//bufio.Reader.Read()：从 Reader 对象中读取字节数据并将其存储到指定的字节数组中。
//
//bufio.Reader.ReadLine()：从 Reader 对象中读取一行文本数据。
//
//bufio.Writer.WriteString()：将指定的字符串写入到 Writer 对象的缓冲区中。
//
//bufio.Writer.Flush()：将 Writer 对象的缓冲区中的数据写入到输出源中。
//
//bufio 包中提供了缓冲读写操作的支持，可以提高读写操作的效率，并且可以更方便地进行数据读写和处理。通过使用 bufio 包，可以避免在读写操作中频繁调用底层系统调用造成的性能损失。
