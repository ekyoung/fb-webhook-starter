package configsrc

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/kelseyhightower/envconfig"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"gopkg.in/yaml.v2"
)

// ErrInvalidAppConfig indicates that an app config is of the wrong type.
var ErrInvalidAppConfig = errors.New("app config must be a struct pointer")

type ConfigSource struct {
	S3Bucket string `envconfig:"s3_bucket"`
	S3Key    string `envconfig:"s3_key"`
}

func Load(prefix string, appConfig interface{}) error {
	t := reflect.ValueOf(appConfig)

	if t.Kind() != reflect.Ptr {
		return ErrInvalidAppConfig
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return ErrInvalidAppConfig
	}

	cfgSource := ConfigSource{}

	envconfigPrefix := fmt.Sprintf("%s_%s", prefix, "config_source")
	err := envconfig.Process(envconfigPrefix, &cfgSource)

	if err != nil {
		return fmt.Errorf("error reading config source from environment: %v", err)
	}

	log.Printf("[CONFIGSRC-debug] Config source: %v\n", cfgSource)

	svc := s3.New(session.New())

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(cfgSource.S3Bucket),
		Key:    aws.String(cfgSource.S3Key),
	})

	if err != nil {
		return fmt.Errorf("error getting app config from S3: %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading S3 GetObject response body: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), appConfig)
	if err != nil {
		return fmt.Errorf("error unmarshalling yaml: %v", err)
	}

	return nil
}
