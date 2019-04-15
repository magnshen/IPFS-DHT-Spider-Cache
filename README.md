# IPFS-DHT-Spider-WebsiteCache

This is Cache tool for the IPFS-DHT-Spider-Website

https://github.com/magnshen/IPFS-DHT-Spider-Website

Cache read and write in Database。Create Database Script in https://github.com/magnshen/IPFS-DHT-Spider-Server

### Build

$ go build update-cache.go

### Usage

Update the real-time data for the website to read，What  real-time data is submitted by the Spider 

update News info :   erery two seconds, automatic cycling

$ ./update-cache 

update Days info :  one time a day ,use  crond.service

$ ./update-cach -daysInfo