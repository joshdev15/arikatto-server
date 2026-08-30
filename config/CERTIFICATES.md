# CERTIFICATES

The app.rsa and app.rsa.pub files that provide the private key and
public key that allow the signing and validation of the tokens by jwt, must
be housed in this folder

## Generation of the private key
_(Example for Linux)_

```bash
openssl genrsa -out app.rsa 1024
```

## Generation of the public key
_(Example for Linux)_

```bash
openssl rsa -in app.rsa -pubout > app.rsa.pub
```

The system looks for the files with the names app.rsa and app.rsa.pub
if you want to change the names you should go to the file.

and change the names of the files you want to load in the main function.

the reference path is in the variables:

 * privateCertificate
 * publicCertificate

from here it is also possible to change the location.