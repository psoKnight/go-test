package main

import (
	"fmt"
	"github.com/sipt/GoJsoner"
)

func main() {
	result, err := GoJsoner.Discard(`
      {
  "global_info": { //全局/公共信息
    "message_type": "MSDP", //string，Megvii Structured Data Protocol，固定不变
    "version": "2020.1", //string，协议版本
    "time_ms": "", //string，数据帧时间戳（毫秒）
    "tracebacks": [ //string数组，数据回溯记录，数据生产者/更改者可以增加log
      ""
    ],
    "device_id": "", //string，设备唯一标识id，该字段对设备非必须
    "data_uuid": "" //string，该帧数据唯一标识id，该字段对设备非必须
  },
  "full_images": [ //大图信息数组
    {
      "original_resolution": { //原始分辨率
        "width_pixels": 0, //int，宽
        "height_pixels": 0 //int，高
      },
      "image_data": { //图片数据
        "image_data_format": 0, //image_data_type枚举定义
        "value": "" //image_data_format对应的值
      },
      "processed_resolution": { //分析处理分辨率
        "width_pixels": 0, //int，宽
        "height_pixels": 0 //int，高
      },
      "saved_resolution": { //存图分辨率
        "width_pixels": 0, //int，宽
        "height_pixels": 0 //int，高
      }
    }
  ],
  "faces": [ //人脸对象数组
    {
      "original_rect": { // 原始坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "track_id": "", //小图track_id
      "full_image_index": 0, //int, 大图数组索引
      "image_data": { //图片数据
        "image_data_format": 0, //image_data_type枚举定义
        "value": "" //image_data_format对应的值
      },
      "cropped_rect": { // 处理后坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "gender": 0, //性别，gender_type枚举定义
      "age": 0, //int, 年龄，0为未知，其他为年龄值
      "wear_hat": 0, //是否戴帽子，ternary_type枚举定义
      "hat_color": 0, //帽子颜色，hat_color_type枚举定义
      "hat_style": 0, //帽子款式，hat_style_type枚举定义
      "beard_status": 0, //胡子状态，ternary_type枚举定义
      "beard_class": 0, //胡子类型，beard_class_type枚举定义
      "wear_glasses": 0, //带眼镜状态，ternary_type枚举定义
      "glasses_color": 0, //眼镜颜色，glasses_color_type枚举定义
      "glasses_style": 0, //眼镜款式，glasses_style_type枚举定义
      "hair_style": 0, //发型，hair_style_type枚举定义
      "hair_color": 0, //发色，hair_color_type枚举定义
      "wear_respirator": 0, //是否带口罩，ternary_type枚举定义
      "respirator_class": 0, //口罩类型，respirator_class_type枚举定义
      "skin_color": 0, //肤色，skin_color_type枚举定义
      "link_info": [ //其他对象关联信息
        {
          "object_class": 0, //绑定对象类型，object_class_type枚举定义，有绑定关系的对象数据不强制要求双向关联
          "object_track_id": "" //绑定对象track_id
        }
      ],
      "recognition_info": [
        {
          "group_id": 1, //int，底库ID
          "group_name": "", //string，底库名称
          "person_id": 0, //int, 底库人员ID
          "person_name": "", //string, 底库人员名称
          "face_score": 99.0, //float, 相似分数
          "gender": 0, //gender_type枚举定义
          "certificates_class": 0, //证件类型，certificates_class_type枚举定义
          "certificate_number": "", //string, 证件号
          "group_info_uuid": "", //string, 底库信息唯一标识
          "native_city_code": 0, //int, 籍贯城市代码，参考GB/T2260行政区划代码
          "ethic_code": 0, //int, 民族代码，参考GB/T3304
          "image_data": { //图片数据
            "image_data_format": 0, //image_data_type枚举定义
            "value": "" //image_data_format对应的值
          }
        }
      ]
    }
  ],
  "pedestrians": [ //人体对象数组
    {
      "original_rect": { // 原始坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "track_id": "", //小图track_id
      "full_image_index": 0, //int, 大图数组索引
      "image_data": { //图片数据
        "image_data_format": 0, //image_data_type枚举定义
        "value": "" //image_data_format对应的值
      },
      "cropped_rect": { // 处理后坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "gender": 0, //性别，gender_type枚举定义
      "age_upper_limit": 0, //int, 年龄上限，0为未知，其他为年龄值
      "age_lower_limit": 0, //int, 年龄下限，0为未知，其他为年龄值
      "bodytype": 0, //体型，body_type枚举定义
      "skin_color": 0, //肤色，skin_color_type枚举定义
      "wear_hat": 0, //是否戴帽子，ternary_type枚举定义
      "hat_color": 0, //帽子颜色，hat_color_type枚举定义
      "hat_style": 0, //帽子款式，hat_style_type枚举定义
      "carry_bag": 0, //是否携带包，ternary_type枚举定义
      "carry_backpack": 0, //是否携带双肩包，ternary_type枚举定义
      "carry_messenger_bag": 0, //是否携带单肩包，ternary_type枚举定义
      "carry_handbag": 0, //是否携带手提包，ternary_type枚举定义
      "carry_drag": 0, //是否携带拉杆箱，ternary_type枚举定义
      "hold_umbrella": 0, //是否打伞，ternary_type枚举定义
      "wear_glasses": 0, //是否戴眼镜，ternary_type枚举定义
      "glasses_color": 0, //眼镜颜色，glasses_color_type枚举定义
      "glasses_style": 0, //眼镜款式，glasses_style_type枚举定义
      "wear_overcoat": 0, //是否穿大衣，ternary_type枚举定义
      "coat_style": 0, //上衣款式，coat_style_type枚举定义
      "coat_length": 0, //上衣长度，coat_length_type枚举定义
      "coat_color": 0, //上衣颜色，coat_color_type枚举定义
      "coat_stripe": 0, //上衣条纹，coat_stripe_type枚举定义
      "wear_lower_skirt": 0, //是否穿裙子，ternary_type枚举定义
      "pants_style": 0, //裤子款式，pants_style_type枚举定义
      "pants_color": 0, //裤子颜色，pants_color_type枚举定义
      "pants_length": 0, //裤子长度，pants_length_type枚举定义
      "shoes_style": 0, //鞋子款式，shoes_style_type枚举定义
      "shoes_color": 0, //鞋子颜色，shoes_color_type枚举定义
      "hair_style": 0, //发型，hair_style_type枚举定义
      "hair_color": 0, //发色，hair_color_type枚举定义
      "ride_bike": 0, //是否骑车，ternary_type枚举定义
      "pedestrian_orientation": 0, //人员朝向，pedestrian_orientation_type枚举定义
      "coat_logo": 0, //上衣有无logo，ternary_type枚举定义
      "speed_status": 0, //运动速度状态，speed_status_type枚举定义
      "wear_helmet": 0, //是否戴头盔，ternary_type枚举定义
      "wear_safetycap": 0, //是否戴安全帽，ternary_type枚举定义
      "action_fall": 0, //行为动作-是否跌倒，ternary_type枚举定义
      "action_smoke": 0, //行为动作-是否抽烟，ternary_type枚举定义
      "action_run": 0, //行为动作-是否奔跑，ternary_type枚举定义
      "action_use_phone": 0, //行为动作-是否打电话，ternary_type枚举定义
      "action_watch_phone": 0, //行为动作-是否低头看手机，ternary_type枚举定义
      "link_info": [ //其他对象关联信息
        {
          "object_class": 0, //绑定对象类型，object_class_type枚举定义，有绑定关系的对象数据不强制要求双向关联
          "object_track_id": "" //绑定对象track_id
        }
      ]
    }
  ],
  "vehicles": [ //车辆对象数组
    {
      "original_rect": { // 原始坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "track_id": "", //小图track_id
      "full_image_index": 0, //int, 大图数组索引
      "image_data": { //图片数据
        "image_data_format": 0, //image_data_type枚举定义
        "value": "" //image_data_format对应的值
      },
      "cropped_rect": { // 处理后坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "vehicle_class": 0, //车辆类型，vehicle_class_type枚举定义
      "vehicle_brand": 0, //车辆品牌，vehicle_brand_type枚举定义
      "vehicle_color": 0, //车辆颜色，vehicle_color_type枚举定义
      "vehicle_orientation": 0, //车辆朝向，vehicle_orientation_type枚举定义
      "vehicle_sunroof": 0, //车辆有无天窗，ternary_type枚举定义
      "vehicle_rack": 0, //车辆有无行李架，ternary_type枚举定义
      "vehicle_spray": 0, //车身有无喷字，ternary_type枚举定义
      "annual_inspection": 0, //车辆有无年检标，ternary_type枚举定义
      "main_safety belt": 0, //主驾驶是否系安全带，ternary_type枚举定义
      "main_sun_visor": 0, //主驾驶是否用遮阳板，ternary_type枚举定义
      "vehicle_pendant": 0, //车辆是否有挂件，ternary_type枚举定义
      "vehicle_decoration": 0, //车辆是否有摆件，ternary_type枚举定义
      "main_calling": 0, //主驾驶是否打电话，ternary_type枚举定义
      "aux_sun_visor": 0, //副驾驶是否用遮阳板，ternary_type枚举定义
      "aux_has_person": 0, //副驾驶是否有人，ternary_type枚举定义
      "aux_safety belt": 0, //副驾驶是否系安全带，ternary_type枚举定义
      "inside_quality": 0, //车窗所有属性是否可见，ternary_type枚举定义
      "has_plate": 0, //车辆有无车牌，ternary_type枚举定义
      "link_info": [ //其他对象关联信息
        {
          "object_class": 0, //绑定对象类型，object_class_type枚举定义，有绑定关系的对象数据不强制要求双向关联
          "object_track_id": "" //绑定对象track_id
        }
      ]
    }
  ],
  "non-motors": [ //非机动车对象数组
    {
      "original_rect": { // 原始坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "track_id": "", //小图track_id
      "full_image_index": 0, //int, 大图数组索引
      "image_data": { //图片数据
        "image_data_format": 0, //image_data_type枚举定义
        "value": "" //image_data_format对应的值
      },
      "cropped_rect": { // 处理后坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "non_motor_class": 0, //非机动车辆类型，non_motor_class_type枚举定义
      "non_motor_orientation": 0, //非机动车辆朝向，non_motor_orientation_type枚举定义
      "non_motor_color": 0, //非机动车辆颜色，vehicle_color_type枚举定义
      "has_plate": 0, //车辆有无车牌，ternary_type枚举定义
      "rider_num": 1, //骑车人数，int类型
      "link_info": [ //其他对象关联信息
        {
          "object_class": 0, //绑定对象类型，object_class_type枚举定义，有绑定关系的对象数据不强制要求双向关联
          "object_track_id": "" //绑定对象track_id
        }
      ]
    }
  ],
  "plates": [ //车牌对象数组
    {
      "original_rect": { // 原始坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "track_id": "", //小图track_id
      "full_image_index": 0, //int, 大图数组索引
      "image_data": { //图片数据
        "image_data_format": 0, //image_data_type枚举定义
        "value": "" //image_data_format对应的值
      },
      "cropped_rect": { // 处理后坐标，小图相对于resolution坐标
        "top": 100, //int
        "left": 100, //int
        "width": 100, //int
        "height": 100 //int
      },
      "plate_class": 0, //车牌种类，plate_class_type枚举定义
      "plate_no": "", //车牌号码，string
      "plate_color": 0, //车牌颜色，plate_color_type枚举定义
      "plate_row": 0, //单双行车牌，plate_row_type枚举定义
      "plate_scores": [ //float, 置信分数
        99.0
      ],
      "plate_occlusion": 0, //车牌遮挡情况，plate_occlusion_type枚举定义
      "plate_dirty": 0, //车牌是否污损，ternary_type枚举定义
      "link_info": [ //其他对象关联信息
        {
          "object_class": 0, //绑定对象类型，object_class_type枚举定义，有绑定关系的对象数据不强制要求双向关联
          "object_track_id": "" //绑定对象track_id
        }
      ]
    }
  ],
  "additional": { //string，附加透传字段，由生产者和消费者自行定义
  }
}
`)
	if err != nil {
		fmt.Printf("Json discard err: %v.", err)
		return
	}

	fmt.Println(result)
}
