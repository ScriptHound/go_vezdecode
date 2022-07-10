package task_fifty_points

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"google.golang.org/protobuf/proto"
)

func checkCityCompatibility(fileData []byte) {
	city := Cities{}
	if err := proto.Unmarshal(fileData, &city); err != nil {
		noticement := "This file isnt compatible with %s schema\n"
		typeName := reflect.TypeOf(city).Name()
		fmt.Printf(noticement, typeName)
	} else {
		noticement := "This file is compatible with %s schema\n"
		typeName := reflect.TypeOf(city).Name()
		fmt.Printf(noticement, typeName)
		fmt.Println(city.String())
	}

}

func checkNameCompatibility(fileData []byte) {
	name := Names{}
	if err := proto.Unmarshal(fileData, &name); err != nil {
		noticement := "This file isnt compatible with %s schema\n"
		typeName := reflect.TypeOf(name).Name()
		fmt.Printf(noticement, typeName)
	} else {
		noticement := "This file is compatible with %s schema\n"
		typeName := reflect.TypeOf(name).Name()
		fmt.Printf(noticement, typeName)
		fmt.Println(name.String())
	}
}

func checkPersonCompatibility(fileData []byte) {
	person := Person{}
	if err := proto.Unmarshal(fileData, &person); err != nil {
		noticement := "This file isnt compatible with %s schema\n"
		typeName := reflect.TypeOf(person).Name()
		fmt.Printf(noticement, typeName)
	} else {
		noticement := "This file is compatible with %s schema\n"
		typeName := reflect.TypeOf(person).Name()
		fmt.Printf(noticement, typeName)
		fmt.Println(person.String())
	}
}

func checkPointCompatibility(fileData []byte) {
	point := Points{}
	if err := proto.Unmarshal(fileData, &point); err != nil {
		noticement := "This file isnt compatible with %s schema\n"
		typeName := reflect.TypeOf(point).Name()
		fmt.Printf(noticement, typeName)
	} else {
		noticement := "This file is compatible with %s schema\n"
		typeName := reflect.TypeOf(point).Name()
		fmt.Printf(noticement, typeName)
		fmt.Println(point.String())
	}
}

func checkTeamCompatibility(fileData []byte) {
	team := Teams{}
	if err := proto.Unmarshal(fileData, &team); err != nil {
		noticement := "This file isnt compatible with %s schema\n"
		typeName := reflect.TypeOf(team).Name()
		fmt.Printf(noticement, typeName)
	} else {
		noticement := "This file is compatible with %s schema\n"
		typeName := reflect.TypeOf(team).Name()
		fmt.Printf(noticement, typeName)
		fmt.Println(team.String())
	}
}

func compatibilityChallenge(filename string) {
	in, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		panic("File not found")
	}
	checkCityCompatibility(in)
	checkNameCompatibility(in)
	checkPersonCompatibility(in)
	checkPointCompatibility(in)
	checkTeamCompatibility(in)

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
	// compatibilityChallenge("protobuff/pb/example1.pb")
}
