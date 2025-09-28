package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    l27 "cobaltdb-local/L-27"
)

var inter *l27.Interpreter

func SetupRouter() *gin.Engine {
    var err error
    inter, err = l27.NewInterpreter("/db")
    if err != nil {
        panic("Error al iniciar el intérprete: " + err.Error())
    }

    router := gin.Default()

    // Endpoint de prueba
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    // Endpoint real
    router.POST("/execute", executeHandler)

    return router
}

type CommandRequest struct {
    Command string `json:"command"`
}

func executeHandler(c *gin.Context) {
    var req CommandRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
        return
    }

    lex := l27.NewLexer(req.Command)
    parser := l27.NewParser(lex)

    cmd, err := parser.ParseCommand()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error de parseo", "details": err.Error()})
        return
    }

    err = inter.Execute(cmd)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error ejecutando comando", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "comando ejecutado correctamente"})
}
