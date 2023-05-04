// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StartVoteBountyMetaData contains all meta data concerning the StartVoteBounty contract.
var StartVoteBountyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voting\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"ApplyCloseBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"voteId\",\"type\":\"uint256\"}],\"name\":\"ClaimBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"CommitCloseBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addedAmount\",\"type\":\"uint256\"}],\"name\":\"IncreaseBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"}],\"name\":\"OpenBounty\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OPEN_BOUNTY_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"applyCloseBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cache\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"cacheBlockHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"}],\"name\":\"calculateIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"receiptIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"headerRlpBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofRlpBytes\",\"type\":\"bytes\"}],\"name\":\"claimBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"commitCloseBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getCommitCloseBountyBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getRefundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"increaseAmount\",\"type\":\"uint256\"}],\"name\":\"increaseBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"}],\"name\":\"openBounty\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"}],\"name\":\"openBounty\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x0x60a03461007157601f61214f38819003918201601f19168301916001600160401b038311848410176100765780849260209460405283398101031261007157516001600160a01b0381168103610071576080526040516120c2908161008d82396080518181816103ad0152610c460152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306239cc71461010757806317bf72c6146101025780631af57b96146100fd5780631ce63f1a146100f8578063269e1d1a146100f35780632fa970d8146100ee5780635d306f34146100e95780636f0d0a5d146100e4578063a24dc3cb146100df578063ad22eb62146100da578063bee8a902146100d5578063c2084b04146100d0578063d6841c94146100cb578063f5ba9296146100c65763fa89401a146100c157600080fd5b61083e565b610763565b610705565b6105ff565b610597565b610505565b6104b9565b610489565b61044f565b6103dc565b610397565b6102b5565b610164565b610138565b346101335760203660031901126101335760043560005260016020526020604060002054604051908152f35b600080fd5b346101335760203660031901126101335760043560005260006020526020604060002054604051908152f35b346101335760203660031901126101335760043560005260036020526020604060002054604051908152f35b600435906001600160a01b038216820361013357565b604435906001600160a01b038216820361013357565b602435906001600160a01b038216820361013357565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b0382111761020357604052565b6101d2565b608081019081106001600160401b0382111761020357604052565b90601f801991011681019081106001600160401b0382111761020357604052565b60405190610251826101e8565b565b6001600160401b03811161020357601f01601f191660200190565b81601f820112156101335780359061028582610253565b926102936040519485610223565b8284526020838301011161013357816000926020809301838601378301015290565b346101335760a0366003190112610133576102ce610190565b6102d66101a6565b6001600160401b039190606435838111610133576102f890369060040161026e565b608435938411610133576103306103809161031a61039396369060040161026e565b9060208151910120906020815191012090610aff565b60208151910120916103726040519384926020840196602435908890606092959493608083019660018060a01b03809316845260208401521660408201520152565b03601f198101835282610223565b5190206040519081529081906020820190565b0390f35b34610133576000366003190112610133576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101335760e0366003190112610133576103f5610190565b6103fd6101a6565b6001600160401b03919060a4358381116101335761041f90369060040161026e565b9060c4359384116101335761043b61044d94369060040161026e565b92608435916064359160243590610b5d565b005b34610133576020366003190112610133576001600160a01b03610470610190565b1660005260026020526020604060002054604051908152f35b34610133576020366003190112610133576004358040908115610133576000526000602052604060002055600080f35b3461013357600036600319011261013357602060405166470de4df8200008152f35b606090600319011261013357600435906024356001600160a01b0381168103610133579060443590565b3461013357610549610516366104db565b604080513360208201908152918101949094526001600160a01b0390921660608401526080830152918160a08101610372565b51902060009080825260036020526040822054156105935760016020524360408320557fd994ec1efb74ba41455cf04e5da2d899411306f2fa57dec2a98fcf3ce0bb30638280a280f35b5080fd5b60a0366003190112610133576105ab6101bc565b6001600160401b03606435818111610133576105cb90369060040161026e565b91608435918211610133576020926105ea6105f793369060040161026e565b916044359060043561129a565b604051908152f35b34610133576103726106e761064e610616366104db565b604080513360208201908152918101949094526001600160a01b038316606085015260808401919091529094909290829060a0820190565b519020600092818452600360205260408420549161066d8315156108ce565b61069d610684826000526001602052604060002090565b548061068f43611043565b1090816106f3575b506108ce565b846106b2826000526003602052604060002090565b557fc236bdbce1e4cfdbe02cc5976c05c076ceedc985cda30253cca17a15cd8a01b58580a233906001600160a01b031661119e565b6106f03361142a565b80f35b90506106fe43611052565b1138610697565b608036600319011261013357610719610190565b6001600160401b036044358181116101335761073990369060040161026e565b91606435918211610133576020926107586105f793369060040161026e565b91602435904261129a565b346101335760803660031901126101335761077c6101bc565b606435801561013357604080513360208201908152600435928201929092526001600160a01b0384166060820152604435608082015261044d9391906107c58160a08101610372565b5190208060005260036020526107eb836040600020546107e68115156108ce565b610b50565b6107ff826000526003602052604060002090565b556040518381527f18ee4fa302d6692472b161e97fddc236d75cf9f3a31fe117ed398dac7d0ceccd90602090a2309033906001600160a01b03166113ab565b3461013357602036600319011261013357610857610190565b6001600160a01b03811660009081526002602052604081208054919283929182156108c95783928392838093555af11561088e5780f35b60405162461bcd60e51b815260206004820152601360248201527211551217d514905394d1915497d19052531151606a1b6044820152606490fd5b505050fd5b1561013357565b634e487b7160e01b600052601160045260246000fd5b60001981146108fa5760010190565b6108d5565b634e487b7160e01b600052603260045260246000fd5b8051156109225760200190565b6108ff565b8051600110156109225760400190565b8051600210156109225760600190565b8051600b1015610922576101800190565b8051600310156109225760800190565b805160101015610922576102200190565b80518210156109225760209160051b010190565b60005b8381106109a05750506000910152565b8181015183820152602001610990565b909291926109bd81610253565b916109cb6040519384610223565b82948284528282011161013357602061025193019061098d565b9060a0828203126101335781516001600160401b0381116101335782019080601f83011215610133578151610a1c926020016109b0565b916020820151916040810151916080606083015192015190565b5190811515820361013357565b51906001600160401b038216820361013357565b6101408183031261013357610a6b81610a36565b92610a7860208301610a36565b92610a8560408401610a43565b92610a9260608201610a43565b92610a9f60808301610a43565b92610aac60a08401610a43565b9260c08101519260e08201519261010083015192610120810151906001600160401b03821161013357019080601f83011215610133578151610af0926020016109b0565b90565b6040513d6000823e3d90fd5b9190604051926020840152604083015260408252606082018281106001600160401b0382111761020357604052565b9066470de4df82000082018092116108fa57565b90600182018092116108fa57565b919082018092116108fa57565b604080516001600160a01b038381166020808401918252838501879052918716606084015260808084018990528352939a999792949093909291610ba260a082610223565b51902096610bba886000526003602052604060002090565b5498610bc78a15156108ce565b610bd090611466565b80518587830151928340831493841597610c03610c1b96610c1696610c0e94610c219c610e72575b50509d9a9b9d6108ce565b6060830151116108ce565b015192610eaa565b610ee7565b91611642565b97610c34610c2f8a51151590565b6108ce565b6060600099019160018060a01b0395867f000000000000000000000000000000000000000000000000000000000000000016915b845180518d101561013357610c9e610c92610c848f8794610979565b51516001600160a01b031690565b6001600160a01b031690565b03610e63577f0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394610cdb87610cd38f8951610979565b510151610915565b5103610e63579084939291610d2487610d1c8f999885610cff8c610d109351610979565b5101518380825183010191016109e5565b50505050998851610979565b510151610927565b5182516305a55c1f60e41b815260048101829052909790600081602481885afa8015610e5e578392610d6c92600092610e31575b508981519101209089815191012090610aff565b87815191012003610e1957505050610de992610dd46102519a9b610c9294610dc2610ddc956000610da78e6000526003602052604060002090565b556001600160a01b0316600090815260026020526040902090565b610dcc8154610b2e565b905551610979565b510151610937565b516001600160a01b031690565b93838516907f62f6344c3fb7a53546ee98642ba528b6b561027f33810bac1cdd4032916fa75d600080a41661119e565b919b610e269196506108eb565b9a9192939490610c68565b610e4d91923d8091833e610e458183610223565b810190610a57565b985050505050505050509038610d58565b610af3565b90949392919a610e26906108eb565b610e889192506000526000602052604060002090565b54143880610bf8565b60405190610e9e826101e8565b60006020838281520152565b610eb2610e91565b50602081519160405192610ec5846101e8565b835201602082015290565b6001600160401b0381116102035760051b60200190565b610ef081610fba565b1561013357610efe81610fe1565b610f0781610ed0565b91610f156040519384610223565b818352601f19610f2483610ed0565b0160005b818110610fa3575050610f49602080920151610f438161113e565b90610b50565b6000905b838210610f5b575050505090565b610f9781610f6b610f9d936110b8565b90610f74610244565b8281528187820152610f86868a610979565b52610f918589610979565b50610b50565b916108eb565b90610f4d565b602090610fae610e91565b82828801015201610f28565b805115610fdb57602060c09101515160001a10610fd657600190565b600090565b50600090565b805115610fdb5760009060208101908151610ffb8161113e565b81018091116108fa579151905181018091116108fa5791905b8281106110215750905090565b61102a816110b8565b81018091116108fa5761103d90916108eb565b90611014565b60ff198101919082116108fa57565b607f198101919082116108fa57565b60bf198101919082116108fa57565b60200390602082116108fa57565b6000198101919082116108fa57565b60f6198101919082116108fa57565b60b6198101919082116108fa57565b919082039182116108fa57565b805160001a9060808210156110ce575050600190565b60b88210156110e957506110e4610af091611052565b610b42565b9060c081101561110d5760b51991600160b783602003016101000a91015104010190565b9060f882101561112457506110e4610af091611061565b60010151602082900360f7016101000a90040160f5190190565b5160001a60808110156111515750600090565b60b881108015611188575b156111675750600190565b60c081101561117c576110e4610af09161109c565b6110e4610af09161108d565b5060c0811015801561115c575060f8811061115c565b60446000928380936111d1966040519363a9059cbb60e01b855260018060a01b0316600485015260248401525af161120f565b156111d857565b60405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b6044820152606490fd5b3d901561123c578060201461122d571561122857600090565b600190565b5060206000803e600051151590565b806000803e6000fd5b9060209161125e8151809281855285808601910161098d565b601f01601f1916010190565b92610af0949261128c9285526020850152608060408501526080840190611245565b916060818403910152611245565b92610af0937f61c6abb5ad2430a2a28d7fe158faa33e09784bd8ce4660816300feea5d448a59959266470de4df8200003414806113a2575b80611398575b6112e1906108ce565b6112f8835160208501208551602087012090610aff565b805160209182012060408051339381019384529081018590526001600160a01b038416606082015260808101919091526113358160a08101610372565b519020968791611359611352846000526003602052604060002090565b54156108ce565b8661136e846000526003602052604060002090565b5560018060a01b031694859461138c60405192839233978b8561126a565b0390a4309033906113ab565b50428210156112d8565b508415156112d2565b600092836064926113e7968295604051946323b872dd60e01b865260018060a01b03809216600487015216602485015260448401525af161120f565b156113ee57565b60405162461bcd60e51b81526020600482015260146024820152731514905394d1915497d19493d357d1905253115160621b6044820152606490fd5b6000808066470de4df82000081945af11561088e57565b6040519061144e82610208565b60006060838281528260208201528260408201520152565b61146e611441565b5061147b610c1682610eaa565b90600b825111156101335761148e611441565b91805160051015610922576114a660c08201516114eb565b6020840152805160081015610922576114d4816114ca6101206114da9401516114eb565b6040860152610947565b516114eb565b606083015260208151910120815290565b80518015159081611521575b5015610133576115069061152d565b90519060208110611515575090565b6020036101000a900490565b602191501115386114f7565b90602082019161153d835161113e565b9251908382018092116108fa57519283039283116108fa579190565b6040519061156682610208565b6060808360008152600060208201528160408201520152565b908151811015610922570160200190565b60405190606082018281106001600160401b03821117610203576040526060604083600081528260208201520152565b906115ca82610ed0565b6115d76040519182610223565b82815280926115e8601f1991610ed0565b019060005b8281106115f957505050565b602090611604611590565b828285010152016115ed565b9061161a82610ed0565b6116276040519182610223565b8281528092611638601f1991610ed0565b0190602036910137565b9061165961165f9392611653611559565b50611f9d565b906119d5565b611667611559565b9080511561185f57610c1681607f6116ab6116a561169861168a6116b697610915565b516001600160f81b03191690565b6001600160f81b03191690565b60f81c90565b111561184d57610eaa565b6116c360048251146108ce565b6116de6116d86116d283610915565b51611863565b15158352565b6116ea6114d482610927565b906020918284015261171d61171761170a61170484610937565b5161191f565b9260409384870152610958565b51610ee7565b9182518061172d575b5050505090565b61173a90959293956115c0565b9260608501938452600092835b875181101561183d57611758611590565b6117a261171761176b611717858d610979565b61178d61178061177a83610915565b51611894565b6001600160a01b03168552565b61179961170482610937565b87850152610927565b8051806117d3575b5050906117c8816117ce938951906117c28383610979565b52610979565b506108eb565b611747565b6117e7909a9495989993979296919a611610565b97818701988952825b8b51811015611826578061181061180a611821938f610979565b516118b3565b61181b828d51610979565b526108eb565b6117f0565b50939950919796509490939192916117c8826117aa565b5095505050505038808080611726565b60016000198251019101908152610eaa565b5090565b600181510361013357602001515160001a8015908115611889575b501561122857600090565b60809150143861187e565b6015815103610133576001600160a01b03906118af906114eb565b1690565b60218151036101335760200151600181018091116108fa575190565b604051602081018181106001600160401b038211176102035760405260008152906000368137565b9061190182610253565b61190e6040519182610223565b8281528092611638601f1991610253565b80511561013357611932610af09161152d565b61193e819392936118f7565b928360200190611958565b601f81116108fa576101000a90565b9290919283156119cf5792915b60209384841061199a57805182528481018091116108fa579381018091116108fa5791601f1981019081116108fa5791611965565b91935091806119a857505050565b6119bc6119b76119c192611070565b611949565b61107e565b905182518216911916179052565b50915050565b906119e1606091611e70565b9060009283906119ef610e91565b50855115611c5c57916000959195925b8251841015611c525783158080611c3c575b610133571580611c20575b61013357611a2d6117178585610979565b96875160028114600014611b2f575050611a51611a4c61170489610915565b611cf6565b96611a67611a60898984611f11565b8092610b50565b975111611b105715611a9f575050611a7f905161107e565b11610133575111611a9657611704610af091610927565b50610af06118cf565b90929195611aad875161107e565b83146101335780611ad0611acc611ac6611aed94610927565b51610fba565b1590565b15611af757611ae1611ae791610927565b51611c96565b926108eb565b92909591956119ff565b611b03611ae791610927565b5160208101519051902090565b50505091509250611b2291505161107e565b1161013357610af06118cf565b601190989194989592939514611b4a575b50611aed906108eb565b855191979695949350908514611bfe5760ff611b75611b6f6116a561168a898961157f565b96610b42565b951690601082101561013357611b94611b8e8383610979565b51611cdf565b15611bb657505050505050611ba9905161107e565b0361013357610af06118cf565b9081611bd1611acc611ac684611aed969c98999a9b9c610979565b15611beb57611be391611ae191610979565b925b90611b40565b611bf891611b0391610979565b92611be5565b959492505050611c0f91505161107e565b0361013357611704610af091610968565b50611c34611c2e8585610979565b51611ca4565b871415611a1c565b50611c4a611b038686610979565b821415611a11565b5095945050505050565b92505050611c8e92507f56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b4219150146108ce565b610af06118cf565b611c9f9061152d565b902090565b80516020811015611cbc575060208101519051902090565b9060200151206040516020810191825260208152611cd9816101e8565b51902090565b6001815103610fdb57602001515160001a60801490565b600091815115611e3957611d0982610915565b5160fc1c80611e025750600291835b611d4284611d318451611d2c8115156108ce565b611e3d565b611d3d818311156108ce565b6110ab565b92611d4c846118f7565b958594835b611d5b8289610b50565b871015611dea5790611dbd611db7838b89611d5b9660018d81811615600014611dc657611d9c6116a561168a611db196611da594600f60f81b961c9061157f565b60041c600f1690565b60f81b168b1a9261157f565b53610b42565b97610b42565b96909150611d51565b611de46116a561168a611db196611da594600f60f81b961c9061157f565b60ff1690565b93505093509350611dfe9150845114611e53565b9190565b60018103611e14575060019183611d18565b60028103611e285750600291600193611d18565b600303611e39576001918293611d18565b8280fd5b908160011b91808304600214901517156108fa57565b15611e5a57565b634e487b7160e01b600052600160045260246000fd5b90611e818251611d2c8115156108ce565b91611e8b836118f7565b9260009081925b818410611ea85750506102519150835114611e53565b9091611ee9611eef91600180871615600014611ef757611ed8611d9c6116a561168a8a600f60f81b951c8961157f565b60f81b1660001a611db1828a61157f565b93610b42565b929190611e92565b611ed8611de46116a561168a8a600f60f81b951c8961157f565b919060005b8381018082116108fa578251811080611f76575b15611f6e576001600160f81b0319908190611f45908561157f565b511690611f52838661157f565b511603611f6757611f62906108eb565b611f16565b9250505090565b509250505090565b5083518210611f2a565b60405190611f8d826101e8565b60018252600160ff1b6020830152565b801561208357607f81111561205e5760ff81111561202f5761ffff80821115612001575062ffffff9081811115611fd357600080fd5b604051608360f81b6020820152911660e81b6001600160e81b0319166021820152610af08160248101610372565b604051604160f91b6020820152911660f01b6001600160f01b0319166021820152610af08160238101610372565b604051608160f81b602082015260f89190911b6001600160f81b0319166021820152610af08160228101610372565b60405160f89190911b6001600160f81b0319166020820152610af08160218101610372565b50610af0611f8056fea264697066735822122028f35b8fc2c380a7484c5b89081c34651180f2f77d351004dd44898db7a5814364736f6c63430008120033",
}

