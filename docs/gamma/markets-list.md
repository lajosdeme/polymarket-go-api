# List markets

```
curl --request GET \
  --url https://gamma-api.polymarket.com/markets
  ```

```
[
  {
    "id": "<string>",
    "question": "<string>",
    "conditionId": "<string>",
    "slug": "<string>",
    "twitterCardImage": "<string>",
    "resolutionSource": "<string>",
    "endDate": "2023-11-07T05:31:56Z",
    "category": "<string>",
    "ammType": "<string>",
    "liquidity": "<string>",
    "sponsorName": "<string>",
    "sponsorImage": "<string>",
    "startDate": "2023-11-07T05:31:56Z",
    "xAxisValue": "<string>",
    "yAxisValue": "<string>",
    "denominationToken": "<string>",
    "fee": "<string>",
    "image": "<string>",
    "icon": "<string>",
    "lowerBound": "<string>",
    "upperBound": "<string>",
    "description": "<string>",
    "outcomes": "<string>",
    "outcomePrices": "<string>",
    "volume": "<string>",
    "active": true,
    "marketType": "<string>",
    "formatType": "<string>",
    "lowerBoundDate": "<string>",
    "upperBoundDate": "<string>",
    "closed": true,
    "marketMakerAddress": "<string>",
    "createdBy": 123,
    "updatedBy": 123,
    "createdAt": "2023-11-07T05:31:56Z",
    "updatedAt": "2023-11-07T05:31:56Z",
    "closedTime": "<string>",
    "wideFormat": true,
    "new": true,
    "mailchimpTag": "<string>",
    "featured": true,
    "archived": true,
    "resolvedBy": "<string>",
    "restricted": true,
    "marketGroup": 123,
    "groupItemTitle": "<string>",
    "groupItemThreshold": "<string>",
    "questionID": "<string>",
    "umaEndDate": "<string>",
    "enableOrderBook": true,
    "orderPriceMinTickSize": 123,
    "orderMinSize": 123,
    "umaResolutionStatus": "<string>",
    "curationOrder": 123,
    "volumeNum": 123,
    "liquidityNum": 123,
    "endDateIso": "<string>",
    "startDateIso": "<string>",
    "umaEndDateIso": "<string>",
    "hasReviewedDates": true,
    "readyForCron": true,
    "commentsEnabled": true,
    "volume24hr": 123,
    "volume1wk": 123,
    "volume1mo": 123,
    "volume1yr": 123,
    "gameStartTime": "<string>",
    "secondsDelay": 123,
    "clobTokenIds": "<string>",
    "disqusThread": "<string>",
    "shortOutcomes": "<string>",
    "teamAID": "<string>",
    "teamBID": "<string>",
    "umaBond": "<string>",
    "umaReward": "<string>",
    "fpmmLive": true,
    "volume24hrAmm": 123,
    "volume1wkAmm": 123,
    "volume1moAmm": 123,
    "volume1yrAmm": 123,
    "volume24hrClob": 123,
    "volume1wkClob": 123,
    "volume1moClob": 123,
    "volume1yrClob": 123,
    "volumeAmm": 123,
    "volumeClob": 123,
    "liquidityAmm": 123,
    "liquidityClob": 123,
    "makerBaseFee": 123,
    "takerBaseFee": 123,
    "customLiveness": 123,
    "acceptingOrders": true,
    "notificationsEnabled": true,
    "score": 123,
    "imageOptimized": {
      "id": "<string>",
      "imageUrlSource": "<string>",
      "imageUrlOptimized": "<string>",
      "imageSizeKbSource": 123,
      "imageSizeKbOptimized": 123,
      "imageOptimizedComplete": true,
      "imageOptimizedLastUpdated": "<string>",
      "relID": 123,
      "field": "<string>",
      "relname": "<string>"
    },
    "iconOptimized": {
      "id": "<string>",
      "imageUrlSource": "<string>",
      "imageUrlOptimized": "<string>",
      "imageSizeKbSource": 123,
      "imageSizeKbOptimized": 123,
      "imageOptimizedComplete": true,
      "imageOptimizedLastUpdated": "<string>",
      "relID": 123,
      "field": "<string>",
      "relname": "<string>"
    },
    "events": [
      {
        "id": "<string>",
        "ticker": "<string>",
        "slug": "<string>",
        "title": "<string>",
        "subtitle": "<string>",
        "description": "<string>",
        "resolutionSource": "<string>",
        "startDate": "2023-11-07T05:31:56Z",
        "creationDate": "2023-11-07T05:31:56Z",
        "endDate": "2023-11-07T05:31:56Z",
        "image": "<string>",
        "icon": "<string>",
        "active": true,
        "closed": true,
        "archived": true,
        "new": true,
        "featured": true,
        "restricted": true,
        "liquidity": 123,
        "volume": 123,
        "openInterest": 123,
        "sortBy": "<string>",
        "category": "<string>",
        "subcategory": "<string>",
        "isTemplate": true,
        "templateVariables": "<string>",
        "published_at": "<string>",
        "createdBy": "<string>",
        "updatedBy": "<string>",
        "createdAt": "2023-11-07T05:31:56Z",
        "updatedAt": "2023-11-07T05:31:56Z",
        "commentsEnabled": true,
        "competitive": 123,
        "volume24hr": 123,
        "volume1wk": 123,
        "volume1mo": 123,
        "volume1yr": 123,
        "featuredImage": "<string>",
        "disqusThread": "<string>",
        "parentEvent": "<string>",
        "enableOrderBook": true,
        "liquidityAmm": 123,
        "liquidityClob": 123,
        "negRisk": true,
        "negRiskMarketID": "<string>",
        "negRiskFeeBips": 123,
        "commentCount": 123,
        "imageOptimized": {
          "id": "<string>",
          "imageUrlSource": "<string>",
          "imageUrlOptimized": "<string>",
          "imageSizeKbSource": 123,
          "imageSizeKbOptimized": 123,
          "imageOptimizedComplete": true,
          "imageOptimizedLastUpdated": "<string>",
          "relID": 123,
          "field": "<string>",
          "relname": "<string>"
        },
        "iconOptimized": {
          "id": "<string>",
          "imageUrlSource": "<string>",
          "imageUrlOptimized": "<string>",
          "imageSizeKbSource": 123,
          "imageSizeKbOptimized": 123,
          "imageOptimizedComplete": true,
          "imageOptimizedLastUpdated": "<string>",
          "relID": 123,
          "field": "<string>",
          "relname": "<string>"
        },
        "featuredImageOptimized": {
          "id": "<string>",
          "imageUrlSource": "<string>",
          "imageUrlOptimized": "<string>",
          "imageSizeKbSource": 123,
          "imageSizeKbOptimized": 123,
          "imageOptimizedComplete": true,
          "imageOptimizedLastUpdated": "<string>",
          "relID": 123,
          "field": "<string>",
          "relname": "<string>"
        },
        "subEvents": [
          "<string>"
        ],
        "markets": "<array>",
        "series": [
          {
            "id": "<string>",
            "ticker": "<string>",
            "slug": "<string>",
            "title": "<string>",
            "subtitle": "<string>",
            "seriesType": "<string>",
            "recurrence": "<string>",
            "description": "<string>",
            "image": "<string>",
            "icon": "<string>",
            "layout": "<string>",
            "active": true,
            "closed": true,
            "archived": true,
            "new": true,
            "featured": true,
            "restricted": true,
            "isTemplate": true,
            "templateVariables": true,
            "publishedAt": "<string>",
            "createdBy": "<string>",
            "updatedBy": "<string>",
            "createdAt": "2023-11-07T05:31:56Z",
            "updatedAt": "2023-11-07T05:31:56Z",
            "commentsEnabled": true,
            "competitive": "<string>",
            "volume24hr": 123,
            "volume": 123,
            "liquidity": 123,
            "startDate": "2023-11-07T05:31:56Z",
            "pythTokenID": "<string>",
            "cgAssetName": "<string>",
            "score": 123,
            "events": "<array>",
            "collections": [
              {
                "id": "<string>",
                "ticker": "<string>",
                "slug": "<string>",
                "title": "<string>",
                "subtitle": "<string>",
                "collectionType": "<string>",
                "description": "<string>",
                "tags": "<string>",
                "image": "<string>",
                "icon": "<string>",
                "headerImage": "<string>",
                "layout": "<string>",
                "active": true,
                "closed": true,
                "archived": true,
                "new": true,
                "featured": true,
                "restricted": true,
                "isTemplate": true,
                "templateVariables": "<string>",
                "publishedAt": "<string>",
                "createdBy": "<string>",
                "updatedBy": "<string>",
                "createdAt": "2023-11-07T05:31:56Z",
                "updatedAt": "2023-11-07T05:31:56Z",
                "commentsEnabled": true,
                "imageOptimized": {
                  "id": "<string>",
                  "imageUrlSource": "<string>",
                  "imageUrlOptimized": "<string>",
                  "imageSizeKbSource": 123,
                  "imageSizeKbOptimized": 123,
                  "imageOptimizedComplete": true,
                  "imageOptimizedLastUpdated": "<string>",
                  "relID": 123,
                  "field": "<string>",
                  "relname": "<string>"
                },
                "iconOptimized": {
                  "id": "<string>",
                  "imageUrlSource": "<string>",
                  "imageUrlOptimized": "<string>",
                  "imageSizeKbSource": 123,
                  "imageSizeKbOptimized": 123,
                  "imageOptimizedComplete": true,
                  "imageOptimizedLastUpdated": "<string>",
                  "relID": 123,
                  "field": "<string>",
                  "relname": "<string>"
                },
                "headerImageOptimized": {
                  "id": "<string>",
                  "imageUrlSource": "<string>",
                  "imageUrlOptimized": "<string>",
                  "imageSizeKbSource": 123,
                  "imageSizeKbOptimized": 123,
                  "imageOptimizedComplete": true,
                  "imageOptimizedLastUpdated": "<string>",
                  "relID": 123,
                  "field": "<string>",
                  "relname": "<string>"
                }
              }
            ],
            "categories": [
              {
                "id": "<string>",
                "label": "<string>",
                "parentCategory": "<string>",
                "slug": "<string>",
                "publishedAt": "<string>",
                "createdBy": "<string>",
                "updatedBy": "<string>",
                "createdAt": "2023-11-07T05:31:56Z",
                "updatedAt": "2023-11-07T05:31:56Z"
              }
            ],
            "tags": [
              {
                "id": "<string>",
                "label": "<string>",
                "slug": "<string>",
                "forceShow": true,
                "publishedAt": "<string>",
                "createdBy": 123,
                "updatedBy": 123,
                "createdAt": "2023-11-07T05:31:56Z",
                "updatedAt": "2023-11-07T05:31:56Z",
                "forceHide": true,
                "isCarousel": true
              }
            ],
            "commentCount": 123,
            "chats": [
              {
                "id": "<string>",
                "channelId": "<string>",
                "channelName": "<string>",
                "channelImage": "<string>",
                "live": true,
                "startTime": "2023-11-07T05:31:56Z",
                "endTime": "2023-11-07T05:31:56Z"
              }
            ]
          }
        ],
        "categories": [
          {
            "id": "<string>",
            "label": "<string>",
            "parentCategory": "<string>",
            "slug": "<string>",
            "publishedAt": "<string>",
            "createdBy": "<string>",
            "updatedBy": "<string>",
            "createdAt": "2023-11-07T05:31:56Z",
            "updatedAt": "2023-11-07T05:31:56Z"
          }
        ],
        "collections": [
          {
            "id": "<string>",
            "ticker": "<string>",
            "slug": "<string>",
            "title": "<string>",
            "subtitle": "<string>",
            "collectionType": "<string>",
            "description": "<string>",
            "tags": "<string>",
            "image": "<string>",
            "icon": "<string>",
            "headerImage": "<string>",
            "layout": "<string>",
            "active": true,
            "closed": true,
            "archived": true,
            "new": true,
            "featured": true,
            "restricted": true,
            "isTemplate": true,
            "templateVariables": "<string>",
            "publishedAt": "<string>",
            "createdBy": "<string>",
            "updatedBy": "<string>",
            "createdAt": "2023-11-07T05:31:56Z",
            "updatedAt": "2023-11-07T05:31:56Z",
            "commentsEnabled": true,
            "imageOptimized": {
              "id": "<string>",
              "imageUrlSource": "<string>",
              "imageUrlOptimized": "<string>",
              "imageSizeKbSource": 123,
              "imageSizeKbOptimized": 123,
              "imageOptimizedComplete": true,
              "imageOptimizedLastUpdated": "<string>",
              "relID": 123,
              "field": "<string>",
              "relname": "<string>"
            },
            "iconOptimized": {
              "id": "<string>",
              "imageUrlSource": "<string>",
              "imageUrlOptimized": "<string>",
              "imageSizeKbSource": 123,
              "imageSizeKbOptimized": 123,
              "imageOptimizedComplete": true,
              "imageOptimizedLastUpdated": "<string>",
              "relID": 123,
              "field": "<string>",
              "relname": "<string>"
            },
            "headerImageOptimized": {
              "id": "<string>",
              "imageUrlSource": "<string>",
              "imageUrlOptimized": "<string>",
              "imageSizeKbSource": 123,
              "imageSizeKbOptimized": 123,
              "imageOptimizedComplete": true,
              "imageOptimizedLastUpdated": "<string>",
              "relID": 123,
              "field": "<string>",
              "relname": "<string>"
            }
          }
        ],
        "tags": [
          {
            "id": "<string>",
            "label": "<string>",
            "slug": "<string>",
            "forceShow": true,
            "publishedAt": "<string>",
            "createdBy": 123,
            "updatedBy": 123,
            "createdAt": "2023-11-07T05:31:56Z",
            "updatedAt": "2023-11-07T05:31:56Z",
            "forceHide": true,
            "isCarousel": true
          }
        ],
        "cyom": true,
        "closedTime": "2023-11-07T05:31:56Z",
        "showAllOutcomes": true,
        "showMarketImages": true,
        "automaticallyResolved": true,
        "enableNegRisk": true,
        "automaticallyActive": true,
        "eventDate": "<string>",
        "startTime": "2023-11-07T05:31:56Z",
        "eventWeek": 123,
        "seriesSlug": "<string>",
        "score": "<string>",
        "elapsed": "<string>",
        "period": "<string>",
        "live": true,
        "ended": true,
        "finishedTimestamp": "2023-11-07T05:31:56Z",
        "gmpChartMode": "<string>",
        "eventCreators": [
          {
            "id": "<string>",
            "creatorName": "<string>",
            "creatorHandle": "<string>",
            "creatorUrl": "<string>",
            "creatorImage": "<string>",
            "createdAt": "2023-11-07T05:31:56Z",
            "updatedAt": "2023-11-07T05:31:56Z"
          }
        ],
        "tweetCount": 123,
        "chats": [
          {
            "id": "<string>",
            "channelId": "<string>",
            "channelName": "<string>",
            "channelImage": "<string>",
            "live": true,
            "startTime": "2023-11-07T05:31:56Z",
            "endTime": "2023-11-07T05:31:56Z"
          }
        ],
        "featuredOrder": 123,
        "estimateValue": true,
        "cantEstimate": true,
        "estimatedValue": "<string>",
        "templates": [
          {
            "id": "<string>",
            "eventTitle": "<string>",
            "eventSlug": "<string>",
            "eventImage": "<string>",
            "marketTitle": "<string>",
            "description": "<string>",
            "resolutionSource": "<string>",
            "negRisk": true,
            "sortBy": "<string>",
            "showMarketImages": true,
            "seriesSlug": "<string>",
            "outcomes": "<string>"
          }
        ],
        "spreadsMainLine": 123,
        "totalsMainLine": 123,
        "carouselMap": "<string>",
        "pendingDeployment": true,
        "deploying": true,
        "deployingTimestamp": "2023-11-07T05:31:56Z",
        "scheduledDeploymentTimestamp": "2023-11-07T05:31:56Z",
        "gameStatus": "<string>"
      }
    ],
    "categories": [
      {
        "id": "<string>",
        "label": "<string>",
        "parentCategory": "<string>",
        "slug": "<string>",
        "publishedAt": "<string>",
        "createdBy": "<string>",
        "updatedBy": "<string>",
        "createdAt": "2023-11-07T05:31:56Z",
        "updatedAt": "2023-11-07T05:31:56Z"
      }
    ],
    "tags": [
      {
        "id": "<string>",
        "label": "<string>",
        "slug": "<string>",
        "forceShow": true,
        "publishedAt": "<string>",
        "createdBy": 123,
        "updatedBy": 123,
        "createdAt": "2023-11-07T05:31:56Z",
        "updatedAt": "2023-11-07T05:31:56Z",
        "forceHide": true,
        "isCarousel": true
      }
    ],
    "creator": "<string>",
    "ready": true,
    "funded": true,
    "pastSlugs": "<string>",
    "readyTimestamp": "2023-11-07T05:31:56Z",
    "fundedTimestamp": "2023-11-07T05:31:56Z",
    "acceptingOrdersTimestamp": "2023-11-07T05:31:56Z",
    "competitive": 123,
    "rewardsMinSize": 123,
    "rewardsMaxSpread": 123,
    "spread": 123,
    "automaticallyResolved": true,
    "oneDayPriceChange": 123,
    "oneHourPriceChange": 123,
    "oneWeekPriceChange": 123,
    "oneMonthPriceChange": 123,
    "oneYearPriceChange": 123,
    "lastTradePrice": 123,
    "bestBid": 123,
    "bestAsk": 123,
    "automaticallyActive": true,
    "clearBookOnStart": true,
    "chartColor": "<string>",
    "seriesColor": "<string>",
    "showGmpSeries": true,
    "showGmpOutcome": true,
    "manualActivation": true,
    "negRiskOther": true,
    "gameId": "<string>",
    "groupItemRange": "<string>",
    "sportsMarketType": "<string>",
    "line": 123,
    "umaResolutionStatuses": "<string>",
    "pendingDeployment": true,
    "deploying": true,
    "deployingTimestamp": "2023-11-07T05:31:56Z",
    "scheduledDeploymentTimestamp": "2023-11-07T05:31:56Z",
    "rfqEnabled": true,
    "eventStartTime": "2023-11-07T05:31:56Z"
  }
]
```

