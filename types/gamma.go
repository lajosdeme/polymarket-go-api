package types

import "time"

// ImageOptimized represents optimized image metadata
type ImageOptimized struct {
	ID                        string `json:"id"`
	ImageURLSource            string `json:"imageUrlSource"`
	ImageURLOptimized         string `json:"imageUrlOptimized"`
	ImageSizeKbSource         *int   `json:"imageSizeKbSource"`
	ImageSizeKbOptimized      *int   `json:"imageSizeKbOptimized"`
	ImageOptimizedComplete    *bool  `json:"imageOptimizedComplete"`
	ImageOptimizedLastUpdated string `json:"imageOptimizedLastUpdated"`
	RelID                     *int   `json:"relID"`
	Field                     string `json:"field"`
	Relname                   string `json:"relname"`
}

// Category represents a market/event category
type Category struct {
	ID             string     `json:"id"`
	Label          *string    `json:"label"`
	ParentCategory *string    `json:"parentCategory"`
	Slug           *string    `json:"slug"`
	PublishedAt    *string    `json:"publishedAt"`
	CreatedBy      *string    `json:"createdBy"`
	UpdatedBy      *string    `json:"updatedBy"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

// Tag represents a metadata tag
type Tag struct {
	ID          string     `json:"id"`
	Label       *string    `json:"label"`
	Slug        *string    `json:"slug"`
	ForceShow   *bool      `json:"forceShow"`
	PublishedAt *string    `json:"publishedAt"`
	CreatedBy   *float64   `json:"createdBy"`
	UpdatedBy   *float64   `json:"updatedBy"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	ForceHide   *bool      `json:"forceHide"`
	IsCarousel  *bool      `json:"isCarousel"`
}

// TagWithCount represents a tag with event count
type TagWithCount struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	Slug       string `json:"slug"`
	EventCount *int   `json:"event_count"`
}

// Collection represents a collection of markets/events
type Collection struct {
	ID                   string          `json:"id"`
	Ticker               *string         `json:"ticker"`
	Slug                 *string         `json:"slug"`
	Title                *string         `json:"title"`
	Subtitle             *string         `json:"subtitle"`
	CollectionType       *string         `json:"collectionType"`
	Description          *string         `json:"description"`
	Tags                 *string         `json:"tags"`
	Image                *string         `json:"image"`
	Icon                 *string         `json:"icon"`
	HeaderImage          *string         `json:"headerImage"`
	Layout               *string         `json:"layout"`
	Active               *bool           `json:"active"`
	Closed               *bool           `json:"closed"`
	Archived             *bool           `json:"archived"`
	New                  *bool           `json:"new"`
	Featured             *bool           `json:"featured"`
	Restricted           *bool           `json:"restricted"`
	IsTemplate           *bool           `json:"isTemplate"`
	TemplateVariables    *string         `json:"templateVariables"`
	PublishedAt          *string         `json:"publishedAt"`
	CreatedBy            *string         `json:"createdBy"`
	UpdatedBy            *string         `json:"updatedBy"`
	CreatedAt            *time.Time      `json:"createdAt"`
	UpdatedAt            *time.Time      `json:"updatedAt"`
	CommentsEnabled      *bool           `json:"commentsEnabled"`
	ImageOptimized       *ImageOptimized `json:"imageOptimized"`
	IconOptimized        *ImageOptimized `json:"iconOptimized"`
	HeaderImageOptimized *ImageOptimized `json:"headerImageOptimized"`
}

