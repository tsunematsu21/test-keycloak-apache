## Getting started
### 1. Create OpenID client by Keycloak console
1. Start container  
  `docker-compose up --build keycloak`
2. Access to console  
  http://localhost:8080
3. Create OpenID client  
  `Clients > Create client`  
  Client ID: `<YOUR CLIENT ID>`  
  Client authentication: `ON`  
  Press `Save` button
4. Copy client secret  
  `Clients > Client details > Credentials > Client secret`  

### 2. Create .env file for docker-compose

```properties
OIDC_PROVIDER_METADATA_URL=http://<YOUR HOST MACHINE IP ADDRESS>:8080/realms/master/.well-known/openid-configuration
OIDC_CLIENT_ID=<YOUR CLIENT ID>
OIDC_CLIENT_SECRET=<YOUR CLIENT SECRET>
```

### 3. Start containers

```
docker-compose up --build
```

### 4. Access web site
http://localhost/

## Links
* [Keycloak console - http://localhost:18080](http://localhost:18080)
* [Web site - http://localhost/](http://localhost/)
  * [APP 1 - http://localhost/app1](http://localhost/app1) (require auth)
  * [APP 2 - http://localhost/app2](http://localhost/app2) (not require auth)
