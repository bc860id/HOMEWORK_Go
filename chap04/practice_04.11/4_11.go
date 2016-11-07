
// 第4章 練習問題4.11
// 出来ませんでした.
package main
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"bytes"
	"time"

	"net/http"
	"encoding/json"
)

const file_create	= "issue_create.txt"
const file_update	= "issue_update.txt"
const username		= "bc860id"
const url_to_issue	= "https://api.github.com/repos/" + username + "/HOMEWORK_Go/issues"
/*
const url_to_issue	= "https://github.com/" + username + "/HOMEWORK_Go/issues"
const url_to_issue	= "https://api.github.com/search/issues"
*/
var title			string

const (
	ope_create = iota
	ope_read
	ope_update
	ope_close
)

const (
	arg_command = iota
	arg_operation
	arg_title
	arg_issue_num
	arg_state
)

type st_issues struct {
	issue				[]*st_issue
}

type st_issue struct {
	url					string		`json:"url"`
	repository_url		string		`json:"repository_url"`
	labels_url			string		`json:"labels_url"`
	comments_url		string		`json:"labels_url"`
	events_url			string		`json:"events_url"`
	html_url			string		`json:"html_url"`
	id					int			`json:"id"`
	number				int			`json:"number"`
	title				string		`json:"title"`
	user				*st_user	`json:"user"`
	labels				[]*st_label	`json:"lavels"`
	state				string		`json:"state"`
	locked				bool
	assignee			*st_user	`json:"assignee"`
	assignees			[]*st_user	`json:"assignees"`
	milestone			*st_milestone
	comments			int
	created_at			time.Time
	updated_at			time.Time
	closed_at			time.Time
	body				string		`json:"body"`
	//closed_by			*st_user
}

/*
type st_issue struct {
	id					int
	url					string
	repository_url		string
	labels_url			string
	comments_url		string
	events_url			string
	html_url			string
	number				int
	state				string
	title				string
	body				string
	user				st_user
	labels				[]st_label
	assignee			st_user
	milestone			st_milestone
	locked				bool
	comments			int
	pull_request		st_pull_request
	closed_at			time.Time
	created_at			time.Time
	updated_at			time.Time
}
*/

type st_user struct {
	login				string
	id					int
	avatar_url			string
	gravatar_id			string
	url					string
	html_url			string
	followers_url		string
	following_url		string
	gists_url			string
	starred_url			string
	subscriptions_url	string
	organizations_url	string
	repos_url			string
	events_url			string
	received_events_url	string
	user_type			string		`json:"type"`
	site_admin			bool
}

type st_label struct {
	id					int
	url					string
	name				string
	color				string
	is_default			bool		`json:"default"`
}

type st_milestone struct {
	url					string
	html_url			string
	labels_url			string
	id					int
	number				int
	state				string
	title				string
	description			string
	creator				*st_user
	open_issues			int
	closed_issues		int
	created_at			time.Time	`json:"created_at"`
	updated_at			time.Time	`json:"updated_at"`
	closed_at			time.Time	`json:"closed_at"`
	due_on				time.Time	`json:"due_on"`
}

type st_pull_request struct {
	url					string
	html_url			string
	diff_url			string
	patch_url			string
}

func main() {
	var commandname	string
	var operation	int
	var err			error
	length_arg := len(os.Args)

	fmt.Println("len:", length_arg)

	if ( length_arg < 2 ) {
		//_, commandname := filepath.Split(os.Args[0])
		commandname = filepath.Base(os.Args[0])
		fmt.Printf("usage: %s create|read|close title\n", commandname)
		fmt.Printf("       %s update title No. state(open|closed)\n", commandname)
		return
	}

	title = os.Args[arg_title]
	switch ( os.Args[arg_operation] ) {
	case "create"	:
		if ( length_arg < 3 ) {
			fmt.Printf("usage: %s create title\n", commandname)
			return
		}
		operation = ope_create
		err = createIssue(title, file_create)
		break

	case "read"		:
		operation = ope_read
		break

	case "update"	:
		if ( length_arg < 5 ) {
			fmt.Printf("usage: %s update title No. state(open|closed)\n", commandname)
			return
		}
		operation = ope_update
		err = updateIssue(title, os.Args[arg_issue_num], os.Args[arg_state], file_update)
		break

	case "close"	:
		operation = ope_close
		break

	default			:
		os.Exit(1)
	}

	fmt.Println("process end. ", operation, err)
}

func createIssue(title, filename string) (err error) {
	var buf bytes.Buffer

	file_temp := "./" + filename

	err = os.Remove(file_temp)
	//fmt.Println("create:", err)

	cmd_edit := exec.Command("vim", file_temp)
	cmd_edit.Stdin = os.Stdin
	cmd_edit.Stdout = os.Stdout
	cmd_edit.Run()

	byte_body, err := exec.Command("cat", file_temp).Output()
	buf.Write(byte_body)
	body := buf.String()
	data := "'{\"title\":" + "\"'" + title + "'\", " +
		"\"body\":" + "\"'" + body + "'\"}'"

	err_curl := exec.Command("curl", "-u", username,
		"-i", "-H", "\"Content-Type: application/json\"",
		"-X POST", "--data", data, url_to_issue).Run()

	if err_curl != nil {
		fmt.Printf("%s %s\n", body, err_curl)
		err = err_curl
	}

	return err
}

func updateIssue(title, issue_num, state, filename string) (err error) {
	//var buf bytes.Buffer
	//var result st_issues
	var result st_issue
	//var result st_user

	//file_temp := "./" + filename

	resp, err_http := http.Get(url_to_issue + "/" + issue_num)
	if err_http != nil {
		return err_http
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		os.Exit(1)
	}
	err_decode := json.NewDecoder(resp.Body).Decode(&result)
	if err_decode != nil {
		resp.Body.Close()
		return err_decode
	}

	resp.Body.Close()
	//fmt.Println(result.body)
	//fmt.Println(result)
	fmt.Printf("%+v\n", result)

	/*
	err = os.Remove(file_temp)
	//file, err_create := os.Create(file_temp)
	file, _ := os.Create(file_temp)
	defer file.Close()

	issue, err := exec.Command("curl", url_to_issue + "/" + issue_num).Output()
	fmt.Fprintf(file, "%s", issue)

	cmd := exec.Command("vim", file_temp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	*/

	return err
}


