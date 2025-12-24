# Historical Timeseries Data
## Historical Timeseries Data

Fetches historical price data for a specified market token.

```
GET https://clob.polymarket.com/prices-history
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

The CLOB provides detailed price history for each traded token.

HTTP REQUEST GET /<clob-endpoint>/prices-history
We also have a Interactive Notebook to visualize the data from this endpoint available [here](https://colab.research.google.com/drive/1s4TCOR4K7fRP7EwAH1YmOactMakx24Cs?usp=sharing#scrollTo=mYCJBcfB9Zu4).

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

