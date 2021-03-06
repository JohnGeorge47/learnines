package esqueries

import "context"

type Iquery interface {
	PingEs(conn string, ctx context.Context) (*string, error)
	Search(tosearch string, start int, end int, ctx context.Context) (*[]Game, error)
}
