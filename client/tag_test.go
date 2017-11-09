package client

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/quintilesims/layer0/common/models"
	"github.com/stretchr/testify/assert"
)

func TestListTags(t *testing.T) {
	expected := models.Tags{
		{EntityID: "eid"},
		{EntityID: "eid"},
	}

	query := url.Values{}
	query.Set(TagQueryParamType, "environment")
	query.Set(TagQueryParamID, "eid")

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.Path, "/tag")
		assert.Equal(t, query, r.URL.Query())

		MarshalAndWrite(t, w, expected, 200)
	}

	client, server := newClientAndServer(handler)
	defer server.Close()

	result, err := client.ListTags(query)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, result)
}