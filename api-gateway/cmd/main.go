package main

import (
	"api-gateway/discovery"
	"api-gateway/internal/service"
	"api-gateway/pkg/util"
	"api-gateway/routes"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	InitConfig()
	go startListen() //转载路由
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
	fmt.Println("gateway listen on :" + viper.GetString("server.port"))
}

func startListen() {
	// etcd注册
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	// 服务名
	userServiceName := viper.GetString("domain.user")
	videoServiceName := viper.GetString("domain.video")
	interactServiceName := viper.GetString("domain.interact")
	relationServiceName := viper.GetString("domain.relation")
	// RPC 连接
	connUser, err := RPCConnect(ctx, userServiceName, etcdRegister)
	if err != nil {
		return
	}
	userService := service.NewUserServiceClient(connUser)

	connVideo, err := RPCConnect(ctx, videoServiceName, etcdRegister)
	if err != nil {
		return
	}
	videoService := service.NewVideoServiceClient(connVideo)

	connInteract, err := RPCConnect(ctx, interactServiceName, etcdRegister)
	if err != nil {
		return
	}
	interactService := service.NewInteractServiceClient(connInteract)

	connRelation, err := RPCConnect(ctx, relationServiceName, etcdRegister)
	if err != nil {
		return
	}
	relationService := service.NewRelationServiceClient(connRelation)

	ginRouter := routes.NewRouter(userService, videoService, interactService, relationService)
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("绑定HTTP到 %s 失败！可能是端口已经被占用，或用户权限不足")
		fmt.Println(err)
	}
	go func() {
		// 优雅关闭
		util.GracefullyShutdown(server)
	}()
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("gateway启动失败, err: ", err)
	}
}

func RPCConnect(ctx context.Context, serviceName string, etcdRegister *discovery.Resolver) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", etcdRegister.Scheme(), serviceName)
	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
