# coding: utf-8
lib = File.expand_path("../lib", __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require "gis/protobuf/version"

Gem::Specification.new do |spec|
  spec.name          = "gis-protobuf"
  spec.version       = Gis::Protobuf::VERSION
  spec.authors       = ["Tomasz Jakub Rup"]
  spec.email         = ["tomasz.rup@gmail.com"]

  spec.summary       = "GIS ProtoBuf module"
  spec.homepage      = "https://github.com/tomi77/protobuf-gis/"
  spec.license       = "MIT"

  spec.files         = `git ls-files -z`.split("\x0").reject do |f|
    f.match(%r{^(test|spec|features)/})
  end
  spec.bindir        = "exe"
  spec.executables   = spec.files.grep(%r{^exe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]

  spec.add_runtime_dependency "google-protobuf", "~> 3.3.0"

  spec.add_development_dependency "bundler", "~> 1.15"
  spec.add_development_dependency "rake", "~> 10.0"
  spec.add_development_dependency "rspec", "~> 3.0"
end
