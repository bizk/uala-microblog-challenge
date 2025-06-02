echo "1 - Creating users = [1,2,3]"

curl -X 'POST' \
  'http://localhost:8080/api/v1/users/1' \
  -H 'accept: application/json' \
  -d ''
echo 

curl -X 'POST' \
  'http://localhost:8080/api/v1/users/2' \
  -H 'accept: application/json' \
  -d ''
echo 

curl -X 'POST' \
  'http://localhost:8080/api/v1/users/3' \
  -H 'accept: application/json' \
  -d ''

echo 

echo
echo "2 - Checking it exists"

curl -X 'GET' \
  'http://localhost:8080/api/v1/users/1' \
  -H 'accept: application/json' \
  -d ''

curl -X 'GET' \
  'http://localhost:8080/api/v1/users/2' \
  -H 'accept: application/json' \
  -d ''

curl -X 'GET' \
  'http://localhost:8080/api/v1/users/3' \
  -H 'accept: application/json' \
  -d ''

echo
echo "3 - Creating tweets..."

curl -X 'POST' \
  'http://localhost:8080/api/v1/tweets' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "content": "hola soy el 1",
  "user_id": "1"
}'

curl -X 'POST' \
  'http://localhost:8080/api/v1/tweets' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "content": "hola soy el 2",
  "user_id": "2"
}'

curl -X 'POST' \
  'http://localhost:8080/api/v1/tweets' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "content": "hola soy el 3",
  "user_id": "3"
}'

curl -X 'POST' \
  'http://localhost:8080/api/v1/tweets' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "content": "Proceda 1",
  "user_id": "1"
}'

curl -X 'POST' \
  'http://localhost:8080/api/v1/tweets' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "content": "Proceda 1.2",
  "user_id": "1"
}'


curl -X 'POST' \
  'http://localhost:8080/api/v1/tweets' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "content": "Proceda 2",
  "user_id": "2"
}'


curl -X 'POST' \
  'http://localhost:8080/api/v1/tweets' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "content": "Proceda 1.3",
  "user_id": "1"
}'

echo
echo "4 - Getting 1 timeline (no follow)..."

curl -X 'GET' \
  'http://localhost:8080/api/v1/tweets/timeline?user_id=1' \
  -H 'accept: application/json'
echo


echo "5 - Getting 21timeline (follow 1)"
curl -X 'POST' \
  'http://localhost:8080/api/v1/tweets/follow' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "followed_id": "2",
  "follower_id": "1"
}'
echo 
curl -X 'GET' \
  'http://localhost:8080/api/v1/tweets/timeline?user_id=1' \
  -H 'accept: application/json'
echo