// StartVoteBountyABI is the input ABI used to generate the binding from.
// Deprecated: Use StartVoteBountyMetaData.ABI instead.
var StartVoteBountyABI = StartVoteBountyMetaData.ABI

// StartVoteBountyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StartVoteBountyMetaData.Bin instead.
var StartVoteBountyBin = StartVoteBountyMetaData.Bin

// DeployStartVoteBounty deploys a new Ethereum contract, binding an instance of StartVoteBounty to it.
func DeployStartVoteBounty(auth *bind.TransactOpts, backend bind.ContractBackend, voting common.Address) (common.Address, *types.Transaction, *StartVoteBounty, error) {
	parsed, err := StartVoteBountyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StartVoteBountyBin), backend, voting)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StartVoteBounty{StartVoteBountyCaller: StartVoteBountyCaller{contract: contract}, StartVoteBountyTransactor: StartVoteBountyTransactor{contract: contract}, StartVoteBountyFilterer: StartVoteBountyFilterer{contract: contract}}, nil
}

// StartVoteBounty is an auto generated Go binding around an Ethereum contract.
type StartVoteBounty struct {
	StartVoteBountyCaller     // Read-only binding to the contract
	StartVoteBountyTransactor // Write-only binding to the contract
	StartVoteBountyFilterer   // Log filterer for contract events
}

