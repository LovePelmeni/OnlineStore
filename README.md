# *Store API Docs*

Visualized API Documentation [API Documentation]("http://localhost:8000/swagger/")


# Currently Is In Development Process.


### *Project Details*

  You can take a look at the approximate architecture for this Application, I've wrote.
  There going to be several service responsible for the managing orders and Product states. Also a couple of Services on the front end for the visualization.
<p align="center">
  <a href="*"><img src="./docs/Снимок экрана 2022-06-30 в 18.27.19.png" alt="FastAPI"></a>


---

# *Initial Requirements*


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
