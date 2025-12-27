# Health check
```
curl --request GET \
  --url https://data-api.polymarket.com/
  ```

  ```
  {
  "data": "OK"
}
```

# Get current positions for a user
```
curl --request GET \
  --url 'https://data-api.polymarket.com/positions?sizeThreshold=1&limit=100&sortBy=TOKENS&sortDirection=DESC'
  ```

  ```
  [
  {
    "proxyWallet": "0x56687bf447db6ffa42ffe2204a05edaa20f55839",
    "asset": "<string>",
    "conditionId": "0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917",
    "size": 123,
    "avgPrice": 123,
    "initialValue": 123,
    "currentValue": 123,
    "cashPnl": 123,
    "percentPnl": 123,
    "totalBought": 123,
    "realizedPnl": 123,
    "percentRealizedPnl": 123,
    "curPrice": 123,
    "redeemable": true,
    "mergeable": true,
    "title": "<string>",
    "slug": "<string>",
    "icon": "<string>",
    "eventSlug": "<string>",
    "outcome": "<string>",
    "outcomeIndex": 123,
    "oppositeOutcome": "<string>",
    "oppositeAsset": "<string>",
    "endDate": "<string>",
    "negativeRisk": true
  }
]
```

Query Parameters
​
user
string
required

User address (required)
User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
market
string[]

Comma-separated list of condition IDs. Mutually exclusive with eventId.

0x-prefixed 64-hex string
​
eventId
integer[]

Comma-separated list of event IDs. Mutually exclusive with market.
Required range: x >= 1
​
sizeThreshold
number
default:1
Required range: x >= 0
​
redeemable
boolean
default:false
​
mergeable
boolean
default:false
​
limit
integer
default:100
Required range: 0 <= x <= 500
​
offset
integer
default:0
Required range: 0 <= x <= 10000
​
sortBy
enum<string>
default:TOKENS
Available options: CURRENT, 
INITIAL, 
TOKENS, 
CASHPNL, 
PERCENTPNL, 
TITLE, 
RESOLVING, 
PRICE, 
AVGPRICE 
​
sortDirection
enum<string>
default:DESC
Available options: ASC, 
DESC 
​
title
string
Maximum string length: 100

## Response
Success
​
proxyWallet
string

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
asset
string
​
conditionId
string

0x-prefixed 64-hex string
Example:

"0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917"
​
size
number
​
avgPrice
number
​
initialValue
number
​
currentValue
number
​
cashPnl
number
​
percentPnl
number
​
totalBought
number
​
realizedPnl
number
​
percentRealizedPnl
number
​
curPrice
number
​
redeemable
boolean
​
mergeable
boolean
​
title
string
​
slug
string
​
icon
string
​
eventSlug
string
​
outcome
string
​
outcomeIndex
integer
​
oppositeOutcome
string
​
oppositeAsset
string
​
endDate
string
​
negativeRisk
boolean

# Get trades for a user or markets
```
curl --request GET \
  --url 'https://data-api.polymarket.com/trades?limit=100&takerOnly=true'
  ```

  ```
  [
  {
    "proxyWallet": "0x56687bf447db6ffa42ffe2204a05edaa20f55839",
    "side": "BUY",
    "asset": "<string>",
    "conditionId": "0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917",
    "size": 123,
    "price": 123,
    "timestamp": 123,
    "title": "<string>",
    "slug": "<string>",
    "icon": "<string>",
    "eventSlug": "<string>",
    "outcome": "<string>",
    "outcomeIndex": 123,
    "name": "<string>",
    "pseudonym": "<string>",
    "bio": "<string>",
    "profileImage": "<string>",
    "profileImageOptimized": "<string>",
    "transactionHash": "<string>"
  }
]
```

Query Parameters
​
limit
integer
default:100
Required range: 0 <= x <= 10000
​
offset
integer
default:0
Required range: 0 <= x <= 10000
​
takerOnly
boolean
default:true
​
filterType
enum<string>

Must be provided together with filterAmount.
Available options: CASH, 
TOKENS 
​
filterAmount
number

Must be provided together with filterType.
Required range: x >= 0
​
market
string[]

Comma-separated list of condition IDs. Mutually exclusive with eventId.

0x-prefixed 64-hex string
​
eventId
integer[]

Comma-separated list of event IDs. Mutually exclusive with market.
Required range: x >= 1
​
user
string

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
side
enum<string>
Available options: BUY, 
SELL 

## Response
Success
​
proxyWallet
string

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
side
enum<string>
Available options: BUY, 
SELL 
​
asset
string
​
conditionId
string

0x-prefixed 64-hex string
Example:

