package testdata

import (
	"fmt"
	"net"
	"testing"
)

func TestGetNet(t *testing.T) {
	//获取所有网卡信息
	interfaces, err := net.Interfaces()

	if err!=nil{
		fmt.Printf("获取网卡信息出错，错误为%v\n",err)
		return
	}

	for _, inter := range interfaces {
		addrs,err := inter.Addrs()

		if err!=nil{
			fmt.Printf("获取地址信息出错，错误为%v\n",err)
			continue
		}
	
		fmt.Println(inter.Name,addrs)

		//得到所有ip
		for _, addr := range addrs {
			ipNet,ok := addr.(*net.IPNet)
			if !ok{
				continue
			}
			fmt.Println("正确ip地址",ipNet)

			//过滤得到ipv4
			ip4 := ipNet.IP.To4()
			if ip4 == nil{
				continue
			}
			fmt.Println("ipv4地址为",ip4)
		}
	}
}