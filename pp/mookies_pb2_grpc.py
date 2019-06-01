# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import mookies_pb2 as mookies__pb2


class MenuServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetMenu = channel.unary_unary(
        '/mookiespb.MenuService/GetMenu',
        request_serializer=mookies__pb2.Empty.SerializeToString,
        response_deserializer=mookies__pb2.Menu.FromString,
        )


class MenuServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetMenu(self, request, context):
    """Unary get menu
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_MenuServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetMenu': grpc.unary_unary_rpc_method_handler(
          servicer.GetMenu,
          request_deserializer=mookies__pb2.Empty.FromString,
          response_serializer=mookies__pb2.Menu.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'mookiespb.MenuService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class OrderServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SubmitOrder = channel.unary_unary(
        '/mookiespb.OrderService/SubmitOrder',
        request_serializer=mookies__pb2.SubmitOrderRequest.SerializeToString,
        response_deserializer=mookies__pb2.SubmitOrderResponse.FromString,
        )
    self.CompleteOrder = channel.unary_unary(
        '/mookiespb.OrderService/CompleteOrder',
        request_serializer=mookies__pb2.CompleteOrderRequest.SerializeToString,
        response_deserializer=mookies__pb2.CompleteOrderResponse.FromString,
        )
    self.SubscribeToOrders = channel.unary_stream(
        '/mookiespb.OrderService/SubscribeToOrders',
        request_serializer=mookies__pb2.SubscribeToOrderRequest.SerializeToString,
        response_deserializer=mookies__pb2.Order.FromString,
        )
    self.SubscribeToCompleteOrders = channel.unary_stream(
        '/mookiespb.OrderService/SubscribeToCompleteOrders',
        request_serializer=mookies__pb2.SubscribeToOrderRequest.SerializeToString,
        response_deserializer=mookies__pb2.Order.FromString,
        )
    self.ActiveOrders = channel.unary_unary(
        '/mookiespb.OrderService/ActiveOrders',
        request_serializer=mookies__pb2.Empty.SerializeToString,
        response_deserializer=mookies__pb2.OrdersResponse.FromString,
        )


class OrderServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def SubmitOrder(self, request, context):
    """Unary 
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CompleteOrder(self, request, context):
    """validate order can be unary

    completeorder unary
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SubscribeToOrders(self, request, context):
    """server streaming
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SubscribeToCompleteOrders(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ActiveOrders(self, request, context):
    """all active orders
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_OrderServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SubmitOrder': grpc.unary_unary_rpc_method_handler(
          servicer.SubmitOrder,
          request_deserializer=mookies__pb2.SubmitOrderRequest.FromString,
          response_serializer=mookies__pb2.SubmitOrderResponse.SerializeToString,
      ),
      'CompleteOrder': grpc.unary_unary_rpc_method_handler(
          servicer.CompleteOrder,
          request_deserializer=mookies__pb2.CompleteOrderRequest.FromString,
          response_serializer=mookies__pb2.CompleteOrderResponse.SerializeToString,
      ),
      'SubscribeToOrders': grpc.unary_stream_rpc_method_handler(
          servicer.SubscribeToOrders,
          request_deserializer=mookies__pb2.SubscribeToOrderRequest.FromString,
          response_serializer=mookies__pb2.Order.SerializeToString,
      ),
      'SubscribeToCompleteOrders': grpc.unary_stream_rpc_method_handler(
          servicer.SubscribeToCompleteOrders,
          request_deserializer=mookies__pb2.SubscribeToOrderRequest.FromString,
          response_serializer=mookies__pb2.Order.SerializeToString,
      ),
      'ActiveOrders': grpc.unary_unary_rpc_method_handler(
          servicer.ActiveOrders,
          request_deserializer=mookies__pb2.Empty.FromString,
          response_serializer=mookies__pb2.OrdersResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'mookiespb.OrderService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
