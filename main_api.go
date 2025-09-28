package main

import (
    "cobaltdb-local/api"
)

func main() {
    router := api.SetupRouter()
    router.Run(":8080")
}
