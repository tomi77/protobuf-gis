
/** This code is included by protobuf-gis package **/

/**
 * Creates an GeoJSON repesentation of a "LineString2D/LineString3D" type object.
 */
proto.gis.protobuf.LineString2D.prototype.toGeoJSON = proto.gis.protobuf.LineString3D.prototype.toGeoJSON = function() {
  let coordinates = this.getPointList().map(function (point) {
    return point.toGeoJSON().coordinates;
  });
  return {
    type: 'LineString',
    coordinates: coordinates
  };
};

/**
 * Creates an GeoJSON repesentation of a "MultiLineString2D/MultiLineString3D" type object.
 */
proto.gis.protobuf.MultiLineString2D.prototype.toGeoJSON = proto.gis.protobuf.MultiLineString3D.prototype.toGeoJSON = function() {
  let coordinates = this.getLineStringList().map(function (lineString) {
    return lineString.toGeoJSON().coordinates;
  });
  return {
    type: 'MultiLineString',
    coordinates: coordinates
  };
};
