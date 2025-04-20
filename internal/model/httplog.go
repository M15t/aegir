package model

import "gorm.io/datatypes"

// HTTPLog represents a log entry for an HTTP request and response.
type HTTPLog struct {
	Base
	ClientIP        string         `json:"client_ip" gorm:"type:varchar(45)"`      // IP address of the client
	RequestMethod   string         `json:"request_method" gorm:"type:varchar(10)"` // HTTP method (GET, POST, etc.)
	RequestURL      string         `json:"request_url" gorm:"type:text"`           // Requested URL
	RequestHeaders  datatypes.JSON `json:"request_headers" gorm:"type:jsonb"`      // Request headers
	RequestBody     datatypes.JSON `json:"request_body" gorm:"type:jsonb"`         // Body of the request
	ResponseStatus  int            `json:"response_status"`                        // HTTP status code returned
	ResponseHeaders datatypes.JSON `json:"response_headers" gorm:"type:jsonb"`     // Response headers
	ResponseBody    datatypes.JSON `json:"response_body" gorm:"type:jsonb"`        // Body of the response
	Duration        int            `json:"duration"`                               // Time taken to process the request in milliseconds
	ErrorMessage    string         `json:"error_message" gorm:"type:text"`         // Error message if any occurred during processing
}
