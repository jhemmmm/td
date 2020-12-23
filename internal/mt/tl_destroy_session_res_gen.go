// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// DestroySessionOk represents TL type `destroy_session_ok#e22045fc`.
type DestroySessionOk struct {
	// SessionID field of DestroySessionOk.
	SessionID int64
}

// DestroySessionOkTypeID is TL type id of DestroySessionOk.
const DestroySessionOkTypeID = 0xe22045fc

// String implements fmt.Stringer.
func (d *DestroySessionOk) String() string {
	if d == nil {
		return "DestroySessionOk(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DestroySessionOk")
	sb.WriteString("{\n")
	sb.WriteString("\tSessionID: ")
	sb.WriteString(fmt.Sprint(d.SessionID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DestroySessionOk) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session_ok#e22045fc as nil")
	}
	b.PutID(DestroySessionOkTypeID)
	b.PutLong(d.SessionID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DestroySessionOk) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session_ok#e22045fc to nil")
	}
	if err := b.ConsumeID(DestroySessionOkTypeID); err != nil {
		return fmt.Errorf("unable to decode destroy_session_ok#e22045fc: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session_ok#e22045fc: field session_id: %w", err)
		}
		d.SessionID = value
	}
	return nil
}

// construct implements constructor of DestroySessionResClass.
func (d DestroySessionOk) construct() DestroySessionResClass { return &d }

// Ensuring interfaces in compile-time for DestroySessionOk.
var (
	_ bin.Encoder = &DestroySessionOk{}
	_ bin.Decoder = &DestroySessionOk{}

	_ DestroySessionResClass = &DestroySessionOk{}
)

// DestroySessionNone represents TL type `destroy_session_none#62d350c9`.
type DestroySessionNone struct {
	// SessionID field of DestroySessionNone.
	SessionID int64
}

// DestroySessionNoneTypeID is TL type id of DestroySessionNone.
const DestroySessionNoneTypeID = 0x62d350c9

// String implements fmt.Stringer.
func (d *DestroySessionNone) String() string {
	if d == nil {
		return "DestroySessionNone(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DestroySessionNone")
	sb.WriteString("{\n")
	sb.WriteString("\tSessionID: ")
	sb.WriteString(fmt.Sprint(d.SessionID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DestroySessionNone) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session_none#62d350c9 as nil")
	}
	b.PutID(DestroySessionNoneTypeID)
	b.PutLong(d.SessionID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DestroySessionNone) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session_none#62d350c9 to nil")
	}
	if err := b.ConsumeID(DestroySessionNoneTypeID); err != nil {
		return fmt.Errorf("unable to decode destroy_session_none#62d350c9: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session_none#62d350c9: field session_id: %w", err)
		}
		d.SessionID = value
	}
	return nil
}

// construct implements constructor of DestroySessionResClass.
func (d DestroySessionNone) construct() DestroySessionResClass { return &d }

// Ensuring interfaces in compile-time for DestroySessionNone.
var (
	_ bin.Encoder = &DestroySessionNone{}
	_ bin.Decoder = &DestroySessionNone{}

	_ DestroySessionResClass = &DestroySessionNone{}
)

// DestroySessionResClass represents DestroySessionRes generic type.
//
// Example:
//  g, err := DecodeDestroySessionRes(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *DestroySessionOk: // destroy_session_ok#e22045fc
//  case *DestroySessionNone: // destroy_session_none#62d350c9
//  default: panic(v)
//  }
type DestroySessionResClass interface {
	bin.Encoder
	bin.Decoder
	construct() DestroySessionResClass
	fmt.Stringer
}

// DecodeDestroySessionRes implements binary de-serialization for DestroySessionResClass.
func DecodeDestroySessionRes(buf *bin.Buffer) (DestroySessionResClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DestroySessionOkTypeID:
		// Decoding destroy_session_ok#e22045fc.
		v := DestroySessionOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", err)
		}
		return &v, nil
	case DestroySessionNoneTypeID:
		// Decoding destroy_session_none#62d350c9.
		v := DestroySessionNone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", bin.NewUnexpectedID(id))
	}
}

// DestroySessionRes boxes the DestroySessionResClass providing a helper.
type DestroySessionResBox struct {
	DestroySessionRes DestroySessionResClass
}

// Decode implements bin.Decoder for DestroySessionResBox.
func (b *DestroySessionResBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DestroySessionResBox to nil")
	}
	v, err := DecodeDestroySessionRes(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DestroySessionRes = v
	return nil
}

// Encode implements bin.Encode for DestroySessionResBox.
func (b *DestroySessionResBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DestroySessionRes == nil {
		return fmt.Errorf("unable to encode DestroySessionResClass as nil")
	}
	return b.DestroySessionRes.Encode(buf)
}