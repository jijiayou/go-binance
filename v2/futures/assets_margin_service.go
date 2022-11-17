package futures

import (
	"context"
	"net/http"
)

// ChangeMultiAssetsMarginModeService change user's multiAssetsMargin mode
type ChangeMultiAssetsMarginModeService struct {
	c                 *Client
	multiAssetsMargin string
}

func (s *ChangeMultiAssetsMarginModeService) MultiAssetsMarginMode(multiAssetsMargin string) *ChangeMultiAssetsMarginModeService {
	s.multiAssetsMargin = multiAssetsMargin
	return s
}

// Do send request
func (s *ChangeMultiAssetsMarginModeService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/multiAssetsMargin",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"multiAssetsMargin": s.multiAssetsMargin,
	})
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}
