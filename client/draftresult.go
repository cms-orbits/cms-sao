// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": draftresult Resource Client
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.4.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetDraftresultPath computes a request path to the get action of draftresult.
func GetDraftresultPath(resultID string) string {
	param0 := resultID

	return fmt.Sprintf("/sao/v1/draft-results/%s", param0)
}

// Get complete Entry Draft Result data for the given entry and testcase ID.
func (c *Client) GetDraftresult(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetDraftresultRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetDraftresultRequest create the request corresponding to the get action endpoint of the draftresult resource.
func (c *Client) NewGetDraftresultRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowDraftresultPath computes a request path to the show action of draftresult.
func ShowDraftresultPath() string {

	return fmt.Sprintf("/sao/v1/draft-results/")
}

// List the Results delimited and grouped by contest, task, entry or user identifier
func (c *Client) ShowDraftresult(ctx context.Context, path string, contest *int, entry *int, page *int, pageSize *int, sort *string, task *int, user *int) (*http.Response, error) {
	req, err := c.NewShowDraftresultRequest(ctx, path, contest, entry, page, pageSize, sort, task, user)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowDraftresultRequest create the request corresponding to the show action endpoint of the draftresult resource.
func (c *Client) NewShowDraftresultRequest(ctx context.Context, path string, contest *int, entry *int, page *int, pageSize *int, sort *string, task *int, user *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if contest != nil {
		tmp11 := strconv.Itoa(*contest)
		values.Set("contest", tmp11)
	}
	if entry != nil {
		tmp12 := strconv.Itoa(*entry)
		values.Set("entry", tmp12)
	}
	if page != nil {
		tmp13 := strconv.Itoa(*page)
		values.Set("page", tmp13)
	}
	if pageSize != nil {
		tmp14 := strconv.Itoa(*pageSize)
		values.Set("page_size", tmp14)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	if task != nil {
		tmp15 := strconv.Itoa(*task)
		values.Set("task", tmp15)
	}
	if user != nil {
		tmp16 := strconv.Itoa(*user)
		values.Set("user", tmp16)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}