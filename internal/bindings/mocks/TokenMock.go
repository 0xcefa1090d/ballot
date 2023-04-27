// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

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

// TokenMockMetaData contains all meta data concerning the TokenMock contract.
var TokenMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x0x60e0604090808252346200049b578062000fc38038038091620000238285620004bc565b83396020928391810103126200049b57518251916200004283620004a0565b600a8352692a32b9ba102a37b5b2b760b11b818401528351916200006683620004a0565b600580845264054535432360dc1b8385015284516001600160401b03959094908686116200048557600095806200009e8854620004e0565b92601f9384811162000436575b508790848311600114620003ce578992620003c2575b50508160011b916000199060031b1c19161786555b815190878211620003ae578190600193620000f28554620004e0565b8281116200035b575b5087918311600114620002f7578892620002eb575b5050600019600383901b1c191690821b1781555b60126080524660a05286519081868754926200014084620004e0565b9384845288808501978383169283600014620002cc575050506001146200028c575b506200017192500382620004bc565b5190208551838101917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8352878201527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260a0815260c081019581871090871117620002785785875251902060c0526002548181018091116200026457907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92916002553384526003825285842081815401905584523393a351610aa590816200051e823960805181610552015260a051816108ce015260c051816108f50152f35b634e487b7160e01b84526011600452602484fd5b634e487b7160e01b85526041600452602485fd5b8791508880528189209089915b858310620002b35750506200017193508201013862000162565b8054838801850152869450899390920191810162000299565b60ff1916895262000171961515901b8501019250389150620001629050565b01519050388062000110565b8489528789208594509190601f1984168a5b8a8282106200034457505084116200032a575b505050811b01815562000124565b015160001960f88460031b161c191690553880806200031c565b838501518655889790950194938401930162000309565b9091925084895287892083808601881c8201928a8710620003a4575b918695889295949301891c01915b82811062000395575050620000fb565b8b815586955087910162000385565b9250819262000377565b634e487b7160e01b87526041600452602487fd5b015190503880620000c1565b898052888a209250601f1984168a5b8a8282106200041f57505090846001959493921062000405575b505050811b018655620000d6565b015160001960f88460031b161c19169055388080620003f7565b6001859682939686015181550195019301620003dd565b90915088805287892084808501881c8201928a86106200047b575b90859493929101881c01905b8181106200046c5750620000ab565b8a81558493506001016200045d565b9250819262000451565b634e487b7160e01b600052604160045260246000fd5b600080fd5b604081019081106001600160401b038211176200048557604052565b601f909101601f19168101906001600160401b038211908210176200048557604052565b90600182811c9216801562000512575b6020831014620004fc57565b634e487b7160e01b600052602260045260246000fd5b91607f1691620004f056fe6080604081815260048036101561001557600080fd5b600092833560e01c90816306fdde031461071457508063095ea7b3146106a657806318160ddd1461068757806323b872dd146105b157806330adf81f14610576578063313ce567146105385780633644e5151461051457806370a08231146104dc5780637ecebe00146104a457806395d89b41146103c5578063a9059cbb14610341578063d505accf146100fd5763dd62ed3e146100b257600080fd5b346100f957816003193601126100f95760209282916100cf610875565b6100d7610890565b6001600160a01b03918216845291865283832091168252845220549051908152f35b8280fd5b5091903461033d5760e036600319011261033d57610119610875565b90610122610890565b91604435606435926084359260ff8416809403610339574285106102f6576101486108c9565b9560018060a01b038092169586895260209560058752848a209889549960018b01905585519285898501957f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c987528b89870152169a8b606086015288608086015260a085015260c084015260c0835260e0830167ffffffffffffffff94848210868311176102e2578188528451902061010085019261190160f01b845261010286015261012285015260428152610160840194818610908611176102cf57848752519020835261018082015260a4356101a082015260c4356101c0909101528780528490889060809060015afa156102c55786511696871515806102bc575b1561028a5786977f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259596975283528087208688528352818188205551908152a380f35b83606492519162461bcd60e51b8352820152600e60248201526d24a72b20a624a22fa9a4a3a722a960911b6044820152fd5b50848814610247565b81513d88823e3d90fd5b634e487b7160e01b8c5260418d5260248cfd5b50634e487b7160e01b8c5260418d5260248cfd5b815162461bcd60e51b81526020818a0152601760248201527f5045524d49545f444541444c494e455f455850495245440000000000000000006044820152606490fd5b8680fd5b5080fd5b50503461033d578060031936011261033d5760209161035e610875565b8260243591338452600386528184206103788482546108a6565b90556001600160a01b0316808452600386529220805482019055825190815233907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908590a35160018152f35b50503461033d578160031936011261033d578051908260018054916103e9836107ba565b8086529282811690811561047c5750600114610420575b5050506104128261041c9403836107f4565b519182918261082c565b0390f35b94508085527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b8286106104645750505061041282602061041c9582010194610400565b80546020878701810191909152909501948101610447565b61041c97508693506020925061041294915060ff191682840152151560051b82010194610400565b50503461033d57602036600319011261033d5760209181906001600160a01b036104cc610875565b1681526005845220549051908152f35b50503461033d57602036600319011261033d5760209181906001600160a01b03610504610875565b1681526003845220549051908152f35b50503461033d578160031936011261033d576020906105316108c9565b9051908152f35b50503461033d578160031936011261033d576020905160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50503461033d578160031936011261033d57602090517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98152f35b509134610684576060366003190112610684576105cc610875565b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6105f5610890565b6001600160a01b0392831680855260208781528686203387528152868620549097919488936044359389938560018201610661575b505050868852600385528288206106428582546108a6565b9055169586815260038452208181540190558551908152a35160018152f35b61066a916108a6565b90888a528652838920338a5286528389205538808561062a565b80fd5b50503461033d578160031936011261033d576020906002549051908152f35b50346100f957816003193601126100f9576020926106c2610875565b918360243592839233825287528181209460018060a01b0316948582528752205582519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925843392a35160018152f35b849084346100f957826003193601126100f957828054610733816107ba565b8085529160019180831690811561047c575060011461075e575050506104128261041c9403836107f4565b80809650527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b8286106107a25750505061041282602061041c9582010194610400565b80546020878701810191909152909501948101610785565b90600182811c921680156107ea575b60208310146107d457565b634e487b7160e01b600052602260045260246000fd5b91607f16916107c9565b90601f8019910116810190811067ffffffffffffffff82111761081657604052565b634e487b7160e01b600052604160045260246000fd5b6020808252825181830181905290939260005b82811061086157505060409293506000838284010152601f8019910116010190565b81810186015184820160400152850161083f565b600435906001600160a01b038216820361088b57565b600080fd5b602435906001600160a01b038216820361088b57565b919082039182116108b357565b634e487b7160e01b600052601160045260246000fd5b6000467f00000000000000000000000000000000000000000000000000000000000000000361091757507f000000000000000000000000000000000000000000000000000000000000000090565b60405181548291610927826107ba565b8082528160209485820194600190878282169182600014610a515750506001146109f8575b50610959925003826107f4565b51902091604051918201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f845260408301527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608301524660808301523060a083015260a0825260c082019082821067ffffffffffffffff8311176109e4575060405251902090565b634e487b7160e01b81526041600452602490fd5b87805286915087907f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b858310610a3957505061095993508201013861094c565b80548388018501528694508893909201918101610a22565b60ff1916885261095995151560051b850101925038915061094c905056fea2646970667358221220d506083b23bc7fddbfc1d6d633d6bbbbcc2e54b6352bdde7d25dfaf196d1227464736f6c63430008120033",
}

// TokenMockABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMockMetaData.ABI instead.
var TokenMockABI = TokenMockMetaData.ABI

// TokenMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenMockMetaData.Bin instead.
var TokenMockBin = TokenMockMetaData.Bin

// DeployTokenMock deploys a new Ethereum contract, binding an instance of TokenMock to it.
func DeployTokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, _initialSupply *big.Int) (common.Address, *types.Transaction, *TokenMock, error) {
	parsed, err := TokenMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenMockBin), backend, _initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenMock{TokenMockCaller: TokenMockCaller{contract: contract}, TokenMockTransactor: TokenMockTransactor{contract: contract}, TokenMockFilterer: TokenMockFilterer{contract: contract}}, nil
}

// TokenMock is an auto generated Go binding around an Ethereum contract.
type TokenMock struct {
	TokenMockCaller     // Read-only binding to the contract
	TokenMockTransactor // Write-only binding to the contract
	TokenMockFilterer   // Log filterer for contract events
}

// TokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenMockSession struct {
	Contract     *TokenMock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenMockCallerSession struct {
	Contract *TokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenMockTransactorSession struct {
	Contract     *TokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenMockRaw struct {
	Contract *TokenMock // Generic contract binding to access the raw methods on
}

// TokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenMockCallerRaw struct {
	Contract *TokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// TokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenMockTransactorRaw struct {
	Contract *TokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenMock creates a new instance of TokenMock, bound to a specific deployed contract.
func NewTokenMock(address common.Address, backend bind.ContractBackend) (*TokenMock, error) {
	contract, err := bindTokenMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenMock{TokenMockCaller: TokenMockCaller{contract: contract}, TokenMockTransactor: TokenMockTransactor{contract: contract}, TokenMockFilterer: TokenMockFilterer{contract: contract}}, nil
}

// NewTokenMockCaller creates a new read-only instance of TokenMock, bound to a specific deployed contract.
func NewTokenMockCaller(address common.Address, caller bind.ContractCaller) (*TokenMockCaller, error) {
	contract, err := bindTokenMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenMockCaller{contract: contract}, nil
}

// NewTokenMockTransactor creates a new write-only instance of TokenMock, bound to a specific deployed contract.
func NewTokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenMockTransactor, error) {
	contract, err := bindTokenMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenMockTransactor{contract: contract}, nil
}

// NewTokenMockFilterer creates a new log filterer instance of TokenMock, bound to a specific deployed contract.
func NewTokenMockFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenMockFilterer, error) {
	contract, err := bindTokenMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenMockFilterer{contract: contract}, nil
}

