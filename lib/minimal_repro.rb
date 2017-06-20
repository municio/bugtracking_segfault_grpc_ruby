require 'base64'
require 'pp'

$LOAD_PATH.unshift(File.dirname(__FILE__ ))

require_relative 'munic/type/event_pb.rb'

if defined?(GRPC)
  raise "grpc loaded. sanity check failed"
end

thds = []

2.times do
  thds << Thread.new do
    tr = Munic::Type::Track.new
    tr.namespaces[0] = Munic::Type::Namespace.new
    raw = Munic::Type::Track.encode(tr)

    1000000.times do
      Munic::Type::Track.decode(raw)
    end
  end
end

thds.each(&:join)
