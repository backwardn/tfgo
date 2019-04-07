# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import rtf_pb2 as rtf__pb2


class RTFStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.DefineAndCall = channel.stream_stream(
        '/RTF/DefineAndCall',
        request_serializer=rtf__pb2.RTFStatement.SerializeToString,
        response_deserializer=rtf__pb2.RTFResponse.FromString,
        )


class RTFServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def DefineAndCall(self, request_iterator, context):
    """accept a stream of RTFStatement that define the Python function body.
    Returns a stream of RTFResponse.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RTFServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'DefineAndCall': grpc.stream_stream_rpc_method_handler(
          servicer.DefineAndCall,
          request_deserializer=rtf__pb2.RTFStatement.FromString,
          response_serializer=rtf__pb2.RTFResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'RTF', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
