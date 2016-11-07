
// 第4章 練習問題4.12
// 出来ませんでした。
package main
import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"encoding/json"
)

const (
	arg_command = iota
	arg_operation
	arg_title
	arg_issue_num
	arg_state
)

type st_xkcd struct {
	month			string		`json:"month"`
	num				int			`json:"num"`
	link			string		`json:"link"`
	year			string		`json:"year"`
	news			string		`json:"news"`
	safe_title		string		`json:"safe_title"`
	transcript		string		`json:"transcript"`
	alt				string		`json:"alt"`
	img				string		`json:"img"`
	title			string		`json:"title"`
	day				string		`json:"day"`
}

func main() {
	var data	st_xkcd

	exec.Command("curl",
		"https://xkcd.com/[1-30]/info.0.json", "-o", "./xkcd#1.json").Run()

	for i := 1; i <= 30; i++ {
		filename := "xkcd" + strconv.Itoa(i) + ".json"
		file, err := os.Open(filename)
		if err != nil {
			continue
		}
		defer file.Close()
		json.NewDecoder(file).Decode(&data)
		fmt.Printf("%+v\n", data)

		fmt.Printf("%s\n", data.transcript)



		//os.Remove(filename)
		//exec.Command("rm", "-f", filename).Run()
	}
	//os.Remove("xkcd*.json")
	//exec.Command("rm", "-f", "xkcd*.json").Run()
	/*
	for i := 1; i <= 30; i++ {
		filename := "./xkcd" + strconv.Itoa(i) + ".json"
		os.Remove(filename)
	}
	*/
	/*
	os.Remove("xkcd1.json")
	os.Remove("xkcd2.json")
	os.Remove("xkcd3.json")
	os.Remove("xkcd4.json")
	os.Remove("xkcd5.json")
	os.Remove("xkcd6.json")
	os.Remove("xkcd7.json")
	os.Remove("xkcd8.json")
	os.Remove("xkcd9.json")
	os.Remove("xkcd10.json")
	os.Remove("xkcd11.json")
	os.Remove("xkcd12.json")
	os.Remove("xkcd13.json")
	os.Remove("xkcd14.json")
	os.Remove("xkcd15.json")
	os.Remove("xkcd16.json")
	os.Remove("xkcd17.json")
	os.Remove("xkcd18.json")
	os.Remove("xkcd19.json")
	os.Remove("xkcd20.json")
	os.Remove("xkcd21.json")
	os.Remove("xkcd22.json")
	os.Remove("xkcd23.json")
	os.Remove("xkcd24.json")
	os.Remove("xkcd25.json")
	os.Remove("xkcd26.json")
	os.Remove("xkcd27.json")
	os.Remove("xkcd28.json")
	os.Remove("xkcd29.json")
	os.Remove("./xkcd30.json")
	*/

	fmt.Println("end")
}


