# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from chirpstack_api.as_pb.external.api import internal_pb2 as chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class InternalServiceStub(object):
    """InternalService is the service providing API endpoints for internal usage.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Login = channel.unary_unary(
                '/api.InternalService/Login',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.LoginRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.LoginResponse.FromString,
                )
        self.Profile = channel.unary_unary(
                '/api.InternalService/Profile',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ProfileResponse.FromString,
                )
        self.Branding = channel.unary_unary(
                '/api.InternalService/Branding',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.BrandingResponse.FromString,
                )
        self.GlobalSearch = channel.unary_unary(
                '/api.InternalService/GlobalSearch',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.GlobalSearchRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.GlobalSearchResponse.FromString,
                )
        self.CreateAPIKey = channel.unary_unary(
                '/api.InternalService/CreateAPIKey',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.CreateAPIKeyRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.CreateAPIKeyResponse.FromString,
                )
        self.DeleteAPIKey = channel.unary_unary(
                '/api.InternalService/DeleteAPIKey',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.DeleteAPIKeyRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.ListAPIKeys = channel.unary_unary(
                '/api.InternalService/ListAPIKeys',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ListAPIKeysRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ListAPIKeysResponse.FromString,
                )


class InternalServiceServicer(object):
    """InternalService is the service providing API endpoints for internal usage.
    """

    def Login(self, request, context):
        """Log in a user
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Profile(self, request, context):
        """Get the current user's profile
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Branding(self, request, context):
        """Get the branding for the UI
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GlobalSearch(self, request, context):
        """Perform a global search.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateAPIKey(self, request, context):
        """CreateAPIKey creates the given API key.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteAPIKey(self, request, context):
        """DeleteAPIKey deletes the API key.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListAPIKeys(self, request, context):
        """ListAPIKeys lists the available API keys.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_InternalServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Login': grpc.unary_unary_rpc_method_handler(
                    servicer.Login,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.LoginRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.LoginResponse.SerializeToString,
            ),
            'Profile': grpc.unary_unary_rpc_method_handler(
                    servicer.Profile,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ProfileResponse.SerializeToString,
            ),
            'Branding': grpc.unary_unary_rpc_method_handler(
                    servicer.Branding,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.BrandingResponse.SerializeToString,
            ),
            'GlobalSearch': grpc.unary_unary_rpc_method_handler(
                    servicer.GlobalSearch,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.GlobalSearchRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.GlobalSearchResponse.SerializeToString,
            ),
            'CreateAPIKey': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateAPIKey,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.CreateAPIKeyRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.CreateAPIKeyResponse.SerializeToString,
            ),
            'DeleteAPIKey': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteAPIKey,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.DeleteAPIKeyRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'ListAPIKeys': grpc.unary_unary_rpc_method_handler(
                    servicer.ListAPIKeys,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ListAPIKeysRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ListAPIKeysResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.InternalService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class InternalService(object):
    """InternalService is the service providing API endpoints for internal usage.
    """

    @staticmethod
    def Login(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.InternalService/Login',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.LoginRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.LoginResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Profile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.InternalService/Profile',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ProfileResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Branding(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.InternalService/Branding',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.BrandingResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GlobalSearch(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.InternalService/GlobalSearch',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.GlobalSearchRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.GlobalSearchResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateAPIKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.InternalService/CreateAPIKey',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.CreateAPIKeyRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.CreateAPIKeyResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteAPIKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.InternalService/DeleteAPIKey',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.DeleteAPIKeyRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListAPIKeys(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.InternalService/ListAPIKeys',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ListAPIKeysRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_internal__pb2.ListAPIKeysResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
