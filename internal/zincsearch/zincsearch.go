package zincsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ZincSearch[T any] struct {
	url       string
	user      string
	password  string
	indexName string
}

func New[T any](url, user, password, indexName string) *ZincSearch[T] {
	return &ZincSearch[T]{
		url:       url,
		user:      user,
		password:  password,
		indexName: indexName,
	}
}

func (z *ZincSearch[T]) UploadBulkV2(records []T) error {
	data := map[string]any{
		"index":   z.indexName,
		"records": records,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := z.url + "/api/_bulkv2"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(z.user, z.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			var authErrorResponse AuthErrorResponse
			json.NewDecoder(resp.Body).Decode(&authErrorResponse)
			return errors.New(authErrorResponse.Auth)
		}

		var errorResponse ErrorResponse
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		return errors.New(errorResponse.Error)
	}

	return nil
}

func (z *ZincSearch[T]) GetAll(params GetAllSearchParams) (*GetAllResponse[T], error) {
	query, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := z.url + fmt.Sprintf("/api/%s/_search", z.indexName)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(z.user, z.password)
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result GetAllResponse[T]
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
