package entity

import (
	"encoding/json"
	"strconv"
	"strings"

	"lraft/utils"
)

type ElectionPriority int64

var (
	Disabled   = ElectionPriority(-1)
	NotElected = ElectionPriority(0)
	MinValue   = ElectionPriority(1)
)

type Endpoint struct {
	ip   string
	port int64
	desc string
}

func NewEndpoint(ip string, port int64) Endpoint {
	return Endpoint{
		ip:   ip,
		port: port,
		desc: ip + ":" + strconv.FormatInt(port, 10),
	}
}

func (e Endpoint) GetIP() string {
	return e.ip
}

func (e Endpoint) GetPort() int64 {
	return e.port
}

func (e Endpoint) GetDesc() string {
	return e.desc
}

func (e Endpoint) Copy() Endpoint {
	return Endpoint{
		ip:   e.ip,
		port: e.port,
		desc: e.desc,
	}
}

func (e Endpoint) Equal(other Endpoint) bool {
	return strings.Compare(e.ip, other.ip) == 0 && e.port == other.port
}

var EmptyPeers = make([]*PeerId, 0)

type PeerId struct {
	endpoint Endpoint
	idx      int64
	priority ElectionPriority
	checksum uint64
	desc     string
}

func NewPeerId(endpoint Endpoint, idx int64, priority ElectionPriority) *PeerId {
	return &PeerId{
		endpoint: endpoint,
		idx:      idx,
		priority: priority,
		checksum: 0,
		desc:     "",
	}
}

func (p *PeerId) GetIP() string {
	return p.endpoint.ip
}

func (p *PeerId) GetPort() int64 {
	return p.endpoint.port
}

func (p *PeerId) GetEndpoint() Endpoint {
	return p.endpoint
}

func (p *PeerId) Copy() *PeerId {
	return &PeerId{
		endpoint: p.endpoint,
		idx:      p.idx,
		priority: p.priority,
		checksum: p.checksum,
		desc:     p.desc,
	}
}

func (p *PeerId) Parse(s string) bool {
	if s == "" {
		return false
	}
	tmps := strings.Split(s, ":")
	if len(tmps) < 2 || len(tmps) > 4 {
		return false
	}
	port := utils.ParseToInt64(tmps[1])
	p.endpoint = NewEndpoint(tmps[0], port)
	switch len(tmps) {
	case 3:
		p.idx = utils.ParseToInt64(tmps[2])
	case 4:
		if tmps[2] == "" {
			p.idx = int64(0)
		} else {
			p.idx = utils.ParseToInt64(tmps[2])
		}
		p.priority = ElectionPriority(utils.ParseToInt64(tmps[3]))
	default:
		return false
	}
	return true
}

func (p *PeerId) Checksum() uint64 {
	str := p.GetDesc()
	if p.checksum == 0 {
		p.checksum = utils.Checksum([]byte(str))
	}
	return p.checksum
}

func (p *PeerId) GetDesc() string {
	if p.desc != "" {
		return p.desc
	}
	p.desc += p.endpoint.desc
	if p.idx != 0 {
		p.desc += ":" + strconv.FormatInt(p.idx, 10)
	}
	if p.priority != Disabled {
		if p.priority == 0 {
			p.desc += ":"
		}
		p.desc += ":" + strconv.FormatInt(int64(p.priority), 10)
	}
	return p.desc
}

func (p *PeerId) IsPriorityNotElected() bool {
	return p.priority == NotElected
}

func (p *PeerId) IsPriorityDisabled() bool {
	return p.priority == Disabled
}

func (p *PeerId) IsEmpty() bool {
	return strings.Compare(utils.IPAny, p.endpoint.ip) == 0 && p.endpoint.port == 0 && p.idx == 0
}

func (p *PeerId) Equal(other *PeerId) bool {
	return p.endpoint.Equal(other.endpoint) && p.idx == other.idx && p.priority == other.priority
}

func (p *PeerId) Encode() []byte {
	b, err := json.Marshal(p)
	utils.CheckErr(err)
	return b
}

func (p *PeerId) Decode(b []byte) {
	err := json.Unmarshal(b, p)
	utils.CheckErr(err)
}
