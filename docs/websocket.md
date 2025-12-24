# WSS Overview

Overview and general information about the Polymarket Websocket
​
## Overview
The Polymarket CLOB API provides websocket (wss) channels through which clients can get pushed updates. These endpoints allow clients to maintain almost real-time views of their orders, their trades and markets in general. There are two available channels user and market.

## Subscription
To subscribe send a message including the following authentication and intent information upon opening the connection.
Field	Type	Description
auth	Auth	see next page for auth information
markets	string[]	array of markets (condition IDs) to receive events for (for user channel)
assets_ids	string[]	array of asset ids (token IDs) to receive events for (for market channel)
type	string	id of channel to subscribe to (USER or MARKET)

Where the auth field is of type Auth which has the form described in the WSS Authentication section below.

### Subscribe to more assets
Once connected, the client can subscribe and unsubscribe to asset_ids by sending the following message:
Field	Type	Description
assets_ids	string[]	array of asset ids (token IDs) to receive events for (for market channel)
operation	string	”subscribe” or “unsubscribe”

# WSS Quickstart
The following code samples and explanation will show you how to subsribe to the Marker and User channels of the Websocket. You’ll need your API keys to do this so we’ll start with that.

## Getting your API keys
```
from py_clob_client.client import ClobClient

host: str = "https://clob.polymarket.com"
key: str = "" #This is your Private Key. If using email login export from https://reveal.magic.link/polymarket otherwise export from your Web3 Application
chain_id: int = 137 #No need to adjust this
POLYMARKET_PROXY_ADDRESS: str = '' #This is the address you deposit/send USDC to to FUND your Polymarket account.

#Select from the following 3 initialization options to matches your login method, and remove any unused lines so only one client is initialized.

### Initialization of a client using a Polymarket Proxy associated with an Email/Magic account. If you login with your email use this example.
client = ClobClient(host, key=key, chain_id=chain_id, signature_type=1, funder=POLYMARKET_PROXY_ADDRESS)

### Initialization of a client using a Polymarket Proxy associated with a Browser Wallet(Metamask, Coinbase Wallet, etc)
client = ClobClient(host, key=key, chain_id=chain_id, signature_type=2, funder=POLYMARKET_PROXY_ADDRESS)

### Initialization of a client that trades directly from an EOA. 
client = ClobClient(host, key=key, chain_id=chain_id)

print( client.derive_api_key() )
```

## Using those keys to connect to the Market or User Websocket
```
from websocket import WebSocketApp
import json
import time
import threading

MARKET_CHANNEL = "market"
USER_CHANNEL = "user"


class WebSocketOrderBook:
    def __init__(self, channel_type, url, data, auth, message_callback, verbose):
        self.channel_type = channel_type
        self.url = url
        self.data = data
        self.auth = auth
        self.message_callback = message_callback
        self.verbose = verbose
        furl = url + "/ws/" + channel_type
        self.ws = WebSocketApp(
            furl,
            on_message=self.on_message,
            on_error=self.on_error,
            on_close=self.on_close,
            on_open=self.on_open,
        )
        self.orderbooks = {}

    def on_message(self, ws, message):
        print(message)
        pass

    def on_error(self, ws, error):
        print("Error: ", error)
        exit(1)

    def on_close(self, ws, close_status_code, close_msg):
        print("closing")
        exit(0)

    def on_open(self, ws):
        if self.channel_type == MARKET_CHANNEL:
            ws.send(json.dumps({"assets_ids": self.data, "type": MARKET_CHANNEL}))
        elif self.channel_type == USER_CHANNEL and self.auth:
            ws.send(
                json.dumps(
                    {"markets": self.data, "type": USER_CHANNEL, "auth": self.auth}
                )
            )
        else:
            exit(1)

        thr = threading.Thread(target=self.ping, args=(ws,))
        thr.start()


    def subscribe_to_tokens_ids(self, assets_ids):
        if self.channel_type == MARKET_CHANNEL:
            self.ws.send(json.dumps({"assets_ids": assets_ids, "operation": "subscribe"}))

    def unsubscribe_to_tokens_ids(self, assets_ids):
        if self.channel_type == MARKET_CHANNEL:
            self.ws.send(json.dumps({"assets_ids": assets_ids, "operation": "unsubscribe"}))


    def ping(self, ws):
        while True:
            ws.send("PING")
            time.sleep(10)

    def run(self):
        self.ws.run_forever()


if __name__ == "__main__":
    url = "wss://ws-subscriptions-clob.polymarket.com"
    #Complete these by exporting them from your initialized client. 
    api_key = ""
    api_secret = ""
    api_passphrase = ""

    asset_ids = [
        "109681959945973300464568698402968596289258214226684818748321941747028805721376",
    ]
    condition_ids = [] # no really need to filter by this one

    auth = {"apiKey": api_key, "secret": api_secret, "passphrase": api_passphrase}

    market_connection = WebSocketOrderBook(
        MARKET_CHANNEL, url, asset_ids, auth, None, True
    )
    user_connection = WebSocketOrderBook(
        USER_CHANNEL, url, condition_ids, auth, None, True
    )

    market_connection.subscribe_to_tokens_ids(["123"])
    # market_connection.unsubscribe_to_tokens_ids(["123"])

    market_connection.run()
    # user_connection.run()
    ```

