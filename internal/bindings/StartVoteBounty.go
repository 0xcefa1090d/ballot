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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voting\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"ApplyCloseBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"voteId\",\"type\":\"uint256\"}],\"name\":\"ClaimBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"CommitCloseBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addedAmount\",\"type\":\"uint256\"}],\"name\":\"IncreaseBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"}],\"name\":\"OpenBounty\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OPEN_BOUNTY_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"applyCloseBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cache\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"cacheBlockHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"}],\"name\":\"calculateIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"receiptIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"headerRlpBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofRlpBytes\",\"type\":\"bytes\"}],\"name\":\"claimBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"commitCloseBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getCommitCloseBountyBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getRefundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"increaseAmount\",\"type\":\"uint256\"}],\"name\":\"increaseBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"}],\"name\":\"openBounty\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x0x60a03461007157601f6120c938819003918201601f19168301916001600160401b038311848410176100765780849260209460405283398101031261007157516001600160a01b03811681036100715760805260405161203c908161008d823960805181818161039d0152610cd60152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306239cc7146100f757806317bf72c6146100f25780631af57b96146100ed5780631ce63f1a146100e8578063269e1d1a146100e35780632fa970d8146100de5780635d306f34146100d95780636f0d0a5d146100d4578063a24dc3cb146100cf578063ad22eb62146100ca578063c2084b04146100c5578063d6841c94146100c0578063f5ba9296146100bb5763fa89401a146100b657600080fd5b6108ce565b6107f3565b61068d565b610587565b6104f5565b6104a9565b610479565b61043f565b6103cc565b610387565b6102a5565b610154565b610128565b346101235760203660031901126101235760043560005260016020526020604060002054604051908152f35b600080fd5b346101235760203660031901126101235760043560005260006020526020604060002054604051908152f35b346101235760203660031901126101235760043560005260036020526020604060002054604051908152f35b600435906001600160a01b038216820361012357565b604435906001600160a01b038216820361012357565b602435906001600160a01b038216820361012357565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b038211176101f357604052565b6101c2565b608081019081106001600160401b038211176101f357604052565b90601f801991011681019081106001600160401b038211176101f357604052565b60405190610241826101d8565b565b6001600160401b0381116101f357601f01601f191660200190565b81601f820112156101235780359061027582610243565b926102836040519485610213565b8284526020838301011161012357816000926020809301838601378301015290565b346101235760a0366003190112610123576102be610180565b6102c6610196565b6001600160401b039190606435838111610123576102e890369060040161025e565b608435938411610123576103206103709161030a61038396369060040161025e565b9060208151910120906020815191012090610b8f565b60208151910120916103626040519384926020840196602435908890606092959493608083019660018060a01b03809316845260208401521660408201520152565b03601f198101835282610213565b5190206040519081529081906020820190565b0390f35b34610123576000366003190112610123576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101235760e0366003190112610123576103e5610180565b6103ed610196565b6001600160401b03919060a4358381116101235761040f90369060040161025e565b9060c4359384116101235761042b61043d94369060040161025e565b92608435916064359160243590610bed565b005b34610123576020366003190112610123576001600160a01b03610460610180565b1660005260026020526020604060002054604051908152f35b34610123576020366003190112610123576004358040908115610123576000526000602052604060002055600080f35b3461012357600036600319011261012357602060405166470de4df8200008152f35b606090600319011261012357600435906024356001600160a01b0381168103610123579060443590565b3461012357610539610506366104cb565b604080513360208201908152918101949094526001600160a01b0390921660608401526080830152918160a08101610362565b51902060009080825260036020526040822054156105835760016020524360408320557fd994ec1efb74ba41455cf04e5da2d899411306f2fa57dec2a98fcf3ce0bb30638280a280f35b5080fd5b346101235761036261066f6105d661059e366104cb565b604080513360208201908152918101949094526001600160a01b038316606085015260808401919091529094909290829060a0820190565b51902060009281845260036020526040842054916105f583151561095e565b61062561060c826000526001602052604060002090565b5480610617436110d3565b10908161067b575b5061095e565b8461063a826000526003602052604060002090565b557fc236bdbce1e4cfdbe02cc5976c05c076ceedc985cda30253cca17a15cd8a01b58580a233906001600160a01b031661122e565b610678336113a4565b80f35b9050610686436110e2565b113861061f565b6080366003190112610123576106a1610180565b6001600160401b0390602435604435838111610123576106c590369060040161025e565b90606435938411610123576107ce7f19f082eb5536f4081e8e296229b3b1aedefd1a248e83e7d37d176a76f92ddafd936107da9361070a61038397369060040161025e565b9166470de4df8200003414806107ea575b6107249061095e565b61073b825160208401208451602086012090610b8f565b8051602091820120604080513393810193845242918101919091526001600160a01b0384166060820152608081019190915261077a8160a08101610362565b51902096879161079e610797846000526003602052604060002090565b541561095e565b866107b3846000526003602052604060002090565b5560018060a01b031694859460405191829133968a846112fa565b0390a430903390611325565b6040519081529081906020820190565b5084151561071b565b346101235760803660031901126101235761080c6101ac565b606435801561012357604080513360208201908152600435928201929092526001600160a01b0384166060820152604435608082015261043d9391906108558160a08101610362565b51902080600052600360205261087b8360406000205461087681151561095e565b610be0565b61088f826000526003602052604060002090565b556040518381527f18ee4fa302d6692472b161e97fddc236d75cf9f3a31fe117ed398dac7d0ceccd90602090a2309033906001600160a01b0316611325565b34610123576020366003190112610123576108e7610180565b6001600160a01b03811660009081526002602052604081208054919283929182156109595783928392838093555af11561091e5780f35b60405162461bcd60e51b815260206004820152601360248201527211551217d514905394d1915497d19052531151606a1b6044820152606490fd5b505050fd5b1561012357565b634e487b7160e01b600052601160045260246000fd5b600019811461098a5760010190565b610965565b634e487b7160e01b600052603260045260246000fd5b8051156109b25760200190565b61098f565b8051600110156109b25760400190565b8051600210156109b25760600190565b8051600b10156109b2576101800190565b8051600310156109b25760800190565b8051601010156109b2576102200190565b80518210156109b25760209160051b010190565b60005b838110610a305750506000910152565b8181015183820152602001610a20565b90929192610a4d81610243565b91610a5b6040519384610213565b829482845282820111610123576020610241930190610a1d565b9060a0828203126101235781516001600160401b0381116101235782019080601f83011215610123578151610aac92602001610a40565b916020820151916040810151916080606083015192015190565b5190811515820361012357565b51906001600160401b038216820361012357565b6101408183031261012357610afb81610ac6565b92610b0860208301610ac6565b92610b1560408401610ad3565b92610b2260608201610ad3565b92610b2f60808301610ad3565b92610b3c60a08401610ad3565b9260c08101519260e08201519261010083015192610120810151906001600160401b03821161012357019080601f83011215610123578151610b8092602001610a40565b90565b6040513d6000823e3d90fd5b9190604051926020840152604083015260408252606082018281106001600160401b038211176101f357604052565b9066470de4df820000820180921161098a57565b906001820180921161098a57565b9190820180921161098a57565b604080516001600160a01b038381166020808401918252838501879052918716606084015260808084018990528352939a999792949093909291610c3260a082610213565b51902096610c4a886000526003602052604060002090565b5498610c578a151561095e565b610c60906113e0565b80518587830151928340831493841597610c93610cab96610ca696610c9e94610cb19c610f02575b50509d9a9b9d61095e565b60608301511161095e565b015192610f3a565b610f77565b916115bc565b97610cc4610cbf8a51151590565b61095e565b6060600099019160018060a01b0395867f000000000000000000000000000000000000000000000000000000000000000016915b845180518d101561012357610d2e610d22610d148f8794610a09565b51516001600160a01b031690565b6001600160a01b031690565b03610ef3577f0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394610d6b87610d638f8951610a09565b5101516109a5565b5103610ef3579084939291610db487610dac8f999885610d8f8c610da09351610a09565b510151838082518301019101610a75565b50505050998851610a09565b5101516109b7565b5182516305a55c1f60e41b815260048101829052909790600081602481885afa8015610eee578392610dfc92600092610ec1575b508981519101209089815191012090610b8f565b87815191012003610ea957505050610e7992610e646102419a9b610d2294610e52610e6c956000610e378e6000526003602052604060002090565b556001600160a01b0316600090815260026020526040902090565b610e5c8154610bbe565b905551610a09565b5101516109c7565b516001600160a01b031690565b93838516907f62f6344c3fb7a53546ee98642ba528b6b561027f33810bac1cdd4032916fa75d600080a41661122e565b919b610eb691965061097b565b9a9192939490610cf8565b610edd91923d8091833e610ed58183610213565b810190610ae7565b985050505050505050509038610de8565b610b83565b90949392919a610eb69061097b565b610f189192506000526000602052604060002090565b54143880610c88565b60405190610f2e826101d8565b60006020838281520152565b610f42610f21565b50602081519160405192610f55846101d8565b835201602082015290565b6001600160401b0381116101f35760051b60200190565b610f808161104a565b1561012357610f8e81611071565b610f9781610f60565b91610fa56040519384610213565b818352601f19610fb483610f60565b0160005b818110611033575050610fd9602080920151610fd3816111ce565b90610be0565b6000905b838210610feb575050505090565b61102781610ffb61102d93611148565b90611004610234565b8281528187820152611016868a610a09565b526110218589610a09565b50610be0565b9161097b565b90610fdd565b60209061103e610f21565b82828801015201610fb8565b80511561106b57602060c09101515160001a1061106657600190565b600090565b50600090565b80511561106b576000906020810190815161108b816111ce565b810180911161098a5791519051810180911161098a5791905b8281106110b15750905090565b6110ba81611148565b810180911161098a576110cd909161097b565b906110a4565b60ff1981019190821161098a57565b607f1981019190821161098a57565b60bf1981019190821161098a57565b602003906020821161098a57565b60001981019190821161098a57565b60f61981019190821161098a57565b60b61981019190821161098a57565b9190820391821161098a57565b805160001a90608082101561115e575050600190565b60b88210156111795750611174610b80916110e2565b610bd2565b9060c081101561119d5760b51991600160b783602003016101000a91015104010190565b9060f88210156111b45750611174610b80916110f1565b60010151602082900360f7016101000a90040160f5190190565b5160001a60808110156111e15750600090565b60b881108015611218575b156111f75750600190565b60c081101561120c57611174610b809161112c565b611174610b809161111d565b5060c081101580156111ec575060f881106111ec565b6044600092838093611261966040519363a9059cbb60e01b855260018060a01b0316600485015260248401525af161129f565b1561126857565b60405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b6044820152606490fd5b3d90156112cc57806020146112bd57156112b857600090565b600190565b5060206000803e600051151590565b806000803e6000fd5b906020916112ee81518092818552858086019101610a1d565b601f01601f1916010190565b9161131790610b80949284526060602085015260608401906112d5565b9160408184039101526112d5565b60009283606492611361968295604051946323b872dd60e01b865260018060a01b03809216600487015216602485015260448401525af161129f565b1561136857565b60405162461bcd60e51b81526020600482015260146024820152731514905394d1915497d19493d357d1905253115160621b6044820152606490fd5b6000808066470de4df82000081945af11561091e57565b604051906113c8826101f8565b60006060838281528260208201528260408201520152565b6113e86113bb565b506113f5610ca682610f3a565b90600b82511115610123576114086113bb565b918051600510156109b25761142060c0820151611465565b60208401528051600810156109b25761144e81611444610120611454940151611465565b60408601526109d7565b51611465565b606083015260208151910120815290565b8051801515908161149b575b501561012357611480906114a7565b9051906020811061148f575090565b6020036101000a900490565b60219150111538611471565b9060208201916114b783516111ce565b92519083820180921161098a575192830392831161098a579190565b604051906114e0826101f8565b6060808360008152600060208201528160408201520152565b9081518110156109b2570160200190565b60405190606082018281106001600160401b038211176101f3576040526060604083600081528260208201520152565b9061154482610f60565b6115516040519182610213565b8281528092611562601f1991610f60565b019060005b82811061157357505050565b60209061157e61150a565b82828501015201611567565b9061159482610f60565b6115a16040519182610213565b82815280926115b2601f1991610f60565b0190602036910137565b906115d36115d993926115cd6114d3565b50611f17565b9061194f565b6115e16114d3565b908051156117d957610ca681607f61162561161f611612611604611630976109a5565b516001600160f81b03191690565b6001600160f81b03191690565b60f81c90565b11156117c757610f3a565b61163d600482511461095e565b61165861165261164c836109a5565b516117dd565b15158352565b61166461144e826109b7565b906020918284015261169761169161168461167e846109c7565b51611899565b92604093848701526109e8565b51610f77565b918251806116a7575b5050505090565b6116b4909592939561153a565b9260608501938452600092835b87518110156117b7576116d261150a565b61171c6116916116e5611691858d610a09565b6117076116fa6116f4836109a5565b5161180e565b6001600160a01b03168552565b61171361167e826109c7565b878501526109b7565b80518061174d575b505090611742816117489389519061173c8383610a09565b52610a09565b5061097b565b6116c1565b611761909a9495989993979296919a61158a565b97818701988952825b8b518110156117a0578061178a61178461179b938f610a09565b5161182d565b611795828d51610a09565b5261097b565b61176a565b509399509197965094909391929161174282611724565b50955050505050388080806116a0565b60016000198251019101908152610f3a565b5090565b600181510361012357602001515160001a8015908115611803575b50156112b857600090565b6080915014386117f8565b6015815103610123576001600160a01b039061182990611465565b1690565b602181510361012357602001516001810180911161098a575190565b604051602081018181106001600160401b038211176101f35760405260008152906000368137565b9061187b82610243565b6118886040519182610213565b82815280926115b2601f1991610243565b805115610123576118ac610b80916114a7565b6118b881939293611871565b9283602001906118d2565b601f811161098a576101000a90565b9290919283156119495792915b602093848410611914578051825284810180911161098a5793810180911161098a5791601f19810190811161098a57916118df565b919350918061192257505050565b61193661193161193b92611100565b6118c3565b61110e565b905182518216911916179052565b50915050565b9061195b606091611dea565b906000928390611969610f21565b50855115611bd657916000959195925b8251841015611bcc5783158080611bb6575b610123571580611b9a575b610123576119a76116918585610a09565b96875160028114600014611aa95750506119cb6119c661167e896109a5565b611c70565b966119e16119da898984611e8b565b8092610be0565b975111611a8a5715611a195750506119f9905161110e565b11610123575111611a105761167e610b80916109b7565b50610b80611849565b90929195611a27875161110e565b83146101235780611a4a611a46611a40611a67946109b7565b5161104a565b1590565b15611a7157611a5b611a61916109b7565b51611c10565b9261097b565b9290959195611979565b611a7d611a61916109b7565b5160208101519051902090565b50505091509250611a9c91505161110e565b1161012357610b80611849565b601190989194989592939514611ac4575b50611a679061097b565b855191979695949350908514611b785760ff611aef611ae961161f61160489896114f9565b96610bd2565b951690601082101561012357611b0e611b088383610a09565b51611c59565b15611b3057505050505050611b23905161110e565b0361012357610b80611849565b9081611b4b611a46611a4084611a67969c98999a9b9c610a09565b15611b6557611b5d91611a5b91610a09565b925b90611aba565b611b7291611a7d91610a09565b92611b5f565b959492505050611b8991505161110e565b036101235761167e610b80916109f8565b50611bae611ba88585610a09565b51611c1e565b871415611996565b50611bc4611a7d8686610a09565b82141561198b565b5095945050505050565b92505050611c0892507f56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b42191501461095e565b610b80611849565b611c19906114a7565b902090565b80516020811015611c36575060208101519051902090565b9060200151206040516020810191825260208152611c53816101d8565b51902090565b600181510361106b57602001515160001a60801490565b600091815115611db357611c83826109a5565b5160fc1c80611d7c5750600291835b611cbc84611cab8451611ca681151561095e565b611db7565b611cb78183111561095e565b61113b565b92611cc684611871565b958594835b611cd58289610be0565b871015611d645790611d37611d31838b89611cd59660018d81811615600014611d4057611d1661161f611604611d2b96611d1f94600f60f81b961c906114f9565b60041c600f1690565b60f81b168b1a926114f9565b53610bd2565b97610bd2565b96909150611ccb565b611d5e61161f611604611d2b96611d1f94600f60f81b961c906114f9565b60ff1690565b93505093509350611d789150845114611dcd565b9190565b60018103611d8e575060019183611c92565b60028103611da25750600291600193611c92565b600303611db3576001918293611c92565b8280fd5b908160011b918083046002149015171561098a57565b15611dd457565b634e487b7160e01b600052600160045260246000fd5b90611dfb8251611ca681151561095e565b91611e0583611871565b9260009081925b818410611e225750506102419150835114611dcd565b9091611e63611e6991600180871615600014611e7157611e52611d1661161f6116048a600f60f81b951c896114f9565b60f81b1660001a611d2b828a6114f9565b93610bd2565b929190611e0c565b611e52611d5e61161f6116048a600f60f81b951c896114f9565b919060005b83810180821161098a578251811080611ef0575b15611ee8576001600160f81b0319908190611ebf90856114f9565b511690611ecc83866114f9565b511603611ee157611edc9061097b565b611e90565b9250505090565b509250505090565b5083518210611ea4565b60405190611f07826101d8565b60018252600160ff1b6020830152565b8015611ffd57607f811115611fd85760ff811115611fa95761ffff80821115611f7b575062ffffff9081811115611f4d57600080fd5b604051608360f81b6020820152911660e81b6001600160e81b0319166021820152610b808160248101610362565b604051604160f91b6020820152911660f01b6001600160f01b0319166021820152610b808160238101610362565b604051608160f81b602082015260f89190911b6001600160f81b0319166021820152610b808160228101610362565b60405160f89190911b6001600160f81b0319166020820152610b808160218101610362565b50610b80611efa56fea264697066735822122008d4e914ea3f395ab09085b78b1480c2af75c1cfe285d85877803517f31dbb3564736f6c63430008120033",
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
// Solidity: function calculateIdentifier(address creator, uint256 createdAt, address rewardToken, string metadata, bytes script) pure returns(bytes32)
func (_StartVoteBounty *StartVoteBountyCaller) CalculateIdentifier(opts *bind.CallOpts, creator common.Address, createdAt *big.Int, rewardToken common.Address, metadata string, script []byte) ([32]byte, error) {
	var out []interface{}
	err := _StartVoteBounty.contract.Call(opts, &out, "calculateIdentifier", creator, createdAt, rewardToken, metadata, script)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateIdentifier is a free data retrieval call binding the contract method 0x1ce63f1a.
//
// Solidity: function calculateIdentifier(address creator, uint256 createdAt, address rewardToken, string metadata, bytes script) pure returns(bytes32)
func (_StartVoteBounty *StartVoteBountySession) CalculateIdentifier(creator common.Address, createdAt *big.Int, rewardToken common.Address, metadata string, script []byte) ([32]byte, error) {
	return _StartVoteBounty.Contract.CalculateIdentifier(&_StartVoteBounty.CallOpts, creator, createdAt, rewardToken, metadata, script)
}

// CalculateIdentifier is a free data retrieval call binding the contract method 0x1ce63f1a.
//
// Solidity: function calculateIdentifier(address creator, uint256 createdAt, address rewardToken, string metadata, bytes script) pure returns(bytes32)
func (_StartVoteBounty *StartVoteBountyCallerSession) CalculateIdentifier(creator common.Address, createdAt *big.Int, rewardToken common.Address, metadata string, script []byte) ([32]byte, error) {
	return _StartVoteBounty.Contract.CalculateIdentifier(&_StartVoteBounty.CallOpts, creator, createdAt, rewardToken, metadata, script)
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
// Solidity: function applyCloseBounty(uint256 createdAt, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) ApplyCloseBounty(opts *bind.TransactOpts, createdAt *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "applyCloseBounty", createdAt, rewardToken, digest)
}

// ApplyCloseBounty is a paid mutator transaction binding the contract method 0xc2084b04.
//
// Solidity: function applyCloseBounty(uint256 createdAt, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountySession) ApplyCloseBounty(createdAt *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.ApplyCloseBounty(&_StartVoteBounty.TransactOpts, createdAt, rewardToken, digest)
}

