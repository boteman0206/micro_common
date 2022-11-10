package consulSer

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/cast"
	"log"
	"micro_product/micro_common/utils"
	"net/http"
	"strings"
)

/**
注册
*/

func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}

// 注册服务
func RegisterServer() {

	config := consulapi.DefaultConfig()

	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	checkPort := 8082

	registration := new(consulapi.AgentServiceRegistration)
	registration.Port = 8802
	registration.ID = "micro_product" + utils.GetIP() + ":" + cast.ToString(registration.Port)
	registration.Name = "micro_product"
	registration.Tags = []string{"micro_product", "v1.0"}
	registration.Address = "127.0.0.1"

	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s", //check失败后30秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

	http.HandleFunc("/check", consulCheck)
	err = http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

}

//获取服务
func GetConsulServices(serName string) *consulapi.AgentService {

	client, err := consulapi.NewClient(consulapi.DefaultConfig()) //非默认情况下需要设置实际的参数
	if err != nil {
		log.Println("NewClient consul err： ", err)
	}
	services, err := client.Agent().Services()
	if err != nil {
		log.Println("Services consul err： ", err)
	}

	for k, v := range services {
		if strings.HasPrefix(k, serName) {
			return v
		}
	}

	return nil
}
