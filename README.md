- Aplikasi Backend menggunakan bahasa pemrograman Golang dengan framework fiber dan database mysql
- config database di file config.json
- jalankan command docker-compose up untuk membuat mysql image container 
  atau bisa buat database secara manual dengan menyesuaikan username password dan port yang ada di config.json
- eksekusi file backend.sql untuk membuat table di database backend 
- masuk ke direktori aplikasi dan jalankan command go run cmd/web/main.go

- POSTMAN Api : https://api.postman.com/collections/4131954-6e21cb0a-563a-4be7-988c-4c1d875c900a?access_key=PMAT-01HJ525KTC88QP8DC77VKQ4M9T

- API public
    - register new user
      POST    : localhost:3000/api/users 
      type    : json
      payload : {
                "password":"admin",
                "name":"admin",
                "id":"1"
                }
    - login user
      POST    : localhost:3000/api/users/login
      type    : json
      payload : {
                "username":"admin",
                "password":"admin",
                "id":"1"
            }
      response : 
                {
                "data": {
                    "token": "960c09ff-da89-4ddb-867c-40c1a359167f"
                        }
                }      

- API must login 
    - Notes   : when access api below please attach token from response login above into headers with key "Authorization"
    - POST    : localhost:3000/api/businesses/create
      type    : json
      payload : {
                "location": "Pabrik sepatu",
                "latitude": 123456789,
                "longitude": 987654321,
                "term": "Nilai Term",
                "radius": 100,
                "categories": ["nama","saya"],
                "locale": "Nilai Locale",
                "price": [100,111],
                "open_now": true,
                "open_at": 1609459200,
                "attributes":["sonia","wibisono"],
                "sort_by": "Nilai SortBy",
                "device_platform": "Nilai DevicePlatform",
                "reservation_date": "2023-12-31",
                "reservation_time": "12:00:00",
                "reservation_covers": 2,
                "matches_party_size_param": true,
                "limit": 10,
                "offset": 0
            } 
    - PUT :  localhost:3000/api/businesses/update/83de2991-b038-4e27-b505-b4f7fc79d18f ("83de2991-b038-4e27-b505-b4f7fc79d18f is UUID")
      type    : json
      payload : {
                "location": "Pabrik Lemari",
                "latitude": 123456789,
                "longitude": 987654321,
                "term": "Nilai Term",
                "radius": 100,
                "categories": ["categories1","categories2"],
                "locale": "Nilai Locale",
                "price": [100,111,222],
                "open_now": true,
                "open_at": 1609459200,
                "attributes":["attr1","attr2"],
                "sort_by": "Nilai SortBy",
                "device_platform": "Nilai DevicePlatform",
                "reservation_date": "2023-12-31",
                "reservation_time": "12:00:00",
                "reservation_covers": 2,
                "matches_party_size_param": true,
                "limit": 10,
                "offset": 0
            }     
    - DELETE localhost:3000/api/businesses/delete/29f1b0e6-47de-475a-b251-b15b35f2dea7 ("29f1b0e6-47de-475a-b251-b15b35f2dea7 is UUID")            
    - GET localhost:3000/api/businesses/list?location=location_name&term=term_value&latitude=latitude_value&longitude=longitude_value&radius=radius_value
      response : {
                    "data": [
                        {
                            "id": "83de2991-b038-4e27-b505-b4f7fc79d18f",
                            "location": "Pabrik sepatu",
                            "latitude": 123456789,
                            "longitude": 987654321,
                            "term": "Nilai Term",
                            "radius": 100,
                            "categories": "[\"value1\", \"value2\"]",
                            "locale": "Nilai Locale",
                            "price": "[100, 111, 222]",
                            "open_now": true,
                            "open_at": 1609459200,
                            "attributes": "[\"attr1\", \"attr2\"]",
                            "sort_by": "Nilai SortBy",
                            "device_platform": "Nilai DevicePlatform",
                            "reservation_date": "2023-12-31",
                            "reservation_time": "12:00:00",
                            "reservation_covers": 2,
                            "matches_party_size_param": true,
                            "limit": 10,
                            "offset": 0
                        }
                    ],
                    "paging": {
                        "page": 1,
                        "size": 10,
                        "total_item": 1,
                        "total_page": 1
                    }
                }  
