package ipcity_test

import (
	"fmt"
	"net"
	"testing"

	geoip2db "github.com/cc14514/go-geoip2-db"
)

func TestIpCity(t *testing.T) {
	//根据ip获取城市
	db, _ := geoip2db.NewGeoipDbByStatik()
	defer db.Close()
	//可以使用本机ip
	record, _ := db.City(net.ParseIP("188.253.7.183"))

	//打印具体信息
	fmt.Printf("具体信息为%#v",record)
}