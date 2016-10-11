
// 第4章 練習問題4.10
package main
import (
	"fmt"
	"log"
	"os"
	"time"
	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	/*----------------------------------------------------------*/
	/* 現在を取得。												*/
	/*----------------------------------------------------------*/
	cur := time.Now().UTC()

	/*----------------------------------------------------------*/
	/* 現在から1ヶ月前を取得。									*/
	/*----------------------------------------------------------*/
	ago_1month := cur.AddDate(0, -1, 0)

	/*----------------------------------------------------------*/
	/* 現在から1年前を取得。									*/
	/*----------------------------------------------------------*/
	ago_1year := cur.AddDate(-1, 0, 0)

	fmt.Printf("\n----- [within 1 month] ------\n")
	for _, item := range result.Items {
		if ( item.CreatedAt.Sub(ago_1month) > 0 ) {
			fmt.Printf("%v #%-5d %9.9s %.128s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Printf("\n----- [within 1 year ] ------\n")
	for _, item := range result.Items {
		if ( (item.CreatedAt.Sub(ago_1month) <= 0) &&
			 (item.CreatedAt.Sub(ago_1year) > 0) ) {
			fmt.Printf("%v #%-5d %9.9s %.128s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Printf("\n----- [before 1 year ] ------\n")
	for _, item := range result.Items {
		if ( item.CreatedAt.Sub(ago_1year) <= 0) {
			fmt.Printf("%v #%-5d %9.9s %.128s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Printf("\n")
}

