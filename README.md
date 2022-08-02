# Webservice Palindrome

Enter two numbers separated by a space character to produce a series of numbers between 1 and 12. Then the program will check how many palindrome numbers are between the numbers 1 to 12

# Run App
```
go run main.go
```

# Run with Docker Compose
Please enter this code in your command
```
docker compose up
```

# Endpoint

Method  | Endpoint
------- | --------
POST    | http://localhost:3000/palindrome

## Example Input JSON

```
{
  "data": "2 50"
}
```
