const GIS = require('../')

describe('gis.protobuf.LineString2D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point2D(); point1.setX(1); point1.setY(2)
    let point2 = new GIS.Point2D(); point2.setX(3); point2.setY(4)

    let lineString = new GIS.LineString2D(); lineString.setPointList([point1, point2])

    expect(lineString.toGeoJSON()).toEqual({
      type: 'LineString',
      coordinates: [[1, 2], [3, 4]]
    })
  })
})

describe('gis.protobuf.LineString3D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point3D(); point1.setX(1); point1.setY(2); point1.setZ(3)
    let point2 = new GIS.Point3D(); point2.setX(4); point2.setY(5); point2.setZ(6)

    let lineString = new GIS.LineString3D(); lineString.setPointList([point1, point2])

    expect(lineString.toGeoJSON()).toEqual({
      type: 'LineString',
      coordinates: [[1, 2, 3], [4, 5, 6]]
    })
  })
})

describe('gis.protobuf.MultiLineString2D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point2D(); point1.setX(1); point1.setY(2)
    let point2 = new GIS.Point2D(); point2.setX(3); point2.setY(4)
    let lineString1 = new GIS.LineString2D(); lineString1.setPointList([point1, point2])

    let point3 = new GIS.Point2D(); point3.setX(5); point3.setY(6)
    let point4 = new GIS.Point2D(); point4.setX(7); point4.setY(8)
    let lineString2 = new GIS.LineString2D(); lineString2.setPointList([point3, point4])

    let multiLineString = new GIS.MultiLineString2D(); multiLineString.setLineStringList([lineString1, lineString2])

    expect(multiLineString.toGeoJSON()).toEqual({
      type: 'MultiLineString',
      coordinates: [[[1, 2], [3, 4]], [[5, 6], [7, 8]]]
    })
  })
})

describe('gis.protobuf.MultiLineString3D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point3D(); point1.setX(1); point1.setY(2); point1.setZ(3)
    let point2 = new GIS.Point3D(); point2.setX(4); point2.setY(5); point2.setZ(6)
    let lineString1 = new GIS.LineString3D(); lineString1.setPointList([point1, point2])

    let point3 = new GIS.Point3D(); point3.setX(7); point3.setY(8); point3.setZ(9)
    let point4 = new GIS.Point3D(); point4.setX(10); point4.setY(11); point4.setZ(12)
    let lineString2 = new GIS.LineString3D(); lineString2.setPointList([point3, point4])

    let multiLineString = new GIS.MultiLineString3D(); multiLineString.setLineStringList([lineString1, lineString2])

    expect(multiLineString.toGeoJSON()).toEqual({
      type: 'MultiLineString',
      coordinates: [[[1, 2, 3], [4, 5, 6]], [[7, 8, 9], [10, 11, 12]]]
    })
  })
})
