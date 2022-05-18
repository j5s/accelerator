package rule

import (
	"io/ioutil"
	"strings"
)

func GetRule(ruleFile string) [][]string {
	var rule [][]string
	data, _ := ioutil.ReadFile(ruleFile)
	lines := strings.Split(string(data)+"\n", "\n")
	for _, line := range lines {
		line = strings.Trim(line, "\r")
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, " ") {
			inst := strings.Split(line, " ")[0]
			var (
				methodName string
				methodDesc string
			)
			if strings.HasPrefix(strings.ToUpper(inst), "INVOKE") {
				methodName = strings.Split(line, " ")[1]
				methodDesc = strings.Split(line, " ")[2]
			} else {
				panic("not support yet")
			}
			instData := make([]string, 3)
			instData[0] = strings.ToUpper(inst)
			instData[1] = methodName
			instData[2] = methodDesc
			rule = append(rule, instData)
		} else {
			panic("parse error")
		}
	}
	return rule
}
