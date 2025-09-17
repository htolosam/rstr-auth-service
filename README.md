# rstr-auth-service


### Plugins gRPC y protobuf
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

micro-servicio de autenticación

🎯 Objetivo del microservicio auth

Gestionar:

Registro y login de usuarios
Emisión y renovación de tokens JWT
Validación de credenciales
Hash de contraseñas

Opcional: refresh tokens, recuperación de contraseña