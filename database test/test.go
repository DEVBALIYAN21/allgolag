package main

import (
    "bufio"
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "godatabasetesting/database"
    "os"
)

func main() {
    var (
        password = "Dev@123"
        user     = "root"
        port     = "3306"
        dbName   = "goconsole"
    )

    connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", user, password, port, dbName)

    sqlObj, connectionError := sql.Open("mysql", connectionString)
    if connectionError != nil {
        fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
        return
    }

    data := database.Database{
        SqlDb: sqlObj,
    }

    fmt.Println("-> Welcome To Remainder App")
    fmt.Println("-> Select a numeric option; \n [1] Create a new Reminder \n [2] Get a reminder \n [3] Delete a reminder")

    consoleReader := bufio.NewScanner(os.Stdin)
    consoleReader.Scan()
    userChoice := consoleReader.Text()

    switch userChoice {
    case "1":
        var (
            titleInput, descriptionInput, aliasInput string
        )
        fmt.Println(" Please provide the following details:")

        fmt.Println("-> What is the title of your reminder?")
        consoleReader.Scan()
        titleInput = consoleReader.Text()

        fmt.Println("-> What is the description of your reminder?")
        consoleReader.Scan()
        descriptionInput = consoleReader.Text()

        fmt.Println("-> What is The NickName of your reminder? ")
        consoleReader.Scan()
        aliasInput = consoleReader.Text()

        data.CreateReminder(titleInput, descriptionInput, aliasInput)

    case "2":
        fmt.Println("-> Please provide the NickName for your reminder:")
        consoleReader.Scan()
        aliasInput := consoleReader.Text()

        data.RetrieveReminder(aliasInput)

    case "3":
        fmt.Println("-> Please provide the NickName for the reminder you want to delete:")
        consoleReader.Scan()
        deleteAlias := consoleReader.Text()

        data.DeleteReminder(deleteAlias)

    default:
        fmt.Printf("-> Option: %v is not a valid numeric option. Try 1 , 2 , 3", userChoice)
    }
}
