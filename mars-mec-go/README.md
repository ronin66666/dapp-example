##1、服务的安装
步骤如下:

服务的实现
定义一个组合了service.Service的结构体,它就是一个服务了

服务的预加载
在服务实现的package的init函数中使用 node.setup就可以预加载服务了, 预加载之前根据需要设置服务的Name, 与后面配置servicelist中的名字对应,默认是服务结构体的名称.

服务配置
在config/cluster.json 中可以配置每个节点的servicelist, 添加节点想要装载的服务名即可.

main.go中导入服务package
这点是很容易忽略的点,一不小心就容易搞忘记了, 一定要在main.go中import

##2、rpc的使用细节
// 异步, rpc首选, 不会阻塞本服务
AsyncCall(serviceMethod string,args interface{},callback interface{}) error
// 同步等待调用结果
Call(serviceMethod string,args interface{},reply interface{}) error
// 无结果,不阻塞
Go(serviceMethod string,args interface{}) error
// 在明确节点时调用,可以稍微减少开销;  Service名相同时, 避免广播
AsyncCallNode(nodeId int,serviceMethod string,args interface{},callback interface{}) error
CallNode(nodeId int,serviceMethod string,args interface{},reply interface{}) error
GoNode(nodeId int,serviceMethod string,args interface{}) error
// 原数据,减少参数/结果的序列化和反序列化, 大量转发时使用.
RawGoNode(rpcProcessorType RpcProcessorType,nodeId int,rpcMethodId uint32,serviceName string,rawArgs IRawInputArgs) error
// 广播
CastGo(serviceMethod string,args interface{})

RPC函数的定义方式看过官方文档了就大致清楚了,这里强调一下, 有时候被调用方可能是一个很费时的操作,这个时候就需要考虑如下这种方式的RPC函数实现:

## 3. 项目启动
1. 配置节点
   program arguments : --start nodeid=1
2. 配置工作区
   working directory: F:\work\go\mars-mec-go\examples
3. 项目启动
   /examples/server.go