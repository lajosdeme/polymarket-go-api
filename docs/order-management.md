# Orders Overview

Detailed instructions for creating, placing, and managing orders using Polymarket’s CLOB API.
All orders are expressed as limit orders (can be marketable). The underlying order primitive must be in the form expected and executable by the on-chain binary limit order protocol contract. Preparing such an order is quite involved (structuring, hashing, signing), thus Polymarket suggests using the open source typescript, python and golang libraries.
​
## Allowances
To place an order, allowances must be set by the funder address for the specified maker asset for the Exchange contract. When buying, this means the funder must have set a USDC allowance greater than or equal to the spending amount. When selling, the funder must have set an allowance for the conditional token that is greater than or equal to the selling amount. This allows the Exchange contract to execute settlement according to the signed order instructions created by a user and matched by the operator.

## Signature Types
Polymarket’s CLOB supports 3 signature types. Orders must identify what signature type they use. The available typescript and python clients abstract the complexity of signing and preparing orders with the following signature types by allowing a funder address and signer type to be specified on initialization. The supported signature types are:
Type	ID	Description
EOA	0	EIP712 signature signed by an EOA
POLY_PROXY	1	EIP712 signatures signed by a signer associated with funding Polymarket proxy wallet
POLY_GNOSIS_SAFE	2	EIP712 signatures signed by a signer associated with funding Polymarket gnosis safe wallet

Validity Checks
Orders are continually monitored to make sure they remain valid. Specifically, this includes continually tracking underlying balances, allowances and on-chain order cancellations. Any maker that is caught intentionally abusing these checks (which are essentially real time) will be blacklisted. Additionally, there are rails on order placement in a market. Specifically, you can only place orders that sum to less than or equal to your available balance for each market. For example if you have 500 USDC in your funding wallet, you can place one order to buy 1000 YES in marketA @ $.50, then any additional buy orders to that market will be rejected since your entire balance is reserved for the first (and only) buy order. More explicitly the max size you can place for an order is:

maxOrderSize=underlyingAssetBalance−∑(orderSize−orderFillAmount)

## Place Single Order

Detailed instructions for creating, placing, and managing orders using Polymarket’s CLOB API.
### Create and Place an Order
 **This endpoint requires a L2 Header **

Create and place an order using the Polymarket CLOB API clients. All orders are represented as “limit” orders, but “market” orders are also supported. To place a market order, simply ensure your price is marketable against current resting limit orders, which are executed on input at the best price.

```
POST /<clob-endpoint>/order
```

Request Payload Parameters
Name	Required	Type	Description
order	yes	Order	signed object
owner	yes	string	api key of order owner
orderType	yes	string	order type (“FOK”, “GTC”, “GTD”)

An order object is the form:
Name	Required	Type	Description
salt	yes	integer	random salt used to create unique order
maker	yes	string	maker address (funder)
signer	yes	string	signing address
taker	yes	string	taker address (operator)
tokenId	yes	string	ERC1155 token ID of conditional token being traded
makerAmount	yes	string	maximum amount maker is willing to spend
takerAmount	yes	string	minimum amount taker will pay the maker in return
expiration	yes	string	unix expiration timestamp
nonce	yes	string	maker’s exchange nonce of the order is associated
feeRateBps	yes	string	fee rate basis points as required by the operator
side	yes	string	buy or sell enum index
signatureType	yes	integer	signature type enum index
signature	yes	string	hex encoded signature

Order types

    FOK: A Fill-Or-Kill order is an market order to buy (in dollars) or sell (in shares) shares that must be executed immediately in its entirety; otherwise, the entire order will be cancelled.
    FAK: A Fill-And-Kill order is a market order to buy (in dollars) or sell (in shares) that will be executed immediately for as many shares as are available; any portion not filled at once is cancelled.
    GTC: A Good-Til-Cancelled order is a limit order that is active until it is fulfilled or cancelled.
    GTD: A Good-Til-Date order is a type of order that is active until its specified date (UTC seconds timestamp), unless it has already been fulfilled or cancelled. There is a security threshold of one minute. If the order needs to expire in 90 seconds the correct expiration value is: now + 1 minute + 30 seconds

Response Format
Name	Type	Description
success	boolean	boolean indicating if server-side err (success = false) -> server-side error
errorMsg	string	error message in case of unsuccessful placement (in case success = false, e.g. client-side error, the reason is in errorMsg)
orderId	string	id of order
orderHashes	string[]	hash of settlement transaction order was marketable and triggered a match

