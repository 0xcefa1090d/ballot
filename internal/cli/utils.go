package cli

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/common"
)

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

func selectAccount(extSigner *external.ExternalSigner) (accounts.Account, error) {
	if accountAddr != "" {
		account := accounts.Account{Address: common.HexToAddress(accountAddr), URL: accounts.URL{}}
		if extSigner.Contains(account) {
			return account, nil
		}
		return accounts.Account{}, fmt.Errorf("the provided account is not available: %s", accountAddr)
	}

	available := extSigner.Accounts()
	switch len(available) {
	case 0:
		return accounts.Account{}, errors.New("the external signer does not have any available accounts")
	case 1:
		return available[0], nil
	default:
		for idx, acct := range available {
			fmt.Printf("[%d]: %s\n", idx, acct.Address)
		}

		fmt.Print("Select an available account: ")

		selection, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return accounts.Account{}, err
		}

		idx, err := strconv.Atoi(strings.TrimSpace(selection))
		if err != nil {
			return accounts.Account{}, err
		}

		if idx < 0 || len(available) < idx {
			return accounts.Account{}, errors.New("invalid selection")
		}

		return available[idx], nil
	}
}
