/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type SmContextCreateError struct {
	Error        *ProblemDetails  `json:"error"`
	N1SmMsg      *RefToBinaryData `json:"n1SmMsg,omitempty"`
	RecoveryTime *time.Time       `json:"recoveryTime,omitempty"`
}