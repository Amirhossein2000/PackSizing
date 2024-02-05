Welcome to the packSizing API 🚀

## Deployment

Access deployment at [packSizing API](https://8pf9yz9pgz.eu-west-1.awsapprunner.com).

## API Documentation

Check API using the provided [Postman Collection](./packSizing.postman_collection.json).

## Available APIs

### 1. Calculate

Calculate the optimum combination for a given count.

```bash
curl --location 'https://8pf9yz9pgz.eu-west-1.awsapprunner.com/calc' \
--header 'Content-Type: application/json' \
--data '{
    "count": 500000
}'
```

**Response:**
```json
{
    "Answer": "(9429x53) + (7x31) + (2x23) = 500000"
}
```

### 2. Update Packs

Update the pack sizes available.

```bash
curl --location --request PUT 'https://8pf9yz9pgz.eu-west-1.awsapprunner.com/packs' \
--header 'Content-Type: application/json' \
--data '[
    53,
    31,
    23
]'
```