package main

var baseURL string = "https://kr.indeed.com/jobs?=q=python&limit=50"

func main() {
	pages := getPages()
}

func getPages() int {
	res, err := http.get(baseURL)
	checkErr(err)
	

}


func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}