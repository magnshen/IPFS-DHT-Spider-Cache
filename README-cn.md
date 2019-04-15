# IPFS-DHT-Spider-网站缓存更新工具

这是IPFS-DHT-Spider-Website中接口使用的缓存的更新工具。缓存的存取在数据库中，建表脚本在https://github.com/magnshen/IPFS-DHT-Spider-Server

### 编译

$ go build update-cache.go

### 用法

缓存中的最新数据每2秒更新1次，工具已实现自动循环。每日数据每日更新1次或多次，请使用主机的定时任务，例如 crontab

更新最新数据：

$ ./update-cache 

更新每日数据：

$ ./update-cach -daysInfo



