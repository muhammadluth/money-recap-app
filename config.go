package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"money-recap-app/model/database"
	"money-recap-app/model/dto"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func getPropertiesService() dto.PropertiesService {
	isLocalDev := flag.Bool("local", false, "=(true/false)")
	flag.Parse()

	fmt.Println("Starting Load Env " + time.Now().Format("2006-01-02 15:04:05"))

	if err := godotenv.Load("config/.env"); err != nil {
		panic(err)
	}

	var appModeDebug bool
	if *isLocalDev {
		appModeDebug = true
	}

	servicePort, err := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	if err != nil {
		panic(err)
	}

	dbPgPort, err := strconv.Atoi(os.Getenv("DB_PG_PORT"))
	if err != nil {
		panic(err)
	}

	propertiesService := dto.PropertiesService{
		AppModeDebug: appModeDebug,
		Service: dto.ServiceConfig{
			ServicePort: servicePort,
		},
		DbPostgres: dto.DbConfig{
			DBHost:     os.Getenv("DB_PG_HOST"),
			DBPort:     dbPgPort,
			DBName:     os.Getenv("DB_PG_NAME"),
			DBUser:     os.Getenv("DB_PG_USER"),
			DBPassword: os.Getenv("DB_PG_PASSWORD"),
		},
		Firebase: dto.FirebaseConfig{
			FirebaseAdminSDKConfigName: os.Getenv("FIREBASE_ADMIN_SDK_CONFIG_NAME"),
		},
		Cors: dto.CorsConfig{
			AllowOrigins: os.Getenv("CORS_ALLOW_ORIGINS"),
		},
	}

	if err := validator.New().Struct(propertiesService); err != nil {
		panic(err)
	}

	fmt.Println("Finish Load Env " + time.Now().Format("2006-01-02 15:04:05"))
	return propertiesService
}

func serviceConfig() fiber.Config {
	config := fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
	}
	return config
}

func dbPostgresConnect(appModeDebug bool, dbConfig dto.DbConfig) *bun.DB {
	dbAddr := fmt.Sprintf("%s:%d", dbConfig.DBHost, dbConfig.DBPort)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&timeout=5s",
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbAddr,
		dbConfig.DBName)
	pgConnectorConn := pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
	)
	sqlOpenDB := sql.OpenDB(pgConnectorConn)
	db := bun.NewDB(sqlOpenDB, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(appModeDebug)))
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func generateConfigFirebaseAdminSDK(dbPostgres *bun.DB, firebaseConfig dto.FirebaseConfig) {
	fmt.Println("Starting Generate Config Firebase Admin SDK " + time.Now().Format("2006-01-02 15:04:05"))

	var configData database.JsonConfig
	if err := dbPostgres.NewSelect().Model(&configData).
		Where("name = ?", firebaseConfig.FirebaseAdminSDKConfigName).
		Scan(context.Background()); err != nil {
		panic(err)
	}

	asJSONByteData, err := json.MarshalIndent(configData.Config, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("config/%s.json", configData.Name),
		asJSONByteData, 0644); err != nil {
		panic(err)
	}
	fmt.Println("Finish Generate Config Firebase Admin SDK " + time.Now().Format("2006-01-02 15:04:05"))
}
