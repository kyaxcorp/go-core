package scalar

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"time"
)

func MarshalTimeRFC3339(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// item.Format(time.RFC3339)
		w.Write([]byte("\"" + t.Format(time.RFC3339) + "\""))
	})
}

func UnmarshalTimeRFC3339(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339, v)
	default:
		return time.Time{}, fmt.Errorf("%T is not a string", v)
	}
}