"0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917"
​
size
number
​
price
number
​
timestamp
integer<int64>
​
title
string
​
slug
string
​
icon
string
​
eventSlug
string
​
outcome
string
​
outcomeIndex
integer
​
name
string
​
pseudonym
string
​
bio
string
​
profileImage
string
​
profileImageOptimized
string
​
transactionHash
string

# Get user activity
```
curl --request GET \
  --url 'https://data-api.polymarket.com/activity?limit=100&sortBy=TIMESTAMP&sortDirection=DESC'
  ```

  ```
  [
  {
    "proxyWallet": "0x56687bf447db6ffa42ffe2204a05edaa20f55839",
    "timestamp": 123,
    "conditionId": "0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917",
    "type": "TRADE",
    "size": 123,
    "usdcSize": 123,
    "transactionHash": "<string>",
    "price": 123,
    "asset": "<string>",
    "side": "BUY",
    "outcomeIndex": 123,
    "title": "<string>",
    "slug": "<string>",
    "icon": "<string>",
    "eventSlug": "<string>",
    "outcome": "<string>",
    "name": "<string>",
    "pseudonym": "<string>",
    "bio": "<string>",
    "profileImage": "<string>",
    "profileImageOptimized": "<string>"
  }
]
```

Query Parameters
​
limit
integer
default:100
Required range: 0 <= x <= 500
​
offset
integer
default:0
Required range: 0 <= x <= 10000
​
user
string
required

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
market
string[]

Comma-separated list of condition IDs. Mutually exclusive with eventId.

0x-prefixed 64-hex string
​
eventId
integer[]

Comma-separated list of event IDs. Mutually exclusive with market.
Required range: x >= 1
​
type
enum<string>[]
Available options: TRADE, 
SPLIT, 
MERGE, 
REDEEM, 
REWARD, 
CONVERSION 
​
start
integer
Required range: x >= 0
​
end
integer
Required range: x >= 0
​
sortBy
enum<string>
default:TIMESTAMP
Available options: TIMESTAMP, 
TOKENS, 
CASH 
​
sortDirection
enum<string>
default:DESC
Available options: ASC, 
DESC 
​
side
enum<string>
Available options: BUY, 
SELL 

## Response
Success
​
proxyWallet
string

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
timestamp
integer<int64>
​
conditionId
string

0x-prefixed 64-hex string
Example:

"0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917"
​
type
enum<string>
Available options: TRADE, 
SPLIT, 
MERGE, 
REDEEM, 
REWARD, 
CONVERSION 
​
size
number
​
usdcSize
number
​
transactionHash
string
​
price
number
​
asset
string
​
side
enum<string>
Available options: BUY, 
SELL 
​
outcomeIndex
integer
​
title
string
​
slug
string
​
icon
string
​
eventSlug
string
​
outcome
string
​
name
string
​
pseudonym
string
​
bio
string
​
profileImage
string
​
profileImageOptimized
string

# Get top holders for markets
```
curl --request GET \
  --url 'https://data-api.polymarket.com/holders?limit=20&minBalance=1'
  ```

```
[
  {
    "token": "<string>",
    "holders": [
      {
        "proxyWallet": "0x56687bf447db6ffa42ffe2204a05edaa20f55839",
        "bio": "<string>",
        "asset": "<string>",
        "pseudonym": "<string>",
        "amount": 123,
        "displayUsernamePublic": true,
        "outcomeIndex": 123,
        "name": "<string>",
        "profileImage": "<string>",
        "profileImageOptimized": "<string>"
      }
    ]
  }
]
```

Query Parameters
​
limit
integer
default:20

Maximum number of holders to return per token. Capped at 20.
Required range: 0 <= x <= 20
​
market
string[]
required

Comma-separated list of condition IDs.

0x-prefixed 64-hex string
​
minBalance
integer
default:1
Required range: 0 <= x <= 999999

## Response
Success
​
token
string
​
holders
object[]
Child attributes:
    holders.proxyWallet
    string

    User Profile Address (0x-prefixed, 40 hex chars)
    Example:

    "0x56687bf447db6ffa42ffe2204a05edaa20f55839"
    ​
    holders.bio
    string
    ​
    holders.asset
    string
    ​
    holders.pseudonym
    string
    ​
    holders.amount
    number
    ​
    holders.displayUsernamePublic
    boolean
    ​
    holders.outcomeIndex
    integer
    ​
    holders.name
    string
    ​
    holders.profileImage
    string
    ​
    holders.profileImageOptimized
    string

