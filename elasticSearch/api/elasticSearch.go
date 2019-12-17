package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strconv"
	"strings"
)

type EsClient struct {
	client *elasticsearch.Client
}

type Content struct {
	KeywordId  string `json:"keyword_id"`
	GroupId    string `json:"group_id"`
	Name       string `json:"name"`
	Answer     string `json:"answer"`
	Single     string `json:"single"`
	LinkTable  string `json:"link_table"`
	Status     string `json:"status"`
	TemplateId string `json:"template_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
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

func (o *EsClient) AddDocument(index string, ty string, idInt int, data string) (err error) {
	//doc, err := json.Marshal(data)
	if err != nil {
		return
	}
	id := strconv.Itoa(idInt)
	req := esapi.IndexRequest{
		Index:        index,
		DocumentID:   id,
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
		Index: indexes,
		Body:  strings.NewReader(param),
	}
	res, err := req.Do(context.Background(), o.client)

	//res, err := o.client.Search(
	//	o.client.Search.WithContext(context.Background()),
	//	o.client.Search.WithIndex("test_go"),
	//	o.client.Search.WithBody(strings.NewReader(query)),
	//	o.client.Search.WithTrackTotalHits(true),
	//	o.client.Search.WithPretty(),
	//)
	if err != nil {
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

func (o *EsClient) GetByQuery(query string, indexName string) (contents []*Content, err error) {
	indexes := make([]string, 0)
	indexes = append(indexes, indexName)
	req := esapi.SearchRequest{
		Index: indexes,
		Body:  strings.NewReader(query),
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
			KeywordId:  GetInterfaceValue(id),
			GroupId:    GetInterfaceValue(source.(map[string]interface{})["group_id"]),
			Name:       GetInterfaceValue(source.(map[string]interface{})["name"]),
			Answer:     GetInterfaceValue(source.(map[string]interface{})["answer"].(string)),
			Single:     GetInterfaceValue(source.(map[string]interface{})["single"]),
			LinkTable:  GetInterfaceValue(source.(map[string]interface{})["link_table"]),
			Status:     GetInterfaceValue(source.(map[string]interface{})["status"]),
			TemplateId: GetInterfaceValue(source.(map[string]interface{})["template_id"]),
			CreatedAt:  GetInterfaceValue(source.(map[string]interface{})["created_at"]),
			UpdatedAt:  GetInterfaceValue(source.(map[string]interface{})["updated_at"]),
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

func GetInterfaceValue(intF interface{}) (data string) {
	switch intF.(type) {
	case int:
		value := intF.(int)
		data = strconv.Itoa(value)
	case int8:
		value := intF.(int8)
		data = strconv.FormatInt(int64(value), 10)
	case int16:
		value := intF.(int16)
		data = strconv.FormatInt(int64(value), 10)
	case int32:
		value := intF.(int32)
		data = strconv.FormatInt(int64(value), 10)
	case int64:
		value := intF.(int64)
		data = strconv.FormatInt(value, 10)
	case float32:
		value := intF.(float32)
		data = strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		value := intF.(float64)
		data = strconv.FormatFloat(value, 'f', -1, 64)
	case string:
		data = intF.(string)
	}
	return
}
