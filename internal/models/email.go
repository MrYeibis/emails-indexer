package models

type Email struct {
	MessageID               string `json:"messageId"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Subject                 string `json:"subject"`
	MimeVersion             string `json:"mimeVersion"`
	ContentType             string `json:"contentType"`
	ContentTransferEncoding string `json:"contentTransferEncoding"`
	XFrom                   string `json:"x-from"`
	XTo                     string `json:"x-to"`
	XCC                     string `json:"x-cc"`
	XBCC                    string `json:"x-bcc"`
	XFolder                 string `json:"x-folder"`
	XOrigin                 string `json:"x-origin"`
	XFileName               string `json:"x-filename"`
	Content                 string `json:"content"`
}