// bindTokenMock binds a generic wrapper to an already deployed contract.
func bindTokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenMock *TokenMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenMock.Contract.TokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenMock *TokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenMock.Contract.TokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenMock *TokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenMock.Contract.TokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenMock *TokenMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenMock *TokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenMock *TokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenMock.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TokenMock *TokenMockCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TokenMock *TokenMockSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TokenMock.Contract.DOMAINSEPARATOR(&_TokenMock.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TokenMock *TokenMockCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TokenMock.Contract.DOMAINSEPARATOR(&_TokenMock.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_TokenMock *TokenMockCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_TokenMock *TokenMockSession) PERMITTYPEHASH() ([32]byte, error) {
	return _TokenMock.Contract.PERMITTYPEHASH(&_TokenMock.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_TokenMock *TokenMockCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _TokenMock.Contract.PERMITTYPEHASH(&_TokenMock.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_TokenMock *TokenMockCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_TokenMock *TokenMockSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenMock.Contract.Allowance(&_TokenMock.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_TokenMock *TokenMockCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenMock.Contract.Allowance(&_TokenMock.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TokenMock *TokenMockCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TokenMock *TokenMockSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TokenMock.Contract.BalanceOf(&_TokenMock.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TokenMock *TokenMockCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TokenMock.Contract.BalanceOf(&_TokenMock.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenMock *TokenMockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenMock *TokenMockSession) Decimals() (uint8, error) {
	return _TokenMock.Contract.Decimals(&_TokenMock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenMock *TokenMockCallerSession) Decimals() (uint8, error) {
	return _TokenMock.Contract.Decimals(&_TokenMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenMock *TokenMockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenMock *TokenMockSession) Name() (string, error) {
	return _TokenMock.Contract.Name(&_TokenMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenMock *TokenMockCallerSession) Name() (string, error) {
	return _TokenMock.Contract.Name(&_TokenMock.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_TokenMock *TokenMockCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_TokenMock *TokenMockSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _TokenMock.Contract.Nonces(&_TokenMock.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_TokenMock *TokenMockCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _TokenMock.Contract.Nonces(&_TokenMock.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenMock *TokenMockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenMock *TokenMockSession) Symbol() (string, error) {
	return _TokenMock.Contract.Symbol(&_TokenMock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenMock *TokenMockCallerSession) Symbol() (string, error) {
	return _TokenMock.Contract.Symbol(&_TokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenMock *TokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenMock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenMock *TokenMockSession) TotalSupply() (*big.Int, error) {
	return _TokenMock.Contract.TotalSupply(&_TokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenMock *TokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _TokenMock.Contract.TotalSupply(&_TokenMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenMock *TokenMockTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenMock *TokenMockSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.Contract.Approve(&_TokenMock.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenMock *TokenMockTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.Contract.Approve(&_TokenMock.TransactOpts, spender, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TokenMock *TokenMockTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TokenMock.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TokenMock *TokenMockSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TokenMock.Contract.Permit(&_TokenMock.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TokenMock *TokenMockTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TokenMock.Contract.Permit(&_TokenMock.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenMock *TokenMockTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenMock *TokenMockSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.Contract.Transfer(&_TokenMock.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenMock *TokenMockTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.Contract.Transfer(&_TokenMock.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenMock *TokenMockTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenMock *TokenMockSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.Contract.TransferFrom(&_TokenMock.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenMock *TokenMockTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMock.Contract.TransferFrom(&_TokenMock.TransactOpts, from, to, amount)
}

// TokenMockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenMock contract.
type TokenMockApprovalIterator struct {
	Event *TokenMockApproval // Event containing the contract specifics and raw log

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
func (it *TokenMockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMockApproval)
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
		it.Event = new(TokenMockApproval)
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
func (it *TokenMockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMockApproval represents a Approval event raised by the TokenMock contract.
type TokenMockApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_TokenMock *TokenMockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenMockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenMock.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenMockApprovalIterator{contract: _TokenMock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_TokenMock *TokenMockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenMockApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenMock.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMockApproval)
				if err := _TokenMock.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_TokenMock *TokenMockFilterer) ParseApproval(log types.Log) (*TokenMockApproval, error) {
	event := new(TokenMockApproval)
	if err := _TokenMock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenMock contract.
type TokenMockTransferIterator struct {
	Event *TokenMockTransfer // Event containing the contract specifics and raw log

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
func (it *TokenMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMockTransfer)
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
		it.Event = new(TokenMockTransfer)
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
func (it *TokenMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMockTransfer represents a Transfer event raised by the TokenMock contract.
type TokenMockTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_TokenMock *TokenMockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenMockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenMock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenMockTransferIterator{contract: _TokenMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_TokenMock *TokenMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenMockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenMock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMockTransfer)
				if err := _TokenMock.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_TokenMock *TokenMockFilterer) ParseTransfer(log types.Log) (*TokenMockTransfer, error) {
	event := new(TokenMockTransfer)
	if err := _TokenMock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
