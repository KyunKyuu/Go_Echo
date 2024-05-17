package web

type AddressServiceRequest struct {
	UserIDFK   int    `json:"user_id_fk"`
	City       string `validate:"required" json:"city"`
	Province   string `validate:"required" json:"province"`
	PostalCode string `validate:"required" json:"postal_code"`
}

type AddressUpdateRequest struct {
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
}
