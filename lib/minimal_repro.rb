require 'base64'

$LOAD_PATH.unshift(File.dirname(__FILE__ ))

require_relative 'munic/grpc/agent_pb'

if defined?(GRPC)
  raise "grpc loaded. sanity check failed"
end

thds = []

10.times do
  thds << Thread.new do
    event_64 = "COSCi7zW9eXVDRDkgou81vXl1Q0aACoAOg0KCwoDMTIzEgR0ZXN0QMiCgLKHr+TVDUoGCJaf88kFUgYIrKDzyQV6FwoKb2ZmZXJfbmFtZRIJbWV0cm9taWxlggEUChISCGNsb3VkLmluGAcgjLmdmwRasgEKiwEKClVOVFlQRURfTU0SfRJ7CnkKBk1NXzI0MBJvQm0KChD8DRjDv6usySsqByAIPQrXkz8yISCppvsDKNrF+wMyFa0PXpUCcys9QGQjhgSkAVTnA4cBIjosCKw7KBgwVUogNf7//////////wH+//////////8B/P//////////AQCABBBiBRDhoZhBCh8KBU1VTklDEhYKFAoSCSJPkq6Z5EJAEYE+kSdJm17AEJAP"
    state_64 = "CsABEr0B5gHwPnsic3RhdGVfdmVyc2lvbiI6MS4wLCJiYXRjaF9maXJzdF9ldmVudF9pZCI6OTg3MTM3NTc5NTMyOTQzNzE2LA0qBGxhnikASGVuZ3RoIjoxLCJ0cmFjayI6eyIFPVhyZWNvcmRlZF9hdCI6MTQ5NzY1NTkwMwGNEHJ1bGVzASk8MF9ib290XzI2NS1FVkVOVAEUlGFjdGl2ZSI6dHJ1ZSwiY291bnQiOjAsInZlcnNpb24iOjJ9fX19EgsI/umeygUQuLOAHw=="

    event_raw = Base64.decode64(event_64)
    state_raw = Base64.decode64(state_64)

    ev = Munic::Type::Event.decode(event_raw)
    state = Munic::Grpc::State.decode(state_raw)

    events = []
    5000.times do
      events << ev
    end

    req = Munic::Grpc::ProcessRequest.new(key: ev.group_key, events: events, state: state)
    raw = Munic::Grpc::ProcessRequest.encode(req)
    p "length of decoded protobuf is #{raw.size}"
    Munic::Grpc::ProcessRequest.decode(raw)
    p "decode successful"
  end
end

thds.each(&:join)
