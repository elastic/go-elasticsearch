package gensource

var (
	overrideRules map[string][]OverrideRule
)

// OverrideFunc defines a function to override generated code for endpoint.
//
type OverrideFunc func(*Endpoint, ...interface{}) string

// OverrideRule represents an override rule.
//
type OverrideRule struct {
	Func     OverrideFunc
	Matching []string
}

// GetOverride returns an override function for id and API name.
//
func (g *Generator) GetOverride(id, apiName string) OverrideFunc {
	if rr, ok := overrideRules[id]; ok {
		for _, r := range rr {
			if r.Match(apiName) {
				return r.Func
			}
		}
	}
	return nil
}

// Match returns true when API name matches a rule.
//
func (r OverrideRule) Match(apiName string) bool {
	for _, v := range r.Matching {
		if v == "*" {
			return true
		}
		if v == apiName {
			return true
		}
	}
	return false
}

func init() {
	overrideRules = map[string][]OverrideRule{

		"polymorphic-param": {{
			Matching: []string{"search"},
			Func: func(e *Endpoint, i ...interface{}) string {
				if len(i) > 0 {
					switch i[0] {
					case "track_total_hits":
						return "interface{}"
					}
				}
				return ""
			},
		}},

		"url": {
			{
				Matching: []string{"cluster.stats"},
				Func: func(*Endpoint, ...interface{}) string {
					return `
	path.Grow(len("/nodes/_cluster/stats/nodes/") + len(strings.Join(r.NodeID, ",")))
	path.WriteString("/")
	path.WriteString("_cluster")
	path.WriteString("/")
	path.WriteString("stats")
	if len(r.NodeID) > 0 {
		path.WriteString("/")
		path.WriteString("nodes")
		path.WriteString("/")
		path.WriteString(strings.Join(r.NodeID, ","))
	}
`
				},
			},
			{
				Matching: []string{"indices.put_mapping"},
				Func: func(*Endpoint, ...interface{}) string {
					return `
	path.Grow(len(strings.Join(r.Index, ",")) + len("/_mapping") + len(r.DocumentType) + 2)
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Index, ","))
	}
	path.WriteString("/")
	path.WriteString("_mapping")
	if r.DocumentType != "" {
		path.WriteString("/")
		path.WriteString(r.DocumentType)
	}
`
				},
			},
		},
	}
}
