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
	findPerformance()
}

func findPerformance() {
	data, err := ioutil.ReadFile("../data/online_data.txt")
	if err != nil {
		return
	}

	start := time.Now().UnixNano()
	for i := 0; i < util.Threshold; i++ {
		regexBodyImgWithPTag.FindStringIndex(string(data))
	}
	end := time.Now().UnixNano()
	fmt.Print(util.Format("rust", "findPerformance", end-start))
}

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
