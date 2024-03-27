package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"permissions/service"
	"proto/generated/permissions"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// db := utils.GetDB()
	// db.AutoMigrate(&models.Group{}, &models.Permission{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	groupServer := service.GroupServer{}
	permissions.RegisterGroupServiceServer(grpcServer, &groupServer)

	permissionServer := service.PermissionServer{}
	permissions.RegisterPermissionServiceServer(grpcServer, &permissionServer)

	// router := http.NewServeMux()
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello world"))
	// })

	// restServer := http.Server{
	// 	Handler: router,
	// }

	// wg := sync.WaitGroup{}
	// wg.Add(1)

	// go func() {
	// 	log.Println("rest server running on port", os.Getenv("SERVER_PORT"))
	// 	if err := restServer.Serve(lis); err != nil {
	// 		log.Println("failed to serve restServer", err)
	// 		log.Fatal(err)
	// 		wg.Done()
	// 	}
	// }()
	// wg.Wait()

	log.Println("grpc server running on port", os.Getenv("SERVER_PORT"))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Println("Server is running")
}
