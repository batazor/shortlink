/*
Metadata Service. Application layer
*/
package application

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	rpc "github.com/batazor/shortlink/internal/services/metadata/domain"
	meta_store "github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
)

type Service struct {
	Store *meta_store.MetaStore
}

func (r *Service) Get(ctx context.Context, hash string) (*rpc.Meta, error) {
	meta, err := r.Store.Store.Get(ctx, hash)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (r *Service) Set(ctx context.Context, url string) (*rpc.Meta, error) {
	meta := &rpc.Meta{
		Id: url,
	}

	// Request the HTML page.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" {
			meta.Description, _ = s.Attr("content")
		}

		if name, _ := s.Attr("name"); name == "keyworlds" {
			meta.Keywords, _ = s.Attr("content")
		}
	})

	// Write to DB
	err = r.Store.Store.Add(ctx, meta)
	if err != nil {
		return nil, err
	}

	return meta, nil
}
