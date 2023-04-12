// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {ERC20} from "transmissions11/solmate/tokens/ERC20.sol";

contract TokenMock is ERC20("Test Token", "TST20", 18) {
    constructor(uint256 _initialSupply) {
        _mint(msg.sender, _initialSupply);
    }
}