// StartVoteBountyCaller is an auto generated read-only Go binding around an Ethereum contract.
type StartVoteBountyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StartVoteBountyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StartVoteBountyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StartVoteBountyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StartVoteBountyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StartVoteBountySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StartVoteBountySession struct {
	Contract     *StartVoteBounty  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StartVoteBountyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StartVoteBountyCallerSession struct {
	Contract *StartVoteBountyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StartVoteBountyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StartVoteBountyTransactorSession struct {
	Contract     *StartVoteBountyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StartVoteBountyRaw is an auto generated low-level Go binding around an Ethereum contract.
type StartVoteBountyRaw struct {
	Contract *StartVoteBounty // Generic contract binding to access the raw methods on
}

// StartVoteBountyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StartVoteBountyCallerRaw struct {
	Contract *StartVoteBountyCaller // Generic read-only contract binding to access the raw methods on
}

// StartVoteBountyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StartVoteBountyTransactorRaw struct {
	Contract *StartVoteBountyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStartVoteBounty creates a new instance of StartVoteBounty, bound to a specific deployed contract.
func NewStartVoteBounty(address common.Address, backend bind.ContractBackend) (*StartVoteBounty, error) {
	contract, err := bindStartVoteBounty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StartVoteBounty{StartVoteBountyCaller: StartVoteBountyCaller{contract: contract}, StartVoteBountyTransactor: StartVoteBountyTransactor{contract: contract}, StartVoteBountyFilterer: StartVoteBountyFilterer{contract: contract}}, nil
}

// NewStartVoteBountyCaller creates a new read-only instance of StartVoteBounty, bound to a specific deployed contract.
func NewStartVoteBountyCaller(address common.Address, caller bind.ContractCaller) (*StartVoteBountyCaller, error) {
	contract, err := bindStartVoteBounty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StartVoteBountyCaller{contract: contract}, nil
}

// NewStartVoteBountyTransactor creates a new write-only instance of StartVoteBounty, bound to a specific deployed contract.
func NewStartVoteBountyTransactor(address common.Address, transactor bind.ContractTransactor) (*StartVoteBountyTransactor, error) {
	contract, err := bindStartVoteBounty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StartVoteBountyTransactor{contract: contract}, nil
}

// NewStartVoteBountyFilterer creates a new log filterer instance of StartVoteBounty, bound to a specific deployed contract.
func NewStartVoteBountyFilterer(address common.Address, filterer bind.ContractFilterer) (*StartVoteBountyFilterer, error) {
	contract, err := bindStartVoteBounty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StartVoteBountyFilterer{contract: contract}, nil
}

// bindStartVoteBounty binds a generic wrapper to an already deployed contract.
func bindStartVoteBounty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StartVoteBountyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StartVoteBounty *StartVoteBountyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StartVoteBounty.Contract.StartVoteBountyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StartVoteBounty *StartVoteBountyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.StartVoteBountyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StartVoteBounty *StartVoteBountyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.StartVoteBountyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StartVoteBounty *StartVoteBountyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StartVoteBounty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StartVoteBounty *StartVoteBountyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StartVoteBounty *StartVoteBountyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.contract.Transact(opts, method, params...)
}

// OPENBOUNTYCOST is a free data retrieval call binding the contract method 0xa24dc3cb.
//
// Solidity: function OPEN_BOUNTY_COST() view returns(uint256)
func (_StartVoteBounty *StartVoteBountyCaller) OPENBOUNTYCOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StartVoteBounty.contract.Call(opts, &out, "OPEN_BOUNTY_COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPENBOUNTYCOST is a free data retrieval call binding the contract method 0xa24dc3cb.
//
// Solidity: function OPEN_BOUNTY_COST() view returns(uint256)
func (_StartVoteBounty *StartVoteBountySession) OPENBOUNTYCOST() (*big.Int, error) {
	return _StartVoteBounty.Contract.OPENBOUNTYCOST(&_StartVoteBounty.CallOpts)
}

// OPENBOUNTYCOST is a free data retrieval call binding the contract method 0xa24dc3cb.
//
// Solidity: function OPEN_BOUNTY_COST() view returns(uint256)
func (_StartVoteBounty *StartVoteBountyCallerSession) OPENBOUNTYCOST() (*big.Int, error) {
	return _StartVoteBounty.Contract.OPENBOUNTYCOST(&_StartVoteBounty.CallOpts)
}

// VOTING is a free data retrieval call binding the contract method 0x269e1d1a.
//
// Solidity: function VOTING() view returns(address)
func (_StartVoteBounty *StartVoteBountyCaller) VOTING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StartVoteBounty.contract.Call(opts, &out, "VOTING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VOTING is a free data retrieval call binding the contract method 0x269e1d1a.
//
// Solidity: function VOTING() view returns(address)
func (_StartVoteBounty *StartVoteBountySession) VOTING() (common.Address, error) {
	return _StartVoteBounty.Contract.VOTING(&_StartVoteBounty.CallOpts)
}

// VOTING is a free data retrieval call binding the contract method 0x269e1d1a.
//
// Solidity: function VOTING() view returns(address)
func (_StartVoteBounty *StartVoteBountyCallerSession) VOTING() (common.Address, error) {
	return _StartVoteBounty.Contract.VOTING(&_StartVoteBounty.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x17bf72c6.
//
// Solidity: function cache(uint256 ) view returns(bytes32)
func (_StartVoteBounty *StartVoteBountyCaller) Cache(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StartVoteBounty.contract.Call(opts, &out, "cache", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Cache is a free data retrieval call binding the contract method 0x17bf72c6.
//
// Solidity: function cache(uint256 ) view returns(bytes32)
func (_StartVoteBounty *StartVoteBountySession) Cache(arg0 *big.Int) ([32]byte, error) {
	return _StartVoteBounty.Contract.Cache(&_StartVoteBounty.CallOpts, arg0)
}

// Cache is a free data retrieval call binding the contract method 0x17bf72c6.
//
// Solidity: function cache(uint256 ) view returns(bytes32)
func (_StartVoteBounty *StartVoteBountyCallerSession) Cache(arg0 *big.Int) ([32]byte, error) {
	return _StartVoteBounty.Contract.Cache(&_StartVoteBounty.CallOpts, arg0)
}

// CalculateIdentifier is a free data retrieval call binding the contract method 0x1ce63f1a.
//
// Solidity: function calculateIdentifier(address creator, uint256 startTime, address rewardToken, string metadata, bytes script) pure returns(bytes32)
func (_StartVoteBounty *StartVoteBountyCaller) CalculateIdentifier(opts *bind.CallOpts, creator common.Address, startTime *big.Int, rewardToken common.Address, metadata string, script []byte) ([32]byte, error) {
	var out []interface{}
	err := _StartVoteBounty.contract.Call(opts, &out, "calculateIdentifier", creator, startTime, rewardToken, metadata, script)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateIdentifier is a free data retrieval call binding the contract method 0x1ce63f1a.
//
// Solidity: function calculateIdentifier(address creator, uint256 startTime, address rewardToken, string metadata, bytes script) pure returns(bytes32)
func (_StartVoteBounty *StartVoteBountySession) CalculateIdentifier(creator common.Address, startTime *big.Int, rewardToken common.Address, metadata string, script []byte) ([32]byte, error) {
	return _StartVoteBounty.Contract.CalculateIdentifier(&_StartVoteBounty.CallOpts, creator, startTime, rewardToken, metadata, script)
}

// CalculateIdentifier is a free data retrieval call binding the contract method 0x1ce63f1a.
//
// Solidity: function calculateIdentifier(address creator, uint256 startTime, address rewardToken, string metadata, bytes script) pure returns(bytes32)
func (_StartVoteBounty *StartVoteBountyCallerSession) CalculateIdentifier(creator common.Address, startTime *big.Int, rewardToken common.Address, metadata string, script []byte) ([32]byte, error) {
	return _StartVoteBounty.Contract.CalculateIdentifier(&_StartVoteBounty.CallOpts, creator, startTime, rewardToken, metadata, script)
}

// GetCommitCloseBountyBlockNumber is a free data retrieval call binding the contract method 0x06239cc7.
//
// Solidity: function getCommitCloseBountyBlockNumber(bytes32 ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountyCaller) GetCommitCloseBountyBlockNumber(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StartVoteBounty.contract.Call(opts, &out, "getCommitCloseBountyBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommitCloseBountyBlockNumber is a free data retrieval call binding the contract method 0x06239cc7.
//
// Solidity: function getCommitCloseBountyBlockNumber(bytes32 ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountySession) GetCommitCloseBountyBlockNumber(arg0 [32]byte) (*big.Int, error) {
	return _StartVoteBounty.Contract.GetCommitCloseBountyBlockNumber(&_StartVoteBounty.CallOpts, arg0)
}

// GetCommitCloseBountyBlockNumber is a free data retrieval call binding the contract method 0x06239cc7.
//
// Solidity: function getCommitCloseBountyBlockNumber(bytes32 ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountyCallerSession) GetCommitCloseBountyBlockNumber(arg0 [32]byte) (*big.Int, error) {
	return _StartVoteBounty.Contract.GetCommitCloseBountyBlockNumber(&_StartVoteBounty.CallOpts, arg0)
}

// GetRefundAmount is a free data retrieval call binding the contract method 0x5d306f34.
//
// Solidity: function getRefundAmount(address ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountyCaller) GetRefundAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StartVoteBounty.contract.Call(opts, &out, "getRefundAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRefundAmount is a free data retrieval call binding the contract method 0x5d306f34.
//
// Solidity: function getRefundAmount(address ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountySession) GetRefundAmount(arg0 common.Address) (*big.Int, error) {
	return _StartVoteBounty.Contract.GetRefundAmount(&_StartVoteBounty.CallOpts, arg0)
}

// GetRefundAmount is a free data retrieval call binding the contract method 0x5d306f34.
//
// Solidity: function getRefundAmount(address ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountyCallerSession) GetRefundAmount(arg0 common.Address) (*big.Int, error) {
	return _StartVoteBounty.Contract.GetRefundAmount(&_StartVoteBounty.CallOpts, arg0)
}

// GetRewardAmount is a free data retrieval call binding the contract method 0x1af57b96.
//
// Solidity: function getRewardAmount(bytes32 ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountyCaller) GetRewardAmount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StartVoteBounty.contract.Call(opts, &out, "getRewardAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardAmount is a free data retrieval call binding the contract method 0x1af57b96.
//
// Solidity: function getRewardAmount(bytes32 ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountySession) GetRewardAmount(arg0 [32]byte) (*big.Int, error) {
	return _StartVoteBounty.Contract.GetRewardAmount(&_StartVoteBounty.CallOpts, arg0)
}

// GetRewardAmount is a free data retrieval call binding the contract method 0x1af57b96.
//
// Solidity: function getRewardAmount(bytes32 ) view returns(uint256)
func (_StartVoteBounty *StartVoteBountyCallerSession) GetRewardAmount(arg0 [32]byte) (*big.Int, error) {
	return _StartVoteBounty.Contract.GetRewardAmount(&_StartVoteBounty.CallOpts, arg0)
}

// ApplyCloseBounty is a paid mutator transaction binding the contract method 0xc2084b04.
//
// Solidity: function applyCloseBounty(uint256 startTime, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) ApplyCloseBounty(opts *bind.TransactOpts, startTime *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "applyCloseBounty", startTime, rewardToken, digest)
}

// ApplyCloseBounty is a paid mutator transaction binding the contract method 0xc2084b04.
//
// Solidity: function applyCloseBounty(uint256 startTime, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountySession) ApplyCloseBounty(startTime *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.ApplyCloseBounty(&_StartVoteBounty.TransactOpts, startTime, rewardToken, digest)
}

// ApplyCloseBounty is a paid mutator transaction binding the contract method 0xc2084b04.
//
// Solidity: function applyCloseBounty(uint256 startTime, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) ApplyCloseBounty(startTime *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.ApplyCloseBounty(&_StartVoteBounty.TransactOpts, startTime, rewardToken, digest)
}

