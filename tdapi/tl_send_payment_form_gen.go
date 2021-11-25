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

// SendPaymentFormRequest represents TL type `sendPaymentForm#5b9133ff`.
type SendPaymentFormRequest struct {
	// Chat identifier of the Invoice message
	ChatID int64
	// Message identifier
	MessageID int64
	// Payment form identifier returned by getPaymentForm
	PaymentFormID int64
	// Identifier returned by validateOrderInfo, or an empty string
	OrderInfoID string
	// Identifier of a chosen shipping option, if applicable
	ShippingOptionID string
	// The credentials chosen by user for payment
	Credentials InputCredentialsClass
	// Chosen by the user amount of tip in the smallest units of the currency
	TipAmount int64
}

// SendPaymentFormRequestTypeID is TL type id of SendPaymentFormRequest.
const SendPaymentFormRequestTypeID = 0x5b9133ff

// Ensuring interfaces in compile-time for SendPaymentFormRequest.
var (
	_ bin.Encoder     = &SendPaymentFormRequest{}
	_ bin.Decoder     = &SendPaymentFormRequest{}
	_ bin.BareEncoder = &SendPaymentFormRequest{}
	_ bin.BareDecoder = &SendPaymentFormRequest{}
)

func (s *SendPaymentFormRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageID == 0) {
		return false
	}
	if !(s.PaymentFormID == 0) {
		return false
	}
	if !(s.OrderInfoID == "") {
		return false
	}
	if !(s.ShippingOptionID == "") {
		return false
	}
	if !(s.Credentials == nil) {
		return false
	}
	if !(s.TipAmount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendPaymentFormRequest) String() string {
	if s == nil {
		return "SendPaymentFormRequest(nil)"
	}
	type Alias SendPaymentFormRequest
	return fmt.Sprintf("SendPaymentFormRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendPaymentFormRequest) TypeID() uint32 {
	return SendPaymentFormRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendPaymentFormRequest) TypeName() string {
	return "sendPaymentForm"
}

// TypeInfo returns info about TL type.
func (s *SendPaymentFormRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendPaymentForm",
		ID:   SendPaymentFormRequestTypeID,
	}
	if s == nil {
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
		{
			Name:       "PaymentFormID",
			SchemaName: "payment_form_id",
		},
		{
			Name:       "OrderInfoID",
			SchemaName: "order_info_id",
		},
		{
			Name:       "ShippingOptionID",
			SchemaName: "shipping_option_id",
		},
		{
			Name:       "Credentials",
			SchemaName: "credentials",
		},
		{
			Name:       "TipAmount",
			SchemaName: "tip_amount",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendPaymentFormRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPaymentForm#5b9133ff as nil")
	}
	b.PutID(SendPaymentFormRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendPaymentFormRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPaymentForm#5b9133ff as nil")
	}
	b.PutLong(s.ChatID)
	b.PutLong(s.MessageID)
	b.PutLong(s.PaymentFormID)
	b.PutString(s.OrderInfoID)
	b.PutString(s.ShippingOptionID)
	if s.Credentials == nil {
		return fmt.Errorf("unable to encode sendPaymentForm#5b9133ff: field credentials is nil")
	}
	if err := s.Credentials.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendPaymentForm#5b9133ff: field credentials: %w", err)
	}
	b.PutLong(s.TipAmount)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendPaymentFormRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPaymentForm#5b9133ff to nil")
	}
	if err := b.ConsumeID(SendPaymentFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendPaymentFormRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPaymentForm#5b9133ff to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field message_id: %w", err)
		}
		s.MessageID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field payment_form_id: %w", err)
		}
		s.PaymentFormID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field order_info_id: %w", err)
		}
		s.OrderInfoID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field shipping_option_id: %w", err)
		}
		s.ShippingOptionID = value
	}
	{
		value, err := DecodeInputCredentials(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field credentials: %w", err)
		}
		s.Credentials = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field tip_amount: %w", err)
		}
		s.TipAmount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendPaymentFormRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPaymentForm#5b9133ff as nil")
	}
	b.ObjStart()
	b.PutID("sendPaymentForm")
	b.FieldStart("chat_id")
	b.PutLong(s.ChatID)
	b.FieldStart("message_id")
	b.PutLong(s.MessageID)
	b.FieldStart("payment_form_id")
	b.PutLong(s.PaymentFormID)
	b.FieldStart("order_info_id")
	b.PutString(s.OrderInfoID)
	b.FieldStart("shipping_option_id")
	b.PutString(s.ShippingOptionID)
	b.FieldStart("credentials")
	if s.Credentials == nil {
		return fmt.Errorf("unable to encode sendPaymentForm#5b9133ff: field credentials is nil")
	}
	if err := s.Credentials.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendPaymentForm#5b9133ff: field credentials: %w", err)
	}
	b.FieldStart("tip_amount")
	b.PutLong(s.TipAmount)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendPaymentFormRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPaymentForm#5b9133ff to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendPaymentForm"); err != nil {
				return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field message_id: %w", err)
			}
			s.MessageID = value
		case "payment_form_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field payment_form_id: %w", err)
			}
			s.PaymentFormID = value
		case "order_info_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field order_info_id: %w", err)
			}
			s.OrderInfoID = value
		case "shipping_option_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field shipping_option_id: %w", err)
			}
			s.ShippingOptionID = value
		case "credentials":
			value, err := DecodeTDLibJSONInputCredentials(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field credentials: %w", err)
			}
			s.Credentials = value
		case "tip_amount":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendPaymentForm#5b9133ff: field tip_amount: %w", err)
			}
			s.TipAmount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SendPaymentFormRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetMessageID returns value of MessageID field.
func (s *SendPaymentFormRequest) GetMessageID() (value int64) {
	return s.MessageID
}

// GetPaymentFormID returns value of PaymentFormID field.
func (s *SendPaymentFormRequest) GetPaymentFormID() (value int64) {
	return s.PaymentFormID
}

// GetOrderInfoID returns value of OrderInfoID field.
func (s *SendPaymentFormRequest) GetOrderInfoID() (value string) {
	return s.OrderInfoID
}

// GetShippingOptionID returns value of ShippingOptionID field.
func (s *SendPaymentFormRequest) GetShippingOptionID() (value string) {
	return s.ShippingOptionID
}

// GetCredentials returns value of Credentials field.
func (s *SendPaymentFormRequest) GetCredentials() (value InputCredentialsClass) {
	return s.Credentials
}

// GetTipAmount returns value of TipAmount field.
func (s *SendPaymentFormRequest) GetTipAmount() (value int64) {
	return s.TipAmount
}

// SendPaymentForm invokes method sendPaymentForm#5b9133ff returning error if any.
func (c *Client) SendPaymentForm(ctx context.Context, request *SendPaymentFormRequest) (*PaymentResult, error) {
	var result PaymentResult

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}