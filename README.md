To run Redis locally inside a docker container, run
(start docker desktop first)
```bash
docker run --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```

Go to [localhost:6379](http://localhost:6379) to see the Redis GUI.

Our local redis instance is running on port 8001.
The GUI is running on port 6379.
