require 'base64'
require 'pp'

$LOAD_PATH.unshift(File.dirname(__FILE__ ))

require_relative 'munic/type/event_pb.rb'

if defined?(GRPC)
  raise "grpc loaded. sanity check failed"
end

thds = []

10.times do
  thds << Thread.new do
    tr = Munic::Type::Track.new
    fields = Munic::Type::KnownFields::MunicFields.new
    tr.namespaces['MUNIC'] = Munic::Type::Namespace.new
    tr.namespaces['MUNIC'].base = fields

    tracks = []
    50000.times do
      tracks << tr
    end

    list = Munic::Type::TrackList.new(tracks: tracks)
    raw = Munic::Type::TrackList.encode(list)
    p "length of decoded protobuf is #{raw.size}"
    Munic::Type::TrackList.decode(raw)
    p "decode successful"
  end
end

thds.each(&:join)
