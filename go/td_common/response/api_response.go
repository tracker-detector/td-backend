package response

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Self    string      `json:"self,omitempty"`
	Next    string      `json:"next,omitempty"`
	Pages   int         `json:"pages,omitempty"`
}

func NewSuccessResponse(data interface{}) *APIResponse {
	return &APIResponse{
		Success: true,
		Data:    data,
	}
}

func NewErrorResponse(message string) *APIResponse {
	return &APIResponse{
		Success: false,
		Message: message,
	}
}

func NewPagedSuccessResponse(data interface{}, self, next string, pages int) *APIResponse {
	return &APIResponse{
		Success: true,
		Data:    data,
		Self:    self,
		Next:    next,
		Pages:   pages,
	}
}
