package response

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type page struct {
	Page
	List interface{} `json:"list"`
}
