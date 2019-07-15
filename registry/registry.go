package registry

var (
	Registries = make(map[string]RegistryCreator)
	//Default    = defaults.NewRegistry()
)

type RegistryCreator func() Registry

// 服务注册与发现:注册，注销，查询，罗列，监听，保活
// 服务名字，Tag，Meta信息
type Registry interface {
	Init(...Option) error
	Options() Options
	Register()
	Deregister()
	Query(name string)
	Catalog() []string
	Watch()
}
