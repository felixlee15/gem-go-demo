package graph

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// format ISO8601
		_, _ = io.WriteString(w, fmt.Sprintf("%q", t.Format(time.RFC3339)))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		// parse ISO8601
		return time.Parse(time.RFC3339, v)
	case time.Time:
		return v, nil
	default:
		return time.Time{}, fmt.Errorf("time should be RFC3339 string, got %T", v)
	}
}
