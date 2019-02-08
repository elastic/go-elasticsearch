package esapi

// API contains the Elasticsearch APIs
//
type API struct {
	Info Info
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
		Info: newInfoFunc(t),
	}
}
