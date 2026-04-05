# tao-hello

```
  _   _  U _____ u  _       _       U  ___ u      _____      _      U  ___ u 
 |'| |'| \| ___"|/ |"|     |"|       \/"_ \/     |_ " _| U  /"\  u   \/"_ \/ 
/| |_| |\ |  _|" U | | u U | | u     | | | |       | |    \/ _ \/    | | | | 
U|  _  |u | |___  \| |/__ \| |/__.-,_| |_| |      /| |\   / ___ \.-,_| |_| | 
 |_| |_|  |_____|  |_____| |_____|\_)-\___/      u |_|U  /_/   \_\\_)-\___/  
 //   \\  <<   >>  //  \\  //  \\      \\        _// \\_  \\    >>     \\    
(_") ("_)(__) (__)(_")("_)(_")("_)    (__)      (__) (__)(__)  (__)   (__)  
```

Tao Universe 示例组件，演示如何创建自定义 Tao 单元。适合作为新组件开发的参考模板。

## 安装

```bash
go get github.com/taouniverse/tao-hello
```

## 使用

### 导入

```go
import _ "github.com/taouniverse/tao-hello"
```

### 配置

单实例模式：

```yaml
hello:
  print: "Hello, Tao!"
  times: 3
  run_after: []
```

多实例模式：

```yaml
hello:
  default_instance: greeting
  greeting:
    print: "Hello, Tao!"
    times: 3
  farewell:
    print: "Goodbye, Tao!"
    times: 1
```

## 工厂模式

| API | 说明 |
|-----|------|
| `hello.Factory` | `*tao.BaseFactory[struct{}]` 工厂实例 |

## 作为模板

`tao-hello` 是创建新 Tao 组件的最佳参考。核心结构如下：

```
tao-hello/
├── config.go          # Config 接口实现 + InstanceConfig 定义
├── hello.go           # init() 注册 + 构造器函数
├── config_test.go     # 配置解析测试
└── hello_test.go      # 集成测试
```

关键模式：
1. **两层 Config**：`InstanceConfig` 存储具体字段，`Config` 嵌入 `BaseMultiConfig[InstanceConfig]`
2. **构造器**：`NewHello(name string, cfg InstanceConfig) (instance, closer, error)` — 返回三元组
3. **init() 注册**：通过 `tao.Register(ConfigKey, &Config{}, NewHello)` 将组件注册到 Tao Universe
4. **便捷函数（可选）**：提供 `Xxx()` 和 `GetXxx(name)` 快捷访问默认/指定实例

### 最小化模板代码

```go
// config.go
type InstanceConfig struct {
    // 你的配置字段
}

type Config struct {
    tao.BaseMultiConfig[InstanceConfig]
    RunAfters []string `json:"run_after,omitempty"`
}

func (c *Config) Name() string        { return ConfigKey }
func (c *Config) ValidSelf()          { /* 设置默认值 */ }
func (c *Config) ToTask() tao.Task     { /* 任务逻辑 */ }
func (c *Config) RunAfter() []string   { return c.RunAfters }
```

```go
// xxx.go
var M = &Config{}
var Factory *tao.BaseFactory[*YourType]

func init() {
    var err error
    Factory, err = tao.Register(ConfigKey, M, NewXxx)
    if err != nil {
        panic(err.Error())
    }
}

func NewXxx(name string, cfg InstanceConfig) (*YourType, func() error, error) {
    // 创建实例
    instance := createInstance(cfg)
    
    closer := func() error {
        // 清理资源
        return nil
    }
    
    return instance, closer, nil
}

// 可选：便捷函数
func Xxx() (*YourType, error) {
    return Factory.Get(M.GetDefaultInstanceName())
}

func GetXxx(name string) (*YourType, error) {
    return Factory.Get(name)
}
```

## 单元测试

| 测试文件 | 说明 | 运行条件 |
|---------|------|---------|
| `config_test.go` | Config 解析与默认值验证 | 无需外部依赖 |
| `hello_test.go` | 完整集成测试 | 无需外部服务 |

### 运行单元测试

```bash
# 全部测试可直接运行
go test -v ./...
```
