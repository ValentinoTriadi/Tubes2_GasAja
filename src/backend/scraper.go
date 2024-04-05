package main 
 
import ( 
	"encoding/csv" 
	"github.com/gocolly/colly" 
	"log" 
	"os"
	"fmt"
) 
 
// initializing a data structure to keep the scraped data 
type PokemonProduct struct { 
	url, image, name, price string 
} 
 
func main() { 
	fmt.Println("Scraping the website...")

	// initializing the slice of structs to store the data to scrape 
	var pokemonProducts []PokemonProduct 
 
	// creating a new Colly instance 
	c := colly.NewCollector() 
 
	// visiting the target page 
	c.Visit("https://scrapeme.live/shop/") 
 
	// scraping logic 
	c.OnHTML("li.product", func(e *colly.HTMLElement) { 
		pokemonProduct := PokemonProduct{} 
 
		pokemonProduct.url = e.ChildAttr("a", "href") 
		pokemonProduct.image = e.ChildAttr("img", "src") 
		pokemonProduct.name = e.ChildText("h2") 
		pokemonProduct.price = e.ChildText(".price")
		fmt.Println(pokemonProduct)
		
		pokemonProducts = append(pokemonProducts, pokemonProduct) 
	}) 
	fmt.Println(pokemonProducts)
 
	// opening the CSV file 
	file, err := os.Create("products.csv") 
	if err != nil { 
		log.Fatalln("Failed to create output CSV file", err) 
	} 
	defer file.Close() 
 
	// initializing a file writer 
	writer := csv.NewWriter(file) 
 
	// writing the CSV headers 
	headers := []string{ 
		"url", 
		"image", 
		"name", 
		"price", 
	} 
	writer.Write(headers) 
 
	// writing each Pokemon product as a CSV row 
	for _, pokemonProduct := range pokemonProducts { 
		// converting a PokemonProduct to an array of strings 
		record := []string{ 
			pokemonProduct.url, 
			pokemonProduct.image, 
			pokemonProduct.name, 
			pokemonProduct.price, 
		} 
 
		// adding a CSV record to the output file 
		writer.Write(record) 
	} 
	defer writer.Flush() 

	fmt.Println("Done...")
}
