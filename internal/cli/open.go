package cli

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/0xcefa1090d/bounty/internal/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type (
	reward struct {
		Token string `yaml:"token"`
		Value int    `yaml:"value"`
	}
	action struct {
		Target   string `yaml:"target"`
		Calldata string `yaml:"calldata"`
	}
	openInput struct {
		Metadata string   `yaml:"metadata"`
		Reward   reward   `yaml:"reward"`
		Script   []action `yaml:"script"`
	}
)

var (
	agentABI    abi.ABI
	rawAgentABI string = `[{"constant":false,"inputs":[{"name":"_target","type":"address"},{"name":"_ethValue","type":"uint256"},{"name":"_data","type":"bytes"}],"name":"execute","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
)

var openCmd = &cobra.Command{
	Use:   "open infile",
	Short: "Open a new vote creation bounty",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := ethclient.Dial(node)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		extSigner, err := external.NewExternalSigner(signer)
		if err != nil {
			log.Fatal(err)
		}
		defer extSigner.Close()

		contract, err := bindings.NewStartVoteBounty(common.HexToAddress("0x3f4d4C07D616fa6822c37c0d66b6Bd5F5Db666d4"), conn)
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewClefTransactor(extSigner, extSigner.Accounts()[0])
		if auth.Value, err = contract.OPENBOUNTYCOST(&bind.CallOpts{}); err != nil {
			log.Fatal(err)
		}

		data, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		input := openInput{}
		if err := yaml.Unmarshal(data, &input); err != nil {
			log.Fatal(err)
		}

		evmScript := prepareEVMScript(common.HexToAddress("0x40907540d8a6c65c637785e8f8b742ae6b0b9968"), input.Script)

		txn, err := contract.OpenBounty(auth, common.HexToAddress(input.Reward.Token), big.NewInt(int64(input.Reward.Value)), input.Metadata, evmScript)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(txn.Hash())
	},
}

func init() {
	if err := agentABI.UnmarshalJSON([]byte(rawAgentABI)); err != nil {
		log.Fatal(err)
	}
}

func prepareEVMScript(agent common.Address, script []action) []byte {
	evmScript := []byte{00, 00, 00, 01}

	for _, op := range script {
		calldata, err := agentABI.Pack("execute", common.HexToAddress(op.Target), big.NewInt(0), common.FromHex(op.Calldata))
		if err != nil {
			log.Fatal(err)
		}
		length := common.LeftPadBytes(big.NewInt(int64(len(calldata))).Bytes(), 8)
		evmScript = bytes.Join([][]byte{evmScript, agent.Bytes(), length, calldata}, []byte{})
	}
	return evmScript
}
