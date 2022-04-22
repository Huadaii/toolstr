package cronStr

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/robfig/cron"
	uuid "github.com/satori/go.uuid"
)

var CronArray sync.Map

func get16MD5Encode(data string) string {
	return getMD5Encode(data)[8:24]
}

//获取uuid
func getUuid() string {
	u := uuid.NewV4()
	return get16MD5Encode(u.String())
}

//返回一个32位md5加密后的字符串
func getMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func CronStart(cronStr string, cronf func()) string {
	c := cron.New()
	err := c.AddFunc(cronStr, cronf)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Start()
	cronId := getUuid()
	//添加cron定时器
	CronArray.Store(cronId, *c)
	return cronId
}

func CronStop(uuid string) {
	//获取指定cron定时器关闭
	getCron, ok := CronArray.Load(uuid)
	if ok {
		cronNew := getCron.(cron.Cron)
		cronNew.Stop()
	}
	select {}
}
