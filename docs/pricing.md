# Pricing
## Get market price
Retrieves the market price for a specific token and side

```
GET https://clob.polymarket.com/price
```

```
curl --request GET \
  --url https://clob.polymarket.com/price
  ```

  ```
  {
  "price": "1800.50"
}
```

Query Parameters
​
token_id
string
required

The unique identifier for the token
​
side
enum<string>
required

The side of the market (BUY or SELL)
Available options: BUY, 
SELL 
Response

Successful response
​
price
string
required

The market price (as string to maintain precision)
Example:

"1800.50"

## Get multiple market prices
Retrieves market prices for multiple tokens and sides

```
https://clob.polymarket.com/prices
```

```
curl --request GET \
  --url https://clob.polymarket.com/prices
  ```

```
{
  "1234567890": {
    "BUY": "1800.50",
    "SELL": "1801.00"
  },
  "0987654321": {
    "BUY": "50.25",
    "SELL": "50.30"
  }
}
```

Successful response

Map of token_id to side to price
​
{key}
object
Child attributes:
{key}.{key}
string

## Get multiple market prices by request
Retrieves market prices for specified tokens and sides via POST request

```
POST https://clob.polymarket.com/prices
```

```
curl --request POST \
  --url https://clob.polymarket.com/prices \
  --header 'Content-Type: application/json' \
  --data '
[
  {
    "token_id": "1234567890",
    "side": "BUY"
  },
  {
    "token_id": "0987654321",
    "side": "SELL"
  }
]
'
```

```
{
  "1234567890": {
    "BUY": "1800.50",
    "SELL": "1801.00"
  },
  "0987654321": {
    "BUY": "50.25",
    "SELL": "50.30"
  }
}
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
required

The side of the market (BUY or SELL)
Available options: BUY, 
SELL 
Example:

"BUY"

Response

Successful response

Map of token_id to side to price

## Get midpoint price

Retrieves the midpoint price for a specific token

```
GET https://clob.polymarket.com/midpoint
```

```
curl --request GET \
  --url https://clob.polymarket.com/midpoint
  ```

```
{
  "mid": "1800.75"
}
```

Query Parameters
​
token_id
string
required

The unique identifier for the token
Response

Successful response
​
mid
string
required

The midpoint price (as string to maintain precision)
Example:

"1800.75"

## Get price history for a traded token
Fetches historical price data for a specified market token

```
POST https://clob.polymarket.com/prices-history
```

```
curl --request GET \
  --url https://clob.polymarket.com/prices-history
  ```

```
{
  "history": [
    {
      "t": 1697875200,
      "p": 1800.75
    }
  ]
}
```

Query Parameters
​
market
string
required

The CLOB token ID for which to fetch price history
​
startTs
number

The start time, a Unix timestamp in UTC
​
endTs
number

The end time, a Unix timestamp in UTC
​
interval
enum<string>

A string representing a duration ending at the current time. Mutually exclusive with startTs and endTs
Available options: 1m, 
1w, 
1d, 
6h, 
1h, 
max 
​
fidelity
number

The resolution of the data, in minutes
Response

A list of timestamp/price pairs
​
history
object[]
required
Child attributes:
history.t
number
required

UTC timestamp
Example:

1697875200
​
history.p
number
required

Price
Example:

1800.75