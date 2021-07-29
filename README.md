# tugas-acp

### Getting Started

Run this projects locally by following this step

##### Installation

- Install Go and MySql
  - download Go from [here](https://golang.org/doc/install)
  - download MySql from [here](https://dev.mysql.com/downloads/installer/)
- Clone this project
```
git clone https://github.com/ACP-10/tugas-acp.git
```

##### Run the Project

- Install the dependencies
```
cd tugas-acp
go get
```
- Run the project
```
go run main.go
```

##### Try the Project
- Go to the POSTMAN documentation [here](https://documenter.getpostman.com/view/10004593/TzscqnfK)
- Open POSTMAN and access use
```
http://localhost:8000
```

### Structure Project

ERD and list API's

##### ERD

![alt text](https://ibb.co/XjSxDBR)

List tables and the relationship

```seq
Category->Product: has many
CartItem->Product: has many
Cart->CartItem: has many
Payment->Cart: has
Customer->Cart: has
```

##### Feature of API's

- Customer
  - Register
  - Login
- Category
  - Create
  - Read
  - Update
  - Delete
- Product
  - Create
  - Read by Category
  - Update
  - Delete
- Cart Item
  - Create
  - Read
  - Update
  - Delete
- Cart
  - Create
  - Read
  - Update
  - Delete
- Payment
  - Pay

### Contribution

Really ? you want to contribute ? Thank you !!!!

- Fork please, and
- Just create new branch
```
git checkout -b [your_branch_name]
```
- Push to your branch
```
git add .
git commit -m "I help you because seems you need help"
git push origin [your_branch_name]
```
- Open pull request

### Developers

[Ikhda Muhammad Wildani](https://www.linkedin.com/in/ikhda-muhammad-wildani-b98a03164/)
[Mohammad Chico Hernando](https://www.linkedin.com/in/mohammad-chico-hernando-33b587218/)
[Your Name](https://google.com) - Come join !
