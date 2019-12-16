package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

type EsClient struct {
	client *elasticsearch.Client
}


type Content struct {
	Id string `json:"id"`
	GroupId float64 `json:"group_id"`
	Name string `json:"name"`
	Answer string `json:"answer"`
	Single float64 `json:"single"`
	LinkTable string `json:"link_table"`
	Status float64 `json:"status"`
	TemplateId float64 `json:"template_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt float64 `json:"updated_at"`
	Words float64 `json:"words"`
}

func NewClient(ip string, port string) (client *EsClient, err error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + ip + ":" + port,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return
	}
	client = new(EsClient)
	client.client = es
	return
}

func (o *EsClient) GetClientInfo() (ret string, err error) {
	info, err := o.client.Info()
	if err != nil {
		fmt.Println(err)
		return
	}
	ret = info.String()
	return
}

func (o *EsClient) AddDocument(index string, ty string, data string) (err error) {
	//doc, err := json.Marshal(data)
	if err != nil {
		return
	}
	req := esapi.IndexRequest{
		Index:        index,
		DocumentType: ty,
		Body:         strings.NewReader(string(data)),
	}

	res, err := req.Do(context.Background(), o.client)
	fmt.Println(res)
	defer res.Body.Close()
	return
}

func (o *EsClient) GetDocument(indexName string, query string) {
	indexes := make([]string, 0)
	indexes = append(indexes, "test_go")
	param := `{"query":{"match":{"name":"类型类似"}},"highlight":{"fields":{"name":{}}}}`
	req := esapi.SearchRequest{
		Index:indexes,
		Body:strings.NewReader(param),
	}
	res, err := req.Do(context.Background(), o.client)

	//res, err := o.client.Search(
	//	o.client.Search.WithContext(context.Background()),
	//	o.client.Search.WithIndex("test_go"),
	//	o.client.Search.WithBody(strings.NewReader(query)),
	//	o.client.Search.WithTrackTotalHits(true),
	//	o.client.Search.WithPretty(),
	//)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func (o *EsClient) GetByField(fieldName string, fieldValue string, indexName string) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				fieldName: fieldValue,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := o.client.Search(
		o.client.Search.WithContext(context.Background()),
		o.client.Search.WithIndex(indexName),
		o.client.Search.WithBody(&buf),
		o.client.Search.WithTrackTotalHits(true),
		o.client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	var r map[string]interface{}
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	fmt.Println(r)
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
}

func (o *EsClient)GetByQuery(query string, indexName string) (contents []*Content, err error) {
	indexes := make([]string, 0)
	indexes = append(indexes, indexName)
	req := esapi.SearchRequest{
		Index:indexes,
		Body:strings.NewReader(query),
	}
	res, err := req.Do(context.Background(), o.client)
	defer res.Body.Close()
	var r map[string]interface{}
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)

		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		id := hit.(map[string]interface{})["_id"]
		source := hit.(map[string]interface{})["_source"]
		content := &Content{
			Id: id.(string),
			GroupId:source.(map[string]interface{})["group_id"].(float64),
			Name:source.(map[string]interface{})["name"].(string),
			Answer:source.(map[string]interface{})["answer"].(string),
			Single:source.(map[string]interface{})["single"].(float64),
			LinkTable:source.(map[string]interface{})["link_table"].(string),
			Status:source.(map[string]interface{})["status"].(float64),
			TemplateId:source.(map[string]interface{})["template_id"].(float64),
			CreatedAt:source.(map[string]interface{})["created_at"].(string),
			UpdatedAt:source.(map[string]interface{})["updated_at"].(float64),
			Words:source.(map[string]interface{})["words"].(float64),
		}
		contents = append(contents, content)
	}
	return
}


func (o *EsClient) GetClient() *elasticsearch.Client {
	return o.client
}

func (o *EsClient) AddIndex(index string, indexConfig interface{}) (err error) {
	config, err := json.Marshal(indexConfig)
	if err != nil {
		return
	}
	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  strings.NewReader(string(config)),
	}
	res, err := req.Do(context.Background(), o.client)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		return err
	}
	return
}

func (o *EsClient) ExistIndex(indexName string) (exist bool, err error) {
	indexes := make([]string, 0)
	indexes = append(indexes, indexName)
	req := esapi.IndicesExistsRequest{
		Index: indexes,
	}
	res, err := req.Do(context.Background(), o.client)
	if err != nil {
		return
	}
	if res.StatusCode == 200 {
		exist = true
	}
	return
}
