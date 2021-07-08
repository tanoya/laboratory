package util

var (
	Threshold               int    = 10000
	RegexAllTagStr          string = `<[^<>]*>`
	RegexBodyImgWithPTagStr string = `<img[^<>]*\ssrc=(?:'|")?([^\s'"<>]+)(?:'|")?[^<>]*>`
	RegexWhitespaceStr      string = `[\s　]*`
)
