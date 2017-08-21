# protobuf_gis

## Installation

~~~sh
go get github.com/tomi77/protobuf-gis/go
~~~


## Additional functions

All messages has function `GetGeoJSON` that returns a GeoJSON representation of a message.

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

Structure with two fields of type `float32`: `x` and `y`.

GeoJSON representation of this message is

~~~go
Point2DGeoJSON{"Point", Coordinates2D{x, y}}
~~~

#### Point3D

Structure with three fields of type `float32`: `x`, `y` and `z`.

GeoJSON representation of this message is

~~~go
Point3DGeoJSON{"Point", Coordinates3D{x, y, z}}
~~~

#### MultiPoint2D

Structure with one repeated field of type `gis_protobuf.Point2D`: `point`.

GeoJSON representation of this message is

~~~go
MultiPoint2DGeoJSON{"MultiPoint", []Coordinates2D{
  Coordinates2D{x1, y1},
  Coordinates2D{x2, y2},
}}
~~~

#### MultiPoint3D

Structure with one repeated field of type `gis_protobuf.Point3D`: `point`.

GeoJSON representation of this message is

~~~go
MultiPoint3DGeoJSON{"MultiPoint", []Coordinates3D{
  Coordinates3D{x1, y1, z1},
  Coordinates3D{x2, y2, z2},
}}
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

Structure with one repeated field of type `gis_protobuf.Point2D`: `point`.

GeoJSON representation of this message is

~~~go
LineString2DGeoJSON{"LineString", []Coordinates2D{
  Coordinates2D{x1, y1},
  Coordinates2D{x2, y2},
}}
~~~

#### LineString3D

Structure with one repeated field of type `gis_protobuf.Point3D`: `point`.

GeoJSON representation of this message is

~~~go
LineString3DGeoJSON{"LineString", []Coordinates3D{
  Coordinates3D{x1, y1, z1},
  Coordinates3D{x2, y2, z2},
}}
~~~

#### MultiLineString2D

Structure with one repeated field of type `gis_protobuf.LineString2D`: `line_string`.

GeoJSON representation of this message is

~~~go
MultiLineString2D{
  []*LineString2D{
    &LineString2D{[]*Point2D{&Point2D{x1, y1}, &Point2D{x2, y2}}},
    &LineString2D{[]*Point2D{&Point2D{x3, y3}, &Point2D{x4, y4}}},
  },
}
~~~

#### MultiLineString3D

Structure with one repeated field of type `gis_protobuf.LineString3D`: `line_string`.

GeoJSON representation of this message is

~~~go
MultiLineString3D{
  []*LineString3D{
    &LineString3D{[]*Point3D{&Point3D{x1, y1, z1}, &Point3D{x2, y2, z2}}},
    &LineString3D{[]*Point3D{&Point3D{x3, y3, z3}, &Point3D{x4, y4, z4}}},
  },
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

Structure with one repeated field of type `gis_protobuf.Point2D`: `point`.

GeoJSON representation of this message is

~~~go
Polygon2DGeoJSON{"Polygon", [][]Coordinates2D{
  []Coordinates2D{
    Coordinates2D{x1, y1},
    Coordinates2D{x2, y2},
  },
}}
~~~

#### Polygon3D

Structure with one repeated field of type `gis_protobuf.Point3D`: `point`.

GeoJSON representation of this message is

~~~go
Polygon3DGeoJSON{"Polygon", [][]Coordinates3D{
  []Coordinates3D{
    Coordinates3D{x1, y1},
    Coordinates3D{x2, y2},
  },
}}
~~~

#### MultiPolygon2D

Structure with one repeated field of type `gis_protobuf.Polygon2D`: `polygon`.

GeoJSON representation of this message is

~~~go
MultiPolygon2DGeoJSON{"MultiPolygon", [][][]Coordinates2D{
  [][]Coordinates2D{
    []Coordinates2D{
      Coordinates2D{x1, y1},
      Coordinates2D{x2, y2},
    },
  },
  [][]Coordinates2D{
    []Coordinates2D{
      Coordinates2D{x3, y3},
      Coordinates2D{x4, y4},
    },
  },
}}
~~~

#### MultiPolygon3D

Structure with one repeated field of type `gis_protobuf.Polygon3D`: `polygon`.

GeoJSON representation of this message is

~~~go
MultiPolygon3DGeoJSON{"MultiPolygon", [][][]Coordinates3D{
  [][]Coordinates3D{
    []Coordinates3D{
      Coordinates3D{x1, y1, z1},
      Coordinates3D{x2, y2, z2},
    },
  },
  [][]Coordinates3D{
    []Coordinates3D{
      Coordinates3D{x3, y3, z3},
      Coordinates3D{x4, y4, z4},
    },
  },
}}
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
protoc --go_out=src -I /usr/src/ -I . proto/test.proto
~~~

Code:

~~~go
import "github.com/tomi77/protobuf-gis/go"
import "proto"

func Test() {
  p := Point2D{1, 2}
  t := Test{&p}
  var g = p.GetGeoJSON()
}
~~~
