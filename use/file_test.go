package use

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// 判断当前文件是否存在

func TestExists(t *testing.T) {
	_, err := os.Stat("file_test.go")
	if os.IsNotExist(err) {
		t.Log("文件不存在")
	}
}

// 获取当前路径

func TestPath(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	t.Log(path)
}

// 获可执行文件目录

func TestRootPath(t *testing.T) {
	executable, err := os.Executable()
	if err != nil {
		t.Error(err)
	}
	exeDir := filepath.Dir(executable)
	t.Log(exeDir)
}

// 获取项目目录

func TestProjectPath(t *testing.T) {
	_, filePath, _, _ := runtime.Caller(0)
	t.Log(filePath)
	// 获取当前目录，排除最后的文件
	fileDir := filepath.Dir(filePath)
	t.Log(fileDir)
}

// 获取文件权限

func TestPower(t *testing.T) {
	fileInfo, err := os.Stat("file_test.go")
	if err != nil {
		return
	}
	permissions := fileInfo.Mode().Perm()
	t.Logf("%o", permissions)
}

// 获取文件大小

func TestSize(t *testing.T) {
	fileInfo, err := os.Stat("file_test.go")
	if err != nil {
		return
	}
	size := fileInfo.Size()
	t.Logf("%d", size)
}

// 获取文件的扩展名

func TestExt(t *testing.T) {
	ext := filepath.Ext("file_test.go")
	t.Log(ext)
}

// 创建目录

func TestCreatePath(t *testing.T) {
	// 获取当前项目的目录
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)
	// 在当前目录下继续创建tmp目录
	err := os.MkdirAll(currentDir+"/tmp", 0755)
	if err != nil {
		t.Error(err)
	}
}

// 在上述创建的目录下创建一个名为tmp.log的文件
// 1、判断文件tmp.log是否存在
// 2、tmp.log不存在时进行创建

func TestCreateFile(t *testing.T) {
	// 获取当前项目的目录
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)
	_, err := os.Stat(currentDir + "/tmp/tmp.log")
	if os.IsNotExist(err) {
		// 创建文件
		file, err := os.Create(currentDir + "/tmp/tmp.log")
		defer file.Close()
		if err != nil {
			return
		}
		t.Log(currentDir + "/tmp/tmp.log 创建成功")
	}
	t.Log("文件已经存在。。。")
}

// 向上面创建的文件中添加 hello golang 的内容

func TestWriteFile(t *testing.T) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)

	file, err := os.OpenFile(currentDir+"/tmp/tmp.log", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("hello golang")
	if err != nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
	t.Log("写入成功。。。")
}

// 上面在文件中写入数据采用了覆盖的模式，意思是新写入的内容将覆盖文件原有的内容
// 下面的任务是，采用追加的方式，写入 all

func TestAppendWrite(t *testing.T) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)

	file, err := os.OpenFile(currentDir+"/tmp/tmp.log", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(" all")
	if err != nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
	t.Log("追加成功。。。")
}

// 读取上面文件的全部内容

func TestReadAll(t *testing.T) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)
	file, err := os.Open(currentDir + "/tmp/tmp.log")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	all, err := io.ReadAll(reader)
	if err != nil {
		return
	}
	t.Log(string(all))
}

// 上面文件内容很少，每次读取全部内容没什么问题
// 但是如果读取一个内容很多的文件的全部内容的话，会有性能问题
// 请按行读取文件内容

func TestReadFileByLine(t *testing.T) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)
	file, _ := os.Open(currentDir + "/tmp/tmp.log")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 读取一行
		t.Log(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Error("读取文件失败", err)
	}

}

// 删除上面的文件

// os.Remove 只能删除空目录，或者指定的文件，否则会报错
// os.RemoveAll 删除非空目录不会报错

func TestDeleteFile(t *testing.T) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filePath)
	err := os.RemoveAll(currentDir + "/tmp/tmp.log")
	if err != nil {
		return
	}
	t.Log("删除成功")
}
