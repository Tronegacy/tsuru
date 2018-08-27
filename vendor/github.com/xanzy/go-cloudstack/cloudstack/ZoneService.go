//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type AddVmwareDcParams struct {
	p map[string]interface{}
}

func (p *AddVmwareDcParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["vcenter"]; found {
		u.Set("vcenter", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddVmwareDcParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *AddVmwareDcParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddVmwareDcParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

func (p *AddVmwareDcParams) SetVcenter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vcenter"] = v
	return
}

func (p *AddVmwareDcParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddVmwareDcParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewAddVmwareDcParams(name string, vcenter string, zoneid string) *AddVmwareDcParams {
	p := &AddVmwareDcParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["vcenter"] = vcenter
	p.p["zoneid"] = zoneid
	return p
}

// Adds a VMware datacenter to specified zone
func (s *ZoneService) AddVmwareDc(p *AddVmwareDcParams) (*AddVmwareDcResponse, error) {
	resp, err := s.cs.newRequest("addVmwareDc", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddVmwareDcResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddVmwareDcResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Vcenter string `json:"vcenter"`
	Zoneid  int64  `json:"zoneid"`
}

type CreateZoneParams struct {
	p map[string]interface{}
}

func (p *CreateZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["dns1"]; found {
		u.Set("dns1", v.(string))
	}
	if v, found := p.p["dns2"]; found {
		u.Set("dns2", v.(string))
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["guestcidraddress"]; found {
		u.Set("guestcidraddress", v.(string))
	}
	if v, found := p.p["internaldns1"]; found {
		u.Set("internaldns1", v.(string))
	}
	if v, found := p.p["internaldns2"]; found {
		u.Set("internaldns2", v.(string))
	}
	if v, found := p.p["ip6dns1"]; found {
		u.Set("ip6dns1", v.(string))
	}
	if v, found := p.p["ip6dns2"]; found {
		u.Set("ip6dns2", v.(string))
	}
	if v, found := p.p["localstorageenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("localstorageenabled", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networktype"]; found {
		u.Set("networktype", v.(string))
	}
	if v, found := p.p["securitygroupenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("securitygroupenabled", vv)
	}
	return u
}

func (p *CreateZoneParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *CreateZoneParams) SetDns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns1"] = v
	return
}

func (p *CreateZoneParams) SetDns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns2"] = v
	return
}

func (p *CreateZoneParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
	return
}

func (p *CreateZoneParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateZoneParams) SetGuestcidraddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestcidraddress"] = v
	return
}

func (p *CreateZoneParams) SetInternaldns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internaldns1"] = v
	return
}

func (p *CreateZoneParams) SetInternaldns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internaldns2"] = v
	return
}

func (p *CreateZoneParams) SetIp6dns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns1"] = v
	return
}

func (p *CreateZoneParams) SetIp6dns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns2"] = v
	return
}

func (p *CreateZoneParams) SetLocalstorageenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["localstorageenabled"] = v
	return
}

func (p *CreateZoneParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateZoneParams) SetNetworktype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networktype"] = v
	return
}

func (p *CreateZoneParams) SetSecuritygroupenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupenabled"] = v
	return
}

// You should always use this function to get a new CreateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewCreateZoneParams(dns1 string, internaldns1 string, name string, networktype string) *CreateZoneParams {
	p := &CreateZoneParams{}
	p.p = make(map[string]interface{})
	p.p["dns1"] = dns1
	p.p["internaldns1"] = internaldns1
	p.p["name"] = name
	p.p["networktype"] = networktype
	return p
}

