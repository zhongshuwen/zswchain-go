package ship

import (
	"github.com/zhongshuwen/zswchain-go"
)

// Request
var RequestVariant = zsw.NewVariantDefinition([]zsw.VariantType{
	{"get_status_request_v0", (*GetStatusRequestV0)(nil)},
	{"get_blocks_request_v0", (*GetBlocksRequestV0)(nil)},
	{"get_blocks_ack_request_v0", (*GetBlocksAckRequestV0)(nil)},
})

type Request struct {
	zsw.BaseVariant
}

func (r *Request) UnmarshalBinary(decoder *zsw.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, RequestVariant)
}

// Result
var ResultVariant = zsw.NewVariantDefinition([]zsw.VariantType{
	{"get_status_result_v0", (*GetStatusResultV0)(nil)},
	{"get_blocks_result_v0", (*GetBlocksResultV0)(nil)},
})

type Result struct {
	zsw.BaseVariant
}

func (r *Result) UnmarshalBinary(decoder *zsw.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ResultVariant)
}

// TransactionTrace
var TransactionTraceVariant = zsw.NewVariantDefinition([]zsw.VariantType{
	{"transaction_trace_v0", (*TransactionTraceV0)(nil)},
})

type TransactionTrace struct {
	zsw.BaseVariant
}

type TransactionTraceArray struct {
	Elem []*TransactionTrace
}

func (t *TransactionTraceArray) AsTransactionTracesV0() (out []*TransactionTraceV0) {
	if t == nil || t.Elem == nil {
		return nil
	}
	for _, e := range t.Elem {
		switch v := e.Impl.(type) {
		case *TransactionTraceV0:
			out = append(out, v)

		default:
			panic("wrong type for conversion")
		}
	}
	return out
}

func (r *TransactionTraceArray) UnmarshalBinary(decoder *zsw.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return zsw.UnmarshalBinary(data, &r.Elem)
}

func (r *TransactionTrace) UnmarshalBinary(decoder *zsw.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, TransactionTraceVariant)
}

// ActionTrace
var ActionTraceVariant = zsw.NewVariantDefinition([]zsw.VariantType{
	{"action_trace_v0", (*ActionTraceV0)(nil)},
})

type ActionTrace struct {
	zsw.BaseVariant
}

func (r *ActionTrace) UnmarshalBinary(decoder *zsw.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ActionTraceVariant)
}

// PartialTransaction
var PartialTransactionVariant = zsw.NewVariantDefinition([]zsw.VariantType{
	{"partial_transaction_v0", (*PartialTransactionV0)(nil)},
})

type PartialTransaction struct {
	zsw.BaseVariant
}

func (r *PartialTransaction) UnmarshalBinary(decoder *zsw.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, PartialTransactionVariant)
}

// TableDelta
var TableDeltaVariant = zsw.NewVariantDefinition([]zsw.VariantType{
	{"table_delta_v0", (*TableDeltaV0)(nil)},
})

type TableDelta struct {
	zsw.BaseVariant
}

func (d *TableDelta) UnmarshalBinary(decoder *zsw.Decoder) error {
	return d.BaseVariant.UnmarshalBinaryVariant(decoder, TableDeltaVariant)
}

type TableDeltaArray struct {
	Elem []*TableDelta
}

func (d *TableDeltaArray) UnmarshalBinary(decoder *zsw.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return zsw.UnmarshalBinary(data, &d.Elem)
}

func (t *TableDeltaArray) AsTableDeltasV0() (out []*TableDeltaV0) {
	if t == nil || t.Elem == nil {
		return nil
	}
	for _, e := range t.Elem {
		switch v := e.Impl.(type) {
		case *TableDeltaV0:
			out = append(out, v)

		default:
			panic("wrong type for conversion")
		}
	}
	return out
}

// Transaction
var TransactionVariant = zsw.NewVariantDefinition([]zsw.VariantType{
	{"transaction_id", (*zsw.Checksum256)(nil)},
	{"packed_transaction", (*zsw.PackedTransaction)(nil)},
})

type Transaction struct {
	zsw.BaseVariant
}

func (d *Transaction) UnmarshalBinary(decoder *zsw.Decoder) error {
	return d.BaseVariant.UnmarshalBinaryVariant(decoder, TransactionVariant)
}

// ActionReceipt
var ActionReceiptVariant = zsw.NewVariantDefinition([]zsw.VariantType{
	{"action_receipt_v0", (*ActionReceiptV0)(nil)},
})

type ActionReceipt struct {
	zsw.BaseVariant
}

func (r *ActionReceipt) UnmarshalBinary(decoder *zsw.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ActionReceiptVariant)
}
