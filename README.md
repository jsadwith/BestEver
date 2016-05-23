# BestEver Prototype
Compiles best things ever

Setup
=====
1. Run `go get` to retrieve packages
2. Run `go install` to compile
3. Run `go build && ./BestEver` to build and run (`go run .main.co` also works)
4. Access at `http://localhost:8081`

Usage
=====
* `/api/category (POST)` - Create a new category
* `/api/category/{categoryId} (GET)` - Get a category
* `/api/category/search/{query} (GET)` - Search for a category ("all" will return all categories)

API Definitions
=====

/api/category (POST)
----------
Create a new category

`http://localhost:8081/api/category`

Payload
```
{
    "name": "Headphones",
    "description": "Music ears"
}
```

Response
```
{
    "ID": 4,
    "CreatedAt": "2016-05-23T15:11:21.833486307-07:00",
    "UpdatedAt": "2016-05-23T15:11:21.833486307-07:00",
    "DeletedAt": null,
    "name": "Headphones",
    "description": "Music ears",
}
```

/api/category/{categoryId} (GET)
----------
Get a category

`http://localhost:8081/api/category/{categoryId}`

Payload
`{categoryId}`: The category's ID

Response
```
{
    "ID": 4,
    "CreatedAt": "2016-05-23T15:11:21.833486307-07:00",
    "UpdatedAt": "2016-05-23T15:11:21.833486307-07:00",
    "DeletedAt": null,
    "name": "Headphones",
    "description": "Music ears",
}
```

/api/category/search/{query} (GET)
----------
Search for a category ("all" will return all categories)

`http://localhost:8081/api/category/search/{query}`

Payload
`{query}`: Category name search term ("all" will return all categories)

Response
```
[
    {
        "ID": 4,
        "CreatedAt": "2016-05-23T15:11:21.833486307-07:00",
        "UpdatedAt": "2016-05-23T15:11:21.833486307-07:00",
        "DeletedAt": null,
        "name": "Headphones",
        "description": "Music ears",
    }
]
```



