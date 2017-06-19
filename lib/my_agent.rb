require 'singleton'
require 'logger'

require_relative 'munic/grpc/agent_services_pb'

class MyAgent < Munic::Grpc::ProcessService::Service
  include Singleton

  LOG = Logger.new(STDOUT)

  def process(process_request, _context)
    LOG.info "received process_request with state #{process_request.state.inspect} and munic base name space  #{process_request.events.first.track.namespaces['MUNIC'].base['POSITION'].inspect} with _context: #{_context}"
    Munic::Grpc::ProcessResponse.new
  end
end
