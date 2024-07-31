import redis
import redis.exceptions
import json

def main():
  try:  
    r = redis.Redis(
        host="localhost",
        port=6379,
        password="",
        db=0
    )
    r.ping()
    print("Connected to Redis...")

    user = {
      "name": "John Doe",
      "email": "johndoe@gmail.com",
      "username": "johndoe",
    }

    err = SetData(r, "user.data", user, 60) # 60 seconds
    if err is not None:
      print("Error while saving user data to Redis")
      return

    userData = GetData(r, "user.data")
    if userData is None:
      print("Error while getting user data from Redis")
      return
    
    print(userData['email'])
  
    r.close()
  except redis.exceptions.ConnectionError as e:
    print("Error connecting to Redis server. Error: {}".format(e))
    return

def SetData(r: redis.Redis, key: str, value: dict, ex: int = 0) -> redis.exceptions.DataError:
  try:
    strVal = json.dumps(value)
    r.set(key, strVal, ex)    
  except redis.exceptions.DataError as e:
    return e

def GetData(r: redis.Redis, key: str) -> dict:
  try:
    strVal = r.get(key)
    if strVal is None:
      return None
    return json.loads(strVal)
  except redis.exceptions.DataError as e:
    return None

if __name__ == "__main__":
  main()