## Query Parameters
​
limit
integer
Required range: x >= 0
​
offset
integer
Required range: x >= 0
​
order
string

Comma-separated list of fields to order by
​
ascending
boolean
​
id
integer[]
​
slug
string[]
​
clob_token_ids
string[]
​
condition_ids
string[]
​
market_maker_address
string[]
​
liquidity_num_min
number
​
liquidity_num_max
number
​
volume_num_min
number
​
volume_num_max
number
​
start_date_min
string<date-time>
​
start_date_max
string<date-time>
​
end_date_min
string<date-time>
​
end_date_max
string<date-time>
​
tag_id
integer
​
related_tags
boolean
​
cyom
boolean
​
uma_resolution_status
string
​
game_id
string
​
sports_market_types
string[]
​
rewards_min_size
number
​
question_ids
string[]
​
include_tag
boolean
​
closed
boolean

## Response
List of markets
​
id
string
​
question
string | null
​
conditionId
string
​
slug
string | null
​
twitterCardImage
string | null
​
resolutionSource
string | null
​
endDate
string<date-time> | null
​
category
string | null
​
ammType
string | null
​
liquidity
string | null
​
sponsorName
string | null
​
sponsorImage
string | null
​
startDate
string<date-time> | null
​
xAxisValue
string | null
​
yAxisValue
string | null
​
denominationToken
string | null
​
fee
string | null
​
image
string | null
​
icon
string | null
​
lowerBound
string | null
​
upperBound
string | null
​
description
string | null
​
outcomes
string | null
​
outcomePrices
string | null
​
volume
string | null
​
active
boolean | null
​
marketType
string | null
​
formatType
string | null
​
lowerBoundDate
string | null
​
upperBoundDate
string | null
​
closed
boolean | null
​
marketMakerAddress
string
​
createdBy
integer | null
​
updatedBy
integer | null
​
createdAt
string<date-time> | null
​
updatedAt
string<date-time> | null
​
closedTime
string | null
​
wideFormat
boolean | null
​
new
boolean | null
​
mailchimpTag
string | null
​
featured
boolean | null
​
archived
boolean | null
​
resolvedBy
string | null
​
restricted
boolean | null
​
marketGroup
integer | null
​
groupItemTitle
string | null
​
groupItemThreshold
string | null
​
questionID
string | null
​
umaEndDate
string | null
​
enableOrderBook
boolean | null
​
orderPriceMinTickSize
number | null
​
orderMinSize
number | null
​
umaResolutionStatus
string | null
​
curationOrder
integer | null
​
volumeNum
number | null
​
liquidityNum
number | null
​
endDateIso
string | null
​
startDateIso
string | null
​
umaEndDateIso
string | null
​
hasReviewedDates
boolean | null
​
readyForCron
boolean | null
​
commentsEnabled
boolean | null
​
volume24hr
number | null
​
volume1wk
number | null
​
volume1mo
number | null
​
volume1yr
number | null
​
gameStartTime
string | null
​
secondsDelay
integer | null
​
clobTokenIds
string | null
​
disqusThread
string | null
​
shortOutcomes
string | null
​
teamAID
string | null
​
teamBID
string | null
​
umaBond
string | null
​
umaReward
string | null
​
fpmmLive
boolean | null
​
volume24hrAmm
number | null
​
volume1wkAmm
number | null
​
volume1moAmm
number | null
​
volume1yrAmm
number | null
​
volume24hrClob
number | null
​
volume1wkClob
number | null
​
volume1moClob
number | null
​
volume1yrClob
number | null
​
volumeAmm
number | null
​
volumeClob
number | null
​
liquidityAmm
number | null
​
liquidityClob
number | null
​
makerBaseFee
integer | null
​
takerBaseFee
integer | null
​
customLiveness
integer | null
​
acceptingOrders
boolean | null
​
notificationsEnabled
boolean | null
​
score
integer | null
​
imageOptimized
object

