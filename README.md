# GOLANG_net_interfaces

This app is written in GOLANG. It's a RESTful service, which allows you to get information about network interfaces.

# Installation

First of all you need to download a dockerimage file which consists current RESTful service:
```bash
docker search net_int_api
docker pull alexsorokin28/net_int_api
```
An image with tag latest will be uploaded on your computer. It goes without saying that you've already downloaded an appropriate version of Docker.
Now you have to build and run a container depended on this image.
```bash
docker run -d -p 40:8080 alexsorokin28/net_int_api
```
By using -p tag container port 8080 will be published to host port 40. Make sure, that container exists by using:
```bash
docker ps
```
Remember [container-id], it will be used soon.

# Usage

Our server is automatically started with containter build, so now you just need to give client app any arguments.
```bash
docker exec -it [container-id] ./client
```
Client app will print a help message if you will mistake in any of arguments. By default, server is 127.0.0.1 and port is 8080.

# Example

```bash
docker exec -it c2f807b3d96a ./client --version --server 127.0.0.1 --port 8080
docker exec -it c2f807b3d96a ./client list --server 127.0.0.1 --port 8080
```
# Help
```bash
AVAILABLE OPTIONS:
help(-h) - shows helpful information 
show(-s) [interface_name] - shows information about specified network interface
list(-l) - shows all names of all available network interfaces 
--version - shows API version of service
USAGE:
./client [command] [command_args] --server [ip_address] --port [port_value]

```
