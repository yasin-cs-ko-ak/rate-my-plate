// Restaurent Log File analysis

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type FoodMenu struct {
	FoodMenuID string `json:"foodmenu_id"`
	EaterID    string `json:"eater_id"`
}

type KeyValueMap struct {
	Key   string
	Value int
}

func TopThree() error {
	// Open our jsonFile
	jsonFile, err := os.Open("log.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened log.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var foodmenus []FoodMenu
	var eater_id []FoodMenu

	// we unmarshal our byteArray which contains
	json.Unmarshal(byteValue, &foodmenus)
	json.Unmarshal(byteValue, &eater_id)

	// create a map to store the count of each foodmenu_id and eater_id
	foodmenuMap := make(map[string]int)
	eater_idMap := make(map[string]int)

	// iterate through the foodmenus array and store the count of each foodmenu_id
	for i := 0; i < len(foodmenus); i++ {
		foodmenuMap[foodmenus[i].FoodMenuID]++
	}

	// iterate through the eater_id array and store the count of each eater_id
	for i := 0; i < len(eater_id); i++ {
		eater_idMap[eater_id[i].EaterID]++
	}
	sortedMap := sortMap(foodmenuMap)

	// iterate through the foodmenuMap and print the foodmenu_id and its count
	fmt.Println("----Top 3 FoodMenuIDs----")
	for i, kv := range sortedMap {
		if i == 3 {
			break
		}
		fmt.Println("FoodMenuIDs: " + kv.Key + " Count: " + fmt.Sprint(kv.Value))

	}
	var idArray []string
	for _, tempData := range foodmenus {
		idArray = append(idArray, tempData.EaterID)
	}
	unique(foodmenus)

	return nil
}

// sorting foodmenu_id
func sortMap(foodmenuMap map[string]int) []KeyValueMap {

	var ss []KeyValueMap
	for k, v := range foodmenuMap {
		ss = append(ss, KeyValueMap{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

// checking if same eater_id has same foodmenu_id more than once
func unique(s []FoodMenu) []FoodMenu {
	inResult := make(map[FoodMenu]bool)
	var result []FoodMenu
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		} else {
			fmt.Println("Error EaterID: " + str.EaterID + " has FoodMenuID: " + str.FoodMenuID + " more than once")
		}
	}
	return result
}

func main() {
	err := TopThree()
	if err != nil {
		fmt.Println(err)
	}

}
