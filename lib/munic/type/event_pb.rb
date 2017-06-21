# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: munic/type/event.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
require 'munic/type/entity_pb'
require 'munic/type/track_pb'
require 'munic/type/connStateChange_pb'
require 'munic/type/message_pb'
require 'munic/type/ack_pb'
require 'munic/internals/type/eventMetadata_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "munic.type.Event" do
    optional :id, :uint64, 1
    optional :thread_id, :uint64, 2
    optional :thread_root, :message, 3, "munic.type.Entity"
    repeated :parent_ids, :uint64, 4
    optional :sender, :message, 5, "munic.type.Entity"
    repeated :recipients, :message, 6, "munic.type.Entity"
    optional :group_key, :message, 7, "munic.type.GroupKey"
    optional :connection_id, :uint64, 8
    optional :event_time, :message, 9, "google.protobuf.Timestamp"
    optional :ingest_time, :message, 10, "google.protobuf.Timestamp"
    map :tags, :string, :string, 15
    optional :internal, :message, 16, "munic.internals.type.EventMetadata"
    oneof :content do
      optional :track, :message, 11, "munic.type.Track"
      optional :connStateChange, :message, 12, "munic.type.ConnStateChange"
      optional :message, :message, 13, "munic.type.Message"
      optional :ack, :message, 14, "munic.type.Ack"
    end
  end
  add_message "munic.type.GroupKey" do
    oneof :key do
      optional :asset, :message, 1, "munic.type.Asset"
    end
  end
  add_message "munic.type.EventList" do
    repeated :events, :message, 1, "munic.type.Event"
  end
  add_message "munic.type.TrackList" do
    repeated :tracks, :message, 1, "munic.type.Track"
  end
end

module Munic
  module Type
    Event = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.Event").msgclass
    GroupKey = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.GroupKey").msgclass
    EventList = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.EventList").msgclass
    TrackList = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.TrackList").msgclass
  end
end
