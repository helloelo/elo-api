package model

// Organization is the model for organization
type Organization struct {
	IDOrganization uint64 `db:"id_organization" json:"id_organisation"`
	Name           string `db:"name" json:"organization_name"`
	Domain         string `db:"domain" json:"organization_domain"`
}
