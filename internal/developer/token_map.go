package developer

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Token struct {
	Key        []byte    `json:"key"`
	ValidUntil time.Time `json:"valid_until"`
}

type TokenMap map[Environment][]Token

func (t TokenMap) Value() (driver.Value, error) {
	j, err := json.Marshal(t)
	return j, err
}

func (t *TokenMap) Scan(src interface{}) error {
	source, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion .(string) failed, got %T", src)
	}
	var i TokenMap
	if err := json.Unmarshal([]byte(source), &i); err != nil {
		return err
	}

	*t = i

	return nil
}
