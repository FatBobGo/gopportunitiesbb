package handler

import "fmt"

func errParamIsRequired(name string, typ string) error {
		return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if !r.Remote {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return fmt.Errorf("param: salary must be greater than zero")
	}
	return nil	
}