// ApplyCloseBounty is a paid mutator transaction binding the contract method 0xc2084b04.
//
// Solidity: function applyCloseBounty(uint256 createdAt, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) ApplyCloseBounty(createdAt *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.ApplyCloseBounty(&_StartVoteBounty.TransactOpts, createdAt, rewardToken, digest)
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
// Solidity: function claimBounty(address creator, uint256 createdAt, address rewardToken, bytes32 digest, uint256 receiptIndex, bytes headerRlpBytes, bytes proofRlpBytes) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) ClaimBounty(opts *bind.TransactOpts, creator common.Address, createdAt *big.Int, rewardToken common.Address, digest [32]byte, receiptIndex *big.Int, headerRlpBytes []byte, proofRlpBytes []byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "claimBounty", creator, createdAt, rewardToken, digest, receiptIndex, headerRlpBytes, proofRlpBytes)
}

// ClaimBounty is a paid mutator transaction binding the contract method 0x2fa970d8.
//
// Solidity: function claimBounty(address creator, uint256 createdAt, address rewardToken, bytes32 digest, uint256 receiptIndex, bytes headerRlpBytes, bytes proofRlpBytes) returns()
func (_StartVoteBounty *StartVoteBountySession) ClaimBounty(creator common.Address, createdAt *big.Int, rewardToken common.Address, digest [32]byte, receiptIndex *big.Int, headerRlpBytes []byte, proofRlpBytes []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.ClaimBounty(&_StartVoteBounty.TransactOpts, creator, createdAt, rewardToken, digest, receiptIndex, headerRlpBytes, proofRlpBytes)
}

