package packetfabric

import (
	"fmt"
	"net/http"
)

const bgpSessionURI = "/v2/bgp-settings/%s/prefixes"
const bgpSessionPrefixesURI = "/v2/bgp-settings/%s/prefixes"
const bgpSessionCloudRouterURI = "/v2/services/cloud-routers/%s/connections/%s/bgp"
const bgpSessionSettingsByUUIDURI = "/v2/services/cloud-routers/%s/connections/%s/bgp/%s"

// This struct represents a Bgp Session for an existing Cloud Router connection
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_bgp_create
type BgpSession struct {
	Md5             string      `json:"md5,omitempty"`
	L3Address       string      `json:"l3_address,omitempty"`
	PrimarySubnet   string      `json:"primary_subnet,omitempty"`
	SecondarySubnet string      `json:"secondary_subnet,omitempty"`
	AddressFamily   string      `json:"address_family"`
	RemoteAddress   string      `json:"remote_address,omitempty"`
	RemoteAsn       int         `json:"remote_asn"`
	MultihopTTL     int         `json:"multihop_ttl,omitempty"`
	LocalPreference int         `json:"local_preference,omitempty"`
	Med             int         `json:"med,omitempty"`
	Community       int         `json:"community,omitempty"`
	AsPrepend       int         `json:"as_prepend,omitempty"`
	Orlonger        bool        `json:"orlonger,omitempty"`
	BfdInterval     int         `json:"bfd_interval,omitempty"`
	BfdMultiplier   int         `json:"bfd_multiplier,omitempty"`
	Disabled        bool        `json:"disabled,omitempty"`
	Prefixes        []BgpPrefix `json:"prefixes,omitempty"`
	Nat             *BgpNat     `json:"nat,omitempty"`
}

type BgpSessionUpdate struct {
	AddressFamily   string      `json:"address_family"`
	BgpSettingsUUID string      `json:"bgp_settings_uuid"`
	Disabled        bool        `json:"disabled"`
	MultihopTTL     int         `json:"multihop_ttl,omitempty"`
	Orlonger        bool        `json:"orlonger,omitempty"`
	RemoteAddress   string      `json:"remote_address,omitempty"`
	RemoteAsn       int         `json:"remote_asn"`
	L3Address       string      `json:"l3_address,omitempty"`
	PrimarySubnet   string      `json:"primary_subnet,omitempty"`
	SecondarySubnet string      `json:"secondary_subnet,omitempty"`
	Prefixes        []BgpPrefix `json:"prefixes"`
	Nat             *BgpNat     `json:"nat,omitempty"`
}

type BgpDnatMapping struct {
	PrivateIP         string `json:"private_ip,omitempty"`
	PublicIP          string `json:"public_ip,omitempty"`
	ConditionalPrefix string `json:"conditional_prefix,omitempty"`
}

type BgpNat struct {
	PreNatSources []interface{}    `json:"pre_nat_sources,omitempty"`
	PoolPrefixes  []interface{}    `json:"pool_prefixes,omitempty"`
	Direction     string           `json:"direction,omitempty"`
	NatType       string           `json:"nat_type,omitempty"`
	DnatMappings  []BgpDnatMapping `json:"dnat_mappings,omitempty"`
}

// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_create
type BgpPrefix struct {
	BgpPrefixUUID   string `json:"bgp_prefix_uuid,omitempty"`
	Prefix          string `json:"prefix,omitempty"`
	MatchType       string `json:"match_type,omitempty"`
	AsPrepend       int    `json:"as_prepend,omitempty"`
	Med             int    `json:"med,omitempty"`
	LocalPreference int    `json:"local_preference,omitempty"`
	Type            string `json:"type,omitempty"`
	Order           int    `json:"order,omitempty"`
}

type BgpSessionCreateResp struct {
	BgpSettingsUUID string      `json:"bgp_settings_uuid"`
	AddressFamily   string      `json:"address_family"`
	RemoteAddress   string      `json:"remote_address"`
	RemoteAsn       int         `json:"remote_asn"`
	MultihopTTL     int         `json:"multihop_ttl"`
	LocalPreference int         `json:"local_preference"`
	Community       string      `json:"community"`
	AsPrepend       int         `json:"as_prepend"`
	Med             int         `json:"med"`
	Md5             string      `json:"md5"`
	Orlonger        bool        `json:"orlonger"`
	BfdInterval     int         `json:"bfd_interval"`
	BfdMultiplier   int         `json:"bfd_multiplier"`
	Disabled        bool        `json:"disabled"`
	Nat             *BgpNat     `json:"nat"`
	Prefixes        []BgpPrefix `json:"prefixes"`
	BgpState        string      `json:"bgp_state"`
	TimeCreated     string      `json:"time_created"`
	TimeUpdated     string      `json:"time_updated"`
}

type BgpSessionBySettingsUUID struct {
	BgpSettingsUUID string      `json:"bgp_settings_uuid"`
	AddressFamily   string      `json:"address_family"`
	RemoteAddress   string      `json:"remote_address"`
	RemoteAsn       int         `json:"remote_asn"`
	MultihopTTL     int         `json:"multihop_ttl"`
	LocalPreference int         `json:"local_preference"`
	Md5             string      `json:"md5"`
	Med             int         `json:"med"`
	L3Address       string      `json:"l3_address,omitempty"`
	PrimarySubnet   string      `json:"primary_subnet,omitempty"`
	SecondarySubnet string      `json:"secondary_subnet,omitempty"`
	Community       interface{} `json:"community"`
	AsPrepend       int         `json:"as_prepend"`
	Orlonger        bool        `json:"orlonger"`
	BfdInterval     int         `json:"bfd_interval"`
	BfdMultiplier   int         `json:"bfd_multiplier"`
	Disabled        bool        `json:"disabled"`
	BgpState        string      `json:"bgp_state"`
	Prefixes        []BgpPrefix `json:"prefixes,omitempty"`
	Subnet          interface{} `json:"subnet"`
	PublicIP        string      `json:"public_ip"`
	Nat             *BgpNat     `json:"nat,omitempty"`
}

