# List tags
```
curl --request GET \
  --url https://gamma-api.polymarket.com/tags
  ```

## Response:
```
[
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
include_template
boolean
​
is_carousel
boolean

## Response
List of tags
​
id
string
​
label
string | null
​
slug
string | null
​
forceShow
boolean | null
​
publishedAt
string | null
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
forceHide
boolean | null
​
isCarousel
boolean | null

# Get tag by id
```
curl --request GET \
  --url https://gamma-api.polymarket.com/tags/{id}
  ```

```
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
```

Path Parameters
​
id
integer
required


Query Parameters
​
include_template
boolean

## Response
Tag
​
id
string
​
label
string | null
​
slug
string | null
​
forceShow
boolean | null
​
publishedAt
string | null
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
forceHide
boolean | null
​
isCarousel
boolean | null

# Get tag by slug
```
curl --request GET \
  --url https://gamma-api.polymarket.com/tags/slug/{slug}
  ```

```
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
```

Path Parameters
​
slug
string
required


Query Parameters
​
include_template
boolean

## Response
Tag
​
id
string
​
label
string | null
​
slug
string | null
​
forceShow
boolean | null
​
publishedAt
string | null
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
forceHide
boolean | null
​
isCarousel
boolean | null

# Get related tags (relationships) by tag id
```
curl --request GET \
  --url https://gamma-api.polymarket.com/tags/{id}/related-tags
  ```

```
[
  {
    "id": "<string>",
    "tagID": 123,
    "relatedTagID": 123,
    "rank": 123
  }
]
```

Path Parameters
​
id
integer
required


Query Parameters
​
omit_empty
boolean
​
status
enum<string>
Available options: active, 
closed, 
all 

## Response 
Related tag relationships
​
id
string
​
tagID
integer | null
​
relatedTagID
integer | null
​
rank
integer | null

# Get related tags (relationships) by tag slug
```
curl --request GET \
  --url https://gamma-api.polymarket.com/tags/slug/{slug}/related-tags
```

```
[
  {
    "id": "<string>",
    "tagID": 123,
    "relatedTagID": 123,
    "rank": 123
  }
]
```

Path Parameters
​
slug
string
required


Query Parameters
​
omit_empty
boolean
​
status
enum<string>
Available options: active, 
closed, 
all 

## Response
Related tag relationships
​
id
string
​
tagID
integer | null
​
relatedTagID
integer | null
​
rank
integer | null

# Get tags related to a tag id
```
curl --request GET \
  --url https://gamma-api.polymarket.com/tags/{id}/related-tags/tags
```

```
[
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
]
```

Path Parameters
​
id
integer
required
Query Parameters
​
omit_empty
boolean
​
status
enum<string>
Available options: active, 
closed, 
all 

## Response
Related tags
​
id
string
​
label
string | null
​
slug
string | null
​
forceShow
boolean | null
​
publishedAt
string | null
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
forceHide
boolean | null
​
isCarousel
boolean | null

# Get tags related to a tag slug
```
curl --request GET \
  --url https://gamma-api.polymarket.com/tags/slug/{slug}/related-tags/tags
```

```
[
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
]
```

Path Parameters
​
slug
string
required
Query Parameters
​
omit_empty
boolean
​
status
enum<string>
Available options: active, 
closed, 
all 

## Response 
Related tags
​
id
string
​
label
string | null
​
slug
string | null
​
forceShow
boolean | null
​
publishedAt
string | null
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
forceHide
boolean | null
​
isCarousel
boolean | null