// ClaimBounty is a paid mutator transaction binding the contract method 0x2fa970d8.
//
// Solidity: function claimBounty(address creator, uint256 createdAt, address rewardToken, bytes32 digest, uint256 receiptIndex, bytes headerRlpBytes, bytes proofRlpBytes) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) ClaimBounty(creator common.Address, createdAt *big.Int, rewardToken common.Address, digest [32]byte, receiptIndex *big.Int, headerRlpBytes []byte, proofRlpBytes []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.ClaimBounty(&_StartVoteBounty.TransactOpts, creator, createdAt, rewardToken, digest, receiptIndex, headerRlpBytes, proofRlpBytes)
}

// CommitCloseBounty is a paid mutator transaction binding the contract method 0xad22eb62.
//
// Solidity: function commitCloseBounty(uint256 createdAt, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) CommitCloseBounty(opts *bind.TransactOpts, createdAt *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "commitCloseBounty", createdAt, rewardToken, digest)
}

// CommitCloseBounty is a paid mutator transaction binding the contract method 0xad22eb62.
//
// Solidity: function commitCloseBounty(uint256 createdAt, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountySession) CommitCloseBounty(createdAt *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.CommitCloseBounty(&_StartVoteBounty.TransactOpts, createdAt, rewardToken, digest)
}

