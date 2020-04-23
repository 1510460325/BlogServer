package info

type PagingInfo struct {
	List  interface{} `json:"data"`
	Total uint32      `json:"total"`
}