Insert Error Messages
If the errorMsg field of the response object from placement is not an empty string, the order was not able to be immediately placed. This might be because of a delay or because of a failure. If the success is not true, then there was an issue placing the order. The following errorMessages are possible:
​
Error
Error	Success	Message	Description
INVALID_ORDER_MIN_TICK_SIZE	yes	order is invalid. Price breaks minimum tick size rules	order price isn’t accurate to correct tick sizing
INVALID_ORDER_MIN_SIZE	yes	order is invalid. Size lower than the minimum	order size must meet min size threshold requirement
INVALID_ORDER_DUPLICATED	yes	order is invalid. Duplicated. Same order has already been placed, can’t be placed again	
INVALID_ORDER_NOT_ENOUGH_BALANCE	yes	not enough balance / allowance	funder address doesn’t have sufficient balance or allowance for order
INVALID_ORDER_EXPIRATION	yes	invalid expiration	expiration field expresses a time before now
INVALID_ORDER_ERROR	yes	could not insert order	system error while inserting order
EXECUTION_ERROR	yes	could not run the execution	system error while attempting to execute trade
ORDER_DELAYED	no	order match delayed due to market conditions	order placement delayed
DELAYING_ORDER_ERROR	yes	error delaying the order	system error while delaying order
FOK_ORDER_NOT_FILLED_ERROR	yes	order couldn’t be fully filled, FOK orders are fully filled/killed	FOK order not fully filled so can’t be placed
MARKET_NOT_READY	no	the market is not yet ready to process new orders	system not accepting orders for market yet

Insert Statuses
When placing an order, a status field is included. The status field provides additional information regarding the order’s state as a result of the placement. Possible values include:
​
Status
Status	Description
matched	order placed and matched with an existing resting order
live	order placed and resting on the book
delayed	order marketable, but subject to matching delay
unmatched	order marketable, but failure delaying, placement successful

## Place Multiple Orders (Batching)

Instructions for placing multiple orders(Batch)

 **This endpoint requires a L2 Header** 