// CommitCloseBounty is a paid mutator transaction binding the contract method 0xad22eb62.
//
// Solidity: function commitCloseBounty(uint256 createdAt, address rewardToken, bytes32 digest) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) CommitCloseBounty(createdAt *big.Int, rewardToken common.Address, digest [32]byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.CommitCloseBounty(&_StartVoteBounty.TransactOpts, createdAt, rewardToken, digest)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0xf5ba9296.
//
// Solidity: function increaseBounty(uint256 createdAt, address rewardToken, bytes32 digest, uint256 increaseAmount) returns()
func (_StartVoteBounty *StartVoteBountyTransactor) IncreaseBounty(opts *bind.TransactOpts, createdAt *big.Int, rewardToken common.Address, digest [32]byte, increaseAmount *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "increaseBounty", createdAt, rewardToken, digest, increaseAmount)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0xf5ba9296.
//
// Solidity: function increaseBounty(uint256 createdAt, address rewardToken, bytes32 digest, uint256 increaseAmount) returns()
func (_StartVoteBounty *StartVoteBountySession) IncreaseBounty(createdAt *big.Int, rewardToken common.Address, digest [32]byte, increaseAmount *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.IncreaseBounty(&_StartVoteBounty.TransactOpts, createdAt, rewardToken, digest, increaseAmount)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0xf5ba9296.
//
// Solidity: function increaseBounty(uint256 createdAt, address rewardToken, bytes32 digest, uint256 increaseAmount) returns()
func (_StartVoteBounty *StartVoteBountyTransactorSession) IncreaseBounty(createdAt *big.Int, rewardToken common.Address, digest [32]byte, increaseAmount *big.Int) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.IncreaseBounty(&_StartVoteBounty.TransactOpts, createdAt, rewardToken, digest, increaseAmount)
}

