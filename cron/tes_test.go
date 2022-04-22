package cronStr

import (
	"log"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	cid := CronStart("*/2 * * * * *", func() { log.Println("tme") })
	time.Sleep(10 * time.Second)
	CronStop(cid)
}
