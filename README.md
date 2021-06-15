[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/proplants/plantbook/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/proplants/plantbook/branch/main/graph/badge.svg)](https://codecov.io/gh/proplants/plantbook)
[![lint workflow](https://github.com/proplants/plantbook/actions/workflows/lint.yml/badge.svg)](https://github.com/proplants/plantbook/actions?query=workflow%3Agolangci-lint)

# Plantbook, plants social network.

![](assets/README_content/github_plants.png)

## For generate plantbook-server

nedd to be installed [go-swagger#install](https://goswagger.io/install.html)

```bash
# after change spec run command below...
swagger generate server -t internal/api -A plantbook -f ./api/swagger/swagger.yaml --main-package=../../../cmd/plantbook-server
# for help type: swagger generate server --help  
-t internal/api # path for put generated files
-A plantbook # application name
-f ./api/swagger/swagger.yaml # path to swagger speca
--main-package=../../../cmd/plantbook-server # path for put main file location
```

## Contents

- [Plantbook, plants social network.](#plantbook-plants-social-network)
  - [For generate plantbook-server](#for-generate-plantbook-server)
  - [Contents](#contents)
  - [1. Introduction](#1-introduction)
  - [2. Components](#2-components)
  - [3. Main classes](#3-main-classes)

<a id=introduction>

## 1. Introduction

Plantbook is an online service that intended for the plant owners and gardeners who look for a tool that would allow
them to take care about their plants better. Every plant can optionally have a name and eventually a blog about
its life, that can contain photos, some sad or happy stories.

<a id=components>

## 2. Components

![](assets/README_content/components.png)

<a id=main-classes>

## 3. Main classes

![](assets/README_content/classes.png)

