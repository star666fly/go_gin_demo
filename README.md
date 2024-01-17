# 修改代理
go env -w GOPROXY=https://goproxy.cn
# 初始化项目
go mod init
# 引入依赖
go get -u github.com/gin-gonic/gin  
go get -u gorm.io/gorm  
go get -u gorm.io/driver/mysql  
go get -u github.com/swaggo/swag/cmd/swag  
go get -u github.com/swaggo/gin-swagger  
go get -u github.com/swaggo/files  
go get -u github.com/go-redis/redis  

# mysql建表：运行ddl.sql文件

# 初始化swagger
swag fmt  
swag init

# 项目swagger地址：http://localhost:8081/swagger/index.html#/