iconOptimized
object

events
object[]
Child attributes:
    events.id
    string
    ​
    events.ticker
    string | null
    ​
    events.slug
    string | null
    ​
    events.title
    string | null
    ​
    events.subtitle
    string | null
    ​
    events.description
    string | null
    ​
    events.resolutionSource
    string | null
    ​
    events.startDate
    string<date-time> | null
    ​
    events.creationDate
    string<date-time> | null
    ​
    events.endDate
    string<date-time> | null
    ​
    events.image
    string | null
    ​
    events.icon
    string | null
    ​
    events.active
    boolean | null
    ​
    events.closed
    boolean | null
    ​
    events.archived
    boolean | null
    ​
    events.new
    boolean | null
    ​
    events.featured
    boolean | null
    ​
    events.restricted
    boolean | null
    ​
    events.liquidity
    number | null
    ​
    events.volume
    number | null
    ​
    events.openInterest
    number | null
    ​
    events.sortBy
    string | null
    ​
    events.category
    string | null
    ​
    events.subcategory
    string | null
    ​
    events.isTemplate
    boolean | null
    ​
    events.templateVariables
    string | null
    ​
    events.published_at
    string | null
    ​
    events.createdBy
    string | null
    ​
    events.updatedBy
    string | null
    ​
    events.createdAt
    string<date-time> | null
    ​
    events.updatedAt
    string<date-time> | null
    ​
    events.commentsEnabled
    boolean | null
    ​
    events.competitive
    number | null
    ​
    events.volume24hr
    number | null
    ​
    events.volume1wk
    number | null
    ​
    events.volume1mo
    number | null
    ​
    events.volume1yr
    number | null
    ​
    events.featuredImage
    string | null
    ​
    events.disqusThread
    string | null
    ​
    events.parentEvent
    string | null
    ​
    events.enableOrderBook
    boolean | null
    ​
    events.liquidityAmm
    number | null
    ​
    events.liquidityClob
    number | null
    ​
    events.negRisk
    boolean | null
    ​
    events.negRiskMarketID
    string | null
    ​
    events.negRiskFeeBips
    integer | null
    ​
    events.commentCount
    integer | null
    ​
    events.imageOptimized
    object

    events.iconOptimized
    object

    events.featuredImageOptimized
    object

    events.subEvents
    string[] | null

    events.markets
    array
    ​
    events.series
    object[]
    Child attributes:
        events.series.id
        string
        ​
        events.series.ticker
        string | null
        ​
        events.series.slug
        string | null
        ​
        events.series.title
        string | null
        ​
        events.series.subtitle
        string | null
        ​
        events.series.seriesType
        string | null
        ​
        events.series.recurrence
        string | null
        ​
        events.series.description
        string | null
        ​
        events.series.image
        string | null
        ​
        events.series.icon
        string | null
        ​
        events.series.layout
        string | null
        ​
        events.series.active
        boolean | null
        ​
        events.series.closed
        boolean | null
        ​
        events.series.archived
        boolean | null
        ​
        events.series.new
        boolean | null
        ​
        events.series.featured
        boolean | null
        ​
        events.series.restricted
        boolean | null
        ​
        events.series.isTemplate
        boolean | null
        ​
        events.series.templateVariables
        boolean | null
        ​
        events.series.publishedAt
        string | null
        ​
        events.series.createdBy
        string | null
        ​
        events.series.updatedBy
        string | null
        ​
        events.series.createdAt
        string<date-time> | null
        ​
        events.series.updatedAt
        string<date-time> | null
        ​
        events.series.commentsEnabled
        boolean | null
        ​
        events.series.competitive
        string | null
        ​
        events.series.volume24hr
        number | null
        ​
        events.series.volume
        number | null
        ​
        events.series.liquidity
        number | null
        ​
        events.series.startDate
        string<date-time> | null
        ​
        events.series.pythTokenID
        string | null
        ​
        events.series.cgAssetName
        string | null
        ​
        events.series.score
        integer | null
        ​
        events.series.events
        array
        ​
        events.series.collections
        object[]
    
    events.categories
    object[]
    Child attributes:
        events.categories.id
        string
        ​
        events.categories.label
        string | null
        ​
        events.categories.parentCategory
        string | null
        ​
        events.categories.slug
        string | null
        ​
        events.categories.publishedAt
        string | null
        ​
        events.categories.createdBy
        string | null
        ​
        events.categories.updatedBy
        string | null
        ​
        events.categories.createdAt
        string<date-time> | null
        ​
        events.categories.updatedAt
        string<date-time> | null

    events.collections
    object[]
    Child attributes:
        events.collections.id
        string
        ​
        events.collections.ticker
        string | null
        ​
        events.collections.slug
        string | null
        ​
        events.collections.title
        string | null
        ​
        events.collections.subtitle
        string | null
        ​
        events.collections.collectionType
        string | null
        ​
        events.collections.description
        string | null
        ​
        events.collections.tags
        string | null
        ​
        events.collections.image
        string | null
        ​
        events.collections.icon
        string | null
        ​
        events.collections.headerImage
        string | null
        ​
        events.collections.layout
        string | null
        ​
        events.collections.active
        boolean | null
        ​
        events.collections.closed
        boolean | null
        ​
        events.collections.archived
        boolean | null
        ​
        events.collections.new
        boolean | null
        ​
        events.collections.featured
        boolean | null
        ​
        events.collections.restricted
        boolean | null
        ​
        events.collections.isTemplate
        boolean | null
        ​
        events.collections.templateVariables
        string | null
        ​
        events.collections.publishedAt
        string | null
        ​
        events.collections.createdBy
        string | null
        ​
        events.collections.updatedBy
        string | null
        ​
        events.collections.createdAt
        string<date-time> | null
        ​
        events.collections.updatedAt
        string<date-time> | null
        ​
        events.collections.commentsEnabled
        boolean | null
        ​
        events.collections.imageOptimized
        object

    events.tags
    object[]

    events.cyom
    boolean | null
    ​
    events.closedTime
    string<date-time> | null
    ​
    events.showAllOutcomes
    boolean | null
    ​
    events.showMarketImages
    boolean | null
    ​
    events.automaticallyResolved
    boolean | null
    ​
    events.enableNegRisk
    boolean | null
    ​
    events.automaticallyActive
    boolean | null
    ​
    events.eventDate
    string | null
    ​
    events.startTime
    string<date-time> | null
    ​
    events.eventWeek
    integer | null
    ​
    events.seriesSlug
    string | null
    ​
    events.score
    string | null
    ​
    events.elapsed
    string | null
    ​
    events.period
    string | null
    ​
    events.live
    boolean | null
    ​
    events.ended
    boolean | null
    ​
    events.finishedTimestamp
    string<date-time> | null
    ​
    events.gmpChartMode
    string | null
    ​
    events.eventCreators
    object[]

    events.tweetCount
    integer | null

    events.chats
    object[]

    events.featuredOrder
    integer | null
    ​
    events.estimateValue
    boolean | null
    ​
    events.cantEstimate
    boolean | null
    ​
    events.estimatedValue
    string | null
    ​
    events.templates
    object[]

    events.spreadsMainLine
    number | null
    ​
    events.totalsMainLine
    number | null
    ​
    events.carouselMap
    string | null
    ​
    events.pendingDeployment
    boolean | null
    ​
    events.deploying
    boolean | null
    ​
    events.deployingTimestamp
    string<date-time> | null
    ​
    events.scheduledDeploymentTimestamp
    string<date-time> | null
    ​
    events.gameStatus
    string | null

