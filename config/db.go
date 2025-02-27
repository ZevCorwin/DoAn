package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// ConnectDB kết nối tới MongoDB và trả về một đối tượng *mongo.Database
func ConnectDB() *mongo.Database {
	LoadEnv() // Nạp biến môi trường từ file .env

	// Tạo URI kết nối MongoDB
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")

	//if dbHost == "" || dbPort == "" || dbName == "" {
	//	log.Fatal("Lỗi cấu hình: DB_HOST, DB_PORT hoặc DB_NAME không được để trống")
	//}

	// Lấy MONGO_URI từ biến môi trường
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("Lỗi cấu hình: MONGO_URI không được để trống")
	}

	//dbURI := fmt.Sprintf("mongodb://%s:%s/%s", dbHost, dbPort, dbName)
	//fmt.Println("MongoDB URI:", dbURI) // Log URI để kiểm tra

	// Cấu hình client MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(10 * time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Không thể kết nối tới MongoDB: %v", err)
	}

	// Kiểm tra kết nối với MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Ping đến MongoDB thất bại: %v", err)
	}

	// Kết nối thành công
	fmt.Printf("Kết nối thành công đến MongoDB tại URI: %s\n", mongoURI)
	return client.Database("chatapp")
}
