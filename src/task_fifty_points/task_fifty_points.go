package task_fifty_points

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
)

func checkCityCompatibility(fileData []byte, filename string) {
	city := Cities{}
	typeName := reflect.TypeOf(city).Name()
	if err := proto.Unmarshal(fileData, &city); err != nil {
		printIfCompatible(typeName, false, filename)
	} else {
		deserialized := city.String()
		if !strings.Contains(deserialized, "\\x") {
			printIfCompatible(typeName, true, filename)
			fmt.Println(deserialized)
		} else {
			printIfCompatible(typeName, false, filename)
		}
	}
}

func checkNameCompatibility(fileData []byte, filename string) {
	names := Names{}
	typeName := reflect.TypeOf(names).Name()
	if err := proto.Unmarshal(fileData, &names); err != nil {
		printIfCompatible(typeName, false, filename)
	} else {
		deserialized := names.String()
		if !strings.Contains(deserialized, "\\x") {
			printIfCompatible(typeName, true, filename)
			fmt.Println(deserialized)
		} else {
			printIfCompatible(typeName, false, filename)
		}
	}
}

func checkPersonCompatibility(fileData []byte, filename string) {
	person := Person{}
	typeName := reflect.TypeOf(person).Name()
	if err := proto.Unmarshal(fileData, &person); err != nil {
		printIfCompatible(typeName, false, filename)
	} else {
		deserialized := person.String()
		if !strings.Contains(deserialized, "\\x") {
			printIfCompatible(typeName, true, filename)
			fmt.Println(deserialized)
		} else {
			printIfCompatible(typeName, false, filename)
		}
	}
}

func checkPointCompatibility(fileData []byte, filename string) {
	points := Points{}
	typeName := reflect.TypeOf(points).Name()
	if err := proto.Unmarshal(fileData, &points); err != nil {
		printIfCompatible(typeName, false, filename)
	} else {
		deserialized := points.String()
		if !strings.Contains(deserialized, "\\x") {
			printIfCompatible(typeName, true, filename)
			fmt.Println(deserialized)
		} else {
			printIfCompatible(typeName, false, filename)
		}
	}
}

func checkTeamCompatibility(fileData []byte, filename string) {
	team := Teams{}
	typeName := reflect.TypeOf(team).Name()
	if err := proto.Unmarshal(fileData, &team); err != nil {
		printIfCompatible(typeName, false, filename)
	} else {
		deserialized := team.String()
		if !strings.Contains(deserialized, "\\x") {
			printIfCompatible(typeName, true, filename)
			fmt.Println(deserialized)
		} else {
			printIfCompatible(typeName, false, filename)
		}
	}
}

func printIfCompatible(schema string, compatible bool, filename string) {
	var notification string
	if compatible {
		notification = "This file %s is compatible with %s schema\n"
	} else {
		notification = "This file %s isnt compatible with %s schema\n"
	}
	fmt.Printf(notification, filename, schema)
}

func compatibilityChallenge(filename string) {
	in, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		panic("File not found")
	}
	checkCityCompatibility(in, filename)
	checkNameCompatibility(in, filename)
	checkPersonCompatibility(in, filename)
	checkPointCompatibility(in, filename)
	checkTeamCompatibility(in, filename)

}

func FiftyPointsMain() {
	filenames, err := ioutil.ReadDir("protobuff/pb/")
	if err != nil {
		panic(err)
	}
	decodedFilenames := make([]string, len(filenames))
	for _, filename := range filenames {
		decodedFilenames = append(decodedFilenames, filename.Name())
	}
	decodedFilenames = decodedFilenames[4:]
	for _, filename := range decodedFilenames {
		fmt.Printf("\n@@@@@@@@@FILE %s @@@@@@@@@@@@\n\n\n", filename)
		fullName := fmt.Sprintf("protobuff/pb/%s", filename)
		compatibilityChallenge(fullName)
	}
}
