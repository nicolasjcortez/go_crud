# heroku authorizations:create
# $env:HEROKU_API_KEY = "#####"; 
docker login --username=_ --password=$(heroku auth:token) registry.heroku.com
heroku git:remote -a gocrud
git remote -v

docker build --rm -f "Dockerfile" -t gocrud:v1 "."
docker tag gocrud:v1 registry.heroku.com/gocrud/web
docker push registry.heroku.com/gocrud/web

heroku container:release web
heroku open 
heroku logs -t

