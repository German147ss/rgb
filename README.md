# FRONTEND
cd assets/
npm start

# BACKEND
cd cmd/rgb
go run main.go


# It’s also possible to create migrations executable file:

1 cd migrations/
2 go build -o migrations *.go


And run it using:

1 ./migrations init
2 ./migrations up