// CacheBlockHash is a paid mutator transaction binding the contract method 0x6f0d0a5d.
//
// Solidity: function cacheBlockHash(uint256 blockNumber) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) CacheBlockHash(opts *bind.TransactOpts, blockNumber *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "cacheBlockHash", blockNumber)
}

// CacheBlockHash is a paid mutator transaction binding the contract method 0x6f0d0a5d.
//
// Solidity: function cacheBlockHash(uint256 blockNumber) returns()
func (_StartVoteBounty *StartVoteBountySession) CacheBlockHash(blockNumber *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.CacheBlockHash(&_StartVoteBounty.TransactOpts, blockNumber)
}

// CacheBlockHash is a paid mutator transaction binding the contract method 0x6f0d0a5d.
//
// Solidity: function cacheBlockHash(uint256 blockNumber) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) CacheBlockHash(blockNumber *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.CacheBlockHash(&_StartVoteBounty.TransactOpts, blockNumber)
}

// ClaimBounty is a paid mutator transaction binding the contract method 0x2fa970d8.
//
// Solidity: function claimBounty(address creator, uint256 startTime, address rewardToken, bytes32 digest, uint256 receiptIndex, bytes headerRlpBytes, bytes proofRlpBytes) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) ClaimBounty(opts *bind.TransactOpts, creator common.Address, startTime *big.Int, rewardToken common.Address, digest [32]byte, receiptIndex *big.Int, headerRlpBytes []byte, proofRlpBytes []byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "claimBounty", creator, startTime, rewardToken, digest, receiptIndex, headerRlpBytes, proofRlpBytes)
}

