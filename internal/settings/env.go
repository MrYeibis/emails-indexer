package settings

import (
	"errors"
	"os"
)

type Env struct {
	ZincSearchURL       string
	ZincSearchUser      string
	ZincSearchPassword  string
	ZincSearchIndexName string
}

func GetEnvVariables() (*Env, error) {
	zincSearchURL, ok := os.LookupEnv("ZINCSEARCH_URL")
	if !ok {
		return nil, errors.New("ZINCSEARCH_URL not found")
	}

	zincSearchUser, ok := os.LookupEnv("ZINCSEARCH_USER")
	if !ok {
		return nil, errors.New("ZINCSEARCH_USER not found")
	}

	zincSearchPassword, ok := os.LookupEnv("ZINCSEARCH_PASSWORD")
	if !ok {
		return nil, errors.New("ZINCSEARCH_PASSWORD not found")
	}

	zincSearchIndexName, ok := os.LookupEnv("ZINCSEARCH_INDEX_NAME")
	if !ok {
		return nil, errors.New("ZINCSEARCH_INDEX_NAME not found")
	}

	return &Env{
		ZincSearchURL:       zincSearchURL,
		ZincSearchUser:      zincSearchUser,
		ZincSearchPassword:  zincSearchPassword,
		ZincSearchIndexName: zincSearchIndexName,
	}, nil
}
