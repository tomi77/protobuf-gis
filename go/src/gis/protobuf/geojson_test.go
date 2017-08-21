package gis_protobuf

import (
  "testing"
  "reflect"
  "encoding/json"
)

func TestPoint2DGetGeoJSON(t *testing.T) {
  p := Point2D{1, 2}
  var g = p.GetGeoJSON()
  var e = Point2DGeoJSON{"Point", Coordinates2D{1, 2}}
  if g != e {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestPoint2DMarschaling(t *testing.T) {
  p := Point2D{1, 2}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"Point","coordinates":[1,2]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestPoint3DGetGeoJSON(t *testing.T) {
  p := Point3D{1, 2, 3}
  var g = p.GetGeoJSON()
  var e = Point3DGeoJSON{"Point", Coordinates3D{1, 2, 3}}
  if g != e {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestPoint3DMarschaling(t *testing.T) {
  p := Point3D{1, 2, 3}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"Point","coordinates":[1,2,3]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestMultiPoint2DGetGeoJSON(t *testing.T) {
  p := MultiPoint2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}}
  var g = p.GetGeoJSON()
  var e = MultiPoint2DGeoJSON{"MultiPoint", []Coordinates2D{
    Coordinates2D{1, 2},
    Coordinates2D{3, 4},
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestMultiPoint2DMarschaling(t *testing.T) {
  p := MultiPoint2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"MultiPoint","coordinates":[[1,2],[3,4]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestMultiPoint3DGetGeoJSON(t *testing.T) {
  p := MultiPoint3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}}
  var g = p.GetGeoJSON()
  var e = MultiPoint3DGeoJSON{"MultiPoint", []Coordinates3D{
    Coordinates3D{1, 2, 3},
    Coordinates3D{4, 5, 6},
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestMultiPoint3DMarschaling(t *testing.T) {
  p := MultiPoint3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"MultiPoint","coordinates":[[1,2,3],[4,5,6]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestLineString2DGetGeoJSON(t *testing.T) {
  p := LineString2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}}
  var g = p.GetGeoJSON()
  var e = LineString2DGeoJSON{"LineString", []Coordinates2D{
    Coordinates2D{1, 2},
    Coordinates2D{3, 4},
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestLineString2DMarscheling(t *testing.T) {
  p := LineString2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"LineString","coordinates":[[1,2],[3,4]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestLineString3DGetGeoJSON(t *testing.T) {
  p := LineString3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}}
  var g = p.GetGeoJSON()
  var e = LineString3DGeoJSON{"LineString", []Coordinates3D{
    Coordinates3D{1, 2, 3},
    Coordinates3D{4, 5, 6},
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestLineString3DMarschaling(t *testing.T) {
  p := LineString3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"LineString","coordinates":[[1,2,3],[4,5,6]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestMultiLineString2DGetGeoJSON(t *testing.T) {
  p := MultiLineString2D{
    []*LineString2D{
      &LineString2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}},
      &LineString2D{[]*Point2D{&Point2D{5, 6}, &Point2D{7, 8}}},
    },
  }
  var g = p.GetGeoJSON()
  var e = MultiLineString2DGeoJSON{"MultiLineString", [][]Coordinates2D{
    []Coordinates2D{Coordinates2D{1, 2}, Coordinates2D{3, 4}},
    []Coordinates2D{Coordinates2D{5, 6}, Coordinates2D{7, 8}},
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestMultiLineString2DMarschaling(t *testing.T) {
  p := MultiLineString2D{
    []*LineString2D{
      &LineString2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}},
      &LineString2D{[]*Point2D{&Point2D{5, 6}, &Point2D{7, 8}}},
    },
  }
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"MultiLineString","coordinates":[[[1,2],[3,4]],[[5,6],[7,8]]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestMultiLineString3DGetGeoJSON(t *testing.T) {
  p := MultiLineString3D{
    []*LineString3D{
      &LineString3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}},
      &LineString3D{[]*Point3D{&Point3D{7, 8, 9}, &Point3D{10, 11, 12}}},
    },
  }
  var g = p.GetGeoJSON()
  var e = MultiLineString3DGeoJSON{"MultiLineString", [][]Coordinates3D{
    []Coordinates3D{Coordinates3D{1, 2, 3}, Coordinates3D{4, 5, 6}},
    []Coordinates3D{Coordinates3D{7, 8, 9}, Coordinates3D{10, 11, 12}},
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestMultiLineString3DMarschaling(t *testing.T) {
  p := MultiLineString3D{
    []*LineString3D{
      &LineString3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}},
      &LineString3D{[]*Point3D{&Point3D{7, 8, 9}, &Point3D{10, 11, 12}}},
    },
  }
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"MultiLineString","coordinates":[[[1,2,3],[4,5,6]],[[7,8,9],[10,11,12]]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestPolygon2DGetGeoJSON(t *testing.T) {
  p := Polygon2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}}
  var g = p.GetGeoJSON()
  var e = Polygon2DGeoJSON{"Polygon", [][]Coordinates2D{
    []Coordinates2D{
      Coordinates2D{1, 2},
      Coordinates2D{3, 4},
    },
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestPolygon2DMarschaling(t *testing.T) {
  p := Polygon2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"Polygon","coordinates":[[[1,2],[3,4]]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestPolygon3DGetGeoJSON(t *testing.T) {
  p := Polygon3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}}
  var g = p.GetGeoJSON()
  var e = Polygon3DGeoJSON{"Polygon", [][]Coordinates3D{
    []Coordinates3D{
      Coordinates3D{1, 2, 3},
      Coordinates3D{4, 5, 6},
    },
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestPolygon3DMarschaling(t *testing.T) {
  p := Polygon3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"Polygon","coordinates":[[[1,2,3],[4,5,6]]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestMultiPolygon2DGetGeoJSON(t *testing.T) {
  p := MultiPolygon2D{[]*Polygon2D{
    &Polygon2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}},
    &Polygon2D{[]*Point2D{&Point2D{5, 6}, &Point2D{7, 8}}},
  }}
  var g = p.GetGeoJSON()
  var e = MultiPolygon2DGeoJSON{"MultiPolygon", [][][]Coordinates2D{
    [][]Coordinates2D{
      []Coordinates2D{
        Coordinates2D{1, 2},
        Coordinates2D{3, 4},
      },
    },
    [][]Coordinates2D{
      []Coordinates2D{
        Coordinates2D{5, 6},
        Coordinates2D{7, 8},
      },
    },
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestMultiPolygon2DMarschaling(t *testing.T) {
  p := MultiPolygon2D{[]*Polygon2D{
    &Polygon2D{[]*Point2D{&Point2D{1, 2}, &Point2D{3, 4}}},
    &Polygon2D{[]*Point2D{&Point2D{5, 6}, &Point2D{7, 8}}},
  }}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"MultiPolygon","coordinates":[[[[1,2],[3,4]]],[[[5,6],[7,8]]]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}

func TestMultiPolygon3DGetGeoJSON(t *testing.T) {
  p := MultiPolygon3D{[]*Polygon3D{
    &Polygon3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}},
    &Polygon3D{[]*Point3D{&Point3D{7, 8, 9}, &Point3D{10, 11, 12}}},
  }}
  var g = p.GetGeoJSON()
  var e = MultiPolygon3DGeoJSON{"MultiPolygon", [][][]Coordinates3D{
    [][]Coordinates3D{
      []Coordinates3D{
        Coordinates3D{1, 2, 3},
        Coordinates3D{4, 5, 6},
      },
    },
    [][]Coordinates3D{
      []Coordinates3D{
        Coordinates3D{7, 8, 9},
        Coordinates3D{10, 11, 12},
      },
    },
  }}
  if !reflect.DeepEqual(g, e) {
    t.Error("Expected ", e, ", got", g)
  }
}

func TestMultiPolygon3DMarschaling(t *testing.T) {
  p := MultiPolygon3D{[]*Polygon3D{
    &Polygon3D{[]*Point3D{&Point3D{1, 2, 3}, &Point3D{4, 5, 6}}},
    &Polygon3D{[]*Point3D{&Point3D{7, 8, 9}, &Point3D{10, 11, 12}}},
  }}
  var g = p.GetGeoJSON()
  var j, _ = json.Marshal(g)
  e := `{"type":"MultiPolygon","coordinates":[[[[1,2,3],[4,5,6]]],[[[7,8,9],[10,11,12]]]]}`
  if string(j) != e {
    t.Error("Expected ", e, ", got", string(j))
  }
}
