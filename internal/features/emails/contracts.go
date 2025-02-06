package emails

type filter struct {
	match      *string
	sortFields *[]string
	from       *uint
	maxResults *uint
}

type errorResponse struct {
	Error string `json:"error"`
}

type paginatedEmailsResponse struct {
	Total   uint    `json:"total"`
	Records []email `json:"records"`
}

type email struct {
	MessageID               string `json:"messageId"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Subject                 string `json:"subject"`
	Content                 string `json:"content"`
	ContentTransferEncoding string `json:"contentTransferEncoding"`
	ContentType             string `json:"contentType"`
	MimeVersion             string `json:"mimeVersion"`
	XBCC                    string `json:"xBcc"`
	XCC                     string `json:"xCc"`
	XFileName               string `json:"xFilename"`
	XFolder                 string `json:"xFolder"`
	XFrom                   string `json:"xFrom"`
	XOrigin                 string `json:"xOrigin"`
	XTo                     string `json:"xTo"`
}
