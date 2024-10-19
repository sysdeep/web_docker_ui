# hdu - Hard docker UI

проект-эксперимент по реализации управления докером с помощью хардкорного html-интерфейса

## WebUI

вариант с сервером

## TkUI

вариант desktop-UI

## Todo

- [ ] список контейнеров для volume
- [ ] список контейнеров для image

## Линки

- https://habr.com/ru/articles/449038/
- https://echo.labstack.com/docs
- https://pkg.go.dev/github.com/docker/docker/client#section-readme
- https://jenil.github.io/chota
- https://prettier.io/docs/en/options
- https://vonheikemen.github.io/devlog/tools/how-to-survive-without-multiple-cursors-in-vim/

### favicon

- https://habr.com/ru/companies/htmlacademy/articles/578224/
- https://favicon.io/

## Packages

```bash
go get github.com/docker/docker/client
go get github.com/labstack/echo/v4
```

## Docker API Version issue

Если возникнет ошибка версии api докера и клиента, можно в переменной окружения указать реальную версию

```bash
DOCKER_API_VERSION=1.43 go run ./cmd/app/
```

## GO

- Use go get -u ./... if you want to update all dependencies to the latest version in the current directory and its subdirectories.
- Use go get -t -u ./... to update all dependencies in the current directory and its subdirectories including test dependencies.
- Use go get -u all if you want to update all packages in the main module and all its dependencies including test dependencies.
- Use go get -u to update all dependencies in the current directory only.

After running the go get command, it is a good practice to also run go mod tidy. This command ensures that the go.mod file is in sync with your source code and cleans it up by removing any unused or unnecessary dependencies that may have been created after updating packages to new version.
