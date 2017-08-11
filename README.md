# protobuf-gis

[![Build Status](https://travis-ci.org/tomi77/protobuf-gis.svg?branch=master)](https://travis-ci.org/tomi77/protobuf-gis)

### Optional global installation

`.proto` files will be installed in a `/usr/incude/` folder.

    git clone https://github.com/tomi77/protobuf-gis.git
    cd protobuf-gis
    make

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
