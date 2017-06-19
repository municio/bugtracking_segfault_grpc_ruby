#!/usr/bin/env ruby

$LOAD_PATH.unshift(File.dirname(__FILE__ ) + '/../lib/')
require 'bundler/setup'
Bundler.require(:default)

require_relative '../lib/server'

AgentServer.new.run
