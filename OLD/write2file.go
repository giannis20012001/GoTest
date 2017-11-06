package OLD

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/7/2017.
 */


import (
	"encoding/json"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	"os/user"
	"reflect"

)

type LinkInfo struct {
	Title		string
	Link		string
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)

	}

	l1 := LinkInfo{"Google", "https://www.google.com/"}
	l2 := LinkInfo{"DuckDuckGo", "https://duckduckgo.com/"}
	log.Println(l1, usr.HomeDir + "/testdata")
	SaveFileAsJson(l1, usr.HomeDir + "/testdata")
	log.Println(l2, usr.HomeDir + "/testdata")
	SaveFileAsJson(l2, usr.HomeDir + "/testdata")

}

func SaveFileAsJson(l LinkInfo, dir string) {
	dirTmp := filepath.FromSlash(dir)
	if _, err := os.Stat(dirTmp); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dirTmp, 0755)

		} else {
			log.Println(err)

		}

	}

	path := dirTmp + "/linkinfo.json"
	//os.Remove(path)

	b, err := json.Marshal(l)
	if err != nil { log.Println(err) }

	ioutil.WriteFile(path, b, 0644)

}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()

	} else {
		return t.Name()

	}

}