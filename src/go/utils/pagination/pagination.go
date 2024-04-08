package pagination

import protobase "market.goravel.dev/proto/base"

func Default() *protobase.Pagination {
	return &protobase.Pagination{
		Page:  1,
		Limit: 10,
	}
}
