package tlsconfig_test

import (
	"testing"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizeIDPrefix(t *testing.T) {
	authorizer := tlsconfig.AuthorizeIDPrefix(spiffeid.RequireFromString("spiffe://example.org/spire/agent"))

	exact := spiffeid.RequireFromString("spiffe://example.org/spire/agent")
	subpath := spiffeid.RequireFromString("spiffe://example.org/spire/agent/node-a")
	segmentBoundaryMismatch := spiffeid.RequireFromString("spiffe://example.org/spire/agentish")
	otherTrustDomain := spiffeid.RequireFromString("spiffe://other.org/spire/agent")

	assert.NoError(t, authorizer(exact, nil))
	assert.NoError(t, authorizer(subpath, nil))
	assert.EqualError(t, authorizer(segmentBoundaryMismatch, nil),
		`unexpected ID "spiffe://example.org/spire/agentish"`)
	assert.EqualError(t, authorizer(otherTrustDomain, nil),
		`unexpected ID "spiffe://other.org/spire/agent"`)
}

func TestAuthorizeIDPrefix_WithoutPath(t *testing.T) {
	authorizer := tlsconfig.AuthorizeIDPrefix(spiffeid.RequireFromString("spiffe://example.org"))

	withoutPath := spiffeid.RequireFromString("spiffe://example.org")
	withPath := spiffeid.RequireFromString("spiffe://example.org/spire/agent")
	otherTrustDomain := spiffeid.RequireFromString("spiffe://other.org/spire/agent")

	assert.NoError(t, authorizer(withoutPath, nil))
	assert.NoError(t, authorizer(withPath, nil))
	assert.EqualError(t, authorizer(otherTrustDomain, nil),
		`unexpected ID "spiffe://other.org/spire/agent"`)
}
