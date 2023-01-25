# `github.com/banaio/golang/samples/reverse_proxy`

## Run

```sh
$ ( go run main.go & sleep 1 && curl http://localhost:8989 && echo && kill %1 )
2019/03/30 19:54:40 forwarding to -> httpsrs.aspsp.ob.forgerock.financial:443

<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Page Not Found</title>
    <link rel="stylesheet" href="https://cdn.forgerock.com/bootstrap/dist/css/bootstrap.min.css"/>
</head>
<body>
<main class="jumbotron rounded-0">
    <section class="container">
        <h1>404 - Page Not Found</h1>
        <p class="lead">No resource was found at / in the <strong>ForgeRock Open Banking Sandbox</strong>.
        <p class="mt-5 mb-0">
            <a href="https://www.forgerock.com"
               class="btn btn-lg btn-outline-primary rounded-0">Go to ForgeRock.com ›</a>
        </p>
    </section>
</main>
</body>
</html>
Terminated: 15
```
