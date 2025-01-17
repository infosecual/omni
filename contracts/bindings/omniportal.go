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

// OmniPortalInitParams is an auto generated low-level Go binding around an user-defined struct.
type OmniPortalInitParams struct {
	Owner                common.Address
	FeeOracle            common.Address
	OmniChainId          uint64
	OmniCChainId         uint64
	XmsgMaxGasLimit      uint64
	XmsgMinGasLimit      uint64
	XmsgMaxDataSize      uint16
	XreceiptMaxErrorSize uint16
	XsubValsetCutoff     uint8
	CChainXMsgOffset     uint64
	CChainXBlockOffset   uint64
	ValSetId             uint64
	Validators           []XTypesValidator
}

// XTypesBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type XTypesBlockHeader struct {
	SourceChainId     uint64
	ConsensusChainId  uint64
	ConfLevel         uint8
	Offset            uint64
	SourceBlockHeight uint64
	SourceBlockHash   [32]byte
}

// XTypesChain is an auto generated low-level Go binding around an user-defined struct.
type XTypesChain struct {
	ChainId uint64
	Shards  []uint64
}

// XTypesMsg is an auto generated low-level Go binding around an user-defined struct.
type XTypesMsg struct {
	DestChainId uint64
	ShardId     uint64
	Offset      uint64
	Sender      common.Address
	To          common.Address
	Data        []byte
	GasLimit    uint64
}

// XTypesMsgContext is an auto generated low-level Go binding around an user-defined struct.
type XTypesMsgContext struct {
	SourceChainId uint64
	Sender        common.Address
}

// XTypesSigTuple is an auto generated low-level Go binding around an user-defined struct.
type XTypesSigTuple struct {
	ValidatorAddr common.Address
	Signature     []byte
}

// XTypesSubmission is an auto generated low-level Go binding around an user-defined struct.
type XTypesSubmission struct {
	AttestationRoot [32]byte
	ValidatorSetId  uint64
	BlockHeader     XTypesBlockHeader
	Msgs            []XTypesMsg
	Proof           [][32]byte
	ProofFlags      []bool
	Signatures      []XTypesSigTuple
}

// XTypesValidator is an auto generated low-level Go binding around an user-defined struct.
type XTypesValidator struct {
	Addr  common.Address
	Power uint64
}