// OpenBounty is a paid mutator transaction binding the contract method 0xd6841c94.
//
// Solidity: function openBounty(address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountyTransactor) OpenBounty(opts *bind.TransactOpts, rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.contract.Transact(opts, "openBounty", rewardToken, rewardAmount, metadata, script)
}

// OpenBounty is a paid mutator transaction binding the contract method 0xd6841c94.
//
// Solidity: function openBounty(address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountySession) OpenBounty(rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.OpenBounty(&_StartVoteBounty.TransactOpts, rewardToken, rewardAmount, metadata, script)
}

// OpenBounty is a paid mutator transaction binding the contract method 0xd6841c94.
//
// Solidity: function openBounty(address rewardToken, uint256 rewardAmount, string metadata, bytes script) payable returns(bytes32)
func (_StartVoteBounty *StartVoteBountyTransactorSession) OpenBounty(rewardToken common.Address, rewardAmount *big.Int, metadata string, script []byte) (*types.Transaction, error) {
	return _StartVoteBounty.Contract.OpenBounty(&_StartVoteBounty.TransactOpts, rewardToken, rewardAmount, metadata, script)
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
	Metadata     string
	Script       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOpenBounty is a free log retrieval operation binding the contract event 0x19f082eb5536f4081e8e296229b3b1aedefd1a248e83e7d37d176a76f92ddafd.
//
// Solidity: event OpenBounty(bytes32 indexed identifier, address indexed creator, address indexed rewardToken, uint256 rewardAmount, string metadata, bytes script)
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

// WatchOpenBounty is a free log subscription operation binding the contract event 0x19f082eb5536f4081e8e296229b3b1aedefd1a248e83e7d37d176a76f92ddafd.
//
// Solidity: event OpenBounty(bytes32 indexed identifier, address indexed creator, address indexed rewardToken, uint256 rewardAmount, string metadata, bytes script)
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

// ParseOpenBounty is a log parse operation binding the contract event 0x19f082eb5536f4081e8e296229b3b1aedefd1a248e83e7d37d176a76f92ddafd.
//
// Solidity: event OpenBounty(bytes32 indexed identifier, address indexed creator, address indexed rewardToken, uint256 rewardAmount, string metadata, bytes script)
func (_StartVoteBounty *StartVoteBountyFilterer) ParseOpenBounty(log types.Log) (*StartVoteBountyOpenBounty, error) {
	event := new(StartVoteBountyOpenBounty)
	if err := _StartVoteBounty.contract.UnpackLog(event, "OpenBounty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
