package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Host          string
	Port          int
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	IdleTimeout   time.Duration
	HeaderTimeout time.Duration
}

func Load() (Config, error){

	port,err := getIntEnv("PORT",8080)
	if err !=nil {
		return Config{} , err
	}

	readTimeout , err := getDurationEnv("READ_TIMEOUT" , 5 * time.Second)
	if err !=nil {
		return Config{} ,err
	}

	 writeTimeout , err := getDurationEnv("WRITE_TIMEOUT", 20*time.Second)
	 if err!=nil{
		return Config{} ,err
	 }

	 idleTimeout, err := getDurationEnv("IDLE_TIMEOUT", 30*time.Second)
	if err != nil {
		return Config{}, err
	}

	headerTimeout, err := getDurationEnv("HEADER_TIMEOUT", 5*time.Second)
	if err != nil {
		return Config{}, err
	}

	return Config{
		Host: getEnv("HOST","localhost"),
		Port: port,
		ReadTimeout: readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout: idleTimeout,
		HeaderTimeout: headerTimeout,
	} ,nil
}


func getEnv(key , defaultValue string) string {
	value :=os.Getenv(key)

	if value ==""{
		return defaultValue	
	}
	return value
}

func getIntEnv(key string, defaultValue int)(int, error){
	value := os.Getenv(key)

	if value ==""{
		return defaultValue,nil
	}

	result,err := strconv.Atoi(value)

	if err!=nil {
		return 0 , fmt.Errorf("invalid %s: %q must be an integer" , key , value)
	}

	return result , nil
}

func getDurationEnv(key string , defaultValue time.Duration)(time.Duration , error){
	value :=os.Getenv(key)

	if value =="" {
		return defaultValue, nil
	}
	result , err := time.ParseDuration(value)

	if err !=nil {
		return 0, fmt.Errorf(
			"invalid %s: %q must be a duration such as 5s or 1m",
			key,
			value,
		)
	}

	return result, nil
}