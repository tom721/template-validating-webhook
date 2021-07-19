package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"template-validating-webhook/pkg/schemas"
)

func CheckInstanceUpdatable(w http.ResponseWriter, r *http.Request) {
	var request interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("request json decoding error")
		return
	}

	result := validate(request)

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

func validate(request interface{}) bool {
	return true
}
