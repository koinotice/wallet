pragma solidity ^0.5.0;

/**
 * @title SafeMath
 * @dev Math operations with safety checks that revert on error
 */
library SafeMath {
    int256 constant private INT256_MIN = -2**255;

    /**
    * @dev Multiplies two unsigned integers, reverts on overflow.
    */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b);

        return c;
    }

    /**
    * @dev Multiplies two signed integers, reverts on overflow.
    */
    function mul(int256 a, int256 b) internal pure returns (int256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }

        require(!(a == -1 && b == INT256_MIN)); // This is the only case of overflow not detected by the check below

        int256 c = a * b;
        require(c / a == b);

        return c;
    }

    /**
    * @dev Integer division of two unsigned integers truncating the quotient, reverts on division by zero.
    */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        // Solidity only automatically asserts when dividing by 0
        require(b > 0);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold

        return c;
    }

    /**
    * @dev Integer division of two signed integers truncating the quotient, reverts on division by zero.
    */
    function div(int256 a, int256 b) internal pure returns (int256) {
        require(b != 0); // Solidity only automatically asserts when dividing by 0
        require(!(b == -1 && a == INT256_MIN)); // This is the only case of overflow

        int256 c = a / b;

        return c;
    }

    /**
    * @dev Subtracts two unsigned integers, reverts on overflow (i.e. if subtrahend is greater than minuend).
    */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a);
        uint256 c = a - b;

        return c;
    }

    /**
    * @dev Subtracts two signed integers, reverts on overflow.
    */
    function sub(int256 a, int256 b) internal pure returns (int256) {
        int256 c = a - b;
        require((b >= 0 && c <= a) || (b < 0 && c > a));

        return c;
    }

    /**
    * @dev Adds two unsigned integers, reverts on overflow.
    */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a);

        return c;
    }

    /**
    * @dev Adds two signed integers, reverts on overflow.
    */
    function add(int256 a, int256 b) internal pure returns (int256) {
        int256 c = a + b;
        require((b >= 0 && c >= a) || (b < 0 && c < a));

        return c;
    }

    /**
    * @dev Divides two unsigned integers and returns the remainder (unsigned integer modulo),
    * reverts when dividing by zero.
    */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0);
        return a % b;
    }
}

contract ERC20 {
    function totalSupply()
        public
        view
        returns (uint256);

    function balanceOf(
        address who
        )
        public
        view
        returns (uint256);

    function allowance(
        address owner,
        address spender
        )
        public
        view
        returns (uint256);

    function transfer(
        address to,
        uint256 value
        )
        public
        returns (bool);

    function transferFrom(
        address from,
        address to,
        uint256 value
        )
        public
        returns (bool);

    function approve(
        address spender,
        uint256 value
        )
        public
        returns (bool);
}

