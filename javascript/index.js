import redis from "redis";

async function main() {
  try {
    const client = redis.createClient({
      url: "redis://localhost:6379/0",
    });

    client.on("error", (err) => console.log("Redis Client Error", err));

    await client.connect();

    const user = {
      name: "John Doe",
      email: "johndoe@gmail.com",
      username: "johndoe",
    };

    const err = await SetData(client, "user.data", user, 60);
    if (err) {
      console.log("Error setting user data to redis");
    }
    const userData = await GetData(client, "user.data");
    if (!userData) {
      console.log("Error getting user data from redis");
    }

    console.log(userData.email);

    await client.quit();
  } catch (err) {
    console.log("Error", err);
  }
}

async function SetData(client, key, value, ex) {
  try {
    await client.set(key, JSON.stringify(value), {
      EX: ex,
    });
  } catch (error) {
    return error;
  }
}

async function GetData(client, key) {
  try {
    const data = await client.get(key);
    return JSON.parse(data);
  } catch (error) {
    return;
  }
}

main();
