package config

import (
	"absen/models"
	"absen/utils"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func init() {
	InitDB()
}

var (
	redisClient       *redis.Client
	DB                *gorm.DB
	DB_USERNAME       string = utils.GetConfig("DB_USERNAME")
	DB_PASSWORD       string = utils.GetConfig("DB_PASSWORD")
	DB_NAME           string = utils.GetConfig("DB_NAME")
	DB_HOST           string = utils.GetConfig("DB_HOST")
	DB_PORT           string = utils.GetConfig("DB_PORT")
	ACCESS_KEY_ID     string = utils.GetConfig("ACCESS_KEY_ID")
	SECRET_ACCESS_KEY string = utils.GetConfig("SECRET_ACCESS_KEY")
)

// connect to the database
func InitDB() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to the database")
	InitMigrate()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Koneksi Redis gagal:", err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Present{})
}

func CreateS3Client() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ""),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return svc, nil
}

func SeedUser() (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte("testsecret"), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name:     "testname",
		Username: "testusername",
		Role:     "isUser",
		Email:    "test@gmail.com",
		Password: string(password),
	}

	result := DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	if err := result.Last(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CloseDB() error {
	database, err := DB.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}
