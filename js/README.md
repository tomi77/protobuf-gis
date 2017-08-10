# protobuf-gis

## Installation

~~~sh
npm install protobuf-gis --save
~~~

or

~~~sh
yarn add protobuf-gis
~~~

## Additional functions

All messages has function `toGeoJSON` that returns a GeoJSON representation of a message.

## Messages

### Point

~~~
syntax = "proto3";

import "gis/protobuf/point.proto";

package test;

message Test {
  gis.protobuf.Point2D point2d = 1;
  gis.protobuf.Point3D point3d = 2;
  gis.protobuf.MultiPoint2D multi_point2d = 3;
  gis.protobuf.MultiPoint3D multi_point3d = 4;
}
~~~

#### Point2D

Structure with two fields of type `float`: `x` and `y`.

GeoJSON representation of this message is

~~~js
{
  'type': 'Point',
  'coordinates': [x, y]
}
~~~

#### Point3D

Structure with three fields of type `float`: `x`, `y` and `z`.

GeoJSON representation of this message is

~~~js
{
  'type': 'Point',
  'coordinates': [x, y, z]
}
~~~

#### MultiPoint2D

Structure with one repeated field of type `gis.protobuf.Point2D`: `point`.

GeoJSON representation of this message is

~~~js
{
  'type': 'MultiPoint',
  'coordinates': [[x1, y1], [x2, y2]]
}
~~~

#### MultiPoint3D

Structure with one repeated field of type `gis.protobuf.Point3D`: `point`.

GeoJSON representation of this message is

~~~js
{
  'type': 'MultiPoint',
  'coordinates': [[x1, y1, z1], [x2, y2, z2]]
}
~~~

### LineString

~~~
syntax = "proto3";

import "gis/protobuf/linestring.proto";

package test;

message Test {
  gis.protobuf.LineString2D line_string2d = 1;
  gis.protobuf.LineString3D line_string3d = 2;
  gis.protobuf.MultiLineString2D multi_line_string2d = 3;
  gis.protobuf.MultiLineString3D multi_line_string3d = 4;
}
~~~

#### LineString2D

Structure with one repeated field of type `gis.protobuf.Point2D`: `point`.

GeoJSON representation of this message is

~~~js
{
  'type': 'LineString',
  'coordinates': [[x1, y1], [x2, y2]]
}
~~~

#### LineString3D

Structure with one repeated field of type `gis.protobuf.Point3D`: `point`.

GeoJSON representation of this message is

~~~js
{
  'type': 'LineString',
  'coordinates': [[x1, y1, z1], [x2, y2, z2]]
}
~~~

#### MultiLineString2D

Structure with one repeated field of type `gis.protobuf.LineString2D`: `line_string`.

GeoJSON representation of this message is

~~~js
{
  'type': 'MultiLineString',
  'coordinates': [[[x1, y1], [x2, y2]], [[x3, y3], [x4, y4]]]
}
~~~

#### MultiLineString3D

Structure with one repeated field of type `gis.protobuf.LineString3D`: `line_string`.

GeoJSON representation of this message is

~~~js
{
  'type': 'MultiLineString',
  'coordinates': [[[x1, y1, z1], [x2, y2, z2]], [[x3, y3, z3], [x4, y4, z4]]]
}
~~~

### Polygon

~~~
syntax = "proto3";

import "gis/protobuf/point.proto";

package test;

message Test {
  gis.protobuf.Polygon2D polugon2d = 1;
  gis.protobuf.Polygon3D polugon3d = 2;
  gis.protobuf.MultiPolygon2D multi_polugon2d = 3;
  gis.protobuf.MultiPolygon3D multi_polugon3d = 4;
}
~~~

#### Polygon2D

Structure with one repeated field of type `gis.protobuf.Point2D`: `point`.

GeoJSON representation of this message is

~~~js
{
  'type': 'Polygon',
  'coordinates': [[[x1, y1], [x2, y2]]]
}
~~~

#### Polygon3D

Structure with one repeated field of type `gis.protobuf.Point3D`: `point`.

GeoJSON representation of this message is

~~~js
{
  'type': 'Polygon',
  'coordinates': [[[x1, y1, z1], [x2, y2, z2]]]
}
~~~

#### MultiPolygon2D

Structure with one repeated field of type `gis.protobuf.Polygon2D`: `polygon`.

GeoJSON representation of this message is

~~~js
{
  'type': 'MultiPolygon',
  'coordinates': [[[[x1, y1], [x2, y2]]], [[[x3, y3], [x4, y4]]]]
}
~~~

#### MultiPolygon3D

Structure with one repeated field of type `gis.protobuf.Polygon3D`: `polygon`.

GeoJSON representation of this message is

~~~js
{
  'type': 'MultiPolygon',
  'coordinates': [[[[x1, y1, z1], [x2, y2, z2]]], [[[x3, y3, z3], [x4, y4, z4]]]]
}
~~~

## Usage

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
