package env

import (
	"os"

	"github.com/c4pt0r/cfg"
	log "github.com/ngaut/logging"
	"github.com/ngaut/zkhelper"
)

type Env interface {
	ProductName() string
	DashboardAddr() string
  ZkPathPrefix() string
	NewZkConn() (zkhelper.Conn, error)
}

type CodisEnv struct {
	zkAddr        string
	dashboardAddr string
	productName   string
  zkPathPrefix  string
}

func LoadCodisEnv(cfg *cfg.Cfg) Env {
	if cfg == nil {
		log.Fatal("config error")
	}

	productName, err := cfg.ReadString("product", "test")
	if err != nil {
		log.Fatal(err)
	}

	zkAddr, err := cfg.ReadString("zk", "localhost:2181")
	if err != nil {
		log.Fatal(err)
	}

	zkPathPrefix, err := cfg.ReadString("zk_path_prefix", "/zk/codis/db_")
	if err != nil {
		log.Fatal(err)
	}

	hostname, _ := os.Hostname()
	dashboardAddr, err := cfg.ReadString("dashboard_addr", hostname+":18087")
	if err != nil {
		log.Fatal(err)
	}

	return &CodisEnv{
		zkAddr:        zkAddr,
		dashboardAddr: dashboardAddr,
		productName:   productName,
    zkPathPrefix:  zkPathPrefix,
	}
}

func (e *CodisEnv) ProductName() string {
	return e.productName
}

func (e *CodisEnv) DashboardAddr() string {
	return e.dashboardAddr
}

func (e *CodisEnv) ZkPathPrefix() string {
  return e.zkPathPrefix
}

func (e *CodisEnv) NewZkConn() (zkhelper.Conn, error) {
	return zkhelper.ConnectToZk(e.zkAddr)
}

