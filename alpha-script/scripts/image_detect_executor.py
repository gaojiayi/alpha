# 图像检测执行脚本
import os
import sys

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
sys.path.append(BASE_DIR)
# print(os.path.abspath(__file__))
# print(os.path.dirname(os.path.abspath(__file__)))
# print(BASE_DIR)
# print(sys.path)
# 该操作是把上一层目录添加到path中   此处是把alpha-script添加到path中，这样可以加载到config了
# sys.path.append("..")
from imageai.Detection import ObjectDetection
import os
from config import config
from utils import helper
import json

#  1 修改session
#  2 修改import tensorflow as tf
#  3 修改keras版本至2.3.1
#  4 修改vsite-packages/imageai/Detection/keras_retinanet/backend/tensorflow_backend.py:tensorflow.image.resize(*args, **kwargs)
#  5 修改RandomNormal
import argparse

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='manual to this script')
    parser.add_argument('--imageName', type=str, default=None)
    args = parser.parse_args()
    if args.imageName is None:
        raise Exception('magePath 参数不正确')
    # print(args.imageName, type(args.imageName))

    # try:
    detector = ObjectDetection()
    detector.setModelTypeAsRetinaNet()
    detector.setModelPath(BASE_DIR + "/models/detect/resnet50_coco_best_v2.0.1.h5")
    detector.loadModel(detection_speed=config.DETECTION_SPEED)
    unique_str = helper.tid_maker()
    # 判断路径是否存在
    isExists = os.path.exists(config.OUTPUT_IMAGE_PATH)
    if not isExists:
        os.makedirs(config.OUTPUT_IMAGE_PATH)
    detections, objects_path = detector.detectObjectsFromImage(
        input_image=config.UPDATE_PATH + args.imageName,
        output_image_path=config.OUTPUT_IMAGE_PATH + unique_str + ".jpg",
        extract_detected_objects=True)

    array = []
    for eachObject, eachObjectPath in zip(detections, objects_path):
        ret = {}
        ret['object'] = "{}".format(eachObject["name"])
        ret['percet'] = "{}%".format(round(float(eachObject["percentage_probability"]), 2))
        ret["path"] = unique_str + ".jpg-objects/" + "{}".format(eachObjectPath).split("/")[-1]
        array.append(ret)
    #特殊符号便于截取
    print("data:")
    print(json.dumps(array))
    # except Exception as e:
    #     print(e)
