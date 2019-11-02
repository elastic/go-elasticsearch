// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package estransport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// Discoverable defines the interface for transports supporting node discovery.
//
type Discoverable interface {
	DiscoverNodes() error
}

// nodeInfo represents the information about node in a cluster.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-info.html
//
type nodeInfo struct {
	ID         string
	Name       string
	URL        *url.URL
	Roles      []string
	Attributes map[string]interface{}
	HTTP       struct {
		PublishAddress string `json:"publish_address"`
	}
}

// DiscoverNodes reloads the client connections by fetching information from the cluster.
//
func (c *Client) DiscoverNodes() error {
	nodes, err := c.getNodesInfo()
	if err != nil {
		return err
	}

	fmt.Println("DiscoverNodes:")
	for i, n := range nodes {
		fmt.Printf("%d. %+v\n", i+1, n)
	}
	fmt.Println("------------------------------------------------------------")

	if len(nodes) < 2 {
		c.pool = newSingleConnectionPool(nodes[0].URL)
		return nil
	}

	var orig []*url.URL
	if cprr, ok := c.pool.(*roundRobinConnectionPool); ok {
		orig = cprr.orig
	}

	cp := &roundRobinConnectionPool{orig: orig, curr: -1}

	cp.Lock()
	defer cp.Unlock()

	for _, node := range nodes {
		fmt.Printf("Checking node %s", node.URL)
		var (
			isDataNode   bool
			isIngestNode bool
		)

		roles := append(node.Roles[:0:0], node.Roles...)
		sort.Strings(roles)

		if i := sort.SearchStrings(roles, "data"); i < len(roles) && roles[i] == "data" {
			isDataNode = true
		}
		if i := sort.SearchStrings(roles, "ingest"); i < len(roles) && roles[i] == "ingest" {
			isIngestNode = true
		}

		fmt.Printf("; roles=%s", node.Roles)

		// Skip master only nodes
		//
		if !isDataNode || !isIngestNode {
			fmt.Printf("; SKIPPING\n")
			continue
		} else {
			fmt.Printf("\n")
		}

		cp.live = append(cp.live, &Connection{
			URL:        node.URL,
			ID:         node.ID,
			Name:       node.Name,
			Roles:      node.Roles,
			Attributes: node.Attributes,
		})
	}

	// TODO(karmi): c.pool.Replace(cp)
	c.pool = cp

	return nil
}

func (c *Client) getNodesInfo() ([]nodeInfo, error) {
	var (
		out    []nodeInfo
		scheme = c.urls[0].Scheme
	)

	req, err := http.NewRequest("GET", "/_nodes/http", nil)
	if err != nil {
		return out, err
	}

	conn, err := c.getConnection()
	if err != nil {
		return out, err
	}

	c.setReqURL(conn.URL, req)
	c.setReqAuth(conn.URL, req)
	c.setReqUserAgent(req)

	res, err := c.transport.RoundTrip(req)
	if err != nil {
		return out, err
	}
	defer res.Body.Close()

	if res.StatusCode > 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return out, fmt.Errorf("server error: %s: %s", res.Status, body)
	}

	var env map[string]json.RawMessage
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		return out, err
	}

	// fmt.Printf("%s\n", env["nodes"])
	var nodes map[string]nodeInfo
	if err := json.Unmarshal(env["nodes"], &nodes); err != nil {
		return out, err
	}

	for id, node := range nodes {
		// fmt.Printf("%+v\n", node)
		node.ID = id
		u, err := c.getNodeURL(node, scheme)
		if err != nil {
			return out, err
		}
		node.URL = u
		out = append(out, node)
	}

	return out, nil
}

func (c *Client) getNodeURL(node nodeInfo, scheme string) (*url.URL, error) {
	var (
		host string
		port string

		addrs = strings.Split(node.HTTP.PublishAddress, "/")
		ports = strings.Split(node.HTTP.PublishAddress, ":")
	)

	if len(addrs) > 1 {
		host = addrs[0]
	} else {
		host = strings.Split(addrs[0], ":")[0]
	}
	port = ports[len(ports)-1]

	u := &url.URL{
		Scheme: scheme,
		Host:   host + ":" + port,
	}

	if c.username != "" || c.password != "" {
		u.User = url.UserPassword(c.username, c.password)
	}

	return u, nil
}
