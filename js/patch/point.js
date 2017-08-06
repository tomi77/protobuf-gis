
/** This code is included by protobuf-gis package **/

/**
 * Creates an GeoJSON repesentation of a "Point2D/Point3D" type object.
 */
proto.gis.protobuf.Point2D.prototype.toGeoJSON = proto.gis.protobuf.Point3D.prototype.toGeoJSON = function() {
  let coordinates = null;
  if (this.hasPoint()) {
    coordinates = this.getPoint().toGeoJSON();
  }
  return {
    type: 'Point',
    coordinates: coordinates
  };
};

/**
 * Creates an GeoJSON repesentation of a "MultiPoint2D/MultiPoint3D" type object.
 */
proto.gis.protobuf.MultiPoint2D.prototype.toGeoJSON = proto.gis.protobuf.MultiPoint3D.prototype.toGeoJSON = function() {
  let coordinates = null;
  if (this.hasPointList()) {
    coordinates = this.getPointList().map(function (point) {
      return point.toGeoJSON();
    });
  }
  return {
    type: 'MultiPoint',
    coordinates: coordinates
  };
};
