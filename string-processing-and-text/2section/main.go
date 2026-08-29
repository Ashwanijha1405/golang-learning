package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value interface{} // a.k.a `any`
	isSet bool
}

/*
%v
%+v
%#v
%T
%s
%d
%f (%.2f)
%t
%q
%%
*/

func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %s, isSet: %t", c.Key, c.Value, c.isSet)
}

func main() {
	appName := "EnvParser"
	version := 1.2
	port := 8000
	isEnabled := true

	status := fmt.Sprintf("Application: %s (Version: %.1f) running on port %d. Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)

	item1 := ConfigItem{Key: "API_URL", Value: "http://localhost:3000/api", isSet: true}
	item2 := ConfigItem{Key: "TIMEOUT_MS", Value: 5000, isSet: true}
	item3 := ConfigItem{Key: "DEBUG_MODE", Value: false, isSet: false}

	fmt.Printf("Item 1 (%%V): %v\n", item1)
	fmt.Printf("Item 2 (%%+v): %+v\n", item2)
	fmt.Printf("Item 3 (%%#V): %#v\n", item3)

	err := errors.New("test")

	fmt.Errorf("here is the error on the port %d: %w", port, err)
}
