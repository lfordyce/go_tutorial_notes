package redis

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack/v4"
)

type Type string

const (
	SiteType   Type = "site"
	SwitchType Type = "switch"
)

type Event interface {
	GetID() string
	GetType() Type
	SetID(id string)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func NewEvent(t Type) (Event, error) {
	b := &Base{
		Type: t,
	}
	switch t {
	case SiteType:
		return &SiteEvent{
			Base: b,
		}, nil
	case SwitchType:
		return &SwitchEvent{
			Base: b,
		}, nil
	}
	return nil, fmt.Errorf("type %v not supported", t)
}

type Base struct {
	ID   string `json:"id"`
	Type Type   `json:"type"`
}

func (o *Base) GetID() string {
	return o.ID
}

func (o *Base) SetID(id string) {
	o.ID = id
}

func (o *Base) GetType() Type {
	return o.Type
}

func (o *Base) String() string {
	return fmt.Sprintf("id:%s type:%s", o.ID, o.Type)
}

type SwitchEvent struct {
	*Base
	Site string `json:"site"`
}

func (s *SwitchEvent) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
	//return msgpack.Marshal(s)
}

func (s *SwitchEvent) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
	//return msgpack.Unmarshal(data, s)
}

type SiteEvent struct {
	*Base
	ActiveSite string
}

func (s *SiteEvent) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(s)
}

func (s *SiteEvent) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, s)
}
