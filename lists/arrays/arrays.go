package arrays

import "fmt"

type Product struct {
	price  float64
	weight float64
	name   string
}

func main() {
	hobbiesArray := []string{"Gym", "Swim", "Clean"}
	fmt.Println(hobbiesArray)
	fmt.Println(hobbiesArray[0])
	hobbiesSlice := hobbiesArray[1:]
	fmt.Println(hobbiesSlice)
	hobbiesSecondSlice := hobbiesArray[:2]
	fmt.Println(hobbiesSecondSlice)
	hobbiesSecondSlice = hobbiesSecondSlice[1:3]
	fmt.Println(cap(hobbiesSecondSlice))
	fmt.Println(hobbiesSecondSlice)

	goalArray := []string{}
	goalArray = append(goalArray, "Win")
	goalArray = append(goalArray, "Lose")
	goalArray[1] = "Win"
	goalArray = append(goalArray, "Win")
	fmt.Println(goalArray)

	productsArray := []Product{}
	productsArray = append(productsArray, Product{
		price:  10.99,
		name:   "Water",
		weight: 1.5,
	})
	fmt.Println(productsArray)

}
