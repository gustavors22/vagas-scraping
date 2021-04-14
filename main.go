package main

import (
	"fmt"

	"github.com/gustavors22/vagas-scraping/crawler"
)

func main() {
	start()
}

func start() {
	var option int
	var url string

	fmt.Println("***************")
	fmt.Println("* 1 - Junior  *")
	fmt.Println("* 2 - Pleno   *")
	fmt.Println("* 3 - Senior  *")
	fmt.Println("* 4 - Estagio *")
	fmt.Println("***************")

	fmt.Printf("=> ")
	fmt.Scanf("%d", &option)

	switch option {
		case 1:
			url = "https://programathor.com.br/jobs?expertise=J%C3%BAnior"
		case 2:
			url = "https://programathor.com.br/jobs?expertise=Pleno"
		case 3:
			url = "https://programathor.com.br/jobs?expertise=S%C3%AAnior"
		case 4:
			url = "https://programathor.com.br/jobs?contract_type=Est%C3%A1gio"
		default:
			fmt.Println("Invalido")
	}

	crawler.Crawler(url)
}