# WSS Authentication
 **Only connections to user channel require authentication. **

Field	Optional	Description
apikey	yes	Polygon account’s CLOB api key
secret	yes	Polygon account’s CLOB api secret
passphrase	yes	Polygon account’s CLOB api passphrase

## User Channel
Authenticated channel for updates related to user activities (orders, trades), filtered for authenticated user by apikey.

SUBSCRIBE
```
<wss-channel> user
```

### Trade Message
Emitted when:

    - when a market order is matched (“MATCHED”)
    - when a limit order for the user is included in a trade (“MATCHED”)
    - subsequent status changes for trade (“MINED”, “CONFIRMED”, “RETRYING”, “FAILED”)

Structure
Name	Type	Description
asset_id	string	asset id (token ID) of order (market order)
event_type	string	”trade”
id	string	trade id
last_update	string	time of last update to trade
maker_orders	MakerOrder[]	array of maker order details
market	string	market identifier (condition ID)
matchtime	string	time trade was matched
outcome	string	outcome
owner	string	api key of event owner
price	string	price
side	string	BUY/SELL
size	string	size
status	string	trade status
taker_order_id	string	id of taker order
timestamp	string	time of event
trade_owner	string	api key of trade owner
type	string	”TRADE”

Where a MakerOrder object is of the form:
Name	Type	Description
asset_id	string	asset of the maker order
matched_amount	string	amount of maker order matched in trade
order_id	string	maker order ID
outcome	string	outcome
owner	string	owner of maker order
price	string	price of maker order

```
{
  "asset_id": "52114319501245915516055106046884209969926127482827954674443846427813813222426",
  "event_type": "trade",
  "id": "28c4d2eb-bbea-40e7-a9f0-b2fdb56b2c2e",
  "last_update": "1672290701",
  "maker_orders": [
    {
      "asset_id": "52114319501245915516055106046884209969926127482827954674443846427813813222426",
      "matched_amount": "10",
      "order_id": "0xff354cd7ca7539dfa9c28d90943ab5779a4eac34b9b37a757d7b32bdfb11790b",
      "outcome": "YES",
      "owner": "9180014b-33c8-9240-a14b-bdca11c0a465",
      "price": "0.57"
    }
  ],
  "market": "0xbd31dc8a20211944f6b70f31557f1001557b59905b7738480ca09bd4532f84af",
  "matchtime": "1672290701",
  "outcome": "YES",
  "owner": "9180014b-33c8-9240-a14b-bdca11c0a465",
  "price": "0.57",
  "side": "BUY",
  "size": "10",
  "status": "MATCHED",
  "taker_order_id": "0x06bc63e346ed4ceddce9efd6b3af37c8f8f440c92fe7da6b2d0f9e4ccbc50c42",
  "timestamp": "1672290701",
  "trade_owner": "9180014b-33c8-9240-a14b-bdca11c0a465",
  "type": "TRADE"
}
```

### Order Message
Emitted when:
    - When an order is placed (PLACEMENT)
    - When an order is updated (some of it is matched) (UPDATE)
    - When an order is canceled (CANCELLATION)

Structure
Name	Type	Description
asset_id	string	asset ID (token ID) of order
associate_trades	string[]	array of ids referencing trades that the order has been included in
event_type	string	”order”
id	string	order id
market	string	condition ID of market
order_owner	string	owner of order
original_size	string	original order size
outcome	string	outcome
owner	string	owner of orders
price	string	price of order
side	string	BUY/SELL
size_matched	string	size of order that has been matched
timestamp	string	time of event
type	string	PLACEMENT/UPDATE/CANCELLATION