// Series represents a series of related events
type Series struct {
	ID                string       `json:"id"`
	Ticker            *string      `json:"ticker"`
	Slug              *string      `json:"slug"`
	Title             *string      `json:"title"`
	Subtitle          *string      `json:"subtitle"`
	SeriesType        *string      `json:"seriesType"`
	Recurrence        *string      `json:"recurrence"`
	Description       *string      `json:"description"`
	Image             *string      `json:"image"`
	Icon              *string      `json:"icon"`
	Layout            *string      `json:"layout"`
	Active            *bool        `json:"active"`
	Closed            *bool        `json:"closed"`
	Archived          *bool        `json:"archived"`
	New               *bool        `json:"new"`
	Featured          *bool        `json:"featured"`
	Restricted        *bool        `json:"restricted"`
	IsTemplate        *bool        `json:"isTemplate"`
	TemplateVariables *bool        `json:"templateVariables"`
	PublishedAt       *string      `json:"publishedAt"`
	CreatedBy         *string      `json:"createdBy"`
	UpdatedBy         *string      `json:"updatedBy"`
	CreatedAt         *time.Time   `json:"createdAt"`
	UpdatedAt         *time.Time   `json:"updatedAt"`
	CommentsEnabled   *bool        `json:"commentsEnabled"`
	Competitive       *string      `json:"competitive"`
	Volume24hr        *float64     `json:"volume24hr"`
	Volume            *float64     `json:"volume"`
	Liquidity         *float64     `json:"liquidity"`
	StartDate         *time.Time   `json:"startDate"`
	PythTokenID       *string      `json:"pythTokenID"`
	CgAssetName       *string      `json:"cgAssetName"`
	Score             *float64     `json:"score"`
	Events            []any        `json:"events"`
	Collections       []Collection `json:"collections"`
	Categories        []Category   `json:"categories"`
	Tags              []Tag        `json:"tags"`
	CommentCount      *float64     `json:"commentCount"`
	Chats             []Chat       `json:"chats"`
}

