package main

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/service/defaultcomponents"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"
)

func components() (component.Factories, error) {
	var errs []error
	factories, err := defaultcomponents.Components()
	if err != nil {
		return component.Factories{}, err
	}

	processors := []component.ProcessorFactory{
	}
	for _, pr := range factories.Processors {
		processors = append(processors, pr)
	}
	exporters := []component.ExporterFactory{
		splunkhecexporter.NewFactory(),
	}
	for _, exporter := range factories.Exporters {
		exporters = append(exporters, exporter)
	}
	factories.Processors, err = component.MakeProcessorFactoryMap(processors...)
	factories.Exporters, err = component.MakeExporterFactoryMap(exporters...)
	if err != nil {
		errs = append(errs, err)
	}

	return factories, componenterror.CombineErrors(errs)
}
