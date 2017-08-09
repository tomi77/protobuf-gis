# protobuf-gis

[![Build Status](https://travis-ci.org/tomi77/protobuf-gis.svg?branch=master)](https://travis-ci.org/tomi77/protobuf-gis)

## Installation

### NodeJS

    npm install protobuf-gis --save

or

    yarn add protobuf-gis

## Additional functions

All messages has function `toGeoJSON` that returns a GeoJSON representation of a message.

## Usage

### NodeJS

In `proto/test.proto` file:

    syntax = "proto3";

    import "node_modules/protobuf-gis/gis/protobuf/point.proto";

    package test;

    message Test {
      gis.protobuf.Point2D point = 1;
    }

Build:

    protoc --js_out=import_style=commonjs,binary:. -I . proto/test.proto

Code:

    const Test = require('./proto/test_pb').Test
    const Point2D = require('protobuf-gis').Point2D

    let point = new Point2D()
    point.setX(10)
    point.setY(20)

    let test = new Test()
    test.setPoint(point)

    b = test.serializeBinary()
    console.log(b)

    test = Test.deserializeBinary(b)
    point = test.getPoint()

    console.log(point.toGeoJSON())
