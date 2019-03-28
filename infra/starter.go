package infra

//基础公共资源可能存在依赖 定义启动上下文传递对应的依赖
//别名 传递上下文环境
type StarterContext map[string]interface{}

//基础资源启动器接口
//4个方法管理资源生命周期 1个辅助方法
//每种资源不一定全部实现5个方法
type Starter interface {
	//1.应用启动资源初始化
	// Init(ctx StarterContext)
	Init(StarterContext)
	//2.应用基础资源安装
	Setup(StarterContext)
	//3.应用基础资源启动
	Start(StarterContext)
	//4.基础资源启动有些需要阻塞,有些不需要阻塞
	//web监听端口启动后就要把阻塞起来
	//数据库连接 redis连接 启动完成后不需要阻塞
	//检测启动器是否可阻塞 返回true表示阻塞starter 返回false不阻塞starter
	StartBlocking() bool
	//5.优雅关闭应用系统资源停止和销毁
	Stop(StarterContext)
}

//方便每个基础资源接口的实现 定义该空实现
//具体实现可以嵌套该空实现 覆盖自己需要实现的方法
//基础空启动器实现
type BaseStarter struct{}

func (b *BaseStarter) Init(ctx StarterContext)  {}
func (b *BaseStarter) Setup(ctx StarterContext) {}
func (b *BaseStarter) Start(ctx StarterContext) {}
func (b *BaseStarter) StartBlocking() bool      { return false }
func (b *BaseStarter) Stop(ctx StarterContext)  {}

//使用忽略变量验证 BaseStarter 是否实现了接口
var _ Starter = new(BaseStarter)

//启动器注册器管理所有的Starter
//注册器在系统的整个生命周期只要有1个即可
type starterRegister struct {
	starters []Starter
}

//启动注册器
func (r *starterRegister) Registe(s Starter) {
	r.starters = append(r.starters, s)
}

//获取所有基础资源启动器
func (r *starterRegister) AllStarters() []Starter {
	return r.starters
}

//单例
var StarterRegister *starterRegister = new(starterRegister)

//便捷函数方便外包注册
func Registe(s Starter) {
	StarterRegister.Registe(s)
}

//基础资源生命周期管理
func SystemRun() {
	var ctx StarterContext = StarterContext{}
	//1.初始化所有基础资源
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(ctx)
	}
	//2.安装
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(ctx)
	}
	//3.启动
	for _, starter := range StarterRegister.AllStarters() {
		starter.Start(ctx)
	}
}
