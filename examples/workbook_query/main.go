package main

import (
	"encoding/json"
	"fmt"
	"github.com/tiketdatarisal/tableau"
	"math/rand"
	"time"
)

func main() {
	cfg := tableau.Config{
		Host:       "https://kagebunshin.kepo.ninja/",
		Version:    "3.12",
		Username:   "risal.risal@tiket.com",
		Password:   "9n230dikfsn0",
		ContentUrl: "",
	}

	client, err := tableau.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	err = client.Authentication.SignIn()
	if err != nil {
		panic(err)
	}

	workbooks, err := client.WorkbooksViews.QueryWorkbooksForSite(nil)
	if err != nil {
		panic(err)
	}

	for _, workbook := range workbooks {
		fmt.Printf("ID: %s, Name: %s\n", *workbook.ID, *workbook.Name)
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := rnd.Int63n(int64(len(workbooks) - 1))
	workbook, err := client.WorkbooksViews.QueryWorkbook(*workbooks[idx].ID)
	if err != nil {
		panic(err)
	}

	raw, err := json.MarshalIndent(workbook, "", "  ")
	fmt.Printf("Workbook:\n%s\n", string(raw))
}
