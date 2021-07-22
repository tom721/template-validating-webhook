package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"template-validating-webhook/pkg/schemas"
)

func CheckInstanceUpdatable(w http.ResponseWriter, r *http.Request) {
	var req interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("request json decoding error")
		return
	}

	result := validate(req)

	body := schemas.ResponseBody{
		Kind:       "AdmissionReview",
		ApiVersion: "admission.k8s.io/v1beta1",
		Response: schemas.Response{
			Allowed: result,
		},
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Println("response json encoding error")
	}
}

func validate(req interface{}) bool {

	data := make(map[string]interface{})
	byteReq, _ := json.Marshal(req)
	if err := json.Unmarshal([]byte(byteReq), &data); err != nil {
		fmt.Println("request unmarshalling is failed")
		return false
	}

	request := data["request"].(map[string]interface{})
	object := request["object"].(map[string]interface{})
	newInstanceName := object["metadata"].(map[string]interface{})["name"].(string)
	newTemplateName := object["spec"].(map[string]interface{})["template"].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string)

	oldObject := request["oldObject"].(map[string]interface{})
	oldInstanceName := oldObject["metadata"].(map[string]interface{})["name"].(string)
	oldTemplateName := oldObject["spec"].(map[string]interface{})["template"].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string)

	if newInstanceName != oldInstanceName {
		return true
	}

	if newInstanceName == oldInstanceName {
		if newTemplateName == oldTemplateName {
			return true
		} else {
			return false
		}
	}

	fmt.Println("validating is failed")
	return false
}
