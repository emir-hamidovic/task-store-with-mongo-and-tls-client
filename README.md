# Client for Task server [emir.hamidovic/restserver](https://gitlab.com/emir.hamidovic/golang_rest)

Due to mTLS implementation, client and server need their keys and certs, how to create those is mentioned here [README.md](https://gitlab.com/emir.hamidovic/golang_rest/-/blob/main/README.md)

How to run this client:
`go run rest.go -cert ../restserver/localhost.pem -clientcert clientcert.pem -clientkey clientkey.pem`

Where `cert` in my case is where I have server's self-signed certificate
`clientcert` and `clientkey` are client's certificate and key, if they already exist in the directory under these names, these flags can be omitted since these are the default values anyways for these flags.
