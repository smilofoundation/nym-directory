package presence

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/nymtech/nym-directory/models"
)

// CocoHostSanitizer cleans untrusted input fields
type CocoHostSanitizer interface {
	Sanitize(models.CocoHostInfo) models.CocoHostInfo
}

// MixHostSanitizer cleans untrusted input fields
type MixHostSanitizer interface {
	Sanitize(models.MixHostInfo) models.MixHostInfo
}

// MixProviderHostSanitizer cleans untrusted input fields
type MixProviderHostSanitizer interface {
	Sanitize(models.MixProviderHostInfo) models.MixProviderHostInfo
}

// NewCoconodeSanitizer constructor...
func NewCoconodeSanitizer(p *bluemonday.Policy) CoconodeSanitizer {
	return CoconodeSanitizer{
		policy: p,
	}
}

// CoconodeSanitizer kills untrusted input in CocoHostInfo structs
type CoconodeSanitizer struct {
	policy *bluemonday.Policy
}

// Sanitize CocoHostInfo input
func (s *CoconodeSanitizer) Sanitize(input models.CocoHostInfo) models.CocoHostInfo {
	input.PubKey = s.policy.Sanitize(input.PubKey)
	input.Host = s.policy.Sanitize(input.Host)
	input.Type = s.policy.Sanitize(input.Type)
	return input
}

// NewMixnodeSanitizer constructor...
func NewMixnodeSanitizer(p *bluemonday.Policy) MixnodeSanitizer {
	return MixnodeSanitizer{
		policy: p,
	}
}

// MixnodeSanitizer kills untrusted input in MixHostInfo structs
type MixnodeSanitizer struct {
	policy *bluemonday.Policy
}

// Sanitize MixHostInfo input
func (s *MixnodeSanitizer) Sanitize(input models.MixHostInfo) models.MixHostInfo {
	input.PubKey = s.policy.Sanitize(input.PubKey)
	input.Host = s.policy.Sanitize(input.Host)
	return input
}

// NewMixproviderSanitizer constructor...
func NewMixproviderSanitizer(p *bluemonday.Policy) MixproviderSanitizer {
	return MixproviderSanitizer{
		policy: p,
	}
}

// MixproviderSanitizer kills untrusted input in MixProviderHostInfo structs
type MixproviderSanitizer struct {
	policy *bluemonday.Policy
}

// Sanitize MixProviderHostInfo input
func (s *MixproviderSanitizer) Sanitize(input models.MixProviderHostInfo) models.MixProviderHostInfo {
	input.PubKey = s.policy.Sanitize(input.PubKey)
	input.ClientListener = s.policy.Sanitize(input.ClientListener)
	input.MixnetListener = s.policy.Sanitize(input.MixnetListener)
	for i, client := range input.RegisteredClients {
		input.RegisteredClients[i].PubKey = s.policy.Sanitize(client.PubKey)
	}
	return input
}