// Creates a Zone.
func (s *ZoneService) CreateZone(p *CreateZoneParams) (*CreateZoneResponse, error) {
	resp, err := s.cs.newRequest("createZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateZoneResponse struct {
	Allocationstate string `json:"allocationstate"`
	Capacity        []struct {
		Capacitytotal int64  `json:"capacitytotal"`
		Capacityused  int64  `json:"capacityused"`
		Clusterid     string `json:"clusterid"`
		Clustername   string `json:"clustername"`
		Percentused   string `json:"percentused"`
		Podid         string `json:"podid"`
		Podname       string `json:"podname"`
		Type          int    `json:"type"`
		Zoneid        string `json:"zoneid"`
		Zonename      string `json:"zonename"`
	} `json:"capacity"`
	Description           string            `json:"description"`
	Dhcpprovider          string            `json:"dhcpprovider"`
	Displaytext           string            `json:"displaytext"`
	Dns1                  string            `json:"dns1"`
	Dns2                  string            `json:"dns2"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainname            string            `json:"domainname"`
	Guestcidraddress      string            `json:"guestcidraddress"`
	Id                    string            `json:"id"`
	Internaldns1          string            `json:"internaldns1"`
	Internaldns2          string            `json:"internaldns2"`
	Ip6dns1               string            `json:"ip6dns1"`
	Ip6dns2               string            `json:"ip6dns2"`
	Localstorageenabled   bool              `json:"localstorageenabled"`
	Name                  string            `json:"name"`
	Networktype           string            `json:"networktype"`
	Resourcedetails       map[string]string `json:"resourcedetails"`
	Securitygroupsenabled bool              `json:"securitygroupsenabled"`
	Tags                  []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
	Zonetoken string `json:"zonetoken"`
}

type DedicateZoneParams struct {
	p map[string]interface{}
}

func (p *DedicateZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DedicateZoneParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DedicateZoneParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *DedicateZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new DedicateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDedicateZoneParams(domainid string, zoneid string) *DedicateZoneParams {
	p := &DedicateZoneParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	p.p["zoneid"] = zoneid
	return p
}

// Dedicates a zones.
func (s *ZoneService) DedicateZone(p *DedicateZoneParams) (*DedicateZoneResponse, error) {
	resp, err := s.cs.newRequest("dedicateZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DedicateZoneResponse struct {
	JobID           string `json:"jobid"`
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type DeleteZoneParams struct {
	p map[string]interface{}
}

func (p *DeleteZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteZoneParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDeleteZoneParams(id string) *DeleteZoneParams {
	p := &DeleteZoneParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Zone.
func (s *ZoneService) DeleteZone(p *DeleteZoneParams) (*DeleteZoneResponse, error) {
	resp, err := s.cs.newRequest("deleteZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteZoneResponse struct {
	Displaytext string `json:"displaytext"`
	Success     bool   `json:"success"`
}

func (r *DeleteZoneResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteZoneResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableOutOfBandManagementForZoneParams struct {
	p map[string]interface{}
}

func (p *DisableOutOfBandManagementForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DisableOutOfBandManagementForZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new DisableOutOfBandManagementForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDisableOutOfBandManagementForZoneParams(zoneid string) *DisableOutOfBandManagementForZoneParams {
	p := &DisableOutOfBandManagementForZoneParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Disables out-of-band management for a zone
func (s *ZoneService) DisableOutOfBandManagementForZone(p *DisableOutOfBandManagementForZoneParams) (*DisableOutOfBandManagementForZoneResponse, error) {
	resp, err := s.cs.newRequest("disableOutOfBandManagementForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableOutOfBandManagementForZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DisableOutOfBandManagementForZoneResponse struct {
	JobID       string `json:"jobid"`
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type EnableOutOfBandManagementForZoneParams struct {
	p map[string]interface{}
}

func (p *EnableOutOfBandManagementForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *EnableOutOfBandManagementForZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new EnableOutOfBandManagementForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewEnableOutOfBandManagementForZoneParams(zoneid string) *EnableOutOfBandManagementForZoneParams {
	p := &EnableOutOfBandManagementForZoneParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Enables out-of-band management for a zone
func (s *ZoneService) EnableOutOfBandManagementForZone(p *EnableOutOfBandManagementForZoneParams) (*EnableOutOfBandManagementForZoneResponse, error) {
	resp, err := s.cs.newRequest("enableOutOfBandManagementForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableOutOfBandManagementForZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type EnableOutOfBandManagementForZoneResponse struct {
	JobID       string `json:"jobid"`
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type ListDedicatedZonesParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedZonesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListDedicatedZonesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListDedicatedZonesParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
	return
}

func (p *ListDedicatedZonesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListDedicatedZonesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDedicatedZonesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDedicatedZonesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListDedicatedZonesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListDedicatedZonesParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewListDedicatedZonesParams() *ListDedicatedZonesParams {
	p := &ListDedicatedZonesParams{}
	p.p = make(map[string]interface{})
	return p
}

// List dedicated zones.
func (s *ZoneService) ListDedicatedZones(p *ListDedicatedZonesParams) (*ListDedicatedZonesResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedZones", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedZonesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedZonesResponse struct {
	Count          int              `json:"count"`
	DedicatedZones []*DedicatedZone `json:"dedicatedzone"`
}

type DedicatedZone struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type ListVmwareDcsParams struct {
	p map[string]interface{}
}

func (p *ListVmwareDcsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVmwareDcsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVmwareDcsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVmwareDcsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVmwareDcsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListVmwareDcsParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewListVmwareDcsParams(zoneid string) *ListVmwareDcsParams {
	p := &ListVmwareDcsParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetVmwareDcID(keyword string, zoneid string, opts ...OptionFunc) (string, int, error) {
	p := &ListVmwareDcsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["zoneid"] = zoneid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVmwareDcs(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.VmwareDcs[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VmwareDcs {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Retrieves VMware DC(s) associated with a zone.
func (s *ZoneService) ListVmwareDcs(p *ListVmwareDcsParams) (*ListVmwareDcsResponse, error) {
	resp, err := s.cs.newRequest("listVmwareDcs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVmwareDcsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVmwareDcsResponse struct {
	Count     int         `json:"count"`
	VmwareDcs []*VmwareDc `json:"vmwaredc"`
}

type VmwareDc struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Vcenter string `json:"vcenter"`
	Zoneid  int64  `json:"zoneid"`
}

type ListZonesParams struct {
	p map[string]interface{}
}

func (p *ListZonesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["available"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("available", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networktype"]; found {
		u.Set("networktype", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *ListZonesParams) SetAvailable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["available"] = v
	return
}

func (p *ListZonesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListZonesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListZonesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListZonesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListZonesParams) SetNetworktype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networktype"] = v
	return
}

func (p *ListZonesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListZonesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListZonesParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
	return
}

func (p *ListZonesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListZonesParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewListZonesParams() *ListZonesParams {
	p := &ListZonesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListZonesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListZones(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Zones[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Zones {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneByName(name string, opts ...OptionFunc) (*Zone, int, error) {
	id, count, err := s.GetZoneID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetZoneByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneByID(id string, opts ...OptionFunc) (*Zone, int, error) {
	p := &ListZonesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListZones(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Zones[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Zone UUID: %s!", id)
}

// Lists zones
func (s *ZoneService) ListZones(p *ListZonesParams) (*ListZonesResponse, error) {
	resp, err := s.cs.newRequest("listZones", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListZonesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListZonesResponse struct {
	Count int     `json:"count"`
	Zones []*Zone `json:"zone"`
}

type Zone struct {
	Allocationstate string `json:"allocationstate"`
	Capacity        []struct {
		Capacitytotal int64  `json:"capacitytotal"`
		Capacityused  int64  `json:"capacityused"`
		Clusterid     string `json:"clusterid"`
		Clustername   string `json:"clustername"`
		Percentused   string `json:"percentused"`
		Podid         string `json:"podid"`
		Podname       string `json:"podname"`
		Type          int    `json:"type"`
		Zoneid        string `json:"zoneid"`
		Zonename      string `json:"zonename"`
	} `json:"capacity"`
	Description           string            `json:"description"`
	Dhcpprovider          string            `json:"dhcpprovider"`
	Displaytext           string            `json:"displaytext"`
	Dns1                  string            `json:"dns1"`
	Dns2                  string            `json:"dns2"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainname            string            `json:"domainname"`
	Guestcidraddress      string            `json:"guestcidraddress"`
	Id                    string            `json:"id"`
	Internaldns1          string            `json:"internaldns1"`
	Internaldns2          string            `json:"internaldns2"`
	Ip6dns1               string            `json:"ip6dns1"`
	Ip6dns2               string            `json:"ip6dns2"`
	Localstorageenabled   bool              `json:"localstorageenabled"`
	Name                  string            `json:"name"`
	Networktype           string            `json:"networktype"`
	Resourcedetails       map[string]string `json:"resourcedetails"`
	Securitygroupsenabled bool              `json:"securitygroupsenabled"`
	Tags                  []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
	Zonetoken string `json:"zonetoken"`
}

type ReleaseDedicatedZoneParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ReleaseDedicatedZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewReleaseDedicatedZoneParams(zoneid string) *ReleaseDedicatedZoneParams {
	p := &ReleaseDedicatedZoneParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Release dedication of zone
func (s *ZoneService) ReleaseDedicatedZone(p *ReleaseDedicatedZoneParams) (*ReleaseDedicatedZoneResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type ReleaseDedicatedZoneResponse struct {
	JobID       string `json:"jobid"`
	Displaytext string `json:"displaytext"`
	Success     bool   `json:"success"`
}

type RemoveVmwareDcParams struct {
	p map[string]interface{}
}

func (p *RemoveVmwareDcParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *RemoveVmwareDcParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new RemoveVmwareDcParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewRemoveVmwareDcParams(zoneid string) *RemoveVmwareDcParams {
	p := &RemoveVmwareDcParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Remove a VMware datacenter from a zone.
func (s *ZoneService) RemoveVmwareDc(p *RemoveVmwareDcParams) (*RemoveVmwareDcResponse, error) {
	resp, err := s.cs.newRequest("removeVmwareDc", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveVmwareDcResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveVmwareDcResponse struct {
	Displaytext string `json:"displaytext"`
	Success     bool   `json:"success"`
}

func (r *RemoveVmwareDcResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias RemoveVmwareDcResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateZoneParams struct {
	p map[string]interface{}
}

func (p *UpdateZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), vv)
			i++
		}
	}
	if v, found := p.p["dhcpprovider"]; found {
		u.Set("dhcpprovider", v.(string))
	}
	if v, found := p.p["dns1"]; found {
		u.Set("dns1", v.(string))
	}
	if v, found := p.p["dns2"]; found {
		u.Set("dns2", v.(string))
	}
	if v, found := p.p["dnssearchorder"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("dnssearchorder", vv)
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["guestcidraddress"]; found {
		u.Set("guestcidraddress", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["internaldns1"]; found {
		u.Set("internaldns1", v.(string))
	}
	if v, found := p.p["internaldns2"]; found {
		u.Set("internaldns2", v.(string))
	}
	if v, found := p.p["ip6dns1"]; found {
		u.Set("ip6dns1", v.(string))
	}
	if v, found := p.p["ip6dns2"]; found {
		u.Set("ip6dns2", v.(string))
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["localstorageenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("localstorageenabled", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateZoneParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *UpdateZoneParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *UpdateZoneParams) SetDhcpprovider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpprovider"] = v
	return
}

func (p *UpdateZoneParams) SetDns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns1"] = v
	return
}

func (p *UpdateZoneParams) SetDns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns2"] = v
	return
}

func (p *UpdateZoneParams) SetDnssearchorder(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dnssearchorder"] = v
	return
}

func (p *UpdateZoneParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
	return
}

func (p *UpdateZoneParams) SetGuestcidraddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestcidraddress"] = v
	return
}

func (p *UpdateZoneParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateZoneParams) SetInternaldns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internaldns1"] = v
	return
}

func (p *UpdateZoneParams) SetInternaldns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internaldns2"] = v
	return
}

func (p *UpdateZoneParams) SetIp6dns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns1"] = v
	return
}

func (p *UpdateZoneParams) SetIp6dns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns2"] = v
	return
}

func (p *UpdateZoneParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
	return
}

func (p *UpdateZoneParams) SetLocalstorageenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["localstorageenabled"] = v
	return
}

