const GIS = require('../')

describe('gis.protobuf.Point2D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point = new GIS.Point2D(); point.setX(1); point.setY(2)

    expect(point.toGeoJSON()).toEqual({
      type: 'Point',
      coordinates: [1, 2]
    })
  })
})

describe('gis.protobuf.Point3D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point = new GIS.Point3D(); point.setX(1); point.setY(2); point.setZ(3)

    expect(point.toGeoJSON()).toEqual({
      type: 'Point',
      coordinates: [1, 2, 3]
    })
  })
})

describe('gis.protobuf.MultiPoint2D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point2D(); point1.setX(1); point1.setY(2)
    let point2 = new GIS.Point2D(); point2.setX(3); point2.setY(4)

    let multiPoint = new GIS.MultiPoint2D(); multiPoint.setPointList([point1, point2])

    expect(multiPoint.toGeoJSON()).toEqual({
      type: 'MultiPoint',
      coordinates: [[1, 2], [3, 4]]
    })
  })
})

describe('gis.protobuf.MultiPoint3D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point3D(); point1.setX(1); point1.setY(2); point1.setZ(3)
    let point2 = new GIS.Point3D(); point2.setX(4); point2.setY(5); point2.setZ(6)

    let multiPoint = new GIS.MultiPoint3D(); multiPoint.setPointList([point1, point2])

    expect(multiPoint.toGeoJSON()).toEqual({
      type: 'MultiPoint',
      coordinates: [[1, 2, 3], [4, 5, 6]]
    })
  })
})
