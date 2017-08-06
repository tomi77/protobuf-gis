
/** This code is included by protobuf-gis package **/

/**
 * Creates an GeoJSON repesentation of a "Polygon2D/Polygon3D" type object.
 */
proto.gis.protobuf.Polygon2D.prototype.toGeoJSON = proto.gis.protobuf.Polygon3D.prototype.toGeoJSON = function() {
  let coordinates = null;
  if (this.hasPointList()) {
    coordinates = this.getPointList().map(function (point) {
      return point.toGeoJSON();
    });
  }
  return {
    type: 'Polygon',
    coordinates: [[coordinates]]
  };
};

/**
 * Creates an GeoJSON repesentation of a "MultiPolygon2D/MultiPolygon3D" type object.
 */
proto.gis.protobuf.MultiPolygon2D.prototype.toGeoJSON = proto.gis.protobuf.MultiPolygon3D.prototype.toGeoJSON = function() {
  let coordinates = null;
  if (this.getPolygonList()) {
    coordinates = this.getPolygonList().map(function (polygon) {
      if (polygon.hasPointList()) {
        return polygon.toGeoJSON().coordinates;
      } else {
        return null
      }
    });
  }
  return {
    type: 'MultiPolygon',
    coordinates: coordinates
  };
};