// OmniPortalMetaData contains all meta data concerning the OmniPortal contract.
var OmniPortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ActionXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ActionXSubmit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyPauseAll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSubQuorumDenominator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSubQuorumNumerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addValidatorSet\",\"inputs\":[{\"name\":\"valSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collectFees\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXBlockOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXMsgOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structOmniPortal.InitParams\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omniChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"omniCChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMaxGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMinGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMaxDataSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"xreceiptMaxErrorSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"xsubValsetCutoff\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"cChainXMsgOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cChainXBlockOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"valSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"actionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"actionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedDest\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedShard\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestValSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniCChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outXMsgOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXCallTo\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXSubmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXSubmitFrom\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeOracle\",\"inputs\":[{\"name\":\"feeOracle_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInXBlockOffset\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInXMsgOffset\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNetwork\",\"inputs\":[{\"name\":\"network_\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Chain[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMaxDataSize\",\"inputs\":[{\"name\":\"numBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMaxGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMinGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXReceiptMaxErrorSize\",\"inputs\":[{\"name\":\"numBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXSubValsetCutoff\",\"inputs\":[{\"name\":\"xsubValsetCutoff_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXCallTo\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXSubmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXSubmitFrom\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"valSetTotalPower\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"conf\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xmsg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXTypes.MsgContext\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMaxDataSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMaxGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMinGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xreceiptMaxErrorSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubValsetCutoff\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubmit\",\"inputs\":[{\"name\":\"xsub\",\"type\":\"tuple\",\"internalType\":\"structXTypes.Submission\",\"components\":[{\"name\":\"attestationRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeader\",\"type\":\"tuple\",\"internalType\":\"structXTypes.BlockHeader\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"consensusChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"confLevel\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Msg[]\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlags\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.SigTuple[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeOracleSet\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesCollected\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InXBlockOffsetSet\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InXMsgOffsetSet\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSetAdded\",\"inputs\":[{\"name\":\"setId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallToPaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallToUnpaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsg\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMaxDataSizeSet\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMaxGasLimitSet\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMinGasLimitSet\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceipt\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"err\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceiptMaxErrorSizeSet\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubValsetCutoffSet\",\"inputs\":[{\"name\":\"cutoff\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitFromPaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitFromUnpaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleProofInvalidMultiproof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000d6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000735760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d35780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61497580620000e66000396000f3fe60806040526004361061036b5760003560e01c80638532eb9f116101c6578063b4d5afd1116100f7578063c3d8ad6711610095578063d051c97d1161006f578063d051c97d14610af6578063d533b44514610b37578063f2fde38b14610b57578063f45cc7b814610b7757600080fd5b8063c3d8ad6714610a98578063c4ab80bc14610aad578063cf84c81814610acd57600080fd5b8063bff0e84d116100d1578063bff0e84d14610a25578063c21dda4f14610a45578063c26dfc0514610a58578063c2f9b96814610a7857600080fd5b8063b4d5afd1146109b0578063b521466d146109e5578063bb8590ad14610a0557600080fd5b8063a480ca7911610164578063afe821981161013e578063afe8219814610923578063afe8af9c14610943578063b187bd2614610979578063b2b2f5bd1461098e57600080fd5b8063a480ca79146108b3578063a8a98962146108d3578063aaf1bc97146108f357600080fd5b806397b52062116101a057806397b520621461083c5780639a8a05921461085c578063a10ac97a1461086f578063a32eb7c61461089157600080fd5b80638532eb9f146107b15780638da5cb5b146107d15780638dd9523c1461080e57600080fd5b80633f4ba83a116102a0578063575420501161023e57806374eba9391161021857806374eba9391461074057806378fe53071461076057806383d0cbd9146107875780638456cb591461079c57600080fd5b806357542050146106ca57806366a1eaf31461070b578063715018a61461072b57600080fd5b806349cc3bf61161027a57806349cc3bf614610643578063500b19e71461065d57806354d26bba1461069557806355e2448e146106aa57600080fd5b80633f4ba83a146105cd5780633fd3b15e146105e2578063461ab4881461062357600080fd5b8063241b71bb1161030d57806330632e8b116102e757806330632e8b1461052557806336d219121461054557806336d853f91461056c5780633aa873301461058c57600080fd5b8063241b71bb1461046057806324278bbe146104905780632f32700e146104c057600080fd5b806310a5a7f71161034957806310a5a7f7146103d3578063110ff5f1146103f35780631d3eb6e31461042b57806323dbce501461044b57600080fd5b80630360d20f1461037057806306c3dc5f1461039c578063103ba701146103b1575b600080fd5b34801561037c57600080fd5b50610385600281565b60405160ff90911681526020015b60405180910390f35b3480156103a857600080fd5b50610385600381565b3480156103bd57600080fd5b506103d16103cc366004613cf7565b610b9e565b005b3480156103df57600080fd5b506103d16103ee366004613d32565b610bb2565b3480156103ff57600080fd5b50600154610413906001600160401b031681565b6040516001600160401b039091168152602001610393565b34801561043757600080fd5b506103d1610446366004613d4f565b610c11565b34801561045757600080fd5b506103d1610d2c565b34801561046c57600080fd5b5061048061047b366004613dc3565b610d76565b6040519015158152602001610393565b34801561049c57600080fd5b506104806104ab366004613d32565b60056020526000908152604090205460ff1681565b3480156104cc57600080fd5b50604080518082018252600080825260209182015281518083018352600b546001600160401b0381168083526001600160a01b03600160401b909204821692840192835284519081529151169181019190915201610393565b34801561053157600080fd5b506103d1610540366004613ddc565b610d87565b34801561055157600080fd5b5060015461041390600160401b90046001600160401b031681565b34801561057857600080fd5b506103d1610587366004613d32565b611028565b34801561059857600080fd5b506104136105a7366004613e17565b60066020908152600092835260408084209091529082529020546001600160401b031681565b3480156105d957600080fd5b506103d1611039565b3480156105ee57600080fd5b506104136105fd366004613e17565b60086020908152600092835260408084209091529082529020546001600160401b031681565b34801561062f57600080fd5b5061048061063e366004613e50565b611074565b34801561064f57600080fd5b506000546103859060ff1681565b34801561066957600080fd5b5060025461067d906001600160a01b031681565b6040516001600160a01b039091168152602001610393565b3480156106a157600080fd5b506103d1611090565b3480156106b657600080fd5b50600b546001600160401b03161515610480565b3480156106d657600080fd5b506104136106e5366004613e8c565b600a6020908152600092835260408084209091529082529020546001600160401b031681565b34801561071757600080fd5b506103d1610726366004613ec1565b6110da565b34801561073757600080fd5b506103d161147a565b34801561074c57600080fd5b5061041361075b366004613dc3565b61148e565b34801561076c57600080fd5b5060005461041390600160681b90046001600160401b031681565b34801561079357600080fd5b506103d16114bd565b3480156107a857600080fd5b506103d1611507565b3480156107bd57600080fd5b506103d16107cc366004613efc565b611542565b3480156107dd57600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661067d565b34801561081a57600080fd5b5061082e610829366004613fcb565b611655565b604051908152602001610393565b34801561084857600080fd5b506103d1610857366004614032565b6116d6565b34801561086857600080fd5b5046610413565b34801561087b57600080fd5b5061082e6000805160206148e083398151915281565b34801561089d57600080fd5b5061082e60008051602061492083398151915281565b3480156108bf57600080fd5b506103d16108ce36600461407d565b6116e9565b3480156108df57600080fd5b506103d16108ee36600461407d565b611771565b3480156108ff57600080fd5b5061048061090e366004613d32565b60046020526000908152604090205460ff1681565b34801561092f57600080fd5b506103d161093e366004613d32565b611782565b34801561094f57600080fd5b5061041361095e366004613d32565b6009602052600090815260409020546001600160401b031681565b34801561098557600080fd5b506104806117dc565b34801561099a57600080fd5b5061082e6000805160206148c083398151915281565b3480156109bc57600080fd5b506000546109d2906301000000900461ffff1681565b60405161ffff9091168152602001610393565b3480156109f157600080fd5b506103d1610a00366004614098565b611832565b348015610a1157600080fd5b506103d1610a20366004613d32565b611843565b348015610a3157600080fd5b506103d1610a40366004614098565b611854565b6103d1610a533660046140bc565b611865565b348015610a6457600080fd5b506000546109d290610100900461ffff1681565b348015610a8457600080fd5b506103d1610a93366004613d32565b611c3f565b348015610aa457600080fd5b506103d1611c9e565b348015610ab957600080fd5b506103d1610ac8366004614032565b611ce8565b348015610ad957600080fd5b50600054610413906501000000000090046001600160401b031681565b348015610b0257600080fd5b50610413610b11366004613e17565b60076020908152600092835260408084209091529082529020546001600160401b031681565b348015610b4357600080fd5b506103d1610b52366004613d32565b611cfb565b348015610b6357600080fd5b506103d1610b7236600461407d565b611d55565b348015610b8357600080fd5b5060005461041390600160a81b90046001600160401b031681565b610ba6611d90565b610baf81611deb565b50565b610bba611d90565b610bda610bd56000805160206148c083398151915283611e87565b611ed0565b6040516001600160401b038216907fcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a90600090a250565b333014610c5d5760405162461bcd60e51b815260206004820152601560248201527427b6b734a837b93a30b61d1037b7363c9039b2b63360591b60448201526064015b60405180910390fd5b600154600b546001600160401b03908116600160401b9092041614610cbe5760405162461bcd60e51b815260206004820152601760248201527627b6b734a837b93a30b61d1037b7363c9031b1b430b4b760491b6044820152606401610c54565b600b54600160401b90046001600160a01b031615610d1e5760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a206f6e6c792063636861696e2073656e64657200006044820152606401610c54565b610d288282611f76565b5050565b610d34611d90565b610d4b600080516020614920833981519152611ed0565b6040517f3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed890600090a1565b6000610d81826120f0565b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b0316600081158015610dcc5750825b90506000826001600160401b03166001148015610de85750303b155b905081158015610df6575080155b15610e145760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610e3e57845460ff60401b1916600160401b1785555b610e53610e4e602088018861407d565b612157565b610e6b610e66604088016020890161407d565b612168565b610e83610e7e60a0880160808901613d32565b61220c565b610e9b610e9660c0880160a08901613d32565b6122c4565b610eb3610eae60e0880160c08901614098565b612378565b610ecc610ec7610100880160e08901614098565b61241c565b610ee6610ee161012088016101008901613cf7565b611deb565b610f0e610efb61018088016101608901613d32565b610f09610180890189614145565b6124bc565b610f1e6060870160408801613d32565b6001805467ffffffffffffffff19166001600160401b0392909216919091179055610f4f6080870160608801613d32565b600180546001600160401b0392909216600160401b026fffffffffffffffff000000000000000019909216919091179055610104610fae610f966080890160608a01613d32565b82610fa96101408b016101208c01613d32565b6127eb565b610fd9610fc16080890160608a01613d32565b82610fd46101608b016101408c01613d32565b612860565b50831561102057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b611030611d90565b610baf8161220c565b611041611d90565b6110496128cc565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d1693390600090a1565b6000611089836110848585611e87565b6128e3565b9392505050565b611098611d90565b6110af6000805160206148c083398151915261296a565b6040517f4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e679482790600090a1565b6000805160206149208339815191526110f96060830160408401613d32565b611107826110848484611e87565b156111495760405162461bcd60e51b815260206004820152601260248201527113db5b9a541bdc9d185b0e881c185d5cd95960721b6044820152606401610c54565b611151612a10565b36600061116261010086018661418e565b909250905060408501600061117a8260208901613d32565b600154909150600160401b90046001600160401b03166111a06040840160208501613d32565b6001600160401b0316146111f65760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a2077726f6e672063636861696e20494400000000006044820152606401610c54565b8261123a5760405162461bcd60e51b81526020600482015260146024820152734f6d6e69506f7274616c3a206e6f20786d73677360601b6044820152606401610c54565b6001600160401b03808216600090815260096020526040902054166112a15760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a20756e6b6e6f776e2076616c2073657400000000006044820152606401610c54565b6112a9612a5a565b6001600160401b0316816001600160401b0316101561130a5760405162461bcd60e51b815260206004820152601760248201527f4f6d6e69506f7274616c3a206f6c642076616c207365740000000000000000006044820152606401610c54565b61134e873561131d6101608a018a61418e565b6001600160401b038086166000908152600a6020908152604080832060099092529091205490911660026003612aaa565b6113925760405162461bcd60e51b81526020600482015260156024820152744f6d6e69506f7274616c3a206e6f2071756f72756d60581b6044820152606401610c54565b6113bb87358386866113a86101208d018d61418e565b6113b66101408f018f61418e565b612c5e565b6114075760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a20696e76616c69642070726f6f66000000000000006044820152606401610c54565b60005b838110156114475761143f83868684818110611428576114286141d7565b905060200281019061143a91906141ed565b612cd9565b60010161140a565b505050505061147560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050565b611482611d90565b61148c6000613202565b565b6003818154811061149e57600080fd5b60009182526020909120600290910201546001600160401b0316905081565b6114c5611d90565b6114dc6000805160206148c0833981519152611ed0565b6040517f5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b3990600090a1565b61150f611d90565b611517613273565b6040517f9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e75290600090a1565b3330146115895760405162461bcd60e51b815260206004820152601560248201527427b6b734a837b93a30b61d1037b7363c9039b2b63360591b6044820152606401610c54565b600154600b546001600160401b03908116600160401b90920416146115ea5760405162461bcd60e51b815260206004820152601760248201527627b6b734a837b93a30b61d1037b7363c9031b1b430b4b760491b6044820152606401610c54565b600b54600160401b90046001600160a01b03161561164a5760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a206f6e6c792063636861696e2073656e64657200006044820152606401610c54565b6114758383836124bc565b600254604051632376548f60e21b81526000916001600160a01b031690638dd9523c9061168c908890889088908890600401614236565b602060405180830381865afa1580156116a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116cd919061426e565b95945050505050565b6116de611d90565b611475838383612860565b6116f1611d90565b60405147906001600160a01b0383169082156108fc029083906000818181858888f19350505050158015611729573d6000803e3d6000fd5b50816001600160a01b03167f9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce8260405161176591815260200190565b60405180910390a25050565b611779611d90565b610baf81612168565b61178a611d90565b6117a5610bd560008051602061492083398151915283611e87565b6040516001600160401b038216907fab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa90600090a250565b600061182d6000805160206148e08339815191526000526000805160206149008339815191526020527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff1690565b905090565b61183a611d90565b610baf81612378565b61184b611d90565b610baf816122c4565b61185c611d90565b610baf8161241c565b6000805160206148c083398151915286611883826110848484611e87565b156118c55760405162461bcd60e51b815260206004820152601260248201527113db5b9a541bdc9d185b0e881c185d5cd95960721b6044820152606401610c54565b6001600160401b03881660009081526005602052604090205460ff1661192d5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a20756e737570706f727465642064657374000000006044820152606401610c54565b6001600160a01b0386166119835760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f20706f7274616c207863616c6c00000000006044820152606401610c54565b6000546001600160401b0365010000000000909104811690841611156119eb5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f20686967680000006044820152606401610c54565b6000546001600160401b03600160681b90910481169084161015611a515760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f206c6f77000000006044820152606401610c54565b6000546301000000900461ffff16841115611aae5760405162461bcd60e51b815260206004820152601a60248201527f4f6d6e69506f7274616c3a206461746120746f6f206c617267650000000000006044820152606401610c54565b60ff808816600081815260046020526040902054909116611b115760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a20756e737570706f727465642073686172640000006044820152606401610c54565b6000611b1f8a888888611655565b905080341015611b715760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a20696e73756666696369656e7420666565000000006044820152606401610c54565b6001600160401b03808b166000908152600660209081526040808320868516845290915281208054600193919291611bab9185911661429d565b82546101009290920a6001600160401b038181021990931691831602179091558b811660008181526006602090815260408083208886168085529252918290205491519190931693507fb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa90611c2b9033908e908e908e908e908b906142c4565b60405180910390a450505050505050505050565b611c47611d90565b611c67611c6260008051602061492083398151915283611e87565b61296a565b6040516001600160401b038216907fc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d90600090a250565b611ca6611d90565b611cbd60008051602061492083398151915261296a565b6040517f2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f8287590600090a1565b611cf0611d90565b6114758383836127eb565b611d03611d90565b611d1e611c626000805160206148c083398151915283611e87565b6040516001600160401b038216907f1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e712590600090a250565b611d5d611d90565b6001600160a01b038116611d8757604051631e4fbdf760e01b815260006004820152602401610c54565b610baf81613202565b33611dc27f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461148c5760405163118cdaa760e01b8152336004820152602401610c54565b60008160ff1611611e3e5760405162461bcd60e51b815260206004820152601a60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206375746f66660000000000006044820152606401610c54565b6000805460ff191660ff83169081179091556040519081527f1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f906020015b60405180910390a150565b60008282604051602001611eb292919091825260c01b6001600160c01b031916602082015260280190565b60405160208183030381529060405280519060200120905092915050565b6000818152600080516020614900833981519152602081905260409091205460ff1615611f325760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610c54565b600082815260208290526040808220805460ff191660011790555183917f0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba3191a25050565b611f7e61328a565b3660005b828110156120ea57838382818110611f9c57611f9c6141d7565b9050602002810190611fae919061430f565b6003805460018101825560009190915290925082906002027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01611ff282826143d8565b5050611ffb4690565b6001600160401b03166120116020840184613d32565b6001600160401b03161461205f576001600560006120326020860186613d32565b6001600160401b031681526020810191909152604001600020805460ff19169115159190911790556120e2565b60005b61206f602084018461418e565b90508110156120e05760016004600061208b602087018761418e565b8581811061209b5761209b6141d7565b90506020020160208101906120b09190613d32565b6001600160401b031681526020810191909152604001600020805460ff1916911515919091179055600101612062565b505b600101611f82565b50505050565b6000805160206148e0833981519152600090815260008051602061490083398151915260208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff16806110895750600092835260205250604090205460ff1690565b61215f613389565b610baf816133d2565b6001600160a01b0381166121be5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206665654f7261636c650000006044820152606401610c54565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527fd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd3290602001611e7c565b6000816001600160401b0316116122655760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782067617300000000006044820152606401610c54565b600080546cffffffffffffffff00000000001916650100000000006001600160401b038416908102919091179091556040519081527f1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d90602001611e7c565b6000816001600160401b03161161231d5760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d696e2067617300000000006044820152606401610c54565b6000805467ffffffffffffffff60681b1916600160681b6001600160401b038416908102919091179091556040519081527f8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e90602001611e7c565b60008161ffff16116123cc5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782073697a65000000006044820152606401610c54565b6000805464ffff0000001916630100000061ffff8416908102919091179091556040519081527f65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd206490602001611e7c565b60008161ffff16116124705760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782073697a65000000006044820152606401610c54565b6000805462ffff00191661010061ffff8416908102919091179091556040519081527f620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b090602001611e7c565b808061250a5760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a206e6f2076616c696461746f7273000000000000006044820152606401610c54565b6001600160401b0380851660009081526009602052604090205416156125725760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206475706c69636174652076616c207365740000006044820152606401610c54565b604080518082018252600080825260208083018290526001600160401b0388168252600a9052918220825b8481101561274a578686828181106125b7576125b76141d7565b9050604002018036038101906125cd9190614500565b80519093506001600160a01b03166126275760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206e6f207a65726f2076616c696461746f720000006044820152606401610c54565b600083602001516001600160401b0316116126845760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a206e6f207a65726f20706f776572000000000000006044820152606401610c54565b82516001600160a01b03166000908152602083905260409020546001600160401b0316156126f45760405162461bcd60e51b815260206004820152601f60248201527f4f6d6e69506f7274616c3a206475706c69636174652076616c696461746f72006044820152606401610c54565b6020830151612703908561429d565b60208481015185516001600160a01b03166000908152918590526040909120805467ffffffffffffffff19166001600160401b03909216919091179055935060010161259d565b506001600160401b038781166000818152600960205260408120805467ffffffffffffffff191687851617905554600160a81b900490911610156127ae576000805467ffffffffffffffff60a81b1916600160a81b6001600160401b038a16021790555b6040516001600160401b038816907f3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd10790600090a250505050505050565b6001600160401b03838116600081815260076020908152604080832087861680855290835292819020805467ffffffffffffffff191695871695861790555193845290927f8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c791015b60405180910390a3505050565b6001600160401b03838116600081815260086020908152604080832087861680855290835292819020805467ffffffffffffffff191695871695861790555193845290927fe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b59101612853565b61148c6000805160206148e083398151915261296a565b6000805160206148e0833981519152600090815260008051602061490083398151915260208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff1680612949575060008481526020829052604090205460ff165b80612962575060008381526020829052604090205460ff165b949350505050565b6000818152600080516020614900833981519152602081905260409091205460ff166129cf5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610c54565b600082815260208290526040808220805460ff191690555183917fd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a91a25050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901612a5457604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000805460ff8116600160a81b9091046001600160401b031611612a7e5750600190565b600054612a9f9060ff811690600160a81b90046001600160401b031661455c565b61182d90600161429d565b6000803660005b88811015612c4b57898982818110612acb57612acb6141d7565b9050602002810190612add919061430f565b91508015612b9157368a8a612af360018561457c565b818110612b0257612b026141d7565b9050602002810190612b14919061430f565b9050612b23602082018261407d565b6001600160a01b0316612b39602085018561407d565b6001600160a01b031611612b8f5760405162461bcd60e51b815260206004820152601f60248201527f51756f72756d3a2073696773206e6f7420646564757065642f736f72746564006044820152606401610c54565b505b612b9b828c6133da565b612be75760405162461bcd60e51b815260206004820152601960248201527f51756f72756d3a20696e76616c6964207369676e6174757265000000000000006044820152606401610c54565b876000612bf7602085018561407d565b6001600160a01b03168152602081019190915260400160002054612c24906001600160401b03168461429d565b9250612c328388888861344e565b15612c435760019350505050612c53565b600101612ab1565b506000925050505b979650505050505050565b60408051600180825281830190925260009182919060208083019080368337019050509050612c9986868686612c948d8d613486565b613553565b81600081518110612cac57612cac6141d7565b602002602001018181525050612ccb818b612cc68c6137b4565b6137cc565b9a9950505050505050505050565b6000612ce86020840184613d32565b90506000612cf96020840184613d32565b90506000612d0d6040850160208601613d32565b90506000612d216060860160408701613d32565b9050466001600160401b0316836001600160401b03161480612d4a57506001600160401b038316155b612d965760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a2077726f6e67206465737420636861696e000000006044820152606401610c54565b6001600160401b0380851660009081526007602090815260408083208685168452909152902054612dc99116600161429d565b6001600160401b0316816001600160401b031614612e295760405162461bcd60e51b815260206004820152601860248201527f4f6d6e69506f7274616c3a2077726f6e67206f666673657400000000000000006044820152606401610c54565b612e396060870160408801613cf7565b60ff16600460ff161480612e61575060ff8216612e5c6060880160408901613cf7565b60ff16145b612ead5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a2077726f6e6720636f6e66206c6576656c000000006044820152606401610c54565b612ebd6080870160608801613d32565b6001600160401b038581166000908152600860209081526040808320878516845290915290205491811691161015612f3d57612eff6080870160608801613d32565b6001600160401b03858116600090815260086020908152604080832087851684529091529020805467ffffffffffffffff1916929091169190911790555b6001600160401b038085166000908152600760209081526040808320868516845290915281208054600193919291612f779185911661429d565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550306001600160a01b0316856080016020810190612fb8919061407d565b6001600160a01b03160361309257806001600160401b0316826001600160401b0316856001600160401b03167f8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b600033600060405160240161304b906020808252601e908201527f4f6d6e69506f7274616c3a206e6f207863616c6c20746f20706f7274616c0000604082015260600190565b60408051601f198184030181529181526020820180516001600160e01b031662461bcd60e51b1790525161308294939291906145df565b60405180910390a4505050505050565b604080518082019091526001600160401b0385168152602081016130bc608088016060890161407d565b6001600160a01b039081169091528151600b8054602090940151909216600160401b026001600160e01b03199093166001600160401b0390911617919091179055600080808061311260a08a0160808b0161407d565b6001600160a01b0316146131635761315e61313360a08a0160808b0161407d565b61314360e08b0160c08c01613d32565b6001600160401b031661315960a08c018c61461b565b6137e2565b613178565b61317861317360a08a018a61461b565b6138a2565b600b80546001600160e01b03191690559194509250905060008361319c57826131ad565b604051806020016040528060008152505b9050846001600160401b0316866001600160401b0316896001600160401b03167f8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b85338987604051611c2b94939291906145df565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b61148c6000805160206148e0833981519152611ed0565b6000805b60035481101561337c57600381815481106132ab576132ab6141d7565b906000526020600020906002020191506132c24690565b82546001600160401b039081169116146132fc5781546001600160401b03166000908152600560205260409020805460ff19169055613374565b60005b600183015481101561337257600060046000856001018481548110613326576133266141d7565b6000918252602080832060048304015460039092166008026101000a9091046001600160401b031683528201929092526040019020805460ff19169115159190911790556001016132ff565b505b60010161328e565b50610baf60036000613c5c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661148c57604051631afcd79f60e31b815260040160405180910390fd5b611d5d613389565b60006133e9602084018461407d565b6001600160a01b031661343d83613403602087018761461b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061393992505050565b6001600160a01b0316149392505050565b600061346660ff84166001600160401b038616614348565b61347c60ff84166001600160401b038816614348565b1195945050505050565b60606000826001600160401b038111156134a2576134a2614332565b6040519080825280602002602001820160405280156134cb578160200160208202803683370190505b50905060005b8381101561354b5761352660028686848181106134f0576134f06141d7565b905060200281019061350291906141ed565b60405160200161351291906146a6565b604051602081830303815290604052613963565b828281518110613538576135386141d7565b60209081029190910101526001016134d1565b509392505050565b80516000908584613565816001614775565b61356f8385614775565b1461358d57604051631a8a024960e11b815260040160405180910390fd5b6000816001600160401b038111156135a7576135a7614332565b6040519080825280602002602001820160405280156135d0578160200160208202803683370190505b5090506000806000805b8581101561371d5760008885106136155785846135f681614788565b955081518110613608576136086141d7565b602002602001015161363b565b8a8561362081614788565b965081518110613632576136326141d7565b60200260200101515b905060008d8d84818110613651576136516141d7565b905060200201602081019061366691906147a1565b613693578f8f8561367681614788565b9650818110613687576136876141d7565b905060200201356136ea565b8986106136c45786856136a581614788565b9650815181106136b7576136b76141d7565b60200260200101516136ea565b8b866136cf81614788565b9750815181106136e1576136e16141d7565b60200260200101515b90506136f6828261399a565b878481518110613708576137086141d7565b602090810291909101015250506001016135da565b50841561376f5785811461374457604051631a8a024960e11b815260040160405180910390fd5b836001860381518110613759576137596141d7565b60200260200101519750505050505050506116cd565b86156137885788600081518110613759576137596141d7565b8c8c600081811061379b5761379b6141d7565b9050602002013597505050505050505095945050505050565b6000610d8160018360405160200161351291906147c3565b6000826137d985846139c9565b14949350505050565b600060606000805a90506000806138658960008060019054906101000a900461ffff168b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e6001600160a01b0316613a0490949392919063ffffffff16565b9150915060005a9050613879603f8b614848565b811161388157fe5b828261388d838761457c565b965096509650505050505b9450945094915050565b600060606000805a9050600080306001600160a01b031688886040516138c992919061486a565b6000604051808303816000865af19150503d8060008114613906576040519150601f19603f3d011682016040523d82523d6000602084013e61390b565b606091505b50915091505a61391b908461457c565b92508161392a57805160208201fd5b909450925090505b9250925092565b6000806000806139498686613a8e565b9250925092506139598282613ad8565b5090949350505050565b6000828260405160200161397892919061487a565b60408051601f1981840301815282825280516020918201209083015201611eb2565b60008183106139b6576000828152602084905260409020611089565b6000838152602083905260409020611089565b600081815b845181101561354b576139fa828683815181106139ed576139ed6141d7565b602002602001015161399a565b91506001016139ce565b6000606060008060008661ffff166001600160401b03811115613a2957613a29614332565b6040519080825280601f01601f191660200182016040528015613a53576020820181803683370190505b5090506000808751602089018b8e8ef191503d925086831115613a74578692505b828152826000602083013e90999098509650505050505050565b60008060008351604103613ac85760208401516040850151606086015160001a613aba88828585613b91565b955095509550505050613932565b5050815160009150600290613932565b6000826003811115613aec57613aec6148a9565b03613af5575050565b6001826003811115613b0957613b096148a9565b03613b275760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613b3b57613b3b6148a9565b03613b5c5760405163fce698f760e01b815260048101829052602401610c54565b6003826003811115613b7057613b706148a9565b03610d28576040516335e2f38360e21b815260048101829052602401610c54565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613bcc5750600091506003905082613898565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613c20573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613c4c57506000925060019150829050613898565b9760009750879650945050505050565b5080546000825560020290600052602060002090810190610baf91905b80821115613ca857805467ffffffffffffffff191681556000613c9f6001830182613cac565b50600201613c79565b5090565b508054600082556003016004900490600052602060002090810190610baf91905b80821115613ca85760008155600101613ccd565b803560ff81168114613cf257600080fd5b919050565b600060208284031215613d0957600080fd5b61108982613ce1565b6001600160401b0381168114610baf57600080fd5b8035613cf281613d12565b600060208284031215613d4457600080fd5b813561108981613d12565b60008060208385031215613d6257600080fd5b82356001600160401b0380821115613d7957600080fd5b818501915085601f830112613d8d57600080fd5b813581811115613d9c57600080fd5b8660208260051b8501011115613db157600080fd5b60209290920196919550909350505050565b600060208284031215613dd557600080fd5b5035919050565b600060208284031215613dee57600080fd5b81356001600160401b03811115613e0457600080fd5b82016101a0818503121561108957600080fd5b60008060408385031215613e2a57600080fd5b8235613e3581613d12565b91506020830135613e4581613d12565b809150509250929050565b60008060408385031215613e6357600080fd5b823591506020830135613e4581613d12565b80356001600160a01b0381168114613cf257600080fd5b60008060408385031215613e9f57600080fd5b8235613eaa81613d12565b9150613eb860208401613e75565b90509250929050565b600060208284031215613ed357600080fd5b81356001600160401b03811115613ee957600080fd5b8201610180818503121561108957600080fd5b600080600060408486031215613f1157600080fd5b8335613f1c81613d12565b925060208401356001600160401b0380821115613f3857600080fd5b818601915086601f830112613f4c57600080fd5b813581811115613f5b57600080fd5b8760208260061b8501011115613f7057600080fd5b6020830194508093505050509250925092565b60008083601f840112613f9557600080fd5b5081356001600160401b03811115613fac57600080fd5b602083019150836020828501011115613fc457600080fd5b9250929050565b60008060008060608587031215613fe157600080fd5b8435613fec81613d12565b935060208501356001600160401b0381111561400757600080fd5b61401387828801613f83565b909450925050604085013561402781613d12565b939692955090935050565b60008060006060848603121561404757600080fd5b833561405281613d12565b9250602084013561406281613d12565b9150604084013561407281613d12565b809150509250925092565b60006020828403121561408f57600080fd5b61108982613e75565b6000602082840312156140aa57600080fd5b813561ffff8116811461108957600080fd5b60008060008060008060a087890312156140d557600080fd5b86356140e081613d12565b95506140ee60208801613ce1565b94506140fc60408801613e75565b935060608701356001600160401b0381111561411757600080fd5b61412389828a01613f83565b909450925050608087013561413781613d12565b809150509295509295509295565b6000808335601e1984360301811261415c57600080fd5b8301803591506001600160401b0382111561417657600080fd5b6020019150600681901b3603821315613fc457600080fd5b6000808335601e198436030181126141a557600080fd5b8301803591506001600160401b038211156141bf57600080fd5b6020019150600581901b3603821315613fc457600080fd5b634e487b7160e01b600052603260045260246000fd5b6000823560de1983360301811261420357600080fd5b9190910192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006001600160401b0380871683526060602084015261425a60608401868861420d565b915080841660408401525095945050505050565b60006020828403121561428057600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b6001600160401b038181168382160190808211156142bd576142bd614287565b5092915050565b6001600160a01b0387811682528616602082015260a0604082018190526000906142f1908301868861420d565b6001600160401b039490941660608301525060800152949350505050565b60008235603e1983360301811261420357600080fd5b60008135610d8181613d12565b634e487b7160e01b600052604160045260246000fd5b8082028115828204841417610d8157610d81614287565b600160401b82111561437357614373614332565b8054828255808310156114755760008260005260206000206003850160021c81016003840160021c8201915060188660031b1680156143c3576000198083018054828460200360031b1c16815550505b505b81811015611020578281556001016143c5565b81356143e381613d12565b815467ffffffffffffffff19166001600160401b0391821617825560019081830160208581013536879003601e1901811261441d57600080fd5b860180358481111561442e57600080fd5b6020820194508060051b360385131561444657600080fd5b614450818561435f565b60009384526020842093600282901c92505b828110156144b9576000805b60048110156144ad576144a061448389614325565b6001600160401b03908116600684901b90811b91901b1984161790565b978601979150880161446e565b50858201558601614462565b5060031981168082038183146144f4576000805b828110156144ee576144e16144838a614325565b98870198915089016144cd565b50868501555b50505050505050505050565b60006040828403121561451257600080fd5b604051604081018181106001600160401b038211171561453457614534614332565b60405261454083613e75565b8152602083013561455081613d12565b60208201529392505050565b6001600160401b038281168282160390808211156142bd576142bd614287565b81810381811115610d8157610d81614287565b60005b838110156145aa578181015183820152602001614592565b50506000910152565b600081518084526145cb81602086016020860161458f565b601f01601f19169290920160200192915050565b8481526001600160a01b03841660208201528215156040820152608060608201819052600090614611908301846145b3565b9695505050505050565b6000808335601e1984360301811261463257600080fd5b8301803591506001600160401b0382111561464c57600080fd5b602001915036819003821315613fc457600080fd5b6000808335601e1984360301811261467857600080fd5b83016020810192503590506001600160401b0381111561469757600080fd5b803603821315613fc457600080fd5b60208152600082356146b781613d12565b6001600160401b038082166020850152602085013591506146d782613d12565b8082166040850152604085013591506146ef82613d12565b166060838101919091526001600160a01b039061470d908501613e75565b16608083015261471f60808401613e75565b6001600160a01b03811660a08401525061473c60a0840184614661565b60e060c08501526147526101008501828461420d565b91505061476160c08501613d27565b6001600160401b03811660e085015261354b565b80820180821115610d8157610d81614287565b60006001820161479a5761479a614287565b5060010190565b6000602082840312156147b357600080fd5b8135801515811461108957600080fd5b60c0810182356147d281613d12565b6001600160401b0390811683526020840135906147ee82613d12565b808216602085015260ff61480460408701613ce1565b1660408501526060850135915061481a82613d12565b908116606084015260808401359061483182613d12565b16608083015260a092830135929091019190915290565b60008261486557634e487b7160e01b600052601260045260246000fd5b500490565b8183823760009101908152919050565b60ff60f81b8360f81b1681526000825161489b81600185016020870161458f565b919091016001019392505050565b634e487b7160e01b600052602160045260246000fdfea06a0c1264badca141841b5f52470407dac9adaaa539dd445540986341b73a6876e8952e4b09b8d505aa08998d716721a1dbf0884ac74202e33985da1ed005e9ff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400feccba1cfc4544bf9cd83b76f36ae5c464750b6c43f682e26744ee21ec31fc1ea2646970667358221220c5e3bace3aca4ee5393630aeb4d73670e9bf89a7151df46be819f199c38fab1664736f6c63430008180033",
}

// OmniPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniPortalMetaData.ABI instead.
var OmniPortalABI = OmniPortalMetaData.ABI

// OmniPortalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniPortalMetaData.Bin instead.
var OmniPortalBin = OmniPortalMetaData.Bin

// DeployOmniPortal deploys a new Ethereum contract, binding an instance of OmniPortal to it.
func DeployOmniPortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniPortal, error) {
	parsed, err := OmniPortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniPortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// OmniPortal is an auto generated Go binding around an Ethereum contract.
type OmniPortal struct {
	OmniPortalCaller     // Read-only binding to the contract
	OmniPortalTransactor // Write-only binding to the contract
	OmniPortalFilterer   // Log filterer for contract events
}

// OmniPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniPortalSession struct {
	Contract     *OmniPortal       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniPortalCallerSession struct {
	Contract *OmniPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OmniPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniPortalTransactorSession struct {
	Contract     *OmniPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OmniPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniPortalRaw struct {
	Contract *OmniPortal // Generic contract binding to access the raw methods on
}

// OmniPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniPortalCallerRaw struct {
	Contract *OmniPortalCaller // Generic read-only contract binding to access the raw methods on
}

// OmniPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniPortalTransactorRaw struct {
	Contract *OmniPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniPortal creates a new instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortal(address common.Address, backend bind.ContractBackend) (*OmniPortal, error) {
	contract, err := bindOmniPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// NewOmniPortalCaller creates a new read-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalCaller(address common.Address, caller bind.ContractCaller) (*OmniPortalCaller, error) {
	contract, err := bindOmniPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalCaller{contract: contract}, nil
}

// NewOmniPortalTransactor creates a new write-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniPortalTransactor, error) {
	contract, err := bindOmniPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalTransactor{contract: contract}, nil
}

// NewOmniPortalFilterer creates a new log filterer instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniPortalFilterer, error) {
	contract, err := bindOmniPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFilterer{contract: contract}, nil
}

