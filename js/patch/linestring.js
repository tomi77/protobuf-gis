
/** This code is included by protobuf-gis package **/

/**
 * Creates an GeoJSON repesentation of a "LineString2D/LineString3D" type object.
 */
proto.gis.protobuf.LineString2D.prototype.toGeoJSON = proto.gis.protobuf.LineString3D.prototype.toGeoJSON = function() {
  let coordinates = null;
  if (this.hasPointList()) {
    coordinates = this.getPointList().map(function (point) {
      return point.toGeoJSON();
    });
  }
  return {
    type: 'LineString',
    coordinates: coordinates
  };
};

/**
 * Creates an GeoJSON repesentation of a "MultiLineString2D/MultiLineString3D" type object.
 */
proto.gis.protobuf.MultiLineString2D.prototype.toGeoJSON = proto.gis.protobuf.MultiLineString3D.prototype.toGeoJSON = function() {
  let coordinates = null;
  if (this.getLineStringList()) {
    coordinates = this.getLineStringList().map(function (lineString) {
      if (lineString.hasPointList()) {
        return lineString.toGeoJSON().coordinates;
      } else {
        return null
      }
    });
  }
  return {
    type: 'MultiLineString',
    coordinates: coordinates
  };
};
