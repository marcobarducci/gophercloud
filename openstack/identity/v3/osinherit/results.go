package osinherit

import "github.com/gophercloud/gophercloud"

// AssignmentResult represents the result of an assign operation.
// Call ExtractErr method to determine if the request succeeded or failed.
type AssignmentResult struct {
	gophercloud.ErrResult
}

// ValidateResult represents the result of an validate operation.
// Call ExtractErr method to determine if the request succeeded or failed.
type ValidateResult struct {
	gophercloud.ErrResult
}
