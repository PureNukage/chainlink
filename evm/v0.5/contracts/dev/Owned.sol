pragma solidity 0.5.0;

/**
 * @title The Owned contract
 * @notice A contract with helpers for basic contract ownership.
 */
contract Owned {

  address public owner;
  address private pendingOwner;

  event OwnershipTransferRequested(
    address indexed to,
    address from
  );
  event OwnershipTransfered(
    address indexed to,
    address from
  );

  constructor() public {
    owner = msg.sender;
  }

  function transferOwnership(address _to)
    public
    onlyOwner()
  {
    pendingOwner = _to;

    emit OwnershipTransferRequested(_to, owner);
  }

  function acceptOwnership()
    public
  {
    require(msg.sender == pendingOwner, "Must be requested to accept ownership");

    address oldOwner = owner;
    owner = msg.sender;
    pendingOwner = address(0);

    emit OwnershipTransfered(msg.sender, oldOwner);
  }

  modifier onlyOwner() {
    require(msg.sender == owner, "Only callable by owner");
    _;
  }

}