// ClaimBounty is a paid mutator transaction binding the contract method 0x2fa970d8.
//
// Solidity: function claimBounty(address creator, uint256 startTime, address rewardToken, bytes32 digest, uint256 receiptIndex, bytes headerRlpBytes, bytes proofRlpBytes) returns()
func (_StartVoteBounty *StartVoteBountySession) ClaimBounty(creator common.Address, startTime *big.Int, rewardToken common.Address, digest [32]byte, receiptIndex *big.Int, headerRlpBytes []byte, proofRlpBytes []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.ClaimBounty(&_StartVoteBounty.TransactOpts, creator, startTime, rewardToken, digest, receiptIndex, headerRlpBytes, proofRlpBytes)
}

// ClaimBounty is a paid mutator transaction binding the contract method 0x2fa970d8.
//
// Solidity: function claimBounty(address creator, uint256 startTime, address rewardToken, bytes32 digest, uint256 receiptIndex, bytes headerRlpBytes, bytes proofRlpBytes) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) ClaimBounty(creator common.Address, startTime *big.Int, rewardToken common.Address, digest [32]byte, receiptIndex *big.Int, headerRlpBytes []byte, proofRlpBytes []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.ClaimBounty(&_StartVoteBounty.TransactOpts, creator, startTime, rewardToken, digest, receiptIndex, headerRlpBytes, proofRlpBytes)
}

