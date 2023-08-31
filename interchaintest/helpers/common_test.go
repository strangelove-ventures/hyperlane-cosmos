package helpers_test

import (
	"fmt"
	"regexp"
	"testing"
)

// testing the regex
func TestParseCliOutput(t *testing.T) {
	// STDOUT will include the hash in its own line in the format "txhash: 6A9F49069B8D3E8B4CA92F76C695263C2F0E8C59B1BDD9A65E5A7C699A9F32DA"
	txt := `useless info we dont want,
txhash: 6A9F49069B8D3E8B4CA92F76C695263C2F0E8C59B1BDD9A65E5A7C699A9F32DA
other line something
last line yeah`

	r, _ := regexp.Compile("(?m)^txhash:\\s(?P<hash>.*)$")
	matches := r.FindStringSubmatch(txt)
	hashIndex := r.SubexpIndex("hash")
	fmt.Print(matches[hashIndex])
}
