# coding=utf-8

import logging

import grpc

from pb import demo_pb2
from pb import demo_pb2_grpc


def run():
    # 注意(gRPC Python Team): .close()方法在channel上是可用的。
    # 并且应该在with语句不符合代码需求的情况下使用。
    with grpc.insecure_channel('localhost:8180') as channel:
        stub = demo_pb2_grpc.GoGreeterStub(channel)
        response = stub.SayHello(demo_pb2.HelloRequest(name='python'))
    print("Greeter client received: {}!".format(response.message))


if __name__ == '__main__':
    logging.basicConfig()
    run()