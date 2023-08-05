# Run tests using `go test` and capture the output
$testResult = Invoke-Expression "go test ./... 2>&1"

# Check if the tests were successful
if ($testResult -match "FAIL") {
    Write-Host "Tests failed, deployment aborted."
    Write-Host $testResult   # Output the test results
    exit 1
} else {
    Write-Host "Tests passed, proceeding with local execution..."
}


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