// EventCreator represents the creator of an event
type EventCreator struct {
	ID            string     `json:"id"`
	CreatorName   *string    `json:"creatorName"`
	CreatorHandle *string    `json:"creatorHandle"`
	CreatorURL    *string    `json:"creatorUrl"`
	CreatorImage  *string    `json:"creatorImage"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}

// Chat represents a chat associated with an event
type Chat struct {
	ID           string     `json:"id"`
	ChannelID    *string    `json:"channelId"`
	ChannelName  *string    `json:"channelName"`
	ChannelImage *string    `json:"channelImage"`
	Live         *bool      `json:"live"`
	StartTime    *time.Time `json:"startTime"`
	EndTime      *time.Time `json:"endTime"`
}

// Template represents an event template
type Template struct {
	ID               string `json:"id"`
	EventTitle       string `json:"eventTitle"`
	EventSlug        string `json:"eventSlug"`
	EventImage       string `json:"eventImage"`
	MarketTitle      string `json:"marketTitle"`
	Description      string `json:"description"`
	ResolutionSource string `json:"resolutionSource"`
	NegRisk          bool   `json:"negRisk"`
	SortBy           string `json:"sortBy"`
	ShowMarketImages bool   `json:"showMarketImages"`
	SeriesSlug       string `json:"seriesSlug"`
	Outcomes         string `json:"outcomes"`
}

// GammaEvent represents an event in the Gamma API
type GammaEvent struct {
	ID                           string          `json:"id"`
	Ticker                       *string         `json:"ticker"`
	Slug                         *string         `json:"slug"`
	Title                        *string         `json:"title"`
	Subtitle                     *string         `json:"subtitle"`
	Description                  *string         `json:"description"`
	ResolutionSource             *string         `json:"resolutionSource"`
	StartDate                    *time.Time      `json:"startDate"`
	CreationDate                 *time.Time      `json:"creationDate"`
	EndDate                      *time.Time      `json:"endDate"`
	Image                        *string         `json:"image"`
	Icon                         *string         `json:"icon"`
	Active                       *bool           `json:"active"`
	Closed                       *bool           `json:"closed"`
	Archived                     *bool           `json:"archived"`
	New                          *bool           `json:"new"`
	Featured                     *bool           `json:"featured"`
	Restricted                   *bool           `json:"restricted"`
	Liquidity                    *float64        `json:"liquidity"`
	Volume                       *float64        `json:"volume"`
	OpenInterest                 *float64        `json:"openInterest"`
	SortBy                       *string         `json:"sortBy"`
	Category                     *string         `json:"category"`
	Subcategory                  *string         `json:"subcategory"`
	IsTemplate                   *bool           `json:"isTemplate"`
	TemplateVariables            *string         `json:"templateVariables"`
	PublishedAt                  *string         `json:"published_at"`
	CreatedBy                    *string         `json:"createdBy"`
	UpdatedBy                    *string         `json:"updatedBy"`
	CreatedAt                    *time.Time      `json:"createdAt"`
	UpdatedAt                    *time.Time      `json:"updatedAt"`
	CommentsEnabled              *bool           `json:"commentsEnabled"`
	Competitive                  *float64        `json:"competitive"`
	Volume24hr                   *float64        `json:"volume24hr"`
	Volume1wk                    *float64        `json:"volume1wk"`
	Volume1mo                    *float64        `json:"volume1mo"`
	Volume1yr                    *float64        `json:"volume1yr"`
	FeaturedImage                *string         `json:"featuredImage"`
	DisqusThread                 *string         `json:"disqusThread"`
	ParentEvent                  *string         `json:"parentEvent"`
	EnableOrderBook              *bool           `json:"enableOrderBook"`
	LiquidityAmm                 *int            `json:"liquidityAmm"`
	LiquidityClob                *float64        `json:"liquidityClob"`
	NegRisk                      *bool           `json:"negRisk"`
	NegRiskMarketID              *string         `json:"negRiskMarketID"`
	NegRiskFeeBips               *int            `json:"negRiskFeeBips"`
	CommentCount                 *int            `json:"commentCount"`
	ImageOptimized               *ImageOptimized `json:"imageOptimized"`
	IconOptimized                *ImageOptimized `json:"iconOptimized"`
	FeaturedImageOptimized       *ImageOptimized `json:"featuredImageOptimized"`
	SubEvents                    []string        `json:"subEvents"`
	Markets                      []GammaMarket   `json:"markets"`
	Series                       []Series        `json:"series"`
	Categories                   []Category      `json:"categories"`
	Collections                  []Collection    `json:"collections"`
	Tags                         []Tag           `json:"tags"`
	Cyom                         *bool           `json:"cyom"`
	ClosedTime                   *time.Time      `json:"closedTime"`
	ShowAllOutcomes              *bool           `json:"showAllOutcomes"`
	ShowMarketImages             *bool           `json:"showMarketImages"`
	AutomaticallyResolved        *bool           `json:"automaticallyResolved"`
	EnableNegRisk                *bool           `json:"enableNegRisk"`
	AutomaticallyActive          *bool           `json:"automaticallyActive"`
	EventDate                    *string         `json:"eventDate"`
	StartTime                    *time.Time      `json:"startTime"`
	EventWeek                    *int            `json:"eventWeek"`
	SeriesSlug                   *string         `json:"seriesSlug"`
	Score                        *string         `json:"score"`
	Elapsed                      *string         `json:"elapsed"`
	Period                       *string         `json:"period"`
	Live                         *bool           `json:"live"`
	Ended                        *bool           `json:"ended"`
	FinishedTimestamp            *time.Time      `json:"finishedTimestamp"`
	GmpChartMode                 *string         `json:"gmpChartMode"`
	EventCreators                []EventCreator  `json:"eventCreators"`
	TweetCount                   *float64        `json:"tweetCount"`
	Chats                        []Chat          `json:"chats"`
	FeaturedOrder                *float64        `json:"featuredOrder"`
	EstimateValue                *bool           `json:"estimateValue"`
	CantEstimate                 *bool           `json:"cantEstimate"`
	EstimatedValue               *string         `json:"estimatedValue"`
	Templates                    []Template      `json:"templates"`
	SpreadsMainLine              *float64        `json:"spreadsMainLine"`
	TotalsMainLine               *float64        `json:"totalsMainLine"`
	CarouselMap                  *string         `json:"carouselMap"`
	PendingDeployment            *bool           `json:"pendingDeployment"`
	Deploying                    *bool           `json:"deploying"`
	DeployingTimestamp           *time.Time      `json:"deployingTimestamp"`
	ScheduledDeploymentTimestamp *time.Time      `json:"scheduledDeploymentTimestamp"`
	GameStatus                   *string         `json:"gameStatus"`
}

// GammaMarket represents a market in the Gamma API
type GammaMarket struct {
	ID                           string          `json:"id"`
	Question                     *string         `json:"question"`
	ConditionID                  *string         `json:"conditionId"`
	Slug                         *string         `json:"slug"`
	TwitterCardImage             *string         `json:"twitterCardImage"`
	ResolutionSource             *string         `json:"resolutionSource"`
	EndDate                      *time.Time      `json:"endDate"`
	Category                     *string         `json:"category"`
	AmmType                      *string         `json:"ammType"`
	Liquidity                    *string         `json:"liquidity"`
	SponsorName                  *string         `json:"sponsorName"`
	SponsorImage                 *string         `json:"sponsorImage"`
	StartDate                    *time.Time      `json:"startDate"`
	XAxisValue                   *string         `json:"xAxisValue"`
	YAxisValue                   *string         `json:"yAxisValue"`
	DenominationToken            *string         `json:"denominationToken"`
	Fee                          *string         `json:"fee"`
	Image                        *string         `json:"image"`
	Icon                         *string         `json:"icon"`
	LowerBound                   *string         `json:"lowerBound"`
	UpperBound                   *string         `json:"upperBound"`
	Description                  *string         `json:"description"`
	Outcomes                     *string         `json:"outcomes"`
	OutcomePrices                *string         `json:"outcomePrices"`
	Volume                       *string         `json:"volume"`
	Active                       *bool           `json:"active"`
	MarketType                   *string         `json:"marketType"`
	FormatType                   *string         `json:"formatType"`
	LowerBoundDate               *string         `json:"lowerBoundDate"`
	UpperBoundDate               *string         `json:"upperBoundDate"`
	Closed                       *bool           `json:"closed"`
	MarketMakerAddress           *string         `json:"marketMakerAddress"`
	CreatedBy                    *float64        `json:"createdBy"`
	UpdatedBy                    *float64        `json:"updatedBy"`
	CreatedAt                    *time.Time      `json:"createdAt"`
	UpdatedAt                    *time.Time      `json:"updatedAt"`
	ClosedTime                   *string         `json:"closedTime"`
	WideFormat                   *bool           `json:"wideFormat"`
	New                          *bool           `json:"new"`
	MailchimpTag                 *string         `json:"mailchimpTag"`
	Featured                     *bool           `json:"featured"`
	Archived                     *bool           `json:"archived"`
	ResolvedBy                   *string         `json:"resolvedBy"`
	Restricted                   *bool           `json:"restricted"`
	MarketGroup                  *float64        `json:"marketGroup"`
	GroupItemTitle               *string         `json:"groupItemTitle"`
	GroupItemThreshold           *string         `json:"groupItemThreshold"`
	QuestionID                   *string         `json:"questionID"`
	UmaEndDate                   *string         `json:"umaEndDate"`
	EnableOrderBook              *bool           `json:"enableOrderBook"`
	OrderPriceMinTickSize        *float64        `json:"orderPriceMinTickSize"`
	OrderMinSize                 *float64        `json:"orderMinSize"`
	UmaResolutionStatus          *string         `json:"umaResolutionStatus"`
	CurationOrder                *float64        `json:"curationOrder"`
	VolumeNum                    *float64        `json:"volumeNum"`
	LiquidityNum                 *float64        `json:"liquidityNum"`
	EndDateIso                   *string         `json:"endDateIso"`
	StartDateIso                 *string         `json:"startDateIso"`
	UmaEndDateIso                *string         `json:"umaEndDateIso"`
	HasReviewedDates             *bool           `json:"hasReviewedDates"`
	ReadyForCron                 *bool           `json:"readyForCron"`
	CommentsEnabled              *bool           `json:"commentsEnabled"`
	Volume24hr                   *float64        `json:"volume24hr"`
	Volume1wk                    *float64        `json:"volume1wk"`
	Volume1mo                    *float64        `json:"volume1mo"`
	Volume1yr                    *float64        `json:"volume1yr"`
	GameStartTime                *string         `json:"gameStartTime"`
	SecondsDelay                 *float64        `json:"secondsDelay"`
	ClobTokenIds                 *string         `json:"clobTokenIds"`
	DisqusThread                 *string         `json:"disqusThread"`
	ShortOutcomes                *string         `json:"shortOutcomes"`
	TeamAID                      *string         `json:"teamAID"`
	TeamBID                      *string         `json:"teamBID"`
	UmaBond                      *string         `json:"umaBond"`
	UmaReward                    *string         `json:"umaReward"`
	FpmmLive                     *bool           `json:"fpmmLive"`
	Volume24hrAmm                *float64        `json:"volume24hrAmm"`
	Volume1wkAmm                 *float64        `json:"volume1wkAmm"`
	Volume1moAmm                 *float64        `json:"volume1moAmm"`
	Volume1yrAmm                 *float64        `json:"volume1yrAmm"`
	Volume24hrClob               *float64        `json:"volume24hrClob"`
	Volume1wkClob                *float64        `json:"volume1wkClob"`
	Volume1moClob                *float64        `json:"volume1moClob"`
	Volume1yrClob                *float64        `json:"volume1yrClob"`
	VolumeAmm                    *float64        `json:"volumeAmm"`
	VolumeClob                   *float64        `json:"volumeClob"`
	LiquidityAmm                 *float64        `json:"liquidityAmm"`
	LiquidityClob                *float64        `json:"liquidityClob"`
	MakerBaseFee                 *float64        `json:"makerBaseFee"`
	TakerBaseFee                 *float64        `json:"takerBaseFee"`
	CustomLiveness               *float64        `json:"customLiveness"`
	AcceptingOrders              *bool           `json:"acceptingOrders"`
	NotificationsEnabled         *bool           `json:"notificationsEnabled"`
	Score                        *int            `json:"score"`
	ImageOptimized               *ImageOptimized `json:"imageOptimized"`
	IconOptimized                *ImageOptimized `json:"iconOptimized"`
	Events                       []GammaEvent    `json:"events"`
	Categories                   []Category      `json:"categories"`
	Tags                         []Tag           `json:"tags"`
	Creator                      *string         `json:"creator"`
	Ready                        *bool           `json:"ready"`
	Funded                       *bool           `json:"funded"`
	PastSlugs                    *string         `json:"pastSlugs"`
	ReadyTimestamp               *time.Time      `json:"readyTimestamp"`
	FundedTimestamp              *time.Time      `json:"fundedTimestamp"`
	AcceptingOrdersTimestamp     *time.Time      `json:"acceptingOrdersTimestamp"`
	Competitive                  *float64        `json:"competitive"`
	RewardsMinSize               *float64        `json:"rewardsMinSize"`
	RewardsMaxSpread             *float64        `json:"rewardsMaxSpread"`
	Spread                       *float64        `json:"spread"`
	AutomaticallyResolved        *bool           `json:"automaticallyResolved"`
	OneDayPriceChange            *float64        `json:"oneDayPriceChange"`
	OneHourPriceChange           *float64        `json:"oneHourPriceChange"`
	OneWeekPriceChange           *float64        `json:"oneWeekPriceChange"`
	OneMonthPriceChange          *float64        `json:"oneMonthPriceChange"`
	OneYearPriceChange           *float64        `json:"oneYearPriceChange"`
	LastTradePrice               *float64        `json:"lastTradePrice"`
	BestBid                      *float64        `json:"bestBid"`
	BestAsk                      *float64        `json:"bestAsk"`
	AutomaticallyActive          *bool           `json:"automaticallyActive"`
	ClearBookOnStart             *bool           `json:"clearBookOnStart"`
	ChartColor                   *string         `json:"chartColor"`
	SeriesColor                  *string         `json:"seriesColor"`
	ShowGmpSeries                *bool           `json:"showGmpSeries"`
	ShowGmpOutcome               *bool           `json:"showGmpOutcome"`
	ManualActivation             *bool           `json:"manualActivation"`
	NegRiskOther                 *bool           `json:"negRiskOther"`
	GameID                       *string         `json:"gameId"`
	GroupItemRange               *string         `json:"groupItemRange"`
	SportsMarketType             *string         `json:"sportsMarketType"`
	Line                         *float64        `json:"line"`
	UmaResolutionStatuses        *string         `json:"umaResolutionStatuses"`
	PendingDeployment            *bool           `json:"pendingDeployment"`
	Deploying                    *bool           `json:"deploying"`
	DeployingTimestamp           *time.Time      `json:"deployingTimestamp"`
	ScheduledDeploymentTimestamp *time.Time      `json:"scheduledDeploymentTimestamp"`
	RfqEnabled                   *bool           `json:"rfqEnabled"`
	EventStartTime               *time.Time      `json:"eventStartTime"`
}

// UserProfile represents a user profile in search results
type UserProfile struct {
	ID                    string          `json:"id"`
	Name                  *string         `json:"name"`
	User                  *float64        `json:"user"`
	Referral              *string         `json:"referral"`
	CreatedBy             *float64        `json:"createdBy"`
	UpdatedBy             *float64        `json:"updatedBy"`
	CreatedAt             *time.Time      `json:"createdAt"`
	UpdatedAt             *time.Time      `json:"updatedAt"`
	UtmSource             *string         `json:"utmSource"`
	UtmMedium             *string         `json:"utmMedium"`
	UtmCampaign           *string         `json:"utmCampaign"`
	UtmContent            *string         `json:"utmContent"`
	UtmTerm               *string         `json:"utmTerm"`
	WalletActivated       *bool           `json:"walletActivated"`
	Pseudonym             *string         `json:"pseudonym"`
	DisplayUsernamePublic *bool           `json:"displayUsernamePublic"`
	ProfileImage          *string         `json:"profileImage"`
	Bio                   *string         `json:"bio"`
	ProxyWallet           *string         `json:"proxyWallet"`
	ProfileImageOptimized *ImageOptimized `json:"profileImageOptimized"`
	IsCloseOnly           *bool           `json:"isCloseOnly"`
	IsCertReq             *bool           `json:"isCertReq"`
	CertReqDate           *time.Time      `json:"certReqDate"`
}

// PaginationInfo represents pagination metadata
type PaginationInfo struct {
	HasMore      bool `json:"hasMore"`
	TotalResults int  `json:"totalResults"`
}

// SearchResult represents the combined search response
type SearchResult struct {
	Events     []GammaEvent   `json:"events"`
	Tags       []TagWithCount `json:"tags"`
	Profiles   []UserProfile  `json:"profiles"`
	Pagination PaginationInfo `json:"pagination"`
}

// RelatedTagRelationship represents a relationship between tags
type RelatedTagRelationship struct {
	ID           string   `json:"id"`
	TagID        *float64 `json:"tagID"`
	RelatedTagID *float64 `json:"relatedTagID"`
	Rank         *float64 `json:"rank"`
}

// MarketFilters represents filtering options for markets
type MarketFilters struct {
	Limit               *int     `json:"limit,omitempty"`
	Offset              *int     `json:"offset,omitempty"`
	Order               *string  `json:"order,omitempty"`
	Ascending           *bool    `json:"ascending,omitempty"`
	ID                  []int    `json:"id,omitempty"`
	Slug                []string `json:"slug,omitempty"`
	ClobTokenIDs        []string `json:"clob_token_ids,omitempty"`
	ConditionIDs        []string `json:"condition_ids,omitempty"`
	MarketMakerAddress  []string `json:"market_maker_address,omitempty"`
	LiquidityNumMin     *float64 `json:"liquidity_num_min,omitempty"`
	LiquidityNumMax     *float64 `json:"liquidity_num_max,omitempty"`
	VolumeNumMin        *float64 `json:"volume_num_min,omitempty"`
	VolumeNumMax        *float64 `json:"volume_num_max,omitempty"`
	StartDateMin        *string  `json:"start_date_min,omitempty"`
	StartDateMax        *string  `json:"start_date_max,omitempty"`
	EndDateMin          *string  `json:"end_date_min,omitempty"`
	EndDateMax          *string  `json:"end_date_max,omitempty"`
	TagID               *int     `json:"tag_id,omitempty"`
	RelatedTags         *bool    `json:"related_tags,omitempty"`
	Cyom                *bool    `json:"cyom,omitempty"`
	UmaResolutionStatus *string  `json:"uma_resolution_status,omitempty"`
	GameID              *string  `json:"game_id,omitempty"`
	SportsMarketTypes   []string `json:"sports_market_types,omitempty"`
	RewardsMinSize      *float64 `json:"rewards_min_size,omitempty"`
	QuestionIDs         []string `json:"question_ids,omitempty"`
	IncludeTag          *bool    `json:"include_tag,omitempty"`
	Closed              *bool    `json:"closed,omitempty"`
}

// EventFilters represents filtering options for events
type EventFilters struct {
	Limit           *int     `json:"limit,omitempty"`
	Offset          *int     `json:"offset,omitempty"`
	Order           *string  `json:"order,omitempty"`
	Ascending       *bool    `json:"ascending,omitempty"`
	ID              []int    `json:"id,omitempty"`
	TagID           *int     `json:"tag_id,omitempty"`
	ExcludeTagID    []int    `json:"exclude_tag_id,omitempty"`
	Slug            []string `json:"slug,omitempty"`
	TagSlug         *string  `json:"tag_slug,omitempty"`
	RelatedTags     *bool    `json:"related_tags,omitempty"`
	Active          *bool    `json:"active,omitempty"`
	Archived        *bool    `json:"archived,omitempty"`
	Featured        *bool    `json:"featured,omitempty"`
	Cyom            *bool    `json:"cyom,omitempty"`
	IncludeChat     *bool    `json:"include_chat,omitempty"`
	IncludeTemplate *bool    `json:"include_template,omitempty"`
	Recurrence      *string  `json:"recurrence,omitempty"`
	Closed          *bool    `json:"closed,omitempty"`
	LiquidityMin    *float64 `json:"liquidity_min,omitempty"`
	LiquidityMax    *float64 `json:"liquidity_max,omitempty"`
	VolumeMin       *float64 `json:"volume_min,omitempty"`
	VolumeMax       *float64 `json:"volume_max,omitempty"`
	StartDateMin    *string  `json:"start_date_min,omitempty"`
	StartDateMax    *string  `json:"start_date_max,omitempty"`
	EndDateMin      *string  `json:"end_date_min,omitempty"`
	EndDateMax      *string  `json:"end_date_max,omitempty"`
}

// TagFilters represents filtering options for tags
type TagFilters struct {
	Limit           *int    `json:"limit,omitempty"`
	Offset          *int    `json:"offset,omitempty"`
	Order           *string `json:"order,omitempty"`
	Ascending       *bool   `json:"ascending,omitempty"`
	IncludeTemplate *bool   `json:"include_template,omitempty"`
	IsCarousel      *bool   `json:"is_carousel,omitempty"`
}

// SearchFilters represents filtering options for search
type SearchFilters struct {
	Query             string   `json:"q"` // Required
	Cache             *bool    `json:"cache,omitempty"`
	EventsStatus      *string  `json:"events_status,omitempty"`
	LimitPerType      *int     `json:"limit_per_type,omitempty"`
	Page              *int     `json:"page,omitempty"`
	EventsTag         []string `json:"events_tag,omitempty"`
	KeepClosedMarkets *int     `json:"keep_closed_markets,omitempty"`
	Sort              *string  `json:"sort,omitempty"`
	Ascending         *bool    `json:"ascending,omitempty"`
	SearchTags        *bool    `json:"search_tags,omitempty"`
	SearchProfiles    *bool    `json:"search_profiles,omitempty"`
	Recurrence        *string  `json:"recurrence,omitempty"`
	ExcludeTagID      []int    `json:"exclude_tag_id,omitempty"`
	Optimized         *bool    `json:"optimized,omitempty"`
}

// RelatedTagFilters represents filtering options for related tags
type RelatedTagFilters struct {
	OmitEmpty *bool   `json:"omit_empty,omitempty"`
	Status    *string `json:"status,omitempty"` // active, closed, all
}
