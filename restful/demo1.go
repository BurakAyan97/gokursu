package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*https://transform.tools/json-to-go   sitesinde body atarsan go da oluşturuyo aşağıdaki yapıyı kendi*/

type Todo struct {
	UserId    int    `json:"userId` // jsondaki değrei birebir yazman lazım [JsonProperty] mantığı
	Id        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo

	json.Unmarshal(bodyBytes, &todo)

	fmt.Println(todo)
}

func Demo2() {
	todo := Todo{1, 2, "Alışveriş yapılacak", false}

	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post(
		"https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo

	json.Unmarshal(bodyBytes, &todoResponse)

	fmt.Println(todo)
}
