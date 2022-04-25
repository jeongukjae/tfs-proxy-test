#

## Run TensorFlow Serving Docker image

```sh
$ docker run \
    --rm -it \
    -v `pwd`/saved_model_half_plus_two_cpu:/models/model:ro \
    -p 8500:8500 \
    tensorflow/serving
```
