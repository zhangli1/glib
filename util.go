package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"

	//	"strings"
	"math"
	"net"
	"os"
	"sort"
	"time"
)

//结构体转map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Substr(source interface{}, start_index int, length int) string {
	var s string
	if v, ok := source.(string); ok {
		s = v
	} else {
		s = fmt.Sprintf("%d", source)
	}
	if len([]rune(s)) > length {
		return string([]rune(s)[start_index : start_index+length])
	}
	return string([]rune(s)[start_index : start_index+len([]rune(s))])
}

func TimestampToDate(date_format string, timestamp int) string {
	if date_format == "" {
		date_format = "2006-01-02 15:04:05"
	}
	return time.Unix(int64(timestamp), 0).Format(date_format)
}

func DateToTimestamp(date_format string, date string) int {
	loc, _ := time.LoadLocation("Local")
	if date_format == "" {
		date_format = "2006-01-02 15:04:05"
	}
	theTime, _ := time.ParseInLocation(date_format, date, loc)
	return int(theTime.Unix())

}

//获取变量类型
func GetVariableType(variable interface{}) reflect.Type {
	return reflect.TypeOf(variable)
}

//判断某个key是否存在map或slice中
func IsExistByKey(key string, arr interface{}) bool {
	if _, ok := arr.([]string); ok {
		for _, v := range arr.([]string) {
			if v == key {
				return true
			}
		}
		return false
	} else if _, ok2 := arr.(map[string]map[string]interface{}); ok2 {
		for k, _ := range arr.(map[string]map[string]interface{}) {
			if k == key {
				return true
			}
		}
		return false
	} else if _, ok3 := arr.(map[string]int); ok3 {
		for k, _ := range arr.(map[string]int) {
			if k == key {
				return true
			}
		}
		return false
	} else if _, ok4 := arr.(map[string]map[string]map[string]interface{}); ok4 {
		for k, _ := range arr.(map[string]map[string]map[string]interface{}) {
			if k == key {
				return true
			}
		}
		return false
	} else if _, ok5 := arr.(map[string]map[string]int); ok5 {
		for k, _ := range arr.(map[string]map[string]int) {
			if k == key {
				return true
			}
		}
		return false
	} else if _, ok6 := arr.(map[string][]string); ok6 {
		for k, _ := range arr.(map[string][]string) {
			if k == key {
				return true
			}
		}
		return false
	} else if _, ok7 := arr.(map[string]string); ok7 {
		for k, _ := range arr.(map[string]string) {
			if k == key {
				return true
			}
		}
		return false
	} else {
		for k, _ := range arr.(map[string]interface{}) {
			if k == key {
				return true
			}
		}
		return false
	}
}

//将JSON字符串转成对应的map类型
func JsonToMap(source string) map[string]interface{} {
	var ret_map map[string]interface{}
	err := json.Unmarshal([]byte(source), &ret_map)
	if err != nil {
		fmt.Println("parse json fail", GetCurrentFuncName())
	}
	return ret_map
}

//将Map转JSON字符串
/*func MapToJson(s map[string]interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return ""
	}

	return string(b)
}*/

//将Map转JSON字符串
func MapToJson(s interface{}) string {
	var err, err2, err3, err4, err5, err6, err7 interface{}
	var b, b2, b3, b4, b5, b6 []byte
	if _, ok := s.(map[string]interface{}); ok {
		b, err = json.Marshal(s.(map[string]interface{}))
		if err != nil {
			fmt.Println("json.Marshal failed:", err)
			return ""
		}
		return string(b)
	} else if _, ok2 := s.(map[string]map[string]interface{}); ok2 {
		b2, err2 = json.Marshal(s.(map[string]map[string]interface{}))
		if err2 != nil {
			fmt.Println("json.Marshal list failed:", err2)
			return ""
		}
		return string(b2)
	} else if _, ok3 := s.(map[string]int); ok3 {
		b3, err3 = json.Marshal(s.(map[string]int))
		if err3 != nil {
			fmt.Println("json.Marshal list failed:", err3)
			return ""
		}
		return string(b3)
	} else if _, ok5 := s.(map[string]map[string]map[string]map[string]map[string]map[string]int); ok5 {
		b5, err5 = json.Marshal(s.(map[string]map[string]map[string]map[string]map[string]map[string]int))
		if err5 != nil {
			fmt.Println("json.Marshal list failed:", err5)
			return ""
		}
		return string(b5)
	} else if _, ok6 := s.([]map[string]string); ok6 {
		b6, err6 = json.Marshal(s.([]map[string]string))
		if err6 != nil {
			fmt.Println("json.Marshal list failed:", err6)
			return ""
		}
		return string(b6)
	} else if _, ok7 := s.(map[string]string); ok7 {
		b6, err7 = json.Marshal(s.(map[string]string))
		if err7 != nil {
			fmt.Println("json.Marshal list failed:", err7)
			return ""
		}
		return string(b6)

	} else {
		b4, err4 = json.Marshal(s.([]map[string]interface{}))
		if err4 != nil {
			fmt.Println("json.Marshal list failed:", err4)
			return ""
		}
		return string(b4)
	}
	return ""

}

