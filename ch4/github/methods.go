package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// This queries the GitHub issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// Must close resp.Body on all execution paths
	// Using 'Defer' makes this simpler
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}

	// Earlier examples used json.Unmarshal to decode the entire
	// contents of a byte slice as a single JSON entity.
	// However, this example using the streaming decoder json.Decoder,
	// which allows several JSON entities to be decoded in sequence from the same stream
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
