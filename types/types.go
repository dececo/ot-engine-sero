package types

import "math/big"
import "github.com/sero-cash/go-sero/crypto"

var (
	publishSig = []byte("Publish(string,uint256)")
	solveSig   = []byte("Solve(string,string,string)")
	acceptSig  = []byte("Accept(string)")
	rejectSig  = []byte("Reject(string)")
	confirmSig = []byte("Confirm(string,string)")

	PublishSigHash = crypto.Keccak256Hash(publishSig)
	SolveSigHash   = crypto.Keccak256Hash(solveSig)
	AcceptSigHash  = crypto.Keccak256Hash(acceptSig)
	RejectSigHash  = crypto.Keccak256Hash(rejectSig)
	ConfirmSigHash = crypto.Keccak256Hash(confirmSig)
)

type PublishEvent struct {
	Block     uint64
	Tx        string
	Mission   string
	Reward    *big.Int
	Publisher string
}

type SolveEvent struct {
	Block    uint64
	Tx       string
	Solution string
	Mission  string
	Data     string
	Solver   string
}

type ProcessEvent struct {
	Block    uint64
	Tx       string
	Solution string
	Time     string // type is string, just for output
	Status   string // accept or reject
}
type Process ProcessEvent
type AcceptEvent ProcessEvent
type RejectEvent ProcessEvent

type ConfirmEvent struct {
	Block       uint64
	Tx          string
	Solution    string
	Arbitration string
}
