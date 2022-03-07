package ship

import (
	"github.com/zhongshuwen/zswchain-go"
	"github.com/zhongshuwen/zswchain-go/ecc"
)

// State History Plugin Requests

type GetStatusRequestV0 struct {
}

type GetBlocksAckRequestV0 struct {
	NumMessages uint32
}

type GetBlocksRequestV0 struct {
	StartBlockNum       uint32
	EndBlockNum         uint32
	MaxMessagesInFlight uint32
	HavePositions       []*BlockPosition
	IrreversibleOnly    bool
	FetchBlock          bool
	FetchTraces         bool
	FetchDeltas         bool
}

// State History Plugin Results
type GetStatusResultV0 struct {
	Head                 *BlockPosition
	LastIrreversible     *BlockPosition
	TraceBeginBlock      uint32
	TraceEndBlock        uint32
	ChainStateBeginBlock uint32
	ChainStateEndBlock   uint32
}

type GetBlocksResultV0 struct {
	Head             *BlockPosition
	LastIrreversible *BlockPosition
	ThisBlock        *BlockPosition         `eos:"optional"`
	PrevBlock        *BlockPosition         `eos:"optional"`
	Block            *SignedBlockBytes      `eos:"optional"`
	Traces           *TransactionTraceArray `eos:"optional"`
	Deltas           *TableDeltaArray       `eos:"optional"`
}

// State History Plugin version of EOS structs
type BlockPosition struct {
	BlockNum uint32
	BlockID  zsw.Checksum256
}

type Row struct {
	Present bool
	Data    []byte
}

type ActionTraceV0 struct {
	ActionOrdinal        zsw.Varuint32
	CreatorActionOrdinal zsw.Varuint32
	Receipt              *ActionReceipt `eos:"optional"`
	Receiver             zsw.Name
	Act                  *Action
	ContextFree          bool
	Elapsed              int64
	Console              zsw.SafeString
	AccountRamDeltas     []*zsw.AccountRAMDelta
	Except               string `eos:"optional"`
	ErrorCode            uint64 `eos:"optional"`
}

type Action struct {
	Account       zsw.AccountName
	Name          zsw.ActionName
	Authorization []zsw.PermissionLevel
	Data          []byte
}

type ActionReceiptV0 struct {
	Receiver       zsw.Name
	ActDigest      zsw.Checksum256
	GlobalSequence uint64
	RecvSequence   uint64
	AuthSequence   []AccountAuthSequence
	CodeSequence   zsw.Varuint32
	ABISequence    zsw.Varuint32
}

type AccountAuthSequence struct {
	Account  zsw.Name
	Sequence uint64
}

type TableDeltaV0 struct {
	Name string
	Rows []Row
}

type PartialTransactionV0 struct {
	Expiration            uint32
	RefBlockNum           uint16
	RefBlockPrefix        uint32
	MaxNetUsageWords      zsw.Varuint32
	MaxCpuUsageMs         uint8
	DelaySec              zsw.Varuint32
	TransactionExtensions []*Extension
	Signatures            []ecc.Signature
	ContextFreeData       []byte
}

type TransactionTraceV0 struct {
	ID              zsw.Checksum256 `json:"id"`
	Status          zsw.TransactionStatus
	CPUUsageUS      uint32               `json:"cpu_usage_us"`
	NetUsageWords   zsw.Varuint32        `json:"net_usage_words"`
	Elapsed         zsw.Int64            `json:"elapsed"`
	NetUsage        uint64               `json:"net_usage"`
	Scheduled       bool                 `json:"scheduled"`
	ActionTraces    []*ActionTrace       `json:"action_traces"`
	AccountDelta    *zsw.AccountRAMDelta `json:"account_delta" eos:"optional"`
	Except          string               `json:"except" eos:"optional"`
	ErrorCode       uint64               `json:"error_code" eos:"optional"`
	FailedDtrxTrace *TransactionTrace    `json:"failed_dtrx_trace" eos:"optional"`
	Partial         *PartialTransaction  `json:"partial" eos:"optional"`
}

type SignedBlockHeader struct {
	zsw.BlockHeader
	ProducerSignature ecc.Signature // no pointer!!
}

type TransactionReceipt struct {
	zsw.TransactionReceiptHeader
	Trx *Transaction
}

//type TransactionID zsw.Checksum256

type SignedBlock struct {
	SignedBlockHeader
	Transactions    []*TransactionReceipt
	BlockExtensions []*Extension
}

type SignedBlockBytes SignedBlock

func (s *SignedBlockBytes) AsSignedBlock() *SignedBlock {
	if s == nil {
		return nil
	}
	ss := SignedBlock(*s)
	return &ss
}

func (s *SignedBlockBytes) UnmarshalBinary(decoder *zsw.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return zsw.UnmarshalBinary(data, (*SignedBlock)(s))
}

type Extension struct {
	Type uint16
	Data []byte
}
