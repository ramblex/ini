// Package ini is a basic ini file reader that supports sections and attributes
// which have spaces in them. It currently doesn't support nested sections and
// ignores indentation.
package ini

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// AttributeMap represents the attributes for a section
type AttributeMap map[string]string

// SectionMap represents the whole ini file
type SectionMap map[string]AttributeMap

const SECTION_REGEXP = `^\[(.*)\]$`
const ATTRIBUTE_REGEXP = `^([\w ]+)\s*=\s*([^#]+)`

var section = regexp.MustCompile(SECTION_REGEXP)
var attribute = regexp.MustCompile(ATTRIBUTE_REGEXP)

// ReadIni reads a configuration file into an map of the form:
//    {
//			"section name": {
//	  		"attribute name": "value"
//	  	},
//	  	"section 2 name": { ... },
//	  }
func ReadIni(path string) (SectionMap, error) {
	attributes := make(SectionMap)

	file, err := os.Open(path)
	if err != nil {
		return attributes, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	var currentSection string
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case section.MatchString(line):
			currentSection = section.FindStringSubmatch(line)[1]
			attributes[currentSection] = make(AttributeMap)
		case attribute.MatchString(line):
			matches := attribute.FindStringSubmatch(line)
			attributes[currentSection][strings.TrimSpace(matches[1])] = strings.TrimSpace(matches[2])
		}
	}

	return attributes, nil
}

// String outputs the ini file in a form which is more easily parsable by
// something like grep or awk. Each line will be of the form:
//
//     SECTION NAME__ATTRIBUTE=VALUE
// For example:
// Section Name__Attribute Name=3
func (s SectionMap) String() (output string) {
	for section := range s {
		for attr, value := range s[section] {
			output += fmt.Sprintf("%s__%s=%s\n", section, attr, value)
		}
	}
	return
}
