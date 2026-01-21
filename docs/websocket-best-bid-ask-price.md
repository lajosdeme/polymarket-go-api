# best_bid_ask Message
Emitted When:

    The best bid and ask prices for a market change.

(This message is behind the custom_feature_enabled flag)
​
Structure
Name	Type	Description
event_type	string	”best_bid_ask”
market	string	condition ID of market
asset_id	string	asset ID (token ID)
best_bid	string	current best bid price
best_ask	string	current best ask price
spread	string	spread between best bid and ask
timestamp	string	unix timestamp in milliseconds

```
{
  "event_type": "best_bid_ask",
  "market": "0x0005c0d312de0be897668695bae9f32b624b4a1ae8b140c49f08447fcc74f442",
  "asset_id": "85354956062430465315924116860125388538595433819574542752031640332592237464430",
  "best_bid": "0.73",
  "best_ask": "0.77",
  "spread": "0.04",
  "timestamp": "1766789469958"
}
```