package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/lajosdeme/polymarket-go-api/client"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// GammaAPI handles Gamma API operations for market data
type GammaAPI struct {
	client *client.GammaClient
}

// NewGammaAPI creates a new GammaAPI instance
func NewGammaAPI(client *client.GammaClient) *GammaAPI {
	return &GammaAPI{
		client: client,
	}
}

// GetMarkets retrieves a list of markets with optional filtering
func (g *GammaAPI) GetMarkets(ctx context.Context, filters *types.MarketFilters) ([]types.GammaMarket, error) {
	queryParams := g.client.BuildQueryParams(filters)

	body, err := g.client.DoGet(ctx, "/markets", queryParams)
	if err != nil {
		return nil, err
	}

	var markets []types.GammaMarket
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return markets, nil
}

// GetMarketByID retrieves a specific market by ID
func (g *GammaAPI) GetMarketByID(ctx context.Context, id int, includeTag *bool) (*types.GammaMarket, error) {
	queryParams := map[string]string{}
	if includeTag != nil {
		queryParams["include_tag"] = strconv.FormatBool(*includeTag)
	}

	body, err := g.client.DoGet(ctx, "/markets/"+strconv.Itoa(id), queryParams)
	if err != nil {
		return nil, err
	}

	var market types.GammaMarket
	if err := json.Unmarshal(body, &market); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &market, nil
}

// GetMarketBySlug retrieves a specific market by slug
func (g *GammaAPI) GetMarketBySlug(ctx context.Context, slug string, includeTag *bool) (*types.GammaMarket, error) {
	queryParams := map[string]string{}
	if includeTag != nil {
		queryParams["include_tag"] = strconv.FormatBool(*includeTag)
	}

	body, err := g.client.DoGet(ctx, "/markets/slug/"+slug, queryParams)
	if err != nil {
		return nil, err
	}

	var market types.GammaMarket
	if err := json.Unmarshal(body, &market); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &market, nil
}

// GetMarketTags retrieves tags for a specific market
func (g *GammaAPI) GetMarketTags(ctx context.Context, id int) ([]types.Tag, error) {
	body, err := g.client.DoGet(ctx, "/markets/"+strconv.Itoa(id)+"/tags", nil)
	if err != nil {
		return nil, err
	}

	var tags []types.Tag
	if err := json.Unmarshal(body, &tags); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return tags, nil
}

// GetEvents retrieves a list of events with optional filtering
func (g *GammaAPI) GetEvents(ctx context.Context, filters *types.EventFilters) ([]types.GammaEvent, error) {
	queryParams := g.client.BuildQueryParams(filters)

	body, err := g.client.DoGet(ctx, "/events", queryParams)
	if err != nil {
		return nil, err
	}

	var events []types.GammaEvent
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return events, nil
}

// GetEventByID retrieves a specific event by ID
func (g *GammaAPI) GetEventByID(ctx context.Context, id int, includeChat, includeTemplate *bool) (*types.GammaEvent, error) {
	queryParams := map[string]string{}
	if includeChat != nil {
		queryParams["include_chat"] = strconv.FormatBool(*includeChat)
	}
	if includeTemplate != nil {
		queryParams["include_template"] = strconv.FormatBool(*includeTemplate)
	}

	body, err := g.client.DoGet(ctx, "/events/"+strconv.Itoa(id), queryParams)
	if err != nil {
		return nil, err
	}

	var event types.GammaEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &event, nil
}

// GetEventBySlug retrieves a specific event by slug
func (g *GammaAPI) GetEventBySlug(ctx context.Context, slug string, includeChat, includeTemplate *bool) (*types.GammaEvent, error) {
	queryParams := map[string]string{}
	if includeChat != nil {
		queryParams["include_chat"] = strconv.FormatBool(*includeChat)
	}
	if includeTemplate != nil {
		queryParams["include_template"] = strconv.FormatBool(*includeTemplate)
	}

	body, err := g.client.DoGet(ctx, "/events/slug/"+slug, queryParams)
	if err != nil {
		return nil, err
	}

	var event types.GammaEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &event, nil
}

// GetEventTags retrieves tags for a specific event
func (g *GammaAPI) GetEventTags(ctx context.Context, id int) ([]types.Tag, error) {
	body, err := g.client.DoGet(ctx, "/events/"+strconv.Itoa(id)+"/tags", nil)
	if err != nil {
		return nil, err
	}

	var tags []types.Tag
	if err := json.Unmarshal(body, &tags); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return tags, nil
}

// GetTags retrieves a list of tags with optional filtering
func (g *GammaAPI) GetTags(ctx context.Context, filters *types.TagFilters) ([]types.Tag, error) {
	queryParams := g.client.BuildQueryParams(filters)

	body, err := g.client.DoGet(ctx, "/tags", queryParams)
	if err != nil {
		return nil, err
	}

	var tags []types.Tag
	if err := json.Unmarshal(body, &tags); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return tags, nil
}

// GetTagByID retrieves a specific tag by ID
func (g *GammaAPI) GetTagByID(ctx context.Context, id int, includeTemplate *bool) (*types.Tag, error) {
	queryParams := map[string]string{}
	if includeTemplate != nil {
		queryParams["include_template"] = strconv.FormatBool(*includeTemplate)
	}

	body, err := g.client.DoGet(ctx, "/tags/"+strconv.Itoa(id), queryParams)
	if err != nil {
		return nil, err
	}

	var tag types.Tag
	if err := json.Unmarshal(body, &tag); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &tag, nil
}

