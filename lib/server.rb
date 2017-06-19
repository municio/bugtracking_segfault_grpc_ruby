#!/usr/bin/env ruby

$LOAD_PATH.unshift(File.dirname(__FILE__))
require 'bundler/setup'
Bundler.require(:default)

require 'logger'

require 'my_agent'

class AgentServer
  attr_accessor :server

  def initialize()
    $stdout.sync = true
    @logger = Logger.new(STDOUT)
    @server = GRPC::RpcServer.new(pool_size: 60)
    @server.add_http2_port('localhost:50070', :this_port_is_insecure)
    @server.handle(MyAgent.instance)
  end

  def run()
    Signal.trap("INT") {
      Thread.new{
        @logger.info("Shutting down gracefully...")
        shut_down
      }
    }
    Signal.trap("TERM") {
      Thread.new{
        @logger.info("Shutting down gracefully...")
        shut_down
      }
    }
    @logger.info("Starting server on :50070 (insecure)")
    @server.run_till_terminated
    @logger.info("server shut down")
  end

  def shut_down()
    @server.stop
  end
end
