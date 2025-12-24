# Orderbook
## Get order book summary

Retrieves the order book summary for a specific token
```
GET /book
```

Get order book summary

```
curl --request GET \
  --url https://clob.polymarket.com/book
```

```
{
  "market": "0x1b6f76e5b8587ee896c35847e12d11e75290a8c3934c5952e8a9d6e4c6f03cfa",
  "asset_id": "1234567890",
  "timestamp": "2023-10-01T12:00:00Z",
  "hash": "0xabc123def456...",
  "bids": [
    {
      "price": "1800.50",
      "size": "10.5"
    }
  ],
  "asks": [
    {
      "price": "1800.50",
      "size": "10.5"
    }
  ],
  "min_order_size": "0.001",
  "tick_size": "0.01",
  "neg_risk": false
}
```

Query Parameters
​
`token_id`
string
required
The unique identifier for the token

Response

Successful response
​
`market`
string
required

Market identifier
Example:

"0x1b6f76e5b8587ee896c35847e12d11e75290a8c3934c5952e8a9d6e4c6f03cfa"
​
`asset_id`
string
required

Asset identifier
Example:

"1234567890"
​
`timestamp`
string<date-time>
required

Timestamp of the order book snapshot
Example:

"2023-10-01T12:00:00Z"
​
`hash`
string
required

Hash of the order book state
Example:

"0xabc123def456..."
​
`bids`
object[]
required

Array of bid levels

Show child attributes
​
`asks`
object[]
required

Array of ask levels

child attributes:
`asks.price`
string
required

Price level (as string to maintain precision)
Example:

"1800.50"
​
`asks.size`
string
required

Total size at this price level
Example:

"10.5"
​
`min_order_size`
string
required

Minimum order size for this market
Example:

"0.001"
​
`tick_size`
string
required

Minimum price increment
Example:

"0.01"
​
`neg_risk`
boolean
required

Whether negative risk is enabled
Example:

false

## Get multiple order books summaries by request
Retrieves order book summaries for specified tokens via POST request
```
POST https://clob.polymarket.com/books
```

```
curl --request POST \
  --url https://clob.polymarket.com/books \
  --header 'Content-Type: application/json' \
  --data '
[
  {
    "token_id": "1234567890"
  },
  {
    "token_id": "0987654321"
  }
]
'
```

```
[
  {
    "market": "0x1b6f76e5b8587ee896c35847e12d11e75290a8c3934c5952e8a9d6e4c6f03cfa",
    "asset_id": "1234567890",
    "timestamp": "2023-10-01T12:00:00Z",
    "hash": "0xabc123def456...",
    "bids": [
      {
        "price": "1800.50",
        "size": "10.5"
      }
    ],
    "asks": [
      {
        "price": "1800.50",
        "size": "10.5"
      }
    ],
    "min_order_size": "0.001",
    "tick_size": "0.01",
    "neg_risk": false
  }
]
```

Body
application/json
Maximum array length: 500
​
token_id
string
required

The unique identifier for the token
Example:

"1234567890"
​
side
enum<string>

Optional side parameter for certain operations
Available options: BUY, 
SELL 
Example:

"BUY"
Response

Successful response
​
market
string
required

Market identifier
Example:

"0x1b6f76e5b8587ee896c35847e12d11e75290a8c3934c5952e8a9d6e4c6f03cfa"
​
asset_id
string
required

Asset identifier
Example:

"1234567890"
​
timestamp
string<date-time>
required

Timestamp of the order book snapshot
Example:

"2023-10-01T12:00:00Z"
​
hash
string
required

Hash of the order book state
Example:

"0xabc123def456..."
​
bids
object[]
required

Array of bid levels
Child attributes:
bids.price
string
required

Price level (as string to maintain precision)
Example:

"1800.50"
​
bids.size
string
required

Total size at this price level
Example:

"10.5"

asks
object[]
required

Array of ask levels
Child attributes:
asks.price
string
required

Price level (as string to maintain precision)
Example:

"1800.50"
​
asks.size
string
required

Total size at this price level
Example:

"10.5"

min_order_size
string
required

Minimum order size for this market
Example:

"0.001"
​
tick_size
string
required

Minimum price increment
Example:

"0.01"
​
neg_risk
boolean
required

Whether negative risk is enabled
Example:

false