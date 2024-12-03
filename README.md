
# DS-Delights-Food-API

DS-Delights-Food-API เป็น API สำหรับจัดการเมนูอาหารโดยเฉพาะเจาะจงกับลูกค้าที่ต้องการควบคุมอาหารตามโรคประจำตัว โดยใช้ Go และ Gin Framework ร่วมกับ GORM สำหรับการเชื่อมต่อฐานข้อมูล SQLite

## สารบัญ
- [เริ่มต้นใช้งาน](#เริ่มต้นใช้งาน)
- [API Routes](#api-routes)
  - [POST /menus](#post-menus)
  - [GET /menus](#get-menus)
  - [POST /diseases](#post-diseases)
  - [GET /diseases](#get-diseases)
- [การตั้งค่าและการใช้งาน](#การตั้งค่าและการใช้งาน)

## เริ่มต้นใช้งาน

### 1. ติดตั้ง Dependencies
- คุณต้องติดตั้ง Go และ Gin Framework รวมถึง GORM
- ในโปรเจกต์นี้ใช้ฐานข้อมูล SQLite

```bash
go mod tidy
```

### 2. การเริ่มต้นเซิร์ฟเวอร์
```bash
go run main.go
```

เซิร์ฟเวอร์จะเริ่มที่พอร์ตที่กำหนดไว้ใน `main.go` (ปกติคือ `8080`)

## API Routes

### POST /menus
เพิ่มเมนูใหม่พร้อมกับโรคที่เกี่ยวข้อง

**Request:**
```
POST /menus
Content-Type: application/json
```

**Body:**
```json
{
  "name": "แกงส้มปลาช่อน",
  "disease_ids": [1, 3]
}
```

**Response (200 OK):**
```json
{
  "menu": {
    "ID": 10,
    "name": "แกงส้มปลาช่อน",
    "Diseases": [
      { "ID": 1, "name": "เบาหวาน", "Menus": null },
      { "ID": 3, "name": "โรคหัวใจ", "Menus": null }
    ]
  },
  "message": "Menu created successfully"
}
```

### GET /menus
ดึงข้อมูลเมนูทั้งหมดจากฐานข้อมูล

**Request:**
```
GET /menus
```

**Response (200 OK):**
```json
[
  {
    "ID": 1,
    "name": "ข้าวมันไก่",
    "Diseases": [
      { "ID": 1, "name": "เบาหวาน", "Menus": null },
      { "ID": 5, "name": "ไขมันในเลือดสูง", "Menus": null }
    ]
  },
  {
    "ID": 2,
    "name": "ยำผักรวม",
    "Diseases": [
      { "ID": 2, "name": "ความดันโลหิตสูง", "Menus": null }
    ]
  },
  ...
]
```

### POST /diseases
เพิ่มข้อมูลโรคใหม่

**Request:**
```
POST /diseases
Content-Type: application/json
```

**Body:**
```json
{
  "name": "โรคตับ"
}
```

**Response (200 OK):**
```json
{
  "disease": {
    "ID": 6,
    "name": "โรคตับ",
    "Menus": null
  },
  "message": "Disease created successfully"
}
```

### GET /diseases
ดึงข้อมูลโรคทั้งหมดจากฐานข้อมูล

**Request:**
```
GET /diseases
```

**Response (200 OK):**
```json
[
  { "ID": 1, "name": "เบาหวาน", "Menus": [...] },
  { "ID": 2, "name": "ความดันโลหิตสูง", "Menus": [...] },
  { "ID": 3, "name": "โรคหัวใจ", "Menus": [...] },
  { "ID": 4, "name": "มะเร็ง", "Menus": [...] },
  { "ID": 5, "name": "ไขมันในเลือดสูง", "Menus": [...] },
  { "ID": 6, "name": "โรคตับ", "Menus": [] }
]
```


## การตั้งค่าและการใช้งาน

### การตั้งค่าฐานข้อมูล
ฐานข้อมูล SQLite ถูกตั้งค่าให้ใช้ไฟล์ `database.db`  คุณสามารถตั้งค่าฐานข้อมูลหรือเปลี่ยนแปลงได้ในไฟล์ `database\database.go` ตามต้องการ

### การทดสอบ API ด้วย Postman
- คุณสามารถใช้ Postman หรือเครื่องมืออื่นๆ ในการทดสอบ API ตามที่ได้ระบุไว้ในส่วนของ API Routes


---
