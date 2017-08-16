require "spec_helper"

RSpec.describe "Gis::Protobuf::Polygon2D has getGeoJSON function" do
  it "which returns a GeoJSON object" do
    point = Gis::Protobuf::Polygon2D.new point: [
      Gis::Protobuf::Point2D.new(x: 1.0, y: 2.0),
      Gis::Protobuf::Point2D.new(x: 3.0, y: 4.0),
    ]

    expect(point.toGeoJSON()).to eq type: "Polygon", coordinates: [[[1.0, 2.0], [3.0, 4.0]]]
  end
end

RSpec.describe "Gis::Protobuf::Polygon3D has getGeoJSON function" do
  it "which returns a GeoJSON object" do
    point = Gis::Protobuf::Polygon3D.new point: [
      Gis::Protobuf::Point3D.new(x: 1.0, y: 2.0, z: 3.0),
      Gis::Protobuf::Point3D.new(x: 4.0, y: 5.0, z: 6.0),
    ]

    expect(point.toGeoJSON()).to eq type: "Polygon", coordinates: [[[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]]]
  end
end

RSpec.describe "Gis::Protobuf::MultiPolygon2D has getGeoJSON function" do
  it "which returns a GeoJSON object" do
    point = Gis::Protobuf::MultiPolygon2D.new polygon: [
      Gis::Protobuf::Polygon2D.new(point: [
        Gis::Protobuf::Point2D.new(x: 1.0, y: 2.0),
        Gis::Protobuf::Point2D.new(x: 3.0, y: 4.0),
      ]),
      Gis::Protobuf::Polygon2D.new(point: [
        Gis::Protobuf::Point2D.new(x: 5.0, y: 6.0),
        Gis::Protobuf::Point2D.new(x: 7.0, y: 8.0),
      ])
    ]

    expect(point.toGeoJSON()).to eq type: "MultiPolygon", coordinates: [[[[1.0, 2.0], [3.0, 4.0]]], [[[5.0, 6.0], [7.0, 8.0]]]]
  end
end

RSpec.describe "Gis::Protobuf::MultiPolygon3D has getGeoJSON function" do
  it "which returns a GeoJSON object" do
    point = Gis::Protobuf::MultiPolygon3D.new polygon: [
      Gis::Protobuf::Polygon3D.new(point: [
        Gis::Protobuf::Point3D.new(x: 1.0, y: 2.0, z: 3.0),
        Gis::Protobuf::Point3D.new(x: 4.0, y: 5.0, z: 6.0),
      ]),
      Gis::Protobuf::Polygon3D.new(point: [
        Gis::Protobuf::Point3D.new(x: 7.0, y: 8.0, z: 9.0),
        Gis::Protobuf::Point3D.new(x: 10.0, y: 11.0, z: 12.0),
      ])
    ]

    expect(point.toGeoJSON()).to eq type: "MultiPolygon", coordinates: [[[[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]]], [[[7.0, 8.0, 9.0], [10.0, 11.0, 12.0]]]]
  end
end
