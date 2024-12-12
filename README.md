# adventofdocker
Based on - https://adventofdocker.com/

The commands:-

Build: 
docker build -t hello-world-go .

Start:
docker run -p 8080:8080 hello-world-go

Start with Persistance:
docker run -p 8080:8080 -v mydata:/go/data hello-world-go