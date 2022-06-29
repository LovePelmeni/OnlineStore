# *Store API Docs*

Visualized API Documentation [API Documentation]("http://localhost:8000/swagger/")


# Currently Is In Development Process.


### *Project Details*
  Project is basically consists of 2 Services written in Golang (GIN-GONIC Framework) and Python (FastAPI Framework) using for communication `grpc` protobuf . Deployed with Docker-Swarm.

  If You are interested in Docker-Swarm Version of the project, Go and Check [`SWARM Documentation`]("http://github.com/LovePelmeni/Store/blob/SWARM.md") for the project, I added.

---
# *Requirements*


~ `docker` - `1.41 or higher`

~ `docker-compose` - `3.8 or higher`


---

# *Pre-Deploy Usage*


1. Clone this repo.

```editorconfig
    git clone http://github.com/LovePelmeni/Store.git

```

# *Deploy*

1. Create Network Called `store_project_network`

2. Deploy the Payment Service application as described in [Payment Service API Documentation]("http://github.com/LovePelmeni/PaymentService.git")

3. Go to the `deploy` directory of the project and simply run `Docker-Compose.yaml` file

```editorconfig
    docker-compose up -d
```
4. Test the API Using `Postman` tool or `curl`

```editorconfig
    # using curl
    $ curl -f http://localhost:8000/healthcheck/
```

~ *External Links*


~ My `LinkedIn`: Email: klimkiruk@gmail.com 
~ `Gmail` for contributions: kirklimushin@gmail.com 
