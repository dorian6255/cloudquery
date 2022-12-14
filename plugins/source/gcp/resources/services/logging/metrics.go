// Code generated by codegen; DO NOT EDIT.

package logging

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/logging/v2"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/logging/apiv2"
)

func Metrics() *schema.Table {
	return &schema.Table{
		Name:        "gcp_logging_metrics",
		Description: `https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics#LogMetric`,
		Resolver:    fetchMetrics,
		Multiplex:   client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "filter",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Filter"),
			},
			{
				Name:     "disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Disabled"),
			},
			{
				Name:     "metric_descriptor",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetricDescriptor"),
			},
			{
				Name:     "value_extractor",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ValueExtractor"),
			},
			{
				Name:     "label_extractors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LabelExtractors"),
			},
			{
				Name:     "bucket_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BucketOptions"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("Version"),
			},
		},
	}
}

func fetchMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListLogMetricsRequest{
		Parent: "projects/" + c.ProjectId,
	}
	gcpClient, err := logging.NewMetricsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListLogMetrics(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