//获取当前时间
func GetCurrentTime() int {
	return int(time.Now().Unix())
}

//获取当前毫秒数
func GetCurrentMillisTime() int {
	return int(time.Now().UnixNano() / 1e6)
}

//string to float64
func StringToFloat64(str string) float64 {
	var s float64
	var err interface{}
	if s, err = strconv.ParseFloat(str, 64); err != nil {
		fmt.Println(err)
		return 0.00
	}
	return s
}

//float64 to string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

//string to int
func StringToInt(str string) int {
	var i int
	var err interface{}
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println(err, GetCurrentFuncName())
		return 0
	}
	return i
}

//判断是否是字符串数字
func IsNum(str string) bool {
	_, error := strconv.Atoi(str)
	if error != nil {
		return false
	}
	return true
}

//判断是否是字符串
func IsString(str interface{}) string {
	if _, ok := str.(string); ok {
		return "string"
	} else if _, ok2 := str.(float64); ok2 {
		return "float64"
	} else if _, ok3 := str.([]int); ok3 {
		return "[]int"
	} else {
		return "int"
	}
}

//获取当前运行函数名
func GetCurrentFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

//获取内网非回环地址
func GEtInterIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

//通过日期转成周天
func DateToWeek(date_format string, date string) string {
	loc, _ := time.LoadLocation("Local")
	if date_format == "" {
		date_format = "2006-01-02 15:04:05"
	}
	theTime, _ := time.ParseInLocation(date_format, date, loc)
	return theTime.Weekday().String()
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

//排序
func SortMapByValue(m map[string]int) PairList {
	p2 := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p2[i] = Pair{k, v}
		i += 1
	}
	sort.Stable(p2)
	return p2
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//slice取最大值
func SliceMax(list []float64) float64 {
	max := 0.00
	for _, v := range list {
		if v > max {
			max = v
		}
	}

	return max
}

//偏差计算
func GetDeviation(list []float64) float64 {
	length := float64(len(list))
	sum := 0.00
	for _, v := range list {
		sum += v
	}

	avg := sum / length

	total := 0.00
	for _, v2 := range list {
		total += math.Pow((v2 - avg), 2)
	}

	return math.Sqrt(total / length)
}

//获取列表和
func GetListSum(list []float64) float64 {
	total := 0.00
	for _, v := range list {
		total += v
	}
	return total
}

//读文件
func ReadFile(fileName string) string {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()
	//fd, err := ioutil.ReadAll(fi)
	fd := ReadAll2(fi)
	return string(fd)
}

//读取文件替换方法
func ReadAll2(r io.Reader) []byte {
	buffer := bytes.NewBuffer(make([]byte, 0, 65536))
	io.Copy(buffer, r)
	temp := buffer.Bytes()
	length := len(temp)
	var body []byte
	//are we wasting more than 10% space?
	if cap(temp) > (length + length/10) {
		body = make([]byte, length)
		copy(body, temp)
	} else {
		body = temp
	}
	return body
}

//写文件
func WriteFile(fileName string, content string, filetype int) {
    var f *os.File
    var err error
    if filetype < 1 {
	    f, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    } else {
	    f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
    }
	if err != nil {
		fmt.Println(err)
	}
	if _, err := f.Write([]byte(content)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		fmt.Println(err)
	}
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}

}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
			//debug.PrintStack()
			//handler(string(debug.Stack()))
		}
	}()
	fun()
}

//将[]uinit8转换为string
func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

//检查某个端口号是否占用
func CheckPort(host string, port int) bool {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err == nil {
		return true
	}
	return false

}

//获取随机
func MtRand(number int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(number)
}
