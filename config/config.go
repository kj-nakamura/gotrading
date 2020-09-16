package config

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"time"

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/kelseyhightower/envconfig"
)

type EnvValue struct {
	Environment string `required:"true" split_words:"true" default:"dev"`
	ApiKey      string `required:"true" split_words:"true"`
	ApiSecret   string `split_words:"true"`
	BackTest    bool   `required:"true" split_words:"true"`
	DbName      string `required:"true" split_words:"true"`
	DbHost      string `split_words:"true" default:"mysql"`
	DbUserName  string `required:"true" split_words:"true"`
	DbPassword  string `split_words:"true"`
	IncomingURL string `split_words:"true"`
}

type ConfigValue struct {
	LogFile          string
	ProductCode      string
	TradeDuration    time.Duration
	UsePercent       float64
	DataLimit        int
	StopLimitPercent float64
	NumRanking       int
	Deadline         int
	MaxUseCurrency   float64
	Durations        map[string]time.Duration
	SQLDriver        string
	Port             int
}

// SecretValue secret managerから取得
type SecretValue struct {
	API_SECRET  string
	DB_PASSWORD string
	DB_HOST     string
}

// Env 環境変数から取得
var Env EnvValue

// Config Project内で使う設定
var Config ConfigValue

func init() {
	if err := envconfig.Process("", &Env); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}

	if Env.Environment == "prod" {
		var secretValue SecretValue
		secretName := "prod/trading/"
		region := "ap-northeast-1"

		svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
		input := &secretsmanager.GetSecretValueInput{
			SecretId:     aws.String(secretName),
			VersionStage: aws.String("AWSCURRENT"),
		}

		result, err := svc.GetSecretValue(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case secretsmanager.ErrCodeDecryptionFailure:
					fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
				case secretsmanager.ErrCodeInternalServiceError:
					fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
				case secretsmanager.ErrCodeInvalidParameterException:
					fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
				case secretsmanager.ErrCodeInvalidRequestException:
					fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
				case secretsmanager.ErrCodeResourceNotFoundException:
					fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
			// return err.Error()
		}

		var secretString string

		if result.SecretString != nil {
			secretString = *result.SecretString

			// return secretString
		} else {
			decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
			len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
			if err != nil {
				fmt.Println("Base64 Decode Error:", err)
				// return "Base64 Decode Error"
			}
			secretString = string(decodedBinarySecretBytes[:len])

			// return decodedBinarySecret
		}

		json.Unmarshal([]byte(secretString), &secretValue)
		Env.ApiSecret = secretValue.API_SECRET
		Env.DbPassword = secretValue.DB_PASSWORD
		Env.DbHost = secretValue.DB_HOST
	}

	durations := map[string]time.Duration{
		"1m": time.Minute,
		"1h": time.Hour,
		"1d": 24 * time.Hour,
	}

	Config.Durations = durations
	Config.LogFile = "gotrading.log"
	Config.ProductCode = "BTC_JPY"
	Config.TradeDuration = durations["1h"]
	Config.UsePercent = 0.9
	Config.DataLimit = 365
	Config.StopLimitPercent = 0.8
	Config.NumRanking = 3
	Config.MaxUseCurrency = 100000
	Config.Deadline = 750
	Config.SQLDriver = "mysql"
	Config.Port = 8090
}

func getSecret() string {
	secretName := "prod/trading/"
	region := "ap-northeast-1"

	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return err.Error()
	}

	var secretString, decodedBinarySecret string

	if result.SecretString != nil {
		secretString = *result.SecretString

		return secretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return "Base64 Decode Error"
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])

		return decodedBinarySecret
	}
}
