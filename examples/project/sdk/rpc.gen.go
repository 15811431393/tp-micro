// Code generated by 'micro gen' command.
// DO NOT EDIT!

package sdk

import (
	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/socket"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/discovery"
	"github.com/xiaoenai/tp-micro/model/etcd"

	"github.com/xiaoenai/tp-micro/examples/project/args"
)

var client *micro.Client

// Init initializes client with configs.
func Init(cliConfig micro.CliConfig, etcdConfing etcd.EasyConfig) {
	client = micro.NewClient(
		cliConfig,
		discovery.NewLinker(etcdConfing),
	)
}

// InitWithClient initializes client with specified object.
func InitWithClient(cli *micro.Client) {
	client = cli
}

// Home comment...
func Home(arg *struct{}, setting ...socket.PacketSetting) (*args.HomeResult, *tp.Rerror) {
	result := new(args.HomeResult)
	rerr := client.Pull("/project/home", arg, result, setting...).Rerror()
	return result, rerr
}

// Math_Divide handler
func Math_Divide(arg *args.DivideArg, setting ...socket.PacketSetting) (*args.DivideResult, *tp.Rerror) {
	result := new(args.DivideResult)
	rerr := client.Pull("/project/math/divide", arg, result, setting...).Rerror()
	return result, rerr
}

// Stat comment...
func Stat(arg *args.StatArg, setting ...socket.PacketSetting) *tp.Rerror {
	return client.Push("/project/stat", arg, setting...)
}
