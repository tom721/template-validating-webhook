package apis

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {

	admissionReview := []byte(`{
		"kind": "AdmissionReview", 
		"request": { 
			"object": { 
				"spec" : { 
					"template": { 
						"metadata" : { 
							"name" : "new-template"} } } }, 
			"oldObject": { 
				"spec" : {
					"template": { 
						"metadata" : { 
							"name" : "old-template"} } } } } }`)

	var req map[string]interface{}

	if err := json.Unmarshal([]byte(admissionReview), &req); err != nil {
		fmt.Println("test request unmarshalling is failed")
	}

	request := req["request"].(map[string]interface{})
	object := request["object"].(map[string]interface{})
	newTemplateName := object["spec"].(map[string]interface{})["template"].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string)

	assert.Equal(t, "new-template", newTemplateName)
	assert.Equal(t, true, Validate(req))
}
