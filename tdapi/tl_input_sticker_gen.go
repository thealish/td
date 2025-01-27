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

// InputSticker represents TL type `inputSticker#9b1829b0`.
type InputSticker struct {
	// File with the sticker; must fit in a 512x512 square. For WEBP stickers and masks the
	// file must be in PNG format, which will be converted to WEBP server-side. Otherwise,
	// the file must be local or uploaded within a week. See https://core.telegram
	// org/animated_stickers#technical-requirements for technical requirements
	Sticker InputFileClass
	// Emojis corresponding to the sticker
	Emojis string
	// Sticker format
	Format StickerFormatClass
	// Position where the mask is placed; pass null if not specified
	MaskPosition MaskPosition
}

// InputStickerTypeID is TL type id of InputSticker.
const InputStickerTypeID = 0x9b1829b0

// Ensuring interfaces in compile-time for InputSticker.
var (
	_ bin.Encoder     = &InputSticker{}
	_ bin.Decoder     = &InputSticker{}
	_ bin.BareEncoder = &InputSticker{}
	_ bin.BareDecoder = &InputSticker{}
)

func (i *InputSticker) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Sticker == nil) {
		return false
	}
	if !(i.Emojis == "") {
		return false
	}
	if !(i.Format == nil) {
		return false
	}
	if !(i.MaskPosition.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputSticker) String() string {
	if i == nil {
		return "InputSticker(nil)"
	}
	type Alias InputSticker
	return fmt.Sprintf("InputSticker%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputSticker) TypeID() uint32 {
	return InputStickerTypeID
}

// TypeName returns name of type in TL schema.
func (*InputSticker) TypeName() string {
	return "inputSticker"
}

// TypeInfo returns info about TL type.
func (i *InputSticker) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputSticker",
		ID:   InputStickerTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "Emojis",
			SchemaName: "emojis",
		},
		{
			Name:       "Format",
			SchemaName: "format",
		},
		{
			Name:       "MaskPosition",
			SchemaName: "mask_position",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputSticker) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSticker#9b1829b0 as nil")
	}
	b.PutID(InputStickerTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputSticker) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSticker#9b1829b0 as nil")
	}
	if i.Sticker == nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field sticker is nil")
	}
	if err := i.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field sticker: %w", err)
	}
	b.PutString(i.Emojis)
	if i.Format == nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field format is nil")
	}
	if err := i.Format.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field format: %w", err)
	}
	if err := i.MaskPosition.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field mask_position: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputSticker) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSticker#9b1829b0 to nil")
	}
	if err := b.ConsumeID(InputStickerTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSticker#9b1829b0: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputSticker) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSticker#9b1829b0 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSticker#9b1829b0: field sticker: %w", err)
		}
		i.Sticker = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputSticker#9b1829b0: field emojis: %w", err)
		}
		i.Emojis = value
	}
	{
		value, err := DecodeStickerFormat(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSticker#9b1829b0: field format: %w", err)
		}
		i.Format = value
	}
	{
		if err := i.MaskPosition.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputSticker#9b1829b0: field mask_position: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputSticker) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSticker#9b1829b0 as nil")
	}
	b.ObjStart()
	b.PutID("inputSticker")
	b.Comma()
	b.FieldStart("sticker")
	if i.Sticker == nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field sticker is nil")
	}
	if err := i.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field sticker: %w", err)
	}
	b.Comma()
	b.FieldStart("emojis")
	b.PutString(i.Emojis)
	b.Comma()
	b.FieldStart("format")
	if i.Format == nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field format is nil")
	}
	if err := i.Format.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field format: %w", err)
	}
	b.Comma()
	b.FieldStart("mask_position")
	if err := i.MaskPosition.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputSticker#9b1829b0: field mask_position: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputSticker) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSticker#9b1829b0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputSticker"); err != nil {
				return fmt.Errorf("unable to decode inputSticker#9b1829b0: %w", err)
			}
		case "sticker":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputSticker#9b1829b0: field sticker: %w", err)
			}
			i.Sticker = value
		case "emojis":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputSticker#9b1829b0: field emojis: %w", err)
			}
			i.Emojis = value
		case "format":
			value, err := DecodeTDLibJSONStickerFormat(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputSticker#9b1829b0: field format: %w", err)
			}
			i.Format = value
		case "mask_position":
			if err := i.MaskPosition.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode inputSticker#9b1829b0: field mask_position: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSticker returns value of Sticker field.
func (i *InputSticker) GetSticker() (value InputFileClass) {
	if i == nil {
		return
	}
	return i.Sticker
}

// GetEmojis returns value of Emojis field.
func (i *InputSticker) GetEmojis() (value string) {
	if i == nil {
		return
	}
	return i.Emojis
}

// GetFormat returns value of Format field.
func (i *InputSticker) GetFormat() (value StickerFormatClass) {
	if i == nil {
		return
	}
	return i.Format
}

// GetMaskPosition returns value of MaskPosition field.
func (i *InputSticker) GetMaskPosition() (value MaskPosition) {
	if i == nil {
		return
	}
	return i.MaskPosition
}
