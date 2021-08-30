# JWT

## Introduction

JSON Web Token (JWT) defines a way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed. JWTs can be signed using a *secret* (with the **HMAC** algorithm) or a *public/private key pair* using **RSA** or **ECDSA**.


## When to use JWT
Here are two scenarios where JSON Web Tokens are useful:

* **Authorization**: This is the most common scenario for using JWT. Once a user logs in to your application, or authenticates in some manner, every request that is then sent from the client on behalf of the user will contain the JWT.

* **Information Exchange**: The second use for JWTs is to securely transmit information between different systems. These JWTs can be signed using public/private key pairs so you can verify each system in this transaction in a secure manner and JWTs contain an anti-tamper mechanism as they are signed based off the header and the payload.

## Server

Server creates a simple API that feature a solitary endpoint. Endpoint is protected by our ```isAuthorized``` middleware decorator. In this ```isAuthorized``` function, we check to see that the incoming request features the ```Token``` header in the request and we then subsequently check to see if the token is valid based off our private ```mySigningKey```.

If this is a valid token, we then serve the protected endpoint.

## Client

Client defines a simple API that features a single endpoint. This endpoint, when triggered, generates a new JWT using our secure ```mySigningKey```, it then creates a new http client and sets the ```Token``` header equal to the JWT string that we have just generated.

It then attempts to hit our **server** application which is running on ```http://localhost:9000``` using this signed JWT token. Our server then validates the token we’ve generated in the client and proceeds to serve us our super secret message.

# References
* [JWT.io](https://jwt.io/introduction)
* [JWT Golang Tutorial](https://tutorialedge.net/golang/authenticating-golang-rest-api-with-jwts/)
