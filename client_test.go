package gohttpclient

import (
	"context"
	"testing"
)

func TestGet(t *testing.T) {
	// Your test logic here
	c := NewHttpClient(HttpClientConfig{})

	// params := map[string]any{
	// 	"features": []string{
	// 		"utm_source", "utm_medium",
	// 	},
	// }

	params := map[string]interface{}{
		"offset":   "",
		"group_id": "279425",
		"filter": map[string]interface{}{
			"gender":                   "",
			"age_from":                 "",
			"age_to":                   "",
			"lead_status":              "",
			"date_first_from":          "",
			"date_first_to":            "",
			"date_subscription_from":   "",
			"date_subscription_to":     "",
			"date_delivery_from":       "",
			"date_delivery_to":         "",
			"date_last_delivery_from":  "",
			"date_last_delivery_to":    "",
			"date_first_delivery_from": "",
			"date_first_delivery_to":   "",
			"date_away_from":           "",
			"date_away_to":             "",
			"var_eq_input":             "",
			"var_gt_input":             "",
			"var_lt_input":             "",
			"var_con_input":            "",
		},
	}

	_, err := c.Get(context.Background(), "http://localhost", params, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestPost(t *testing.T) {

	c := NewHttpClient(HttpClientConfig{})

	params := map[string]interface{}{
		"offset":   "",
		"group_id": "279425",
		"filter": map[string]interface{}{
			"gender":                   "",
			"age_from":                 "",
			"age_to":                   "",
			"lead_status":              "",
			"date_first_from":          "",
			"date_first_to":            "",
			"date_subscription_from":   "",
			"date_subscription_to":     "",
			"date_delivery_from":       "",
			"date_delivery_to":         "",
			"date_last_delivery_from":  "",
			"date_last_delivery_to":    "",
			"date_first_delivery_from": "",
			"date_first_delivery_to":   "",
			"date_away_from":           "",
			"date_away_to":             "",
			"var_eq_input":             "",
			"var_gt_input":             "",
			"var_lt_input":             "",
			"var_con_input":            "",
		},
	}

	// send url encoded request
	_, err := c.Post(context.Background(), "http://localhost", params, map[string]string{
		HeaderContentType: HeaderContentTypeValueFormUrlEncoded,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	// send json request
	_, err = c.Post(context.Background(), "http://localhost", params, map[string]string{
		HeaderContentType: HeaderContentTypeValueApplicationJSON,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	params2 := struct {
		Age   int
		Name  string
		Some  []int
		Some2 any
	}{
		Age:  32,
		Name: "igor",
		Some: []int{1, 2, 3},
		Some2: struct {
			Key   string
			Value string
		}{
			Key:   "key1",
			Value: "value1",
		},
	}

	// send url encoded request
	_, err = c.Post(context.Background(), "http://localhost", params2, map[string]string{
		HeaderContentType: HeaderContentTypeValueFormUrlEncoded,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	// send json request
	_, err = c.Post(context.Background(), "http://localhost", params2, map[string]string{
		HeaderContentType: HeaderContentTypeValueApplicationJSON,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

}
