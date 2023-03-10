package main

import (
	"cos/config"
	"cos/discovery"
	"cos/internal/handler"
	"cos/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 初始化配置
	config.InitConfig()
	handler.InitCos()

	// etcd 地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	defer etcdRegister.Stop()
	userNode := discovery.Server{
		Name:    viper.GetString("server.domain"),
		Addr:    grpcAddress,
		Version: viper.GetString("server.version"),
	}
	server := grpc.NewServer()
	defer server.Stop()
	// 绑定服务
	service.RegisterCosServiceServer(server, handler.NewCosService())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err = etcdRegister.Register(userNode, 10); err != nil {
		panic(err)
	}
	if err = server.Serve(lis); err != nil {
		panic(err)
	}

}
