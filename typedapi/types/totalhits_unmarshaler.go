package types

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/totalhitsrelation"
)

// UnmarshalJSON implements Unmarshaler interface, it handles the shortcut for total hits.
func (t *TotalHits) UnmarshalJSON(data []byte) error {
	type stub TotalHits
	tmp := stub{}
	dec := json.NewDecoder(bytes.NewReader(data))
	if _, err := strconv.Atoi(string(data)); err == nil {
		err := dec.Decode(&t.Value)
		if err != nil {
			return err
		}
		t.Relation = totalhitsrelation.Eq
	} else {
		err := dec.Decode(&tmp)
		if err != nil {
			return err
		}
		*t = TotalHits(tmp)
	}

	return nil
}
