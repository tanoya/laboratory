package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/BurntSushi/rure-go"
	util "github.com/tanoya/laboratory/reg/go"
)

var (
	regexAllTag          = rure.MustCompile(util.RegexAllTagStr)
	regexBodyImgWithPTag = rure.MustCompile(util.RegexBodyImgWithPTagStr)
	regexWhitespace      = rure.MustCompile(util.RegexWhitespaceStr)
)

func log(method string, start, end int64) error {
	context := util.Format("rust", method, end-start)
	return util.WriteAppend("rust", []byte(context))
}

func main() {
	replaceAllPerformance()
	findImgPerformance()
	findAllTagPerformance()
	findWhiteSpacePerformance()
}

// findWhiteSpacePerformance 查找空白字符
func findWhiteSpacePerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}

	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		regexWhitespace.FindAllBytes([]byte(data))
	}
	end := time.Now().UnixNano()
	fmt.Print(util.Format("rust", "findWhiteSpacePerformance", end-start))
}

// findAllTagPerformance 查找所有标签
func findAllTagPerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}

	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		regexAllTag.FindAllBytes([]byte(data))

	}
	end := time.Now().UnixNano()
	fmt.Print(util.Format("rust", "findAllTagPerformance", end-start))
}

// findImgPerformance 查找图片标签
func findImgPerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}

	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		regexBodyImgWithPTag.FindAllBytes([]byte(data))
	}
	end := time.Now().UnixNano()
	fmt.Print(util.Format("rust", "findImgPerformance", end-start))
}

// replaceAllPerformance 线上数据[替换]的性能表现
func replaceAllPerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}
	var buf = bytes.NewBuffer(make([]byte, 0, len(data)))

	var ret []byte
	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		index := regexBodyImgWithPTag.FindAllBytes([]byte(data))
		ret = deleteMatch([]byte(data), index, buf)
		index = regexAllTag.FindAllBytes(ret)
		ret = deleteMatch(ret, index, buf)
		index = regexWhitespace.FindAllBytes(ret)
		ret = deleteMatch(ret, index, buf)
	}
	end := time.Now().UnixNano()
	// log("replaceAllPerformance", start, end)
	fmt.Print(util.Format("rust", "replaceAllPerformance", end-start))
}

// deleteMatch 删除匹配的数据（replace all）
func deleteMatch(data []byte, idx []int, buf *bytes.Buffer) []byte {
	if len(idx) < 1 {
		return data
	}
	buf.Reset()
	buf.Write(data[0:idx[0]])
	for i := 1; i < len(idx); i += 2 {
		if i+1 >= len(idx) {
			start := idx[i]
			buf.Write(data[start:])
		} else {
			start := idx[i]
			end := idx[i+1]
			buf.Write(data[start:end])
		}
	}
	data = buf.Bytes()
	return data
}
