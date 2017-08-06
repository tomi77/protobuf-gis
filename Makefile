BUILDDIR = /usr/include/gis/protobuf/

all: $(BUILDDIR) $(BUILDDIR)*.proto

$(BUILDDIR):
	@mkdir -p $@

$(BUILDDIR)*.proto: gis/protobuf/*.proto
	@cp $^ $(BUILDDIR)
