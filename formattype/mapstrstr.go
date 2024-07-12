package formattype

import (
	"github.com/json-iterator/go"
	"strings"
)

type FormatMapStrStr map[string]string

func (m *FormatMapStrStr) UnmarshalJSON(data []byte) error {
	tm := make(map[string]string)
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &tm)
	if err != nil {
		return err
	}
	fm := make(map[string]string)
	for k, v := range tm {
		fm[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
	*m = fm
	return nil
}
