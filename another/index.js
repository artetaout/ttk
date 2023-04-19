const express = require('express')
const app = express()

app.get('/', (req, res) => res.send('Hello World!'))

app.get("/api/courses/:id", (req, res) => {
    res.send(`your req id = ${req.params.id}`)
})

app.get("/api/posts/:year/:month/:day", (req, res) => {
    res.send(`your req date: year=${req.params.year}, month=${req.params.month}, day=${req.params.day}`)
})


app.listen(3000)