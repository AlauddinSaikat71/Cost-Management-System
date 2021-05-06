### Models 

**User DTO**
This model is used to keep track by which employee the cost has been transferred.

* Id : int
* Name : string
* Email : string
* Phone : string
* Password : string
* Created at : Date-Time
* Updated at : Date-Time

**Cost DTO**
This model is used to transfer a Cost object from client to server and validate the object on required criteria

* Id : int
* Title : string
* Description : string
* Amount : float/double
* Payment_Id : int
* Payment : PaymentDTO
* CreatedAt : Date-Time
* CreatedBy : int (UserId)
* UpdatedAt : Date-Time
* UpdatedBy : int (UserId)

**Payment DTO**
This model is used to keep details about the payment process.

* Id : int
* Method : string
* Amount : float/double
* CreatedAt : Date-Time
* CreatedBy : int (UserId)
* PaidBy : int (UserId)
* Meta : string (optional JSON data)

**Attachment DTO**
This model is used to keep proof of transactions.

* Id : int
* FilePath : string
* FileType : string
* CreatedAt : Date-Time
* CreatedBy : int (UserId)

**CostAttachment DTO**
This model is used to see the relation between cost and attachments. We are able to verify the cost by attachments.

* Id : int
* CostId : int
* AttachmentId : int
* CreatedAt : Date-Time
* CreatedBy : int (UserId)

**PaymentAttachment DTO**
This model is used to see the relation between payment and attachments. We are able to verify the payment by attachments.

* Id : int
* PaymentId : id
* AttachmentId : int
* CreatedAt : Date-Time
* CreatedBy : int (UserId)



***



### CRUD requests and responses



POST `/users`

request  body: 

```json
{
    "name": "Saikat",
    "email": " saikat@gamil.com",
    "phone": "01744139232",
    "password": "123456"
}
```

response : 

     ```json
{
    "data": {
        "id": 1,
        "name" : "Saikat",
        "email" : "saikat@gamil.com",
        "phone" : "01744139232",
        "password" : "123456"
        "createdat" : "2021-05-06T15:54:17.6146253+06:00",				
        "updatedat" : "2021-05-06T15:54:17.6146253+06:00"					
    }
}
     ```





GET `/users`

response 

```json
{
    "data":[
        {
          "id": 1,
          "name" : "Saikat",
          "email" : "saikat@gamil.com",
          "phone" : "01744139232",
          "password" : "123456"
          "createdat" : "2021-05-06T15:54:17.6146253+06:00",				
          "updatedat" : "2021-05-06T15:54:17.6146253+06:00"					
        }
    ]
}
```





GET `/users/:id`

response : 

```json
{
			"data": {
          "id":1,
          "name": "Saikat",
          "email": " saikat@gamil.com",
          "phone": "01744139232",
          "password": "123456"
          "createdat" : "2021-05-02T15:19:58.6942203+06:00",
          "updatedat" : "2021-05-02T15:20:23.2903907+06:00"
    	}
}
```





PATCH `/users/:id`

request body :

```json
{
    "name": "Rafiul",
    "email": "rafiul@gamil.com"
}
```

response :

```json
{
    "data": {
        "id":1,
        "name": "Rafiul",
        "email": "rafiul@gamil.com",
        "phone": "01744139232",
        "password": "123456"
        "createdat" : "2021-05-02T15:19:58.6942203+06:00",
        "updatedat" : "2021-05-02T15:20:23.2903907+06:00"
    }
}
```



DELETE  `/users/:id`

response:

```json
{
  	"data": true
}
```
