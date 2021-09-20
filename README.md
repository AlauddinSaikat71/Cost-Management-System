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
* Date : Date-Time
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





PUT    `/users/:id`

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





POST  `/uploads`

request body :

`key: file     value:image.png`

response:

```json
{
	"destination": "uploads/2021/05/19/27c7496b-5260-4afe-ba53-			            08a83117a8a2.PNG",
    "type": "image/png"
}
```





POST  `/attachments`

request body:

```json
{  
	"filepath" : "dvudvud/cieji",
    "filetype" : "image/jpg"
}`
```

response :

```json
  "data": {
    "ID": 1,
    "FilePath": "dvudvud/cieji",
    "FileType": "image/jpg",
    "CreatedAt": "2021-05-19T21:58:57.8858983+06:00",
    "CreatedBy": 0
  }
}
```







POST `/costs`

request body:

```json
{
    "title":"saikat",
    "description" : "no",
    "amount": 600,
    "date":"2020-05-22T18:22:00Z",
    "payment_id":5,
    "created_by": 10
}
```

response :

```json
  "data": {
    "ID": 1,
    "Title": "saikat",
    "Description": "no",
    "Amount": 600,
    "Date": "2020-05-22T00:00:00Z",
    "Payment_Id": 5,
    "Payment": {
      "ID": 0,
      "Method": "",
      "Amount": 0,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "CreatedBy": 0,
      "PaidBy": 0,
      "Meta": ""
    },
    "CreatedAt": "2021-05-19T22:07:42.779+06:00",
    "CreatedBy": 0,
    "UpdatedAt": "2021-05-19T22:07:42.779+06:00",
    "UpdatedBy": 0
  }
}
```



GET `/costs?from=2021-05-20&to=2021-05-22`

response:

``` json
{
  "data": [
      {
      "ID": 1,
      "Title": "dufduhf",
      "Description": "suhu",
      "Amount": 4000,
      "Payment_Id": 5,
      "Payment": {
        "ID": 0,
        "Method": "",
        "Amount": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "CreatedBy": 0,
        "PaidBy": 0,
        "Meta": ""
      },
      "CreatedAt": "2021-05-19T15:06:52.532947+06:00",
      "CreatedBy": 0,
      "UpdatedAt": "2021-05-19T15:06:52.532947+06:00",
      "UpdatedBy": 0
    },
    {
      "ID": 2,
      "Title": "dufduhf",
      "Description": "suhu",
      "Amount": 203,
      "Payment_Id": 5,
      "Payment": {
        "ID": 0,
        "Method": "",
        "Amount": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "CreatedBy": 0,
        "PaidBy": 0,
        "Meta": ""
      },
      "CreatedAt": "2021-05-19T15:07:09.2208497+06:00",
      "CreatedBy": 0,
      "UpdatedAt": "2021-05-19T15:07:09.2208497+06:00",
      "UpdatedBy": 0
    }
   ]
}
```





GET `/costs/:id`

response :

```json
{
  "data": {
    "ID": 1,
    "Title": "dufduhf",
    "Description": "suhu",
    "Amount": 500,
    "Payment_Id": 5,
    "Payment": {
      "ID": 0,
      "Method": "",
      "Amount": 0,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "CreatedBy": 0,
      "PaidBy": 0,
      "Meta": ""
    },
    "CreatedAt": "2021-05-19T20:22:58.9666021+06:00",
    "CreatedBy": 0,
    "UpdatedAt": "2021-05-19T20:22:58.9666021+06:00",
    "UpdatedBy": 0
  }
}
```



PUT   `/costs/:id`

request body :

```json
{
    "title":"dufduhf",
    "description" : "suhu",
    "amount": 500.5555,
    "payment_id":5
}
```

response :

```json
{
  "data": {
    "ID": 8,
    "Title": "dufduhf",
    "Description": "suhu",
    "Amount": 500.5555,
    "Payment_Id": 5,
    "Payment": {
      "ID": 0,
      "Method": "",
      "Amount": 0,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "CreatedBy": 0,
      "PaidBy": 0,
      "Meta": ""
    },
    "CreatedAt": "2021-05-19T22:07:42.779+06:00",
    "CreatedBy": 0,
    "UpdatedAt": "2021-05-22T00:53:51.4793073+06:00",
    "UpdatedBy": 0
  }
}
```



DELETE `/costs/:id`

response:

```json
{
  "data": true
}
```



POST `/costattachments`

request body:

```json
{
    "costid" :5,
    "attachmentid":5,
    "createdby": 5
}
```

response :

``` json
{
  "data": {
    "ID": 1,
    "CostId": 5,
    "Cost": {
      "ID": 0,
      "Title": "",
      "Description": "",
      "Amount": 0,
      "Payment_Id": 0,
      "Payment": {
        "ID": 0,
        "Method": "",
        "Amount": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "CreatedBy": 0,
        "PaidBy": 0,
        "Meta": ""
      },
      "CreatedAt": "0001-01-01T00:00:00Z",
      "CreatedBy": 0,
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "UpdatedBy": 0
    },
    "AttachmentId": 5,
    "Attachment": {
      "ID": 0,
      "FilePath": "",
      "FileType": "",
      "CreatedAt": "0001-01-01T00:00:00Z",
      "CreatedBy": 0
    },
    "CreatedAt": "2021-05-19T22:15:50.7998754+06:00",
    "CreatedBy": 5
  }
}
```









POST   `/payments`

request body:

```json
{
    "method":"bKash",
    "amount" : 9000.00,
    "createdby":2,
    "paidby":3
}
```

response :

```json
{
  "data": {
    "ID": 1,
    "Method": "bKash",
    "Amount": 9000,
    "CreatedAt": "2021-05-22T22:08:27.4257175+06:00",
    "CreatedBy": 2,
    "PaidBy": 3,
    "Meta": ""
  }
}
```

GET    `/payments`

response :

```json
{
  "data": [
    {
      "ID": 1,
      "Method": "bKash",
      "Amount": 5000,
      "CreatedAt": "2021-05-22T21:53:19.426631+06:00",
      "CreatedBy": 2,
      "PaidBy": 3,
      "Meta": "uuefue"
    },
    {
      "ID": 3,
      "Method": "rocket",
      "Amount": 6000,
      "CreatedAt": "2021-05-22T21:53:30.9642308+06:00",
      "CreatedBy": 2,
      "PaidBy": 3,
      "Meta": ""
    }
  ]
}
```

GET `/payments/:id`

response :

```json
{
  "data": {
    "ID": 1,
    "Method": "bKash",
    "Amount": 5000,
    "CreatedAt": "2021-05-22T21:53:19.426631+06:00",
    "CreatedBy": 2,
    "PaidBy": 3,
    "Meta": "uuefue"
  }
}
```

PUT `/payments/:id`

request body:

```json
{
    "method":"rocket",
    "amount" : 6000.00,
    "createdby":2,
    "paidby":3
}
```

response :

```json
{
  "data": {
    "ID": 1,
    "Method": "rocket",
    "Amount": 6000,
    "CreatedAt": "2021-05-22T21:53:19.426631+06:00",
    "CreatedBy": 2,
    "PaidBy": 3,
    "Meta": "uuefue"
  }
}
```

DELETE `/payments/:id`

response :

```json
{
  "data": true
}
```