categories
object[]
Child attributes:
    categories.id
    string
    ​
    categories.label
    string | null
    ​
    categories.parentCategory
    string | null
    ​
    categories.slug
    string | null
    ​
    categories.publishedAt
    string | null
    ​
    categories.createdBy
    string | null
    ​
    categories.updatedBy
    string | null
    ​
    categories.createdAt
    string<date-time> | null
    ​
    categories.updatedAt
    string<date-time> | null

tags
object[]

creator
string | null
​
ready
boolean | null
​
funded
boolean | null
​
pastSlugs
string | null
​
readyTimestamp
string<date-time> | null
​
fundedTimestamp
string<date-time> | null
​
acceptingOrdersTimestamp
string<date-time> | null
​
competitive
number | null
​
rewardsMinSize
number | null
​
rewardsMaxSpread
number | null
​
spread
number | null
​
automaticallyResolved
boolean | null
​
oneDayPriceChange
number | null
​
oneHourPriceChange
number | null
​
oneWeekPriceChange
number | null
​
oneMonthPriceChange
number | null
​
oneYearPriceChange
number | null
​
lastTradePrice
number | null
​
bestBid
number | null
​
bestAsk
number | null
​
automaticallyActive
boolean | null
​
clearBookOnStart
boolean | null
​
chartColor
string | null
​
seriesColor
string | null
​
showGmpSeries
boolean | null
​
showGmpOutcome
boolean | null
​
manualActivation
boolean | null
​
negRiskOther
boolean | null
​
gameId
string | null
​
groupItemRange
string | null
​
sportsMarketType
string | null
​
line
number | null
​
umaResolutionStatuses
string | null
​
pendingDeployment
boolean | null
​
deploying
boolean | null
​
deployingTimestamp
string<date-time> | null
​
scheduledDeploymentTimestamp
string<date-time> | null
​
rfqEnabled
boolean | null
​
eventStartTime
string<date-time> | null