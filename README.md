# Cupcake


## Microsoft auth

https://github.com/AzureAD/microsoft-authentication-library-for-go
Before using MSAL Go, you will need to register your application with the Microsoft identity platform.

https://learn.microsoft.com/pt-br/azure/active-directory/develop/quickstart-register-app#register-an-application

Criar app

criar url de redirecionamento aplicativo "WEB"
http://localhost:3000/auth-response

criar credenciais de cliente

atualizar .env

```
AUTHORITY=https://login.microsoftonline.com/SOME_DIRECTORY_ID
CLIENT_ID=SOME_ID
CLIENT_SECRET=SOME_SECRET
REDIRECT=http://localhost:3000/auth-response
```

