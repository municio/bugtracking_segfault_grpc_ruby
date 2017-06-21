require_relative 'track_pb'

thds = []

2.times do
  thds << Thread.new do
    o = Repro::Outer.new
    o.items[0] = Repro::Inner.new
    raw = Repro::Outer.encode(o)

    100000.times do
      Repro::Outer.decode(raw)
    end
  end
end

thds.each(&:join)
