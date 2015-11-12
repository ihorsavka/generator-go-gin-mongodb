# The Go-Gin-MongoDB generator

A [Yeoman](http://yeoman.io) generator for [Go (aka Golang)](https://golang.org/), [Gin-Gonic](https://gin-gonic.github.io/gin/) and [MongoDB](https://www.mongodb.org/).

The goal of this generator is to bootstrap a RESTish web app to manage Entities stored as MongoDB documents

## Current status

WIP

## Installation

Install [Git](http://git-scm.com), [node.js](http://nodejs.org), and [Go 1.4.1](http://golang.org/).

Install Yeoman:

    npm install -g yo

Install the Angular-Go-Martini generator:

    npm install -g generator-angular-go-martini

## Creating a new web app

In a new directory, generate the service:

    yo go-gin-mongodb

Get the dependencies:

    go get

Build the app:

    go build

Your service will run at [http://localhost:9000](http://localhost:9000).


## Creating a persistent entity

Generate the entity:

    yo go-gin-mongodb:entity [myentity]

You will be asked to specify attributes for the entity, where each attribute has the following:

- a name
- a type (String, Integer, Float, Boolean, Date, Enum)
- for a String attribute, an optional minimum and maximum length
- for a numeric attribute, an optional minimum and maximum value
- for a Date attribute, an optional constraint to either past values or future values
- for an Enum attribute, a list of enumerated values
- whether the attribute is required

Files that are regenerated will appear as conflicts.  Allow the generator to overwrite these files as long as no custom changes have been made.

Run the service:

    go run
