//Replace 
package for_file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

// true - файл существует
func FileExists(fname string) bool {
	f, err := os.Open(fname)
	if err != nil || f == nil {
		return false
	}
	defer f.Close()
	return true
}

// добавляет в конец имени файла (расширения) от 0 до 99 создавая уникальное имя
func GetUniqueFileName(fname string) string {
	n := fname + "_"
	for i := 0; i < 100; i++ {
		if !FileExists(n) {
			return n
		}
		n = fname + fmt.Sprintf("_%d", i)
	}
	return n
}

//модификация оригинальной функции "path.Split"
//добавлена обработка учета вида слеша от типа операционки
//в "dir" возвращает часть "path" до последнего слеша включая его а в "file" после
func PathSplit(path string) (dir, file string) {
	s := "/"
	if runtime.GOOS == "windows" {
		s = "\\"
	}
	i := strings.LastIndex(path, s)
	return path[:i+1], path[i+1:]
}

//добавляет в конец пути слеш если его нет
func TrailingBackSplash(path string) string {
	path = strings.TrimSpace(path)
	if path != "" {
		s := "/"
		if runtime.GOOS == "windows" {
			s = "\\"
		}
		i := strings.LastIndex(path, s)
		if i < len(path)-1 {
			path = path + s
		}
	}
	return path
}

//Заменяет все вхождения "s_old" на "s_new" в файле "f_name"
//"cnt_all" - общее к-во строк в файле, "cnt_replace" - к-во строк в которых произошла замена
func ReplaceSubstrInFile(f_name, s_old, s_new string) (cnt_all, cnt_replace int, err error) {
	if s_old == "" {
		return 0, 0, nil
	}
	fR, err := os.Open(f_name)
	if err != nil {
		return 0, 0, err
	}
	defer fR.Close()
	r := bufio.NewReader(fR)
	n_name := GetUniqueFileName(f_name)
	fW, err := os.Create(n_name)
	if err != nil {
		return 0, 0, err
	}
	defer fW.Close()
	w := bufio.NewWriter(fW)
	cnt_all = 0
	cnt_replace = 0
	for {
		s1, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, errors.New("ReplaceSubstrInFile: " + err.Error())
		}
		ss := s1
		if strings.Contains(s1, s_old) {
			ss = strings.Replace(s1, s_old, s_new, -1)
			cnt_replace++
			//fmt.Println(cnt_replace, " = ", s1)
			//fmt.Println(cnt_replace, " = ", ss)
		}
		w.WriteString(ss)
		cnt_all++
	} //for
	if cnt_replace > 0 {
		fR.Close()
		err := os.Remove(f_name)
		if err != nil {
			return 0, 0, errors.New("ReplaceSubstrInFile: " + err.Error())
		}
		err = w.Flush()
		if err != nil {
			return 0, 0, errors.New("ReplaceSubstrInFile: " + err.Error())
		}
		fW.Close()
		err = os.Rename(n_name, f_name)
		if err != nil {
			return 0, 0, errors.New("ReplaceSubstrInFile: " + err.Error())
		}
	} else {
		fW.Close()
		os.Remove(n_name)
	}
	return cnt_all, cnt_replace, nil
}
