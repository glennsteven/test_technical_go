## Running on your local machine

Linux or MacOS

## Installation guide
#### 1. install go version 1.17
```bash
# please read this link installation guide of go
# https://go.dev/doc/install
```

#### 2. Create directory workspace    
```bash
run command below: 
mkdir $HOME/go
mkdir $HOME/go/src
mkdir $HOME/go/pkg
mkdir $HOME/go/bin
mkdir -p $HOME/go/src/technical_test_go/technical_test_go
chmod -R 775 $HOME/go
cd $HOME/go/src/technical_test_go/technical_test_go
export GOPATH=$HOME/go
```    
```bash
# edit bash profile in $HOME/.bash_profile        
# add below to new line in file .bash_profile         
    PATH=$PATH:$HOME/bin:$HOME/go/bin
    export PATH  
    export GOPATH=$HOME/go 
# run command :
source $HOME/.bash_profile
```

#### 3. Build the application    
```bash
# run command :
clone repository , git clone "url repository"
duplicate file app.yaml.example to app.yaml
make clean
make install
make start
```


### 4. Health check Route PATH
```bash
{{host}}/checklife
```


#### Postman Collection
```go
```

### Database Migration
migration up
```bash
go run main.go db:migrate up
```

migration down
```bash
go run main.go db:migrate down
```

migration reset
```bash
go run main.go db:migrate reset
```

migration reset
```bash
go run main.go db:migrate reset
```

migration redo
```bash
go run main.go db:migrate redo
```

migration status
```bash
go run main.go db:migrate status
```

create migration table
```bash
go run main.go db:migrate create {table-name} sql

# example
go run main.go db:migrate create users sql
```

to show all command
```bash
go run main.go db:migrate
```

#### generate entity only
make sure already have database and table
```bash
go run main.go gen:entity {tableName}

# example
go run main.go gen:entity users
```
