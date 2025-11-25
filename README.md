# 🚀 Restfull API Golang

Project ini merupakan hasil dari belajar saya tentang Golang dan alhamdulillah sudah berhasil membuat restfull API

## ✨ Fitur

- ✅ Memakai library httprouter
- ✅ Memakai library go-viper untuk settingan .envnya
- ✅ Memakai library driver mysql untuk ke databasenya
- ✅ Memakai library validator untuk memvalidasi datanya
- ✅ Memakai library testify untuk testingnya ( saat ini menggunakan metode integration testing )
- ✅ Memakai authentication sederhana yaitu menggunakan header X-API-KEY

---

## 🎯 Tujuan

Project ini dibuat sebagai **latihan pribadi** dan **showcase** di mana untuk menunjukkan kemampuan saya bahwa saya berhasil belajar Golang sampai ke tahap membuat Restfull API.

---

## ⚙️ Prerequisites

Pastikan kamu sudah menginstall:

- Golang
- MySql client
- Extension Rest Client di IDE nya

---

## 🧩 Installation

```bash
git clone https://github.com/adjisdhani/restfull-api-go-lang.git
cd restfull-api-go-lang
go mod tidy
go run main.go
```

---

## 🧩 Installation additional

di sini menggunakan 2 database, yaitu database corenya dan database yang digunakan di dalam integration testingnya

```bash
CREATE TABLE category(
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT null
) ENGINE = INNODB;
```

pastikan juga mengisi value yang ada .env nya bisa mengcopy nanti dari .env-sample nya . untuk melihat dokumentasi dari restfull apinya bisa di cek di file **percobaan.test**

```bash
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
X_API_KEY=
```

---

## 👨‍💻 Author

**Adjis Ramadhani Utomo**
📧 [GitHub Profile](https://github.com/adjisdhani)

---

## 📜 License

Project ini memiliki lisensi open-source . Saya memperbolehkan buat siapapun yang belajar dari project ini :)
