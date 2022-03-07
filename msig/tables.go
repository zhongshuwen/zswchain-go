package msig

import zsw "github.com/zhongshuwen/zswchain-go"

type ProposalRow struct {
	ProposalName       zsw.Name              `json:"proposal_name"`
	RequestedApprovals []zsw.PermissionLevel `json:"requested_approvals"`
	ProvidedApprovals  []zsw.PermissionLevel `json:"provided_approvals"`
	PackedTransaction  zsw.HexBytes          `json:"packed_transaction"`
}
