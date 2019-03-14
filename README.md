# environment
Go 1.8
MongoDB

# recommend HTTP Client
[HTTPie](https://httpie.org/)
# installation
## builds docker compose
```
$ docker-compose up -d 
```

# execution
## get Token
```
prompt$ http GET http://localhost/auth
HTTP/1.1 200 OK
Content-Length: 225
Content-Type: text/plain; charset=utf-8
Date: Thu, 14 Mar 2019 05:31:30 GMT

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTUyNjI4NzIwLCJpYXQiOiIyMDE5LTAzLTE0VDA1OjQ1OjIwLjc3NTY3MTlaIiwibmFtZSI6InRlc3R1c2VyIiwic3ViIjoiNTQ1NDY1NTczNTQifQ.CD8ILq7Uxjz7dDpUU2b9GRbsSjBSN2Lm5feSDFUgnvM
```

## create book
```
prompt$ http POST http://localhost/books "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTUyNjI4NzIwLCJpYXQiOiIyMDE5LTAzLTE0VDA1OjQ1OjIwLjc3NTY3MTlaIiwibmFtZSI6InRlc3R1c2VyIiwic3ViIjoiNTQ1NDY1NTczNTQifQ.CD8ILq7Uxjz7dDpUU2b9GRbsSjBSN2Lm5feSDFUgnvM" name=book1 price=20 difficulty=1 released:=true
HTTP/1.1 201 Created
Content-Length: 137
Content-Type: application/json
Date: Thu, 14 Mar 2019 05:46:00 GMT

{
    "CreatedAt": "2019-03-14T05:46:00.0623065Z",
    "ID": "5c89ea9861e99400286262e5",
    "difficulty": "1",
    "name": "book1",
    "price": "20",
    "released": true
}
```

## get book
### book list
```
prompt$ http GET http://localhost/books
HTTP/1.1 200 OK
Content-Length: 135
Content-Type: application/json
Date: Thu, 14 Mar 2019 05:46:24 GMT

[
    {
        "CreatedAt": "2019-03-14T05:46:00.062Z",
        "ID": "5c89ea9861e99400286262e5",
        "difficulty": "1",
        "name": "book1",
        "price": "20",
        "released": true
    }
]
```

### update book
```
prompt$ http PUT http://localhost/books/5c89ea9861e99400286262e5 "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTUyNjI4NzIwLCJpYXQiOiIyMDE5LTAzLTE0VDA1OjQ1OjIwLjc3NTY3MTlaIiwibmFtZSI6InRlc3R1c2VyIiwic3ViIjoiNTQ1NDY1NTczNTQifQ.CD8ILq7Uxjz7dDpUU2b9GRbsSjBSN2Lm5feSDFUgnvM" name=book2 price=30 difficulty=2 released:=false
HTTP/1.1 200 OK
Content-Length: 20
Content-Type: application/json
Date: Thu, 14 Mar 2019 05:47:13 GMT

{
    "result": "success"
}
```

### detail book
```
prompt$ http GET http://localhost/books/5c89ea9861e99400286262e5
HTTP/1.1 200 OK
Content-Length: 134
Content-Type: application/json
Date: Thu, 14 Mar 2019 05:47:28 GMT

{
    "CreatedAt": "2019-03-14T05:46:00.062Z",
    "ID": "5c89ea9861e99400286262e5",
    "difficulty": "2",
    "name": "book2",
    "price": "30",
    "released": false
}

```

### create review
```
prompt$ http POST http://localhost/books/5c89ea9861e99400286262e5/review name=Aron comment='That is nice book.' value=1
HTTP/1.1 201 Created
Content-Length: 169
Content-Type: application/json
Date: Thu, 14 Mar 2019 05:47:48 GMT

{
    "BookId": "5c89ea9861e99400286262e5",
    "CreatedAt": "2019-03-14T05:47:48.6452016Z",
    "comment": "That is nice book.",
    "id": "5c89eb0461e99400286262e6",
    "name": "Aron",
    "value": "1"
}
```

### delete book
```
prompt$ http DELETE http://localhost/books/5c89ea9861e99400286262e5 "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTUyNjI4NzIwLCJpYXQiOiIyMDE5LTAzLTE0VDA1OjQ1OjIwLjc3NTY3MTlaIiwibmFtZSI6InRlc3R1c2VyIiwic3ViIjoiNTQ1NDY1NTczNTQifQ.CD8ILq7Uxjz7dDpUU2b9GRbsSjBSN2Lm5feSDFUgnvM"
HTTP/1.1 200 OK
Content-Length: 20
Content-Type: application/json
Date: Thu, 14 Mar 2019 05:48:12 GMT

{
    "result": "success"
}
```

# testing
## executes unit test
```
prompt$ docker-compose run go go test
setup
PASS
teardown
ok      0.132s
```