# Golang Opinionated Version Controlled REST API for 2022

## 1. Introduction

Golang is one of the most performant language for building APIs. But, it lacks a definite project structure and 
workflows. Of course, we cannot just build a web API without a proper workflow and structure for code.

This starter is an opinionated view on how to build a REST API that's both versioned and performant at the same time. 
This isn't a standard to be followed. It's a structure that works best for our TEAM. You may or may not decide to use it 
after you read the whole documentation provided here.
___

## 2. Features

The main features of this project are as follows:

1. **Dependency Injection**: Using [Uber-FX]().
2. **Fast Web Server**: Using [GIN-GONIC]().
3. **SQL Migration Support**: Using [SQL Migrate](https://github.com/rubenv/sql-migrate).
4. **ORM**: Using [GORM]() with auto migrations disabled.
5. **Version Controlled APIs**: Using feature driven folder architecture along with MVC.

___
