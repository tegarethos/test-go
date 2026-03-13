## Installation (Using Docker)
Jalankan docker
```
docker compose up --build
```
Ini akan menjalankan:
- Go API
- PostgreSQL database

Akses API
```
http://localhost:8080
```
## API Endpoints
Base URL
```
http://localhost:8080/api/v1
```
Create Todo
```
POST /todos
```
Request
```
{
  "title": "Belajar Golang",
  "description": "Clean Architecture"
}
```
Response
```
{
  "id": 1,
  "title": "Belajar Golang",
  "description": "Clean Architecture",
  "completed": false
}
```
Get All Todos
```
GET /todos
```
Response
```
[
  {
    "id": 1,
    "title": "Belajar Golang",
    "description": "Clean Architecture",
    "completed": false
  }
]
```
Get Todo By ID
```
GET /todos/{id}
```
Example
```
GET /todos/1
```
Update Todo
```
PUT /todos/{id}
```
Request
```
{
  "title": "Update Todo",
  "description": "Update description",
  "completed": true
}
```
Delete Todo
```
DELETE /todos/{id}
```
## Jalankan Container
Jalankan semua container:
```
docker compose up -d
```
Cek semua jalan
```
docker ps
```
Ambil password awal Jenkins
```
docker logs jenkins
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```
Buka browser
- Jenkins → http://localhost:8081 (admin/1570688d051a45d88eff039d77f79e04)
- SonarQube → http://localhost:9000 (admin/admin) -> admin, @Password123
## Setup Jenkins pertama kali
Buka http://localhost:8081, masukkan password dari step 3, lalu:
```
Install suggested plugins → tunggu selesai → Create Admin User → Save
```
Install tambahan plugin:
```
Dashboard → Manage Jenkins → Plugins → Available plugins
```
Cari dan install:
- Docker Pipeline
- SonarQube Scanner
Centang keduanya → Install without restart
### Tambah credentials Docker Hub:
```
Dashboard → Manage Jenkins → Credentials → System → Global credentials → Add Credentials
```
Isi:
- Kind: Username with password
- Username: username Docker Hub kamu
- Password: password Docker Hub kamu
- ID: dockerhub-credentials
- Save
### Tambah credentials SonarQube:
Dulu generate token di SonarQube dulu:
```
http://localhost:9000 → Login → 
My Account → Security → Generate Token → kasih nama → Generate → copy tokennya
```
Balik ke Jenkins:
```
Dashboard → Manage Jenkins → Credentials → Add Credentials
```
Isi:
```
Kind     : Secret text
Scope    : Global
Secret   : [paste token sonarqube]
ID       : sonar-token
Description: sonar-token
```
## Sekarang lanjut buat Pipeline:
```
Dashboard → New Item
```
Isi:
- Name: go-todo-app
- Pilih: Pipeline
- Klik OK
Nanti muncul konfigurasi, scroll ke bawah ke bagian Pipeline, isi:
- Definition: Pipeline script from SCM
- SCM: Git
- Repository URL: https://github.com/username/repo-kamu
- Branch: */main
- Script Path: Jenkinsfile