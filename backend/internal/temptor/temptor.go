package temptor

import "context"

type Temptor interface {
	Generate(ctx context.Context, uid uint64, body string) (address string, err error)
}
