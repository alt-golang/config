package properties

import (
	"testing"
)

func TestParse(t *testing.T) {
	content := `#comment line
    #coment line
default=default
application=application
environment=default
env-instance=default
local=default
local-development=default
application-profile-1=default
application-profile-2=default
newline=application

step1.step2.step3=application
multiline1=1-\
  2-\
  3\
multiline2=r+\
  g+\
  b\  
l2.multiline3=1-\
  2-\
  3\
multiline4=application\
`
	config := Parse(content)
	t.Log(config)
}
