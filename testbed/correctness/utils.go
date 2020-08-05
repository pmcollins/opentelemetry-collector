package correctness

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.opentelemetry.io/collector/testbed/testbed"
)

func CreateConfigYaml(
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	processors map[string]string,
	pipelineType string,
) string {

	// Prepare extra processor config section and comma-separated list of extra processor
	// names to use in corresponding "processors" settings.
	processorsSections := ""
	processorsList := ""
	if len(processors) > 0 {
		first := true
		for name, cfg := range processors {
			processorsSections += cfg + "\n"
			if !first {
				processorsList += ","
			}
			processorsList += name
			first = false
		}
	}

	format := `
receivers:%v
exporters:%v
processors:
  %s

extensions:

service:
  extensions:
  pipelines:
    %s:
      receivers: [%v]
      processors: [%s]
      exporters: [%v]
`

	return fmt.Sprintf(
		format,
		sender.GenConfigYAMLStr(),
		receiver.GenConfigYAMLStr(),
		processorsSections,
		pipelineType,
		sender.ProtocolName(),
		processorsList,
		receiver.ProtocolName(),
	)
}

type PipelineDef struct {
	Receiver     string
	Exporter     string
	TestName     string
	DataSender   testbed.DataSender
	DataReceiver testbed.DataReceiver
	ResourceSpec testbed.ResourceSpec
}

func LoadPictOutputPipelineDefs(fileName string) ([]PipelineDef, error) {
	file, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	defs := make([]PipelineDef, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\t")
		if "Receiver" == s[0] {
			continue
		}

		var aDef PipelineDef
		aDef.Receiver, aDef.Exporter = s[0], s[1]
		defs = append(defs, aDef)
	}

	return defs, err
}

func ConstructTraceSender(t *testing.T, receiver string) testbed.DataSender {
	var sender testbed.DataSender
	switch receiver {
	case "otlp":
		sender = testbed.NewOTLPTraceDataSender(testbed.DefaultHost, testbed.GetAvailablePort(t))
	case "opencensus":
		sender = testbed.NewOCTraceDataSender(testbed.DefaultHost, testbed.GetAvailablePort(t))
	case "jaeger":
		sender = testbed.NewJaegerGRPCDataSender(testbed.DefaultHost, testbed.GetAvailablePort(t))
	case "zipkin":
		sender = testbed.NewZipkinDataSender(testbed.DefaultHost, testbed.GetAvailablePort(t))
	default:
		t.Errorf("unknown receiver type: %s", receiver)
	}
	return sender
}

func ConstructMetricsSender(t *testing.T, receiver string) testbed.DataSender {
	var sender testbed.DataSender
	switch receiver {
	case "otlp":
		sender = testbed.NewOTLPMetricDataSender(testbed.DefaultHost, testbed.GetAvailablePort(t))
	case "opencensus":
		sender = testbed.NewOCMetricDataSender(testbed.DefaultHost, testbed.GetAvailablePort(t))
	// will be uncommented in a subsequent PR
	// case "prometheus":
	// 	sender = testbed.NewPrometheusDataSender(testbed.DefaultHost, testbed.GetAvailablePort(t))
	default:
		t.Errorf("unknown receiver type: %s", receiver)
	}
	return sender
}

func ConstructReceiver(t *testing.T, exporter string) testbed.DataReceiver {
	var receiver testbed.DataReceiver
	switch exporter {
	case "otlp":
		receiver = testbed.NewOTLPDataReceiver(testbed.GetAvailablePort(t))
	case "opencensus":
		receiver = testbed.NewOCDataReceiver(testbed.GetAvailablePort(t))
	case "jaeger":
		receiver = testbed.NewJaegerDataReceiver(testbed.GetAvailablePort(t))
	case "zipkin":
		receiver = testbed.NewZipkinDataReceiver(testbed.GetAvailablePort(t))
	// will be uncommented in a subsequent PR
	// case "prometheus":
	// 	receiver = testbed.NewPrometheusDataReceiver(testbed.GetAvailablePort(t))
	default:
		t.Errorf("unknown exporter type: %s", exporter)
	}
	return receiver
}
