# ORDER TRACK SERVICE 

#### *API Documentation* []("http://localhost:8000/swagger")

--- 

The Branch is the separated microservice that is responsible for order handling, coming from the store app, 
uses Apache Kafka As a Communication Layer for Accepting and Rejecting Orders Requests. 

--- 

## *Schema Preview* 




--- 

## *Requirements* 


#### ~ `docker` - `1.4.1 or higher` 

#### ~ `docker-compose` - `3.8 or higher` 


--- 

## *Usage* 

#### 1. Clone the branch from the repo.

```editorconfig
    $ git clone https://github.com/LovePelmeni/OnlineStore.git -- branch order_tracker 
```

#### 2. Go to the Root Directory and set up `proj_env.env` file according to your requirements and replace the configuration with custom params 

### Using Vim 
```editorconfig 
    $ cd -
    $ vim ./proj_env.env 

    # editing
```

#### 3. Run the Docker-Compose File in the root directory 

```
    $ docker-compose up -d
```

--- 

#### ~ External Links 

`Email for Conrtibutions` - `kirklimushin@gmail.com` & `klimkiruk@gmail.com`