Polymarket’s CLOB supports batch orders, allowing you to place up to 15 orders in a single request. Before using this feature, make sure you’re comfortable placing a single order first. You can find the documentation for that [here](https://docs.polymarket.com/developers/CLOB/orders/create-order).


```
POST /<clob-endpoint>/orders
```
​
Request Payload Parameters
Name	Required	Type	Description
PostOrder	yes	PostOrders[]	list of signed order objects (Signed Order + Order Type + Owner)

A PostOrder object is the form:
Name	Required	Type	Description
order	yes	order	See below table for details on crafting this object
orderType	yes	string	order type (“FOK”, “GTC”, “GTD”, “FAK”)
owner	yes	string	api key of order owner

An order object is the form:
Name	Required	Type	Description
salt	yes	integer	random salt used to create unique order
maker	yes	string	maker address (funder)
signer	yes	string	signing address
taker	yes	string	taker address (operator)
tokenId	yes	string	ERC1155 token ID of conditional token being traded
makerAmount	yes	string	maximum amount maker is willing to spend
takerAmount	yes	string	minimum amount taker will pay the maker in return
expiration	yes	string	unix expiration timestamp
nonce	yes	string	maker’s exchange nonce of the order is associated
feeRateBps	yes	string	fee rate basis points as required by the operator
side	yes	string	buy or sell enum index
signatureType	yes	integer	signature type enum index
signature	yes	string	hex encoded signature

Order types

    FOK: A Fill-Or-Kill order is an market order to buy (in dollars) or sell (in shares) shares that must be executed immediately in its entirety; otherwise, the entire order will be cancelled.
    FAK: A Fill-And-Kill order is a market order to buy (in dollars) or sell (in shares) that will be executed immediately for as many shares as are available; any portion not filled at once is cancelled.
    GTC: A Good-Til-Cancelled order is a limit order that is active until it is fulfilled or cancelled.
    GTD: A Good-Til-Date order is a type of order that is active until its specified date (UTC seconds timestamp), unless it has already been fulfilled or cancelled. There is a security threshold of one minute. If the order needs to expire in 90 seconds the correct expiration value is: now + 1 minute + 30 seconds

Response Format
Name	Type	Description
success	boolean	boolean indicating if server-side err (success = false) -> server-side error
errorMsg	string	error message in case of unsuccessful placement (in case success = false, e.g. client-side error, the reason is in errorMsg)
orderId	string	id of order
orderHashes	string[]	hash of settlement transaction order was marketable and triggered a match

Insert Error Messages
If the errorMsg field of the response object from placement is not an empty string, the order was not able to be immediately placed. This might be because of a delay or because of a failure. If the success is not true, then there was an issue placing the order. The following errorMessages are possible:
​
Error
Error	Success	Message	Description
INVALID_ORDER_MIN_TICK_SIZE	yes	order is invalid. Price breaks minimum tick size rules	order price isn’t accurate to correct tick sizing
INVALID_ORDER_MIN_SIZE	yes	order is invalid. Size lower than the minimum	order size must meet min size threshold requirement
INVALID_ORDER_DUPLICATED	yes	order is invalid. Duplicated. Same order has already been placed, can’t be placed again	
INVALID_ORDER_NOT_ENOUGH_BALANCE	yes	not enough balance / allowance	funder address doesn’t have sufficient balance or allowance for order
INVALID_ORDER_EXPIRATION	yes	invalid expiration	expiration field expresses a time before now
INVALID_ORDER_ERROR	yes	could not insert order	system error while inserting order
EXECUTION_ERROR	yes	could not run the execution	system error while attempting to execute trade
ORDER_DELAYED	no	order match delayed due to market conditions	order placement delayed
DELAYING_ORDER_ERROR	yes	error delaying the order	system error while delaying order
FOK_ORDER_NOT_FILLED_ERROR	yes	order couldn’t be fully filled, FOK orders are fully filled/killed	FOK order not fully filled so can’t be placed
MARKET_NOT_READY	no	the market is not yet ready to process new orders	system not accepting orders for market yet

Insert Statuses
When placing an order, a status field is included. The status field provides additional information regarding the order’s state as a result of the placement. Possible values include:
​
Status
Status	Description
matched	order placed and matched with an existing resting order
live	order placed and resting on the book
delayed	order marketable, but subject to matching delay
unmatched	order marketable, but failure delaying, placement successful

Example payload: 
```
[
    {'order': {'salt': 660377097, 'maker': '0x17A9568474b5fc84B1D1C44f081A0a3aDE750B2b', 'signer': '0x17A9568474b5fc84B1D1C44f081A0a3aDE750B2b', 'taker': '0x0000000000000000000000000000000000000000', 'tokenId': '88613172803544318200496156596909968959424174365708473463931555296257475886634', 'makerAmount': '50000', 'takerAmount': '5000000', 'expiration': '0', 'nonce': '0', 'feeRateBps': '0', 'side': 'BUY', 'signatureType': 0, 'signature': '0xccb8d1298d698ebc0859e6a26044c848ac4a4b0e20a391a4574e42b9c9bf237e5fa09fc00743e3e2d2f8e909a21d60f276ce083cc35c6661410b892f5bcbe2291c'}, 'owner': 'PRIVATEKEY', 'orderType': 'GTC'}, 
    {'order': {'salt': 1207111323, 'maker': '0x17A9568474b5fc84B1D1C44f081A0a3aDE750B2b', 'signer': '0x17A9568474b5fc84B1D1C44f081A0a3aDE750B2b', 'taker': '0x0000000000000000000000000000000000000000', 'tokenId': '93025177978745967226369398316375153283719303181694312089956059680730874301533', 'makerAmount': '50000', 'takerAmount': '5000000', 'expiration': '0', 'nonce': '0', 'feeRateBps': '0', 'side': 'BUY', 'signatureType': 0, 'signature': '0x0feca28666283824c27d7bead0bc441dde6df20dd71ef5ff7c84d3d1d5bf8aa4296fa382769dc11a92abe05b6f731d6c32556e9b4fb29e6eb50131af23a9ac941c'}, 'owner': 'PRIVATEKEY', 'orderType': 'GTC'}
]
```

## Get Order

Get information about an existing order

**This endpoint requires a L2 Header. **

Get single order by id.

```
GET /<clob-endpoint>/data/order/<order_hash>
```

Request Parameters
Name	Required	Type	Description
id	no	string	id of order to get information about

Response Format
Name	Type	Description
order	OpenOrder	order if it exists
An OpenOrder object is of the form:
Name	Type	Description
associate_trades	string[]	any Trade id the order has been partially included in
id	string	order id
status	string	order current status
market	string	market id (condition id)
original_size	string	original order size at placement
outcome	string	human readable outcome the order is for
maker_address	string	maker address (funder)
owner	string	api key
price	string	price
side	string	buy or sell
size_matched	string	size of order that has been matched/filled
asset_id	string	token id
expiration	string	unix timestamp when the order expired, 0 if it does not expire
type	string	order type (GTC, FOK, GTD)
created_at	string	unix timestamp when the order was created

## Get Active Orders
 **This endpoint requires a L2 Header. **

 Get active order(s) for a specific market.

 ```
 GET /<clob-endpoint>/data/orders
 ```

 Request Parameters
Name	Required	Type	Description
id	no	string	id of order to get information about
market	no	string	condition id of market
asset_id	no	string	id of the asset/token

Response Format
Name	Type	Description
null	OpenOrder[]	list of open orders filtered by the query parameters

## Check Order Reward Scoring
Check if an order is eligble or scoring for Rewards purposes
 **This endpoint requires a L2 Header. **

 Returns a boolean value where it is indicated if an order is scoring or not.

 ```
 GET /<clob-endpoint>/order-scoring?order_id={...}
 ```

 Request Parameters
Name	Required	Type	Description
orderId	yes	string	id of order to get information about

Response Format
Name	Type	Description
null	OrdersScoring	order scoring data
An OrdersScoring object is of the form:
Name	Type	Description
scoring	boolean	indicates if the order is scoring or not

## Check if some orders are scoring

 **This endpoint requires a L2 Header. **

 Returns to a dictionary with boolean value where it is indicated if an order is scoring or not.

 ```
 POST /<clob-endpoint>/orders-scoring
 ```

 Request Parameters
Name	Required	Type	Description
orderIds	yes	string[]	ids of the orders to get information about

Response Format
Name	Type	Description
null	OrdersScoring	orders scoring data
An OrdersScoring object is a dictionary that indicates the order by if it score.

## Cancel Order(s)
Multiple endpoints to cancel a single order, multiple orders, all orders or all orders from a single market.

### Cancel a single order
 **This endpoint requires a L2 Header. **

 ```
 DELETE /<clob-endpoint>/order
 ```
​
Request Payload Parameters
Name	Required	Type	Description
orderID	yes	string	ID of order to cancel

Response Format
Name	Type	Description
canceled	string[]	list of canceled orders
not_canceled		a order id -> reason map that explains why that order couldn’t be canceled

### Cancel multiple orders
 **This endpoint requires a L2 Header. **

 ```
 DELETE /<clob-endpoint>/orders
 ```

 Request Payload Parameters
Name	Required	Type	Description
null	yes	string[]	IDs of the orders to cancel
​
Response Format
Name	Type	Description
canceled	string[]	list of canceled orders
not_canceled		a order id -> reason map that explains why that order couldn’t be canceled

### Cancel all orders
 **This endpoint requires a L2 Header. **

Cancel all open orders posted by a user.

```
DELETE /<clob-endpoint>/cancel-all
```

Response Format
Name	Type	Description
canceled	string[]	list of canceled orders
not_canceled		a order id -> reason map that explains why that order couldn’t be canceled

### Cancel orders from market
 **This endpoint requires a L2 Header. **
 
 Cancel orders from market.

 ```
 DELETE /<clob-endpoint>/cancel-market-orders
 ```

 Request Payload Parameters
Name	Required	Type	Description
market	no	string	condition id of the market
asset_id	no	string	id of the asset/token

Response Format
Name	Type	Description
canceled	string[]	list of canceled orders
not_canceled		a order id -> reason map that explains why that order couldn’t be canceled

### Onchain Order Info
​
#### How do I interpret the OrderFilled onchain event?
Given an OrderFilled event:
    - orderHash: a unique hash for the Order being filled
    - maker: the user generating the order and the source of funds for the order
    - taker: the user filling the order OR the Exchange contract if the order fills multiple limit orders
    - makerAssetId: id of the asset that is given out. If 0, indicates that the Order is a BUY, giving USDC in exchange for Outcome tokens. Else, indicates that the Order is a SELL, giving Outcome tokens in exchange for USDC.
    - takerAssetId: id of the asset that is received. If 0, indicates that the Order is a SELL, receiving USDC in exchange for Outcome tokens. Else, indicates that the Order is a BUY, receiving Outcome tokens in exchange for USDC.
    - makerAmountFilled: the amount of the asset that is given out.
    - takerAmountFilled: the amount of the asset that is received.
    - fee: the fees paid by the order maker