// CommitCloseBounty is a paid mutator transaction binding the contract method 0xad22eb62.
//
// Solidity: function commitCloseBounty(uint256 startTime, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) CommitCloseBounty(opts *bind.TransactOpts, startTime *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "commitCloseBounty", startTime, rewardToken, digest)
}

// CommitCloseBounty is a paid mutator transaction binding the contract method 0xad22eb62.
//
// Solidity: function commitCloseBounty(uint256 startTime, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountySession) CommitCloseBounty(startTime *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.CommitCloseBounty(&_StartVoteBounty.TransactOpts, startTime, rewardToken, digest)
}

// CommitCloseBounty is a paid mutator transaction binding the contract method 0xad22eb62.
//
// Solidity: function commitCloseBounty(uint256 startTime, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) CommitCloseBounty(startTime *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.CommitCloseBounty(&_StartVoteBounty.TransactOpts, startTime, rewardToken, digest)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0xf5ba9296.
//
// Solidity: function increaseBounty(uint256 startTime, address rewardToken, bytes32 digest, uint256 increaseAmount) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) IncreaseBounty(opts *bind.TransactOpts, startTime *big.Int, rewardToken common.Address, digest [32]byte, increaseAmount *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "increaseBounty", startTime, rewardToken, digest, increaseAmount)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0xf5ba9296.
//
// Solidity: function increaseBounty(uint256 startTime, address rewardToken, bytes32 digest, uint256 increaseAmount) returns()
func (_StartVoteBounty *StartVoteBountySession) IncreaseBounty(startTime *big.Int, rewardToken common.Address, digest [32]byte, increaseAmount *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.IncreaseBounty(&_StartVoteBounty.TransactOpts, startTime, rewardToken, digest, increaseAmount)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0xf5ba9296.
//
// Solidity: function increaseBounty(uint256 startTime, address rewardToken, bytes32 digest, uint256 increaseAmount) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) IncreaseBounty(startTime *big.Int, rewardToken common.Address, digest [32]byte, increaseAmount *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.IncreaseBounty(&_StartVoteBounty.TransactOpts, startTime, rewardToken, digest, increaseAmount)
}

// OpenBounty is a paid mutator transaction binding the contract method 0xbee8a902.
//
// Solidity: function openBounty(uint256 startTime, address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountyTransactor) OpenBounty(opts *bind.TransactOpts, startTime *big.Int, rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "openBounty", startTime, rewardToken, rewardAmount, metadata, script)
}

// OpenBounty is a paid mutator transaction binding the contract method 0xbee8a902.
//
// Solidity: function openBounty(uint256 startTime, address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountySession) OpenBounty(startTime *big.Int, rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.OpenBounty(&_StartVoteBounty.TransactOpts, startTime, rewardToken, rewardAmount, metadata, script)
}

// OpenBounty is a paid mutator transaction binding the contract method 0xbee8a902.
//
// Solidity: function openBounty(uint256 startTime, address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountyTransactorSession) OpenBounty(startTime *big.Int, rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.OpenBounty(&_StartVoteBounty.TransactOpts, startTime, rewardToken, rewardAmount, metadata, script)
}

// OpenBounty0 is a paid mutator transaction binding the contract method 0xd6841c94.
//
// Solidity: function openBounty(address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountyTransactor) OpenBounty0(opts *bind.TransactOpts, rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "openBounty0", rewardToken, rewardAmount, metadata, script)
}

// OpenBounty0 is a paid mutator transaction binding the contract method 0xd6841c94.
//
// Solidity: function openBounty(address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountySession) OpenBounty0(rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.OpenBounty0(&_StartVoteBounty.TransactOpts, rewardToken, rewardAmount, metadata, script)
}

// OpenBounty0 is a paid mutator transaction binding the contract method 0xd6841c94.
//
// Solidity: function openBounty(address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountyTransactorSession) OpenBounty0(rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.OpenBounty0(&_StartVoteBounty.TransactOpts, rewardToken, rewardAmount, metadata, script)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address creator) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) Refund(opts *bind.TransactOpts, creator common.Address) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "refund", creator)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address creator) returns()
func (_StartVoteBounty *StartVoteBountySession) Refund(creator common.Address) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.Refund(&_StartVoteBounty.TransactOpts, creator)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address creator) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) Refund(creator common.Address) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.Refund(&_StartVoteBounty.TransactOpts, creator)
}

