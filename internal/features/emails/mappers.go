package emails

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/mryeibis/indexer/internal/models"
	"github.com/mryeibis/indexer/internal/zincsearch"
)

func filterFromValues(values *url.Values) filter {
	match := getStringQueryParamFromValues(values, "match")
	from := getUintQueryParamFromValues(values, "from")
	maxResults := getUintQueryParamFromValues(values, "maxResults")
	sortFieldsString := getStringQueryParamFromValues(values, "sortFields")

	var sortFields []string
	if sortFieldsString != nil {
		sortFields = strings.Split(*sortFieldsString, ",")
	}

	return filter{
		match:      match,
		from:       from,
		maxResults: maxResults,
		sortFields: &sortFields,
	}
}

func getAllsearchParamsFromFilter(params filter) *zincsearch.GetAllSearchParams {
	var query map[string]any

	searchType := "alldocuments"
	if params.match != nil {
		searchType = "match"
		query = map[string]any{
			"term":  params.match,
			"field": "_all",
		}
	}

	var sortFields []string
	if params.sortFields != nil {
		sortFields = *params.sortFields
	}

	var from uint = 0
	if params.from != nil {
		from = *params.from
	}

	var maxResults uint = 10
	if params.maxResults != nil {
		maxResults = *params.maxResults
	}

	return &zincsearch.GetAllSearchParams{
		SearchType: searchType,
		Query:      query,
		SortFields: sortFields,
		From:       from,
		MaxResults: maxResults,
	}
}

func getStringQueryParamFromValues(values *url.Values, name string) *string {
	value := values.Get(name)
	if value == "" {
		return nil
	}

	return &value
}

func getUintQueryParamFromValues(values *url.Values, name string) *uint {
	value := getStringQueryParamFromValues(values, name)
	if value == nil {
		return nil
	}

	intValue, err := strconv.Atoi(*value)
	if err != nil {
		return nil
	}

	uintValue := uint(intValue)
	return &uintValue
}

func emailFromModel(e models.Email) email {
	return email{
		MessageID:               e.MessageID,
		Date:                    e.Date,
		From:                    e.From,
		To:                      e.To,
		Subject:                 e.Subject,
		Content:                 e.Content,
		ContentTransferEncoding: e.ContentTransferEncoding,
		ContentType:             e.ContentType,
		MimeVersion:             e.MimeVersion,
		XBCC:                    e.XBCC,
		XCC:                     e.XCC,
		XFileName:               e.XFileName,
		XFolder:                 e.XFolder,
		XFrom:                   e.XFrom,
		XOrigin:                 e.XOrigin,
		XTo:                     e.XTo,
	}
}

func getAllResponseFromZincSearchResponse(response *zincsearch.GetAllResponse[models.Email]) paginatedEmailsResponse {
	var records []email
	for _, record := range response.Hits.Hits {
		records = append(records, emailFromModel(record.Source))
	}

	return paginatedEmailsResponse{
		Total:   response.Hits.Total.Value,
		Records: records,
	}
}