# Get total value of a user's positions
```
curl --request GET \
  --url https://data-api.polymarket.com/value
  ```

  ```
  [
  {
    "user": "0x56687bf447db6ffa42ffe2204a05edaa20f55839",
    "value": 123
  }
]
```

Query Parameters
​
user
string
required

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
market
string[]

0x-prefixed 64-hex string

## Response
Success
​
user
string

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
value
number

# Get closed positions for a user
Fetches closed positions for a user(address)
```
curl --request GET \
  --url 'https://data-api.polymarket.com/closed-positions?limit=10&sortBy=REALIZEDPNL&sortDirection=DESC'
  ```

  ```
  [
  {
    "proxyWallet": "0x56687bf447db6ffa42ffe2204a05edaa20f55839",
    "asset": "<string>",
    "conditionId": "0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917",
    "avgPrice": 123,
    "totalBought": 123,
    "realizedPnl": 123,
    "curPrice": 123,
    "timestamp": 123,
    "title": "<string>",
    "slug": "<string>",
    "icon": "<string>",
    "eventSlug": "<string>",
    "outcome": "<string>",
    "outcomeIndex": 123,
    "oppositeOutcome": "<string>",
    "oppositeAsset": "<string>",
    "endDate": "<string>"
  }
]
```

Query Parameters
​
user
string
required

The address of the user in question
User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
market
string[]

The conditionId of the market in question. Supports multiple csv separated values. Cannot be used with the eventId param.

0x-prefixed 64-hex string
​
title
string

Filter by market title
Maximum string length: 100
​
eventId
integer[]

The event id of the event in question. Supports multiple csv separated values. Returns positions for all markets for those event ids. Cannot be used with the market param.
Required range: x >= 1
​
limit
integer
default:10

The max number of positions to return
Required range: 0 <= x <= 50
​
offset
integer
default:0

The starting index for pagination
Required range: 0 <= x <= 100000
​
sortBy
enum<string>
default:REALIZEDPNL

The sort criteria
Available options: REALIZEDPNL, 
TITLE, 
PRICE, 
AVGPRICE, 
TIMESTAMP 
​
sortDirection
enum<string>
default:DESC

The sort direction
Available options: ASC, 
DESC 

## Response
Success
​
proxyWallet
string

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
asset
string
​
conditionId
string

0x-prefixed 64-hex string
Example:

"0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917"
​
avgPrice
number
​
totalBought
number
​
realizedPnl
number
​
curPrice
number
​
timestamp
integer<int64>
​
title
string
​
slug
string
​
icon
string
​
eventSlug
string
​
outcome
string
​
outcomeIndex
integer
​
oppositeOutcome
string
​
oppositeAsset
string
​
endDate
string

# Get trader leaderboard rankings
Returns trader leaderboard rankings filtered by category, time period, and ordering.

```
curl --request GET \
  --url 'https://data-api.polymarket.com/v1/leaderboard?category=OVERALL&timePeriod=DAY&orderBy=PNL&limit=25'
  ```

  ```
  [
  {
    "rank": "<string>",
    "proxyWallet": "0x56687bf447db6ffa42ffe2204a05edaa20f55839",
    "userName": "<string>",
    "vol": 123,
    "pnl": 123,
    "profileImage": "<string>",
    "xUsername": "<string>",
    "verifiedBadge": true
  }
]
```

Query Parameters
​
category
enum<string>
default:OVERALL

Market category for the leaderboard
Available options: OVERALL, 
POLITICS, 
SPORTS, 
CRYPTO, 
CULTURE, 
MENTIONS, 
WEATHER, 
ECONOMICS, 
TECH, 
FINANCE 
​
timePeriod
enum<string>
default:DAY

Time period for leaderboard results
Available options: DAY, 
WEEK, 
MONTH, 
ALL 
​
orderBy
enum<string>
default:PNL

Leaderboard ordering criteria
Available options: PNL, 
VOL 
​
limit
integer
default:25

Max number of leaderboard traders to return
Required range: 1 <= x <= 50
​
offset
integer
default:0

Starting index for pagination
Required range: 0 <= x <= 1000
​
user
string

Limit leaderboard to a single user by address
User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
userName
string

Limit leaderboard to a single username

## Response
Success
​
rank
string

The rank position of the trader
​
proxyWallet
string

User Profile Address (0x-prefixed, 40 hex chars)
Example:

"0x56687bf447db6ffa42ffe2204a05edaa20f55839"
​
userName
string

The trader's username
​
vol
number

Trading volume for this trader
​
pnl
number

Profit and loss for this trader
​
profileImage
string

URL to the trader's profile image
​
xUsername
string

The trader's X (Twitter) username
​
verifiedBadge
boolean

Whether the trader has a verified badge