// bindOmniPortal binds a generic wrapper to an already deployed contract.
func bindOmniPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.OmniPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transact(opts, method, params...)
}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) ActionXCall(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "ActionXCall")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) ActionXCall() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXCall(&_OmniPortal.CallOpts)
}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) ActionXCall() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXCall(&_OmniPortal.CallOpts)
}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) ActionXSubmit(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "ActionXSubmit")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) ActionXSubmit() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXSubmit(&_OmniPortal.CallOpts)
}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) ActionXSubmit() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXSubmit(&_OmniPortal.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) KeyPauseAll(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "KeyPauseAll")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) KeyPauseAll() ([32]byte, error) {
	return _OmniPortal.Contract.KeyPauseAll(&_OmniPortal.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) KeyPauseAll() ([32]byte, error) {
	return _OmniPortal.Contract.KeyPauseAll(&_OmniPortal.CallOpts)
}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XSubQuorumDenominator(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XSubQuorumDenominator")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XSubQuorumDenominator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumDenominator(&_OmniPortal.CallOpts)
}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XSubQuorumDenominator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumDenominator(&_OmniPortal.CallOpts)
}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XSubQuorumNumerator(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XSubQuorumNumerator")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XSubQuorumNumerator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumNumerator(&_OmniPortal.CallOpts)
}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XSubQuorumNumerator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumNumerator(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalCaller) FeeFor(opts *bind.CallOpts, destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeFor", destChainId, data, gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor(&_OmniPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalCallerSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor(&_OmniPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalCaller) FeeOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalSession) FeeOracle() (common.Address, error) {
	return _OmniPortal.Contract.FeeOracle(&_OmniPortal.CallOpts)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalCallerSession) FeeOracle() (common.Address, error) {
	return _OmniPortal.Contract.FeeOracle(&_OmniPortal.CallOpts)
}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXBlockOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXBlockOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXBlockOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXBlockOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXBlockOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXBlockOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXMsgOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXMsgOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused(opts *bind.CallOpts, actionId [32]byte) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused", actionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused(actionId [32]byte) (bool, error) {
	return _OmniPortal.Contract.IsPaused(&_OmniPortal.CallOpts, actionId)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused(actionId [32]byte) (bool, error) {
	return _OmniPortal.Contract.IsPaused(&_OmniPortal.CallOpts, actionId)
}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused0(opts *bind.CallOpts, actionId [32]byte, chainId_ uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused0", actionId, chainId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused0(actionId [32]byte, chainId_ uint64) (bool, error) {
	return _OmniPortal.Contract.IsPaused0(&_OmniPortal.CallOpts, actionId, chainId_)
}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused0(actionId [32]byte, chainId_ uint64) (bool, error) {
	return _OmniPortal.Contract.IsPaused0(&_OmniPortal.CallOpts, actionId, chainId_)
}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused1() (bool, error) {
	return _OmniPortal.Contract.IsPaused1(&_OmniPortal.CallOpts)
}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused1() (bool, error) {
	return _OmniPortal.Contract.IsPaused1(&_OmniPortal.CallOpts)
}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsSupportedDest(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isSupportedDest", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsSupportedDest(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedDest(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsSupportedDest(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedDest(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsSupportedShard(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isSupportedShard", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsSupportedShard(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedShard(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsSupportedShard(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedShard(&_OmniPortal.CallOpts, arg0)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsXCall(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isXCall")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalSession) IsXCall() (bool, error) {
	return _OmniPortal.Contract.IsXCall(&_OmniPortal.CallOpts)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsXCall() (bool, error) {
	return _OmniPortal.Contract.IsXCall(&_OmniPortal.CallOpts)
}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) LatestValSetId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "latestValSetId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) LatestValSetId() (uint64, error) {
	return _OmniPortal.Contract.LatestValSetId(&_OmniPortal.CallOpts)
}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) LatestValSetId() (uint64, error) {
	return _OmniPortal.Contract.LatestValSetId(&_OmniPortal.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalCaller) Network(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "network", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalSession) Network(arg0 *big.Int) (uint64, error) {
	return _OmniPortal.Contract.Network(&_OmniPortal.CallOpts, arg0)
}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalCallerSession) Network(arg0 *big.Int) (uint64, error) {
	return _OmniPortal.Contract.Network(&_OmniPortal.CallOpts, arg0)
}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OmniCChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "omniCChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) OmniCChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniCChainId(&_OmniPortal.CallOpts)
}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OmniCChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniCChainId(&_OmniPortal.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OmniChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "omniChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) OmniChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniChainId(&_OmniPortal.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OmniChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniChainId(&_OmniPortal.CallOpts)
}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OutXMsgOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "outXMsgOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) OutXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OutXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalSession) Owner() (common.Address, error) {
	return _OmniPortal.Contract.Owner(&_OmniPortal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalCallerSession) Owner() (common.Address, error) {
	return _OmniPortal.Contract.Owner(&_OmniPortal.CallOpts)
}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ValSet(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "valSet", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) ValSet(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _OmniPortal.Contract.ValSet(&_OmniPortal.CallOpts, arg0, arg1)
}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ValSet(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _OmniPortal.Contract.ValSet(&_OmniPortal.CallOpts, arg0, arg1)
}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ValSetTotalPower(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "valSetTotalPower", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) ValSetTotalPower(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.ValSetTotalPower(&_OmniPortal.CallOpts, arg0)
}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ValSetTotalPower(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.ValSetTotalPower(&_OmniPortal.CallOpts, arg0)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalCaller) Xmsg(opts *bind.CallOpts) (XTypesMsgContext, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsg")

	if err != nil {
		return *new(XTypesMsgContext), err
	}

	out0 := *abi.ConvertType(out[0], new(XTypesMsgContext)).(*XTypesMsgContext)

	return out0, err

}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalSession) Xmsg() (XTypesMsgContext, error) {
	return _OmniPortal.Contract.Xmsg(&_OmniPortal.CallOpts)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalCallerSession) Xmsg() (XTypesMsgContext, error) {
	return _OmniPortal.Contract.Xmsg(&_OmniPortal.CallOpts)
}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalCaller) XmsgMaxDataSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMaxDataSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalSession) XmsgMaxDataSize() (uint16, error) {
	return _OmniPortal.Contract.XmsgMaxDataSize(&_OmniPortal.CallOpts)
}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalCallerSession) XmsgMaxDataSize() (uint16, error) {
	return _OmniPortal.Contract.XmsgMaxDataSize(&_OmniPortal.CallOpts)
}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XmsgMaxGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMaxGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XmsgMaxGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMaxGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XmsgMaxGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMaxGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XmsgMinGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMinGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XmsgMinGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMinGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XmsgMinGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMinGasLimit(&_OmniPortal.CallOpts)
}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalCaller) XreceiptMaxErrorSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xreceiptMaxErrorSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalSession) XreceiptMaxErrorSize() (uint16, error) {
	return _OmniPortal.Contract.XreceiptMaxErrorSize(&_OmniPortal.CallOpts)
}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalCallerSession) XreceiptMaxErrorSize() (uint16, error) {
	return _OmniPortal.Contract.XreceiptMaxErrorSize(&_OmniPortal.CallOpts)
}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XsubValsetCutoff(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xsubValsetCutoff")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XsubValsetCutoff() (uint8, error) {
	return _OmniPortal.Contract.XsubValsetCutoff(&_OmniPortal.CallOpts)
}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XsubValsetCutoff() (uint8, error) {
	return _OmniPortal.Contract.XsubValsetCutoff(&_OmniPortal.CallOpts)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalTransactor) AddValidatorSet(opts *bind.TransactOpts, valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "addValidatorSet", valSetId, validators)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalSession) AddValidatorSet(valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.Contract.AddValidatorSet(&_OmniPortal.TransactOpts, valSetId, validators)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalTransactorSession) AddValidatorSet(valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.Contract.AddValidatorSet(&_OmniPortal.TransactOpts, valSetId, validators)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalTransactor) CollectFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "collectFees", to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.CollectFees(&_OmniPortal.TransactOpts, to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalTransactorSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.CollectFees(&_OmniPortal.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalTransactor) Initialize(opts *bind.TransactOpts, p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "initialize", p)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalSession) Initialize(p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.Contract.Initialize(&_OmniPortal.TransactOpts, p)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalTransactorSession) Initialize(p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.Contract.Initialize(&_OmniPortal.TransactOpts, p)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalSession) Pause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Pause(&_OmniPortal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalTransactorSession) Pause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Pause(&_OmniPortal.TransactOpts)
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalTransactor) PauseXCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXCall")
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalSession) PauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCall(&_OmniPortal.TransactOpts)
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCall(&_OmniPortal.TransactOpts)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) PauseXCallTo(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXCallTo", chainId_)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) PauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactor) PauseXSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXSubmit")
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalSession) PauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmit(&_OmniPortal.TransactOpts)
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmit(&_OmniPortal.TransactOpts)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) PauseXSubmitFrom(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXSubmitFrom", chainId_)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) PauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniPortal.Contract.RenounceOwnership(&_OmniPortal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniPortal.Contract.RenounceOwnership(&_OmniPortal.TransactOpts)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalTransactor) SetFeeOracle(opts *bind.TransactOpts, feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setFeeOracle", feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetFeeOracle(&_OmniPortal.TransactOpts, feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetFeeOracle(&_OmniPortal.TransactOpts, feeOracle_)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactor) SetInXBlockOffset(opts *bind.TransactOpts, sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setInXBlockOffset", sourceChainId, shardId, offset)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalSession) SetInXBlockOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXBlockOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetInXBlockOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXBlockOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactor) SetInXMsgOffset(opts *bind.TransactOpts, sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setInXMsgOffset", sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalSession) SetInXMsgOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXMsgOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetInXMsgOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXMsgOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalTransactor) SetNetwork(opts *bind.TransactOpts, network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setNetwork", network_)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalSession) SetNetwork(network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetNetwork(&_OmniPortal.TransactOpts, network_)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetNetwork(network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetNetwork(&_OmniPortal.TransactOpts, network_)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMaxDataSize(opts *bind.TransactOpts, numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMaxDataSize", numBytes)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMaxDataSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxDataSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMaxDataSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxDataSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMaxGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMaxGasLimit", gasLimit)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMaxGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMaxGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMinGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMinGasLimit", gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMinGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMinGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMinGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMinGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactor) SetXReceiptMaxErrorSize(opts *bind.TransactOpts, numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXReceiptMaxErrorSize", numBytes)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalSession) SetXReceiptMaxErrorSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXReceiptMaxErrorSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXReceiptMaxErrorSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXReceiptMaxErrorSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalTransactor) SetXSubValsetCutoff(opts *bind.TransactOpts, xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXSubValsetCutoff", xsubValsetCutoff_)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalSession) SetXSubValsetCutoff(xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXSubValsetCutoff(&_OmniPortal.TransactOpts, xsubValsetCutoff_)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXSubValsetCutoff(xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXSubValsetCutoff(&_OmniPortal.TransactOpts, xsubValsetCutoff_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.TransferOwnership(&_OmniPortal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.TransferOwnership(&_OmniPortal.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalSession) Unpause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Unpause(&_OmniPortal.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalTransactorSession) Unpause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Unpause(&_OmniPortal.TransactOpts)
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXCall")
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalSession) UnpauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCall(&_OmniPortal.TransactOpts)
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCall(&_OmniPortal.TransactOpts)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXCallTo(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXCallTo", chainId_)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) UnpauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXSubmit")
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalSession) UnpauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmit(&_OmniPortal.TransactOpts)
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmit(&_OmniPortal.TransactOpts)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXSubmitFrom(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXSubmitFrom", chainId_)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) UnpauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactor) Xcall(opts *bind.TransactOpts, destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xcall", destChainId, conf, to, data, gasLimit)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalSession) Xcall(destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, conf, to, data, gasLimit)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactorSession) Xcall(destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, conf, to, data, gasLimit)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactor) Xsubmit(opts *bind.TransactOpts, xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xsubmit", xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactorSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// OmniPortalFeeOracleSetIterator is returned from FilterFeeOracleSet and is used to iterate over the raw logs and unpacked data for FeeOracleSet events raised by the OmniPortal contract.
type OmniPortalFeeOracleSetIterator struct {
	Event *OmniPortalFeeOracleSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalFeeOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalFeeOracleSet)
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
		it.Event = new(OmniPortalFeeOracleSet)
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
func (it *OmniPortalFeeOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalFeeOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalFeeOracleSet represents a FeeOracleSet event raised by the OmniPortal contract.
type OmniPortalFeeOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeOracleSet is a free log retrieval operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_OmniPortal *OmniPortalFilterer) FilterFeeOracleSet(opts *bind.FilterOpts) (*OmniPortalFeeOracleSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "FeeOracleSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalFeeOracleSetIterator{contract: _OmniPortal.contract, event: "FeeOracleSet", logs: logs, sub: sub}, nil
}

// WatchFeeOracleSet is a free log subscription operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_OmniPortal *OmniPortalFilterer) WatchFeeOracleSet(opts *bind.WatchOpts, sink chan<- *OmniPortalFeeOracleSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "FeeOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalFeeOracleSet)
				if err := _OmniPortal.contract.UnpackLog(event, "FeeOracleSet", log); err != nil {
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

// ParseFeeOracleSet is a log parse operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_OmniPortal *OmniPortalFilterer) ParseFeeOracleSet(log types.Log) (*OmniPortalFeeOracleSet, error) {
	event := new(OmniPortalFeeOracleSet)
	if err := _OmniPortal.contract.UnpackLog(event, "FeeOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalFeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the OmniPortal contract.
type OmniPortalFeesCollectedIterator struct {
	Event *OmniPortalFeesCollected // Event containing the contract specifics and raw log

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
func (it *OmniPortalFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalFeesCollected)
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
		it.Event = new(OmniPortalFeesCollected)
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
func (it *OmniPortalFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalFeesCollected represents a FeesCollected event raised by the OmniPortal contract.
type OmniPortalFeesCollected struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) FilterFeesCollected(opts *bind.FilterOpts, to []common.Address) (*OmniPortalFeesCollectedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFeesCollectedIterator{contract: _OmniPortal.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *OmniPortalFeesCollected, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalFeesCollected)
				if err := _OmniPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
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

// ParseFeesCollected is a log parse operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) ParseFeesCollected(log types.Log) (*OmniPortalFeesCollected, error) {
	event := new(OmniPortalFeesCollected)
	if err := _OmniPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInXBlockOffsetSetIterator is returned from FilterInXBlockOffsetSet and is used to iterate over the raw logs and unpacked data for InXBlockOffsetSet events raised by the OmniPortal contract.
type OmniPortalInXBlockOffsetSetIterator struct {
	Event *OmniPortalInXBlockOffsetSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalInXBlockOffsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInXBlockOffsetSet)
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
		it.Event = new(OmniPortalInXBlockOffsetSet)
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
func (it *OmniPortalInXBlockOffsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInXBlockOffsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInXBlockOffsetSet represents a InXBlockOffsetSet event raised by the OmniPortal contract.
type OmniPortalInXBlockOffsetSet struct {
	SrcChainId uint64
	ShardId    uint64
	Offset     uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInXBlockOffsetSet is a free log retrieval operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) FilterInXBlockOffsetSet(opts *bind.FilterOpts, srcChainId []uint64, shardId []uint64) (*OmniPortalInXBlockOffsetSetIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "InXBlockOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalInXBlockOffsetSetIterator{contract: _OmniPortal.contract, event: "InXBlockOffsetSet", logs: logs, sub: sub}, nil
}

// WatchInXBlockOffsetSet is a free log subscription operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) WatchInXBlockOffsetSet(opts *bind.WatchOpts, sink chan<- *OmniPortalInXBlockOffsetSet, srcChainId []uint64, shardId []uint64) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "InXBlockOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInXBlockOffsetSet)
				if err := _OmniPortal.contract.UnpackLog(event, "InXBlockOffsetSet", log); err != nil {
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

// ParseInXBlockOffsetSet is a log parse operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) ParseInXBlockOffsetSet(log types.Log) (*OmniPortalInXBlockOffsetSet, error) {
	event := new(OmniPortalInXBlockOffsetSet)
	if err := _OmniPortal.contract.UnpackLog(event, "InXBlockOffsetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInXMsgOffsetSetIterator is returned from FilterInXMsgOffsetSet and is used to iterate over the raw logs and unpacked data for InXMsgOffsetSet events raised by the OmniPortal contract.
type OmniPortalInXMsgOffsetSetIterator struct {
	Event *OmniPortalInXMsgOffsetSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalInXMsgOffsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInXMsgOffsetSet)
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
		it.Event = new(OmniPortalInXMsgOffsetSet)
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
func (it *OmniPortalInXMsgOffsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInXMsgOffsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInXMsgOffsetSet represents a InXMsgOffsetSet event raised by the OmniPortal contract.
type OmniPortalInXMsgOffsetSet struct {
	SrcChainId uint64
	ShardId    uint64
	Offset     uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInXMsgOffsetSet is a free log retrieval operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) FilterInXMsgOffsetSet(opts *bind.FilterOpts, srcChainId []uint64, shardId []uint64) (*OmniPortalInXMsgOffsetSetIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "InXMsgOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalInXMsgOffsetSetIterator{contract: _OmniPortal.contract, event: "InXMsgOffsetSet", logs: logs, sub: sub}, nil
}

// WatchInXMsgOffsetSet is a free log subscription operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) WatchInXMsgOffsetSet(opts *bind.WatchOpts, sink chan<- *OmniPortalInXMsgOffsetSet, srcChainId []uint64, shardId []uint64) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "InXMsgOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInXMsgOffsetSet)
				if err := _OmniPortal.contract.UnpackLog(event, "InXMsgOffsetSet", log); err != nil {
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

// ParseInXMsgOffsetSet is a log parse operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) ParseInXMsgOffsetSet(log types.Log) (*OmniPortalInXMsgOffsetSet, error) {
	event := new(OmniPortalInXMsgOffsetSet)
	if err := _OmniPortal.contract.UnpackLog(event, "InXMsgOffsetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniPortal contract.
type OmniPortalInitializedIterator struct {
	Event *OmniPortalInitialized // Event containing the contract specifics and raw log

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
func (it *OmniPortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInitialized)
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
		it.Event = new(OmniPortalInitialized)
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
func (it *OmniPortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInitialized represents a Initialized event raised by the OmniPortal contract.
type OmniPortalInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniPortal *OmniPortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniPortalInitializedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniPortalInitializedIterator{contract: _OmniPortal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniPortal *OmniPortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniPortalInitialized) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInitialized)
				if err := _OmniPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniPortal *OmniPortalFilterer) ParseInitialized(log types.Log) (*OmniPortalInitialized, error) {
	event := new(OmniPortalInitialized)
	if err := _OmniPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniPortal contract.
type OmniPortalOwnershipTransferredIterator struct {
	Event *OmniPortalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniPortalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalOwnershipTransferred)
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
		it.Event = new(OmniPortalOwnershipTransferred)
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
func (it *OmniPortalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalOwnershipTransferred represents a OwnershipTransferred event raised by the OmniPortal contract.
type OmniPortalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniPortalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalOwnershipTransferredIterator{contract: _OmniPortal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniPortalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalOwnershipTransferred)
				if err := _OmniPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) ParseOwnershipTransferred(log types.Log) (*OmniPortalOwnershipTransferred, error) {
	event := new(OmniPortalOwnershipTransferred)
	if err := _OmniPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OmniPortal contract.
type OmniPortalPausedIterator struct {
	Event *OmniPortalPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalPaused)
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
		it.Event = new(OmniPortalPaused)
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
func (it *OmniPortalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalPaused represents a Paused event raised by the OmniPortal contract.
type OmniPortalPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_OmniPortal *OmniPortalFilterer) FilterPaused(opts *bind.FilterOpts) (*OmniPortalPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalPausedIterator{contract: _OmniPortal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_OmniPortal *OmniPortalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_OmniPortal *OmniPortalFilterer) ParsePaused(log types.Log) (*OmniPortalPaused, error) {
	event := new(OmniPortalPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalPaused0Iterator is returned from FilterPaused0 and is used to iterate over the raw logs and unpacked data for Paused0 events raised by the OmniPortal contract.
type OmniPortalPaused0Iterator struct {
	Event *OmniPortalPaused0 // Event containing the contract specifics and raw log

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
func (it *OmniPortalPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalPaused0)
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
		it.Event = new(OmniPortalPaused0)
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
func (it *OmniPortalPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalPaused0 represents a Paused0 event raised by the OmniPortal contract.
type OmniPortalPaused0 struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused0 is a free log retrieval operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_OmniPortal *OmniPortalFilterer) FilterPaused0(opts *bind.FilterOpts, key [][32]byte) (*OmniPortalPaused0Iterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Paused0", keyRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalPaused0Iterator{contract: _OmniPortal.contract, event: "Paused0", logs: logs, sub: sub}, nil
}

// WatchPaused0 is a free log subscription operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_OmniPortal *OmniPortalFilterer) WatchPaused0(opts *bind.WatchOpts, sink chan<- *OmniPortalPaused0, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Paused0", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalPaused0)
				if err := _OmniPortal.contract.UnpackLog(event, "Paused0", log); err != nil {
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

// ParsePaused0 is a log parse operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_OmniPortal *OmniPortalFilterer) ParsePaused0(log types.Log) (*OmniPortalPaused0, error) {
	event := new(OmniPortalPaused0)
	if err := _OmniPortal.contract.UnpackLog(event, "Paused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OmniPortal contract.
type OmniPortalUnpausedIterator struct {
	Event *OmniPortalUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalUnpaused)
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
		it.Event = new(OmniPortalUnpaused)
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
func (it *OmniPortalUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalUnpaused represents a Unpaused event raised by the OmniPortal contract.
type OmniPortalUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_OmniPortal *OmniPortalFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OmniPortalUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalUnpausedIterator{contract: _OmniPortal.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_OmniPortal *OmniPortalFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_OmniPortal *OmniPortalFilterer) ParseUnpaused(log types.Log) (*OmniPortalUnpaused, error) {
	event := new(OmniPortalUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalUnpaused0Iterator is returned from FilterUnpaused0 and is used to iterate over the raw logs and unpacked data for Unpaused0 events raised by the OmniPortal contract.
type OmniPortalUnpaused0Iterator struct {
	Event *OmniPortalUnpaused0 // Event containing the contract specifics and raw log

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
func (it *OmniPortalUnpaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalUnpaused0)
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
		it.Event = new(OmniPortalUnpaused0)
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
func (it *OmniPortalUnpaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalUnpaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalUnpaused0 represents a Unpaused0 event raised by the OmniPortal contract.
type OmniPortalUnpaused0 struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused0 is a free log retrieval operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_OmniPortal *OmniPortalFilterer) FilterUnpaused0(opts *bind.FilterOpts, key [][32]byte) (*OmniPortalUnpaused0Iterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Unpaused0", keyRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalUnpaused0Iterator{contract: _OmniPortal.contract, event: "Unpaused0", logs: logs, sub: sub}, nil
}

// WatchUnpaused0 is a free log subscription operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_OmniPortal *OmniPortalFilterer) WatchUnpaused0(opts *bind.WatchOpts, sink chan<- *OmniPortalUnpaused0, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Unpaused0", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalUnpaused0)
				if err := _OmniPortal.contract.UnpackLog(event, "Unpaused0", log); err != nil {
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

// ParseUnpaused0 is a log parse operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_OmniPortal *OmniPortalFilterer) ParseUnpaused0(log types.Log) (*OmniPortalUnpaused0, error) {
	event := new(OmniPortalUnpaused0)
	if err := _OmniPortal.contract.UnpackLog(event, "Unpaused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalValidatorSetAddedIterator is returned from FilterValidatorSetAdded and is used to iterate over the raw logs and unpacked data for ValidatorSetAdded events raised by the OmniPortal contract.
type OmniPortalValidatorSetAddedIterator struct {
	Event *OmniPortalValidatorSetAdded // Event containing the contract specifics and raw log

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
func (it *OmniPortalValidatorSetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalValidatorSetAdded)
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
		it.Event = new(OmniPortalValidatorSetAdded)
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
func (it *OmniPortalValidatorSetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalValidatorSetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalValidatorSetAdded represents a ValidatorSetAdded event raised by the OmniPortal contract.
type OmniPortalValidatorSetAdded struct {
	SetId uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetAdded is a free log retrieval operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_OmniPortal *OmniPortalFilterer) FilterValidatorSetAdded(opts *bind.FilterOpts, setId []uint64) (*OmniPortalValidatorSetAddedIterator, error) {

	var setIdRule []interface{}
	for _, setIdItem := range setId {
		setIdRule = append(setIdRule, setIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "ValidatorSetAdded", setIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalValidatorSetAddedIterator{contract: _OmniPortal.contract, event: "ValidatorSetAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorSetAdded is a free log subscription operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_OmniPortal *OmniPortalFilterer) WatchValidatorSetAdded(opts *bind.WatchOpts, sink chan<- *OmniPortalValidatorSetAdded, setId []uint64) (event.Subscription, error) {

	var setIdRule []interface{}
	for _, setIdItem := range setId {
		setIdRule = append(setIdRule, setIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "ValidatorSetAdded", setIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalValidatorSetAdded)
				if err := _OmniPortal.contract.UnpackLog(event, "ValidatorSetAdded", log); err != nil {
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

// ParseValidatorSetAdded is a log parse operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_OmniPortal *OmniPortalFilterer) ParseValidatorSetAdded(log types.Log) (*OmniPortalValidatorSetAdded, error) {
	event := new(OmniPortalValidatorSetAdded)
	if err := _OmniPortal.contract.UnpackLog(event, "ValidatorSetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallPausedIterator is returned from FilterXCallPaused and is used to iterate over the raw logs and unpacked data for XCallPaused events raised by the OmniPortal contract.
type OmniPortalXCallPausedIterator struct {
	Event *OmniPortalXCallPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallPaused)
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
		it.Event = new(OmniPortalXCallPaused)
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
func (it *OmniPortalXCallPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallPaused represents a XCallPaused event raised by the OmniPortal contract.
type OmniPortalXCallPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXCallPaused is a free log retrieval operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_OmniPortal *OmniPortalFilterer) FilterXCallPaused(opts *bind.FilterOpts) (*OmniPortalXCallPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallPaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallPausedIterator{contract: _OmniPortal.contract, event: "XCallPaused", logs: logs, sub: sub}, nil
}

// WatchXCallPaused is a free log subscription operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_OmniPortal *OmniPortalFilterer) WatchXCallPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallPaused", log); err != nil {
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

// ParseXCallPaused is a log parse operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_OmniPortal *OmniPortalFilterer) ParseXCallPaused(log types.Log) (*OmniPortalXCallPaused, error) {
	event := new(OmniPortalXCallPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallToPausedIterator is returned from FilterXCallToPaused and is used to iterate over the raw logs and unpacked data for XCallToPaused events raised by the OmniPortal contract.
type OmniPortalXCallToPausedIterator struct {
	Event *OmniPortalXCallToPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallToPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallToPaused)
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
		it.Event = new(OmniPortalXCallToPaused)
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
func (it *OmniPortalXCallToPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallToPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallToPaused represents a XCallToPaused event raised by the OmniPortal contract.
type OmniPortalXCallToPaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXCallToPaused is a free log retrieval operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXCallToPaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXCallToPausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallToPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallToPausedIterator{contract: _OmniPortal.contract, event: "XCallToPaused", logs: logs, sub: sub}, nil
}

// WatchXCallToPaused is a free log subscription operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXCallToPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallToPaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallToPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallToPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallToPaused", log); err != nil {
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

// ParseXCallToPaused is a log parse operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) ParseXCallToPaused(log types.Log) (*OmniPortalXCallToPaused, error) {
	event := new(OmniPortalXCallToPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallToPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallToUnpausedIterator is returned from FilterXCallToUnpaused and is used to iterate over the raw logs and unpacked data for XCallToUnpaused events raised by the OmniPortal contract.
type OmniPortalXCallToUnpausedIterator struct {
	Event *OmniPortalXCallToUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallToUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallToUnpaused)
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
		it.Event = new(OmniPortalXCallToUnpaused)
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
func (it *OmniPortalXCallToUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallToUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallToUnpaused represents a XCallToUnpaused event raised by the OmniPortal contract.
type OmniPortalXCallToUnpaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXCallToUnpaused is a free log retrieval operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXCallToUnpaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXCallToUnpausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallToUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallToUnpausedIterator{contract: _OmniPortal.contract, event: "XCallToUnpaused", logs: logs, sub: sub}, nil
}

// WatchXCallToUnpaused is a free log subscription operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXCallToUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallToUnpaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallToUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallToUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallToUnpaused", log); err != nil {
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

// ParseXCallToUnpaused is a log parse operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) ParseXCallToUnpaused(log types.Log) (*OmniPortalXCallToUnpaused, error) {
	event := new(OmniPortalXCallToUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallToUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallUnpausedIterator is returned from FilterXCallUnpaused and is used to iterate over the raw logs and unpacked data for XCallUnpaused events raised by the OmniPortal contract.
type OmniPortalXCallUnpausedIterator struct {
	Event *OmniPortalXCallUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallUnpaused)
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
		it.Event = new(OmniPortalXCallUnpaused)
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
func (it *OmniPortalXCallUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallUnpaused represents a XCallUnpaused event raised by the OmniPortal contract.
type OmniPortalXCallUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXCallUnpaused is a free log retrieval operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_OmniPortal *OmniPortalFilterer) FilterXCallUnpaused(opts *bind.FilterOpts) (*OmniPortalXCallUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallUnpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallUnpausedIterator{contract: _OmniPortal.contract, event: "XCallUnpaused", logs: logs, sub: sub}, nil
}

// WatchXCallUnpaused is a free log subscription operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_OmniPortal *OmniPortalFilterer) WatchXCallUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallUnpaused", log); err != nil {
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

// ParseXCallUnpaused is a log parse operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_OmniPortal *OmniPortalFilterer) ParseXCallUnpaused(log types.Log) (*OmniPortalXCallUnpaused, error) {
	event := new(OmniPortalXCallUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgIterator is returned from FilterXMsg and is used to iterate over the raw logs and unpacked data for XMsg events raised by the OmniPortal contract.
type OmniPortalXMsgIterator struct {
	Event *OmniPortalXMsg // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsg)
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
		it.Event = new(OmniPortalXMsg)
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
func (it *OmniPortalXMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsg represents a XMsg event raised by the OmniPortal contract.
type OmniPortalXMsg struct {
	DestChainId uint64
	ShardId     uint64
	Offset      uint64
	Sender      common.Address
	To          common.Address
	Data        []byte
	GasLimit    uint64
	Fees        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterXMsg is a free log retrieval operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_OmniPortal *OmniPortalFilterer) FilterXMsg(opts *bind.FilterOpts, destChainId []uint64, shardId []uint64, offset []uint64) (*OmniPortalXMsgIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsg", destChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgIterator{contract: _OmniPortal.contract, event: "XMsg", logs: logs, sub: sub}, nil
}

// WatchXMsg is a free log subscription operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_OmniPortal *OmniPortalFilterer) WatchXMsg(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsg, destChainId []uint64, shardId []uint64, offset []uint64) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsg", destChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsg)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
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

// ParseXMsg is a log parse operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_OmniPortal *OmniPortalFilterer) ParseXMsg(log types.Log) (*OmniPortalXMsg, error) {
	event := new(OmniPortalXMsg)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMaxDataSizeSetIterator is returned from FilterXMsgMaxDataSizeSet and is used to iterate over the raw logs and unpacked data for XMsgMaxDataSizeSet events raised by the OmniPortal contract.
type OmniPortalXMsgMaxDataSizeSetIterator struct {
	Event *OmniPortalXMsgMaxDataSizeSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMaxDataSizeSet)
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
		it.Event = new(OmniPortalXMsgMaxDataSizeSet)
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
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMaxDataSizeSet represents a XMsgMaxDataSizeSet event raised by the OmniPortal contract.
type OmniPortalXMsgMaxDataSizeSet struct {
	Size uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXMsgMaxDataSizeSet is a free log retrieval operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMaxDataSizeSet(opts *bind.FilterOpts) (*OmniPortalXMsgMaxDataSizeSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMaxDataSizeSetIterator{contract: _OmniPortal.contract, event: "XMsgMaxDataSizeSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMaxDataSizeSet is a free log subscription operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMaxDataSizeSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMaxDataSizeSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMaxDataSizeSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxDataSizeSet", log); err != nil {
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

// ParseXMsgMaxDataSizeSet is a log parse operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMaxDataSizeSet(log types.Log) (*OmniPortalXMsgMaxDataSizeSet, error) {
	event := new(OmniPortalXMsgMaxDataSizeSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxDataSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMaxGasLimitSetIterator is returned from FilterXMsgMaxGasLimitSet and is used to iterate over the raw logs and unpacked data for XMsgMaxGasLimitSet events raised by the OmniPortal contract.
type OmniPortalXMsgMaxGasLimitSetIterator struct {
	Event *OmniPortalXMsgMaxGasLimitSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMaxGasLimitSet)
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
		it.Event = new(OmniPortalXMsgMaxGasLimitSet)
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
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMaxGasLimitSet represents a XMsgMaxGasLimitSet event raised by the OmniPortal contract.
type OmniPortalXMsgMaxGasLimitSet struct {
	GasLimit uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterXMsgMaxGasLimitSet is a free log retrieval operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMaxGasLimitSet(opts *bind.FilterOpts) (*OmniPortalXMsgMaxGasLimitSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMaxGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMaxGasLimitSetIterator{contract: _OmniPortal.contract, event: "XMsgMaxGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMaxGasLimitSet is a free log subscription operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMaxGasLimitSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMaxGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMaxGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMaxGasLimitSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxGasLimitSet", log); err != nil {
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

// ParseXMsgMaxGasLimitSet is a log parse operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMaxGasLimitSet(log types.Log) (*OmniPortalXMsgMaxGasLimitSet, error) {
	event := new(OmniPortalXMsgMaxGasLimitSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMinGasLimitSetIterator is returned from FilterXMsgMinGasLimitSet and is used to iterate over the raw logs and unpacked data for XMsgMinGasLimitSet events raised by the OmniPortal contract.
type OmniPortalXMsgMinGasLimitSetIterator struct {
	Event *OmniPortalXMsgMinGasLimitSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMinGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMinGasLimitSet)
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
		it.Event = new(OmniPortalXMsgMinGasLimitSet)
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
func (it *OmniPortalXMsgMinGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMinGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMinGasLimitSet represents a XMsgMinGasLimitSet event raised by the OmniPortal contract.
type OmniPortalXMsgMinGasLimitSet struct {
	GasLimit uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterXMsgMinGasLimitSet is a free log retrieval operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMinGasLimitSet(opts *bind.FilterOpts) (*OmniPortalXMsgMinGasLimitSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMinGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMinGasLimitSetIterator{contract: _OmniPortal.contract, event: "XMsgMinGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMinGasLimitSet is a free log subscription operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMinGasLimitSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMinGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMinGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMinGasLimitSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMinGasLimitSet", log); err != nil {
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

// ParseXMsgMinGasLimitSet is a log parse operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMinGasLimitSet(log types.Log) (*OmniPortalXMsgMinGasLimitSet, error) {
	event := new(OmniPortalXMsgMinGasLimitSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMinGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXReceiptIterator is returned from FilterXReceipt and is used to iterate over the raw logs and unpacked data for XReceipt events raised by the OmniPortal contract.
type OmniPortalXReceiptIterator struct {
	Event *OmniPortalXReceipt // Event containing the contract specifics and raw log

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
func (it *OmniPortalXReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXReceipt)
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
		it.Event = new(OmniPortalXReceipt)
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
func (it *OmniPortalXReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXReceipt represents a XReceipt event raised by the OmniPortal contract.
type OmniPortalXReceipt struct {
	SourceChainId uint64
	ShardId       uint64
	Offset        uint64
	GasUsed       *big.Int
	Relayer       common.Address
	Success       bool
	Err           []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXReceipt is a free log retrieval operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes err)
func (_OmniPortal *OmniPortalFilterer) FilterXReceipt(opts *bind.FilterOpts, sourceChainId []uint64, shardId []uint64, offset []uint64) (*OmniPortalXReceiptIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XReceipt", sourceChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXReceiptIterator{contract: _OmniPortal.contract, event: "XReceipt", logs: logs, sub: sub}, nil
}

// WatchXReceipt is a free log subscription operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes err)
func (_OmniPortal *OmniPortalFilterer) WatchXReceipt(opts *bind.WatchOpts, sink chan<- *OmniPortalXReceipt, sourceChainId []uint64, shardId []uint64, offset []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XReceipt", sourceChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXReceipt)
				if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
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

// ParseXReceipt is a log parse operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes err)
func (_OmniPortal *OmniPortalFilterer) ParseXReceipt(log types.Log) (*OmniPortalXReceipt, error) {
	event := new(OmniPortalXReceipt)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXReceiptMaxErrorSizeSetIterator is returned from FilterXReceiptMaxErrorSizeSet and is used to iterate over the raw logs and unpacked data for XReceiptMaxErrorSizeSet events raised by the OmniPortal contract.
type OmniPortalXReceiptMaxErrorSizeSetIterator struct {
	Event *OmniPortalXReceiptMaxErrorSizeSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXReceiptMaxErrorSizeSet)
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
		it.Event = new(OmniPortalXReceiptMaxErrorSizeSet)
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
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXReceiptMaxErrorSizeSet represents a XReceiptMaxErrorSizeSet event raised by the OmniPortal contract.
type OmniPortalXReceiptMaxErrorSizeSet struct {
	Size uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXReceiptMaxErrorSizeSet is a free log retrieval operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) FilterXReceiptMaxErrorSizeSet(opts *bind.FilterOpts) (*OmniPortalXReceiptMaxErrorSizeSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XReceiptMaxErrorSizeSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXReceiptMaxErrorSizeSetIterator{contract: _OmniPortal.contract, event: "XReceiptMaxErrorSizeSet", logs: logs, sub: sub}, nil
}

// WatchXReceiptMaxErrorSizeSet is a free log subscription operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) WatchXReceiptMaxErrorSizeSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXReceiptMaxErrorSizeSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XReceiptMaxErrorSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXReceiptMaxErrorSizeSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XReceiptMaxErrorSizeSet", log); err != nil {
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

// ParseXReceiptMaxErrorSizeSet is a log parse operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) ParseXReceiptMaxErrorSizeSet(log types.Log) (*OmniPortalXReceiptMaxErrorSizeSet, error) {
	event := new(OmniPortalXReceiptMaxErrorSizeSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceiptMaxErrorSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubValsetCutoffSetIterator is returned from FilterXSubValsetCutoffSet and is used to iterate over the raw logs and unpacked data for XSubValsetCutoffSet events raised by the OmniPortal contract.
type OmniPortalXSubValsetCutoffSetIterator struct {
	Event *OmniPortalXSubValsetCutoffSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubValsetCutoffSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubValsetCutoffSet)
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
		it.Event = new(OmniPortalXSubValsetCutoffSet)
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
func (it *OmniPortalXSubValsetCutoffSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubValsetCutoffSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubValsetCutoffSet represents a XSubValsetCutoffSet event raised by the OmniPortal contract.
type OmniPortalXSubValsetCutoffSet struct {
	Cutoff uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterXSubValsetCutoffSet is a free log retrieval operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_OmniPortal *OmniPortalFilterer) FilterXSubValsetCutoffSet(opts *bind.FilterOpts) (*OmniPortalXSubValsetCutoffSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubValsetCutoffSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubValsetCutoffSetIterator{contract: _OmniPortal.contract, event: "XSubValsetCutoffSet", logs: logs, sub: sub}, nil
}

// WatchXSubValsetCutoffSet is a free log subscription operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_OmniPortal *OmniPortalFilterer) WatchXSubValsetCutoffSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubValsetCutoffSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubValsetCutoffSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubValsetCutoffSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubValsetCutoffSet", log); err != nil {
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

// ParseXSubValsetCutoffSet is a log parse operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_OmniPortal *OmniPortalFilterer) ParseXSubValsetCutoffSet(log types.Log) (*OmniPortalXSubValsetCutoffSet, error) {
	event := new(OmniPortalXSubValsetCutoffSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubValsetCutoffSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitFromPausedIterator is returned from FilterXSubmitFromPaused and is used to iterate over the raw logs and unpacked data for XSubmitFromPaused events raised by the OmniPortal contract.
type OmniPortalXSubmitFromPausedIterator struct {
	Event *OmniPortalXSubmitFromPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitFromPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitFromPaused)
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
		it.Event = new(OmniPortalXSubmitFromPaused)
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
func (it *OmniPortalXSubmitFromPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitFromPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitFromPaused represents a XSubmitFromPaused event raised by the OmniPortal contract.
type OmniPortalXSubmitFromPaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXSubmitFromPaused is a free log retrieval operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitFromPaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXSubmitFromPausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitFromPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitFromPausedIterator{contract: _OmniPortal.contract, event: "XSubmitFromPaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitFromPaused is a free log subscription operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitFromPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitFromPaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitFromPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitFromPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromPaused", log); err != nil {
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

// ParseXSubmitFromPaused is a log parse operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitFromPaused(log types.Log) (*OmniPortalXSubmitFromPaused, error) {
	event := new(OmniPortalXSubmitFromPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitFromUnpausedIterator is returned from FilterXSubmitFromUnpaused and is used to iterate over the raw logs and unpacked data for XSubmitFromUnpaused events raised by the OmniPortal contract.
type OmniPortalXSubmitFromUnpausedIterator struct {
	Event *OmniPortalXSubmitFromUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitFromUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitFromUnpaused)
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
		it.Event = new(OmniPortalXSubmitFromUnpaused)
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
func (it *OmniPortalXSubmitFromUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitFromUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitFromUnpaused represents a XSubmitFromUnpaused event raised by the OmniPortal contract.
type OmniPortalXSubmitFromUnpaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXSubmitFromUnpaused is a free log retrieval operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitFromUnpaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXSubmitFromUnpausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitFromUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitFromUnpausedIterator{contract: _OmniPortal.contract, event: "XSubmitFromUnpaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitFromUnpaused is a free log subscription operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitFromUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitFromUnpaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitFromUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitFromUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromUnpaused", log); err != nil {
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

// ParseXSubmitFromUnpaused is a log parse operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitFromUnpaused(log types.Log) (*OmniPortalXSubmitFromUnpaused, error) {
	event := new(OmniPortalXSubmitFromUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitPausedIterator is returned from FilterXSubmitPaused and is used to iterate over the raw logs and unpacked data for XSubmitPaused events raised by the OmniPortal contract.
type OmniPortalXSubmitPausedIterator struct {
	Event *OmniPortalXSubmitPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitPaused)
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
		it.Event = new(OmniPortalXSubmitPaused)
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
func (it *OmniPortalXSubmitPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitPaused represents a XSubmitPaused event raised by the OmniPortal contract.
type OmniPortalXSubmitPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXSubmitPaused is a free log retrieval operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitPaused(opts *bind.FilterOpts) (*OmniPortalXSubmitPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitPaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitPausedIterator{contract: _OmniPortal.contract, event: "XSubmitPaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitPaused is a free log subscription operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitPaused", log); err != nil {
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

// ParseXSubmitPaused is a log parse operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitPaused(log types.Log) (*OmniPortalXSubmitPaused, error) {
	event := new(OmniPortalXSubmitPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitUnpausedIterator is returned from FilterXSubmitUnpaused and is used to iterate over the raw logs and unpacked data for XSubmitUnpaused events raised by the OmniPortal contract.
type OmniPortalXSubmitUnpausedIterator struct {
	Event *OmniPortalXSubmitUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitUnpaused)
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
		it.Event = new(OmniPortalXSubmitUnpaused)
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
func (it *OmniPortalXSubmitUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitUnpaused represents a XSubmitUnpaused event raised by the OmniPortal contract.
type OmniPortalXSubmitUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXSubmitUnpaused is a free log retrieval operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitUnpaused(opts *bind.FilterOpts) (*OmniPortalXSubmitUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitUnpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitUnpausedIterator{contract: _OmniPortal.contract, event: "XSubmitUnpaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitUnpaused is a free log subscription operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitUnpaused", log); err != nil {
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

// ParseXSubmitUnpaused is a log parse operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitUnpaused(log types.Log) (*OmniPortalXSubmitUnpaused, error) {
	event := new(OmniPortalXSubmitUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
