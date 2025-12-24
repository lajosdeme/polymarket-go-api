# Spreads
## Get bid-ask spreads

Retrieves bid-ask spreads for multiple tokens

```
POST https://clob.polymarket.com/spreads
```

```
curl --request POST \
  --url https://clob.polymarket.com/spreads \
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
{
  "1234567890": "0.50",
  "0987654321": "0.05"
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

Optional side parameter for certain operations
Available options: BUY, 
SELL 
Example:

"BUY"

Response

Successful response

Map of token_id to spread value
​
{key}
string