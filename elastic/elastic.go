package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/olivere/elastic"
	"github.com/prometheus/common/log"
	"github.com/storefinder/pkg/models"
)

//Proxy elastic search proxy
type Proxy struct {
	config  ProxyConfig
	context context.Context
	client  *elastic.Client
}

//ProxyConfig represents EsProxy configuration
type ProxyConfig struct {
	ElasticURL *url.URL
	Username   string
	Password   string
}

//NewProxy initializes a new proxy
func NewProxy(config ProxyConfig) *Proxy {
	ctx := context.Background()

	c, err := elastic.NewClient(
		elastic.SetURL(config.ElasticURL.String()),
		//elastic.SetBasicAuth(config.Username, config.Password),
	)
	if err != nil {
		log.Infof("Cannot create elastic search client : %v", err)
		log.Error(err)
	}

	return &Proxy{
		config:  config,
		context: ctx,
		client:  c,
	}
}

//CreateIndex creates an elastic search index
func (p *Proxy) CreateIndex(indexName string, mapping string) error {
	exists, _ := p.client.IndexExists(indexName).Do(p.context)
	if !exists {
		result, err := p.client.CreateIndex(indexName).
			BodyString(mapping).
			Pretty(true).
			Do(p.context)

		if err != nil {
			return err
		}
		log.Infof("Acknowledged : %v Shards Acknowledged : %v", result.Acknowledged, result.ShardsAcknowledged)
	} else {
		return fmt.Errorf("Index %s already exists", indexName)
	}
	return nil
}

//DeleteIndex deletes elastic search index
func (p *Proxy) DeleteIndex(indexName string) error {
	exists, _ := p.client.IndexExists(indexName).Do(p.context)
	if exists {
		p.client.DeleteIndex(indexName).Do(p.context)
		log.Infof("Index %s successfully deleted", indexName)
	} else {
		return fmt.Errorf("Index %s doesnot exist, nothing to do", indexName)
	}
	return nil
}

//Index indexes csv data
func (p *Proxy) Index(indexName string, storesToIndex []models.StoreRecord) *models.IndexerResponse {
	var storesIndexed []models.StoreRecord
	var storesNotIndexed []models.StoreRecord

	for _, store := range storesToIndex {
		//Marshal to JSON
		storeRecord, err := json.Marshal(store)
		//index the doc
		result, err := p.client.Index().
			Index(indexName).
			Type("store").
			Id(store.StoreCode).
			BodyJson(storeRecord).
			Do(p.context)
		if err != nil {
			log.Infof("Error adding store record for store code %s to index : %v", store.StoreCode, err)
			storesNotIndexed = append(storesNotIndexed, store)
		}
		log.Infof("Added store %s to index %s", result.Id, result.Index)
		storesIndexed = append(storesIndexed, store)
	}
	return &models.IndexerResponse{
		IndexName:           indexName,
		StoresIndexed:       storesIndexed,
		StoresFailedToIndex: storesNotIndexed,
	}
}

//Search executes a store location elastic query and returns response
func (p *Proxy) Search(request models.StoreQueryRequest, indexName string) (*models.StoreQueryResponse, error) {
	var stores []models.StoreRecord

	gdq := elastic.NewGeoDistanceQuery("location").
		GeoPoint(elastic.GeoPointFromLatLon(request.Lat, request.Lon)).
		Distance(request.Radius)

	bq := elastic.NewBoolQuery().
		Must(
			elastic.NewMatchAllQuery(),
		).
		Filter(
			gdq,
		)

	results, err := p.client.Search(indexName).
		Query(bq).
		Pretty(true).
		From(0).
		Size(20).
		Do(p.context)

	if err != nil {
		log.Infof("Error executing search query %v", err)
		return nil, fmt.Errorf("Error executing search query")
	}
	if results.Hits.TotalHits > 0 {
		for _, hit := range results.Hits.Hits {
			var sr models.StoreRecord

			if err := json.Unmarshal(*hit.Source, &sr); err != nil {
				log.Infof("Error serializing hit source to StoreRecord %v", err)
				return nil, fmt.Errorf("Error serializing hit source to StoreRecord")
			}
			stores = append(stores, sr)
		}
	}

	return &models.StoreQueryResponse{
		Hits:         results.Hits.TotalHits,
		TookInMillis: results.TookInMillis,
		Stores:       stores,
	}, nil
}
