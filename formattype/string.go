package formattype

import (
	"github.com/json-iterator/go"
	"strings"
)

type FormatString string

func (s *FormatString) UnmarshalJSON(data []byte) error {
	var ts string
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &ts)
	if err != nil {
		return err
	}
	fs := strings.TrimSpace(ts)
	*s = FormatString(fs)
	return nil
}
