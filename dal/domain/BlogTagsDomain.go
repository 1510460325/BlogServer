package domain

type BlogTagsDomain struct {
	BaseDomain
	Name *string `json:"name"`
}
