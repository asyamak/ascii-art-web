# ASCII ART WEB STYLIZE #


to initialize project:

**~go run cmd/main.go**

or

~go run cmd/*

Purpose of project:


###### Ascii-art-web stylize is a program which consists in receiving a string as an argument in text area and outputting the string in a graphic representation using ASCII.The result of program will be in html form. ######


**allowed range of ASCII is :** 32-126 (all alpa-numeric symbols in latin alphabet + additional common signs).

## Docker: ##

Build image (You can specify a repository and tag at which to save the new image if the build succeeds): 

**~docker build -t ascii-art-web .**

Run image: 

**~docker run -d -p 8080:8080 --name ascii-web ascii-art-web**

**~docker images**

**~docker container ls**

Use this to delete everything: 

*~docker system prune -a --volumes*

To start existing but not working right now container: 

*~docker start  <container ID or name>*

To stop existing and working right now container:

*~docker stop <container ID or name>*

To show all containers created in the system:

*~docker ps -a*





The project was written by ***Asya(m.a_k)*** and ***Aray(araya)*** 

students of ***Alem school***, 6-flow students	:grinning: