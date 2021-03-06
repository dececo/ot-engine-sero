package process

import (
	"github.com/sero-cash/go-sero/core/types"
	otTypes "github.com/dececo/ot-engine-sero/types"
	"math/big"
	"log"
	"fmt"
	"strings"
	"github.com/sero-cash/go-sero/accounts/abi"
	"github.com/dececo/ot-engine-sero/contracts"
)

func ParseOTLog(vLog types.Log) (err error) {
	switch vLog.Topics[0].String() {
	case otTypes.PublishSigHash.Hex():
		fmt.Println("Publish")
		row, err1 := Publish(vLog)
		fmt.Println(row)
		if err1 != nil {
			return err1
		}
	case otTypes.SolveSigHash.Hex():
		fmt.Println("Solve")
		row, err1 := Solve(vLog)
		fmt.Println(row)
		if err1 != nil {
			return err1
		}
	case otTypes.AcceptSigHash.Hex():
		fmt.Println("Accept")
		row, err1 := Accept(vLog)
		fmt.Println(row)
		if err1 != nil {
			return err1
		}
	case otTypes.RejectSigHash.Hex():
		fmt.Println("Reject")
		row, err1 := Reject(vLog)
		fmt.Println(row)
		if err1 != nil {
			return err1
		}
	case otTypes.ConfirmSigHash.Hex():
		fmt.Println("Confirm")
		row, err1 := Confirm(vLog)
		fmt.Println(row)
		if err1 != nil {
			return err1
		}
	default:
		fmt.Println("UNKNOWN Event Log")
	}
	return
}

func Publish(vLog types.Log) (p otTypes.PublishEvent, err error) {
	event := struct {
		MissionId string
		Reward    *big.Int
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return p, err
	}
	err = contractAbi.Unpack(&event, "Publish", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return p, err
	}

	fmt.Printf("missionId: %s, reward: %s\n\n", event.MissionId, event.Reward.String())

	p.Mission = event.MissionId
	p.Reward = event.Reward
	p.Block = vLog.BlockNumber
	p.Tx = vLog.TxHash.String()
	p.Publisher = vLog.Address.Base58()
	return
}

func Solve(vLog types.Log) (s otTypes.SolveEvent, err error) {
	event := struct {
		SolutionId string
		MissionId  string
		Data       string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return s, err
	}
	err = contractAbi.Unpack(&event, "Solve", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return s, err
	}

	fmt.Printf("solutionId: %s, missionId: %s, data: %s\n", event.SolutionId, event.MissionId, event.Data)
	s.Solution = event.SolutionId
	s.Mission = event.MissionId
	s.Data = event.Data
	s.Block = vLog.BlockNumber
	s.Tx = vLog.TxHash.String()
	return
}

func Accept(vLog types.Log) (a otTypes.AcceptEvent, err error) {
	event := struct {
		SolutionId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return a, err
	}
	err = contractAbi.Unpack(&event, "Accept", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return a, err
	}

	fmt.Printf("solutionId: %s\n", event.SolutionId)
	a.Solution = event.SolutionId
	a.Block = vLog.BlockNumber
	a.Tx = vLog.TxHash.String()
	return
}

func Reject(vLog types.Log) (r otTypes.RejectEvent, err error) {
	event := struct {
		SolutionId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return r, err
	}
	err = contractAbi.Unpack(&event, "Reject", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return r, err
	}

	fmt.Printf("solutionId: %s\n", event.SolutionId)
	r.Solution = event.SolutionId
	r.Block = vLog.BlockNumber
	r.Tx = vLog.TxHash.String()
	return
}

func Confirm(vLog types.Log) (c otTypes.ConfirmEvent, err error) {
	fmt.Println("Confirm")
	event := struct {
		SolutionId    string
		ArbitrationId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return c, err
	}
	err = contractAbi.Unpack(&event, "Confirm", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return c, err
	}

	fmt.Printf("solutionId: %s, missionId: %s\n", event.SolutionId, event.ArbitrationId)
	c.Solution = event.SolutionId
	c.Arbitration = event.ArbitrationId
	c.Block = vLog.BlockNumber
	c.Tx = vLog.TxHash.String()
	return
}
