package filters

import (
	"io/ioutil"

	"github.com/Knetic/govaluate"
	"github.com/logicmonitor/k8s-argus/pkg/constants"
	"github.com/logicmonitor/lm-sdk-go/models"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	filter filters
)

type filters struct {
	filtersConfig config
}

type config struct {
	Filter filterExpression `yaml:"filter"`
}

type filterExpression struct {
	POD        string `yaml:"pods"`
	SERVICES   string `yaml:"services"`
	DEPLOYMENT string `yaml:"deployments"`
	NODE       string `yaml:"nodes"`
}

// package init block so that filter-config will be loaded on application start
func init() {
	filter = filters{}
	filter.setConfig(readFilterConfig())
}

func (f *filters) setConfig(config *config) {
	f.filtersConfig = *config
}

func (config config) get(resource string) filterExpression {
	switch resource {
	case "filter":
		return config.Filter
	}
	return filterExpression{}
}

func (expression filterExpression) get(resource string) string {
	switch resource {
	case constants.Pods:
		return expression.POD
	case constants.Deployments:
		return expression.DEPLOYMENT
	case constants.Services:
		return expression.SERVICES
	case constants.Nodes:
		return expression.NODE
	}
	return ""
}

func readFilterConfig() *config {
	configBytes, err := ioutil.ReadFile("/etc/argus/filters-config.yaml")
	if err != nil {
		log.Fatalf("Failed to read filters config file: /etc/argus/filters-config.yaml")
	}
	config := &config{}
	log.Infof("config bytes %s ", configBytes)
	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		log.Fatalf("Couldn't parse filters-config file.")
	}
	log.Infof("Filter config read: %v", config)
	return config
}

// getEvaluationParamsForResource generates evaluation parameters based on labels and specified resource
func getEvaluationParamsForResource(device *models.Device, labels map[string]string) (map[string]interface{}, error) {
	evaluationParams := make(map[string]interface{}, 8)

	for key, value := range labels {
		evaluationParams[key] = value
	}

	evaluationParams["name"] = *device.DisplayName
	return evaluationParams, nil
}

func getFilterExpressionForResource(resource string) string {
	return filter.filtersConfig.get("filter").get(resource)
}

// EvaluateFiltering evaluates filtering expression based on labels and specified resource
func EvaluateFiltering(resource string, device *models.Device, labels map[string]string) bool {
	filterExpression := getFilterExpressionForResource(resource)
	log.Infof("Filter exprrssion for resource %s is %s", resource, filterExpression)

	if len(filterExpression) == 0 {
		log.Infof("No filtering specified for resouce %s ", resource)
		return false
	}

	if isFilterAll(filterExpression) {
		return true
	}

	evaluationParams, _ := getEvaluationParamsForResource(device, labels)
	log.Infof("Evaluation params  %+v:", evaluationParams)

	expression, err := govaluate.NewEvaluableExpression(filterExpression)

	result, err := expression.Evaluate(evaluationParams)

	if err != nil {
		log.Errorf("Invalid filter expression for resource %s -> %s", resource, filterExpression)
		return false
	}

	if result.(bool) {
		log.Infof("Condition is evaluated to true")
		return true
	}
	return false

}

func isFilterAll(expression string) bool {
	return expression == "*"
}
