package emailStr

import (
	"log"
	"testing"
)

func TestCron(t *testing.T) {
	err := SendMail("wuhuuuuuuuu", "tesss", "ss@163.com", "testheader", "testsecondheader", []string{"s@qq.com", "s@qq.com"})
	if err != nil {
		log.Println(err)
	}
}
