package adapter

// Target 是适配的目标接口
type Target interface {
	Request() string
}

// Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// AdapteeImpl 是被适配的目标类
type adapteeImpl struct {
}

// SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

// Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}

// NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}
