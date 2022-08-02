package scalar

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"io"
)

func MarshalNonNilUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if u == uuid.Nil {
			// todo: what to do in this case?!
		}
		w.Write([]byte("\"" + u.String() + "\""))
	})
}

func UnmarshalNonNilUUID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		_uuid, _err := uuid.Parse(v)
		if _err != nil {
			return uuid.Nil, _err
		}
		if _uuid == uuid.Nil {
			return uuid.Nil, fmt.Errorf("%T nil uuid is not permitted", v)
		}
		return _uuid, nil
	default:
		return uuid.Nil, fmt.Errorf("%T is not a string", v)
	}
}