// StartVoteBountyApplyCloseBountyIterator is returned from FilterApplyCloseBounty and is used to iterate over the raw logs and unpacked data for ApplyCloseBounty events raised by the StartVoteBounty contract.
type StartVoteBountyApplyCloseBountyIterator struct {
	Event *StartVoteBountyApplyCloseBounty // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StartVoteBountyApplyCloseBountyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StartVoteBountyApplyCloseBounty)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StartVoteBountyApplyCloseBounty)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StartVoteBountyApplyCloseBountyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StartVoteBountyApplyCloseBountyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StartVoteBountyApplyCloseBounty represents a ApplyCloseBounty event raised by the StartVoteBounty contract.
type StartVoteBountyApplyCloseBounty struct {
	Identifier [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApplyCloseBounty is a free log retrieval operation binding the contract event 0xc236bdbce1e4cfdbe02cc5976c05c076ceedc985cda30253cca17a15cd8a01b5.
//
// Solidity: event ApplyCloseBounty(bytes32 indexed identifier)
func (_StartVoteBounty *StartVoteBountyFilterer) FilterApplyCloseBounty(opts *bind.FilterOpts, identifier [][32]byte) (*StartVoteBountyApplyCloseBountyIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _StartVoteBounty.contract.FilterLogs(opts, "ApplyCloseBounty", identifierRule)
	if err != nil {
		return nil, err
	}
	return &StartVoteBountyApplyCloseBountyIterator{contract: _StartVoteBounty.contract, event: "ApplyCloseBounty", logs: logs, sub: sub}, nil
}

// WatchApplyCloseBounty is a free log subscription operation binding the contract event 0xc236bdbce1e4cfdbe02cc5976c05c076ceedc985cda30253cca17a15cd8a01b5.
//
// Solidity: event ApplyCloseBounty(bytes32 indexed identifier)
func (_StartVoteBounty *StartVoteBountyFilterer) WatchApplyCloseBounty(opts *bind.WatchOpts, sink chan<- *StartVoteBountyApplyCloseBounty, identifier [][32]byte) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _StartVoteBounty.contract.WatchLogs(opts, "ApplyCloseBounty", identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StartVoteBountyApplyCloseBounty)
				if err := _StartVoteBounty.contract.UnpackLog(event, "ApplyCloseBounty", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApplyCloseBounty is a log parse operation binding the contract event 0xc236bdbce1e4cfdbe02cc5976c05c076ceedc985cda30253cca17a15cd8a01b5.
//
// Solidity: event ApplyCloseBounty(bytes32 indexed identifier)
func (_StartVoteBounty *StartVoteBountyFilterer) ParseApplyCloseBounty(log types.Log) (*StartVoteBountyApplyCloseBounty, error) {
	event := new(StartVoteBountyApplyCloseBounty)
	if err := _StartVoteBounty.contract.UnpackLog(event, "ApplyCloseBounty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StartVoteBountyClaimBountyIterator is returned from FilterClaimBounty and is used to iterate over the raw logs and unpacked data for ClaimBounty events raised by the StartVoteBounty contract.
type StartVoteBountyClaimBountyIterator struct {
	Event *StartVoteBountyClaimBounty // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StartVoteBountyClaimBountyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StartVoteBountyClaimBounty)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StartVoteBountyClaimBounty)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StartVoteBountyClaimBountyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StartVoteBountyClaimBountyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StartVoteBountyClaimBounty represents a ClaimBounty event raised by the StartVoteBounty contract.
type StartVoteBountyClaimBounty struct {
	Identifier [32]byte
	Claimant   common.Address
	VoteId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimBounty is a free log retrieval operation binding the contract event 0x62f6344c3fb7a53546ee98642ba528b6b561027f33810bac1cdd4032916fa75d.
//
// Solidity: event ClaimBounty(bytes32 indexed identifier, address indexed claimant, uint256 indexed voteId)
func (_StartVoteBounty *StartVoteBountyFilterer) FilterClaimBounty(opts *bind.FilterOpts, identifier [][32]byte, claimant []common.Address, voteId []*big.Int) (*StartVoteBountyClaimBountyIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}
	var voteIdRule []interface{}
	for _, voteIdItem := range voteId {
		voteIdRule = append(voteIdRule, voteIdItem)
	}

	logs, sub, err := _StartVoteBounty.contract.FilterLogs(opts, "ClaimBounty", identifierRule, claimantRule, voteIdRule)
	if err != nil {
		return nil, err
	}
	return &StartVoteBountyClaimBountyIterator{contract: _StartVoteBounty.contract, event: "ClaimBounty", logs: logs, sub: sub}, nil
}

// WatchClaimBounty is a free log subscription operation binding the contract event 0x62f6344c3fb7a53546ee98642ba528b6b561027f33810bac1cdd4032916fa75d.
//
// Solidity: event ClaimBounty(bytes32 indexed identifier, address indexed claimant, uint256 indexed voteId)
func (_StartVoteBounty *StartVoteBountyFilterer) WatchClaimBounty(opts *bind.WatchOpts, sink chan<- *StartVoteBountyClaimBounty, identifier [][32]byte, claimant []common.Address, voteId []*big.Int) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}
	var voteIdRule []interface{}
	for _, voteIdItem := range voteId {
		voteIdRule = append(voteIdRule, voteIdItem)
	}

	logs, sub, err := _StartVoteBounty.contract.WatchLogs(opts, "ClaimBounty", identifierRule, claimantRule, voteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StartVoteBountyClaimBounty)
				if err := _StartVoteBounty.contract.UnpackLog(event, "ClaimBounty", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimBounty is a log parse operation binding the contract event 0x62f6344c3fb7a53546ee98642ba528b6b561027f33810bac1cdd4032916fa75d.
//
// Solidity: event ClaimBounty(bytes32 indexed identifier, address indexed claimant, uint256 indexed voteId)
func (_StartVoteBounty *StartVoteBountyFilterer) ParseClaimBounty(log types.Log) (*StartVoteBountyClaimBounty, error) {
	event := new(StartVoteBountyClaimBounty)
	if err := _StartVoteBounty.contract.UnpackLog(event, "ClaimBounty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StartVoteBountyCommitCloseBountyIterator is returned from FilterCommitCloseBounty and is used to iterate over the raw logs and unpacked data for CommitCloseBounty events raised by the StartVoteBounty contract.
type StartVoteBountyCommitCloseBountyIterator struct {
	Event *StartVoteBountyCommitCloseBounty // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StartVoteBountyCommitCloseBountyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StartVoteBountyCommitCloseBounty)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StartVoteBountyCommitCloseBounty)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StartVoteBountyCommitCloseBountyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StartVoteBountyCommitCloseBountyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StartVoteBountyCommitCloseBounty represents a CommitCloseBounty event raised by the StartVoteBounty contract.
type StartVoteBountyCommitCloseBounty struct {
	Identifier [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitCloseBounty is a free log retrieval operation binding the contract event 0xd994ec1efb74ba41455cf04e5da2d899411306f2fa57dec2a98fcf3ce0bb3063.
//
// Solidity: event CommitCloseBounty(bytes32 indexed identifier)
func (_StartVoteBounty *StartVoteBountyFilterer) FilterCommitCloseBounty(opts *bind.FilterOpts, identifier [][32]byte) (*StartVoteBountyCommitCloseBountyIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _StartVoteBounty.contract.FilterLogs(opts, "CommitCloseBounty", identifierRule)
	if err != nil {
		return nil, err
	}
	return &StartVoteBountyCommitCloseBountyIterator{contract: _StartVoteBounty.contract, event: "CommitCloseBounty", logs: logs, sub: sub}, nil
}

// WatchCommitCloseBounty is a free log subscription operation binding the contract event 0xd994ec1efb74ba41455cf04e5da2d899411306f2fa57dec2a98fcf3ce0bb3063.
//
// Solidity: event CommitCloseBounty(bytes32 indexed identifier)
func (_StartVoteBounty *StartVoteBountyFilterer) WatchCommitCloseBounty(opts *bind.WatchOpts, sink chan<- *StartVoteBountyCommitCloseBounty, identifier [][32]byte) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _StartVoteBounty.contract.WatchLogs(opts, "CommitCloseBounty", identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StartVoteBountyCommitCloseBounty)
				if err := _StartVoteBounty.contract.UnpackLog(event, "CommitCloseBounty", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitCloseBounty is a log parse operation binding the contract event 0xd994ec1efb74ba41455cf04e5da2d899411306f2fa57dec2a98fcf3ce0bb3063.
//
// Solidity: event CommitCloseBounty(bytes32 indexed identifier)
func (_StartVoteBounty *StartVoteBountyFilterer) ParseCommitCloseBounty(log types.Log) (*StartVoteBountyCommitCloseBounty, error) {
	event := new(StartVoteBountyCommitCloseBounty)
	if err := _StartVoteBounty.contract.UnpackLog(event, "CommitCloseBounty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StartVoteBountyIncreaseBountyIterator is returned from FilterIncreaseBounty and is used to iterate over the raw logs and unpacked data for IncreaseBounty events raised by the StartVoteBounty contract.
type StartVoteBountyIncreaseBountyIterator struct {
	Event *StartVoteBountyIncreaseBounty // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StartVoteBountyIncreaseBountyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StartVoteBountyIncreaseBounty)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StartVoteBountyIncreaseBounty)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StartVoteBountyIncreaseBountyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StartVoteBountyIncreaseBountyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StartVoteBountyIncreaseBounty represents a IncreaseBounty event raised by the StartVoteBounty contract.
type StartVoteBountyIncreaseBounty struct {
	Identifier  [32]byte
	AddedAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIncreaseBounty is a free log retrieval operation binding the contract event 0x18ee4fa302d6692472b161e97fddc236d75cf9f3a31fe117ed398dac7d0ceccd.
//
// Solidity: event IncreaseBounty(bytes32 indexed identifier, uint256 addedAmount)
func (_StartVoteBounty *StartVoteBountyFilterer) FilterIncreaseBounty(opts *bind.FilterOpts, identifier [][32]byte) (*StartVoteBountyIncreaseBountyIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _StartVoteBounty.contract.FilterLogs(opts, "IncreaseBounty", identifierRule)
	if err != nil {
		return nil, err
	}
	return &StartVoteBountyIncreaseBountyIterator{contract: _StartVoteBounty.contract, event: "IncreaseBounty", logs: logs, sub: sub}, nil
}

// WatchIncreaseBounty is a free log subscription operation binding the contract event 0x18ee4fa302d6692472b161e97fddc236d75cf9f3a31fe117ed398dac7d0ceccd.
//
// Solidity: event IncreaseBounty(bytes32 indexed identifier, uint256 addedAmount)
func (_StartVoteBounty *StartVoteBountyFilterer) WatchIncreaseBounty(opts *bind.WatchOpts, sink chan<- *StartVoteBountyIncreaseBounty, identifier [][32]byte) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _StartVoteBounty.contract.WatchLogs(opts, "IncreaseBounty", identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StartVoteBountyIncreaseBounty)
				if err := _StartVoteBounty.contract.UnpackLog(event, "IncreaseBounty", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseBounty is a log parse operation binding the contract event 0x18ee4fa302d6692472b161e97fddc236d75cf9f3a31fe117ed398dac7d0ceccd.
//
// Solidity: event IncreaseBounty(bytes32 indexed identifier, uint256 addedAmount)
func (_StartVoteBounty *StartVoteBountyFilterer) ParseIncreaseBounty(log types.Log) (*StartVoteBountyIncreaseBounty, error) {
	event := new(StartVoteBountyIncreaseBounty)
	if err := _StartVoteBounty.contract.UnpackLog(event, "IncreaseBounty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StartVoteBountyOpenBountyIterator is returned from FilterOpenBounty and is used to iterate over the raw logs and unpacked data for OpenBounty events raised by the StartVoteBounty contract.
type StartVoteBountyOpenBountyIterator struct {
	Event *StartVoteBountyOpenBounty // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StartVoteBountyOpenBountyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StartVoteBountyOpenBounty)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StartVoteBountyOpenBounty)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StartVoteBountyOpenBountyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StartVoteBountyOpenBountyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StartVoteBountyOpenBounty represents a OpenBounty event raised by the StartVoteBounty contract.
type StartVoteBountyOpenBounty struct {
	Identifier   [32]byte
	Creator      common.Address
	RewardToken  common.Address
	RewardAmount *big.Int
	StartTime    *big.Int
	Metadata     string
	Script       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOpenBounty is a free log retrieval operation binding the contract event 0x61c6abb5ad2430a2a28d7fe158faa33e09784bd8ce4660816300feea5d448a59.
//
// Solidity: event OpenBounty(bytes32 indexed identifier, address indexed creator, address indexed rewardToken, uint256 rewardAmount, uint256 startTime, string metadata, bytes script)
func (_StartVoteBounty *StartVoteBountyFilterer) FilterOpenBounty(opts *bind.FilterOpts, identifier [][32]byte, creator []common.Address, rewardToken []common.Address) (*StartVoteBountyOpenBountyIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _StartVoteBounty.contract.FilterLogs(opts, "OpenBounty", identifierRule, creatorRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &StartVoteBountyOpenBountyIterator{contract: _StartVoteBounty.contract, event: "OpenBounty", logs: logs, sub: sub}, nil
}

// WatchOpenBounty is a free log subscription operation binding the contract event 0x61c6abb5ad2430a2a28d7fe158faa33e09784bd8ce4660816300feea5d448a59.
//
// Solidity: event OpenBounty(bytes32 indexed identifier, address indexed creator, address indexed rewardToken, uint256 rewardAmount, uint256 startTime, string metadata, bytes script)
func (_StartVoteBounty *StartVoteBountyFilterer) WatchOpenBounty(opts *bind.WatchOpts, sink chan<- *StartVoteBountyOpenBounty, identifier [][32]byte, creator []common.Address, rewardToken []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _StartVoteBounty.contract.WatchLogs(opts, "OpenBounty", identifierRule, creatorRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StartVoteBountyOpenBounty)
				if err := _StartVoteBounty.contract.UnpackLog(event, "OpenBounty", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpenBounty is a log parse operation binding the contract event 0x61c6abb5ad2430a2a28d7fe158faa33e09784bd8ce4660816300feea5d448a59.
//
// Solidity: event OpenBounty(bytes32 indexed identifier, address indexed creator, address indexed rewardToken, uint256 rewardAmount, uint256 startTime, string metadata, bytes script)
func (_StartVoteBounty *StartVoteBountyFilterer) ParseOpenBounty(log types.Log) (*StartVoteBountyOpenBounty, error) {
	event := new(StartVoteBountyOpenBounty)
	if err := _StartVoteBounty.contract.UnpackLog(event, "OpenBounty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
