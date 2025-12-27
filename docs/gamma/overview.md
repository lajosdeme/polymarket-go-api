# Overview
All market data necessary for market resolution is available on-chain (ie ancillaryData in UMA 00 request), but Polymarket also provides a hosted service, Gamma, that indexes this data and provides additional market metadata (ie categorization, indexed volume, etc). This service is made available through a REST API. For public users, this resource read only and can be used to fetch useful information about markets for things like non-profit research projects, alternative trading interfaces, automated trading systems etc.

## Endpoint
https://gamma-api.polymarket.com

## Gamma Structure
Gamma provides some organizational models. These include events, and markets. The most fundamental element is always markets and the other models simply provide additional organization.
​
### Detail
    Market
        Contains data related to a market that is traded on. Maps onto a pair of clob token ids, a market address, a question id and a condition id
    Event
        Contains a set of markets
        Variants:
            Event with 1 market (i.e., resulting in an SMP)
            Event with 2 or more markets (i.e., resulting in an GMP)

​
### Example
    [Event] Where will Barron Trump attend College?
        [Market] Will Barron attend Georgetown?
        [Market] Will Barron attend NYU?
        [Market] Will Barron attend UPenn?
        [Market] Will Barron attend Harvard?
        [Market] Will Barron attend another college?

## How to Fetch Markets
Both the getEvents and getMarkets are paginated. See pagination section for details.
This guide covers the three recommended approaches for fetching market data from the Gamma API, each optimized for different use cases.
​
### Overview
There are three main strategies for retrieving market data:

    By Slug - Best for fetching specific individual markets or events
    By Tags - Ideal for filtering markets by category or sport
    Via Events Endpoint - Most efficient for retrieving all active markets

### 1. Fetch by Slug
Use Case: When you need to retrieve a specific market or event that you already know about. Individual markets and events are best fetched using their unique slug identifier. The slug can be found directly in the Polymarket frontend URL.
​
How to Extract the Slug
From any Polymarket URL, the slug is the path segment after /event/ or /market/:
```
https://polymarket.com/event/fed-decision-in-october?tid=1758818660485
                            ↑
                  Slug: fed-decision-in-october
```

API Endpoints
For Events: 
```
GET /events/slug/ 
```

For Markets: 
```
GET /markets/slug/
```

#### Example
```
curl "https://gamma-api.polymarket.com/events/slug/fed-decision-in-october"
```

## 2. Fetch by Tags
Use Case: When you want to filter markets by category, sport, or topic. Tags provide a powerful way to categorize and filter markets. You can discover available tags and then use them to filter your market requests.
​
### Discover Available Tags
General Tags: GET /tags Sports Tags & Metadata: GET /sports The /sports endpoint returns comprehensive metadata for sports including tag IDs, images, resolution sources, and series information.
​
### Using Tags in Market Requests
Once you have tag IDs, you can use them with the tag_id parameter in both markets and events endpoints. 
Markets with Tags: 
```
GET /markets 
```
Events with Tags: 
```
GET /events
```

```
curl "https://gamma-api.polymarket.com/events?tag_id=100381&limit=1&closed=false"
```

Additional Tag Filtering
You can also:

    Use related_tags=true to include related tag markets
    Exclude specific tags with exclude_tag_id


## 3. Fetch All Active Markets
Use Case: When you need to retrieve all available active markets, typically for broader analysis or market discovery. The most efficient approach is to use the /events endpoint and work backwards, as events contain their associated markets. Events Endpoint: GET /events Markets Endpoint: GET /markets
​
### Key Parameters

    order=id - Order by event ID
    ascending=false - Get newest events first
    closed=false - Only active markets
    limit - Control response size
    offset - For pagination

### Example
```
curl "https://gamma-api.polymarket.com/events?order=id&ascending=false&closed=false&limit=100"
```

This approach gives you all active markets ordered from newest to oldest, allowing you to systematically process all available trading opportunities.

### Pagination
For large datasets, use pagination with limit and offset parameters:

    limit=50 - Return 50 results per page
    offset=0 - Start from the beginning (increment by limit for subsequent pages)

#### Example
```
# Page 1: First 50 results (offset=0)
curl "https://gamma-api.polymarket.com/events?order=id&ascending=false&closed=false&limit=50&offset=0"

# Page 2: Next 50 results (offset=50)
curl "https://gamma-api.polymarket.com/events?order=id&ascending=false&closed=false&limit=50&offset=50"

# Page 3: Next 50 results (offset=100)
curl "https://gamma-api.polymarket.com/events?order=id&ascending=false&closed=false&limit=50&offset=100"

# Paginating through markets with tag filtering
curl "https://gamma-api.polymarket.com/markets?tag_id=100381&closed=false&limit=25&offset=0"

# Next page of markets with tag filtering
curl "https://gamma-api.polymarket.com/markets?tag_id=100381&closed=false&limit=25&offset=25"
```

### Best Practices

    For Individual Markets: Always use the slug method for best performance
    For Category Browsing: Use tag filtering to reduce API calls
    For Complete Market Discovery: Use the events endpoint with pagination
    Always Include closed=false: Unless you specifically need historical data
    Implement Rate Limiting: Respect API limits for production applications

## Health check
```
curl --request GET \
  --url https://data-api.polymarket.com/
  ```

Response:
```
{
  "data": "OK"
}
```