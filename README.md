# WordCount

## Description
Given an Input text, returns the top ten most used words and their respective counts in the given string.

## Dependencies
Gin frame work is used. Gin is a HTTP web framework written in Go.
```
go get -u "gopkg.in/gin-gonic/gin.v1"
```
Stopwords library is used for removal of most frequent words (stop words) from a text content.
```
go get -u "github.com/bbalet/stopwords"
```

## API Details
* **API URL** - [http://localhost:3000/wordsList](http://localhost:3000/wordsList)
* **Method** - POST
* **Sample Input** - {"text" : "MongoDB is an open-source document database that provides high performance, high availability, and automatic scaling.A record in MongoDB is a document, which is a data structure composed of field and value pairs. MongoDB documents are similar to JSON objects. The values of fields may include other documents, arrays, and arrays of documents."}
* **Output** - {
  "result": {
    "arrays": 2,
    "automatic": 1,
    "document": 2,
    "documents": 3,
    "field": 1,
    "high": 2,
    "mongodb": 3,
    "open": 1,
    "pairs": 1,
    "performance": 1
  },
  "status": 200
}
