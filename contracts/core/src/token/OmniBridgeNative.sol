// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OmniBridgeCommon } from "./OmniBridgeCommon.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { OmniBridgeL1 } from "./OmniBridgeL1.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";
import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title OmniBridgeNative
 * @notice The Omni side of Omni's native token bridge. Partner to OmniBridgeL1, which is deployed to Ethereum.
 *         This contract is predeployed to Omni's EVM, prefunded with native OMNI tokens to match totalL1Supply, such
 *         that each L1 token has a "sibling" native token on Omni.
 * @dev This contract is predeployed, and requires storage slots to be set in genesis.
 *      Genesis storage slots must:
 *          - set _owner on proxy
 *          - set _initialized on proxy to 1, to disable the initializer
 *          - set _initialized on implementation to 0xffffffffffffffff, to disabled all initializers
 */
contract OmniBridgeNative is OmniBridgeCommon {
    /**
     * @notice Emitted when an account deposits OMNI, bridging it to Ethereum.
     */
    event Bridge(address indexed payor, address indexed to, uint256 amount);

    /**
     * @notice Emitted when OMNI tokens are withdrawn for an account.
     *         If success is false, the amount is claimable by the account.
     */
    event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success);

    /**
     * @notice Emitted when an account claims OMNI tokens that failed to be withdrawn.
     */
    event Claimed(address indexed claimant, address indexed to, uint256 amount);

    /**
     * @notice Emitted on setup(...)
     */
    event Setup(uint64 l1ChainId, address omni, address l1Bridge);

    /**
     * @notice xcall gas limit for OmniBridgeL1.withdraw
     */
    uint64 public constant XCALL_WITHDRAW_GAS_LIMIT = 80_000;

    /**
     * @notice L1 chain id, configurable so that this contract can be used on testnets.
     */
    uint64 public l1ChainId;

    /**
     * @notice The OmniPortal contract.
     */
    IOmniPortal public omni;

    /**
     * @notice Total OMNI tokens deposited to OmniBridgeL1.
     *
     *         If l1BridgeBalance == totalL1Supply, all OMNI tokens are on Omni's EVM.
     *         If l1BridgeBalance == 0, withdraws to L1 are blocked.
     *
     *         Without validator rewards, totalL1Deposits == 0 would mean all OMNI tokens are on Ethereum.
     *         However with validator rewards, some OMNI may remain on Omni.
     *
     *         This balance is synced on each withdraw to Omni, and decremented on each bridge to back L1.
     */
    uint256 public l1BridgeBalance;

    /**
     * @notice The address of the OmniBridgeL1 contract deployed to Ethereum.
     */
    address public l1Bridge;

    /**
     * @notice Track claimable amount per address. Claimable amount increases when withdraw transfer fails.
     */
    mapping(address => uint256) public claimable;

    function initialize(address owner_) external initializer {
        __Ownable_init(owner_);
    }

    /**
     * @notice Withdraw `amount` native OMNI to `to`. Only callable via xcall from OmniBridgeL1.
     * @param payor     The address of the account with OMNI on L1.
     * @param to        The address to receive the OMNI on Omni.
     * @param amount    The amount of OMNI to withdraw.
     * @param l1Balance The OMNI balance of the L1 bridge contract, synced on each withdraw.
     */
    function withdraw(address payor, address to, uint256 amount, uint256 l1Balance)
        external
        whenNotPaused(ACTION_WITHDRAW)
    {
        XTypes.MsgContext memory xmsg = omni.xmsg();

        require(msg.sender == address(omni), "OmniBridge: not xcall"); // this protects against reentrancy
        require(xmsg.sender == l1Bridge, "OmniBridge: not bridge");
        require(xmsg.sourceChainId == l1ChainId, "OmniBridge: not L1");

        l1BridgeBalance = l1Balance;

        (bool success,) = to.call{ value: amount }("");

        if (!success) claimable[payor] += amount;

        emit Withdraw(payor, to, amount, success);
    }

    /**
     * @notice Bridge `amount` OMNI to `to` on L1.
     */
    function bridge(address to, uint256 amount) external payable whenNotPaused(ACTION_BRIDGE) {
        _bridge(to, amount);
    }

    /**
     * @dev Trigger a withdraw of `amount` OMNI to `to` on L1, via xcall.
     */
    function _bridge(address to, uint256 amount) internal {
        require(to != address(0), "OmniBridge: no bridge to zero");
        require(amount > 0, "OmniBridge: amount must be > 0");
        require(amount <= l1BridgeBalance, "OmniBridge: no liquidity");

        uint256 fee = bridgeFee(to, amount);
        require(msg.value == amount + fee, "OmniBridge: incorrect funds");

        l1BridgeBalance -= amount;

        omni.xcall{ value: fee }(
            l1ChainId,
            ConfLevel.Finalized,
            l1Bridge,
            abi.encodeCall(OmniBridgeL1.withdraw, (to, amount)),
            XCALL_WITHDRAW_GAS_LIMIT
        );

        emit Bridge(msg.sender, to, amount);
    }

    /**
     * @notice Return the xcall fee required to bridge `amount` to `to`.
     */
    function bridgeFee(address to, uint256 amount) public view returns (uint256) {
        return omni.feeFor(l1ChainId, abi.encodeCall(OmniBridgeL1.withdraw, (to, amount)), XCALL_WITHDRAW_GAS_LIMIT);
    }

    /**
     * @notice Claim OMNI tokens that failed to be withdrawn via xmsg from OmniBridgeL1.
     * @dev We require this be made by xcall, because an account on Omni may not be authorized to
     *      claim for that address on L1. Consider the case wherein the address of the L1 contract that initiated the
     *      withdraw is the same as the address of a contract on Omni deployed and owned by a malicious actor.
     */
    function claim(address to) external whenNotPaused(ACTION_WITHDRAW) {
        XTypes.MsgContext memory xmsg = omni.xmsg();

        require(msg.sender == address(omni), "OmniBridge: not xcall");
        require(xmsg.sourceChainId == l1ChainId, "OmniBridge: not L1");
        require(to != address(0), "OmniBridge: no claim to zero");

        address claimant = xmsg.sender;
        require(claimable[claimant] > 0, "OmniBridge: nothing to claim");

        uint256 amount = claimable[claimant];
        claimable[claimant] = 0;

        (bool success,) = to.call{ value: amount }("");
        require(success, "OmniBridge: transfer failed");

        emit Claimed(claimant, to, amount);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                          Admin functions                                 //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Setup core contract parameters, done by owner immediately after pre-deployment.
     * @param l1ChainId_    The chain id of the L1 network.
     * @param omni_         The address of the OmniPortal contract.
     * @param l1Bridge_     The address of the L1 OmniBridge contract.
     */
    function setup(uint64 l1ChainId_, address omni_, address l1Bridge_) external onlyOwner {
        l1ChainId = l1ChainId_;
        omni = IOmniPortal(omni_);
        l1Bridge = l1Bridge_;
        emit Setup(l1ChainId_, omni_, l1Bridge_);
    }
}
