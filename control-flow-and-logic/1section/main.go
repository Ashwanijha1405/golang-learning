package main 

import "fmt"

func main() {

	//In go there is only way to loop that is by "for" loop

	fmt.Println("------------------ C - style loop -------------------------")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	

	fmt.Println("------------------ While Style -------------------------")
	k := 3 
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("------------------- Infinite Loop ----------------------")
	counter := 0
	for {
		fmt.Println("counter: ", counter)
		counter++
		if counter > 5 {
			break
		}
	}

	fmt.Println("------------------- Skipping ----------------------")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("------------------- Arrays ----------------------")
	items := [3]string{"Go", "Python", "Javascript"}
	for index, value := range items {
		fmt.Println(index, value)
	}
    
	fmt.Println("---------------- Variation 1 --------------------")
	for _, value := range items {
		fmt.Println(value)
	}

	fmt.Println("---------------- Variation 2 --------------------")
	for index, _ := range items {
		fmt.Println(index)
	}
}