contract HashDice {
    struct BetOrder{
        address payable owner;
        uint256 startBlock;
        uint256 orderId;
        uint256 totalValue;
        uint256 maxCompensate;
        uint256 gain;
        uint8 betType;
        bool closed;
        uint32[] betValue;
    }

    struct Room{
        address payable creator; //RoomOwner
        address erc20Addr; //ERC20 token address
        string symbol; //Token name
        string name; //Room name
        uint256 id; //Unique room ID
        uint256[4] nominator; //0:Odd, 1:Even, 2:char, 3:tail
        uint256[4] denominator; //0:Odd, 1:Even, 2:char, 3:tail
        bool active;
        uint256 decimical; //Token decimical, such as 10**18, all bet order will multiple this value
        uint256 currentOrderId; //Current bet order id
        uint256 lastClosedOrderId; //Last time closed order id
        uint256 currentMaxCompensate;
        uint256 lastLockedValue; //Locked token value after last closed order
        mapping(uint256 => BetOrder) orders;
    }

    event RoomOpened(
        address creator, //RoomOwner
        address erc20Addr, //ERC20 token address
        string symbol, //Token name
        string name, //Room name
        uint256 id, //Unique room ID
        uint256 decimical, //Token decimical, all bet order will multiple this value
        uint256 lastLockedValue //Locked token value after last round
    );

    event RoomClosed(
        uint256 roomId
    );

    event Withdrawed(
        uint256 roomId,
        uint256 value,
        uint256 remained
    );

    event Deposited(
        uint256 roomId,
        uint256 value,
        uint256 remained
    );

    event NewBetOrder(
        uint256 roomId,
        uint256 orderId,
        uint256 value
    );

    event CloseRoundTooLate(
        uint256 roomId,
        uint256 orderId,
        uint256 blockNum,
        uint256 startBlock
    );

    event PayBetOwner(
        uint256 roomId,
        uint256 orderId,
        uint256 value
    );

    address public gOwner;
    uint256 public gRoomId;
    uint256 public gProfitRatio;
    uint256 public gRoundWait;
    uint256 public gRoundPeriod;

    mapping(uint256 => Room) rooms;
    mapping(address => uint256) profits;

    modifier onlyOwner() { require(msg.sender == gOwner, "Only allow contract owner to operate!"); _; }

    constructor() public payable{
        gOwner = msg.sender;
		gRoomId = 0;
        gProfitRatio = 100; //Default is 1%, UNIT is 10000
        gRoundWait = 10; //Default wait 10 block for any bet order
        gRoundPeriod = 10; //Default bet for 10 blocks
 	}

    function _CheckTransferFrom(
        address erc20Addr,
        address from,
        address to,
        uint256 value
    ) internal view {
        if (erc20Addr == address(0)) {
            require(value <= msg.value, "Cannot above current ether holdings!");
        } else {
            //First check whether msg.sender has enough tokens
            ERC20 token = ERC20(erc20Addr);
            uint256 totalBalance = token.balanceOf(from);
            require(value <= totalBalance, "Cannot above current token holdings!");
            uint256 spendable = token.allowance(
                from,
                to
            );
            if (spendable != 0) {
                spendable = (totalBalance < spendable) ? totalBalance : spendable;
            }
            require(spendable >= value, "Must approve locked tokens before.");
        }
    }

    function _DoRealTransferFrom(
        address erc20Addr,
        address from,
        address to,
        uint256 value
    ) internal {
        if (erc20Addr != address(0)) {
            ERC20 token = ERC20(erc20Addr);
            //Final tranfer bet valule to contract
            token.transferFrom(from, to, value);
        }
    }

    function _CheckTransfer(
        address erc20Addr,
        address from,
        uint256 value
    ) internal view {
        if (erc20Addr == address(0)) {
            require(from.balance >= value, "Not enough ether :(!");
        } else {
            ERC20 token = ERC20(erc20Addr);
            uint256 balance = token.balanceOf(from);
            require(balance >= value, "Not enough token :(!");
        }
    }

    function _DoRealTransfer(
        address erc20Addr,
        address payable to,
        uint256 value
    ) internal {
        if (erc20Addr == address(0)) {
            to.transfer(value);
        } else {
            ERC20 token = ERC20(erc20Addr);
            //Final tranfer bet valule to contract
            token.transfer(to, value);
        }
    }

    function GetBetOrder(
        uint256 roomId,
        uint256 orderId
    ) public view returns(address, uint256, uint256, uint256, uint256, bool, uint32[] memory) {
        require(roomId > 0 && orderId > 0);
        Room storage room = rooms[roomId];
        BetOrder storage order = room.orders[orderId];
        return (order.owner, order.startBlock, order.totalValue, order.gain, order.betType, 
            order.closed, order.betValue);
    }

    function GetRoomInfo(
        uint256 roomId
    ) public view returns(address, address, string memory, string memory, 
            uint256[4] memory, uint256[4] memory, bool, uint256, uint256, uint256, uint256, uint256) {
        require(roomId > 0);
        Room storage room = rooms[roomId];
        return (room.creator, room.erc20Addr, room.symbol, room.name,
            room.nominator, room.denominator, room.active, room.currentOrderId, room.lastClosedOrderId,
            room.currentMaxCompensate, room.lastLockedValue, room.decimical);
    }

    function GetRooms() public view returns (uint256) {
        return gRoomId;
    }

    function GetProfitRatio() public view returns (uint256) {
        return gProfitRatio;
    }

    function SetProfitRatio(uint256 ratio) public onlyOwner {
        require(ratio < 10000, "Cannot exceed profit UNIT!");
        gProfitRatio = ratio;
    }

    function GetRoundWait() public view returns (uint256) {
        return gRoundWait;
    }

    function SetRoundWait(uint256 value) public onlyOwner {
        gRoundWait = value;
    }

    function GetRoundPeriod() public view returns (uint256) {
        return gRoundPeriod;
    }

    function SetRoundPeriod(uint256 value) public onlyOwner {
        gRoundPeriod = value;
    }

    function OpenRoom(
        address erc20Addr, 
        string memory symbol, 
        uint256 initLockedValue, 
        uint256 decimical,
        string memory roomName, 
        uint256[4] memory nominator,
        uint256[4] memory denominator
    ) public payable returns(uint256 roomId){
        require(initLockedValue > 0 && decimical > 0);
        //Must meet (0 < nonimator <= denominator)
        for (uint i=0; i<4; i++) {
            require(nominator[i] > 0 && nominator[i] <= denominator[i]);
        }

        //Do the real token transfer, after that the token is locked in the contract.
        _CheckTransferFrom(erc20Addr, msg.sender, address(this), initLockedValue);
        gRoomId = SafeMath.add(gRoomId,1);

        Room memory newRoom;
        newRoom.creator = msg.sender;
        newRoom.id = gRoomId;
        newRoom.symbol = symbol;
        newRoom.erc20Addr = erc20Addr;
        newRoom.name = roomName;
        newRoom.nominator = nominator;
        newRoom.denominator = denominator;
        newRoom.currentMaxCompensate = 0;
        newRoom.lastLockedValue = initLockedValue;
        newRoom.active = true;
        newRoom.decimical = decimical;
        newRoom.currentOrderId = 0;
        newRoom.lastClosedOrderId = 0;

        rooms[gRoomId] = newRoom; //Automatically copy it from memory to storage
        _DoRealTransferFrom(erc20Addr, msg.sender, address(this), initLockedValue);
        emit RoomOpened(msg.sender, erc20Addr, symbol, roomName, gRoomId, decimical, initLockedValue);
        return gRoomId;
    }

    function CloseRoom(
        uint256 roomId
    ) public payable returns (bool) {
        require(roomId > 0);
        Room storage room = rooms[roomId];
        //Only allow room owner to withdraw money
        require(room.creator == msg.sender, 
            "Only allow room owner to withdraw money!");
        //Try to close active bet order first
        CloseBetOrders(roomId);
        if(room.lastClosedOrderId != room.currentOrderId) {
            //There exsits some future bet orders to be closed
            return false;
        }
        //Finally return back the locked tokens to room creator
        if(room.lastLockedValue > 0){
            //Do the real token transfer, after that the token is locked in the contract.
            uint256 lastLockedValue = room.lastLockedValue;
            _CheckTransfer(room.erc20Addr, address(this), room.lastLockedValue);
            room.lastLockedValue = 0;
            _DoRealTransfer(room.erc20Addr, room.creator, lastLockedValue);
        }
        room.active = false;
        emit RoomClosed(roomId);
        return true;
    }

    function GetProfit(
        address erc20Addr
    ) public view returns (uint256) {
        return profits[erc20Addr];
    }

    function WithdrawProfit(
        address erc20Addr,
        uint256 value,
        address payable receiver
    ) public payable onlyOwner {
        require(value > 0 && receiver != address(0));
        require(value <= profits[erc20Addr], "Withdraw profit over money!");
        //Do the real token transfer
        _CheckTransfer(erc20Addr, address(this), value);
        profits[erc20Addr] = SafeMath.sub(profits[erc20Addr], value);
        _DoRealTransfer(erc20Addr, receiver, value);
    }

    function Withdraw(
        uint256 roomId,
        uint256 value,
        address payable receiver
    ) public payable {
        require(roomId > 0);
        Room storage room = rooms[roomId];
        //Only allow room owner to withdraw money
        require(room.creator == msg.sender, 
            "Only allow room owner to withdraw money!");
        require(value <= room.lastLockedValue - room.currentMaxCompensate, "Withdraw over money!");
        //Finally return back the locked tokens to room creator
        if(room.lastLockedValue > 0){
            //Do the real token transfer
            _CheckTransfer(room.erc20Addr, address(this), value);
            room.lastLockedValue = SafeMath.sub(room.lastLockedValue, value);
            _DoRealTransfer(room.erc20Addr, receiver, value);
        }
        emit Withdrawed(roomId, value, room.lastLockedValue);
    }

    function Deposit(
        uint256 roomId,
        uint256 value
    ) public payable {
        require(roomId > 0 && value > 0);
        Room storage room = rooms[roomId];
        //Do the real token transfer, after that the token is locked in the contract.
        _CheckTransferFrom(room.erc20Addr, msg.sender, address(this), value);
        room.lastLockedValue = SafeMath.add(room.lastLockedValue, value);
        _DoRealTransferFrom(room.erc20Addr, msg.sender, address(this), value);

        emit Deposited(roomId, value, room.lastLockedValue);
    }

    /*
     * We allow anyone to call CloseBetOrders to trigger the hash result
     * --> Bet order owner should call it asap if he wins and detect room or game owner corruption.
     * --> Room owner should call it asap if he wins and detect game owner corruption.
     * --> Normally Game owner should call it after some blocks of open round.
     */
    function CloseBetOrders(
        uint256 roomId
    ) public payable {
        require(roomId > 0, "wrong roomId!");
        Room storage room = rooms[roomId];
        require(room.active, "The room is not actived!");
        uint256 orderId = room.lastClosedOrderId + 1;
        for (; orderId<=room.currentOrderId; orderId++) {
            BetOrder storage order = room.orders[orderId];
            //Close all bet orders whose start block is 10 less than current block
            if (block.number > SafeMath.add(order.startBlock, gRoundPeriod)){
                _CloseBetOrder(roomId, orderId);
            } else {
                break;
            }
        }
    }

    function CloseBetOrder(
        uint256 roomId,
        uint256 orderId
    ) public payable {
        require(roomId > 0 && orderId > 0);
        Room storage room = rooms[roomId];
        require(room.active, "The room is not actived!");
        BetOrder storage order = room.orders[orderId];
        require(!order.closed, "The bet order is already closed!");
        _CloseBetOrder(roomId, orderId);
    }

    function _CloseBetOrder(
        uint256 roomId,
        uint256 orderId
    ) internal {
        Room storage room = rooms[roomId];
        BetOrder storage order = room.orders[orderId];
        if(order.closed) {
            return;
        }

        if(block.number > SafeMath.add(order.startBlock, 256 - gRoundPeriod)){
            //It's too late to close one bet order, we cannot get the block hash
            //the only choice is to return back bet money.
            //Directly return back bet values
            _CheckTransfer(room.erc20Addr, address(this), order.totalValue);
            order.closed = true;
            order.gain = 0;
            room.currentMaxCompensate = SafeMath.sub(room.currentMaxCompensate, order.maxCompensate);
            room.lastClosedOrderId = order.orderId;
            _DoRealTransfer(room.erc20Addr, order.owner, order.totalValue);
            emit CloseRoundTooLate(roomId, orderId, block.number, order.startBlock);
        } else {
            uint8 betType = order.betType;
            if (betType == 0) {
                _CloseBetOrderType0(roomId, orderId);
            } else {
                _CloseBetOrderNormalType(roomId, orderId);
            }
        }
    }

    function _CloseBetOrderType0(
        uint256 roomId,
        uint256 orderId
    ) internal {
        Room storage room = rooms[roomId];
        BetOrder storage order = room.orders[orderId];

        //Now we should compensate each one according to block hashes.
        uint256 round = gRoundPeriod;
        uint256 gain = 0;
        for (uint8 i=0; i<round; i++){
            uint256 blockHash = uint256(blockhash(order.startBlock + i));
            uint8 blockHashTail = uint8(blockHash & 0xF);
            bool isChar = (blockHashTail >= 10 && blockHashTail <=15);
            bool isOdd = (!isChar && ((blockHashTail & 0x1) == 0x1));
            bool isEven = (!isChar && !isOdd);

            if((order.betValue[0] > 0) && isOdd){
                //bet order owner wins odd!
                gain = SafeMath.add(gain, 
                    order.betValue[0] * room.denominator[0] / room.nominator[0]);
            }
            if((order.betValue[1] > 0) && isEven){
                //bet order owner wins even!
                gain = SafeMath.add(gain, 
                    order.betValue[1] * room.denominator[1] / room.nominator[1]);
            }
            if((order.betValue[2] > 0) && isEven){
                //bet order owner wins char!
                gain = SafeMath.add(gain, 
                    order.betValue[2] * room.denominator[2] / room.nominator[2]);
            }
            if (order.betValue[3+blockHashTail] > 0) {
                //bet order owner wins tail!
                gain = SafeMath.add(gain, 
                    order.betValue[3+blockHashTail] * room.denominator[3] / room.nominator[3]);
            }
        }

        //require(SafeMath.add(room.lastLockedValue, round.totalBetValue) >= round.compensate, 
        //    "No enough money for compesate!");
        //In case of no money for compensation, though it should not happen because we alreay limit
        //  submmit order if exceed maximum compensation.
        gain = SafeMath.mul(gain, room.decimical);
        if (gain > 0) {
            _CheckTransfer(room.erc20Addr, address(this), gain);
        }
        order.closed = true;
        order.gain = gain;
        room.currentMaxCompensate = SafeMath.sub(room.currentMaxCompensate, order.maxCompensate);
        room.lastClosedOrderId = order.orderId;
        //Then pay the profit ratio
        uint256 profit = order.totalValue * gProfitRatio / 10000;
        profits[room.erc20Addr] = SafeMath.add(profits[room.erc20Addr], profit);
        //Finally pay the room owner
        room.lastLockedValue = SafeMath.add(room.lastLockedValue, order.totalValue - profit);
        room.lastLockedValue = SafeMath.sub(room.lastLockedValue, gain);
        if (gain > 0) {
            //Pay the betor
            _DoRealTransfer(room.erc20Addr, order.owner, gain);
            emit PayBetOwner(roomId, orderId, gain);
        }

    }

    function _CloseBetOrderNormalType(
        uint256 roomId,
        uint256 orderId
    ) internal {
        Room storage room = rooms[roomId];
        BetOrder storage order = room.orders[orderId];

        //Now we should compensate each one according to block hashes.
        uint256 round = gRoundPeriod;
        uint256 gain = 0;
        for (uint8 i=0; i<round; i++){
            uint256 blockHash = uint256(blockhash(order.startBlock + i));
            uint8 blockHashTail = uint8(blockHash & 0xF);
            bool isChar = (blockHashTail >= 10 && blockHashTail <=15);
            bool isOdd = (!isChar && ((blockHashTail & 0x1) == 0x1));
            bool isEven = (!isChar && !isOdd);

            uint8 count = 0;
            if (((order.betType & 0x1) == 0x1)){
                if((order.betValue[count*round + i] > 0) && isOdd){
                    //bet order owner wins odd!
                    gain = SafeMath.add(gain, 
                        order.betValue[count*round + i] * room.denominator[0] / room.nominator[0]);
                }
                count++;
            }
            if (((order.betType & 0x2) == 0x2)){
                if((order.betValue[count*round + i] > 0) && isEven){
                    //bet order owner wins even!
                    gain = SafeMath.add(gain, 
                        order.betValue[count*round + i] * room.denominator[1] / room.nominator[1]);
                }
                count++;
            }
            if (((order.betType & 0x4) == 0x4)){
                if((order.betValue[count*round + i] > 0) && isChar){
                    //bet order owner wins char!
                    gain = SafeMath.add(gain, 
                        order.betValue[count*round + i] * room.denominator[2] / room.nominator[2]);
                }
                count++;
            }
        }
        uint256 n = 0;
        uint8 betType8 = uint8(order.betType & 0xF);
        for (uint8 i=0; i<3; i++) {
            uint8 subType = betType8 & (1 * 2**i);
            if (subType != 0) {
                n++;
            }
        }
        n = n * round;
        if (((order.betType & 0x8) == 0x8)){
            uint32 total_tail = order.betValue[n++];
            for (uint32 i=0; i<total_tail; i++) {
                for (uint8 block_num=0; block_num<round; block_num++) {
                    uint256 blockHash = uint256(blockhash(order.startBlock + block_num));
                    uint8 blockHashTail = uint8(blockHash & 0xF);
                    if (order.betValue[n] == blockHashTail) {
                        //bet order owner wins char!
                        gain = SafeMath.add(gain, 
                            order.betValue[n+1] * room.denominator[3] / room.nominator[3]);
                    }
                }
                n += 2;
            }
        }

        //require(SafeMath.add(room.lastLockedValue, round.totalBetValue) >= round.compensate, 
        //    "No enough money for compesate!");
        //In case of no money for compensation, though it should not happen because we alreay limit
        //  submmit order if exceed maximum compensation.
        gain = SafeMath.mul(gain, room.decimical);
        if (gain > 0) {
            _CheckTransfer(room.erc20Addr, address(this), gain);
        }
        order.closed = true;
        order.gain = gain;
        room.currentMaxCompensate = SafeMath.sub(room.currentMaxCompensate, order.maxCompensate);
        room.lastClosedOrderId = order.orderId;
        //Then pay the profit ratio
        uint256 profit = order.totalValue * gProfitRatio / 10000;
        profits[room.erc20Addr] = SafeMath.add(profits[room.erc20Addr], profit);
        //Finally pay the room owner
        room.lastLockedValue = SafeMath.add(room.lastLockedValue, order.totalValue - profit);
        room.lastLockedValue = SafeMath.sub(room.lastLockedValue, gain);
        if (gain > 0) {
            //Pay the betor
            _DoRealTransfer(room.erc20Addr, order.owner, gain);
            emit PayBetOwner(roomId, orderId, gain);
        }
    }

    /*
     * @param betType, betType format is 4 bit |tail|char|even|odd|
     * @param betValue, betValue is 10 values per each bit in betType
     *                  plus [n, tail1, value1, tail2, value2, ... tailn, valuen]
     */
    function SubmitBetOrder(
        uint256 roomId, 
        uint256 startBlock,
        uint8 betType, 
        uint32[] memory betValue
    ) public payable returns (uint256){
        require(roomId > 0);
        if(startBlock == 0){
            startBlock = block.number + gRoundWait;
        } else {
            require(block.number + gRoundWait <= startBlock,
                "startBlock is too small!");
        }
        Room storage room = rooms[roomId];
        betType = betType & 0xF;
        require(betType != 0, "Invalid type!");
        require(room.active, "The room is not actived!");
 
        //Pre-calculate the maximum compensate and check whether 
        //room owner can pay the maximum compensate
        uint256 betBalance = 0;
        uint256 n = 0;
        uint256 maxCompensate = 0;
        uint256 nominator = 1;
        uint256 denominator = 1;
        uint256 round = gRoundPeriod;
        for (uint8 i=0; i<3; i++) {
            if ((betType & (1 * 2**i)) != 0) {
                require(betValue.length >= (n+1)*round, "Missing bet values!");
                nominator = room.nominator[i];
                denominator = room.denominator[i];
                for (uint8 j=0; j<round; j++){
                    maxCompensate = SafeMath.add(maxCompensate, betValue[n*round+j] * denominator / nominator);
                    betBalance = SafeMath.add(betBalance, betValue[n*round + j]);
                }
                n++;
            }
        }
        require(betValue.length >= n*round, "Missing bet values!");
        //Now check the bet tail info, format is
        //[n, tail1, value1, tail2, value2, ... tailn, valuen]
        n = n*round;
        if (((betType & 0x8) == 0x8) && betValue.length > n) {
            uint32 total_tail = betValue[n++];
            require(betValue.length >= n + total_tail*2, "betTailInfo format wrong!");
            nominator = room.nominator[3];
            denominator = room.denominator[3];
            for (uint32 i=0; i<total_tail; i++) {
                uint256 bet_value = betValue[n+1] * round;
                uint256 compensate = bet_value * denominator / nominator;
                maxCompensate = SafeMath.add(maxCompensate, compensate);
                betBalance = SafeMath.add(betBalance, bet_value);
                n += 2;
            }
        }

        uint256 decimical = room.decimical;
        betBalance = SafeMath.mul(betBalance, decimical);
        require(betBalance > 0, "Must bet something!");
        maxCompensate = SafeMath.mul(maxCompensate, decimical);
        require(maxCompensate <= room.lastLockedValue - room.currentMaxCompensate, "Cannot bet over compensate!");

        _CheckTransferFrom(room.erc20Addr, msg.sender, address(this), betBalance);
        room.currentMaxCompensate = SafeMath.add(room.currentMaxCompensate, maxCompensate);

        room.currentOrderId = SafeMath.add(room.currentOrderId, 1);
        BetOrder memory newOrder;
        newOrder.orderId = room.currentOrderId;
        newOrder.totalValue = betBalance;
        newOrder.betType = betType;
        newOrder.betValue = betValue;
        newOrder.owner = msg.sender;
        newOrder.closed = false;
        newOrder.gain = 0;
        newOrder.maxCompensate = maxCompensate;
        newOrder.startBlock = startBlock;
        room.orders[newOrder.orderId] = newOrder;
        _DoRealTransferFrom(room.erc20Addr, msg.sender, address(this), betBalance);
        emit NewBetOrder(roomId, newOrder.orderId, betBalance);
        return newOrder.orderId;
    }

    /*
     * @param betValue, betValue is [odd_val, even_val, char_val, tail1_val, tail2_val, ... tail16_val]
     *            Internally this type of bet order is marked as betType 0.
     */
    function SubmitBetOrderType0(
        uint256 roomId, 
        uint256 startBlock,
        uint32[] memory betValue
    ) public payable returns (uint256){
        require(roomId > 0);
        if(startBlock == 0){
            startBlock = block.number + gRoundWait;
        } else {
            require(block.number + gRoundWait <= startBlock,
                "startBlock is too small!");
        }
        Room storage room = rooms[roomId];
        require(room.active, "The room is not actived!");
        require(betValue.length == 19, "betValue length must be exactly 19!");
 
        //Pre-calculate the maximum compensate and check whether 
        //room owner can pay the maximum compensate
        uint256 betBalance = 0;
        uint256 maxCompensate = 0;
        uint256 nominator = 1;
        uint256 denominator = 1;
        uint256 round = gRoundPeriod;
        for (uint8 i=0; i<3; i++) {
            if (betValue[i] > 0) {
                nominator = room.nominator[i];
                denominator = room.denominator[i];
                maxCompensate = SafeMath.add(maxCompensate, betValue[i] * round * denominator / nominator);
                betBalance = SafeMath.add(betBalance, betValue[i] * round);
            }
        }
        nominator = room.nominator[3];
        denominator = room.denominator[3];
        for (uint8 i=3; i<19; i++) {
            if (betValue[i] > 0) {
                maxCompensate = SafeMath.add(maxCompensate, betValue[i] * round * denominator / nominator);
                betBalance = SafeMath.add(betBalance, betValue[i] * round);
            }
        }
        uint256 decimical = room.decimical;
        betBalance = SafeMath.mul(betBalance, decimical);
        require(betBalance > 0, "Must bet something!");
        maxCompensate = SafeMath.mul(maxCompensate, decimical);
        require(maxCompensate <= room.lastLockedValue - room.currentMaxCompensate, "Cannot bet over compensate!");

        _CheckTransferFrom(room.erc20Addr, msg.sender, address(this), betBalance);
        room.currentMaxCompensate = SafeMath.add(room.currentMaxCompensate, maxCompensate);

        room.currentOrderId = SafeMath.add(room.currentOrderId, 1);
        BetOrder memory newOrder;
        newOrder.orderId = room.currentOrderId;
        newOrder.totalValue = betBalance;
        newOrder.betType = 0;
        newOrder.betValue = betValue;
        newOrder.owner = msg.sender;
        newOrder.closed = false;
        newOrder.gain = 0;
        newOrder.maxCompensate = maxCompensate;
        newOrder.startBlock = startBlock;
        room.orders[newOrder.orderId] = newOrder;
        _DoRealTransferFrom(room.erc20Addr, msg.sender, address(this), betBalance);
        emit NewBetOrder(roomId, newOrder.orderId, betBalance);
        return newOrder.orderId;

    }
}