// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// APNList APNs that are allowed for this subscriber
// swagger:model apn_list
type APNList []string

// Validate validates this apn list
func (m APNList) Validate(formats strfmt.Registry) error {
	return nil
}