// This struct represents a Bgp Session create response
type BgpSessionAssociatedResp struct {
	BgpSettingsUUID string      `json:"bgp_settings_uuid"`
	AddressFamily   string      `json:"address_family"`
	RemoteAddress   string      `json:"remote_address"`
	RemoteAsn       int         `json:"remote_asn"`
	MultihopTTL     int         `json:"multihop_ttl"`
	LocalPreference int         `json:"local_preference"`
	Community       string      `json:"community"`
	AsPrepend       int         `json:"as_prepend"`
	Med             int         `json:"med"`
	Orlonger        bool        `json:"orlonger"`
	BfdInterval     int         `json:"bfd_interval"`
	BfdMultiplier   int         `json:"bfd_multiplier"`
	Disabled        bool        `json:"disabled"`
	TimeCreated     string      `json:"time_created"`
	TimeUpdated     string      `json:"time_updated"`
	Prefixes        []BgpPrefix `json:"prefixes,omitempty"`
	Nat             *BgpNat     `json:"nat,omitempty"`
}

type BgpDeleteMessage struct {
	Message string `json:"message"`
}

// This function represents the Action to Create a Bgp Session using an existing Bgp Settigs UUID
// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_create
func (c *PFClient) CreateBgpSession(bgpSession BgpSession, cID, connID string) (*BgpSessionCreateResp, error) {
	formatedURI := fmt.Sprintf(bgpSessionCloudRouterURI, cID, connID)
	expectedResp := &BgpSessionCreateResp{}
	_, err := c.sendRequest(formatedURI, postMethod, bgpSession, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) CreateBgpSessionPrefixes(prefixes []BgpPrefix, bgpSessionUUID string) ([]BgpPrefix, error) {
	formatedURI := fmt.Sprintf(bgpSessionPrefixesURI, bgpSessionUUID)
	expectedResp := make([]BgpPrefix, 0)
	_, err := c.sendRequest(formatedURI, postMethod, prefixes, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) ReadBgpSessionPrefixes(bgpSettingsUUID string) ([]BgpPrefix, error) {
	formatedURI := fmt.Sprintf(bgpSessionPrefixesURI, bgpSettingsUUID)
	expectedResp := make([]BgpPrefix, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// This function represents the Action to Retrieve a list of Bgp Sessions by Bgp Settings UUID
// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_prefixes_list
func (c *PFClient) ReadBgpSession(bgpSetUUID string) ([]BgpSessionAssociatedResp, error) {
	formatedURI := fmt.Sprintf(bgpSessionURI, bgpSetUUID)
	expectedResp := make([]BgpSessionAssociatedResp, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)

	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetBgpSessionBy(cID, cloudConnID, bgpSettingsUUID string) (*BgpSessionBySettingsUUID, error) {
	formatedURI := fmt.Sprintf(bgpSessionSettingsByUUIDURI, cID, cloudConnID, bgpSettingsUUID)
	expectedResp := &BgpSessionBySettingsUUID{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}

// This function represents the Action to Update a given Cloud Router BGP session
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_bgp_update
func (c *PFClient) UpdateBgpSession(bgpSession BgpSession, cID, connCID string) (*http.Response, *BgpSessionCreateResp, error) {
	formatedURI := fmt.Sprintf(bgpSessionCloudRouterURI, cID, connCID)
	expectedResp := &BgpSessionCreateResp{}
	resp, err := c.sendRequest(formatedURI, putMethod, bgpSession, expectedResp)
	if err != nil {
		return nil, nil, err
	}
	return resp.(*http.Response), expectedResp, err
}

func (c *PFClient) DeleteBgpPrefixes(prefixesUUID []string, bgpSettingsUUID string) ([]BgpPrefix, error) {
	formatedURI := fmt.Sprintf(bgpSessionPrefixesURI, bgpSettingsUUID)
	expectedResp := make([]BgpPrefix, 0)
	_, err := c.sendRequest(formatedURI, deleteMethod, prefixesUUID, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// This function represents the Action to Delete a single BGP Session by a Circuit ID,
// Cloud Connection Circuit ID and BGP Settings UUID
// https://docs.packetfabric.com/api/v2/redoc/#operation/cloud_routers_bgp_delete_by_uuid
func (c *PFClient) DeleteBgpSession(cID, cloudConnCID, bgpSettingsUUID string) (*BgpDeleteMessage, error) {
	formatedURI := fmt.Sprintf(bgpSessionSettingsByUUIDURI, cID, cloudConnCID, bgpSettingsUUID)
	expectedResp := &BgpDeleteMessage{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// This function represents the Action to Return a list of Bgp settings instances associated with the current Account.
// https://docs.packetfabric.com/api/v2/redoc/#operation/bgp_session_settings_list
func (c *PFClient) ListBgpSessions(cID, connCID string) ([]BgpSessionAssociatedResp, error) {
	formatedURI := fmt.Sprintf(bgpSessionCloudRouterURI, cID, connCID)
	expectedResp := make([]BgpSessionAssociatedResp, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
