package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"time"

	util "github.com/tanoya/laboratory/reg/go"
)

var (
	regexAllTag          = regexp.MustCompile(util.RegexAllTagStr)
	regexBodyImgWithPTag = regexp.MustCompile(util.RegexBodyImgWithPTagStr)
	regexWhitespace      = regexp.MustCompile(util.RegexWhitespaceStr)
)

func main() {
	replaceAllPerformance()
	findImgPerformance()
	findAllTagPerformance()
	findWhitespacePerformance()
}

// findWhitespacePerformance 查找空格
func findWhitespacePerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}

	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		regexWhitespace.FindStringIndex(string(data))
	}
	end := time.Now().UnixNano()
	fmt.Print(util.Format("sdk", "findWhitespace", end-start))
}

// findAllTagPerformance 查找所有标签
func findAllTagPerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}

	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		regexAllTag.FindStringIndex(string(data))
	}
	end := time.Now().UnixNano()
	fmt.Print(util.Format("sdk", "findAllTagPerformance", end-start))
}

// findImgPerformance 查找图片标签内容
func findImgPerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}

	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		regexBodyImgWithPTag.FindStringIndex(string(data))
	}
	end := time.Now().UnixNano()
	fmt.Print(util.Format("sdk", "findImgPerformance", end-start))
}

// replaceAllPerformance 替换内容
func replaceAllPerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}

	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		tmp := regexBodyImgWithPTag.ReplaceAllString(string(data), "")
		tmp = regexAllTag.ReplaceAllString(tmp, "")
		regexWhitespace.ReplaceAllString(tmp, "")
	}
	end := time.Now().UnixNano()
	// log("replaceAllPerformance", start, end)
	fmt.Print(util.Format("sdk", "replaceAllPerformance", end-start))
}
