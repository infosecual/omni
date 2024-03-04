// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_allowlist_Test
 * @dev Test suite for the AVS allowlist functionality
 */
contract OmniAVS_allowlist_Test is Base {
    /// @dev Test that an operator can be added to the allowlist
    function test_addToAllowlist_succeeds() public {
        address operator = _operator(0);
        _addToAllowlist(operator);
        assertTrue(omniAVS.isInAllowlist(operator));
    }

    /// @dev Test that an operator can be removed from the allowlist
    function test_removeFromAllowlist_succeeds() public {
        address operator1 = _operator(0);
        address operator2 = _operator(1);

        _addToAllowlist(operator1);
        _addToAllowlist(operator2);
        assertTrue(omniAVS.isInAllowlist(operator1));
        assertTrue(omniAVS.isInAllowlist(operator2));

        _removeFromAllowlist(operator1);
        assertFalse(omniAVS.isInAllowlist(operator1));
        assertTrue(omniAVS.isInAllowlist(operator2));
    }

    /// @dev Test that only the owner can add to the allowlist
    function test_addToAllowlist_notOwner_reverts() public {
        address operator = _operator(0);

        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.addToAllowlist(operator);
    }

    /// @dev Test that only the owner can remove from the allowlist
    function test_removeFromAllowlist_notOwner_reverts() public {
        address operator = _operator(0);

        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.removeFromAllowlist(operator);
    }

    /// @dev Test that an operator can register if in allow list
    function test_registerOperator_succeeds() public {
        address operator = _operator(0);

        _addToAllowlist(operator);
        _depositIntoSupportedStrategy(operator, 1 ether);
        _registerAsOperator(operator);
        _registerOperatorWithAVS(operator);

        IOmniAVS.Operator[] memory operators = omniAVS.operators();

        assertEq(operators.length, 1);
        assertEq(operators[0].addr, operator);
    }

    /// @dev Test that an operator cannot register if not in allow list
    function test_registerOperator_notAllowed_reverts() public {
        address operator = _operator(0);

        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;

        vm.expectRevert("OmniAVS: not allowed");
        vm.prank(operator);
        omniAVS.registerOperatorToAVS(operator, emptySig);
    }

    /// @dev Test that the owner can set the ethStakeInbox
    function test_setEthStakeInbox_succeeds() public {
        address newInbox = 0x1234567890123456789012345678901234567890;

        vm.prank(omniAVSOwner);
        omniAVS.setEthStakeInbox(newInbox);
        assertEq(omniAVS.ethStakeInbox(), newInbox);
    }

    /// @dev Test that only the owner can set the ethStakeInbox
    function test_setEthStakeInbox_notOwner_reverts() public {
        address newInbox = 0x1234567890123456789012345678901234567890;

        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setEthStakeInbox(newInbox);
    }

    /// @dev Test the the ethStakeInbox cannot be set to the zero address
    function test_setEthStakeInbox_zeroAddress_reverts() public {
        vm.prank(omniAVSOwner);
        vm.expectRevert("OmniAVS: zero address");
        omniAVS.setEthStakeInbox(address(0));
    }
}