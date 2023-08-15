package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

/*
# 是否支持聚类
enableProfilo: true
# 是否开启戴口罩聚类
enableRespiratorCluster: false
# 是否开启人体聚类。
enablePedestrianCluster: true
profilo:
# 是否开启实时聚档功能
enableRealtimeFaceCluster: true
iot:
# iot gateway 云翼地址.
# e.g. 1.2.3.4:8500/admin:admin
existingGatewayAddr: "1.2.3.4:8500/admin:admin"

*/

type CoreProfilo struct {
	// 啊啊啊啊啊啊啊啊啊啊啊啊啊
	EnableRealtimeFaceCluster bool `yaml:"enableRealtimeFaceCluster"` // aaaaaaaaa
}
type IOT struct {
	ExistingGatewayAddr string `yaml:"existingGatewayAddr"`
}

type Cube struct {
	EnableProfilo           bool        `yaml:"enableProfilo"`
	EnableRespiratorCluster bool        `yaml:"enableRespiratorCluster"`
	EnablePedestrianCluster bool        `yaml:"enablePedestrianCluster"`
	Profilo                 CoreProfilo `yaml:"profilo"`
	IOT                     IOT         `yaml:"iot"`
}

func main() {
	wanger := &Cube{
		EnableProfilo:           true,
		EnableRespiratorCluster: true,
		EnablePedestrianCluster: true,
		Profilo:                 CoreProfilo{EnableRealtimeFaceCluster: true},
		IOT:                     IOT{ExistingGatewayAddr: fmt.Sprintf("%s:8500/admin:admin", "9.9.9.9")},
	}

	yamlData, err := yaml.Marshal(wanger)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fileName := "cube.yaml"
	err = ioutil.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		return
	}
}