// GetTagBySlug retrieves a specific tag by slug
func (g *GammaAPI) GetTagBySlug(ctx context.Context, slug string, includeTemplate *bool) (*types.Tag, error) {
	queryParams := map[string]string{}
	if includeTemplate != nil {
		queryParams["include_template"] = strconv.FormatBool(*includeTemplate)
	}

	body, err := g.client.DoGet(ctx, "/tags/slug/"+slug, queryParams)
	if err != nil {
		return nil, err
	}

	var tag types.Tag
	if err := json.Unmarshal(body, &tag); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &tag, nil
}

// GetRelatedTagsByTagID retrieves tag relationships by tag ID
func (g *GammaAPI) GetRelatedTagsByTagID(ctx context.Context, id int, filters *types.RelatedTagFilters) ([]types.RelatedTagRelationship, error) {
	queryParams := g.client.BuildQueryParams(filters)

	body, err := g.client.DoGet(ctx, "/tags/"+strconv.Itoa(id)+"/related-tags", queryParams)
	if err != nil {
		return nil, err
	}

	var relationships []types.RelatedTagRelationship
	if err := json.Unmarshal(body, &relationships); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return relationships, nil
}

// GetRelatedTagsByTagSlug retrieves tag relationships by tag slug
func (g *GammaAPI) GetRelatedTagsByTagSlug(ctx context.Context, slug string, filters *types.RelatedTagFilters) ([]types.RelatedTagRelationship, error) {
	queryParams := g.client.BuildQueryParams(filters)

	body, err := g.client.DoGet(ctx, "/tags/slug/"+slug+"/related-tags", queryParams)
	if err != nil {
		return nil, err
	}

	var relationships []types.RelatedTagRelationship
	if err := json.Unmarshal(body, &relationships); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return relationships, nil
}

// GetRelatedTagsByTagIDTags retrieves actual tags related to a tag ID
func (g *GammaAPI) GetRelatedTagsByTagIDTags(ctx context.Context, id int, filters *types.RelatedTagFilters) ([]types.Tag, error) {
	queryParams := g.client.BuildQueryParams(filters)

	body, err := g.client.DoGet(ctx, "/tags/"+strconv.Itoa(id)+"/related-tags/tags", queryParams)
	if err != nil {
		return nil, err
	}

	var tags []types.Tag
	if err := json.Unmarshal(body, &tags); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return tags, nil
}

// GetRelatedTagsByTagSlugTags retrieves actual tags related to a tag slug
func (g *GammaAPI) GetRelatedTagsByTagSlugTags(ctx context.Context, slug string, filters *types.RelatedTagFilters) ([]types.Tag, error) {
	queryParams := g.client.BuildQueryParams(filters)

	body, err := g.client.DoGet(ctx, "/tags/slug/"+slug+"/related-tags/tags", queryParams)
	if err != nil {
		return nil, err
	}

	var tags []types.Tag
	if err := json.Unmarshal(body, &tags); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return tags, nil
}

// Search performs a search across markets, events, and profiles
func (g *GammaAPI) Search(ctx context.Context, filters *types.SearchFilters) (*types.SearchResult, error) {
	queryParams := g.client.BuildQueryParams(filters)

	body, err := g.client.DoGet(ctx, "/public-search", queryParams)
	if err != nil {
		return nil, err
	}

	var searchResult types.SearchResult
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &searchResult, nil
}

// GetActiveMarkets retrieves all active markets (convenience method)
func (g *GammaAPI) GetActiveMarkets(ctx context.Context, limit, offset *int) ([]types.GammaMarket, error) {
	filters := &types.MarketFilters{
		Limit:     limit,
		Offset:    offset,
		Closed:    boolPtr(false), // Only active markets
		Order:     stringPtr("id"),
		Ascending: boolPtr(false), // Newest first
	}
	return g.GetMarkets(ctx, filters)
}

// GetActiveEvents retrieves all active events (convenience method)
func (g *GammaAPI) GetActiveEvents(ctx context.Context, limit, offset *int) ([]types.GammaEvent, error) {
	filters := &types.EventFilters{
		Limit:     limit,
		Offset:    offset,
		Closed:    boolPtr(false), // Only active events
		Order:     stringPtr("id"),
		Ascending: boolPtr(false), // Newest first
	}
	return g.GetEvents(ctx, filters)
}

// GetMarketsByTag retrieves markets filtered by tag ID (convenience method)
func (g *GammaAPI) GetMarketsByTag(ctx context.Context, tagID int, limit, offset *int) ([]types.GammaMarket, error) {
	filters := &types.MarketFilters{
		Limit:  limit,
		Offset: offset,
		TagID:  &tagID,
		Closed: boolPtr(false), // Only active markets
	}
	return g.GetMarkets(ctx, filters)
}

// GetEventsByTag retrieves events filtered by tag ID (convenience method)
func (g *GammaAPI) GetEventsByTag(ctx context.Context, tagID int, limit, offset *int) ([]types.GammaEvent, error) {
	filters := &types.EventFilters{
		Limit:  limit,
		Offset: offset,
		TagID:  &tagID,
		Closed: boolPtr(false), // Only active events
	}
	return g.GetEvents(ctx, filters)
}

// Helper functions for pointer creation
func boolPtr(b bool) *bool {
	return &b
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}
