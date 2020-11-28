// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesGetWebPagePreviewRequest represents TL type `messages.getWebPagePreview#8b68b0cc`.
type MessagesGetWebPagePreviewRequest struct {
	// Flags field of MessagesGetWebPagePreviewRequest.
	Flags bin.Fields
	// Message field of MessagesGetWebPagePreviewRequest.
	Message string
	// Entities field of MessagesGetWebPagePreviewRequest.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// MessagesGetWebPagePreviewRequestTypeID is TL type id of MessagesGetWebPagePreviewRequest.
const MessagesGetWebPagePreviewRequestTypeID = 0x8b68b0cc

// Encode implements bin.Encoder.
func (g *MessagesGetWebPagePreviewRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getWebPagePreview#8b68b0cc as nil")
	}
	b.PutID(MessagesGetWebPagePreviewRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field flags: %w", err)
	}
	b.PutString(g.Message)
	if g.Flags.Has(3) {
		b.PutVectorHeader(len(g.Entities))
		for idx, v := range g.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetEntities sets value of Entities conditional field.
func (g *MessagesGetWebPagePreviewRequest) SetEntities(value []MessageEntityClass) {
	g.Flags.Set(3)
	g.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (g *MessagesGetWebPagePreviewRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if !g.Flags.Has(3) {
		return value, false
	}
	return g.Entities, true
}

// Decode implements bin.Decoder.
func (g *MessagesGetWebPagePreviewRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getWebPagePreview#8b68b0cc to nil")
	}
	if err := b.ConsumeID(MessagesGetWebPagePreviewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field message: %w", err)
		}
		g.Message = value
	}
	if g.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field entities: %w", err)
			}
			g.Entities = append(g.Entities, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetWebPagePreviewRequest.
var (
	_ bin.Encoder = &MessagesGetWebPagePreviewRequest{}
	_ bin.Decoder = &MessagesGetWebPagePreviewRequest{}
)