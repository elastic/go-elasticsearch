// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package estransport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Discoverable defines the interface for transports supporting node discovery.
//
type Discoverable interface {
	DiscoverNodes() error
}

// NodeInfo represents the information about node in a cluster.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-info.html
//
type NodeInfo struct {
	ID         string
	Name       string
	Address    string
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
	for _, n := range nodes {
		fmt.Printf("==> %+v\n", n)
	}
	fmt.Println("------------------------------------------------------------")

	var urls []*url.URL
	for _, n := range nodes {
		u, err := url.Parse(n.Address)
		if err != nil {
			return err
		}
		urls = append(urls, u)
	}

	c.pool = newRoundRobinConnectionPool(urls...)

	return nil
}

func (c *Client) getNodesInfo() ([]NodeInfo, error) {
	var (
		out    []NodeInfo
		scheme = c.urls[0].Scheme
	)

	req, err := http.NewRequest("GET", "_nodes/http", nil)
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

	// fmt.Printf("%+v\n", req)

	res, err := c.transport.RoundTrip(req)
	if err != nil {
		c.pool.Remove(conn)
		return out, err
	}
	defer res.Body.Close()

	var env map[string]json.RawMessage
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		return out, err
	}

	// fmt.Printf("%s\n", env["nodes"])
	var nodes map[string]NodeInfo
	if err := json.Unmarshal(env["nodes"], &nodes); err != nil {
		return out, err
	}

	for id, node := range nodes {
		// fmt.Printf("%+v\n", node)
		node.ID = id
		nodeAddr, err := getNodeAddress(node, scheme)
		if err != nil {
			return out, err
		}
		node.Address = nodeAddr
		out = append(out, node)
	}

	return out, nil
}

func getNodeAddress(node NodeInfo, scheme string) (string, error) {
	return scheme + "://" + node.HTTP.PublishAddress, nil
}
