#c4-active

```sh
$   CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/c4-active
```

```sh
$   docker build --no-cache -t c4-active .
```

```sh
$   docker run -d --name c4-active -p 8080:8080 c4-active
```