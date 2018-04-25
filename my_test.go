package main

import (
	"bytes"
	"log"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestConvertGraphqlArgsToJson(t *testing.T) {
	var now = time.Now()
	var str = `a:"a6OGtW2bxPeMjPQl",b:[{c:"u1",action:1},{d:"u2",action:1}],devices:["ee","ff","gg"]`

	re := regexp.MustCompile(`([,[{]*)([^(^{,:)]*)([:])(\[[^{]*\]|[^([{,:)]*)`)
	s := re.FindAllStringSubmatch(str, -1)

	var buffer bytes.Buffer
	buffer.WriteString("{")
	for _, a := range s {
		buffer.WriteString(a[1] + `"` + a[2] + `"` + a[3] + a[4])
	}
	buffer.WriteString("}")

	log.Printf("buffer = %v\n", buffer.String())

	var after = time.Since(now)
	log.Printf("after = %s\n", after)
}
