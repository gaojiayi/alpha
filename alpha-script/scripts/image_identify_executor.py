from imageai.Prediction import ImagePrediction
import json
import argparse
import os
import sys

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
sys.path.append(BASE_DIR)
from config import config

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='manual to this script')
    parser.add_argument('--imageName', type=str, default=None)
    parser.add_argument('--modelType', type=str, default='ResNet')
    args = parser.parse_args()
    if args.imageName is None or args.modelType == '':
        raise Exception('magePath 或者 modelType 参数不正确')

    execution_path = os.getcwd()
    prediction = ImagePrediction()
    if args.modelType == 'SqueezeNet':
        prediction.setModelTypeAsSqueezeNet()
        prediction.setModelPath(BASE_DIR + "/models/identify/squeezenet_weights_tf_dim_ordering_tf_kernels.h5")
    if args.modelType == 'ResNet':
        prediction.setModelTypeAsResNet()
        prediction.setModelPath(BASE_DIR + "/models/identify/resnet50_weights_tf_dim_ordering_tf_kernels.h5")
    if args.modelType == 'DenseNet':
        prediction.setModelTypeAsDenseNet()
        prediction.setModelPath(BASE_DIR + "/models/identify/DenseNet-BC-121-32.h5")
    if args.modelType == 'Inception':
        prediction.setModelTypeAsDenseNet()
        prediction.setModelPath(BASE_DIR + "/models/identify/inception_v3_weights_tf_dim_ordering_tf_kernels.h5")

    # TODO 后面支持多图识别
    prediction.loadModel()
    predictions, probabilities = prediction.predictImage(config.UPDATE_PATH + args.imageName, result_count=5)
    ret = {}
    print("data:")
    for eachPrediction, eachProbability in zip(predictions, probabilities):
        ret["{}".format(eachPrediction)] = "{}%".format(round(float(eachProbability), 2))
    print(json.dumps(ret))
