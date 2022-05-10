package main
import (
	"fmt"
	"net"

	mgrpc "golang-restAPI-FR/Core/Grpc"
	log "github.com/Sirupsen/logrus"
	pb "golang-restAPI-FR/Core/Grpc/Services"
	
	"golang-restAPI-FR/Config"
	"golang-restAPI-FR/Database"
	"golang-restAPI-FR/Core/Router"
	"golang-restAPI-FR/Core/Models"
	"google.golang.org/grpc"
)

// Api server start from here. router is define your api router and public it.
func main() {
	// GORM DATABASE
	Database.Mysql, Database.Err = Database.ConnectToDB("main")
	if Database.Err != nil {
		fmt.Println("status error : ", Database.Err)
	} else {
		fmt.Println("database connected")
	}
	defer Database.Mysql.Close()
	// auto migrate
	Database.Mysql.AutoMigrate(&Models.User{})
	Database.Mysql.AutoMigrate(&Models.Friend{})

	// GRPC
	// Here will enable grpc server, if you don`t want it, you can disable it
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 10000))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		pb.RegisterRouteGuideServer(grpcServer, mgrpc.NewServer())
		grpcServer.Serve(lis)
	}()
	app_env := Config.GoDotEnvVariable("APP_ENV")

	Router.Start(app_env)
}
