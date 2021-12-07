// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// GetMessageThreadRequest represents TL type `getMessageThread#7af23e3e`.
type GetMessageThreadRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the message
	MessageID int64
}

// GetMessageThreadRequestTypeID is TL type id of GetMessageThreadRequest.
const GetMessageThreadRequestTypeID = 0x7af23e3e

// Ensuring interfaces in compile-time for GetMessageThreadRequest.
var (
	_ bin.Encoder     = &GetMessageThreadRequest{}
	_ bin.Decoder     = &GetMessageThreadRequest{}
	_ bin.BareEncoder = &GetMessageThreadRequest{}
	_ bin.BareDecoder = &GetMessageThreadRequest{}
)

func (g *GetMessageThreadRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessageThreadRequest) String() string {
	if g == nil {
		return "GetMessageThreadRequest(nil)"
	}
	type Alias GetMessageThreadRequest
	return fmt.Sprintf("GetMessageThreadRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageThreadRequest) TypeID() uint32 {
	return GetMessageThreadRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageThreadRequest) TypeName() string {
	return "getMessageThread"
}

// TypeInfo returns info about TL type.
func (g *GetMessageThreadRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageThread",
		ID:   GetMessageThreadRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessageThreadRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageThread#7af23e3e as nil")
	}
	b.PutID(GetMessageThreadRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageThreadRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageThread#7af23e3e as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageThreadRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageThread#7af23e3e to nil")
	}
	if err := b.ConsumeID(GetMessageThreadRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageThread#7af23e3e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageThreadRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageThread#7af23e3e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageThread#7af23e3e: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageThread#7af23e3e: field message_id: %w", err)
		}
		g.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessageThreadRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageThread#7af23e3e as nil")
	}
	b.ObjStart()
	b.PutID("getMessageThread")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessageThreadRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageThread#7af23e3e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessageThread"); err != nil {
				return fmt.Errorf("unable to decode getMessageThread#7af23e3e: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageThread#7af23e3e: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageThread#7af23e3e: field message_id: %w", err)
			}
			g.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessageThreadRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetMessageThreadRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetMessageThread invokes method getMessageThread#7af23e3e returning error if any.
func (c *Client) GetMessageThread(ctx context.Context, request *GetMessageThreadRequest) (*MessageThreadInfo, error) {
	var result MessageThreadInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}