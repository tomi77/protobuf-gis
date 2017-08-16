require "spec_helper"

RSpec.describe "Gis::Protobuf::Point2D has getGeoJSON function" do
  it "which returns a GeoJSON object" do
    point = Gis::Protobuf::Point2D.new x: 1.0, y: 2.0

    expect(point.toGeoJSON()).to eq type: "Point", coordinates: [1.0, 2.0]
  end
end

RSpec.describe "Gis::Protobuf::Point3D has getGeoJSON function" do
  it "which returns a GeoJSON object" do
    point = Gis::Protobuf::Point3D.new x: 1.0, y: 2.0, z: 3.0

    expect(point.toGeoJSON()).to eq type: "Point", coordinates: [1.0, 2.0, 3.0]
  end
end

RSpec.describe "Gis::Protobuf::MultiPoint2D has getGeoJSON function" do
  it "which returns a GeoJSON object" do
    multi_point = Gis::Protobuf::MultiPoint2D.new point: [
      Gis::Protobuf::Point2D.new(x: 1.0, y: 2.0),
      Gis::Protobuf::Point2D.new(x: 3.0, y: 4.0)
    ]

    expect(multi_point.toGeoJSON()).to eq type: "MultiPoint", coordinates: [[1.0, 2.0], [3.0, 4.0]]
  end
end

RSpec.describe "Gis::Protobuf::MultiPoint3D has getGeoJSON function" do
  it "which returns a GeoJSON object" do
    multi_point = Gis::Protobuf::MultiPoint3D.new point: [
      Gis::Protobuf::Point3D.new(x: 1.0, y: 2.0, z: 3.0),
      Gis::Protobuf::Point3D.new(x: 4.0, y: 5.0, z: 6.0)
    ]

    expect(multi_point.toGeoJSON()).to eq type: "MultiPoint", coordinates: [[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]]
  end
end
