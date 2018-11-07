package persist

import (
	"WebSpider/crawler/engine"
	"WebSpider/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(item, s.Client, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}
