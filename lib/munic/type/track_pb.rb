# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: munic/type/track.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "munic.type.Track" do
    map :namespaces, :int32, :message, 1, "munic.type.Namespace"
  end
  add_message "munic.type.Namespace" do
  end
end

module Munic
  module Type
    Track = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.Track").msgclass
    Namespace = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.Namespace").msgclass
  end
end
