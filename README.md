# Telkomsel Backend test API documentation

* **End Point:**
/create-product

* **Description:**
This API is used to create the product

* **Method:**
POST

* **URL Query Params:**
none

* **URL Params:**
none

* **JSON Body:**
{
    "name" : "Telkomsel 123",
    "description" : "produk telkomsel 123 adalah produk yang dibuat pada tahun ..",
    "price" : 75000,
    "variety" : "Digital",
    "rating" : 4.8,
    "stock" : 1000,
    "total_sold" : 520
}

* **Success Response:**
{
    "id": 10,
    "name": "Telkomsel 123",
    "desription": "produk telkomsel 123 adalah produk yang dibuat pada tahun ..",
    "price": 75000,
    "variety": "Digital",
    "rating": 4.8,
    "stock": 1000,
    "total_sold": 520
}

* **Error Response:**
  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "internal server error"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "failed create product. Please call customer service"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/products

* **Description:**
This API is used to get Products using pagination

* **Method:**
GET

* **URL Query Params:**
page=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
[
    {
        "id": 3,
        "name": "Telkomsel 123",
        "desription": "produk telkomsel 123 adalah produk yang dibuat pada tahun 123",
        "price": 75000,
        "variety": "Digital",
        "rating": 4.8,
        "stock": 1000,
        "total_sold": 520
    },
    {
        "id": 4,
        "name": "Telkomsel 456",
        "desription": "produk telkomsel 456 adalah produk yang dibuat pada tahun 456",
        "price": 65000,
        "variety": "Test",
        "rating": 4.6,
        "stock": 200,
        "total_sold": 100
    }
]

* **Error Response:**
  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "internal server error"
    }

* **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Failed fetch data"
    } 
------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/product

* **Description:**
This API is used to get detail product data by id

* **Method:**
GET

* **URL Query Params:**
id=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
{
    "id": 10,
    "name": "Telkomsel 123",
    "desription": "produk telkomsel 123 adalah produk yang dibuat pada tahun ..",
    "price": 75000,
    "variety": "Digital",
    "rating": 4.8,
    "stock": 1000,
    "total_sold": 520
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing product id"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "internal server error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/product

* **Description:**
This API is used to update product data. You need to pass product id query param to specify which product you want to update

* **Method:**
PUT

* **URL Query Params:**
id=...

* **URL Params:**
none

* **JSON Body:**
You can add whatever payload as long as it's property registered in product struct. The json form of the product struct is: <br/>
{
    "name" : "Telkomsel 123",
    "description" : "produk telkomsel 123 adalah produk yang dibuat pada tahun ..",
    "price" : 75000,
    "variety" : "Digital",
    "rating" : 4.8,
    "stock" : 1000,
    "total_sold" : 520
}

* **Success Response:**
{
    "code": 200,
    "message": "Product with id 3 updated"
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing product id"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Internal Server Error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/product

* **Description:**
This API is used to delete product data. You need to pass product id query param to specify which product you want to delete

* **Method:**
DELETE

* **URL Query Params:**
id=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
{
    "code": 200,
    "message": "Product with id 5 deleted"
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing product id"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Internal Server Error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------
# How to run this program?
1. create .env file, you can see the env format in env-example file
2. create docker-composer.yml file, you can see the docker-composer file in docker-compose-example file
3. run the command below in your terminal to build docker images
```
docker-compose build
```
4. if step 3 success, then run the command below
```
docker-compose up -d
```
5. if step 4 success, Congrats!! your system successfully running in local docker and you can access it
