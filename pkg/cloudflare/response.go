package cloudflare

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseResultInfo struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Count      int `json:"count"`
	TotalCount int `json:"total_count"`
}

type Response struct {
	Success    bool               `json:"success"`
	Errors     []ResponseError    `json:"errors"`
	Messages   []string           `json:"messages"`
	ResultInfo ResponseResultInfo `json:"result_info"`
}
