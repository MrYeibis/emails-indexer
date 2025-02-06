package zincsearch

type GetAllSearchParams struct {
	SearchType string         `json:"search_type"`
	Query      map[string]any `json:"query"`
	SortFields []string       `json:"sort_fields"`
	Source     []string       `json:"_source"`
	From       uint           `json:"from"`
	MaxResults uint           `json:"max_results"`
}

type GetAllResponse[T any] struct {
	Hits struct {
		Hits []struct {
			Source T `json:"_source"`
		} `json:"hits"`
		MaxScore float64 `json:"max_score"`
		Total    struct {
			Value uint `json:"value"`
		} `json:"total"`
	} `json:"hits"`
}

type AuthErrorResponse struct {
	Auth string `json:"auth"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