func (p *UpdateZoneParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new UpdateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewUpdateZoneParams(id string) *UpdateZoneParams {
	p := &UpdateZoneParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a Zone.
func (s *ZoneService) UpdateZone(p *UpdateZoneParams) (*UpdateZoneResponse, error) {
	resp, err := s.cs.newRequest("updateZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateZoneResponse struct {
	Allocationstate string `json:"allocationstate"`
	Capacity        []struct {
		Capacitytotal int64  `json:"capacitytotal"`
		Capacityused  int64  `json:"capacityused"`
		Clusterid     string `json:"clusterid"`
		Clustername   string `json:"clustername"`
		Percentused   string `json:"percentused"`
		Podid         string `json:"podid"`
		Podname       string `json:"podname"`
		Type          int    `json:"type"`
		Zoneid        string `json:"zoneid"`
		Zonename      string `json:"zonename"`
	} `json:"capacity"`
	Description           string            `json:"description"`
	Dhcpprovider          string            `json:"dhcpprovider"`
	Displaytext           string            `json:"displaytext"`
	Dns1                  string            `json:"dns1"`
	Dns2                  string            `json:"dns2"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainname            string            `json:"domainname"`
	Guestcidraddress      string            `json:"guestcidraddress"`
	Id                    string            `json:"id"`
	Internaldns1          string            `json:"internaldns1"`
	Internaldns2          string            `json:"internaldns2"`
	Ip6dns1               string            `json:"ip6dns1"`
	Ip6dns2               string            `json:"ip6dns2"`
	Localstorageenabled   bool              `json:"localstorageenabled"`
	Name                  string            `json:"name"`
	Networktype           string            `json:"networktype"`
	Resourcedetails       map[string]string `json:"resourcedetails"`
	Securitygroupsenabled bool              `json:"securitygroupsenabled"`
	Tags                  []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
	Zonetoken string `json:"zonetoken"`
}
