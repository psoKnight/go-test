package main

import (
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(len("QyazFh4R+PUM+wAPEeIQ9TDaxSrZzB4LAisfCfr8KyogGdwv/9kYFs8Z9b79CT0QDAcZH98E/urzwvgSIPj/+MkHCBHn7Bnw4PERFjzpGgsN/CALE88V8xkg3tkM1u/0B/IwBRAQ7wINLfAGAykNzhgEBd7n/8L31fL7MwgH/u8G+Qbr7x4K/0nm88QBBgPzIAUO7fEA4yEt//UBF9TtugMCGiz2JvASyzUJDuDx5vL99QUf/g0zF/MKEAPpRAcN9+TlAxkJCvoKxBD/GfMrCBj39ugwsxngNf3k9OwLHQTbA+xIBQP99ev44NEkCgDFEsru/PbLVhD/z/AG7yzY5w=="))
	decodeString, err := base64.StdEncoding.DecodeString("QyazFh4R+PUM+wAPEeIQ9TDaxSrZzB4LAisfCfr8KyogGdwv/9kYFs8Z9b79CT0QDAcZH98E/urzwvgSIPj/+MkHCBHn7Bnw4PERFjzpGgsN/CALE88V8xkg3tkM1u/0B/IwBRAQ7wINLfAGAykNzhgEBd7n/8L31fL7MwgH/u8G+Qbr7x4K/0nm88QBBgPzIAUO7fEA4yEt//UBF9TtugMCGiz2JvASyzUJDuDx5vL99QUf/g0zF/MKEAPpRAcN9+TlAxkJCvoKxBD/GfMrCBj39ugwsxngNf3k9OwLHQTbA+xIBQP99ev44NEkCgDFEsru/PbLVhD/z/AG7yzY5w==")
	if err != nil {
		logrus.Errorf("decode str err: %v", err)
		return
	}
	fmt.Println(len(decodeString))
	fmt.Println(decodeString)
}
