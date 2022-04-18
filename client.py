import grpc
import tensorflow as tf
from tensorflow_serving.apis import predict_pb2
from tensorflow_serving.apis import prediction_service_pb2_grpc
from absl import logging, flags, app

FLAGS = flags.FLAGS
flags.DEFINE_string("host", "localhost:8500", "tfs host")


def main(_):
    """
    TF Serving Client code example
    """
    logging.info(f"Host: {FLAGS.host}")

    channel = grpc.insecure_channel(FLAGS.host)
    stub = prediction_service_pb2_grpc.PredictionServiceStub(channel)
    request = predict_pb2.PredictRequest()
    request.model_spec.name = "model"
    request.model_spec.signature_name = "serving_default"
    request.inputs["x"].CopyFrom(tf.make_tensor_proto(tf.constant([1.0, 2.0, 5.0])))
    response = stub.Predict(request, 10.0)

    logging.info(f"Input: {request}")
    logging.info(f"Output: {response}")


if __name__ == "__main__":
    app.run(main)
