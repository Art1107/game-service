# Game Digital CRUD Service

บริการ HTTP API สำหรับจัดการข้อมูลสินค้าเกมดิจิทัล (Game Digital) ที่เขียนด้วยภาษา Go โดยใช้ Clean Architecture, Singleton Database Connector และฐานข้อมูล MySQL พร้อมรองรับ Unit Testing

---

## คุณสมบัติ

- จัดการสินค้าเกม: ชื่อ, ราคา, รูปภาพ
- โครงสร้างตาม Clean Architecture
- เชื่อมต่อ MySQL ด้วย Singleton Pattern
- API สร้าง อ่าน อัปเดต ลบ (CRUD)
- ทดสอบด้วย Unit Test

---

## การติดตั้งและใช้งาน

### 1. ติดตั้ง Go และ MySQL

- ตรวจสอบ Go:
  ```bash
  go version

สร้างฐานข้อมูล:

CREATE DATABASE IF NOT EXISTS game;

USE game;

CREATE TABLE name (
    name TEXT PRIMARY KEY,
    price VARCHAR(20),
    image TEXT
);

## 2. ตั้งค่า DSN
แก้ไฟล์ config/config.go:

const DSN = "root:password@tcp(127.0.0.1:3306)/game"
⚠️ เปลี่ยน root:password ให้ตรงกับเครื่องของคุณ

## 3. ติดตั้ง dependencies
go mod tidy

## 4. รันเซิร์ฟเวอร์
cd cmd/server
go run main.go
API จะพร้อมที่: http://localhost:8080

## การใช้งาน API
Method	Endpoint	Description
GET	/games	ดูเกมทั้งหมด
GET	/games/:name	ดูเกมตามชื่อ
POST	/games	เพิ่มเกม
PUT	/games/:name	แก้ไขเกม
DELETE	/games/:name	ลบเกม

## ตัวอย่าง POST (เพิ่มเกม)
POST http://localhost:8080/games

{
  "name": "Elden Ring",
  "price": "59.99",
  "image": "eldenring.jpg"
}
