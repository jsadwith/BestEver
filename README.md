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
* `/api/v1/categories (POST)` - Create a new category
* `/api/v1/categories/{categoryId} (GET)` - Get a category
* `/api/v1/categories/search/{query} (GET)` - Search for a category ("all" will return all categories)
* `/api/v1/items (POST)` - Create a new item in a category
* `/api/v1/items/{itemId} (GET)` - Get an item
* `/api/v1/items/search/{categoryId}/{query} (GET)` - Search for an item in a category ("all" will return all items in that category)

API Definitions
=====

/api/v1/categories (POST)
----------
Create a new category

`http://localhost:8081/api/v1/categories`

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

/api/v1/categories/{categoryId} (GET)
----------
Get a category

`http://localhost:8081/api/v1/categories/{categoryId}?expand={resources}`

Payload
`{categoryId}` (required): The category's ID
`{resources}` (optional): Comma-separated list of related resources to expand (Items)

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

/api/v1/categories/search/{query} (GET)
----------
Search for a category ("all" will return all categories)

`http://localhost:8081/api/v1/categories/search/{query}?expand={resources}`

Payload
`{query}` (required): Category name search term ("all" will return all categories)
`{resources}` (optional): Comma-separated list of related resources to expand (Items)

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

/api/v1/items (POST)
----------
Create a new item

`http://localhost:8081/api/v1/items`

Payload
```
{
    "name": "Bose Bluetooth",
    "description": "Bluetooth headphones",
    "category_id": 4
}
```

Response
```
{
    "ID": 1,
    "CreatedAt": "2016-05-23T15:49:43.144672457-07:00",
    "UpdatedAt": "2016-05-23T15:49:43.144672457-07:00",
    "DeletedAt": null,
    "name": "Bose Bluetooth",
    "description": "Music bluetooth",
    "category_id": 2
}
```

/api/v1/items/{itemId} (GET)
----------
Get an item

`http://localhost:8081/api/v1/items/{itemId}?expand={resources}`

Payload
`{itemId}`: The item's ID
`{resources}`: Comma-separated list of related resources to expand

Response
```
{
    "ID": 1,
    "CreatedAt": "2016-05-23T15:49:43.144672457-07:00",
    "UpdatedAt": "2016-05-23T15:49:43.144672457-07:00",
    "DeletedAt": null,
    "name": "Bose Bluetooth",
    "description": "Music bluetooth",
    "category_id": 2
}
```

/api/v1/items/search/{categoryId}/{query} (GET)
----------
Search for an item in a category ("all" will return all items in that category)

`http://localhost:8081/api/v1/items/search/{categoryId}/{query}?expand={resources}`

Payload
`{categoryId}`: Category ID to search with
`{query}`: Item name search term ("all" will return all items in that category)
`{resources}`: Comma-separated list of related resources to expand

Response
```
[
    {
        "ID": 1,
        "CreatedAt": "2016-05-23T15:49:43.144672457-07:00",
        "UpdatedAt": "2016-05-23T15:49:43.144672457-07:00",
        "DeletedAt": null,
        "name": "Bose Bluetooth",
        "description": "Music bluetooth",
        "category_id": 2
    }
]
```





