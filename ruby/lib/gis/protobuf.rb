require "gis/protobuf/point_pb"
require "gis/protobuf/linestring_pb"
require "gis/protobuf/polygon_pb"

Gis::Protobuf::Point2D.class_eval {
  def toGeoJSON()
    return {
      type: "Point",
      coordinates: [self.x, self.y]
    }
  end
}

Gis::Protobuf::Point3D.class_eval {
  def toGeoJSON()
    return {
      type: "Point",
      coordinates: [self.x, self.y, self.z]
    }
  end
}

Gis::Protobuf::MultiPoint2D.class_eval {
  def toGeoJSON()
    return {
      type: "MultiPoint",
      coordinates: self.point.map { |point| point.toGeoJSON()[:coordinates] }
    }
  end
}

Gis::Protobuf::MultiPoint3D.class_eval {
  def toGeoJSON()
    return {
      type: "MultiPoint",
      coordinates: self.point.map { |point| point.toGeoJSON()[:coordinates] }
    }
  end
}

Gis::Protobuf::LineString2D.class_eval {
  def toGeoJSON()
    return {
      type: "LineString",
      coordinates: self.point.map { |point| point.toGeoJSON()[:coordinates] }
    }
  end
}

Gis::Protobuf::LineString3D.class_eval {
  def toGeoJSON()
    return {
      type: "LineString",
      coordinates: self.point.map { |point| point.toGeoJSON()[:coordinates] }
    }
  end
}

Gis::Protobuf::MultiLineString2D.class_eval {
  def toGeoJSON()
    return {
      type: "MultiLineString",
      coordinates: self.line_string.map { |line_string| line_string.toGeoJSON()[:coordinates] }
    }
  end
}

Gis::Protobuf::MultiLineString3D.class_eval {
  def toGeoJSON()
    return {
      type: "MultiLineString",
      coordinates: self.line_string.map { |line_string| line_string.toGeoJSON()[:coordinates] }
    }
  end
}

Gis::Protobuf::Polygon2D.class_eval {
  def toGeoJSON()
    return {
      type: "Polygon",
      coordinates: [self.point.map { |point| point.toGeoJSON()[:coordinates] }]
    }
  end
}

Gis::Protobuf::Polygon3D.class_eval {
  def toGeoJSON()
    return {
      type: "Polygon",
      coordinates: [self.point.map { |point| point.toGeoJSON()[:coordinates] }]
    }
  end
}

Gis::Protobuf::MultiPolygon2D.class_eval {
  def toGeoJSON()
    return {
      type: "MultiPolygon",
      coordinates: self.polygon.map { |polygon| polygon.toGeoJSON()[:coordinates] }
    }
  end
}

Gis::Protobuf::MultiPolygon3D.class_eval {
  def toGeoJSON()
    return {
      type: "MultiPolygon",
      coordinates: self.polygon.map { |polygon| polygon.toGeoJSON()[:coordinates] }
    }
  end
}
