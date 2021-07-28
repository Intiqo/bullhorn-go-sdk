package bullhorn

type CreateResponse struct {
	ChangedEntityId int    `json:"changedEntityId"`
	ChangeType      string `json:"changeType"`
}

type UpdateResponse struct {
	ChangedEntityId int    `json:"changedEntityId"`
	ChangeType      string `json:"changeType"`
}