```
{
  "asset_id": "52114319501245915516055106046884209969926127482827954674443846427813813222426",
  "associate_trades": null,
  "event_type": "order",
  "id": "0xff354cd7ca7539dfa9c28d90943ab5779a4eac34b9b37a757d7b32bdfb11790b",
  "market": "0xbd31dc8a20211944f6b70f31557f1001557b59905b7738480ca09bd4532f84af",
  "order_owner": "9180014b-33c8-9240-a14b-bdca11c0a465",
  "original_size": "10",
  "outcome": "YES",
  "owner": "9180014b-33c8-9240-a14b-bdca11c0a465",
  "price": "0.57",
  "side": "SELL",
  "size_matched": "0",
  "timestamp": "1672290687",
  "type": "PLACEMENT"
}
```
## Market Channel
Public channel for updates related to market updates (level 2 price data).

SUBSCRIBE
```
<wss-channel> market
```

### Book Message
Emitted When:
    - First subscribed to a market
    - When there is a trade that affects the book

Structure
Name	Type	Description
event_type	string	”book”
asset_id	string	asset ID (token ID)
market	string	condition ID of market
timestamp	string	unix timestamp the current book generation in milliseconds (1/1,000 second)
hash	string	hash summary of the orderbook content
buys	OrderSummary[]	list of type (size, price) aggregate book levels for buys
sells	OrderSummary[]	list of type (size, price) aggregate book levels for sells

Where a OrderSummary object is of the form:
Name	Type	Description
price	string	size available at that price level
size	string	price of the orderbook level

```
{
  "event_type": "book",
  "asset_id": "65818619657568813474341868652308942079804919287380422192892211131408793125422",
  "market": "0xbd31dc8a20211944f6b70f31557f1001557b59905b7738480ca09bd4532f84af",
  "bids": [
    { "price": ".48", "size": "30" },
    { "price": ".49", "size": "20" },
    { "price": ".50", "size": "15" }
  ],
  "asks": [
    { "price": ".52", "size": "25" },
    { "price": ".53", "size": "60" },
    { "price": ".54", "size": "10" }
  ],
  "timestamp": "123456789000",
  "hash": "0x0...."
}
```

### price_change Message
⚠️ Breaking Change Notice: The price_change message schema will be updated on September 15, 2025 at 11 PM UTC. Please see the migration guide for details.
Emitted When:
    - A new order is placed
    - An order is cancelled

Structure
Name	Type	Description
event_type	string	”price_change”
market	string	condition ID of market
price_changes	PriceChange[]	array of price change objects
timestamp	string	unix timestamp in milliseconds

Where a PriceChange object is of the form:
Name	Type	Description
asset_id	string	asset ID (token ID)
price	string	price level affected
size	string	new aggregate size for price level
side	string	”BUY” or “SELL”
hash	string	hash of the order
best_bid	string	current best bid price
best_ask	string	current best ask price

```
{
    "market": "0x5f65177b394277fd294cd75650044e32ba009a95022d88a0c1d565897d72f8f1",
    "price_changes": [
        {
            "asset_id": "71321045679252212594626385532706912750332728571942532289631379312455583992563",
            "price": "0.5",
            "size": "200",
            "side": "BUY",
            "hash": "56621a121a47ed9333273e21c83b660cff37ae50",
            "best_bid": "0.5",
            "best_ask": "1"
        },
        {
            "asset_id": "52114319501245915516055106046884209969926127482827954674443846427813813222426",
            "price": "0.5",
            "size": "200",
            "side": "SELL",
            "hash": "1895759e4df7a796bf4f1c5a5950b748306923e2",
            "best_bid": "0",
            "best_ask": "0.5"
        }
    ],
    "timestamp": "1757908892351",
    "event_type": "price_change"
}
```

### tick_size_change Message
Emitted When:

    - The minimum tick size of the market changes. This happens when the book’s price reaches the limits: price > 0.96 or price < 0.04

Structure
Name	Type	Description
event_type	string	”price_change”
asset_id	string	asset ID (token ID)
market	string	condition ID of market
old_tick_size	string	previous minimum tick size
new_tick_size	string	current minimum tick size
side	string	buy/sell
timestamp	string	time of event

```
{
"event_type": "tick_size_change",
"asset_id": "65818619657568813474341868652308942079804919287380422192892211131408793125422",\
"market": "0xbd31dc8a20211944f6b70f31557f1001557b59905b7738480ca09bd4532f84af",
"old_tick_size": "0.01",
"new_tick_size": "0.001",
"timestamp": "100000000"
}
```

### last_trade_price Message
Emitted When:
    - When a maker and taker order is matched creating a trade event.

```
{
"asset_id":"114122071509644379678018727908709560226618148003371446110114509806601493071694",
"event_type":"last_trade_price",
"fee_rate_bps":"0",
"market":"0x6a67b9d828d53862160e470329ffea5246f338ecfffdf2cab45211ec578b0347",
"price":"0.456",
"side":"BUY",
"size":"219.217767",
"timestamp":"1750428146322"
}
```



