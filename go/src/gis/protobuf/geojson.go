package gis_protobuf

type Coordinates2D [2]float32
type Coordinates3D [3]float32

// Point2D

type Point2DGeoJSON struct {
  Type string               `json:"type"`
  Coordinates Coordinates2D `json:"coordinates"`
}

func (m* Point2D) GetGeoJSON() Point2DGeoJSON {
  return Point2DGeoJSON{"Point", Coordinates2D{m.GetX(), m.GetY()}}
}

// Point3D

type Point3DGeoJSON struct {
  Type string               `json:"type"`
  Coordinates Coordinates3D `json:"coordinates"`
}

func (m* Point3D) GetGeoJSON() Point3DGeoJSON {
  return Point3DGeoJSON{"Point", Coordinates3D{m.GetX(), m.GetY(), m.GetZ()}}
}

// MultiPoint2D

type MultiPoint2DGeoJSON struct {
  Type string                 `json:"type"`
  Coordinates []Coordinates2D `json:"coordinates"`
}

func (m* MultiPoint2D) GetGeoJSON() MultiPoint2DGeoJSON {
  var points = m.GetPoint()
  coordinates := make([]Coordinates2D, len(points))
  for i, v := range points {
    coordinates[i] = v.GetGeoJSON().Coordinates
  }
  return MultiPoint2DGeoJSON{"MultiPoint", coordinates}
}

// MultiPoint3D

type MultiPoint3DGeoJSON struct {
  Type string                 `json:"type"`
  Coordinates []Coordinates3D `json:"coordinates"`
}

func (m* MultiPoint3D) GetGeoJSON() MultiPoint3DGeoJSON {
  var points = m.GetPoint()
  coordinates := make([]Coordinates3D, len(points))
  for i, v := range points {
    coordinates[i] = v.GetGeoJSON().Coordinates
  }
  return MultiPoint3DGeoJSON{"MultiPoint", coordinates}
}

// LineString2D

type LineString2DGeoJSON struct {
  Type string                 `json:"type"`
  Coordinates []Coordinates2D `json:"coordinates"`
}

func (m* LineString2D) GetGeoJSON() LineString2DGeoJSON {
  var points = m.GetPoint()
  coordinates := make([]Coordinates2D, len(points))
  for i, v := range points {
    coordinates[i] = v.GetGeoJSON().Coordinates
  }
  return LineString2DGeoJSON{"LineString", coordinates}
}

// LineString3D

type LineString3DGeoJSON struct {
  Type string                 `json:"type"`
  Coordinates []Coordinates3D `json:"coordinates"`
}

func (m* LineString3D) GetGeoJSON() LineString3DGeoJSON {
  var points = m.GetPoint()
  coordinates := make([]Coordinates3D, len(points))
  for i, v := range points {
    coordinates[i] = v.GetGeoJSON().Coordinates
  }
  return LineString3DGeoJSON{"LineString", coordinates}
}

// MultiLineString2D

type MultiLineString2DGeoJSON struct {
  Type string                   `json:"type"`
  Coordinates [][]Coordinates2D `json:"coordinates"`
}

func (m* MultiLineString2D) GetGeoJSON() MultiLineString2DGeoJSON {
  var lineStrings = m.GetLineString()
  coordinates := make([][]Coordinates2D, len(lineStrings))
  for i, v := range lineStrings {
    coordinates[i] = v.GetGeoJSON().Coordinates
  }
  return MultiLineString2DGeoJSON{"MultiLineString", coordinates}
}

// MultiLineString3D

type MultiLineString3DGeoJSON struct {
  Type string                   `json:"type"`
  Coordinates [][]Coordinates3D `json:"coordinates"`
}

func (m* MultiLineString3D) GetGeoJSON() MultiLineString3DGeoJSON {
  var lineStrings = m.GetLineString()
  coordinates := make([][]Coordinates3D, len(lineStrings))
  for i, v := range lineStrings {
    coordinates[i] = v.GetGeoJSON().Coordinates
  }
  return MultiLineString3DGeoJSON{"MultiLineString", coordinates}
}

// Polygon2D

type Polygon2DGeoJSON struct {
  Type string                   `json:"type"`
  Coordinates [][]Coordinates2D `json:"coordinates"`
}

func (m* Polygon2D) GetGeoJSON() Polygon2DGeoJSON {
  var points = m.GetPoint()
  coordinates := make([][]Coordinates2D, 1)
  coordinates[0] = make([]Coordinates2D, len(points))
  for i, v := range points {
    coordinates[0][i] = v.GetGeoJSON().Coordinates
  }
  return Polygon2DGeoJSON{"Polygon", coordinates}
}

// Polygon3D

type Polygon3DGeoJSON struct {
  Type string                   `json:"type"`
  Coordinates [][]Coordinates3D `json:"coordinates"`
}

func (m* Polygon3D) GetGeoJSON() Polygon3DGeoJSON {
  var points = m.GetPoint()
  coordinates := make([][]Coordinates3D, 1)
  coordinates[0] = make([]Coordinates3D, len(points))
  for i, v := range points {
    coordinates[0][i] = v.GetGeoJSON().Coordinates
  }
  return Polygon3DGeoJSON{"Polygon", coordinates}
}

// MultiPolygon2D

type MultiPolygon2DGeoJSON struct {
  Type string                     `json:"type"`
  Coordinates [][][]Coordinates2D `json:"coordinates"`
}

func (m* MultiPolygon2D) GetGeoJSON() MultiPolygon2DGeoJSON {
  var polygons = m.GetPolygon()
  coordinates := make([][][]Coordinates2D, len(polygons))
  for i, v := range polygons {
    coordinates[i] = v.GetGeoJSON().Coordinates
  }
  return MultiPolygon2DGeoJSON{"MultiPolygon", coordinates}
}

// MultiPolygon3D

type MultiPolygon3DGeoJSON struct {
  Type string                     `json:"type"`
  Coordinates [][][]Coordinates3D `json:"coordinates"`
}

func (m* MultiPolygon3D) GetGeoJSON() MultiPolygon3DGeoJSON {
  var polygons = m.GetPolygon()
  coordinates := make([][][]Coordinates3D, len(polygons))
  for i, v := range polygons {
    coordinates[i] = v.GetGeoJSON().Coordinates
  }
  return MultiPolygon3DGeoJSON{"MultiPolygon", coordinates}
}
