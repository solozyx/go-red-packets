package infra

import "log"

type ConfStarter struct {
	BaseStarter
}

func init() {
	Registe(&ConfStarter{})
}

func (c *ConfStarter) Init(ctx StarterContext) {
	log.Printf("配置初始化")
}

func (c *ConfStarter) Setup(ctx StarterContext) {
	log.Printf("配置安装")
}

func (c *ConfStarter) Start(ctx StarterContext) {
	log.Printf("配置启动")
}
