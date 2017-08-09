# protobuf-gis

[![Build Status](https://travis-ci.org/tomi77/protobuf-gis.svg?branch=master)](https://travis-ci.org/tomi77/protobuf-gis)

## Installation

### NodeJS

~~~sh
npm install protobuf-gis --save
~~~

or

~~~sh
yarn add protobuf-gis
~~~

## Additional functions

All messages has function `toGeoJSON` that returns a GeoJSON representation of a message.

## Usage

### NodeJS

In `proto/test.proto` file:

~~~
syntax = "proto3";

import "gis/protobuf/point.proto";

package test;

message Test {
  gis.protobuf.Point2D point = 1;
}
~~~

Build:

~~~sh
protoc --js_out=import_style=commonjs,binary:. -I node_modules/protobuf-gis -I . proto/test.proto
npm explore protobuf-gis -- npm run fix:pb ../../proto/test_pb.js
~~~

`fix:pb` script fixes paths to `protobuf-gis` `*_pb.js` files. Path to `test_pb.js` is relative to `node_modules/protobuf-gis` folder.

Code:

~~~js
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
~~~
