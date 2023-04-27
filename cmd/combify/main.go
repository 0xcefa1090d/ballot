package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type bytecode struct {
	Bytecode string `json:"bytecode"`
}

type contractType struct {
	ContractName       string      `json:"contractName"`
	SourceID           string      `json:"sourceId"`
	DeploymentBytecode bytecode    `json:"deploymentBytecode"`
	ABI                interface{} `json:"abi"`
}

type combined struct {
	ABI interface{} `json:"abi"`
	Bin string      `json:"bin"`
}

func main() {
	cmd := cobra.Command{
		Use:   "combify",
		Short: "EthPMv3 Contract Type to Combined JSON Output",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var input contractType

			decoder := json.NewDecoder(os.Stdin)
			if err := decoder.Decode(&input); err != nil {
				log.Fatal(err)
			}

			output := map[string]map[string]combined{
				"contracts": {
					fmt.Sprintf("%s:%s", input.SourceID, input.ContractName): {
						ABI: input.ABI,
						Bin: input.DeploymentBytecode.Bytecode,
					},
				},
			}

			encoder := json.NewEncoder(os.Stdout)
			if err := encoder.Encode(&output); err != nil {
				log.Fatal(err)
			}
		},
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
