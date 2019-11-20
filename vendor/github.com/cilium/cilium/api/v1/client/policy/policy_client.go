// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteFqdnCache deletes matching DNS lookups from the policy generation cache

Deletes matching DNS lookups from the cache, optionally restricted by
DNS name. The removed IP data will no longer be used in generated
policies.

*/
func (a *Client) DeleteFqdnCache(params *DeleteFqdnCacheParams) (*DeleteFqdnCacheOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFqdnCacheParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteFqdnCache",
		Method:             "DELETE",
		PathPattern:        "/fqdn/cache",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteFqdnCacheReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFqdnCacheOK), nil

}

/*
DeletePolicy deletes a policy sub tree
*/
func (a *Client) DeletePolicy(params *DeletePolicyParams) (*DeletePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePolicy",
		Method:             "DELETE",
		PathPattern:        "/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePolicyOK), nil

}

/*
GetFqdnCache retrieves the list of DNS lookups intercepted from all endpoints

Retrieves the list of DNS lookups intercepted from endpoints,
optionally filtered by endpoint id, DNS name, or CIDR IP range.

*/
func (a *Client) GetFqdnCache(params *GetFqdnCacheParams) (*GetFqdnCacheOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFqdnCacheParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFqdnCache",
		Method:             "GET",
		PathPattern:        "/fqdn/cache",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFqdnCacheReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFqdnCacheOK), nil

}

/*
GetFqdnCacheID retrieves the list of DNS lookups intercepted from an endpoint

Retrieves the list of DNS lookups intercepted from endpoints,
optionally filtered by endpoint id, DNS name, or CIDR IP range.

*/
func (a *Client) GetFqdnCacheID(params *GetFqdnCacheIDParams) (*GetFqdnCacheIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFqdnCacheIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFqdnCacheID",
		Method:             "GET",
		PathPattern:        "/fqdn/cache/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFqdnCacheIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFqdnCacheIDOK), nil

}

/*
GetFqdnNames lists internal DNS selector representations

Retrieves the list of DNS-related fields (names to poll, selectors and
their corresponding regexes).

*/
func (a *Client) GetFqdnNames(params *GetFqdnNamesParams) (*GetFqdnNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFqdnNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFqdnNames",
		Method:             "GET",
		PathPattern:        "/fqdn/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFqdnNamesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFqdnNamesOK), nil

}

/*
GetIP lists information about known IP addresses

Retrieves a list of IPs with known associated information such as
their identities, host addresses, Kubernetes pod names, etc.
The list can optionally filtered by a CIDR IP range.

*/
func (a *Client) GetIP(params *GetIPParams) (*GetIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIP",
		Method:             "GET",
		PathPattern:        "/ip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIPReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIPOK), nil

}

/*
GetIdentity retrieves a list of identities that have metadata matching the provided parameters

Retrieves a list of identities that have metadata matching the provided parameters, or all identities if no parameters are provided.

*/
func (a *Client) GetIdentity(params *GetIdentityParams) (*GetIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIdentity",
		Method:             "GET",
		PathPattern:        "/identity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIdentityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIdentityOK), nil

}

/*
GetIdentityEndpoints retrieves identities which are being used by local endpoints
*/
func (a *Client) GetIdentityEndpoints(params *GetIdentityEndpointsParams) (*GetIdentityEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityEndpointsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIdentityEndpoints",
		Method:             "GET",
		PathPattern:        "/identity/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIdentityEndpointsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIdentityEndpointsOK), nil

}

/*
GetIdentityID retrieves identity
*/
func (a *Client) GetIdentityID(params *GetIdentityIDParams) (*GetIdentityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIdentityID",
		Method:             "GET",
		PathPattern:        "/identity/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIdentityIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIdentityIDOK), nil

}

/*
GetPolicy retrieves entire policy tree

Returns the entire policy tree with all children.

*/
func (a *Client) GetPolicy(params *GetPolicyParams) (*GetPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPolicy",
		Method:             "GET",
		PathPattern:        "/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPolicyOK), nil

}

/*
GetPolicyResolve resolves policy for an identity context
*/
func (a *Client) GetPolicyResolve(params *GetPolicyResolveParams) (*GetPolicyResolveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyResolveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPolicyResolve",
		Method:             "GET",
		PathPattern:        "/policy/resolve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPolicyResolveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPolicyResolveOK), nil

}

/*
GetPolicySelectors sees what selectors match which identities
*/
func (a *Client) GetPolicySelectors(params *GetPolicySelectorsParams) (*GetPolicySelectorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicySelectorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPolicySelectors",
		Method:             "GET",
		PathPattern:        "/policy/selectors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPolicySelectorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPolicySelectorsOK), nil

}

/*
PutPolicy creates or update a policy sub tree
*/
func (a *Client) PutPolicy(params *PutPolicyParams) (*PutPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutPolicy",
		Method:             "PUT",
		PathPattern:        "/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutPolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutPolicyOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}