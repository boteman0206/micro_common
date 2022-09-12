package es

import (
	"bytes"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"sync"
	"time"
)

type DC_ES struct {
	EsClient *elastic.Client
	Buf      *bytes.Buffer
	lock     *sync.Mutex
}

var esClient *elastic.Client
var esIndex = "my_log"

func init() {

	var esUrl = "http://localhost:9200/"
	// 创建Client, 连接ES
	var err error
	esClient, err = elastic.NewClient(
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL(esUrl),
		// 基于http base auth验证机制的账号和密码 不需要账号妈妈
		//elastic.SetBasicAuth("user", "secret"),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置请求失败最大重试次数
		elastic.SetMaxRetries(5),
		// 设置错误日志输出
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//// 设置info日志输出
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))

		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal("连接es失败", err.Error())
	}

	do, i, err := esClient.Ping(esUrl).Do(context.Background())
	if err != nil {
		log.Fatal("连接esPing失败", err.Error())
	}
	fmt.Println(" 获取的es信息： ", do.Name, " i:", i)

	exists, err := esClient.IndexExists(esIndex).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	if !exists {
		result, err := esClient.CreateIndex(esIndex).Do(context.Background())
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("创建了es的index===: ", result.Index)
	}

}

func NewEsClient() *DC_ES {
	return &DC_ES{
		EsClient: esClient,
		Buf:      &bytes.Buffer{},
		lock:     &sync.Mutex{},
	}
}

// 新增
func (c *DC_ES) Add(body interface{}) (bool, error) {
	_, err := c.EsClient.Index().Index(esIndex).Type("_doc").BodyJson(body).Do(context.Background())
	if err != nil {
		fmt.Println("add es error : ", err.Error())
		return false, err
	}
	return true, nil
}

// 新增
func (c *DC_ES) GetBuffer(trace_id, fileName string, line int, v ...interface{}) map[string]interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Buf.Reset()
	logger := c.Buf
	fmt.Fprint(logger, v...)

	m := make(map[string]interface{}, 0)
	m["trace_id"] = trace_id
	m["fileName"] = fileName
	m["line"] = line
	m["message"] = logger.String()

	return m
}
