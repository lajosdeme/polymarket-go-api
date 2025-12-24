# Trades Overview
​
## Overview
All historical trades can be fetched via the Polymarket CLOB REST API. A trade is initiated by a “taker” who creates a marketable limit order. This limit order can be matched against one or more resting limit orders on the associated book. A trade can be in various states as described below. Note: in some cases (due to gas limitations) the execution of a “trade” must be broken into multiple transactions which case separate trade entities will be returned. To associate trade entities, there is a bucket_index field and a match_time field. Trades that have been broken into multiple trade objects can be reconciled by combining trade objects with the same market_order_id, match_time and incrementing bucket_index’s into a top level “trade” client side.

## Statuses
Status	Terminal?	Description
MATCHED	no	trade has been matched and sent to the executor service by the operator, the executor service submits the trade as a transaction to the Exchange contract
MINED	no	trade is observed to be mined into the chain, no finality threshold established
CONFIRMED	yes	trade has achieved strong probabilistic finality and was successful
RETRYING	no	trade transaction has failed (revert or reorg) and is being retried/resubmitted by the operator
FAILED	yes	trade has failed and is not being retried

## Get trades
** This endpoint requires a L2 Header. **

Get trades for the authenticated user based on the provided filters.

HTTP REQUEST 
```
GET /<clob-endpoint>/data/trades
​```

Request Parameters
Name	Required	Type	Description
id	no	string	id of trade to fetch
taker	no	string	address to get trades for where it is included as a taker
maker	no	string	address to get trades for where it is included as a maker
market	no	string	market for which to get the trades (condition ID)
before	no	string	unix timestamp representing the cutoff up to which trades that happened before then can be included
after	no	string	unix timestamp representing the cutoff for which trades that happened after can be included

Response Format
Name	Type	Description
null	Trade[]	list of trades filtered by query parameters

A Trade object is of the form:
Name	Type	Description
id	string	trade id
taker_order_id	string	hash of taker order (market order) that catalyzed the trade
market	string	market id (condition id)
asset_id	string	asset id (token id) of taker order (market order)
side	string	buy or sell
size	string	size
fee_rate_bps	string	the fees paid for the taker order expressed in basic points
price	string	limit price of taker order
status	string	trade status (see above)
match_time	string	time at which the trade was matched
last_update	string	timestamp of last status update
outcome	string	human readable outcome of the trade
maker_address	string	funder address of the taker of the trade
owner	string	api key of taker of the trade
transaction_hash	string	hash of the transaction where the trade was executed
bucket_index	integer	index of bucket for trade in case trade is executed in multiple transactions
maker_orders	MakerOrder[]	list of the maker trades the taker trade was filled against
type	string	side of the trade: TAKER or MAKER

A MakerOrder object is of the form:
Name	Type	Description
order_id	string	id of maker order
maker_address	string	maker address of the order
owner	string	api key of the owner of the order
matched_amount	string	size of maker order consumed with this trade
fee_rate_bps	string	the fees paid for the taker order expressed in basic points
price	string	price of maker order
asset_id	string	token/asset id
outcome	string	human readable outcome of the maker order
side	string	the side of the maker order. Can be buy or sell