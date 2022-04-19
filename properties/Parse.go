package properties

import (
	"strings"
)

func Parse(content string) map[string]interface{} {
	lines := strings.Split(content, "\n")
	object := map[string]interface{}{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		isCommentLine := strings.HasPrefix(strings.TrimSpace(line), "#")
		isAssignmentLine := strings.Contains(line, "=")
		isMultilineOpening := strings.HasSuffix(strings.TrimSpace(line), "\\") && strings.Contains(line, "=")
		isMultilineTrailing := strings.HasSuffix(strings.TrimSpace(line), "\\") && len(strings.TrimSpace(line)) < len(line)
		if !isCommentLine && !isMultilineTrailing && isAssignmentLine {
			tuple := strings.Split(line, "=")
			key := tuple[0]
			value := tuple[1]
			steps := strings.Split(key, ".")
			keymap := map[string]interface{}{}
			nextmap := keymap
			for j := 0; j < len(steps); j++ {
				if j+1 == len(steps) {
					if isMultilineOpening {
						value = strings.Replace(strings.TrimSpace(value), "\\", "", 1)
						k := 1
						for i+k < len(lines) && strings.HasSuffix(strings.TrimSpace(lines[i+k]), "\\") && !strings.Contains(lines[i+k], "=") {
							value += strings.Replace(strings.TrimSpace(lines[i+k]), "\\", "", 1)
							k += 1
						}
					}
					nextmap[steps[j]] = value
				} else {
					newmap := map[string]interface{}{}
					nextmap[steps[j]] = newmap
					nextmap = newmap
				}
				for k, v := range keymap {
					object[k] = v
				}
			}
		}
	}
	return object
}
