Zertifikate erstellen für Traing Fellow
---------------------------------------

- Root Zertifikat erstellen:

```
openssl genrsa -out tfRoot.key 2048
openssl req -x509 -new -nodes -key tfRoot.key -sha256 -days 3650 -out tfRoot.crt -config tfRoot-cert.request.conf
```

- Zertifikat für Server erstellen:

```
#CSR erstellen
openssl req -new -sha256 -nodes -out registrierung.tf.csr -newkey rsa:2048 -keyout registrierung.tf.key -config registrierung.tf.request.conf

openssl x509 -req -in registrierung.tf.csr -CA tfRoot.crt -CAkey tfRoot.key -CAcreateserial -out registrierung.tf.crt -days 3650 -sha256
```