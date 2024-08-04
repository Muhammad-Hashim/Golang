package main

func main() {
	// fmt.Println("Hello, Golang!")
	// mynumber :=45
	// var ptr = &mynumber
	//  fmt.Println("Address of mynumber:", *ptr)

	// var newry = []string{"Hello, Golang"}
	// newry = append(newry, "one", "two", "three")
	// fmt.Println("newry:", newry)

	// var mymap = make([]int, 2)
	// mymap[0] = 10
	// mymap[1] = 20
	// fmt.Println("mymap:", mymap[1])

	//=============================MAPS =============================

	// language := make(map[string]string)
	// language["en"] = "English"
	// language["fr"] = "Français"
	// fmt.Println("Language:", language["fr"])
	// language["es"] = "Español"
	// fmt.Println("All languages:", language)
	// delete(language, "fr")
	// fmt.Println("Languages after deletion:", language)
	// 	for Key, value := range language {
	// 	fmt.Println(Key, value)
	// }
	//=============================MAPS =============================

	//============================ switch =============================

	// fmt.Println("Switching")

	// rand.Seed(time.Now().UnixNano())

	//==========================https =============================
	// url := "https://jsonplaceholder.typicode.com/todos/10"
	// response, Error := http.Get(url)
	// if Error != nil {
	// 	fmt.Println("Error sending request:", Error)
	// 	return
	// }
	// fmt.Println("Response Status:", response.Status)
	// fmt.Println("Response Headers:", response.Header)
	// fmt.Println("Response Body:", response)

	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println(string(data))

	// response.Body.Close()
	// body, err := ioutil.ReadAll(response.Body)
	// client := &http.Client{}
	// req, err := http.NewRequest("GET", url, nil)
	// if err!= nil {
	//     fmt.Println("Error creating request:", err)
	//     return
	// }

	// res, err := client.Do(req)
	// if err!= nil {
	//     fmt.Println("Error sending request:", err)
	//     return
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err!= nil {
	//     fmt.Println("Error reading response body:", err)
	//     return
	// }

	// fmt.Println("Response Body:", string(body))

	//==========================HTTPS Request GET =================

}

func ProformanceGetRequest() {

}
