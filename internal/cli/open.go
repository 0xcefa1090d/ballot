package cli

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/0xcefa1090d/bounty/internal/bindings"
	"github.com/0xcefa1090d/bounty/internal/bindings/mocks"
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
		extSigner, err := external.NewExternalSigner(signer)
		if err != nil {
			log.Fatal(err)
		}
		defer extSigner.Close()

		account, err := selectAccount(extSigner)
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewClefTransactor(extSigner, account)

		data, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		input := openInput{}
		if err := yaml.Unmarshal(data, &input); err != nil {
			log.Fatal(err)
		}

		conn, err := ethclient.Dial(node)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		bountyAddr := common.HexToAddress("0x3f4d4C07D616fa6822c37c0d66b6Bd5F5Db666d4")
		bountyContract, err := bindings.NewStartVoteBounty(bountyAddr, conn)
		if err != nil {
			log.Fatal(err)
		}

		tokenAddr := common.HexToAddress(input.Reward.Token)
		tokenContract, err := mocks.NewTokenMock(tokenAddr, conn)
		if err != nil {
			log.Fatal(err)
		}

		value := big.NewInt(int64(input.Reward.Value))

		if allowance, err := tokenContract.Allowance(&bind.CallOpts{}, auth.From, bountyAddr); err != nil {
			log.Fatal(err)
		} else if allowance.Cmp(value) == -1 {
			txn, err := tokenContract.Approve(auth, bountyAddr, value)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(txn.Hash())
		}

		if auth.Value, err = bountyContract.OPENBOUNTYCOST(&bind.CallOpts{}); err != nil {
			log.Fatal(err)
		}

		evmScript := prepareEVMScript(common.HexToAddress("0x40907540d8a6c65c637785e8f8b742ae6b0b9968"), input.Script)
		txn, err := bountyContract.OpenBounty(auth, tokenAddr, value, input.Metadata, evmScript)
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
