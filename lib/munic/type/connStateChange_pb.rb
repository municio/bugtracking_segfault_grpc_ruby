# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: munic/type/connStateChange.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "munic.type.ConnStateChange" do
    optional :type, :enum, 1, "munic.type.ConnStateChange.Type"
    optional :asset_reason, :enum, 2, "munic.type.ConnStateChange.AssetReason"
    optional :bs_reason, :enum, 3, "munic.type.ConnStateChange.BSReason"
    optional :ip, :string, 4
    optional :protocol, :string, 5
    optional :fullreason, :string, 6
  end
  add_enum "munic.type.ConnStateChange.Type" do
    value :UNKNOWN_TYPE, 0
    value :CONNECT, 1
    value :DISCONNECT, 2
    value :CONNECT_TRY_FAIL, 3
  end
  add_enum "munic.type.ConnStateChange.AssetReason" do
    value :UNKNOWN_ASSET_REASON, 0
    value :COLD_BOOT, 1
    value :SUSPEND_BOOT, 2
    value :IDLE_OUT, 3
    value :PPP_LOST, 4
    value :MODEM_RESET, 5
    value :SIM_ERROR, 6
    value :ROAMING, 7
    value :CONNECTION_LOST, 8
    value :NEW_CONFIG, 9
    value :WRITE_ERROR, 10
    value :READ_ERROR, 11
    value :CLOSED_BY_SERVER, 12
  end
  add_enum "munic.type.ConnStateChange.BSReason" do
    value :UNKNOWN_BS_REASON, 0
    value :SOCKET_CLOSED, 1
    value :NETWORK_ERROR, 2
    value :CLIENT_DISCONNECT, 3
    value :NETWORK_ACTIVITY_TIMEOUT, 4
    value :MESSAGE_ACK_TIMEOUT, 5
    value :BASEVALUE_ACK_TIMEOUT, 6
    value :UNKNOWN_CHANNEL, 7
    value :UNKNOWN_DYNAMIC_CHANNEL, 8
    value :SERVER_SHUTDOWN, 9
    value :AUTH_FAILED_ASSET, 10
    value :AUTH_FAILED_ACCOUNT, 11
    value :AUTH_FAILED_IP, 12
    value :AUTH_FAILED_TRANSPORT, 13
  end
end

module Munic
  module Type
    ConnStateChange = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.ConnStateChange").msgclass
    ConnStateChange::Type = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.ConnStateChange.Type").enummodule
    ConnStateChange::AssetReason = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.ConnStateChange.AssetReason").enummodule
    ConnStateChange::BSReason = Google::Protobuf::DescriptorPool.generated_pool.lookup("munic.type.ConnStateChange.BSReason").enummodule
  end
end
