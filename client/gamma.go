package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/lajosdeme/polymarket-go-api/types"
)

// GammaClient represents the Gamma API client for market data
type GammaClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewGammaClient creates a new Gamma client
func NewGammaClient(baseURL string) *GammaClient {
	if baseURL == "" {
		baseURL = "https://gamma-api.polymarket.com"
	}

	return &GammaClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SetTimeout sets HTTP client timeout
func (c *GammaClient) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// DoGet performs a GET request to the Gamma API
func (c *GammaClient) DoGet(ctx context.Context, path string, queryParams map[string]string) ([]byte, error) {
	// Build URL with query parameters
	requestURL := c.baseURL + path
	if len(queryParams) > 0 {
		values := url.Values{}
		for key, value := range queryParams {
			values.Add(key, value)
		}
		requestURL += "?" + values.Encode()
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// DoPost performs a POST request to the Gamma API
func (c *GammaClient) DoPost(ctx context.Context, path string, body interface{}) ([]byte, error) {
	// Prepare request body
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	// Create request
	requestURL := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// BuildQueryParams converts filter structs to query parameters map
func (c *GammaClient) BuildQueryParams(filters any) map[string]string {
	params := make(map[string]string)

	switch f := filters.(type) {
	case *types.MarketFilters:
		if f.Limit != nil {
			params["limit"] = strconv.Itoa(*f.Limit)
		}
		if f.Offset != nil {
			params["offset"] = strconv.Itoa(*f.Offset)
		}
		if f.Order != nil {
			params["order"] = *f.Order
		}
		if f.Ascending != nil {
			params["ascending"] = strconv.FormatBool(*f.Ascending)
		}
		if len(f.ID) > 0 {
			for _, id := range f.ID {
				params["id"] = strconv.Itoa(id)
			}
		}
		if len(f.Slug) > 0 {
			for _, slug := range f.Slug {
				params["slug"] = slug
			}
		}
		if len(f.ClobTokenIDs) > 0 {
			for _, tokenID := range f.ClobTokenIDs {
				params["clob_token_ids"] = tokenID
			}
		}
		if len(f.ConditionIDs) > 0 {
			for _, conditionID := range f.ConditionIDs {
				params["condition_ids"] = conditionID
			}
		}
		if len(f.MarketMakerAddress) > 0 {
			for _, addr := range f.MarketMakerAddress {
				params["market_maker_address"] = addr
			}
		}
		if f.LiquidityNumMin != nil {
			params["liquidity_num_min"] = fmt.Sprintf("%.f", *f.LiquidityNumMin)
		}
		if f.LiquidityNumMax != nil {
			params["liquidity_num_max"] = fmt.Sprintf("%.f", *f.LiquidityNumMax)
		}
		if f.VolumeNumMin != nil {
			params["volume_num_min"] = fmt.Sprintf("%.f", *f.VolumeNumMin)
		}
		if f.VolumeNumMax != nil {
			params["volume_num_max"] = fmt.Sprintf("%.f", *f.VolumeNumMax)
		}
		if f.StartDateMin != nil {
			params["start_date_min"] = *f.StartDateMin
		}
		if f.StartDateMax != nil {
			params["start_date_max"] = *f.StartDateMax
		}
		if f.EndDateMin != nil {
			params["end_date_min"] = *f.EndDateMin
		}
		if f.EndDateMax != nil {
			params["end_date_max"] = *f.EndDateMax
		}
		if f.TagID != nil {
			params["tag_id"] = strconv.Itoa(*f.TagID)
		}
		if f.RelatedTags != nil {
			params["related_tags"] = strconv.FormatBool(*f.RelatedTags)
		}
		if f.Cyom != nil {
			params["cyom"] = strconv.FormatBool(*f.Cyom)
		}
		if f.UmaResolutionStatus != nil {
			params["uma_resolution_status"] = *f.UmaResolutionStatus
		}
		if f.GameID != nil {
			params["game_id"] = *f.GameID
		}
		if len(f.SportsMarketTypes) > 0 {
			for _, marketType := range f.SportsMarketTypes {
				params["sports_market_types"] = marketType
			}
		}
		if f.RewardsMinSize != nil {
			params["rewards_min_size"] = fmt.Sprintf("%.f", *f.RewardsMinSize)
		}
		if len(f.QuestionIDs) > 0 {
			for _, questionID := range f.QuestionIDs {
				params["question_ids"] = questionID
			}
		}
		if f.IncludeTag != nil {
			params["include_tag"] = strconv.FormatBool(*f.IncludeTag)
		}
		if f.Closed != nil {
			params["closed"] = strconv.FormatBool(*f.Closed)
		}

	case *types.EventFilters:
		if f.Limit != nil {
			params["limit"] = strconv.Itoa(*f.Limit)
		}
		if f.Offset != nil {
			params["offset"] = strconv.Itoa(*f.Offset)
		}
		if f.Order != nil {
			params["order"] = *f.Order
		}
		if f.Ascending != nil {
			params["ascending"] = strconv.FormatBool(*f.Ascending)
		}
		if len(f.ID) > 0 {
			for _, id := range f.ID {
				params["id"] = strconv.Itoa(id)
			}
		}
		if f.TagID != nil {
			params["tag_id"] = strconv.Itoa(*f.TagID)
		}
		if len(f.ExcludeTagID) > 0 {
			for _, tagID := range f.ExcludeTagID {
				params["exclude_tag_id"] = strconv.Itoa(tagID)
			}
		}
		if len(f.Slug) > 0 {
			for _, slug := range f.Slug {
				params["slug"] = slug
			}
		}
		if f.TagSlug != nil {
			params["tag_slug"] = *f.TagSlug
		}
		if f.RelatedTags != nil {
			params["related_tags"] = strconv.FormatBool(*f.RelatedTags)
		}
		if f.Active != nil {
			params["active"] = strconv.FormatBool(*f.Active)
		}
		if f.Archived != nil {
			params["archived"] = strconv.FormatBool(*f.Archived)
		}
		if f.Featured != nil {
			params["featured"] = strconv.FormatBool(*f.Featured)
		}
		if f.Cyom != nil {
			params["cyom"] = strconv.FormatBool(*f.Cyom)
		}
		if f.IncludeChat != nil {
			params["include_chat"] = strconv.FormatBool(*f.IncludeChat)
		}
		if f.IncludeTemplate != nil {
			params["include_template"] = strconv.FormatBool(*f.IncludeTemplate)
		}
		if f.Recurrence != nil {
			params["recurrence"] = *f.Recurrence
		}
		if f.Closed != nil {
			params["closed"] = strconv.FormatBool(*f.Closed)
		}
		if f.LiquidityMin != nil {
			params["liquidity_min"] = fmt.Sprintf("%.f", *f.LiquidityMin)
		}
		if f.LiquidityMax != nil {
			params["liquidity_max"] = fmt.Sprintf("%.f", *f.LiquidityMax)
		}
		if f.VolumeMin != nil {
			params["volume_min"] = fmt.Sprintf("%.f", *f.VolumeMin)
		}
		if f.VolumeMax != nil {
			params["volume_max"] = fmt.Sprintf("%.f", *f.VolumeMax)
		}
		if f.StartDateMin != nil {
			params["start_date_min"] = *f.StartDateMin
		}
		if f.StartDateMax != nil {
			params["start_date_max"] = *f.StartDateMax
		}
		if f.EndDateMin != nil {
			params["end_date_min"] = *f.EndDateMin
		}
		if f.EndDateMax != nil {
			params["end_date_max"] = *f.EndDateMax
		}

	case *types.TagFilters:
		if f.Limit != nil {
			params["limit"] = strconv.Itoa(*f.Limit)
		}
		if f.Offset != nil {
			params["offset"] = strconv.Itoa(*f.Offset)
		}
		if f.Order != nil {
			params["order"] = *f.Order
		}
		if f.Ascending != nil {
			params["ascending"] = strconv.FormatBool(*f.Ascending)
		}
		if f.IncludeTemplate != nil {
			params["include_template"] = strconv.FormatBool(*f.IncludeTemplate)
		}
		if f.IsCarousel != nil {
			params["is_carousel"] = strconv.FormatBool(*f.IsCarousel)
		}

	case *types.SearchFilters:
		params["q"] = f.Query // Required
		if f.Cache != nil {
			params["cache"] = strconv.FormatBool(*f.Cache)
		}
		if f.EventsStatus != nil {
			params["events_status"] = *f.EventsStatus
		}
		if f.LimitPerType != nil {
			params["limit_per_type"] = strconv.Itoa(*f.LimitPerType)
		}
		if f.Page != nil {
			params["page"] = strconv.Itoa(*f.Page)
		}
		if len(f.EventsTag) > 0 {
			for _, tag := range f.EventsTag {
				params["events_tag"] = tag
			}
		}
		if f.KeepClosedMarkets != nil {
			params["keep_closed_markets"] = strconv.Itoa(*f.KeepClosedMarkets)
		}
		if f.Sort != nil {
			params["sort"] = *f.Sort
		}
		if f.Ascending != nil {
			params["ascending"] = strconv.FormatBool(*f.Ascending)
		}
		if f.SearchTags != nil {
			params["search_tags"] = strconv.FormatBool(*f.SearchTags)
		}
		if f.SearchProfiles != nil {
			params["search_profiles"] = strconv.FormatBool(*f.SearchProfiles)
		}
		if f.Recurrence != nil {
			params["recurrence"] = *f.Recurrence
		}
		if len(f.ExcludeTagID) > 0 {
			for _, tagID := range f.ExcludeTagID {
				params["exclude_tag_id"] = strconv.Itoa(tagID)
			}
		}
		if f.Optimized != nil {
			params["optimized"] = strconv.FormatBool(*f.Optimized)
		}

	case *types.RelatedTagFilters:
		if f.OmitEmpty != nil {
			params["omit_empty"] = strconv.FormatBool(*f.OmitEmpty)
		}
		if f.Status != nil {
			params["status"] = *f.Status
		}
	}

	return params
}
