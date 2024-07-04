// Express Framework Setup:
const express = require('express')
const app = express()
const port = 8000

app.use(express.json());
app.use(express.urlencoded({extended: true}))

// Route Handling:
app.get('/', (req, res) => {
    res.status(200).send('Hello! this is Route /')
})

app.get('/get', (req, res) => {
    res.status(200).json({message: "You are Route /GET"})
})

app.post('/post', (req, res) => {
    let myJson = req.body;
    res.status(200).send(myJson)
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body))
})

// Listening for Requests:
app.listen(port, () => {
    console.log(`Listening on PORT::${port}`)
})