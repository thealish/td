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

// ChatFilterInfo represents TL type `chatFilterInfo#c7bff533`.
type ChatFilterInfo struct {
	// Unique chat filter identifier
	ID int32
	// The title of the filter; 1-12 characters without line feeds
	Title string
	// The chosen or default icon name for short filter representation. One of "All",
	// "Unread", "Unmuted", "Bots", "Channels", "Groups", "Private", "Custom", "Setup", "Cat"
	// "Crown", "Favorite", "Flower", "Game", "Home", "Love", "Mask", "Party", "Sport",
	// "Study", "Trade", "Travel", "Work", "Airplane", "Book", "Light", "Like", "Money",
	// "Note", "Palette"
	IconName string
}

// ChatFilterInfoTypeID is TL type id of ChatFilterInfo.
const ChatFilterInfoTypeID = 0xc7bff533

// Ensuring interfaces in compile-time for ChatFilterInfo.
var (
	_ bin.Encoder     = &ChatFilterInfo{}
	_ bin.Decoder     = &ChatFilterInfo{}
	_ bin.BareEncoder = &ChatFilterInfo{}
	_ bin.BareDecoder = &ChatFilterInfo{}
)

func (c *ChatFilterInfo) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ID == 0) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.IconName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatFilterInfo) String() string {
	if c == nil {
		return "ChatFilterInfo(nil)"
	}
	type Alias ChatFilterInfo
	return fmt.Sprintf("ChatFilterInfo%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatFilterInfo) TypeID() uint32 {
	return ChatFilterInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatFilterInfo) TypeName() string {
	return "chatFilterInfo"
}

// TypeInfo returns info about TL type.
func (c *ChatFilterInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatFilterInfo",
		ID:   ChatFilterInfoTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "IconName",
			SchemaName: "icon_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatFilterInfo) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatFilterInfo#c7bff533 as nil")
	}
	b.PutID(ChatFilterInfoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatFilterInfo) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatFilterInfo#c7bff533 as nil")
	}
	b.PutInt32(c.ID)
	b.PutString(c.Title)
	b.PutString(c.IconName)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatFilterInfo) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatFilterInfo#c7bff533 to nil")
	}
	if err := b.ConsumeID(ChatFilterInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatFilterInfo#c7bff533: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatFilterInfo) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatFilterInfo#c7bff533 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatFilterInfo#c7bff533: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatFilterInfo#c7bff533: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatFilterInfo#c7bff533: field icon_name: %w", err)
		}
		c.IconName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatFilterInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatFilterInfo#c7bff533 as nil")
	}
	b.ObjStart()
	b.PutID("chatFilterInfo")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(c.ID)
	b.Comma()
	b.FieldStart("title")
	b.PutString(c.Title)
	b.Comma()
	b.FieldStart("icon_name")
	b.PutString(c.IconName)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatFilterInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatFilterInfo#c7bff533 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatFilterInfo"); err != nil {
				return fmt.Errorf("unable to decode chatFilterInfo#c7bff533: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatFilterInfo#c7bff533: field id: %w", err)
			}
			c.ID = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatFilterInfo#c7bff533: field title: %w", err)
			}
			c.Title = value
		case "icon_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatFilterInfo#c7bff533: field icon_name: %w", err)
			}
			c.IconName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (c *ChatFilterInfo) GetID() (value int32) {
	if c == nil {
		return
	}
	return c.ID
}

// GetTitle returns value of Title field.
func (c *ChatFilterInfo) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// GetIconName returns value of IconName field.
func (c *ChatFilterInfo) GetIconName() (value string) {
	if c == nil {
		return
	}
	return c.IconName
}
